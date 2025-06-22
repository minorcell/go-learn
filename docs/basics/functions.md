---
title: 函数
description: 学习Go语言的函数定义、参数传递、返回值和高级特性
---

# 函数

函数是组织代码的基本单位，让程序更加模块化和可重用。Go语言的函数设计简洁而强大，支持多返回值、命名返回值、变参等特性。

## 本章内容

- 函数的定义和调用方式
- 参数传递：值传递vs指针传递
- 多返回值和命名返回值
- 高级特性：匿名函数、闭包、递归
- 函数作为"一等公民"的应用

## 基本函数

### 函数定义语法

Go语言函数的基本语法为：`func 函数名(参数列表) (返回值列表) { 函数体 }`

```go
// 无参数无返回值
func sayHello() {
    fmt.Println("Hello, World!")
}

// 有参数无返回值
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// 有参数有返回值
func add(a, b int) int {
    return a + b
}

// 多个返回值
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("除数不能为零")
    }
    return a / b, nil
}
```

**函数特点**：
- 函数名首字母大写表示导出（public），小写表示私有（private）
- 参数类型写在参数名后面
- 多个相同类型参数可以简写：`func add(a, b int)`
- 支持多返回值，这是Go的特色功能

### 函数调用

```go
func main() {
    sayHello()                    // 调用无参函数
    greet("Alice")               // 传递参数
    
    sum := add(3, 5)             // 接收单个返回值
    fmt.Printf("3 + 5 = %d\n", sum)
    
    result, err := divide(10, 3) // 接收多个返回值
    if err != nil {
        fmt.Printf("错误: %v\n", err)
    } else {
        fmt.Printf("10 / 3 = %.2f\n", result)
    }
}
```

### 命名返回值

Go语言支持为返回值命名，让函数更加自文档化：

```go
// 命名返回值
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // 自动返回命名的变量
}

// 实用示例：分析学生成绩
func analyzeScore(score int) (grade string, passed bool, message string) {
    switch {
    case score >= 90:
        grade, passed, message = "A", true, "优秀"
    case score >= 80:
        grade, passed, message = "B", true, "良好"
    case score >= 60:
        grade, passed, message = "C", true, "及格"
    default:
        grade, passed, message = "F", false, "不及格"
    }
    return
}
```

**命名返回值的优势**：
- 增强代码可读性，清楚表达函数的输出
- 自动初始化为零值
- 可以在函数体中直接赋值和修改
- `return` 语句可以省略返回值

::: tip 使用建议
命名返回值适合返回值含义明确的函数，但不要滥用，简单函数直接返回即可。
:::

## 参数传递

理解Go语言的参数传递机制对编写高效代码很重要。

### 值传递 vs 指针传递

```go
// 值传递 - 不会修改原始值
func modifyValue(x int) {
    x = x * 2
    fmt.Printf("函数内部 x = %d\n", x)
}

// 指针传递 - 会修改原始值
func modifyPointer(x *int) {
    *x = *x * 2
    fmt.Printf("函数内部 *x = %d\n", *x)
}

func main() {
    a := 10
    fmt.Printf("修改前 a = %d\n", a)
    modifyValue(a)
    fmt.Printf("修改后 a = %d\n", a)     // 值不变，仍为10
    
    b := 10  
    fmt.Printf("修改前 b = %d\n", b)
    modifyPointer(&b)
    fmt.Printf("修改后 b = %d\n", b)     // 值改变，变为20
}
```

**参数传递规则**：
- **基本类型**（int、float、bool、string）：值传递，不会修改原始值
- **指针**：传递内存地址，可以修改原始值
- **切片、映射、通道**：引用类型，传递引用，可以修改内容
- **数组、结构体**：值传递，会复制整个数据

### 引用类型的传递

```go
// 切片传递 - 引用类型
func modifySlice(s []int) {
    for i := range s {
        s[i] = s[i] * 2  // 会修改原始切片
    }
}

// 映射传递 - 引用类型  
func modifyMap(m map[string]int) {
    m["new"] = 100       // 会修改原始映射
}

func main() {
    slice := []int{1, 2, 3, 4}
    fmt.Printf("修改前: %v\n", slice)
    modifySlice(slice)
    fmt.Printf("修改后: %v\n", slice)    // [2 4 6 8]
    
    m := make(map[string]int)
    m["old"] = 50
    modifyMap(m)
    fmt.Printf("映射内容: %v\n", m)      // map[new:100 old:50]
}
```

## 高级特性

### 变参函数

变参函数可以接受可变数量的参数：

```go
// 变参函数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 混合参数（固定参数 + 变参）
func formatMessage(prefix string, args ...interface{}) string {
    message := prefix + ": "
    for _, arg := range args {
        message += fmt.Sprintf("%v ", arg)
    }
    return message
}

func main() {
    // 变参调用
    fmt.Println(sum(1, 2, 3))           // 6
    fmt.Println(sum(1, 2, 3, 4, 5))     // 15
    
    // 传递切片给变参函数
    numbers := []int{10, 20, 30}
    fmt.Println(sum(numbers...))        // 60，注意...语法
    
    // 混合参数示例
    msg := formatMessage("Info", "user", 123, true)
    fmt.Println(msg)  // "Info: user 123 true "
}
```

