---
title: 监控和日志（Monitoring & Logging）
outline: deep
---

# 监控和日志

::: tip
**监控和日志**是生产环境的"眼睛"和"记录仪"。它们让你了解应用运行状态，快速定位问题，并预防潜在故障。
:::

## 为什么需要监控？

想象你开车时没有仪表盘：不知道速度、油量、水温。这样开车既危险又低效。生产环境的应用也是如此，没有监控就是"盲驾"。

### 监控的核心价值

**故障快速发现**：系统出现问题时，监控能在用户发现之前就告警。

**性能趋势分析**：了解系统负载变化，提前扩容或优化。

**用户体验保障**：确保响应时间、可用性等关键指标在合理范围内。

**容量规划**：基于历史数据预测未来的资源需求。

---

## 监控体系概述

现代监控通常包含四个维度，称为"四个黄金信号"：

### 延迟（Latency）
请求的响应时间。例如：
- API 响应时间：95% 的请求在 200ms 内响应
- 数据库查询时间：平均查询时间 50ms

### 流量（Traffic）
系统处理的请求量。例如：
- 每秒 HTTP 请求数（RPS）
- 每分钟处理的消息数

### 错误率（Errors）
失败请求的比例。例如：
- HTTP 5xx 错误率 < 0.1%
- 任务失败率 < 1%

### 饱和度（Saturation）
资源使用率。例如：
- CPU 使用率 < 80%
- 内存使用率 < 90%
- 磁盘使用率 < 85%

---

## Go 应用监控实现

### HTTP 服务监控

让我们从一个简单的 Web 服务开始：

::: details 示例：HTTP 服务监控
```go
package main

import (
    "log"
    "net/http"
    "strconv"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// 定义指标
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration in seconds",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    // 注册指标
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// 监控中间件
func monitoringMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 包装 ResponseWriter 以捕获状态码
        wrapper := &responseWrapper{ResponseWriter: w, statusCode: 200}
        
        next(wrapper, r)
        
        duration := time.Since(start).Seconds()
        status := strconv.Itoa(wrapper.statusCode)
        
        // 记录指标
        httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
        httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
    }
}

type responseWrapper struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWrapper) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```
:::

### 健康检查端点
::: details 示例：健康检查端点
```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    // 检查数据库连接
    if !isDatabaseHealthy() {
        http.Error(w, "Database unhealthy", http.StatusServiceUnavailable)
        return
    }
    
    // 检查外部依赖
    if !isExternalServiceHealthy() {
        http.Error(w, "External service unhealthy", http.StatusServiceUnavailable)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func isDatabaseHealthy() bool {
    // 简单的数据库连接检查
    // return db.Ping() == nil
    return true
}

func isExternalServiceHealthy() bool {
    // 检查外部 API 是否可达
    client := http.Client{Timeout: 2 * time.Second}
    _, err := client.Get("https://api.external-service.com/health")
    return err == nil
}
```
:::
---

## Prometheus 生态系统

**Prometheus** 是云原生监控的事实标准，特别适合 Go 应用。

### 为什么选择 Prometheus？

**Pull 模式**：Prometheus 主动拉取指标，服务只需要暴露 HTTP 端点。

**时间序列数据库**：专门为监控数据设计，高效存储和查询。

**强大的查询语言**：PromQL 可以进行复杂的数据分析。

**丰富的生态**：与 Grafana、Alertmanager 等工具完美集成。

### 基础配置

::: details Prometheus 配置示例
```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'my-go-app'
    static_configs:
      - targets: ['localhost:8080']
    metrics_path: '/metrics'
    scrape_interval: 10s

  - job_name: 'node-exporter'
    static_configs:
      - targets: ['localhost:9100']

rule_files:
  - "alert_rules.yml"

alerting:
  alertmanagers:
    - static_configs:
        - targets: ['localhost:9093']
```
:::

### 关键指标设计

**计数器（Counter）**：只增不减的累积指标，如请求总数、错误总数。

```go
httpRequestsTotal.WithLabelValues("GET", "/api/users", "200").Inc()
```

**仪表盘（Gauge）**：可增可减的瞬时值，如当前连接数、内存使用量。

```go
var currentConnections = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "current_connections",
    Help: "Current number of connections",
})
```

**直方图（Histogram）**：观察值的分布，如响应时间分布。

```go
httpRequestDuration.WithLabelValues("GET", "/api/users").Observe(0.125)
```

**摘要（Summary）**：类似直方图，但提供分位数统计。

---

## 日志系统设计

### 结构化日志

传统的文本日志难以解析和查询，结构化日志（JSON 格式）更适合现代应用：

