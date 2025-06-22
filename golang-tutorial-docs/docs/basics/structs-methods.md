# 结构体和方法

结构体是Go语言中用户自定义类型的主要方式，结合方法可以实现面向对象编程的思想。

## 本章内容

- 结构体的定义和使用
- 方法的定义和接收者
- 结构体嵌入和组合
- 构造函数模式
- 面向对象编程实践

## 结构体基础

### 结构体定义和初始化

```go
package main

import "fmt"

// 定义学生结构体
type Student struct {
    ID      int
    Name    string
    Age     int
    Email   string
    Scores  []int
}

// 定义地址结构体
type Address struct {
    Street   string
    City     string
    Province string
    ZipCode  string
}

// 定义人员结构体
type Person struct {
    Name    string
    Age     int
    Address Address // 嵌套结构体
}

func main() {
    // 方式1：零值初始化
    var s1 Student
    fmt.Printf("零值学生: %+v\n", s1)
    
    // 方式2：字面量初始化
    s2 := Student{
        ID:     1001,
        Name:   "张三",
        Age:    20,
        Email:  "zhangsan@example.com",
        Scores: []int{85, 92, 78},
    }
    fmt.Printf("完整初始化: %+v\n", s2)
    
    // 方式3：部分初始化
    s3 := Student{
        ID:   1002,
        Name: "李四",
        Age:  19,
    }
    fmt.Printf("部分初始化: %+v\n", s3)
    
    // 方式4：按顺序初始化（不推荐）
    s4 := Student{1003, "王五", 21, "wangwu@example.com", []int{90, 88, 95}}
    fmt.Printf("顺序初始化: %+v\n", s4)
    
    // 嵌套结构体初始化
    p1 := Person{
        Name: "赵六",
        Age:  25,
        Address: Address{
            Street:   "长安街1号",
            City:     "北京",
            Province: "北京市",
            ZipCode:  "100000",
        },
    }
    fmt.Printf("嵌套结构体: %+v\n", p1)
    
    // 访问和修改字段
    fmt.Printf("学生姓名: %s\n", s2.Name)
    fmt.Printf("学生年龄: %d\n", s2.Age)
    
    s2.Age = 21
    s2.Scores = append(s2.Scores, 88)
    fmt.Printf("修改后: %+v\n", s2)
    
    // 访问嵌套字段
    fmt.Printf("地址: %s, %s\n", p1.Address.City, p1.Address.Province)
}
```

### 结构体指针

```go
package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Pages  int
    Price  float64
}

// 修改结构体的函数（值传递）
func updateBookValue(b Book) {
    b.Price = b.Price * 1.1 // 10% 涨价
    fmt.Printf("函数内价格: %.2f\n", b.Price)
}

// 修改结构体的函数（指针传递）
func updateBookPointer(b *Book) {
    b.Price = b.Price * 1.1 // 10% 涨价
    fmt.Printf("函数内价格: %.2f\n", b.Price)
}

func main() {
    // 创建结构体
    book := Book{
        Title:  "Go语言编程",
        Author: "张三",
        Pages:  300,
        Price:  59.90,
    }
    
    fmt.Printf("原始价格: %.2f\n", book.Price)
    
    // 值传递 - 不会修改原结构体
    updateBookValue(book)
    fmt.Printf("值传递后: %.2f\n", book.Price)
    
    // 指针传递 - 会修改原结构体
    updateBookPointer(&book)
    fmt.Printf("指针传递后: %.2f\n", book.Price)
    
    // 使用 new 创建结构体指针
    bookPtr := new(Book)
    bookPtr.Title = "Go高级编程"
    bookPtr.Author = "李四"
    bookPtr.Pages = 450
    bookPtr.Price = 79.90
    
    fmt.Printf("new创建: %+v\n", *bookPtr)
    
    // 结构体指针的简化访问
    fmt.Printf("标题: %s\n", bookPtr.Title) // 自动解引用
    fmt.Printf("价格: %.2f\n", (*bookPtr).Price) // 显式解引用
}
```

## 方法

### 方法定义和接收者

