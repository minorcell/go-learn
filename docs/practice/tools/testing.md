# 测试和基准测试 Testing

> 测试不是负担，而是信心的来源——让你的代码在任何时候都值得信赖

## 🤔 为什么Go把测试当作一等公民？

在很多编程语言中，测试往往是后加的功能，需要额外的框架和复杂的配置。但Go不同——**测试是Go语言设计的核心部分**。

想想这个设计哲学的深意：
- `go test`命令内置在工具链中
- 测试文件与源码文件并列存放
- 标准库提供了完整的测试支持
- 基准测试和示例代码都是原生支持

这不是偶然的设计选择，而是Go团队对软件质量的深度思考。他们相信：**好的软件应该从第一行代码开始就考虑测试**。

## 🎯 Go测试哲学

### 简单胜过复杂

::: details 示例：简单胜过复杂
```go
// 这就是一个完整的Go测试
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```
:::
没有复杂的注解，没有魔法方法，就是普通的Go函数。这种简单性让测试变得易写、易读、易维护。

### 测试就在身边

```
myproject/
├── calculator.go
├── calculator_test.go    # 测试文件就在旁边
├── user.go
└── user_test.go
```

测试文件与源码文件并列，这种物理上的接近性提醒开发者：**测试不是额外的工作，而是开发的一部分**。

## 🧪 基础测试实践

### 测试函数的命名规范

::: details 示例：测试函数的命名规范
```go
package calculator

import "testing"

// ✅ 标准的测试函数命名
func TestAdd(t *testing.T) {
    // 测试加法功能
}

func TestAddWithNegativeNumbers(t *testing.T) {
    // 测试负数加法
}

func TestDivide(t *testing.T) {
    // 测试除法功能
}

func TestDivideByZero(t *testing.T) {
    // 测试零除错误
}
```
:::
**命名原则**：
- 必须以`Test`开头
- 函数名应该清楚描述测试的功能
- 使用驼峰命名法
- 具体场景可以用描述性后缀

### 表格驱动测试（Table-Driven Tests）

Go社区最佳实践之一，用一个测试函数覆盖多个测试用例：

::: details 示例：表格驱动测试
```go
func TestAdd(t *testing.T) {
    testCases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed signs", -2, 3, 1},
        {"zero values", 0, 0, 0},
        {"large numbers", 1000000, 2000000, 3000000},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```
:::
**优势解析**：
- **全面性**：一次性测试多种场景
- **可读性**：测试用例一目了然
- **维护性**：添加新用例只需添加数据
- **诊断性**：`t.Run`提供子测试，失败时能精确定位

### 错误处理测试

Go的错误处理模式在测试中同样重要：

::: details 示例：错误处理测试
```go
func TestDivide(t *testing.T) {
    testCases := []struct {
        name        string
        a, b        float64
        expected    float64
        expectError bool
        errorMsg    string
    }{
        {"normal division", 10, 2, 5, false, ""},
        {"divide by zero", 10, 0, 0, true, "division by zero"},
        {"negative result", -10, 2, -5, false, ""},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := Divide(tc.a, tc.b)
            
            // 检查错误期望
            if tc.expectError {
                if err == nil {
                    t.Errorf("expected error but got none")
                    return
                }
                if !strings.Contains(err.Error(), tc.errorMsg) {
                    t.Errorf("expected error containing %q, got %q", 
                        tc.errorMsg, err.Error())
                }
                return
            }
            
            // 检查正常结果
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            
            if result != tc.expected {
                t.Errorf("Divide(%.2f, %.2f) = %.2f; want %.2f", 
                    tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```
:::
## 🏗️ 高级测试技术

### 测试辅助函数

将测试逻辑模块化，提高代码复用性：

::: details 示例：测试辅助函数
```go
// 测试辅助函数，不以Test开头
func assertAdd(t *testing.T, a, b, expected int) {
    t.Helper() // 标记为辅助函数，错误报告时显示调用者位置
    
    result := Add(a, b)
    if result != expected {
        t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
    }
}

func TestAddOperations(t *testing.T) {
    assertAdd(t, 2, 3, 5)
    assertAdd(t, -1, 1, 0)
    assertAdd(t, 0, 0, 0)
}
```
:::
### Setup和Teardown

管理测试的初始化和清理工作：