::: details 示例：结构化日志
```go
package main

import (
    "context"
    "github.com/sirupsen/logrus"
    "net/http"
)

func setupLogger() *logrus.Logger {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
    
    // 生产环境只输出 Info 及以上级别
    logger.SetLevel(logrus.InfoLevel)
    
    return logger
}

func logRequestMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            
            // 生成请求 ID
            requestID := generateRequestID()
            ctx := context.WithValue(r.Context(), "requestID", requestID)
            r = r.WithContext(ctx)
            
            wrapper := &responseWrapper{ResponseWriter: w, statusCode: 200}
            next.ServeHTTP(wrapper, r)
            
            logger.WithFields(logrus.Fields{
                "request_id":    requestID,
                "method":        r.Method,
                "path":          r.URL.Path,
                "status":        wrapper.statusCode,
                "duration_ms":   time.Since(start).Milliseconds(),
                "user_agent":    r.UserAgent(),
                "remote_addr":   r.RemoteAddr,
            }).Info("HTTP request processed")
        })
    }
}
```
:::
### 日志级别策略

**ERROR**：需要立即关注的错误，可能影响服务可用性
- 数据库连接失败
- 外部 API 调用失败
- 业务逻辑错误

**WARN**：需要关注但不会立即影响服务的问题
- 配置使用默认值
- 性能降级
- 资源使用率较高

**INFO**：正常的业务流程记录
- 用户登录/登出
- 重要操作完成
- 服务启动/停止

**DEBUG**：详细的调试信息，通常只在开发环境启用
- 函数入参/出参
- 中间状态变量
- 详细的执行流程

---

## 分布式链路追踪

在微服务架构中，一个用户请求可能经过多个服务。**分布式链路追踪**能够跟踪请求在整个系统中的流转过程。

### OpenTelemetry 实现

::: details 示例：分布式链路追踪
```go
package main

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("my-service")

func processOrder(ctx context.Context, orderID string) error {
    // 创建 span
    ctx, span := tracer.Start(ctx, "process_order")
    defer span.End()
    
    // 添加属性
    span.SetAttributes(
        attribute.String("order.id", orderID),
        attribute.String("user.id", getUserID(ctx)),
    )
    
    // 验证订单
    if err := validateOrder(ctx, orderID); err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, "Order validation failed")
        return err
    }
    
    // 处理支付
    if err := processPayment(ctx, orderID); err != nil {
        span.RecordError(err)
        return err
    }
    
    span.SetStatus(codes.Ok, "Order processed successfully")
    return nil
}

func validateOrder(ctx context.Context, orderID string) error {
    ctx, span := tracer.Start(ctx, "validate_order")
    defer span.End()
    
    // 验证逻辑
    return nil
}
```
:::
### 追踪的价值

**性能瓶颈识别**：看到每个服务的处理时间，找出最慢的环节。

**依赖关系可视化**：了解服务间的调用关系和依赖程度。

**错误传播分析**：当出现错误时，可以看到错误是从哪个服务开始传播的。

---

## 监控工具生态

### Grafana 可视化

Grafana 是最流行的监控数据可视化工具：

**仪表板设计**：
- 系统概览：CPU、内存、网络、磁盘
- 应用性能：响应时间、吞吐量、错误率
- 业务指标：用户数、订单量、收入

**告警规则**：
- 响应时间超过 1 秒
- 错误率超过 5%
- CPU 使用率超过 90%

### ELK/EFK 日志栈

**Elasticsearch**：日志存储和搜索引擎
**Logstash/Fluentd**：日志收集和处理
**Kibana**：日志查询和可视化

### 日志收集流程

```
应用 → Filebeat → Logstash → Elasticsearch → Kibana
```

1. **应用**：输出结构化日志到文件或标准输出
2. **Filebeat**：轻量级日志收集器，实时读取日志文件
3. **Logstash**：日志处理管道，解析、转换、丰富日志数据
4. **Elasticsearch**：存储和索引日志数据
5. **Kibana**：提供查询界面和可视化

---

## 实际案例：电商系统监控

假设我们有一个电商系统，包含以下服务：
- 用户服务（User Service）
- 商品服务（Product Service）
- 订单服务（Order Service）
- 支付服务（Payment Service）

### 关键监控指标

**业务指标**：
- 每分钟订单数
- 支付成功率
- 商品浏览量
- 用户注册数

**技术指标**：
- API 响应时间（P95, P99）
- 数据库连接池使用率
- 缓存命中率
- 消息队列积压数量

### 告警策略

