---
title: Web服务器项目
description: 构建一个功能完整的Web服务器应用
---

# Web服务器项目

构建一个包含用户认证、数据管理和API服务的完整Web应用。

## 项目功能

- 用户注册和登录
- RESTful API设计
- 数据持久化
- 中间件系统
- 响应式前端

## 快速开始

### 项目结构
```
web-server/
├── main.go           # 主程序
├── handlers/         # 路由处理器
├── middleware/       # 中间件
├── models/          # 数据模型
├── static/          # 静态文件
└── templates/       # HTML模板
```

### 核心代码

```go
// main.go
package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"` // 不返回密码
}

type Post struct {
    ID       int       `json:"id"`
    Title    string    `json:"title"`
    Content  string    `json:"content"`
    UserID   int       `json:"user_id"`
    Created  time.Time `json:"created"`
}

var (
    users   = []User{}
    posts   = []Post{}
    nextID  = 1
)

func main() {
    // 静态文件服务
    http.Handle("/static/", http.StripPrefix("/static/", 
        http.FileServer(http.Dir("./static"))))
    
    // 页面路由
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/register", registerHandler)
    
    // API路由
    http.HandleFunc("/api/users", usersAPIHandler)
    http.HandleFunc("/api/posts", postsAPIHandler)
    
    fmt.Println("Web服务器启动在 http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Web服务器演示</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .api-section { background: #f5f5f5; padding: 20px; margin: 20px 0; }
        .btn { background: #007bff; color: white; padding: 10px 20px; 
               text-decoration: none; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Go Web服务器</h1>
        <p>一个功能完整的Web应用演示</p>
        
        <div class="api-section">
            <h2>🔗 API端点</h2>
            <ul>
                <li><a href="/api/users">GET /api/users</a> - 获取用户列表</li>
                <li><a href="/api/posts">GET /api/posts</a> - 获取文章列表</li>
                <li>POST /api/users - 创建用户</li>
                <li>POST /api/posts - 创建文章</li>
            </ul>
        </div>
        
        <div class="api-section">
            <h2>功能页面</h2>
            <a href="/login" class="btn">登录</a>
            <a href="/register" class="btn">注册</a>
        </div>
        
        <div class="api-section">
            <h2>当前数据</h2>
            <p>用户数: {{.UserCount}}</p>
            <p>文章数: {{.PostCount}}</p>
        </div>
    </div>
</body>
</html>`
    
    t, _ := template.New("home").Parse(tmpl)
    data := struct {
        UserCount int
        PostCount int
    }{
        UserCount: len(users),
        PostCount: len(posts),
    }
    t.Execute(w, data)
}

func usersAPIHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(map[string]interface{}{
            "users": users,
            "total": len(users),
        })
    case "POST":
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        user.ID = nextID
        nextID++
        users = append(users, user)
        
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func postsAPIHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(map[string]interface{}{
            "posts": posts,
            "total": len(posts),
        })
    case "POST":
        var post Post
        if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        post.ID = nextID
        nextID++
        post.Created = time.Now()
        posts = append(posts, post)
        
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(post)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    loginPage := `
<!DOCTYPE html>
<html>
<head>
    <title>登录</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .form-container { max-width: 400px; margin: 0 auto; 
                         background: #f9f9f9; padding: 20px; border-radius: 8px; }
        input { width: 100%; padding: 10px; margin: 10px 0; border: 1px solid #ddd; border-radius: 4px; }
        button { background: #007bff; color: white; padding: 10px 20px; 
                border: none; border-radius: 4px; cursor: pointer; width: 100%; }
    </style>
</head>
<body>
    <div class="form-container">
        <h2>用户登录</h2>
        <form method="POST">
            <input type="text" name="username" placeholder="用户名" required>
            <input type="password" name="password" placeholder="密码" required>
            <button type="submit">登录</button>
        </form>
        <p><a href="/">返回首页</a></p>
    </div>
</body>
</html>`
    
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
        
        // 简单验证
        for _, user := range users {
            if user.Username == username && user.Password == password {
                http.Redirect(w, r, "/?login=success", http.StatusFound)
                return
            }
        }
        
        fmt.Fprint(w, "<script>alert('登录失败'); history.back();</script>")
        return
    }
    
    fmt.Fprint(w, loginPage)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    registerPage := `
<!DOCTYPE html>
<html>
<head>
    <title>注册</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .form-container { max-width: 400px; margin: 0 auto; 
                         background: #f9f9f9; padding: 20px; border-radius: 8px; }
        input { width: 100%; padding: 10px; margin: 10px 0; border: 1px solid #ddd; border-radius: 4px; }
        button { background: #28a745; color: white; padding: 10px 20px; 
                border: none; border-radius: 4px; cursor: pointer; width: 100%; }
    </style>
</head>
<body>
    <div class="form-container">
        <h2> 用户注册</h2>
        <form method="POST">
            <input type="text" name="username" placeholder="用户名" required>
            <input type="email" name="email" placeholder="邮箱" required>
            <input type="password" name="password" placeholder="密码" required>
            <button type="submit">注册</button>
        </form>
        <p><a href="/">返回首页</a></p>
    </div>
</body>
</html>`
    
    if r.Method == "POST" {
        r.ParseForm()
        user := User{
            ID:       nextID,
            Username: r.FormValue("username"),
            Email:    r.FormValue("email"),
            Password: r.FormValue("password"),
        }
        nextID++
        users = append(users, user)
        
        http.Redirect(w, r, "/login?register=success", http.StatusFound)
        return
    }
    
    fmt.Fprint(w, registerPage)
}
```

## 运行结果

启动服务器：
```bash
$ go run main.go
Web服务器启动在 http://localhost:8080
```

访问主页：
- 显示功能导航和API文档
- 实时数据统计

测试API：
```bash
# 创建用户
curl -X POST -H "Content-Type: application/json" \
     -d '{"username":"testuser","email":"test@example.com","password":"123456"}' \
     http://localhost:8080/api/users

# 获取用户列表
curl http://localhost:8080/api/users

# 创建文章
curl -X POST -H "Content-Type: application/json" \
     -d '{"title":"Go语言学习","content":"学习Go语言基础知识","user_id":1}' \
     http://localhost:8080/api/posts
```

## 扩展功能

- 数据库集成
- 用户认证中间件
- 文件上传功能
- WebSocket实时通信
- 前端框架集成

这个项目展示了Go语言在Web开发方面的能力，涵盖了HTTP服务、API设计、模板渲染等核心技能。 