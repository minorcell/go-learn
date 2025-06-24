# 构建和部署工具 Build & Deploy

> 从代码到产品的最后一公里——构建和部署是将你的Go代码变成可运行服务的关键环节

## 🤔 为什么构建和部署如此重要？

很多开发者认为写完代码就完成了工作，但实际上，**代码只有部署到生产环境才能创造价值**。Go语言在构建和部署方面有着天然的优势，但如何充分利用这些优势，需要深入理解整个构建部署流程。

### Go构建部署的独特优势

#### 🎯 单一可执行文件

::: details 示例：Go编译后就是一个独立的可执行文件
```bash
# Go编译后就是一个独立的可执行文件
go build -o myapp main.go

# 无需安装运行时环境，直接运行
./myapp
```
:::
这种设计哲学的深层含义：
- **部署简单**：不需要复杂的依赖管理
- **容器友好**：完美适配Docker容器化
- **跨平台**：一次编写，到处运行
- **启动快速**：无需虚拟机预热

#### ⚡ 交叉编译能力

::: details 示例：交叉编译能力
```bash
# 在Linux上为Windows编译
GOOS=windows GOARCH=amd64 go build -o myapp.exe main.go

# 在macOS上为Linux编译
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go

# 为ARM架构编译（如树莓派）
GOOS=linux GOARCH=arm go build -o myapp-arm main.go
```
:::
## 📊 构建部署工具全景

```mermaid
graph TD
    A[Go构建部署] --> B[本地构建]
    A --> C[容器化]
    A --> D[CI/CD]
    A --> E[部署策略]
    A --> F[监控运维]
    
    B --> B1[go build]
    B --> B2[Makefile]
    B --> B3[构建优化]
    
    C --> C1[Docker基础]
    C --> C2[多阶段构建]
    C --> C3[镜像优化]
    
    D --> D1[GitHub Actions]
    D --> D2[GitLab CI]
    D --> D3[Jenkins]
    
    E --> E1[蓝绿部署]
    E --> E2[滚动更新]
    E --> E3[金丝雀发布]
    
    F --> F1[健康检查]
    F --> F2[日志管理]
    F --> F3[性能监控]
```

## 🔧 本地构建实践

### Go Build 深度使用

#### 基础构建命令

::: details 示例：基础构建命令
```bash
# 最简单的构建
go build

# 指定输出文件名
go build -o myapp

# 构建并安装到 $GOPATH/bin
go install

# 显示构建过程
go build -v

# 构建时显示编译器命令
go build -x
```
:::
#### 构建标签（Build Tags）

::: details 示例：构建标签（dev）

```go
// +build dev

package config

// 开发环境配置
const (
    DBHost = "localhost:5432"
    Debug  = true
)
```
:::

::: details 示例：构建标签（prod）
```go
// +build prod

package config

// 生产环境配置
const (
    DBHost = "prod-db.example.com:5432"
    Debug  = false
)
```
:::

::: details 示例：构建标签（dev、prod）
```bash
# 使用构建标签
go build -tags dev      # 开发环境
go build -tags prod     # 生产环境
go build -tags "prod monitoring"  # 多个标签
```
:::
#### 编译优化选项

::: details 示例：编译优化选项
```bash
# 去除调试信息，减小文件大小
go build -ldflags="-s -w" -o myapp

# 静态链接（适合容器部署）
CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp

# 嵌入版本信息
VERSION=$(git describe --tags --always)
BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')
go build -ldflags="-X main.version=${VERSION} -X 'main.buildTime=${BUILD_TIME}'" -o myapp
```
:::
#### 版本信息嵌入

::: details 示例：版本信息嵌入
```go
package main

import (
    "fmt"
    "flag"
)

var (
    version   = "unknown"
    buildTime = "unknown"
    gitCommit = "unknown"
)

func main() {
    var showVersion = flag.Bool("version", false, "Show version information")
    flag.Parse()
    
    if *showVersion {
        fmt.Printf("Version: %s\n", version)
        fmt.Printf("Build Time: %s\n", buildTime)
        fmt.Printf("Git Commit: %s\n", gitCommit)
        return
    }
    
    // 应用逻辑
    fmt.Println("Application is running...")
}
```
:::
### Makefile 构建自动化

