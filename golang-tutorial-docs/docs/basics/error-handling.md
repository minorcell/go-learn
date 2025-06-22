# 错误处理

Go语言采用显式错误处理的设计哲学，通过返回值来处理错误。这种方式虽然代码较长，但让错误处理变得明确和可预测。

## 📖 本章内容

- 错误的基本概念
- 错误创建和处理
- 自定义错误类型
- 错误包装和链
- 错误处理最佳实践

## ❌ 错误基础

### 内置错误类型

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

// 基本错误处理
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

// 多种错误情况
func validateAge(age int) error {
    if age < 0 {
        return errors.New("年龄不能为负数")
    }
    if age > 150 {
        return errors.New("年龄不能超过150岁")
    }
    if age < 18 {
        return errors.New("年龄必须满18岁")
    }
    return nil
}

// 字符串转换示例
func parseAndValidate(input string) (int, error) {
    // 尝试转换字符串为整数
    value, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("无法解析 '%s' 为整数: %v", input, err)
    }
    
    // 验证范围
    if value < 1 || value > 100 {
        return 0, fmt.Errorf("值 %d 超出有效范围 [1, 100]", value)
    }
    
    return value, nil
}

// 文件操作模拟
func readConfig(filename string) (map[string]string, error) {
    // 模拟文件不存在
    if filename == "" {
        return nil, errors.New("文件名不能为空")
    }
    
    if filename == "missing.txt" {
        return nil, fmt.Errorf("文件 '%s' 不存在", filename)
    }
    
    // 模拟权限错误
    if filename == "protected.txt" {
        return nil, fmt.Errorf("没有权限读取文件 '%s'", filename)
    }
    
    // 模拟成功读取
    config := map[string]string{
        "host": "localhost",
        "port": "8080",
        "name": "myapp",
    }
    
    return config, nil
}

