# 网络编程与HTTP

网络编程是Go语言的强项之一，Go内置的`net/http`包提供了功能强大且易用的HTTP客户端和服务器实现。本章将学习如何使用Go进行网络编程。

## 🌐 HTTP客户端

### 基础HTTP请求

Go的`net/http`包提供了简单易用的HTTP客户端功能：

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

func main() {
    // GET请求
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    // 读取响应体
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("状态码: %d\n", resp.StatusCode)
    fmt.Printf("响应头: %v\n", resp.Header)
    fmt.Printf("响应体: %s\n", string(body))
}
```

**运行结果：**
```
状态码: 200
响应头: map[Content-Type:[application/json] ...]
响应体: {
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  },
  "origin": "xxx.xxx.xxx.xxx",
  "url": "https://httpbin.org/get"
}
```

### 自定义HTTP客户端

创建自定义的HTTP客户端可以设置超时、代理等参数：

```go
package main

import (
    "fmt"
    "net/http"
    "time"
    "context"
    "strings"
)

func main() {
    // 创建自定义客户端
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    // GET请求示例
    makeGetRequest(client)
    
    // POST请求示例
    makePostRequest(client)
    
    // 带上下文的请求
    makeRequestWithContext(client)
}

func makeGetRequest(client *http.Client) {
    fmt.Println("=== GET请求示例 ===")
    
    resp, err := client.Get("https://httpbin.org/get?name=Go&version=1.21")
    if err != nil {
        fmt.Printf("GET请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("GET状态码: %d\n", resp.StatusCode)
}

func makePostRequest(client *http.Client) {
    fmt.Println("\n=== POST请求示例 ===")
    
    // JSON数据
    jsonData := `{"name": "Go语言", "type": "编程语言"}`
    
    resp, err := client.Post(
        "https://httpbin.org/post",
        "application/json",
        strings.NewReader(jsonData),
    )
    if err != nil {
        fmt.Printf("POST请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("POST状态码: %d\n", resp.StatusCode)
}

func makeRequestWithContext(client *http.Client) {
    fmt.Println("\n=== 带上下文的请求 ===")
    
    // 创建带超时的上下文
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
    if err != nil {
        fmt.Printf("创建请求失败: %v\n", err)
        return
    }
    
    // 添加请求头
    req.Header.Set("User-Agent", "Go-Learning-Client/1.0")
    req.Header.Set("Accept", "application/json")
    
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("请求失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("上下文请求状态码: %d\n", resp.StatusCode)
}
```

**运行结果：**
```
=== GET请求示例 ===
GET状态码: 200

=== POST请求示例 ===
POST状态码: 200

=== 带上下文的请求 ===
上下文请求状态码: 200
```

## 🖥️ HTTP服务器

### 基础HTTP服务器

使用Go创建HTTP服务器非常简单：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    // 注册路由处理器
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/time", timeHandler)
    http.HandleFunc("/json", jsonHandler)
    
    fmt.Println("服务器启动在 :8080")
    fmt.Println("访问 http://localhost:8080")
    
    // 启动服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `
    <h1>Go HTTP服务器</h1>
    <p>欢迎访问Go语言HTTP服务器示例</p>
    <ul>
        <li><a href="/hello">Hello页面</a></li>
        <li><a href="/time">当前时间</a></li>
        <li><a href="/json">JSON数据</a></li>
    </ul>
    `)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "世界"
    }
    
    fmt.Fprintf(w, "<h1>你好, %s!</h1>", name)
    fmt.Fprintf(w, "<p>请求方法: %s</p>", r.Method)
    fmt.Fprintf(w, "<p>请求路径: %s</p>", r.URL.Path)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(w, "<h1>当前时间</h1>")
    fmt.Fprintf(w, "<p>%s</p>", currentTime)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    jsonData := `{
        "message": "Hello from Go Server",
        "timestamp": "` + time.Now().Format(time.RFC3339) + `",
        "server": "Go HTTP Server"
    }`
    
    fmt.Fprint(w, jsonData)
}
```

**运行结果：**
访问 `http://localhost:8080` 会显示主页，点击链接可以访问不同的页面。

### RESTful API服务器

创建一个完整的RESTful API服务器：

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

// 用户结构体
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// 用户存储（内存中）
type UserStore struct {
    mu    sync.RWMutex
    users map[int]*User
    nextID int
}

func NewUserStore() *UserStore {
    return &UserStore{
        users: make(map[int]*User),
        nextID: 1,
    }
}

func (s *UserStore) CreateUser(user *User) *User {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    user.ID = s.nextID
    s.nextID++
    s.users[user.ID] = user
    return user
}

func (s *UserStore) GetUser(id int) (*User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    user, exists := s.users[id]
    return user, exists
}

func (s *UserStore) GetAllUsers() []*User {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
}

func (s *UserStore) UpdateUser(id int, user *User) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if _, exists := s.users[id]; !exists {
        return false
    }
    
    user.ID = id
    s.users[id] = user
    return true
}

func (s *UserStore) DeleteUser(id int) bool {
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
    store *UserStore
}

func NewAPIServer() *APIServer {
    return &APIServer{
        store: NewUserStore(),
    }
}

func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.handleGetUsers(w, r)
    case http.MethodPost:
        s.handleCreateUser(w, r)
    default:
        http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
    }
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) {
    // 从URL路径中提取用户ID
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 3 {
        http.Error(w, "无效的用户ID", http.StatusBadRequest)
        return
    }
    
    idStr := pathParts[2]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "无效的用户ID", http.StatusBadRequest)
        return
    }
    
    switch r.Method {
    case http.MethodGet:
        s.handleGetUser(w, r, id)
    case http.MethodPut:
        s.handleUpdateUser(w, r, id)
    case http.MethodDelete:
        s.handleDeleteUser(w, r, id)
    default:
        http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
    }
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    users := s.store.GetAllUsers()
    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, "编码错误", http.StatusInternalServerError)
    }
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "无效的JSON数据", http.StatusBadRequest)
        return
    }
    
    createdUser := s.store.CreateUser(&user)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request, id int) {
    user, exists := s.store.GetUser(id)
    if !exists {
        http.Error(w, "用户不存在", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request, id int) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "无效的JSON数据", http.StatusBadRequest)
        return
    }
    
    if !s.store.UpdateUser(id, &user) {
        http.Error(w, "用户不存在", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request, id int) {
    if !s.store.DeleteUser(id) {
        http.Error(w, "用户不存在", http.StatusNotFound)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    server := NewAPIServer()
    
    // 初始化一些测试数据
    server.store.CreateUser(&User{Name: "张三", Email: "zhangsan@example.com"})
    server.store.CreateUser(&User{Name: "李四", Email: "lisi@example.com"})
    
    // 路由设置
    http.HandleFunc("/users", server.handleUsers)
    http.HandleFunc("/users/", server.handleUser)
    
    // 添加根路径处理器
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        
        fmt.Fprint(w, `
        <h1>用户管理API</h1>
        <h2>API端点:</h2>
        <ul>
            <li>GET /users - 获取所有用户</li>
            <li>POST /users - 创建用户</li>
            <li>GET /users/{id} - 获取特定用户</li>
            <li>PUT /users/{id} - 更新用户</li>
            <li>DELETE /users/{id} - 删除用户</li>
        </ul>
        `)
    })
    
    fmt.Println("RESTful API服务器启动在 :8080")
    fmt.Println("访问 http://localhost:8080 查看API文档")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### API测试示例

使用curl命令测试API：

```bash
# 获取所有用户
curl http://localhost:8080/users

# 创建新用户
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"王五","email":"wangwu@example.com"}'

# 获取特定用户
curl http://localhost:8080/users/1

# 更新用户
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三更新","email":"zhangsan_new@example.com"}'

# 删除用户
curl -X DELETE http://localhost:8080/users/1
```

## 🔧 中间件

中间件是处理HTTP请求的强大模式：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// 中间件类型
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 日志中间件
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 调用下一个处理器
        next(w, r)
        
        // 记录请求信息
        duration := time.Since(start)
        log.Printf("%s %s %v", r.Method, r.URL.Path, duration)
    }
}