#### 基础Makefile

::: details 示例：基础Makefile
```makefile
# Makefile
.PHONY: build clean test coverage help

# 变量定义
APP_NAME := myapp
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S UTC')
GIT_COMMIT := $(shell git rev-parse HEAD)

# 构建标志
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION) -X 'main.buildTime=$(BUILD_TIME)' -X main.gitCommit=$(GIT_COMMIT)"

# 默认目标
all: build

# 构建目标
build:
	@echo "Building $(APP_NAME)..."
	CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APP_NAME) .

# 开发环境构建
build-dev:
	@echo "Building $(APP_NAME) for development..."
	go build -tags dev -o bin/$(APP_NAME)-dev .

# 生产环境构建
build-prod:
	@echo "Building $(APP_NAME) for production..."
	CGO_ENABLED=0 go build -tags prod $(LDFLAGS) -o bin/$(APP_NAME) .

# 交叉编译
build-linux:
	@echo "Building $(APP_NAME) for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux .

build-windows:
	@echo "Building $(APP_NAME) for Windows..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APP_NAME).exe .

build-darwin:
	@echo "Building $(APP_NAME) for macOS..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APP_NAME)-darwin .

# 构建所有平台
build-all: build-linux build-windows build-darwin

# 测试
test:
	@echo "Running tests..."
	go test -v ./...

# 测试覆盖率
coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 代码检查
lint:
	@echo "Running linter..."
	golangci-lint run

# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

# 清理
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# 运行
run:
	@echo "Running $(APP_NAME)..."
	go run .

# 运行开发版本
run-dev:
	@echo "Running $(APP_NAME) in development mode..."
	go run -tags dev .

# 帮助
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  build-dev  - Build for development"
	@echo "  build-prod - Build for production"
	@echo "  build-all  - Build for all platforms"
	@echo "  test       - Run tests"
	@echo "  coverage   - Run tests with coverage"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  clean      - Clean build artifacts"
	@echo "  run        - Run the application"
	@echo "  run-dev    - Run in development mode"
```
:::
#### 高级Makefile技巧

::: details 示例：高级Makefile技巧
```makefile
# 检查工具是否安装
check-tools:
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	}

# 依赖管理
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

# 更新依赖
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# 安全检查
security:
	@echo "Running security check..."
	gosec ./...

# 完整质量检查
quality: fmt lint test security
	@echo "All quality checks passed!"

# 发布准备
release: clean quality build-all
	@echo "Release artifacts ready in bin/"
```
:::
## 🐳 容器化实践

### Docker基础使用

#### 简单Dockerfile

::: details 示例：简单Dockerfile
```dockerfile
# 简单但不够优化的Dockerfile
FROM golang:1.21

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
```
:::
#### 多阶段构建优化

::: details 示例：多阶段构建优化
```dockerfile
# 多阶段构建 - 推荐方式
# 构建阶段
FROM golang:1.21-alpine AS builder

# 安装构建依赖
RUN apk add --no-cache git ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 复制模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-s -w -extldflags "-static"' \
    -o main .

# 运行阶段
FROM scratch

# 从构建阶段复制必要文件
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/main /main

# 创建非root用户
USER 65534:65534

# 暴露端口
EXPOSE 8080

# 启动应用
ENTRYPOINT ["/main"]
```
:::
#### 进一步优化的Dockerfile