func main() {
    fmt.Println("=== 基本错误处理 ===")
    
    // 除法测试
    tests := []struct {
        a, b float64
        desc string
    }{
        {10, 2, "正常除法"},
        {10, 0, "除零错误"},
        {15, 3, "正常除法"},
    }
    
    for _, test := range tests {
        result, err := divide(test.a, test.b)
        if err != nil {
            fmt.Printf("%s: 错误 - %v\n", test.desc, err)
        } else {
            fmt.Printf("%s: %.2f / %.2f = %.2f\n", test.desc, test.a, test.b, result)
        }
    }
    
    fmt.Println("\n=== 年龄验证 ===")
    
    ages := []int{-5, 25, 160, 15, 30}
    for _, age := range ages {
        if err := validateAge(age); err != nil {
            fmt.Printf("年龄 %d: 验证失败 - %v\n", age, err)
        } else {
            fmt.Printf("年龄 %d: 验证通过\n", age)
        }
    }
    
    fmt.Println("\n=== 字符串解析 ===")
    
    inputs := []string{"50", "abc", "150", "0", "75"}
    for _, input := range inputs {
        value, err := parseAndValidate(input)
        if err != nil {
            fmt.Printf("输入 '%s': 解析失败 - %v\n", input, err)
        } else {
            fmt.Printf("输入 '%s': 解析成功，值为 %d\n", input, value)
        }
    }
    
    fmt.Println("\n=== 文件读取 ===")
    
    files := []string{"config.txt", "", "missing.txt", "protected.txt"}
    for _, filename := range files {
        config, err := readConfig(filename)
        if err != nil {
            fmt.Printf("文件 '%s': 读取失败 - %v\n", filename, err)
        } else {
            fmt.Printf("文件 '%s': 读取成功 - %v\n", filename, config)
        }
    }
}
```

### 错误检查模式

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

// 用户结构体
type User struct {
    ID       int
    Username string
    Email    string
    Age      int
}

// 用户验证错误
var (
    ErrInvalidUsername = errors.New("用户名无效")
    ErrInvalidEmail    = errors.New("邮箱无效")
    ErrInvalidAge      = errors.New("年龄无效")
    ErrUserNotFound    = errors.New("用户不存在")
    ErrUserExists      = errors.New("用户已存在")
)

// 用户验证
func validateUser(user User) error {
    // 用户名验证
    if len(user.Username) < 3 {
        return fmt.Errorf("%w: 用户名长度至少3个字符", ErrInvalidUsername)
    }
    
    if strings.Contains(user.Username, " ") {
        return fmt.Errorf("%w: 用户名不能包含空格", ErrInvalidUsername)
    }
    
    // 邮箱验证
    if !strings.Contains(user.Email, "@") {
        return fmt.Errorf("%w: 邮箱必须包含@符号", ErrInvalidEmail)
    }
    
    // 年龄验证
    if user.Age < 13 || user.Age > 120 {
        return fmt.Errorf("%w: 年龄必须在13-120之间", ErrInvalidAge)
    }
    
    return nil
}

// 模拟数据库
var userDB = map[int]User{
    1: {ID: 1, Username: "alice", Email: "alice@example.com", Age: 25},
    2: {ID: 2, Username: "bob", Email: "bob@example.com", Age: 30},
}

// 获取用户
func getUser(id int) (User, error) {
    user, exists := userDB[id]
    if !exists {
        return User{}, fmt.Errorf("%w: ID为%d的用户", ErrUserNotFound, id)
    }
    return user, nil
}

// 创建用户
func createUser(user User) error {
    // 验证用户数据
    if err := validateUser(user); err != nil {
        return fmt.Errorf("用户验证失败: %w", err)
    }
    
    // 检查用户是否已存在
    for _, existingUser := range userDB {
        if existingUser.Username == user.Username {
            return fmt.Errorf("%w: 用户名 '%s'", ErrUserExists, user.Username)
        }
        if existingUser.Email == user.Email {
            return fmt.Errorf("%w: 邮箱 '%s'", ErrUserExists, user.Email)
        }
    }
    
    // 分配新ID
    user.ID = len(userDB) + 1
    userDB[user.ID] = user
    
    return nil
}

// 错误类型检查
func handleUserError(err error) {
    if err == nil {
        return
    }
    
    // 使用 errors.Is 检查错误类型
    switch {
    case errors.Is(err, ErrUserNotFound):
        fmt.Printf("🔍 用户查找错误: %v\n", err)
    case errors.Is(err, ErrUserExists):
        fmt.Printf("⚠️ 用户冲突错误: %v\n", err)
    case errors.Is(err, ErrInvalidUsername):
        fmt.Printf("👤 用户名错误: %v\n", err)
    case errors.Is(err, ErrInvalidEmail):
        fmt.Printf("📧 邮箱错误: %v\n", err)
    case errors.Is(err, ErrInvalidAge):
        fmt.Printf("🎂 年龄错误: %v\n", err)
    default:
        fmt.Printf("❌ 未知错误: %v\n", err)
    }
}

// 批量处理用户操作
func processUsers() {
    fmt.Println("=== 用户操作测试 ===")
    
    // 测试获取用户
    fmt.Println("\n--- 获取用户 ---")
    for _, id := range []int{1, 2, 999} {
        user, err := getUser(id)
        if err != nil {
            fmt.Printf("获取用户ID=%d失败: ", id)
            handleUserError(err)
        } else {
            fmt.Printf("获取用户成功: %+v\n", user)
        }
    }
    
    // 测试创建用户
    fmt.Println("\n--- 创建用户 ---")
    newUsers := []User{
        {Username: "charlie", Email: "charlie@example.com", Age: 28},
        {Username: "a", Email: "short@example.com", Age: 25},      // 用户名太短
        {Username: "dave", Email: "invalid-email", Age: 22},         // 邮箱无效
        {Username: "eve", Email: "eve@example.com", Age: 10},        // 年龄无效
        {Username: "alice", Email: "alice2@example.com", Age: 30},   // 用户名已存在
        {Username: "frank", Email: "alice@example.com", Age: 35},    // 邮箱已存在
        {Username: "valid user", Email: "user@example.com", Age: 25}, // 用户名包含空格
    }
    
    for _, newUser := range newUsers {
        fmt.Printf("创建用户 '%s': ", newUser.Username)
        err := createUser(newUser)
        if err != nil {
            handleUserError(err)
        } else {
            fmt.Printf("✅ 创建成功\n")
        }
    }
    
    fmt.Println("\n--- 最终用户列表 ---")
    for id, user := range userDB {
        fmt.Printf("ID=%d: %+v\n", id, user)
    }
}

func main() {
    processUsers()
}
```