```go
package main

import (
    "fmt"
    "math"
)

// 圆形结构体
type Circle struct {
    Radius float64
}

// 矩形结构体
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle的方法 - 值接收者
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}

// Circle的方法 - 指针接收者
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func (c *Circle) String() string {
    return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}

// Rectangle的方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func (r *Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

func main() {
    // 创建圆形
    circle := Circle{Radius: 5.0}
    fmt.Printf("圆形: %s\n", circle.String())
    fmt.Printf("面积: %.2f\n", circle.Area())
    fmt.Printf("周长: %.2f\n", circle.Circumference())
    
    // 缩放圆形
    circle.Scale(2.0)
    fmt.Printf("缩放后: %s\n", circle.String())
    fmt.Printf("新面积: %.2f\n", circle.Area())
    
    // 创建矩形
    rectangle := Rectangle{Width: 10.0, Height: 6.0}
    fmt.Printf("\n矩形: %s\n", rectangle.String())
    fmt.Printf("面积: %.2f\n", rectangle.Area())
    fmt.Printf("周长: %.2f\n", rectangle.Perimeter())
    
    // 缩放矩形
    rectangle.Scale(1.5)
    fmt.Printf("缩放后: %s\n", rectangle.String())
    fmt.Printf("新面积: %.2f\n", rectangle.Area())
    
    // 指针和值的方法调用
    circlePtr := &Circle{Radius: 3.0}
    fmt.Printf("\n指针调用: %s\n", circlePtr.String()) // 自动解引用
    fmt.Printf("面积: %.2f\n", circlePtr.Area())        // 自动解引用
    
    circleValue := Circle{Radius: 4.0}
    circleValue.Scale(2.0) // 自动取地址
    fmt.Printf("值调用指针方法: %s\n", circleValue.String())
}
```

### 方法集和接收者选择

```go
package main

import "fmt"

type Counter struct {
    value int
}

// 值接收者 - 不会修改原结构体
func (c Counter) Get() int {
    return c.value
}

func (c Counter) Add(n int) int {
    c.value += n // 只修改副本
    return c.value
}

// 指针接收者 - 会修改原结构体
func (c *Counter) Set(value int) {
    c.value = value
}

func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Decrement() {
    c.value--
}

func (c *Counter) AddValue(n int) {
    c.value += n
}

func (c *Counter) String() string {
    return fmt.Sprintf("Counter{value: %d}", c.value)
}

func main() {
    // 演示值接收者vs指针接收者
    counter := Counter{value: 10}
    fmt.Printf("初始: %s\n", counter.String())
    
    // 值接收者方法
    result := counter.Add(5)
    fmt.Printf("Add(5)返回: %d\n", result)
    fmt.Printf("Add后counter: %s\n", counter.String()) // 原值未变
    
    // 指针接收者方法
    counter.Set(20)
    fmt.Printf("Set(20)后: %s\n", counter.String())
    
    counter.Increment()
    fmt.Printf("Increment后: %s\n", counter.String())
    
    counter.AddValue(8)
    fmt.Printf("AddValue(8)后: %s\n", counter.String())
    
    // 方法链调用
    counter.Set(0)
    counter.Increment()
    counter.Increment()
    counter.AddValue(5)
    fmt.Printf("链式调用后: %s\n", counter.String())
    
    // 演示接收者类型的重要性
    fmt.Println("\n接收者类型演示:")
    
    // 值类型变量
    c1 := Counter{value: 100}
    c1.Increment() // 自动转换为 (&c1).Increment()
    fmt.Printf("值变量调用指针方法: %s\n", c1.String())
    
    // 指针类型变量
    c2 := &Counter{value: 200}
    fmt.Printf("指针变量调用值方法: %d\n", c2.Get()) // 自动解引用
}
```

## 结构体组合和嵌入

### 结构体嵌入