::: details 示例：进一步优化的Dockerfile
```dockerfile
# 高度优化的生产级Dockerfile
FROM golang:1.21-alpine AS builder

# 添加构建参数
ARG VERSION=unknown
ARG BUILD_TIME=unknown
ARG GIT_COMMIT=unknown

# 安装必要工具
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    && update-ca-certificates

# 创建非root用户
RUN adduser -D -g '' appuser

WORKDIR /app

# 利用Docker层缓存，先复制依赖文件
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 复制源代码
COPY . .

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=${VERSION} -X 'main.buildTime=${BUILD_TIME}' -X main.gitCommit=${GIT_COMMIT}" \
    -o main .

# 最终镜像使用distroless
FROM gcr.io/distroless/static:nonroot

# 复制时区信息
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# 复制二进制文件
COPY --from=builder /app/main /main

# 复制配置文件（如果有）
COPY --from=builder /app/config/prod.yaml /config/

# 使用非root用户
USER nonroot:nonroot

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/main", "-health-check"]

EXPOSE 8080
ENTRYPOINT ["/main"]
```
:::
### Docker Compose 本地开发

::: details 示例：Docker Compose 本地开发
```yaml
# docker-compose.yml
version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.dev
      args:
        - VERSION=dev
    ports:
      - "8080:8080"
    environment:
      - ENV=development
      - DB_HOST=postgres
      - REDIS_HOST=redis
    volumes:
      - .:/app
      - /app/vendor
    depends_on:
      - postgres
      - redis
    restart: unless-stopped

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data

  # 开发工具
  adminer:
    image: adminer:latest
    ports:
      - "8081:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
  redis_data:
```
:::
::: details 示例：Dockerfile.dev
```dockerfile
# Dockerfile.dev - 开发环境专用
FROM golang:1.21-alpine

RUN apk add --no-cache git

WORKDIR /app

# 安装开发工具
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码（开发时会被volume覆盖）
COPY . .

# 使用air进行热重载
CMD ["air"]
```
:::
## 🚀 CI/CD 流水线

### GitHub Actions

#### 基础工作流

::: details 示例：基础工作流
```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

env:
  GO_VERSION: 1.21

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: testdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: ${{ env.GO_VERSION }}

    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Download dependencies
      run: go mod download

    - name: Run tests
      run: |
        go test -v -race -coverprofile=coverage.out ./...
        go tool cover -html=coverage.out -o coverage.html

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
        flags: unittests
        name: codecov-umbrella

    - name: Run linting
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

    - name: Security scan
      uses: securecodewarrior/github-action-gosec@master
      with:
        args: './...'

  build:
    needs: test
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        goos: [linux, windows, darwin]
        goarch: [amd64, arm64]
        exclude:
          - goos: windows
            goarch: arm64

    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: ${{ env.GO_VERSION }}

    - name: Get version
      id: version
      run: |
        echo "VERSION=$(git describe --tags --always --dirty)" >> $GITHUB_OUTPUT
        echo "BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')" >> $GITHUB_OUTPUT
        echo "GIT_COMMIT=${{ github.sha }}" >> $GITHUB_OUTPUT

    - name: Build binary
      run: |
        GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} \
        go build -ldflags="-s -w \
          -X main.version=${{ steps.version.outputs.VERSION }} \
          -X 'main.buildTime=${{ steps.version.outputs.BUILD_TIME }}' \
          -X main.gitCommit=${{ steps.version.outputs.GIT_COMMIT }}" \
          -o myapp-${{ matrix.goos }}-${{ matrix.goarch }}${{ matrix.goos == 'windows' && '.exe' || '' }} .

    - name: Upload artifacts
      uses: actions/upload-artifact@v3
      with:
        name: binaries
        path: myapp-*
```
:::
#### Docker构建和发布