## 🎯 自定义错误类型

### 结构体错误

```go
package main

import (
    "fmt"
    "time"
)

// 验证错误
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("验证错误: 字段 '%s' 的值 '%v' %s", e.Field, e.Value, e.Message)
}

// 网络错误
type NetworkError struct {
    Operation string
    URL       string
    Err       error
    Timestamp time.Time
    Retries   int
}

func (e NetworkError) Error() string {
    return fmt.Sprintf("网络错误 [%s]: %s 操作失败 (重试%d次) - %v", 
        e.Timestamp.Format("15:04:05"), e.Operation, e.Retries, e.Err)
}

func (e NetworkError) Unwrap() error {
    return e.Err
}

// 业务逻辑错误
type BusinessError struct {
    Code    string
    Message string
    Context map[string]interface{}
}

func (e BusinessError) Error() string {
    return fmt.Sprintf("业务错误 [%s]: %s", e.Code, e.Message)
}

func (e BusinessError) ErrorCode() string {
    return e.Code
}

// 用户服务
type UserService struct {
    users map[string]User
}

type User struct {
    ID    string
    Name  string
    Email string
    Age   int
}

func NewUserService() *UserService {
    return &UserService{
        users: make(map[string]User),
    }
}

// 验证用户输入
func (us *UserService) validateUser(user User) error {
    if len(user.Name) < 2 {
        return ValidationError{
            Field:   "Name",
            Value:   user.Name,
            Message: "长度必须至少2个字符",
        }
    }
    
    if user.Age < 0 || user.Age > 150 {
        return ValidationError{
            Field:   "Age",
            Value:   user.Age,
            Message: "必须在0-150之间",
        }
    }
    
    if user.Email == "" {
        return ValidationError{
            Field:   "Email",
            Value:   user.Email,
            Message: "不能为空",
        }
    }
    
    return nil
}

// 创建用户
func (us *UserService) CreateUser(user User) error {
    // 验证输入
    if err := us.validateUser(user); err != nil {
        return err
    }
    
    // 检查用户是否已存在
    if _, exists := us.users[user.ID]; exists {
        return BusinessError{
            Code:    "USER_EXISTS",
            Message: fmt.Sprintf("用户ID '%s' 已存在", user.ID),
            Context: map[string]interface{}{
                "userID": user.ID,
                "action": "create",
            },
        }
    }
    
    // 模拟网络操作失败
    if user.ID == "network_fail" {
        return NetworkError{
            Operation: "CreateUser",
            URL:       "https://api.example.com/users",
            Err:       fmt.Errorf("connection timeout"),
            Timestamp: time.Now(),
            Retries:   3,
        }
    }
    
    us.users[user.ID] = user
    return nil
}

// 获取用户
func (us *UserService) GetUser(id string) (User, error) {
    if id == "" {
        return User{}, ValidationError{
            Field:   "ID",
            Value:   id,
            Message: "不能为空",
        }
    }
    
    user, exists := us.users[id]
    if !exists {
        return User{}, BusinessError{
            Code:    "USER_NOT_FOUND",
            Message: fmt.Sprintf("未找到ID为 '%s' 的用户", id),
            Context: map[string]interface{}{
                "userID": id,
                "action": "get",
            },
        }
    }
    
    return user, nil
}

// 错误处理器
func handleError(err error) {
    if err == nil {
        return
    }
    
    switch e := err.(type) {
    case ValidationError:
        fmt.Printf("🔍 验证失败: %s\n", e.Error())
        fmt.Printf("  - 字段: %s\n", e.Field)
        fmt.Printf("  - 值: %v\n", e.Value)
        
    case NetworkError:
        fmt.Printf("🌐 网络异常: %s\n", e.Error())
        fmt.Printf("  - 操作: %s\n", e.Operation)
        fmt.Printf("  - URL: %s\n", e.URL)
        fmt.Printf("  - 时间: %s\n", e.Timestamp.Format("2006-01-02 15:04:05"))
        
    case BusinessError:
        fmt.Printf("💼 业务错误: %s\n", e.Error())
        fmt.Printf("  - 错误码: %s\n", e.Code)
        if len(e.Context) > 0 {
            fmt.Printf("  - 上下文: %v\n", e.Context)
        }
        
    default:
        fmt.Printf("❌ 未知错误: %s\n", err.Error())
    }
}

func main() {
    service := NewUserService()
    
    fmt.Println("=== 用户服务测试 ===")
    
    // 测试用户创建
    testUsers := []User{
        {ID: "1", Name: "Alice", Email: "alice@example.com", Age: 25},
        {ID: "2", Name: "B", Email: "bob@example.com", Age: 30},        // 名字太短
        {ID: "3", Name: "Charlie", Email: "", Age: 35},                  // 邮箱为空
        {ID: "4", Name: "David", Email: "david@example.com", Age: -5},   // 年龄无效
        {ID: "1", Name: "Alice2", Email: "alice2@example.com", Age: 28}, // ID重复
        {ID: "network_fail", Name: "Network", Email: "net@example.com", Age: 30}, // 网络错误
    }
    
    fmt.Println("\n--- 创建用户测试 ---")
    for _, user := range testUsers {
        fmt.Printf("\n创建用户: %+v\n", user)
        err := service.CreateUser(user)
        if err != nil {
            handleError(err)
        } else {
            fmt.Printf("✅ 用户创建成功\n")
        }
    }
    
    // 测试用户获取
    fmt.Println("\n--- 获取用户测试 ---")
    testIDs := []string{"1", "999", "", "3"}
    
    for _, id := range testIDs {
        fmt.Printf("\n获取用户ID: '%s'\n", id)
        user, err := service.GetUser(id)
        if err != nil {
            handleError(err)
        } else {
            fmt.Printf("✅ 用户信息: %+v\n", user)
        }
    }
    
    // 显示所有用户
    fmt.Println("\n--- 所有用户 ---")
    for id, user := range service.users {
        fmt.Printf("ID=%s: %+v\n", id, user)
    }
}
```