**变参函数特点**：
- 变参必须是最后一个参数
- 在函数内部，变参是切片类型
- 调用时可以传递切片，使用 `...` 展开

### 函数作为值

Go语言中，函数是"一等公民"，可以作为变量、参数和返回值：

```go
// 定义函数类型
type MathFunc func(int, int) int

// 函数作为变量
func main() {
    // 将函数赋值给变量
    var operation MathFunc = add
    result := operation(5, 3)  // 8
    
    // 函数作为参数
    calculate(10, 5, add)      // 传递add函数
    calculate(10, 5, multiply) // 传递multiply函数
}

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

// 高阶函数：接受函数作为参数
func calculate(a, b int, op MathFunc) {
    result := op(a, b)
    fmt.Printf("计算结果: %d\n", result)
}
```

### 匿名函数和闭包

匿名函数可以访问外部作用域的变量，形成闭包：

```go
func main() {
    // 匿名函数
    func() {
        fmt.Println("这是一个匿名函数")
    }()
    
    // 闭包：捕获外部变量
    counter := 0
    increment := func() int {
        counter++  // 访问外部变量
        return counter
    }
    
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
    fmt.Println(increment())  // 3
    
    // 函数工厂：返回闭包
    makeMultiplier := func(factor int) func(int) int {
        return func(n int) int {
            return n * factor  // 捕获factor变量
        }
    }
    
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

**闭包的特点**：
- 可以访问和修改外部作用域的变量
- 外部变量的生命周期会延长
- 常用于回调函数、事件处理等场景

### 递归函数

递归是函数调用自身的编程技巧：

```go
// 计算阶乘
func factorial(n int) int {
    if n <= 1 {
        return 1  // 基础情况
    }
    return n * factorial(n-1)  // 递归调用
}

// 斐波那契数列
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// 实用递归：计算目录大小
func dirSize(path string) int64 {
    var size int64
    // 这里简化了文件系统操作的代码
    // 实际应用中需要使用filepath.Walk等
    return size
}
```

**递归使用要点**：
- 必须有明确的终止条件（基础情况）
- 每次递归都要向终止条件靠近
- 注意栈溢出问题，深度过大时考虑改用循环

## 实践示例：简单计算器

让我们实现一个函数式的计算器来巩固所学知识：

```go
package main

import (
    "fmt"
    "math"
)

// 定义运算函数类型
type Operation func(float64, float64) (float64, error)

// 基本运算函数
func add(a, b float64) (float64, error) {
    return a + b, nil
}

func subtract(a, b float64) (float64, error) {
    return a - b, nil
}

func multiply(a, b float64) (float64, error) {
    return a * b, nil
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("除数不能为零")
    }
    return a / b, nil
}

func power(a, b float64) (float64, error) {
    return math.Pow(a, b), nil
}

// 计算器主函数
func calculator(a, b float64, op Operation) (result float64, err error) {
    return op(a, b)
}

// 批量计算
func batchCalculate(numbers []float64, op Operation) (results []float64, err error) {
    if len(numbers) < 2 {
        return nil, fmt.Errorf("至少需要两个数字")
    }
    
    result := numbers[0]
    for i := 1; i < len(numbers); i++ {
        result, err = op(result, numbers[i])
        if err != nil {
            return nil, err
        }
        results = append(results, result)
    }
    return results, nil
}

func main() {
    // 运算映射表
    operations := map[string]Operation{
        "+": add,
        "-": subtract,
        "*": multiply,
        "/": divide,
        "^": power,
    }
    
    fmt.Println("🧮 函数式计算器")
    
    // 单次计算
    result, err := calculator(10, 3, divide)
    if err != nil {
        fmt.Printf("错误: %v\n", err)
    } else {
        fmt.Printf("10 ÷ 3 = %.2f\n", result)
    }
    
    // 使用映射表计算
    if op, exists := operations["*"]; exists {
        result, _ := calculator(6, 7, op)
        fmt.Printf("6 × 7 = %.0f\n", result)
    }
    
    // 批量计算
    numbers := []float64{100, 2, 5}
    results, err := batchCalculate(numbers, divide)
    if err != nil {
        fmt.Printf("批量计算错误: %v\n", err)
    } else {
        fmt.Printf("连续除法 %v = %v\n", numbers, results)
    }
}
```

## 本章小结

通过本章学习，你应该掌握：

### 核心概念
- **函数定义**：语法简洁，支持多返回值
- **参数传递**：值传递vs指针传递，引用类型的特殊性
- **命名返回值**：增强可读性，自动初始化
- **变参函数**：灵活处理可变数量参数

### 高级特性
- **函数作为值**：可以赋值给变量，作为参数传递
- **匿名函数和闭包**：强大的编程工具
- **递归**：解决特定问题的优雅方式

### 最佳实践
1. **函数设计**：单一职责，参数不宜过多
2. **返回值**：错误处理用多返回值，简单函数可用命名返回值
3. **参数传递**：大型数据用指针或引用类型避免复制
4. **函数命名**：动词开头，表达函数的行为

### Go语言函数特色
- 多返回值让错误处理更自然
- 函数是一等公民，支持函数式编程
- 简洁的语法，没有默认参数和重载
- defer语句提供优雅的资源清理机制

::: tip 练习建议
尝试实现一个文本处理工具，包含多个处理函数（统计字符、转换大小写、查找替换等），练习函数的组织和调用。
:::
