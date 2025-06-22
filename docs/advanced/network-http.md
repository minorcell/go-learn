---
title: 网络编程与HTTP
description: 学习Go语言的HTTP客户端、服务器开发和网络通信
---

# 网络编程与HTTP

网络编程是现代应用开发的核心技能。Go语言内置的`net/http`包提供了功能强大且易用的HTTP客户端和服务器实现，让网络编程变得简单高效。

## 本章内容

- HTTP客户端开发和使用
- HTTP服务器构建和路由
- RESTful API设计和实现
- 中间件和请求处理
- WebSocket实时通信基础

## 网络编程概念

### HTTP协议基础

HTTP是无状态的请求-响应协议，基于TCP/IP通信：

- **客户端-服务器模型**：客户端发起请求，服务器返回响应
- **无状态协议**：每个请求独立，服务器不保存客户端状态
- **方法语义**：GET(查询)、POST(创建)、PUT(更新)、DELETE(删除)
- **状态码**：200(成功)、404(未找到)、500(服务器错误)

### Go网络编程优势

| 特性 | 说明 | 优势 |
|------|------|------|
| **并发支持** | 每个连接一个goroutine | 高并发处理能力 |
| **标准库完整** | 内置HTTP/HTTPS支持 | 无需第三方依赖 |
| **性能优秀** | 高效的网络I/O | 低延迟高吞吐 |
| **部署简单** | 单二进制文件 | 容易部署和运维 |

::: tip 设计原则
Go网络编程遵循"简单、高效、并发"的设计理念：
- 使用标准库优先
- 利用goroutines处理并发
- 注重错误处理和资源管理
:::

## HTTP客户端

### 基础HTTP请求

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

func simpleGet() {
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        fmt.Printf("请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("读取响应失败: %v\n", err)
        return
    }
    
    fmt.Printf("状态码: %d\n", resp.StatusCode)
    fmt.Printf("响应内容: %s\n", string(body))
}

func simplePost() {
    data := strings.NewReader(`{"name": "Go", "type": "language"}`)
    
    resp, err := http.Post(
        "https://httpbin.org/post",
        "application/json",
        data,
    )
    if err != nil {
        fmt.Printf("POST请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("POST状态码: %d\n", resp.StatusCode)
}
```

### 自定义HTTP客户端

创建可配置的HTTP客户端处理复杂场景：

```go
func createCustomClient() *http.Client {
    return &http.Client{
        Timeout: 10 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     30 * time.Second,
        },
    }
}

func makeRequestWithHeaders() {
    client := createCustomClient()
    
    req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
    if err != nil {
        fmt.Printf("创建请求失败: %v\n", err)
        return
    }
    
    // 设置请求头
    req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Bearer your-token")
    
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("自定义请求状态码: %d\n", resp.StatusCode)
}
```

## HTTP服务器

### 基础服务器

```go
func startBasicServer() {
    // 注册路由处理器
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api/health", healthHandler)
    http.HandleFunc("/api/echo", echoHandler)
    
    fmt.Println("服务器启动在 :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("服务器启动失败: %v\n", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <h1>Go HTTP 服务器</h1>
        <p>欢迎访问Go语言HTTP服务器</p>
        <ul>
            <li><a href="/api/health">健康检查</a></li>
            <li><a href="/api/echo">Echo API</a></li>
        </ul>
    `)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"status": "ok", "timestamp": "%s"}`, 
        time.Now().Format(time.RFC3339))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "只支持POST方法", http.StatusMethodNotAllowed)
        return
    }
    
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "读取请求体失败", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()
    
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"echo": %q}`, string(body))
}
```

### 路由和中间件

使用结构化的方式管理路由和中间件：