::: details 示例：Docker构建和发布
```yaml
# .github/workflows/docker.yml
name: Docker Build and Push

on:
  push:
    branches: [ main ]
    tags: [ 'v*' ]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
    - name: Checkout repository
      uses: actions/checkout@v4

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3

    - name: Log in to Container Registry
      uses: docker/login-action@v3
      with:
        registry: ${{ env.REGISTRY }}
        username: ${{ github.actor }}
        password: ${{ secrets.GITHUB_TOKEN }}

    - name: Extract metadata
      id: meta
      uses: docker/metadata-action@v5
      with:
        images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
        tags: |
          type=ref,event=branch
          type=ref,event=pr
          type=semver,pattern={{version}}
          type=semver,pattern={{major}}.{{minor}}

    - name: Build and push Docker image
      uses: docker/build-push-action@v5
      with:
        context: .
        platforms: linux/amd64,linux/arm64
        push: true
        tags: ${{ steps.meta.outputs.tags }}
        labels: ${{ steps.meta.outputs.labels }}
        build-args: |
          VERSION=${{ github.ref_name }}
          BUILD_TIME=${{ github.event.head_commit.timestamp }}
          GIT_COMMIT=${{ github.sha }}
        cache-from: type=gha
        cache-to: type=gha,mode=max
```
:::
### GitLab CI

::: details 示例：GitLab CI
```yaml
# .gitlab-ci.yml
stages:
  - test
  - build
  - package
  - deploy

variables:
  GO_VERSION: "1.21"
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: "/certs"

# 缓存配置
.go-cache:
  cache:
    key: ${CI_COMMIT_REF_SLUG}
    paths:
      - .go/pkg/mod/

before_script:
  - apt-get update -qq && apt-get install -y -qq git ca-certificates
  - export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
  - export GOPATH=$CI_PROJECT_DIR/.go

# 测试阶段
test:
  stage: test
  image: golang:${GO_VERSION}
  extends: .go-cache
  services:
    - postgres:15
  variables:
    POSTGRES_DB: testdb
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres
    DATABASE_URL: "postgres://postgres:postgres@postgres:5432/testdb?sslmode=disable"
  script:
    - go mod download
    - go vet ./...
    - go test -v -race -coverprofile=coverage.out ./...
    - go tool cover -func=coverage.out
  artifacts:
    reports:
      coverage_report:
        coverage_format: cobertura
        path: coverage.xml
    paths:
      - coverage.out
    expire_in: 1 week
  coverage: '/coverage: \d+\.\d+% of statements/'

# 代码质量检查
code_quality:
  stage: test
  image: golang:${GO_VERSION}
  extends: .go-cache
  script:
    - go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    - golangci-lint run --out-format code-climate > gl-code-quality-report.json
  artifacts:
    reports:
      codequality: gl-code-quality-report.json
    expire_in: 1 week
  allow_failure: true

# 构建阶段
build:
  stage: build
  image: golang:${GO_VERSION}
  extends: .go-cache
  script:
    - export VERSION=$(git describe --tags --always --dirty)
    - export BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')
    - |
      CGO_ENABLED=0 go build -ldflags="-s -w \
        -X main.version=${VERSION} \
        -X 'main.buildTime=${BUILD_TIME}' \
        -X main.gitCommit=${CI_COMMIT_SHA}" \
        -o myapp .
  artifacts:
    paths:
      - myapp
    expire_in: 1 week

# Docker打包
package:
  stage: package
  image: docker:latest
  services:
    - docker:dind
  dependencies:
    - build
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - |
      docker build \
        --build-arg VERSION=$(git describe --tags --always) \
        --build-arg BUILD_TIME="$(date -u '+%Y-%m-%d %H:%M:%S UTC')" \
        --build-arg GIT_COMMIT=${CI_COMMIT_SHA} \
        -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA \
        -t $CI_REGISTRY_IMAGE:latest .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:latest
  only:
    - main

# 部署到staging
deploy_staging:
  stage: deploy
  image: alpine:latest
  dependencies: []
  before_script:
    - apk add --no-cache curl
  script:
    - |
      curl -X POST \
        -H "Authorization: Bearer $DEPLOY_TOKEN" \
        -H "Content-Type: application/json" \
        -d '{"image":"'$CI_REGISTRY_IMAGE:$CI_COMMIT_SHA'","environment":"staging"}' \
        $DEPLOY_WEBHOOK_URL
  environment:
    name: staging
    url: https://staging.example.com
  only:
    - main

# 部署到生产环境
deploy_production:
  stage: deploy
  image: alpine:latest
  dependencies: []
  before_script:
    - apk add --no-cache curl
  script:
    - |
      curl -X POST \
        -H "Authorization: Bearer $DEPLOY_TOKEN" \
        -H "Content-Type: application/json" \
        -d '{"image":"'$CI_REGISTRY_IMAGE:$CI_COMMIT_SHA'","environment":"production"}' \
        $DEPLOY_WEBHOOK_URL
  environment:
    name: production
    url: https://example.com
  when: manual
  only:
    - tags
```
:::
## 🎯 部署策略