::: details 示例：Setup和Teardown
```go
func TestMain(m *testing.M) {
    // 全局设置
    setupGlobalResources()
    
    // 运行所有测试
    code := m.Run()
    
    // 全局清理
    teardownGlobalResources()
    
    os.Exit(code)
}

func TestUserService(t *testing.T) {
    // 测试级别的设置
    db := setupTestDatabase(t)
    defer cleanupTestDatabase(t, db)
    
    userService := NewUserService(db)
    
    t.Run("CreateUser", func(t *testing.T) {
        // 子测试逻辑
    })
    
    t.Run("DeleteUser", func(t *testing.T) {
        // 子测试逻辑
    })
}

func setupTestDatabase(t *testing.T) *sql.DB {
    t.Helper()
    // 创建测试数据库连接
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to create test database: %v", err)
    }
    return db
}
```
:::
### Mock和依赖注入

Go的接口系统让Mock变得自然而优雅：

::: details 示例：Mock和依赖注入
```go
// 定义接口
type UserRepository interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

// 生产环境实现
type DatabaseUserRepository struct {
    db *sql.DB
}

func (r *DatabaseUserRepository) GetUser(id int) (*User, error) {
    // 数据库查询逻辑
}

// 测试环境的Mock实现
type MockUserRepository struct {
    users map[int]*User
    error error
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    if m.error != nil {
        return nil, m.error
    }
    return m.users[id], nil
}

func (m *MockUserRepository) SaveUser(user *User) error {
    if m.error != nil {
        return m.error
    }
    m.users[user.ID] = user
    return nil
}

// 使用Mock进行测试
func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepository{
        users: make(map[int]*User),
    }
    
    service := NewUserService(mockRepo)
    
    t.Run("GetUser success", func(t *testing.T) {
        expectedUser := &User{ID: 1, Name: "Alice"}
        mockRepo.users[1] = expectedUser
        
        user, err := service.GetUser(1)
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        
        if user.Name != expectedUser.Name {
            t.Errorf("expected name %s, got %s", expectedUser.Name, user.Name)
        }
    })
    
    t.Run("GetUser error", func(t *testing.T) {
        mockRepo.error = errors.New("database connection failed")
        
        _, err := service.GetUser(2)
        if err == nil {
            t.Error("expected error but got none")
        }
    })
}
```
:::

## 🌐 Web和API测试

### HTTP测试

Go标准库的`httptest`包让Web测试变得简单：

::: details 示例：HTTP测试
```go
func TestUserHandler(t *testing.T) {
    // 创建测试用的HTTP服务器
    handler := NewUserHandler()
    server := httptest.NewServer(handler)
    defer server.Close()
    
    testCases := []struct {
        name           string
        method         string
        path           string
        body           string
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "get user success",
            method:         "GET",
            path:           "/users/1",
            expectedStatus: http.StatusOK,
            expectedBody:   `{"id":1,"name":"Alice"}`,
        },
        {
            name:           "user not found",
            method:         "GET",
            path:           "/users/999",
            expectedStatus: http.StatusNotFound,
            expectedBody:   `{"error":"user not found"}`,
        },
        {
            name:           "create user",
            method:         "POST",
            path:           "/users",
            body:           `{"name":"Bob"}`,
            expectedStatus: http.StatusCreated,
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            var req *http.Request
            var err error
            
            if tc.body != "" {
                req, err = http.NewRequest(tc.method, 
                    server.URL+tc.path, 
                    strings.NewReader(tc.body))
            } else {
                req, err = http.NewRequest(tc.method, 
                    server.URL+tc.path, nil)
            }
            
            if err != nil {
                t.Fatalf("failed to create request: %v", err)
            }
            
            req.Header.Set("Content-Type", "application/json")
            
            resp, err := http.DefaultClient.Do(req)
            if err != nil {
                t.Fatalf("request failed: %v", err)
            }
            defer resp.Body.Close()
            
            // 检查状态码
            if resp.StatusCode != tc.expectedStatus {
                t.Errorf("expected status %d, got %d", 
                    tc.expectedStatus, resp.StatusCode)
            }
            
            // 检查响应体
            if tc.expectedBody != "" {
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    t.Fatalf("failed to read response body: %v", err)
                }
                
                if strings.TrimSpace(string(body)) != tc.expectedBody {
                    t.Errorf("expected body %s, got %s", 
                        tc.expectedBody, string(body))
                }
            }
        })
    }
}
```
:::
### 使用testify增强测试体验

虽然Go内置测试足够强大，但`testify`库提供了更好的断言体验：