```go
package main

import "fmt"

// 基础结构体
type Animal struct {
    Name    string
    Species string
    Age     int
}

// Animal的方法
func (a Animal) Speak() string {
    return fmt.Sprintf("%s makes a sound", a.Name)
}

func (a Animal) Info() string {
    return fmt.Sprintf("%s is a %d-year-old %s", a.Name, a.Age, a.Species)
}

// 嵌入Animal的Dog结构体
type Dog struct {
    Animal // 匿名嵌入
    Breed  string
    Owner  string
}

// Dog的特有方法
func (d Dog) Bark() string {
    return fmt.Sprintf("%s barks: Woof! Woof!", d.Name)
}

// 重写Animal的方法
func (d Dog) Speak() string {
    return d.Bark()
}

// Cat结构体
type Cat struct {
    Animal
    Color     string
    Indoor    bool
}

func (c Cat) Meow() string {
    return fmt.Sprintf("%s meows: Meow~", c.Name)
}

func (c Cat) Speak() string {
    return c.Meow()
}

// 鸟类
type Bird struct {
    Animal
    CanFly    bool
    WingSpan  float64
}

func (b Bird) Chirp() string {
    return fmt.Sprintf("%s chirps: Tweet tweet!", b.Name)
}

func (b Bird) Speak() string {
    return b.Chirp()
}

func (b Bird) Fly() string {
    if b.CanFly {
        return fmt.Sprintf("%s is flying with %.1fm wingspan", b.Name, b.WingSpan)
    }
    return fmt.Sprintf("%s cannot fly", b.Name)
}

func main() {
    // 创建狗
    dog := Dog{
        Animal: Animal{
            Name:    "旺财",
            Species: "犬",
            Age:     3,
        },
        Breed: "金毛",
        Owner: "张三",
    }
    
    // 访问嵌入字段
    fmt.Printf("狗的名字: %s\n", dog.Name) // 直接访问Animal的字段
    fmt.Printf("狗的信息: %s\n", dog.Info()) // 调用Animal的方法
    fmt.Printf("狗叫声: %s\n", dog.Speak()) // 调用重写的方法
    fmt.Printf("狗的品种: %s, 主人: %s\n", dog.Breed, dog.Owner)
    
    // 创建猫
    cat := Cat{
        Animal: Animal{
            Name:    "咪咪",
            Species: "猫",
            Age:     2,
        },
        Color:  "橙色",
        Indoor: true,
    }
    
    fmt.Printf("\n猫的信息: %s\n", cat.Info())
    fmt.Printf("猫叫声: %s\n", cat.Speak())
    fmt.Printf("猫的颜色: %s, 室内猫: %t\n", cat.Color, cat.Indoor)
    
    // 创建鸟
    bird := Bird{
        Animal: Animal{
            Name:    "小黄",
            Species: "金丝雀",
            Age:     1,
        },
        CanFly:   true,
        WingSpan: 0.3,
    }
    
    fmt.Printf("\n鸟的信息: %s\n", bird.Info())
    fmt.Printf("鸟叫声: %s\n", bird.Speak())
    fmt.Printf("飞行能力: %s\n", bird.Fly())
    
    // 演示方法集的继承
    animals := []Animal{dog.Animal, cat.Animal, bird.Animal}
    fmt.Println("\n所有动物信息:")
    for _, animal := range animals {
        fmt.Printf("- %s\n", animal.Info())
        fmt.Printf("  %s\n", animal.Speak()) // 调用原始的Speak方法
    }
}
```

### 组合vs继承

