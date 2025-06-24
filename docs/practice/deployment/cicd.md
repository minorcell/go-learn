---
title: CI/CD 流水线（CI/CD Pipeline）
outline: deep
---

# CI/CD 流水线

::: tip
**CI/CD** 是现代软件开发的基石。通过自动化测试、构建、部署流程，实现快速、可靠的软件交付，让开发者专注于代码本身。
:::

## 什么是 CI/CD？

### 持续集成（CI - Continuous Integration）
开发者频繁地将代码变更合并到主分支，每次合并都会触发自动化构建和测试。这样可以：
- **早期发现问题**：问题在提交后几分钟内就能发现
- **减少合并冲突**：小步快跑，避免大规模代码冲突
- **提高代码质量**：每次提交都经过测试验证

### 持续部署（CD - Continuous Deployment）
在 CI 基础上，自动将通过测试的代码部署到生产环境。包括：
- **自动化部署**：无需人工干预的部署流程
- **环境一致性**：开发、测试、生产环境配置一致
- **快速回滚**：出现问题时能快速恢复到上一个版本

### 现实中的价值

想象一个电商网站的发布场景：
- **传统方式**：每周五晚上手动部署，团队加班到深夜，Monday 才知道是否有问题
- **CI/CD 方式**：代码提交后 10 分钟自动部署到预发布环境，通过测试后自动发布到生产环境

---

## Go 项目的 CI/CD 特点

Go 项目在 CI/CD 中有独特优势：

### 编译时检查
Go 是静态类型语言，编译期能发现大部分类型错误，CI 流程可以快速反馈。

### 快速构建
Go 编译速度快，适合频繁的自动化构建。

### 单一二进制文件
部署时只需要一个可执行文件，简化了部署复杂性。

### 丰富的测试工具
内置测试框架，支持单元测试、基准测试、覆盖率分析。

---

## GitHub Actions 实战

GitHub Actions 是 GitHub 提供的 CI/CD 服务，与代码仓库深度集成。

### 基础工作流

让我们从一个简单的 Go 项目开始：

::: details 示例：GitHub Actions 配置
```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v ./...
    
    - name: Build
      run: go build -v ./...
```
:::
### 工作流程解释

**触发条件**：推送到 main/develop 分支或创建针对 main 分支的 PR 时触发。

**运行环境**：使用 Ubuntu 最新版本，GitHub 提供虚拟机环境。

**步骤说明**：
1. 检出代码到运行环境
2. 安装指定版本的 Go
3. 下载项目依赖
4. 运行测试用例
5. 构建应用

---

## 完整的 CI/CD 流水线

实际项目需要更完整的流程：

::: details 完整的 GitHub Actions 配置
```yaml
# .github/workflows/cicd.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main, develop ]
    tags: [ 'v*' ]
  pull_request:
    branches: [ main ]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  # 代码质量检查
  lint:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: golangci-lint
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

  # 测试
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/test?sslmode=disable
    
    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3

  # 构建和推送镜像
  build:
    needs: [lint, test]
    runs-on: ubuntu-latest
    if: github.event_name != 'pull_request'
    
    steps:
    - uses: actions/checkout@v4
    
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
    
    - name: Build and push
      uses: docker/build-push-action@v5
      with:
        context: .
        push: true
        tags: ${{ steps.meta.outputs.tags }}
        labels: ${{ steps.meta.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max

  # 部署到测试环境
  deploy-staging:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/develop'
    environment: staging
    
    steps:
    - name: Deploy to staging
      run: |
        echo "Deploying to staging environment"
        # 这里会调用部署脚本或 Kubernetes 部署
  
  # 部署到生产环境
  deploy-production:
    needs: build
    runs-on: ubuntu-latest
    if: startsWith(github.ref, 'refs/tags/v')
    environment: production
    
    steps:
    - name: Deploy to production
      run: |
        echo "Deploying to production environment"
        # 生产环境部署逻辑
```
:::

### 流水线阶段说明

**1. 代码检查（Lint）**
使用 golangci-lint 检查代码风格、潜在错误、性能问题等。

**2. 测试（Test）**
运行单元测试、集成测试，生成覆盖率报告。包括启动数据库服务进行真实环境测试。