::: details 示例：使用testify增强测试体验
```go
import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/suite"
)

func TestUserService(t *testing.T) {
    service := NewUserService()
    
    // 基本断言
    user, err := service.CreateUser("Alice")
    require.NoError(t, err) // 失败时立即停止测试
    assert.Equal(t, "Alice", user.Name)
    assert.NotZero(t, user.ID)
    
    // 集合断言
    users, err := service.GetAllUsers()
    require.NoError(t, err)
    assert.Len(t, users, 1)
    assert.Contains(t, users, user)
}

// 测试套件模式
type UserServiceTestSuite struct {
    suite.Suite
    service *UserService
    db      *sql.DB
}

func (s *UserServiceTestSuite) SetupTest() {
    // 每个测试前的设置
    s.db = setupTestDB()
    s.service = NewUserService(s.db)
}

func (s *UserServiceTestSuite) TearDownTest() {
    // 每个测试后的清理
    s.db.Close()
}

func (s *UserServiceTestSuite) TestCreateUser() {
    user, err := s.service.CreateUser("Alice")
    s.Require().NoError(err)
    s.Equal("Alice", user.Name)
}

func TestUserServiceTestSuite(t *testing.T) {
    suite.Run(t, new(UserServiceTestSuite))
}
```
:::
## 📊 基准测试：性能的科学测量

### 基本基准测试

::: details 示例：基本基准测试
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := "hello" + "world"
        _ = result // 避免编译器优化
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        builder.WriteString("hello")
        builder.WriteString("world")
        result := builder.String()
        _ = result
    }
}
```
:::
### 运行基准测试

::: details 示例：运行基准测试
```bash
# 运行所有基准测试
go test -bench=.

# 运行特定基准测试
go test -bench=BenchmarkAdd

# 设置运行时间
go test -bench=. -benchtime=5s

# 内存分析
go test -bench=. -benchmem

# 输出示例：
# BenchmarkAdd-8                1000000000    0.25 ns/op
# BenchmarkStringConcat-8       500000000     3.2 ns/op     0 B/op   0 allocs/op
# BenchmarkStringBuilder-8      200000000     8.1 ns/op    32 B/op   1 allocs/op
```
:::
### 高级基准测试技术

::: details 示例：高级基准测试技术
```go
func BenchmarkMapOperations(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            m := make(map[int]int, size)
            
            // 预填充数据
            for i := 0; i < size; i++ {
                m[i] = i
            }
            
            b.ResetTimer() // 重置计时器，排除设置时间
            
            for i := 0; i < b.N; i++ {
                key := i % size
                _ = m[key] // 避免编译器优化
            }
        })
    }
}

// 内存池基准测试
func BenchmarkByteSliceWithoutPool(b *testing.B) {
    for i := 0; i < b.N; i++ {
        data := make([]byte, 1024)
        _ = data
    }
}

func BenchmarkByteSliceWithPool(b *testing.B) {
    pool := sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024)
        },
    }
    
    for i := 0; i < b.N; i++ {
        data := pool.Get().([]byte)
        pool.Put(data)
    }
}
```
:::
## 📈 测试覆盖率分析

### 生成覆盖率报告

::: details 示例：生成覆盖率报告
```bash
# 生成覆盖率文件
go test -coverprofile=coverage.out ./...

# 查看总体覆盖率
go tool cover -func=coverage.out

# 生成HTML报告
go tool cover -html=coverage.out -o coverage.html

# 查看特定包的覆盖率
go test -cover ./mypackage

# 输出示例：
# github.com/myproject/calculator/add.go:5:   Add             100.0%
# github.com/myproject/calculator/divide.go:8: Divide         85.7%
# total:                                       (statements)   92.3%
```
:::
### 覆盖率最佳实践

::: details 示例：覆盖率最佳实践
```go
// ❌ 为了覆盖率而写的无意义测试
func TestGetUserName(t *testing.T) {
    user := User{Name: "Alice"}
    name := user.GetName()
    if name != "Alice" {
        t.Error("expected Alice")
    }
}

// ✅ 有意义的测试，关注边界条件和错误情况
func TestUserValidation(t *testing.T) {
    testCases := []struct {
        name      string
        user      User
        expectErr bool
    }{
        {"valid user", User{Name: "Alice", Age: 25}, false},
        {"empty name", User{Name: "", Age: 25}, true},
        {"negative age", User{Name: "Bob", Age: -1}, true},
        {"very long name", User{Name: strings.Repeat("a", 1000), Age: 25}, true},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            err := tc.user.Validate()
            hasErr := err != nil
            if hasErr != tc.expectErr {
                t.Errorf("expected error: %v, got error: %v", tc.expectErr, hasErr)
            }
        })
    }
}
```
:::
## 🔄 测试驱动开发（TDD）

### 红-绿-重构循环

::: details 示例：红-绿-重构循环
```go
// 1. 红：先写测试（会失败）
func TestCalculateDiscount(t *testing.T) {
    testCases := []struct {
        amount   float64
        userType string
        expected float64
    }{
        {100, "regular", 100},
        {100, "premium", 90},
        {100, "vip", 80},
    }
    
    for _, tc := range testCases {
        result := CalculateDiscount(tc.amount, tc.userType)
        if result != tc.expected {
            t.Errorf("CalculateDiscount(%.2f, %s) = %.2f; want %.2f",
                tc.amount, tc.userType, result, tc.expected)
        }
    }
}