// 认证中间件
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        if token != "Bearer secret-token" {
            http.Error(w, "未授权", http.StatusUnauthorized)
            return
        }
        
        next(w, r)
    }
}

// CORS中间件
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}

// 组合多个中间件
func ChainMiddleware(middlewares ...Middleware) Middleware {
    return func(next http.HandlerFunc) http.HandlerFunc {
        for i := len(middlewares) - 1; i >= 0; i-- {
            next = middlewares[i](next)
        }
        return next
    }
}

// 处理器函数
func publicHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "这是公开端点")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "这是受保护的端点")
}

func main() {
    // 公开端点（只有日志和CORS中间件）
    http.HandleFunc("/public", 
        ChainMiddleware(LoggingMiddleware, CORSMiddleware)(publicHandler))
    
    // 受保护端点（包含认证中间件）
    http.HandleFunc("/protected", 
        ChainMiddleware(LoggingMiddleware, CORSMiddleware, AuthMiddleware)(protectedHandler))
    
    fmt.Println("服务器启动在 :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**测试中间件：**

```bash
# 访问公开端点
curl http://localhost:8080/public

# 访问受保护端点（无认证）
curl http://localhost:8080/protected

# 访问受保护端点（有认证）
curl -H "Authorization: Bearer secret-token" http://localhost:8080/protected
```

## 🔄 WebSocket

Go也支持WebSocket实时通信：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    
    "golang.org/x/net/websocket"
)

// 客户端连接
type Client struct {
    id     string
    conn   *websocket.Conn
    server *Server
}

// WebSocket服务器
type Server struct {
    clients map[string]*Client
    mutex   sync.RWMutex
}

func NewServer() *Server {
    return &Server{
        clients: make(map[string]*Client),
    }
}

func (s *Server) addClient(client *Client) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.clients[client.id] = client
    log.Printf("客户端 %s 已连接", client.id)
}

