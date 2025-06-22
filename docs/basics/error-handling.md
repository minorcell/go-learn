---
title: 错误处理
description: 学习Go语言的错误处理机制、自定义错误和最佳实践
---

# 错误处理

Go语言采用显式错误处理的设计哲学，通过返回值来处理错误，而不是异常机制。这种方式让错误处理变得明确、可预测，是Go语言"简单胜过复杂"理念的体现。

## 本章内容

- Go错误处理的基本概念和哲学
- 内置错误类型和创建方法
- 自定义错误类型和错误包装
- 错误处理的最佳实践和常见模式
- 实际项目中的错误处理策略

## 错误处理哲学

### Go的错误处理理念

与许多语言使用异常不同，Go选择了**显式错误处理**：

**优点**：
- **明确性**：错误处理逻辑清晰可见
- **可控性**：开发者必须主动处理错误
- **性能**：避免了异常机制的性能开销
- **简单性**：没有复杂的异常层次结构

**Go错误处理的核心原则**：
```go
// 函数返回值的最后一个通常是error
func doSomething() (result string, err error) {
    // 实现
}

// 调用方必须检查错误
result, err := doSomething()
if err != nil {
    // 处理错误
    return err
}
// 使用result
```

## 基本错误处理

### 内置错误类型

Go的 `error` 是一个内置接口：

```go
type error interface {
    Error() string
}
```

### 创建和处理错误

```go
import (
    "errors"
    "fmt"
)

// 1. 使用errors.New创建错误
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

// 2. 使用fmt.Errorf格式化错误
func validateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("年龄不能为负数，输入值: %d", age)
    }
    if age > 150 {
        return fmt.Errorf("年龄%d超出合理范围(0-150)", age)
    }
    return nil
}

// 3. 预定义错误变量
var (
    ErrInvalidInput = errors.New("输入无效")
    ErrNotFound     = errors.New("未找到")
    ErrPermissionDenied = errors.New("权限不足")
)

func processUser(id int) error {
    if id <= 0 {
        return ErrInvalidInput
    }
    if id > 10000 {
        return ErrNotFound
    }
    // 正常处理
    return nil
}
```

### 错误检查模式

```go
// 基本模式：立即检查
result, err := someFunction()
if err != nil {
    return fmt.Errorf("操作失败: %w", err)
}

// 延迟处理模式
func processFile(filename string) (err error) {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("无法打开文件 %s: %w", filename, err)
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            // 如果之前没有错误，使用关闭错误
            if err == nil {
                err = fmt.Errorf("关闭文件失败: %w", closeErr)
            }
        }
    }()
    
    // 处理文件...
    return nil
}

// 错误累积模式
func validateData(data map[string]string) error {
    var errs []string
    
    if data["name"] == "" {
        errs = append(errs, "姓名不能为空")
    }
    if data["email"] == "" {
        errs = append(errs, "邮箱不能为空")
    }
    if !strings.Contains(data["email"], "@") {
        errs = append(errs, "邮箱格式无效")
    }
    
    if len(errs) > 0 {
        return fmt.Errorf("数据验证失败: %s", strings.Join(errs, "; "))
    }
    return nil
}
```

## 自定义错误类型

当需要更丰富的错误信息时，可以实现自定义错误类型：

### 结构体错误类型

```go
// 验证错误
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("字段 '%s' 验证失败: %s (值: %v)", 
        e.Field, e.Message, e.Value)
}

// HTTP错误
type HTTPError struct {
    StatusCode int
    Message    string
    Cause      error
}

func (e HTTPError) Error() string {
    if e.Cause != nil {
        return fmt.Sprintf("HTTP %d: %s (原因: %v)", 
            e.StatusCode, e.Message, e.Cause)
    }
    return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// 实现Unwrap方法支持错误包装
func (e HTTPError) Unwrap() error {
    return e.Cause
}

// 使用示例
func validateUser(user User) error {
    if user.Age < 0 {
        return ValidationError{
            Field:   "age",
            Value:   user.Age,
            Message: "不能为负数",
        }
    }
    return nil
}
```

### 错误包装和链

Go 1.13+ 支持错误包装，可以保留原始错误信息：

```go
import (
    "errors"
    "fmt"
)

// 使用fmt.Errorf的%w动词包装错误
func readUserConfig(userID int) (*Config, error) {
    config, err := loadConfigFile("user.json")
    if err != nil {
        return nil, fmt.Errorf("读取用户%d配置失败: %w", userID, err)
    }
    return config, nil
}

// 错误检查和解包
func handleError(err error) {
    // 检查是否是特定错误
    if errors.Is(err, ErrNotFound) {
        fmt.Println("资源未找到")
        return
    }
    
    // 检查是否是特定类型的错误
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("验证失败: 字段 %s\n", validationErr.Field)
        return
    }
    
    // 处理其他错误
    fmt.Printf("未知错误: %v\n", err)
}
```

## 错误处理模式

### 1. 哨兵错误（Sentinel Errors）

预定义的错误值，用于特定的错误条件：