**3. 构建（Build）**
构建 Docker 镜像并推送到镜像仓库。只在非 PR 时执行。

**4. 部署（Deploy）**
根据分支自动部署到不同环境：
- develop 分支 → 测试环境
- 版本标签 → 生产环境

---

## GitLab CI/CD

GitLab CI/CD 是另一个流行选择，特别适合私有部署：

::: details 示例：GitLab CI 配置示例
```yaml
# .gitlab-ci.yml
stages:
  - test
  - build
  - deploy

variables:
  GO_VERSION: "1.21"
  DOCKER_DRIVER: overlay2

before_script:
  - apt-get update -qq && apt-get install -y -qq git ca-certificates
  - go version

# 测试阶段
test:
  stage: test
  image: golang:$GO_VERSION
  services:
    - postgres:13
  variables:
    POSTGRES_DB: test
    POSTGRES_USER: test
    POSTGRES_PASSWORD: test
    DATABASE_URL: postgres://test:test@postgres:5432/test?sslmode=disable
  script:
    - go mod download
    - go test -v -race -cover ./...
  coverage: '/coverage: \d+\.\d+% of statements/'

# 构建阶段
build:
  stage: build
  image: docker:latest
  services:
    - docker:dind
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  only:
    - main
    - develop

# 部署到测试环境
deploy:staging:
  stage: deploy
  script:
    - echo "Deploying to staging"
    # 部署脚本
  environment:
    name: staging
    url: https://staging.example.com
  only:
    - develop

# 部署到生产环境
deploy:production:
  stage: deploy
  script:
    - echo "Deploying to production"
    # 生产部署脚本
  environment:
    name: production
    url: https://example.com
  only:
    - main
  when: manual
```
:::

### GitLab CI 特点

**内置镜像仓库**：每个项目都有专用的 Docker 镜像仓库。

**环境管理**：可以定义多个部署环境，支持手动批准。

**Runner 支持**：支持自托管 Runner，适合私有环境。

---

## 部署策略

不同的部署策略适用于不同场景：

### 蓝绿部署（Blue-Green Deployment）

**原理**：维护两个相同的生产环境（蓝环境和绿环境），任何时候只有一个环境对外提供服务。

**优点**：
- 零停机部署
- 快速回滚（切换环境）
- 完整的预生产验证

**缺点**：
- 资源消耗翻倍
- 数据库状态管理复杂

**适用场景**：对可用性要求极高的系统，如金融交易系统。

### 滚动部署（Rolling Deployment）

**原理**：逐步替换旧版本实例，直到所有实例都更新完成。

**优点**：
- 资源利用率高
- 部署过程平滑
- 可以控制更新速度

**缺点**：
- 部署时间较长
- 版本混合可能导致兼容性问题

**适用场景**：无状态服务，如 Web API。

### 金丝雀部署（Canary Deployment）

**原理**：新版本只部署到小部分实例，观察运行效果后再逐步扩大范围。

**优点**：
- 风险可控
- 可以收集真实用户反馈
- 问题影响范围小

**缺点**：
- 部署过程复杂
- 需要良好的监控支持

**适用场景**：面向用户的应用，特别是新功能发布。

---

## 环境管理

### 环境分层

典型的环境分层策略：

::: details 示例：环境分层
```
开发环境 (dev) → 测试环境 (test) → 预发布环境 (staging) → 生产环境 (prod)
```
:::

**开发环境**：开发者本地或共享的开发环境，数据可以随意修改。

**测试环境**：QA 团队进行功能测试、集成测试的环境。

**预发布环境**：与生产环境配置完全一致，用于最后验证。

**生产环境**：对外提供服务的正式环境。

### 配置管理

不同环境的配置应该分离：

::: details 示例：配置管理
```go
// config/config.go
package config

import (
    "os"
    "strconv"
)

type Config struct {
    Port         int
    DatabaseURL  string
    RedisURL     string
    Environment  string
}

func Load() *Config {
    port, _ := strconv.Atoi(getEnv("PORT", "8080"))
    
    return &Config{
        Port:        port,
        DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost/myapp"),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
        Environment: getEnv("ENVIRONMENT", "development"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```
:::
### 密钥管理

敏感信息不应该硬编码在代码中：