// 2. 绿：写最简单的实现让测试通过
func CalculateDiscount(amount float64, userType string) float64 {
    switch userType {
    case "premium":
        return amount * 0.9
    case "vip":
        return amount * 0.8
    default:
        return amount
    }
}

// 3. 重构：优化代码结构，保持测试通过
type DiscountRule struct {
    UserType string
    Rate     float64
}

var discountRules = []DiscountRule{
    {"premium", 0.9},
    {"vip", 0.8},
}

func CalculateDiscount(amount float64, userType string) float64 {
    for _, rule := range discountRules {
        if rule.UserType == userType {
            return amount * rule.Rate
        }
    }
    return amount
}
```
:::
## 🛠️ 测试工具生态

### 常用测试库

::: details 示例：常用测试库
```go
// 1. testify - 断言和Mock框架
import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
)

// 2. GoMock - 自动生成Mock
//go:generate mockgen -source=user.go -destination=mocks/user_mock.go

// 3. httpexpect - HTTP API测试
import "github.com/gavv/httpexpect/v2"

func TestUserAPI(t *testing.T) {
    e := httpexpect.New(t, "http://localhost:8080")
    
    e.GET("/users/1").
        Expect().
        Status(http.StatusOK).
        JSON().Object().
        Value("name").String().Equal("Alice")
}

// 4. ginkgo - BDD风格测试框架
import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
    Context("when adding numbers", func() {
        It("should return the sum", func() {
            result := Add(2, 3)
            Expect(result).To(Equal(5))
        })
    })
})
```
:::
### CI/CD集成

::: details 示例：CI/CD集成
```yaml
# .github/workflows/test.yml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        go-version: [1.20, 1.21]
        
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: ${{ matrix.go-version }}
        
    - name: Run tests
      run: |
        go test -v -race -coverprofile=coverage.out ./...
        
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
        
    - name: Run benchmarks
      run: go test -bench=. -benchmem ./...
```
:::
## 🎯 测试最佳实践总结

### 1. 测试命名和组织

**好的测试名称**：
- `TestUserService_CreateUser_Success`
- `TestUserService_CreateUser_DuplicateEmail_ReturnsError`
- `TestCalculateDiscount_PremiumUser_Returns10PercentOff`

**测试文件组织**：
```
user/
├── user.go
├── user_test.go           # 单元测试
├── user_integration_test.go # 集成测试
└── user_benchmark_test.go   # 基准测试
```

### 2. 测试金字塔

```
        E2E Tests (少量)
       /               \
   Integration Tests (适量)
  /                       \
Unit Tests (大量，快速，独立)
```

### 3. 测试原则

**FIRST原则**：
- **Fast**：测试应该快速运行
- **Independent**：测试之间不应该相互依赖
- **Repeatable**：测试结果应该可重复
- **Self-validating**：测试应该有明确的通过/失败结果
- **Timely**：测试应该及时编写

### 4. 常见陷阱和解决方案

::: details 示例：常见陷阱和解决方案
```go
// ❌ 测试依赖于外部状态
func TestGetCurrentTime(t *testing.T) {
    result := GetCurrentTime()
    expected := time.Now()
    if result != expected {
        t.Error("time mismatch") // 这个测试不稳定
    }
}

// ✅ 依赖注入，控制外部依赖
type TimeProvider interface {
    Now() time.Time
}

func TestGetCurrentTime(t *testing.T) {
    mockTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
    mockProvider := &MockTimeProvider{fixedTime: mockTime}
    
    service := NewTimeService(mockProvider)
    result := service.GetCurrentTime()
    
    if result != mockTime {
        t.Errorf("expected %v, got %v", mockTime, result)
    }
}
```
:::
---

💡 **测试心法**：写测试不是为了完成任务，而是为了让代码更可靠。好的测试是活文档，它们告诉未来的你（和团队成员）代码应该如何工作。

**下一步**：学习[性能分析工具](/practice/tools/profiling)，掌握Go程序的性能优化技术。