```go
var (
    ErrUserNotFound = errors.New("用户不存在")
    ErrInvalidCredentials = errors.New("凭据无效")
    ErrAccessDenied = errors.New("访问被拒绝")
)

func authenticateUser(username, password string) error {
    user := findUser(username)
    if user == nil {
        return ErrUserNotFound
    }
    
    if !checkPassword(user, password) {
        return ErrInvalidCredentials
    }
    
    if !user.IsActive {
        return ErrAccessDenied
    }
    
    return nil
}

// 使用时检查特定错误
err := authenticateUser("john", "secret")
switch err {
case ErrUserNotFound:
    fmt.Println("请先注册")
case ErrInvalidCredentials:
    fmt.Println("密码错误")
case ErrAccessDenied:
    fmt.Println("账户已被禁用")
case nil:
    fmt.Println("登录成功")
default:
    fmt.Printf("登录失败: %v\n", err)
}
```

### 2. 错误类型断言

根据错误类型执行不同的处理逻辑：

```go
type NetworkError struct {
    Timeout bool
    Temporary bool
    Message string
}

func (e NetworkError) Error() string {
    return e.Message
}

func (e NetworkError) IsTimeout() bool {
    return e.Timeout
}

func (e NetworkError) IsTemporary() bool {
    return e.Temporary
}

// 智能重试逻辑
func performRequest() error {
    for retry := 0; retry < 3; retry++ {
        err := makeNetworkCall()
        if err == nil {
            return nil
        }
        
        // 检查是否是网络错误
        if netErr, ok := err.(NetworkError); ok {
            if netErr.IsTemporary() {
                fmt.Printf("临时错误，重试中... (第%d次)\n", retry+1)
                time.Sleep(time.Second * time.Duration(retry+1))
                continue
            }
            if netErr.IsTimeout() {
                fmt.Println("请求超时，稍后重试")
                return err
            }
        }
        
        // 非网络错误或不可重试的错误
        return err
    }
    return errors.New("重试次数超限")
}
```

### 3. 函数选项模式处理错误

```go
type DatabaseConfig struct {
    Host     string
    Port     int
    Database string
    Timeout  time.Duration
}

type Option func(*DatabaseConfig) error

func WithHost(host string) Option {
    return func(c *DatabaseConfig) error {
        if host == "" {
            return errors.New("主机地址不能为空")
        }
        c.Host = host
        return nil
    }
}

func WithPort(port int) Option {
    return func(c *DatabaseConfig) error {
        if port <= 0 || port > 65535 {
            return fmt.Errorf("端口号无效: %d", port)
        }
        c.Port = port
        return nil
    }
}

func NewDatabaseConfig(options ...Option) (*DatabaseConfig, error) {
    config := &DatabaseConfig{
        Host:    "localhost",
        Port:    5432,
        Timeout: 30 * time.Second,
    }
    
    for _, option := range options {
        if err := option(config); err != nil {
            return nil, fmt.Errorf("配置选项错误: %w", err)
        }
    }
    
    return config, nil
}
```

## 实践示例：用户管理服务

让我们通过一个用户管理服务展示错误处理的最佳实践：

