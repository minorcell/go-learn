---
title: "黑夜中的哨兵：Go 应用监控与告警"
description: "应用部署到生产环境，如同孤舟驶入黑夜。本手册将指导你如何部署“哨兵”——建立基于指标、日志和追踪的观测体系，使用 Prometheus 和 Grafana，在问题发生前洞察一切。"
---

# 黑夜中的哨兵：Go 应用监控与告警

当你的 Go 应用被部署到生产环境，它就进入了一片深邃的黑夜。开发者无法再直接触及它，也无法预知它会遇到何种风暴。没有监控的应用，就像一艘没有雷达和瞭望员的船，任何一点小问题都可能演变成一场灾难。

本手册旨在教你如何部署忠诚的"哨兵"——一个强大的监控与告警系统。我们将围绕**可观测性 (Observability) 的三大支柱**来构建这个系统，它们是你洞察应用内部状态的眼睛和耳朵。

## 1. 可观测性的三大支柱

现代监控系统早已超越了简单的"CPU/内存使用率"检查。一个真正可观测的系统，建立在三个核心数据源之上：

- **指标 (Metrics)**: **可聚合的**数值型数据。它们是你的仪表盘，告诉你应用的宏观状态，如"当前 QPS 是多少？"或"95分位延迟是多少？"。指标非常高效，适合长期存储和趋势分析。

- **日志 (Logs)**: **离散的**、带有时间戳的事件记录。当指标显示"有异常"时，日志会告诉你"发生了什么具体事件"。结构化的日志（如 JSON 格式）是现代日志系统的基石。

- **追踪 (Traces)**: **以请求为维度**的事件链。它能串联起一个请求在分布式系统中所经过的所有服务和组件，清晰地展示出延迟发生在哪个环节。追踪是诊断微服务架构中性能问题的终极武器。

这三大支柱相辅相成，共同构成了一幅完整的应用健康全景图。

## 2. 核心武器：Prometheus + Grafana

- **Prometheus**: 云原生监控领域的王者。它是一个时间序列数据库和监控系统，通过"拉模式 (Pull)"从你的应用暴露的 `/metrics` 端点上抓取指标。
- **Grafana**: 领先的可视化平台。它能连接到 Prometheus 等多种数据源，通过灵活的查询和丰富的图表，将枯燥的数据变为直观的仪表盘。

## 3. 实战：武装你的 Go 应用

我们将创建一个简单的 HTTP 服务，并为其配备完整的观测能力。

### 3.1. 应用代码 (`main.go`)

这个例子将演示如何在一个文件中同时实现三大支柱的埋点。

```go
package main

import (
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// 指标 (Metrics)
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	// 日志 (Logging)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("GET /hello", loggingMiddleware(helloHandler))

	slog.Info("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server failed to start", "error", err)
	}
}

// HTTP 中间件，用于统一处理监控埋点
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 追踪 (Tracing) - 生成一个唯一的请求 ID
		traceID := uuid.New().String()
		ctxLogger := slog.With("trace_id", traceID)

		ctxLogger.Info("request started")
		start := time.Now()

		// 包装 ResponseWriter 以捕获状态码
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(rw, r)

		duration := time.Since(start)
		status := http.StatusText(rw.statusCode)

		// 指标记录
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
		httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration.Seconds())

		ctxLogger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration_ms", duration.Milliseconds(),
		)
	}
}

// 业务处理器
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 模拟随机延迟和随机错误
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	if rand.Intn(10) < 2 { // 20% 的几率出错
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

// responseWriter 包装器
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
```

### 3.2. 运行完整的监控栈 (`docker-compose.yml`)

为了方便本地测试，我们使用 Docker Compose 将应用、Prometheus 和 Grafana 一起运行。

```yaml
version: '3.8'

services:
  # 我们的 Go 应用
  app:
    build: .
    ports:
      - "8080:8080"
    restart: always

  # Prometheus 用于收集指标
  prometheus:
    image: prom/prometheus:v2.53.0
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
      - ./alert.rules.yml:/etc/prometheus/alert.rules.yml
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
    restart: always
    depends_on:
      - app

  # Grafana 用于可视化
  grafana:
    image: grafana/grafana:11.1.0
    ports:
      - "3000:3000"
    restart: always
    depends_on:
      - prometheus
```

### 3.3. Prometheus 配置

**`prometheus.yml`**
```yaml
global:
  scrape_interval: 15s

rule_files:
  - "/etc/prometheus/alert.rules.yml"

scrape_configs:
  - job_name: 'go-app'
    static_configs:
      - targets: ['app:8080'] # 'app' 是 docker-compose 中的服务名
```

**`alert.rules.yml`**
```yaml
groups:
  - name: go-app-alerts
    rules:
      - alert: HighErrorRate
        expr: (sum(rate(http_requests_total{status=~"5.."}[5m])) by (job) / sum(rate(http_requests_total[5m])) by (job)) > 0.1
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "High HTTP 5xx Error Rate on {{ $labels.job }}"
          description: "The error rate is over 10% for the last 5 minutes."

      - alert: HighRequestLatency
        expr: histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le, job)) > 0.5
        for: 2m
        labels:
          severity: warning
        annotations:
          summary: "High P95 Latency on {{ $labels.job }}"
          description: "The 95th percentile latency is over 500ms for the last 5 minutes."
```

## 4. 建立你的指挥中心

1.  **启动**: 在包含 `main.go`, `Dockerfile`, `docker-compose.yml` 等文件的目录下，运行 `docker-compose up`。
2.  **生成流量**: 使用工具（如 `hey` 或 `wrk`）向 `http://localhost:8080/hello` 发送一些请求。
3.  **查看指标**: 访问 `http://localhost:9090`，在 Prometheus UI 中查询 `http_requests_total`。
4.  **配置 Grafana**: 访问 `http://localhost:3000` (默认用户名/密码: `admin`/`admin`)。
    - 添加 Prometheus 数据源 (URL: `http://prometheus:9090`)。
    - 创建一个新的仪表盘，添加图表来可视化 QPS ( `sum(rate(http_requests_total[1m]))` ) 和 P95 延迟。
5.  **观察日志**: 在 `docker-compose up` 的终端输出中，你将看到带有 `trace_id` 的 JSON 格式日志。

## 5. 设计有效的告警

告警不是越多越好。**嘈杂的、无法采取行动的告警比没有告警更糟糕**。好的告警应该直接关联到用户体验。

- **告警于症状，而非原因**: 优先告警于高错误率、高延迟等直接影响用户的"症状"，而不是 CPU 使用率过高等"原因"。后者应该是排查问题时参考的指标。
- **使用 `for` 关键字**: 在 `alert.rules.yml` 中，`for: 1m` 表示只有当条件持续满足 1 分钟后才触发告警，这能有效防止因短暂尖峰而产生的"抖动"告警。
- **提供明确的行动指南**: 告警信息中应包含足够的信息（如服务名、问题描述），甚至可以链接到处理预案 (Runbook)。

通过部署这些"哨兵"，你就为你的应用建立了一套强大的防御体系。现在，即使在最黑暗的生产环境之夜，你也能洞察一切，从容应对。