func (s *Server) removeClient(clientID string) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    delete(s.clients, clientID)
    log.Printf("客户端 %s 已断开", clientID)
}

func (s *Server) broadcast(message string) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    for _, client := range s.clients {
        if err := websocket.Message.Send(client.conn, message); err != nil {
            log.Printf("发送消息失败: %v", err)
        }
    }
}

func (s *Server) handleConnection(ws *websocket.Conn) {
    clientID := fmt.Sprintf("client_%d", len(s.clients)+1)
    client := &Client{
        id:     clientID,
        conn:   ws,
        server: s,
    }
    
    s.addClient(client)
    defer s.removeClient(clientID)
    
    // 发送欢迎消息
    welcomeMsg := fmt.Sprintf("欢迎 %s!", clientID)
    websocket.Message.Send(ws, welcomeMsg)
    
    // 广播新用户加入
    s.broadcast(fmt.Sprintf("%s 加入了聊天室", clientID))
    
    // 监听消息
    for {
        var message string
        if err := websocket.Message.Receive(ws, &message); err != nil {
            log.Printf("接收消息错误: %v", err)
            break
        }
        
        broadcastMsg := fmt.Sprintf("%s: %s", clientID, message)
        log.Printf("收到消息: %s", broadcastMsg)
        s.broadcast(broadcastMsg)
    }
    
    // 广播用户离开
    s.broadcast(fmt.Sprintf("%s 离开了聊天室", clientID))
}

func main() {
    server := NewServer()
    
    // WebSocket处理器
    http.Handle("/chat", websocket.Handler(server.handleConnection))
    
    // 静态文件服务器（聊天页面）
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Go WebSocket 聊天室</title>
        </head>
        <body>
            <div id="messages"></div>
            <input type="text" id="messageInput" placeholder="输入消息...">
            <button onclick="sendMessage()">发送</button>
            
            <script>
                const ws = new WebSocket('ws://localhost:8080/chat');
                const messages = document.getElementById('messages');
                
                ws.onmessage = function(event) {
                    const div = document.createElement('div');
                    div.textContent = event.data;
                    messages.appendChild(div);
                };
                
                function sendMessage() {
                    const input = document.getElementById('messageInput');
                    ws.send(input.value);
                    input.value = '';
                }
                
                document.getElementById('messageInput').addEventListener('keypress', function(e) {
                    if (e.key === 'Enter') {
                        sendMessage();
                    }
                });
            </script>
        </body>
        </html>
        `)
    })
    
    fmt.Println("WebSocket服务器启动在 :8080")
    fmt.Println("访问 http://localhost:8080 使用聊天室")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 💡 最佳实践

### 1. 错误处理
```go
func handleAPI(w http.ResponseWriter, r *http.Request) {
    data, err := processData()
    if err != nil {
        http.Error(w, "处理数据失败", http.StatusInternalServerError)
        log.Printf("API错误: %v", err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### 2. 请求超时
```go
server := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}
```

### 3. 优雅关闭
```go
func gracefulShutdown() {
    server := &http.Server{Addr: ":8080"}
    
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("服务器启动失败: %v", err)
        }
    }()
    
    // 等待中断信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("正在关闭服务器...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("服务器强制关闭:", err)
    }
    
    log.Println("服务器已关闭")
}
```

## 🎯 实践练习

1. **HTTP客户端练习**：创建一个程序，调用公开API（如天气API）并解析响应
2. **RESTful API练习**：扩展用户管理API，添加分页、搜索功能
3. **中间件练习**：实现速率限制中间件
4. **WebSocket练习**：创建一个实时聊天应用

Go的网络编程能力强大而简洁，是构建现代Web应用和微服务的理想选择。通过这些示例，您可以开始构建自己的网络应用程序。 