### Kubernetes部署

#### 基础部署配置

::: details 示例：基础部署配置
```yaml
# k8s/namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: myapp

---
# k8s/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: myapp-config
  namespace: myapp
data:
  app.yaml: |
    server:
      port: 8080
      host: "0.0.0.0"
    database:
      host: postgres-service
      port: "5432"
      name: myapp
    redis:
      host: redis-service
      port: "6379"

---
# k8s/secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: myapp-secret
  namespace: myapp
type: Opaque
data:
  database-password: cGFzc3dvcmQ=  # base64 encoded
  jwt-secret: c2VjcmV0a2V5  # base64 encoded

---
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  namespace: myapp
  labels:
    app: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_PASSWORD
          valueFrom:
            secretKeyRef:
              name: myapp-secret
              key: database-password
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: myapp-secret
              key: jwt-secret
        volumeMounts:
        - name: config
          mountPath: /config
          readOnly: true
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
      volumes:
      - name: config
        configMap:
          name: myapp-config

---
# k8s/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp-service
  namespace: myapp
spec:
  selector:
    app: myapp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: ClusterIP

---
# k8s/ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: myapp-ingress
  namespace: myapp
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - api.example.com
    secretName: myapp-tls
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myapp-service
            port:
              number: 80
```
:::
#### HPA（水平自动扩展）

::: details 示例：HPA（水平自动扩展）
```yaml
# k8s/hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: myapp-hpa
  namespace: myapp
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: myapp
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```
:::
### 健康检查实现