```go
type Server struct {
    router *http.ServeMux
    addr   string
}

func NewServer(addr string) *Server {
    return &Server{
        router: http.NewServeMux(),
        addr:   addr,
    }
}

func (s *Server) routes() {
    s.router.HandleFunc("/api/users", s.handleUsers())
    s.router.HandleFunc("/api/users/", s.handleUserByID())
}

func (s *Server) Start() error {
    s.routes()
    
    fmt.Printf("服务器启动在 %s\n", s.addr)
    return http.ListenAndServe(s.addr, s.middleware(s.router))
}

// 中间件：日志记录
func (s *Server) middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 添加CORS头
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // 处理OPTIONS预检请求
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
        
        fmt.Printf("[%s] %s %s - %v\n", 
            time.Now().Format("2006-01-02 15:04:05"),
            r.Method, r.URL.Path, time.Since(start))
    })
}
```

## 实战项目：RESTful API服务

让我们构建一个完整的用户管理API服务：

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "sync"
    "time"
)

// 数据模型
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

type UpdateUserRequest struct {
    Name  string `json:"name,omitempty"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"age,omitempty"`
}

type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Message string      `json:"message,omitempty"`
}

// 内存数据存储
type UserStore struct {
    mu    sync.RWMutex
    users map[int]*User
    nextID int
}

func NewUserStore() *UserStore {
    return &UserStore{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

func (s *UserStore) Create(req CreateUserRequest) *User {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    user := &User{
        ID:        s.nextID,
        Name:      req.Name,
        Email:     req.Email,
        Age:       req.Age,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    s.users[s.nextID] = user
    s.nextID++
    
    return user
}

func (s *UserStore) GetAll() []*User {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    
    return users
}

func (s *UserStore) GetByID(id int) (*User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    user, exists := s.users[id]
    return user, exists
}

func (s *UserStore) Update(id int, req UpdateUserRequest) (*User, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    user, exists := s.users[id]
    if !exists {
        return nil, false
    }
    
    if req.Name != "" {
        user.Name = req.Name
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    if req.Age != 0 {
        user.Age = req.Age
    }
    user.UpdatedAt = time.Now()
    
    return user, true
}

func (s *UserStore) Delete(id int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if _, exists := s.users[id]; !exists {
        return false
    }
    
    delete(s.users, id)
    return true
}

// API服务器
type APIServer struct {
    store  *UserStore
    router *http.ServeMux
    addr   string
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{
        store:  NewUserStore(),
        router: http.NewServeMux(),
        addr:   addr,
    }
}

func (s *APIServer) setupRoutes() {
    s.router.HandleFunc("/api/users", s.handleUsers)
    s.router.HandleFunc("/api/users/", s.handleUserByID)
    s.router.HandleFunc("/api/stats", s.handleStats)
    s.router.HandleFunc("/", s.handleHome)
}

func (s *APIServer) Start() error {
    s.setupRoutes()
    
    // 添加示例数据
    s.seedData()
    
    fmt.Printf("🚀 API服务器启动在 %s\n", s.addr)
    fmt.Printf("📝 API文档: http://localhost%s\n", s.addr)
    
    return http.ListenAndServe(s.addr, s.corsMiddleware(s.loggingMiddleware(s.router)))
}

func (s *APIServer) seedData() {
    users := []CreateUserRequest{
        {Name: "张三", Email: "zhangsan@example.com", Age: 25},
        {Name: "李四", Email: "lisi@example.com", Age: 30},
        {Name: "王五", Email: "wangwu@example.com", Age: 28},
    }
    
    for _, user := range users {
        s.store.Create(user)
    }
    
    fmt.Println("📊 已添加示例数据")
}

// 中间件
func (s *APIServer) corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

func (s *APIServer) loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        fmt.Printf("[%s] %s %s - %v\n", 
            time.Now().Format("15:04:05"), r.Method, r.URL.Path, time.Since(start))
    })
}

// 路由处理器
func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        s.writeErrorResponse(w, http.StatusNotFound, "页面未找到")
        return
    }
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>用户管理API</title>
            <style>body { font-family: Arial, sans-serif; margin: 40px; }</style>
        </head>
        <body>
            <h1>🚀 用户管理API服务</h1>
            <h2>📖 API文档</h2>
            <ul>
                <li><strong>GET /api/users</strong> - 获取所有用户</li>
                <li><strong>POST /api/users</strong> - 创建用户</li>
                <li><strong>GET /api/users/{id}</strong> - 获取指定用户</li>
                <li><strong>PUT /api/users/{id}</strong> - 更新用户</li>
                <li><strong>DELETE /api/users/{id}</strong> - 删除用户</li>
                <li><strong>GET /api/stats</strong> - 获取统计信息</li>
            </ul>
            <h2>🧪 测试示例</h2>
            <pre>