### 错误链和包装

```go
package main

import (
    "errors"
    "fmt"
)

// 自定义错误类型
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("数据库错误: %s操作在表'%s'失败: %v", e.Operation, e.Table, e.Err)
}

func (e DatabaseError) Unwrap() error {
    return e.Err
}

// 服务层错误
type ServiceError struct {
    Service string
    Method  string
    Err     error
}

func (e ServiceError) Error() string {
    return fmt.Sprintf("服务错误: %s.%s() - %v", e.Service, e.Method, e.Err)
}

func (e ServiceError) Unwrap() error {
    return e.Err
}

// 模拟数据库层
func dbQuery(table, query string) error {
    // 模拟不同的数据库错误
    switch table {
    case "users":
        if query == "invalid" {
            return errors.New("SQL语法错误")
        }
        return nil
    case "orders":
        return errors.New("连接超时")
    case "products":
        return errors.New("表不存在")
    default:
        return nil
    }
}

// 数据访问层
func getUserFromDB(userID string) error {
    err := dbQuery("users", userID)
    if err != nil {
        return DatabaseError{
            Operation: "SELECT",
            Table:     "users",
            Err:       err,
        }
    }
    return nil
}

func getOrdersFromDB(userID string) error {
    err := dbQuery("orders", userID)
    if err != nil {
        return DatabaseError{
            Operation: "SELECT",
            Table:     "orders",
            Err:       err,
        }
    }
    return nil
}

func getProductsFromDB(category string) error {
    err := dbQuery("products", category)
    if err != nil {
        return DatabaseError{
            Operation: "SELECT",
            Table:     "products",
            Err:       err,
        }
    }
    return nil
}

// 服务层
func getUserProfile(userID string) error {
    err := getUserFromDB(userID)
    if err != nil {
        return ServiceError{
            Service: "UserService",
            Method:  "GetProfile",
            Err:     err,
        }
    }
    return nil
}

func getUserOrders(userID string) error {
    err := getOrdersFromDB(userID)
    if err != nil {
        return ServiceError{
            Service: "OrderService",
            Method:  "GetUserOrders",
            Err:     err,
        }
    }
    return nil
}

func getProductsByCategory(category string) error {
    err := getProductsFromDB(category)
    if err != nil {
        return ServiceError{
            Service: "ProductService",
            Method:  "GetByCategory",
            Err:     err,
        }
    }
    return nil
}

// 应用层 - 组合多个服务调用
func getUserDashboard(userID string) error {
    // 获取用户资料
    if err := getUserProfile(userID); err != nil {
        return fmt.Errorf("获取用户面板失败(用户资料): %w", err)
    }
    
    // 获取用户订单
    if err := getUserOrders(userID); err != nil {
        return fmt.Errorf("获取用户面板失败(用户订单): %w", err)
    }
    
    return nil
}

// 错误分析函数
func analyzeError(err error) {
    if err == nil {
        fmt.Println("✅ 操作成功")
        return
    }
    
    fmt.Printf("❌ 错误分析: %v\n", err)
    
    // 检查错误链
    fmt.Println("\n错误链分析:")
    currentErr := err
    level := 0
    
    for currentErr != nil {
        indent := ""
        for i := 0; i < level; i++ {
            indent += "  "
        }
        
        fmt.Printf("%s- %v (类型: %T)\n", indent, currentErr, currentErr)
        
        // 检查具体错误类型
        switch e := currentErr.(type) {
        case ServiceError:
            fmt.Printf("%s  服务: %s, 方法: %s\n", indent, e.Service, e.Method)
        case DatabaseError:
            fmt.Printf("%s  操作: %s, 表: %s\n", indent, e.Operation, e.Table)
        }
        
        // 获取下一层错误
        currentErr = errors.Unwrap(currentErr)
        level++
        
        if level > 10 { // 防止无限循环
            fmt.Printf("%s... (错误链过长，停止展示)\n", indent)
            break
        }
    }
    
    // 检查特定错误类型
    fmt.Println("\n错误类型检查:")
    
    var dbErr DatabaseError
    if errors.As(err, &dbErr) {
        fmt.Printf("- 发现数据库错误: 操作=%s, 表=%s\n", dbErr.Operation, dbErr.Table)
    }
    
    var serviceErr ServiceError
    if errors.As(err, &serviceErr) {
        fmt.Printf("- 发现服务错误: 服务=%s, 方法=%s\n", serviceErr.Service, serviceErr.Method)
    }
    
    // 检查根本原因
    fmt.Println("\n根本原因分析:")
    if errors.Is(err, errors.New("SQL语法错误")) {
        fmt.Println("- 根本原因: SQL语法问题")
    } else if errors.Is(err, errors.New("连接超时")) {
        fmt.Println("- 根本原因: 网络连接问题")
    } else if errors.Is(err, errors.New("表不存在")) {
        fmt.Println("- 根本原因: 数据库结构问题")
    }
}

func main() {
    fmt.Println("=== 错误链和包装演示 ===")
    
    // 测试不同场景
    testCases := []struct {
        name string
        fn   func() error
    }{
        {
            name: "正常用户查询",
            fn:   func() error { return getUserProfile("123") },
        },
        {
            name: "SQL错误用户查询", 
            fn:   func() error { return getUserProfile("invalid") },
        },
        {
            name: "用户订单查询(超时)",
            fn:   func() error { return getUserOrders("456") },
        },
        {
            name: "产品查询(表不存在)",
            fn:   func() error { return getProductsByCategory("electronics") },
        },
        {
            name: "用户面板(组合操作)",
            fn:   func() error { return getUserDashboard("789") },
        },
    }
    
    for i, tc := range testCases {
        fmt.Printf("\n=== 测试 %d: %s ===\n", i+1, tc.name)
        err := tc.fn()
        analyzeError(err)
    }
    
    // 演示 errors.Is 和 errors.As 的使用
    fmt.Println("\n=== errors.Is 和 errors.As 演示 ===")
    
    err := getUserOrders("test")
    
    // 使用 errors.Is 检查特定错误
    timeoutErr := errors.New("连接超时")
    if errors.Is(err, timeoutErr) {
        fmt.Println("✅ 检测到连接超时错误")
    }
    
    // 使用 errors.As 提取特定类型
    var dbErr DatabaseError
    if errors.As(err, &dbErr) {
        fmt.Printf("✅ 提取到数据库错误: %+v\n", dbErr)
    }
    
    var svcErr ServiceError
    if errors.As(err, &svcErr) {
        fmt.Printf("✅ 提取到服务错误: %+v\n", svcErr)
    }
}
```