::: details 示例：健康检查实现
```go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
)

// 健康检查状态
type HealthStatus struct {
    Status    string            `json:"status"`
    Timestamp time.Time         `json:"timestamp"`
    Services  map[string]string `json:"services"`
    Version   string            `json:"version"`
    Uptime    string            `json:"uptime"`
}

type HealthChecker struct {
    startTime time.Time
    version   string
    checks    map[string]func(context.Context) error
}

func NewHealthChecker(version string) *HealthChecker {
    return &HealthChecker{
        startTime: time.Now(),
        version:   version,
        checks:    make(map[string]func(context.Context) error),
    }
}

func (h *HealthChecker) AddCheck(name string, check func(context.Context) error) {
    h.checks[name] = check
}

func (h *HealthChecker) LivenessHandler(w http.ResponseWriter, r *http.Request) {
    // 活性检查：只要进程在运行就返回成功
    status := HealthStatus{
        Status:    "ok",
        Timestamp: time.Now(),
        Version:   h.version,
        Uptime:    time.Since(h.startTime).String(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(status)
}

func (h *HealthChecker) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
    // 就绪检查：检查所有依赖服务
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    services := make(map[string]string)
    allHealthy := true
    
    for name, check := range h.checks {
        if err := check(ctx); err != nil {
            services[name] = "unhealthy: " + err.Error()
            allHealthy = false
        } else {
            services[name] = "healthy"
        }
    }
    
    status := HealthStatus{
        Timestamp: time.Now(),
        Services:  services,
        Version:   h.version,
        Uptime:    time.Since(h.startTime).String(),
    }
    
    if allHealthy {
        status.Status = "ok"
        w.WriteHeader(http.StatusOK)
    } else {
        status.Status = "unhealthy"
        w.WriteHeader(http.StatusServiceUnavailable)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(status)
}

// 数据库健康检查
func (h *HealthChecker) DatabaseCheck(db *sql.DB) func(context.Context) error {
    return func(ctx context.Context) error {
        return db.PingContext(ctx)
    }
}

// Redis健康检查
func (h *HealthChecker) RedisCheck(client *redis.Client) func(context.Context) error {
    return func(ctx context.Context) error {
        return client.Ping(ctx).Err()
    }
}

// 外部API健康检查
func (h *HealthChecker) ExternalAPICheck(url string) func(context.Context) error {
    return func(ctx context.Context) error {
        req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
        if err != nil {
            return err
        }
        
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        
        if resp.StatusCode >= 400 {
            return fmt.Errorf("external API returned status %d", resp.StatusCode)
        }
        
        return nil
    }
}

// 使用示例
func main() {
    // 初始化健康检查器
    healthChecker := NewHealthChecker(version)
    
    // 添加检查项
    healthChecker.AddCheck("database", healthChecker.DatabaseCheck(db))
    healthChecker.AddCheck("redis", healthChecker.RedisCheck(redisClient))
    healthChecker.AddCheck("external-api", healthChecker.ExternalAPICheck("https://api.example.com/health"))
    
    // 注册健康检查端点
    http.HandleFunc("/health", healthChecker.LivenessHandler)
    http.HandleFunc("/ready", healthChecker.ReadinessHandler)
    
    // 启动服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
:::
## 🎯 构建部署最佳实践

### 1. 构建优化清单

#### 二进制文件优化
- [ ] 使用 `-ldflags="-s -w"` 去除调试信息
- [ ] 启用 `CGO_ENABLED=0` 进行静态编译
- [ ] 嵌入版本信息和构建时间
- [ ] 使用构建标签区分环境

#### 容器优化
- [ ] 使用多阶段构建减小镜像大小
- [ ] 使用 distroless 或 scratch 基础镜像
- [ ] 利用 Docker 层缓存
- [ ] 添加健康检查和非root用户

### 2. 安全性考虑
::: details 示例：安全性最佳实践
```dockerfile
# 安全性最佳实践
FROM golang:1.21-alpine AS builder

# 使用非root用户构建
RUN adduser -D -g '' appuser

# 扫描基础镜像漏洞
FROM gcr.io/distroless/static:nonroot

# 复制二进制文件时设置正确的权限
COPY --from=builder --chown=nonroot:nonroot /app/main /main

# 使用非root用户运行
USER nonroot:nonroot

# 只暴露必要的端口
EXPOSE 8080

# 使用 ENTRYPOINT 而不是 CMD
ENTRYPOINT ["/main"]
```
:::
### 3. 监控和可观测性
::: details 示例：监控和可观测性
```go
// 集成Prometheus指标
package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

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
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

func prometheusMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 包装ResponseWriter以捕获状态码
        wrapped := &responseWrapper{ResponseWriter: w, statusCode: 200}
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start).Seconds()
        status := fmt.Sprintf("%d", wrapped.statusCode)
        
        httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
        httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
    })
}

func main() {
    // 注册Prometheus指标端点
    http.Handle("/metrics", promhttp.Handler())
    
    // 应用中间件
    mux := http.NewServeMux()
    mux.HandleFunc("/api/users", usersHandler)
    
    handler := prometheusMiddleware(mux)
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```
:::
---

💡 **构建部署心法**：
1. **自动化优先**：能自动化的就不要手动
2. **安全第一**：永远使用非root用户和最小权限
3. **可观测性**：监控、日志、追踪一个都不能少
4. **渐进式部署**：蓝绿部署、金丝雀发布降低风险
5. **基础设施即代码**：所有配置都应该版本化

**恭喜！**：你已经掌握了Go应用从开发到生产的完整工具链。接下来可以学习[实战项目](/practice/projects/)，将这些工具应用到真实项目中。