::: details 告警规则配置
```yaml
# alert_rules.yml
groups:
- name: ecommerce_alerts
  rules:
  # 业务告警
  - alert: HighOrderFailureRate
    expr: sum(rate(orders_failed_total[5m])) / sum(rate(orders_total[5m])) > 0.1
    for: 5m
    labels:
      severity: critical
    annotations:
      summary: "订单失败率过高"
      description: "订单失败率 {{ $value | humanizePercentage }} 超过 10%"

  # 性能告警
  - alert: HighResponseTime
    expr: histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le)) > 1
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "响应时间过长"
      description: "95% 响应时间 {{ $value }}s 超过 1 秒"

  # 资源告警
  - alert: HighMemoryUsage
    expr: process_memory_usage_ratio > 0.9
    for: 10m
    labels:
      severity: warning
    annotations:
      summary: "内存使用率过高"
      description: "内存使用率 {{ $value | humanizePercentage }} 超过 90%"
```
:::

### 故障处理流程

1. **告警触发**：监控系统发现异常，发送告警到 Slack/Email
2. **问题定位**：通过 Grafana 仪表板查看系统状态
3. **日志分析**：在 Kibana 中搜索相关错误日志
4. **链路追踪**：使用 Jaeger 分析请求调用链
5. **问题修复**：定位根因并修复
6. **验证恢复**：确认指标恢复正常

---

## 性能监控最佳实践

### Go 应用特有监控

**Goroutine 数量**：
::: details 示例：Goroutine 数量
```go
var goroutineCount = prometheus.NewGaugeFunc(
    prometheus.GaugeOpts{
        Name: "goroutines_count",
        Help: "Number of goroutines",
    },
    func() float64 {
        return float64(runtime.NumGoroutine())
    },
)
```
:::
**GC 性能**：
::: details 示例：GC 性能
```go
var gcDuration = prometheus.NewGaugeFunc(
    prometheus.GaugeOpts{
        Name: "gc_duration_seconds",
        Help: "GC pause duration",
    },
    func() float64 {
        var stats runtime.MemStats
        runtime.ReadMemStats(&stats)
        return float64(stats.PauseNs[(stats.NumGC+255)%256]) / 1e9
    },
)
```
:::
**内存使用**：
::: details 示例：内存使用
```go
var memoryUsage = prometheus.NewGaugeFunc(
    prometheus.GaugeOpts{
        Name: "memory_usage_bytes",
        Help: "Memory usage in bytes",
    },
    func() float64 {
        var stats runtime.MemStats
        runtime.ReadMemStats(&stats)
        return float64(stats.Alloc)
    },
)
```
:::
### 监控数据保留策略

**实时数据**（15秒间隔）：保留 1 天
**短期数据**（1分钟间隔）：保留 7 天
**中期数据**（5分钟间隔）：保留 30 天
**长期数据**（1小时间隔）：保留 1 年

---

## 常见监控陷阱

### 过度监控

**问题**：监控太多指标，产生噪音，掩盖真正重要的信号。

**解决方案**：
- 专注于影响用户体验的关键指标
- 使用 SLI/SLO（服务等级指标/目标）指导监控策略
- 定期审查和清理不必要的告警

### 告警疲劳

**问题**：告警太多或者误报率高，导致重要告警被忽视。

**解决方案**：
- 设置合理的告警阈值
- 使用告警分级（Critical/Warning/Info）
- 实施告警抑制和静默机制

### 监控盲区

**问题**：某些重要组件或流程没有被监控覆盖。

**解决方案**：
- 绘制系统架构图，确保每个组件都有监控
- 监控业务流程，不仅仅是技术指标
- 定期进行监控有效性测试

---

## 成本优化

### 存储成本控制

**日志分层存储**：
- 热数据（最近 7 天）：SSD 存储，快速检索
- 温数据（8-30 天）：标准存储
- 冷数据（30 天以上）：归档存储

**指标数据采样**：
- 高频指标：根据重要性调整采集频率
- 低价值指标：降低精度或删除

### 云服务成本优化

**按需扩缩容**：根据负载自动调整监控基础设施规模

**预留实例**：对于稳定的监控工作负载，使用预留实例降低成本

**多云策略**：比较不同云提供商的监控服务价格

---

## 💡 监控体系建设路径

1. **起步阶段**：基础监控（CPU、内存、磁盘、网络）
2. **发展阶段**：应用监控（响应时间、错误率、吞吐量）
3. **成熟阶段**：业务监控（用户行为、转化率、收入）
4. **高级阶段**：智能监控（异常检测、预测分析）

监控不是一蹴而就的，需要随着系统复杂度的增加而逐步完善。重要的是先建立基础监控，确保系统可观测，然后再逐步细化和优化。

📊 接下来推荐阅读：[性能优化](/practice/deployment/performance)，学习如何基于监控数据进行系统性能调优。