```go
package main

import "fmt"

// 引擎结构体
type Engine struct {
    Type       string
    Horsepower int
    FuelType   string
}

func (e Engine) Start() string {
    return fmt.Sprintf("%s engine started", e.Type)
}

func (e Engine) Stop() string {
    return fmt.Sprintf("%s engine stopped", e.Type)
}

func (e Engine) Info() string {
    return fmt.Sprintf("%s engine: %d HP, fuel: %s", e.Type, e.Horsepower, e.FuelType)
}

// 轮子结构体
type Wheels struct {
    Count int
    Size  string
    Type  string
}

func (w Wheels) Info() string {
    return fmt.Sprintf("%d %s %s wheels", w.Count, w.Size, w.Type)
}

// 车辆基础结构体
type Vehicle struct {
    Make   string
    Model  string
    Year   int
    Engine Engine // 组合
    Wheels Wheels // 组合
}

func (v Vehicle) Info() string {
    return fmt.Sprintf("%d %s %s", v.Year, v.Make, v.Model)
}

func (v Vehicle) Start() string {
    return fmt.Sprintf("%s: %s", v.Info(), v.Engine.Start())
}

func (v Vehicle) Stop() string {
    return fmt.Sprintf("%s: %s", v.Info(), v.Engine.Stop())
}

// 汽车 - 嵌入Vehicle
type Car struct {
    Vehicle
    Doors       int
    Sunroof     bool
    AirConditioning bool
}

func (c Car) OpenDoors() string {
    return fmt.Sprintf("Opening %d doors of %s", c.Doors, c.Info())
}

// 摩托车 - 嵌入Vehicle
type Motorcycle struct {
    Vehicle
    SidecarAttached bool
    HelmetRequired  bool
}

func (m Motorcycle) Wheelie() string {
    return fmt.Sprintf("%s is doing a wheelie!", m.Info())
}

// 卡车 - 嵌入Vehicle
type Truck struct {
    Vehicle
    LoadCapacity int // 载重量(kg)
    TrailerAttached bool
}

func (t Truck) Load(weight int) string {
    if weight > t.LoadCapacity {
        return fmt.Sprintf("Cannot load %dkg, exceeds capacity of %dkg", weight, t.LoadCapacity)
    }
    return fmt.Sprintf("Loaded %dkg cargo on %s", weight, t.Info())
}

func main() {
    // 创建汽车
    car := Car{
        Vehicle: Vehicle{
            Make:  "丰田",
            Model: "凯美瑞",
            Year:  2023,
            Engine: Engine{
                Type:       "V6",
                Horsepower: 300,
                FuelType:   "汽油",
            },
            Wheels: Wheels{
                Count: 4,
                Size:  "18英寸",
                Type:  "合金",
            },
        },
        Doors:           4,
        Sunroof:         true,
        AirConditioning: true,
    }
    
    fmt.Println("=== 汽车信息 ===")
    fmt.Printf("车辆: %s\n", car.Info())
    fmt.Printf("引擎: %s\n", car.Engine.Info())
    fmt.Printf("轮子: %s\n", car.Wheels.Info())
    fmt.Printf("车门: %d, 天窗: %t, 空调: %t\n", car.Doors, car.Sunroof, car.AirConditioning)
    fmt.Printf("启动: %s\n", car.Start())
    fmt.Printf("开门: %s\n", car.OpenDoors())
    fmt.Printf("停车: %s\n", car.Stop())
    
    // 创建摩托车
    motorcycle := Motorcycle{
        Vehicle: Vehicle{
            Make:  "哈雷",
            Model: "Street 750",
            Year:  2023,
            Engine: Engine{
                Type:       "V-Twin",
                Horsepower: 53,
                FuelType:   "汽油",
            },
            Wheels: Wheels{
                Count: 2,
                Size:  "17英寸",
                Type:  "运动型",
            },
        },
        SidecarAttached: false,
        HelmetRequired:  true,
    }
    
    fmt.Println("\n=== 摩托车信息 ===")
    fmt.Printf("车辆: %s\n", motorcycle.Info())
    fmt.Printf("引擎: %s\n", motorcycle.Engine.Info())
    fmt.Printf("轮子: %s\n", motorcycle.Wheels.Info())
    fmt.Printf("边车: %t, 头盔: %t\n", motorcycle.SidecarAttached, motorcycle.HelmetRequired)
    fmt.Printf("启动: %s\n", motorcycle.Start())
    fmt.Printf("特技: %s\n", motorcycle.Wheelie())
    
    // 创建卡车
    truck := Truck{
        Vehicle: Vehicle{
            Make:  "沃尔沃",
            Model: "FH16",
            Year:  2023,
            Engine: Engine{
                Type:       "柴油",
                Horsepower: 750,
                FuelType:   "柴油",
            },
            Wheels: Wheels{
                Count: 18,
                Size:  "22.5英寸",
                Type:  "载重型",
            },
        },
        LoadCapacity:    25000,
        TrailerAttached: true,
    }
    
    fmt.Println("\n=== 卡车信息 ===")
    fmt.Printf("车辆: %s\n", truck.Info())
    fmt.Printf("载重: %dkg, 拖车: %t\n", truck.LoadCapacity, truck.TrailerAttached)
    fmt.Printf("装货: %s\n", truck.Load(20000))
    fmt.Printf("超载测试: %s\n", truck.Load(30000))
}
```