# 获取所有用户
curl http://localhost%s/api/users

# 创建用户
curl -X POST http://localhost%s/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","email":"new@example.com","age":25}'

# 获取用户详情
curl http://localhost%s/api/users/1
            </pre>
        </body>
        </html>
    `, s.addr, s.addr, s.addr)
}

func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.getAllUsers(w, r)
    case http.MethodPost:
        s.createUser(w, r)
    default:
        s.writeErrorResponse(w, http.StatusMethodNotAllowed, "方法不允许")
    }
}

func (s *APIServer) handleUserByID(w http.ResponseWriter, r *http.Request) {
    // 从URL路径提取ID
    path := strings.TrimPrefix(r.URL.Path, "/api/users/")
    id, err := strconv.Atoi(path)
    if err != nil {
        s.writeErrorResponse(w, http.StatusBadRequest, "无效的用户ID")
        return
    }
    
    switch r.Method {
    case http.MethodGet:
        s.getUserByID(w, r, id)
    case http.MethodPut:
        s.updateUser(w, r, id)
    case http.MethodDelete:
        s.deleteUser(w, r, id)
    default:
        s.writeErrorResponse(w, http.StatusMethodNotAllowed, "方法不允许")
    }
}

func (s *APIServer) getAllUsers(w http.ResponseWriter, r *http.Request) {
    users := s.store.GetAll()
    s.writeJSONResponse(w, http.StatusOK, APIResponse{
        Success: true,
        Data:    users,
        Message: fmt.Sprintf("成功获取 %d 个用户", len(users)),
    })
}

func (s *APIServer) createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.writeErrorResponse(w, http.StatusBadRequest, "无效的JSON数据")
        return
    }
    
    // 简单验证
    if req.Name == "" || req.Email == "" || req.Age <= 0 {
        s.writeErrorResponse(w, http.StatusBadRequest, "姓名、邮箱和年龄都是必填项")
        return
    }
    
    user := s.store.Create(req)
    s.writeJSONResponse(w, http.StatusCreated, APIResponse{
        Success: true,
        Data:    user,
        Message: "用户创建成功",
    })
}

func (s *APIServer) getUserByID(w http.ResponseWriter, r *http.Request, id int) {
    user, exists := s.store.GetByID(id)
    if !exists {
        s.writeErrorResponse(w, http.StatusNotFound, "用户不存在")
        return
    }
    
    s.writeJSONResponse(w, http.StatusOK, APIResponse{
        Success: true,
        Data:    user,
        Message: "成功获取用户信息",
    })
}

func (s *APIServer) updateUser(w http.ResponseWriter, r *http.Request, id int) {
    var req UpdateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.writeErrorResponse(w, http.StatusBadRequest, "无效的JSON数据")
        return
    }
    
    user, exists := s.store.Update(id, req)
    if !exists {
        s.writeErrorResponse(w, http.StatusNotFound, "用户不存在")
        return
    }
    
    s.writeJSONResponse(w, http.StatusOK, APIResponse{
        Success: true,
        Data:    user,
        Message: "用户更新成功",
    })
}

func (s *APIServer) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
    if !s.store.Delete(id) {
        s.writeErrorResponse(w, http.StatusNotFound, "用户不存在")
        return
    }
    
    s.writeJSONResponse(w, http.StatusOK, APIResponse{
        Success: true,
        Message: "用户删除成功",
    })
}

func (s *APIServer) handleStats(w http.ResponseWriter, r *http.Request) {
    users := s.store.GetAll()
    
    stats := map[string]interface{}{
        "total_users": len(users),
        "server_time": time.Now().Format(time.RFC3339),
        "uptime":     "运行中",
    }
    
    s.writeJSONResponse(w, http.StatusOK, APIResponse{
        Success: true,
        Data:    stats,
        Message: "统计信息获取成功",
    })
}

// 工具方法
func (s *APIServer) writeJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}

func (s *APIServer) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    s.writeJSONResponse(w, statusCode, APIResponse{
        Success: false,
        Error:   message,
    })
}

func main() {
    server := NewAPIServer(":8080")
    
    if err := server.Start(); err != nil {
        fmt.Printf("服务器启动失败: %v\n", err)
    }
}
```

