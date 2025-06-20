package main

import (
	"fmt"
	"math"
	"time"
)

/*
05_structs_methods.go - Go语言基础：结构体和方法
学习内容：
1. 结构体定义和使用
2. 方法定义
3. 值接收者 vs 指针接收者
4. 嵌套结构体
5. 构造函数模式
*/

// 定义基本结构体
type Person struct {
	Name string
	Age  int
	City string
}

// 定义带有私有字段的结构体
type BankAccount struct {
	ownerName string  // 小写字母开头，私有字段
	balance   float64 // 私有字段
	AccountID string  // 公开字段
}

// 为Person定义方法（值接收者）
func (p Person) Introduce() string {
	return fmt.Sprintf("你好，我是%s，今年%d岁，来自%s", p.Name, p.Age, p.City)
}

// 为Person定义方法（值接收者）
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// 为BankAccount定义方法（指针接收者）
func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
		fmt.Printf("存入 %.2f 元，当前余额: %.2f 元\n", amount, ba.balance)
	}
}

// 为BankAccount定义方法（指针接收者）
func (ba *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && ba.balance >= amount {
		ba.balance -= amount
		fmt.Printf("取出 %.2f 元，当前余额: %.2f 元\n", amount, ba.balance)
		return true
	}
	fmt.Printf("余额不足或金额无效，当前余额: %.2f 元\n", ba.balance)
	return false
}

// 获取余额（值接收者，因为不修改数据）
func (ba BankAccount) GetBalance() float64 {
	return ba.balance
}

// 获取账户信息
func (ba BankAccount) GetAccountInfo() string {
	return fmt.Sprintf("账户ID: %s, 所有者: %s, 余额: %.2f",
		ba.AccountID, ba.ownerName, ba.balance)
}

// 构造函数模式
func NewBankAccount(ownerName, accountID string, initialBalance float64) *BankAccount {
	return &BankAccount{
		ownerName: ownerName,
		AccountID: accountID,
		balance:   initialBalance,
	}
}

func NewPerson(name string, age int, city string) Person {
	return Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// 嵌套结构体示例
type Address struct {
	Street   string
	City     string
	Province string
	ZipCode  string
}

type Employee struct {
	Person   // 嵌入Person结构体
	Address  // 嵌入Address结构体
	ID       int
	Salary   float64
	HireDate time.Time
}

// 为Employee定义方法
func (e Employee) GetFullInfo() string {
	return fmt.Sprintf("员工信息:\n  姓名: %s\n  年龄: %d\n  员工ID: %d\n  薪资: %.2f\n  地址: %s, %s, %s",
		e.Name, e.Age, e.ID, e.Salary, e.Street, e.City, e.Province)
}

func (e *Employee) GiveSalaryRaise(percentage float64) {
	e.Salary *= (1 + percentage/100)
	fmt.Printf("%s的薪资调整为: %.2f (涨幅: %.1f%%)\n", e.Name, e.Salary, percentage)
}

// 几何图形示例
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Circle的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("圆形(半径: %.2f)", c.Radius)
}

// Rectangle的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("矩形(宽: %.2f, 高: %.2f)", r.Width, r.Height)
}

func main() {
	fmt.Println("=== Go语言基础：结构体和方法 ===")

	// 1. 基本结构体使用
	fmt.Println("\n1. 基本结构体使用：")

	// 结构体的多种创建方式
	var p1 Person
	p1.Name = "张三"
	p1.Age = 25
	p1.City = "北京"

	p2 := Person{Name: "李四", Age: 30, City: "上海"}
	p3 := Person{"王五", 28, "广州"}    // 按字段顺序初始化
	p4 := NewPerson("赵六", 35, "深圳") // 使用构造函数

	fmt.Printf("p1: %+v\n", p1) // %+v 显示字段名
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p3: %v\n", p3)
	fmt.Printf("p4: %v\n", p4)

	// 2. 方法调用
	fmt.Println("\n2. 方法调用：")

	fmt.Println(p1.Introduce())
	fmt.Printf("%s 是否成年: %t\n", p1.Name, p1.IsAdult())
	fmt.Printf("%s 是否成年: %t\n", p2.Name, p2.IsAdult())

	// 3. 指针和值接收者
	fmt.Println("\n3. 指针和值接收者：")

	// 创建银行账户
	account1 := NewBankAccount("张三", "ACC001", 1000.0)
	account2 := &BankAccount{
		ownerName: "李四",
		AccountID: "ACC002",
		balance:   500.0,
	}

	fmt.Println(account1.GetAccountInfo())
	fmt.Println(account2.GetAccountInfo())

	// 进行一些银行操作
	account1.Deposit(500)
	account1.Withdraw(200)
	account1.Withdraw(2000) // 余额不足

	fmt.Printf("最终余额: %.2f\n", account1.GetBalance())

	// 4. 嵌套结构体
	fmt.Println("\n4. 嵌套结构体：")

	emp := Employee{
		Person: Person{
			Name: "王小明",
			Age:  29,
			City: "杭州",
		},
		Address: Address{
			Street:   "西湖大道123号",
			City:     "杭州",
			Province: "浙江",
			ZipCode:  "310000",
		},
		ID:       1001,
		Salary:   8000.0,
		HireDate: time.Now(),
	}

	fmt.Println(emp.GetFullInfo())

	// 访问嵌入字段
	fmt.Printf("员工姓名: %s\n", emp.Name)        // 直接访问Person的字段
	fmt.Printf("员工城市: %s\n", emp.Person.City) // 或者明确指定
	fmt.Printf("地址街道: %s\n", emp.Street)      // 直接访问Address的字段

	// 调用嵌入类型的方法
	fmt.Println(emp.Introduce()) // 调用Person的方法

	// 修改员工信息
	emp.GiveSalaryRaise(10.0) // 涨薪10%

	// 5. 几何图形示例
	fmt.Println("\n5. 几何图形示例：")

	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	fmt.Printf("%s - 面积: %.2f, 周长: %.2f\n",
		circle.String(), circle.Area(), circle.Perimeter())
	fmt.Printf("%s - 面积: %.2f, 周长: %.2f\n",
		rectangle.String(), rectangle.Area(), rectangle.Perimeter())

	// 6. 结构体切片
	fmt.Println("\n6. 结构体切片：")

	people := []Person{
		{"Alice", 25, "北京"},
		{"Bob", 30, "上海"},
		{"Charlie", 22, "广州"},
		{"David", 17, "深圳"},
	}

	fmt.Println("所有人员信息:")
	for i, person := range people {
		fmt.Printf("%d. %s\n", i+1, person.Introduce())
	}

	// 统计成年人数
	adultCount := 0
	for _, person := range people {
		if person.IsAdult() {
			adultCount++
		}
	}
	fmt.Printf("成年人数量: %d/%d\n", adultCount, len(people))
}
