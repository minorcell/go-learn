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

/*
项目案例3：Web服务器
功能：
1. 静态文件服务
2. RESTful API
3. JSON响应
4. HTML模板
5. 中间件
*/

// 用户数据结构
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// 模拟数据库
var users = []User{
	{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25},
	{ID: 2, Name: "李四", Email: "lisi@example.com", Age: 30},
	{ID: 3, Name: "王五", Email: "wangwu@example.com", Age: 28},
}

// 消息结构
type Message struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 日志中间件
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("开始处理请求: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("请求处理完成: %s %s - 耗时: %v", r.Method, r.URL.Path, duration)
	}
}

// CORS中间件
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// 首页处理器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web服务器示例</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background-color: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #333; text-align: center; }
        .api-list { background: #f8f9fa; padding: 20px; border-radius: 5px; margin: 20px 0; }
        .api-item { margin: 10px 0; padding: 10px; background: white; border-left: 4px solid #007bff; }
        .method { font-weight: bold; color: #007bff; }
        .url { font-family: monospace; background: #e9ecef; padding: 2px 6px; border-radius: 3px; }
        .footer { text-align: center; margin-top: 30px; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Go Web服务器示例</h1>
        <p>欢迎使用Go语言构建的Web服务器！这个示例展示了Go的HTTP服务器功能。</p>
        
        <div class="api-list">
            <h2>📋 可用的API端点：</h2>
            
            <div class="api-item">
                <span class="method">GET</span> <span class="url">/</span> - 首页
            </div>
            
            <div class="api-item">
                <span class="method">GET</span> <span class="url">/api/users</span> - 获取所有用户
            </div>
            
            <div class="api-item">
                <span class="method">GET</span> <span class="url">/api/users/{id}</span> - 获取指定用户
            </div>
            
            <div class="api-item">
                <span class="method">POST</span> <span class="url">/api/users</span> - 创建新用户
            </div>
            
            <div class="api-item">
                <span class="method">PUT</span> <span class="url">/api/users/{id}</span> - 更新用户
            </div>
            
            <div class="api-item">
                <span class="method">DELETE</span> <span class="url">/api/users/{id}</span> - 删除用户
            </div>
            
            <div class="api-item">
                <span class="method">GET</span> <span class="url">/api/status</span> - 服务器状态
            </div>
            
            <div class="api-item">
                <span class="method">GET</span> <span class="url">/time</span> - 当前时间
            </div>
        </div>
        
        <h2>🧪 测试示例：</h2>
        <pre style="background: #f8f9fa; padding: 15px; border-radius: 5px; overflow-x: auto;">
# 获取所有用户
curl http://localhost:8080/api/users

# 获取指定用户
curl http://localhost:8080/api/users/1

# 创建新用户
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","email":"new@example.com","age":25}'

# 获取服务器状态
curl http://localhost:8080/api/status
        </pre>
        
        <div class="footer">
            <p>⚡ 服务器运行时间: {{.StartTime}}</p>
            <p>🌐 当前时间: {{.CurrentTime}}</p>
        </div>
    </div>
</body>
</html>`

	tmpl, err := template.New("home").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "模板解析错误", http.StatusInternalServerError)
		return
	}

	data := struct {
		StartTime   string
		CurrentTime string
	}{
		StartTime:   startTime.Format("2006-01-02 15:04:05"),
		CurrentTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

// 获取所有用户
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Message{
		Status:  "success",
		Message: "获取用户列表成功",
		Data:    users,
	}

	json.NewEncoder(w).Encode(response)
}

// 根据ID获取用户
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 从URL路径中提取ID
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := Message{
			Status:  "error",
			Message: "无效的用户ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 查找用户
	for _, user := range users {
		if user.ID == id {
			response := Message{
				Status:  "success",
				Message: "获取用户信息成功",
				Data:    user,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// 用户不存在
	response := Message{
		Status:  "error",
		Message: "用户不存在",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// 创建新用户
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		response := Message{
			Status:  "error",
			Message: "只支持POST方法",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		response := Message{
			Status:  "error",
			Message: "JSON格式错误",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 生成新ID
	maxID := 0
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	newUser.ID = maxID + 1

	// 添加到用户列表
	users = append(users, newUser)

	response := Message{
		Status:  "success",
		Message: "用户创建成功",
		Data:    newUser,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// 更新用户
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "PUT" {
		response := Message{
			Status:  "error",
			Message: "只支持PUT方法",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 从URL路径中提取ID
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := Message{
			Status:  "error",
			Message: "无效的用户ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		response := Message{
			Status:  "error",
			Message: "JSON格式错误",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 查找并更新用户
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser

			response := Message{
				Status:  "success",
				Message: "用户更新成功",
				Data:    updatedUser,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// 用户不存在
	response := Message{
		Status:  "error",
		Message: "用户不存在",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// 删除用户
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		response := Message{
			Status:  "error",
			Message: "只支持DELETE方法",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 从URL路径中提取ID
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := Message{
			Status:  "error",
			Message: "无效的用户ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// 查找并删除用户
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			response := Message{
				Status:  "success",
				Message: "用户删除成功",
				Data:    user,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// 用户不存在
	response := Message{
		Status:  "error",
		Message: "用户不存在",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// 服务器状态
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uptime := time.Since(startTime)
	status := struct {
		Status     string `json:"status"`
		Uptime     string `json:"uptime"`
		StartTime  string `json:"start_time"`
		UserCount  int    `json:"user_count"`
		GoVersion  string `json:"go_version"`
		ServerTime string `json:"server_time"`
	}{
		Status:     "running",
		Uptime:     uptime.String(),
		StartTime:  startTime.Format("2006-01-02 15:04:05"),
		UserCount:  len(users),
		GoVersion:  "Go 1.21+",
		ServerTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	response := Message{
		Status:  "success",
		Message: "服务器运行正常",
		Data:    status,
	}

	json.NewEncoder(w).Encode(response)
}

// 时间处理器
func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()
	timeInfo := struct {
		Timestamp int64  `json:"timestamp"`
		DateTime  string `json:"datetime"`
		UTC       string `json:"utc"`
		Unix      int64  `json:"unix"`
		Weekday   string `json:"weekday"`
		TimeZone  string `json:"timezone"`
	}{
		Timestamp: now.Unix(),
		DateTime:  now.Format("2006-01-02 15:04:05"),
		UTC:       now.UTC().Format("2006-01-02 15:04:05"),
		Unix:      now.Unix(),
		Weekday:   now.Weekday().String(),
		TimeZone:  now.Format("MST"),
	}

	response := Message{
		Status:  "success",
		Message: "获取时间信息成功",
		Data:    timeInfo,
	}

	json.NewEncoder(w).Encode(response)
}

var startTime time.Time

func main() {
	startTime = time.Now()

	fmt.Println("=== Go语言项目案例：Web服务器 ===")
	fmt.Printf("服务器启动时间: %s\n", startTime.Format("2006-01-02 15:04:05"))

	// 路由设置
	http.HandleFunc("/", loggingMiddleware(corsMiddleware(homeHandler)))
	http.HandleFunc("/api/users", loggingMiddleware(corsMiddleware(handleUsers)))
	http.HandleFunc("/api/users/", loggingMiddleware(corsMiddleware(handleUserByID)))
	http.HandleFunc("/api/status", loggingMiddleware(corsMiddleware(statusHandler)))
	http.HandleFunc("/time", loggingMiddleware(corsMiddleware(timeHandler)))

	port := ":8080"
	fmt.Printf("🚀 服务器启动成功！\n")
	fmt.Printf("📍 访问地址: http://localhost%s\n", port)
	fmt.Printf("📖 API文档: http://localhost%s/\n", port)
	fmt.Printf("⏰ 服务器状态: http://localhost%s/api/status\n", port)
	fmt.Println("按 Ctrl+C 停止服务器")

	log.Fatal(http.ListenAndServe(port, nil))
}

// 处理用户相关请求的路由函数
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersHandler(w, r)
	case "POST":
		createUserHandler(w, r)
	default:
		response := Message{
			Status:  "error",
			Message: "不支持的HTTP方法",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
	}
}

// 处理特定用户ID的请求
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUserHandler(w, r)
	case "PUT":
		updateUserHandler(w, r)
	case "DELETE":
		deleteUserHandler(w, r)
	default:
		response := Message{
			Status:  "error",
			Message: "不支持的HTTP方法",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
	}
}