## 📝 本章小结

在这一章中，我们学习了Go语言的错误处理机制：

### 🔹 错误基础
- **error接口** - Go的内置错误类型
- **显式处理** - 通过返回值处理错误
- **errors包** - 创建和操作错误的工具
- **fmt.Errorf** - 格式化错误信息

### 🔹 错误检查
- **nil检查** - 判断是否有错误发生
- **errors.Is** - 检查错误是否为特定类型
- **errors.As** - 提取特定类型的错误
- **预定义错误** - 使用全局错误变量

### 🔹 自定义错误
- **结构体错误** - 实现error接口的结构体
- **错误上下文** - 携带额外信息的错误
- **错误分类** - 按业务逻辑分类错误
- **错误方法** - 提供额外的错误信息访问

### 🔹 错误包装
- **错误链** - 通过Unwrap构建错误层次
- **fmt.Errorf %w** - 包装底层错误
- **errors.Unwrap** - 获取被包装的错误
- **错误传播** - 在调用栈中传递错误信息

### 🔹 最佳实践
- 总是检查错误返回值
- 使用有意义的错误消息
- 在适当的层级处理错误
- 不要忽略或隐藏错误
- 使用错误包装保留上下文
- 定义特定领域的错误类型

### 🔹 设计原则
- **明确性** - 错误处理应该明确可见
- **一致性** - 整个代码库保持一致的错误处理风格
- **信息丰富** - 错误信息应该帮助调试和排错
- **分层处理** - 在不同层级采用不同的错误处理策略

## 🎉 基础语法完成

恭喜！你已经完成了Go语言基础语法的学习。现在你应该掌握了：

- ✅ 变量和类型系统
- ✅ 控制流程和循环
- ✅ 函数定义和调用
- ✅ 数组、切片和映射
- ✅ 结构体和方法
- ✅ 接口和多态
- ✅ 错误处理机制

## 🎯 下一步

完成基础语法后，建议继续学习Go语言的进阶主题：

- **并发编程** - goroutine和channel
- **包管理** - 模块系统和依赖管理
- **标准库** - 文件操作、网络编程、JSON处理
- **实战项目** - 构建完整的应用程序

继续你的Go语言学习之旅吧！🚀 