## WebSocket基础

WebSocket提供了全双工通信能力，适合实时应用：

```go
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
    defer ws.Close()
    
    fmt.Println("新的WebSocket连接")
    
    for {
        var message string
        err := websocket.Message.Receive(ws, &message)
        if err != nil {
            fmt.Printf("WebSocket读取错误: %v\n", err)
            break
        }
        
        fmt.Printf("收到消息: %s\n", message)
        
        response := fmt.Sprintf("服务器回复: %s", message)
        err = websocket.Message.Send(ws, response)
        if err != nil {
            fmt.Printf("WebSocket发送错误: %v\n", err)
            break
        }
    }
}

func startWebSocketServer() {
    http.Handle("/ws", websocket.Handler(echoHandler))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head><title>WebSocket测试</title></head>
            <body>
                <div id="messages"></div>
                <input type="text" id="input" placeholder="输入消息">
                <button onclick="send()">发送</button>
                
                <script>
                    const ws = new WebSocket('ws://localhost:8080/ws');
                    const messages = document.getElementById('messages');
                    
                    ws.onmessage = function(event) {
                        messages.innerHTML += '<p>' + event.data + '</p>';
                    };
                    
                    function send() {
                        const input = document.getElementById('input');
                        ws.send(input.value);
                        input.value = '';
                    }
                </script>
            </body>
            </html>
        `)
    })
    
    fmt.Println("WebSocket服务器启动在 :8080")
    http.ListenAndServe(":8080", nil)
}
```

## 最佳实践

### 1. 错误处理和日志

```go
func handleWithErrorLogging(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                fmt.Printf("处理器panic: %v\n", err)
                http.Error(w, "内部服务器错误", http.StatusInternalServerError)
            }
        }()
        
        handler(w, r)
    }
}
```

### 2. 请求限流

```go
type RateLimiter struct {
    tokens chan struct{}
}

func NewRateLimiter(rate int) *RateLimiter {
    rl := &RateLimiter{
        tokens: make(chan struct{}, rate),
    }
    
    // 填充令牌
    for i := 0; i < rate; i++ {
        rl.tokens <- struct{}{}
    }
    
    return rl
}

func (rl *RateLimiter) Allow() bool {
    select {
    case <-rl.tokens:
        return true
    default:
        return false
    }
}

func rateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !rl.Allow() {
                http.Error(w, "请求过于频繁", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

### 3. 优雅关闭

```go
func gracefulServer() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: nil,
    }
    
    // 启动服务器
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("服务器启动失败: %v\n", err)
        }
    }()
    
    // 等待中断信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    fmt.Println("关闭服务器...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        fmt.Printf("服务器强制关闭: %v\n", err)
    }
    
    fmt.Println("服务器已关闭")
}
```

## 本章小结

Go网络编程的核心要点：

- **HTTP客户端**：使用net/http包进行HTTP通信，支持自定义配置
- **HTTP服务器**：构建高性能的Web服务和API
- **RESTful设计**：遵循REST原则设计清晰的API接口
- **中间件模式**：实现横切关注点如日志、认证、限流
- **错误处理**：完善的错误处理和资源管理

::: tip 练习建议
1. 实现一个完整的RESTful API服务
2. 添加身份验证和权限控制
3. 集成数据库存储用户数据
4. 实现WebSocket聊天室功能
:::