```go
package main

import (
    "errors"
    "fmt"
    "regexp"
    "strings"
)

// 自定义错误类型
type UserError struct {
    Code    string
    Message string
    Field   string
}

func (e UserError) Error() string {
    if e.Field != "" {
        return fmt.Sprintf("[%s] %s: %s", e.Code, e.Field, e.Message)
    }
    return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// 预定义错误
var (
    ErrUserExists    = UserError{Code: "USER_EXISTS", Message: "用户已存在"}
    ErrUserNotFound  = UserError{Code: "USER_NOT_FOUND", Message: "用户不存在"}
    ErrInvalidEmail  = UserError{Code: "INVALID_EMAIL", Message: "邮箱格式无效", Field: "email"}
    ErrInvalidAge    = UserError{Code: "INVALID_AGE", Message: "年龄必须在18-100之间", Field: "age"}
)

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

type UserService struct {
    users map[int]*User
    nextID int
}

func NewUserService() *UserService {
    return &UserService{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

// 验证用户数据
func (s *UserService) validateUser(user *User) error {
    var errors []string
    
    // 姓名验证
    if strings.TrimSpace(user.Name) == "" {
        errors = append(errors, "姓名不能为空")
    }
    
    // 邮箱验证
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(user.Email) {
        errors = append(errors, "邮箱格式无效")
    }
    
    // 年龄验证
    if user.Age < 18 || user.Age > 100 {
        errors = append(errors, "年龄必须在18-100之间")
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("用户数据验证失败: %s", strings.Join(errors, "; "))
    }
    
    return nil
}

// 创建用户
func (s *UserService) CreateUser(name, email string, age int) (*User, error) {
    user := &User{
        Name:  name,
        Email: email,
        Age:   age,
    }
    
    // 验证用户数据
    if err := s.validateUser(user); err != nil {
        return nil, fmt.Errorf("创建用户失败: %w", err)
    }
    
    // 检查邮箱是否已存在
    for _, existingUser := range s.users {
        if existingUser.Email == email {
            return nil, fmt.Errorf("创建用户失败: %w", ErrUserExists)
        }
    }
    
    // 创建用户
    user.ID = s.nextID
    s.nextID++
    s.users[user.ID] = user
    
    return user, nil
}

// 获取用户
func (s *UserService) GetUser(id int) (*User, error) {
    user, exists := s.users[id]
    if !exists {
        return nil, fmt.Errorf("获取用户失败: %w", ErrUserNotFound)
    }
    return user, nil
}

// 更新用户
func (s *UserService) UpdateUser(id int, name, email string, age int) (*User, error) {
    user, err := s.GetUser(id)
    if err != nil {
        return nil, fmt.Errorf("更新用户失败: %w", err)
    }
    
    // 保存原始数据以便回滚
    original := *user
    
    // 更新数据
    user.Name = name
    user.Email = email
    user.Age = age
    
    // 验证更新后的数据
    if err := s.validateUser(user); err != nil {
        // 回滚到原始数据
        *user = original
        return nil, fmt.Errorf("更新用户失败: %w", err)
    }
    
    return user, nil
}

// 删除用户
func (s *UserService) DeleteUser(id int) error {
    if _, exists := s.users[id]; !exists {
        return fmt.Errorf("删除用户失败: %w", ErrUserNotFound)
    }
    
    delete(s.users, id)
    return nil
}

func main() {
    service := NewUserService()
    
    fmt.Println("🔧 用户管理服务错误处理示例")
    
    // 1. 创建有效用户
    user1, err := service.CreateUser("张三", "zhangsan@example.com", 25)
    if err != nil {
        fmt.Printf("❌ 创建用户失败: %v\n", err)
    } else {
        fmt.Printf("✅ 创建用户成功: %+v\n", user1)
    }
    
    // 2. 创建无效用户（邮箱重复）
    _, err = service.CreateUser("李四", "zhangsan@example.com", 30)
    if err != nil {
        fmt.Printf("❌ 创建用户失败: %v\n", err)
        
        // 检查特定错误类型
        var userErr UserError
        if errors.As(err, &userErr) && userErr.Code == "USER_EXISTS" {
            fmt.Println("   → 这是用户已存在错误")
        }
    }
    
    // 3. 创建无效用户（数据验证失败）
    _, err = service.CreateUser("", "invalid-email", 150)
    if err != nil {
        fmt.Printf("❌ 创建用户失败: %v\n", err)
    }
    
    // 4. 获取不存在的用户
    _, err = service.GetUser(999)
    if err != nil {
        fmt.Printf("❌ 获取用户失败: %v\n", err)
        
        // 检查错误类型
        if errors.Is(err, ErrUserNotFound) {
            fmt.Println("   → 这是用户不存在错误")
        }
    }
    
    // 5. 更新用户数据
    if user1 != nil {
        updatedUser, err := service.UpdateUser(user1.ID, "张三（已更新）", "zhangsan.new@example.com", 26)
        if err != nil {
            fmt.Printf("❌ 更新用户失败: %v\n", err)
        } else {
            fmt.Printf("✅ 更新用户成功: %+v\n", updatedUser)
        }
    }
    
    // 6. 删除用户
    if user1 != nil {
        err = service.DeleteUser(user1.ID)
        if err != nil {
            fmt.Printf("❌ 删除用户失败: %v\n", err)
        } else {
            fmt.Printf("✅ 删除用户成功\n")
        }
    }
}
```

## 本章小结

通过本章学习，你应该掌握：

### 核心概念
- **显式错误处理**：通过返回值处理错误，而非异常
- **error接口**：Go的错误处理核心，只有一个Error()方法
- **错误创建**：errors.New()、fmt.Errorf()等方式
- **错误包装**：使用%w动词保留原始错误信息

### 错误处理模式
1. **哨兵错误**：预定义的特定错误值
2. **自定义错误类型**：包含更丰富信息的错误结构
3. **错误包装**：保持错误链，便于调试
4. **错误检查**：使用errors.Is()和errors.As()

### 最佳实践
- **立即处理**：获得错误后立即检查和处理
- **向上传播**：在适当的层级处理错误，必要时向上传播
- **添加上下文**：包装错误时添加有用的上下文信息
- **错误分类**：使用不同的错误类型表示不同的错误条件

### Go错误处理的优势
- **明确性**：错误处理逻辑一目了然
- **强制性**：编译器强制检查错误处理
- **性能**：无异常机制的性能开销
- **调试友好**：错误信息明确，便于定位问题

### 错误处理策略
1. **快速失败**：在错误发生时立即返回
2. **重试机制**：对于临时性错误进行重试
3. **优雅降级**：提供备选方案或默认行为
4. **日志记录**：记录错误信息用于调试和监控

::: tip 练习建议
尝试实现一个简单的银行账户系统，练习各种错误处理模式：余额不足、账户不存在、无效操作等。
:::