**GitHub Actions Secrets**：
::: details 示例：GitHub Actions Secrets
```yaml
- name: Deploy
  env:
    DATABASE_PASSWORD: ${{ secrets.DB_PASSWORD }}
    API_KEY: ${{ secrets.API_KEY }}
  run: ./deploy.sh
```
:::
**Kubernetes Secrets**：
::: details 示例：Kubernetes Secrets
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
data:
  database-password: <base64-encoded-password>
  api-key: <base64-encoded-key>
```
:::
---

## 监控和通知

### 部署状态监控

部署完成后要验证应用是否正常运行：

::: details 示例：部署状态监控
```yaml
- name: Health Check
  run: |
    for i in {1..30}; do
      if curl -f http://localhost:8080/health; then
        echo "Application is healthy"
        exit 0
      fi
      sleep 10
    done
    echo "Health check failed"
    exit 1
```
:::
### 失败通知

当构建或部署失败时，及时通知相关人员：

::: details 示例：失败通知
```yaml
- name: Notify on failure
  if: failure()
  uses: 8398a7/action-slack@v3
  with:
    status: failure
    channel: '#deployments'
    webhook_url: ${{ secrets.SLACK_WEBHOOK }}
```
:::
---

## 性能优化技巧

### 构建缓存

利用缓存减少重复工作：

::: details 示例：构建缓存
```yaml
- name: Cache Go modules
  uses: actions/cache@v3
  with:
    path: ~/go/pkg/mod
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      ${{ runner.os }}-go-
```
:::
### 并行执行

将独立的任务并行执行：

::: details 示例：并行执行
```yaml
jobs:
  test:
    strategy:
      matrix:
        go-version: [1.20, 1.21]
        os: [ubuntu-latest, windows-latest]
    runs-on: ${{ matrix.os }}
```
:::
### Docker 层缓存

使用多阶段构建和缓存优化 Docker 构建：

::: details 示例：Docker 层缓存
```dockerfile
# 缓存友好的 Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app

# 先复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 再复制源码
COPY . .
RUN go build -o main .
```
:::
---

## 常见问题与解决方案

### 测试数据库问题

**问题**：测试需要数据库，但 CI 环境没有。

**解决方案**：
::: details 示例：测试数据库问题
```yaml
services:
  postgres:
    image: postgres:13
    env:
      POSTGRES_PASSWORD: postgres
    options: >-
      --health-cmd pg_isready
      --health-interval 10s
      --health-timeout 5s
      --health-retries 5
```
:::
### 依赖下载慢

**问题**：每次构建都要重新下载 Go 模块。

**解决方案**：使用模块代理和缓存
::: details 示例：依赖下载慢
```yaml
env:
  GOPROXY: https://proxy.golang.org,direct
  GOSUMDB: sum.golang.org

- name: Cache Go modules
  uses: actions/cache@v3
  # ... 缓存配置
```
:::
### 构建时间过长

**问题**：项目变大后，构建时间越来越长。

**解决方案**：
1. 使用并行构建
2. 优化 Dockerfile 层缓存
3. 只在必要时构建（条件触发）

---

## 平台对比

| 特性 | GitHub Actions | GitLab CI | Jenkins |
|------|----------------|-----------|---------|
| 托管方式 | 云托管 | 云托管/私有部署 | 私有部署 |
| 与代码仓库集成 | 深度集成 | 深度集成 | 需要配置 |
| 配置文件 | YAML | YAML | Groovy/YAML |
| 免费额度 | 2000分钟/月 | 400分钟/月 | 无限制 |
| 学习曲线 | 简单 | 中等 | 复杂 |
| 生态系统 | 丰富 | 丰富 | 最丰富 |

---

## 💡 最佳实践总结

1. **小步快跑**：频繁提交小的变更，而不是大规模修改
2. **测试优先**：确保测试覆盖率，测试失败时不允许部署
3. **环境一致性**：使用容器确保各环境配置一致
4. **监控部署**：部署后要验证应用健康状态
5. **快速回滚**：确保能够快速回滚到上一个稳定版本

CI/CD 不仅仅是技术工具，更是一种开发文化。它让团队能够更快地交付高质量的软件，减少人为错误，提高开发效率。

🚀 接下来推荐阅读：[监控和日志](/practice/deployment/monitoring)，学习如何观测生产环境中的应用运行状态。