## 构造函数模式

### 构造函数和工厂函数

```go
package main

import (
    "fmt"
    "time"
)

// 用户结构体
type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
    IsActive  bool
}

// 简单构造函数
func NewUser(username, email string) *User {
    return &User{
        ID:        generateID(), // 假设有ID生成函数
        Username:  username,
        Email:     email,
        CreatedAt: time.Now(),
        IsActive:  true,
    }
}

// 模拟ID生成
var idCounter = 1000
func generateID() int {
    idCounter++
    return idCounter
}

// 用户方法
func (u *User) Activate() {
    u.IsActive = true
}

func (u *User) Deactivate() {
    u.IsActive = false
}

func (u *User) String() string {
    status := "inactive"
    if u.IsActive {
        status = "active"
    }
    return fmt.Sprintf("User{ID: %d, Username: %s, Email: %s, Status: %s, Created: %s}",
        u.ID, u.Username, u.Email, status, u.CreatedAt.Format("2006-01-02"))
}

// 配置结构体
type DatabaseConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    Database string
    MaxConns int
    Timeout  time.Duration
}

// 使用函数选项模式的构造函数
type DBOption func(*DatabaseConfig)

func WithHost(host string) DBOption {
    return func(config *DatabaseConfig) {
        config.Host = host
    }
}

func WithPort(port int) DBOption {
    return func(config *DatabaseConfig) {
        config.Port = port
    }
}

func WithCredentials(username, password string) DBOption {
    return func(config *DatabaseConfig) {
        config.Username = username
        config.Password = password
    }
}

func WithDatabase(database string) DBOption {
    return func(config *DatabaseConfig) {
        config.Database = database
    }
}

func WithMaxConnections(maxConns int) DBOption {
    return func(config *DatabaseConfig) {
        config.MaxConns = maxConns
    }
}

func WithTimeout(timeout time.Duration) DBOption {
    return func(config *DatabaseConfig) {
        config.Timeout = timeout
    }
}

// 数据库配置构造函数
func NewDatabaseConfig(options ...DBOption) *DatabaseConfig {
    // 默认配置
    config := &DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        Username: "root",
        Password: "",
        Database: "test",
        MaxConns: 10,
        Timeout:  30 * time.Second,
    }
    
    // 应用选项
    for _, option := range options {
        option(config)
    }
    
    return config
}

func (db *DatabaseConfig) String() string {
    return fmt.Sprintf("Database{Host: %s:%d, User: %s, DB: %s, MaxConns: %d, Timeout: %v}",
        db.Host, db.Port, db.Username, db.Database, db.MaxConns, db.Timeout)
}

// 银行账户示例
type BankAccount struct {
    accountNumber string
    holderName    string
    balance       float64
    accountType   string
    isLocked      bool
}

// 私有构造函数（通过首字母小写）
func newBankAccount(number, holder, accountType string, initialBalance float64) *BankAccount {
    return &BankAccount{
        accountNumber: number,
        holderName:    holder,
        balance:       initialBalance,
        accountType:   accountType,
        isLocked:      false,
    }
}

// 公开的工厂函数
func CreateSavingsAccount(holder string, initialDeposit float64) *BankAccount {
    accountNumber := fmt.Sprintf("SAV-%d", generateID())
    return newBankAccount(accountNumber, holder, "储蓄账户", initialDeposit)
}

func CreateCheckingAccount(holder string) *BankAccount {
    accountNumber := fmt.Sprintf("CHK-%d", generateID())
    return newBankAccount(accountNumber, holder, "支票账户", 0.0)
}

// 银行账户方法
func (ba *BankAccount) Deposit(amount float64) error {
    if ba.isLocked {
        return fmt.Errorf("账户已锁定")
    }
    if amount <= 0 {
        return fmt.Errorf("存款金额必须大于0")
    }
    ba.balance += amount
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if ba.isLocked {
        return fmt.Errorf("账户已锁定")
    }
    if amount <= 0 {
        return fmt.Errorf("取款金额必须大于0")
    }
    if amount > ba.balance {
        return fmt.Errorf("余额不足")
    }
    ba.balance -= amount
    return nil
}

func (ba *BankAccount) GetBalance() float64 {
    return ba.balance
}

func (ba *BankAccount) Lock() {
    ba.isLocked = true
}

func (ba *BankAccount) Unlock() {
    ba.isLocked = false
}

func (ba *BankAccount) String() string {
    status := "正常"
    if ba.isLocked {
        status = "锁定"
    }
    return fmt.Sprintf("账户{号码: %s, 持有人: %s, 类型: %s, 余额: %.2f, 状态: %s}",
        ba.accountNumber, ba.holderName, ba.accountType, ba.balance, status)
}

func main() {
    // 简单构造函数示例
    fmt.Println("=== 用户创建 ===")
    user1 := NewUser("alice", "alice@example.com")
    user2 := NewUser("bob", "bob@example.com")
    
    fmt.Printf("用户1: %s\n", user1.String())
    fmt.Printf("用户2: %s\n", user2.String())
    
    user1.Deactivate()
    fmt.Printf("停用后: %s\n", user1.String())
    
    // 函数选项模式示例
    fmt.Println("\n=== 数据库配置 ===")
    
    // 使用默认配置
    defaultDB := NewDatabaseConfig()
    fmt.Printf("默认配置: %s\n", defaultDB.String())
    
    // 自定义配置
    customDB := NewDatabaseConfig(
        WithHost("192.168.1.100"),
        WithPort(5432),
        WithCredentials("admin", "secret123"),
        WithDatabase("production"),
        WithMaxConnections(50),
        WithTimeout(60*time.Second),
    )
    fmt.Printf("自定义配置: %s\n", customDB.String())
    
    // 部分自定义
    devDB := NewDatabaseConfig(
        WithHost("dev.example.com"),
        WithDatabase("development"),
        WithMaxConnections(5),
    )
    fmt.Printf("开发配置: %s\n", devDB.String())
    
    // 工厂函数示例
    fmt.Println("\n=== 银行账户 ===")
    
    savings := CreateSavingsAccount("张三", 1000.0)
    checking := CreateCheckingAccount("李四")
    
    fmt.Printf("储蓄账户: %s\n", savings.String())
    fmt.Printf("支票账户: %s\n", checking.String())
    
    // 操作账户
    savings.Deposit(500.0)
    savings.Withdraw(200.0)
    fmt.Printf("操作后储蓄账户: %s\n", savings.String())
    
    checking.Deposit(1500.0)
    checking.Lock()
    err := checking.Withdraw(100.0)
    if err != nil {
        fmt.Printf("取款失败: %v\n", err)
    }
    fmt.Printf("锁定的支票账户: %s\n", checking.String())
    
    checking.Unlock()
    checking.Withdraw(100.0)
    fmt.Printf("解锁后支票账户: %s\n", checking.String())
}
```

## 本章小结

在这一章中，我们学习了Go语言的结构体和方法：

### 结构体基础
- **类型定义** - struct关键字定义自定义类型
- **字段访问** - 点号操作符访问字段
- **初始化方式** - 字面量、零值、部分初始化
- **指针操作** - 结构体指针和自动解引用

### 方法系统
- **接收者类型** - 值接收者vs指针接收者
- **方法定义** - func (receiver Type) methodName()
- **方法调用** - 自动地址转换和解引用
- **方法集** - 类型可调用的方法集合

### 结构体组合
- **嵌入** - 匿名字段实现继承效果
- **组合** - 包含其他结构体作为字段
- **方法重写** - 重新定义嵌入类型的方法
- **字段提升** - 直接访问嵌入类型的字段

### 设计模式
- **构造函数** - NewXxx函数创建实例
- **函数选项** - 灵活的配置模式
- **工厂函数** - 封装创建逻辑
- **私有字段** - 通过首字母控制可见性

### 最佳实践
- 优先使用组合而非继承
- 根据是否修改状态选择接收者类型
- 提供合理的构造函数
- 遵循Go的命名约定
