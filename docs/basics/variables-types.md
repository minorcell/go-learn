---
title: 变量和数据类型
description: 学习Go语言的变量声明、数据类型和基本概念
---

# 变量和类型

学习任何编程语言，都要从变量和数据类型开始。Go语言的类型系统简洁而强大，设计哲学是"简单胜过复杂"。

## 本章内容

- 变量声明的多种方式
- Go语言的基本数据类型
- 常量的定义和使用
- 类型转换和零值概念

## 变量声明

Go语言提供了4种声明变量的方式，每种都有其适用场景：

### 1. 标准声明：var 关键字

```go
var name string        // 声明后赋值
var version string = "1.22"  // 声明时初始化
var year = 2024       // 类型推断
```

**使用场景**：
- 包级别变量必须用 `var`
- 需要明确指定类型时
- 声明零值变量时

### 2. 短变量声明：:= 操作符

```go
language := "Golang"   // 仅在函数内部使用
x, y := 10, 20        // 同时声明多个变量
```

**使用场景**：
- 函数内部的局部变量
- 类型明确，无需显式声明
- 大部分情况的首选方式

### 3. 批量声明

当需要声明多个相关变量时，可以使用括号组织：

```go
var (
    name    string = "Alice"
    age     int    = 25
    height  float64 = 1.68
    married bool   = false
)
```

::: tip 实用建议
在实际开发中，优先使用 `:=` 进行局部变量声明，它简洁且类型安全。只有在需要零值或明确类型时才使用 `var`。
:::

## 基本数据类型

Go语言是强类型语言，所有变量都有明确的类型。基本类型分为四大类：

### 整数类型

Go提供了有符号和无符号整数类型，区别在于是否支持负数：

| 类型 | 范围 | 说明 |
|------|------|------|
| `int8` | -128 到 127 | 1字节有符号整数 |
| `int16` | -32768 到 32767 | 2字节有符号整数 |
| `int32` | -21亿 到 21亿 | 4字节有符号整数 |
| `int64` | 约-922万亿 到 922万亿 | 8字节有符号整数 |
| `uint8` | 0 到 255 | 1字节无符号整数 |
| `int` | 平台相关 | **推荐使用**，32或64位 |

```go
// 实际使用示例
count := 100        // 推荐：使用 int
var age uint8 = 25  // 明确范围较小时使用具体类型
```

### 浮点数类型

Go只有两种浮点数类型：

```go
var price float32 = 19.99    // 单精度，4字节
var pi float64 = 3.141592653 // 双精度，8字节（推荐）
```

**选择建议**：除非对内存占用有严格要求，否则总是使用 `float64`，它精度更高且是Go的默认浮点类型。

### 布尔类型

布尔类型只有两个值：`true` 和 `false`

```go
isActive := true
canLogin := age >= 18 && hasAccount
```

**重要特点**：
- 不能与数字类型互转（与C/C++不同）
- 零值为 `false`
- 常用于条件判断和标志位

### 字符串类型

Go的字符串是UTF-8编码的字节序列，不可变类型：

```go
name := "Go语言"
greeting := `多行字符串
可以包含换行
和特殊字符 "引号"`
```

**字符串特点**：
- 用双引号 `"` 或反引号 `` ` `` 定义
- 反引号内的字符串不转义，所见即所得
- 字符串不可变，修改会创建新字符串
- `len()` 返回字节数，不是字符数

## 常量

常量是编译时确定的值，运行时不可修改：

### 基本常量

```go
const pi = 3.14159
const greeting = "Hello, Go!"

// 常量组
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)
```

### iota 枚举器

`iota` 是Go的常量生成器，在 `const` 声明中自动递增：

```go
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
)

// 实用例子：定义文件大小单位
const (
    B = 1 << (10 * iota)  // 1
    KB                    // 1024
    MB                    // 1048576
    GB                    // 1073741824
)
```

**iota 特点**：
- 每个 `const` 声明块中，`iota` 从0开始
- 可以参与表达式计算
- 常用于定义枚举值和标志位

## 类型转换

Go是强类型语言，不同类型之间不会自动转换，必须显式转换：

### 数值类型转换

```go
var intNum int = 42
var floatNum float64 = float64(intNum)  // int → float64
var int32Num int32 = int32(intNum)      // int → int32
```

### 字符串与数字转换

Go标准库 `strconv` 包提供了字符串转换功能：

```go
import "strconv"

// 字符串 → 数字
str := "123"
num, err := strconv.Atoi(str)  // 返回值和错误
if err != nil {
    // 处理转换错误
}

// 数字 → 字符串
numStr := strconv.Itoa(42)           // "42"
floatStr := strconv.FormatFloat(3.14, 'f', 2, 64)  // "3.14"
```

::: warning 注意
类型转换可能导致精度丢失（如 float64 → int）或溢出。在转换前要确保数据范围合适。
:::

## 零值概念

Go语言的一个重要特性是**零值初始化**：声明但未初始化的变量会自动设置为类型的零值。

| 类型 | 零值 |
|------|------|
| 数值类型 | `0` |
| 布尔类型 | `false` |
| 字符串 | `""` (空字符串) |
| 指针、切片、映射、通道、函数 | `nil` |

```go
var counter int     // 0，可以直接使用
var isReady bool    // false
var message string  // ""

counter++  // 1，无需初始化就能使用
```

**零值的意义**：
- 避免了未初始化变量的问题
- 让代码更安全，减少运行时错误
- 简化了变量声明，很多情况下零值就是我们想要的初始值

## 实践示例：用户信息管理

让我们通过一个实际例子来巩固所学知识：

```go
package main

import "fmt"

func main() {
    // 用户基本信息
    const appName = "用户管理系统"
    
    name := "张三"
    age := 28
    height := 175.5  // cm
    weight := 70.2   // kg
    isVIP := false
    
    // 计算BMI指数
    heightInMeters := height / 100.0
    bmi := weight / (heightInMeters * heightInMeters)
    
    // 根据BMI判断健康状况
    var healthStatus string
    switch {
    case bmi < 18.5:
        healthStatus = "偏瘦"
    case bmi < 24:
        healthStatus = "正常"
    case bmi < 28:
        healthStatus = "偏胖"
    default:
        healthStatus = "肥胖"
    }
    
    // 输出用户信息
    fmt.Printf("=== %s ===\n", appName)
    fmt.Printf("用户：%s（%d岁）\n", name, age)
    fmt.Printf("BMI：%.1f（%s）\n", bmi, healthStatus)
    fmt.Printf("会员状态：%t\n", isVIP)
}
```

## 本章小结

通过本章学习，你应该掌握：

### 核心概念
- **变量声明**：`var` 和 `:=` 的使用场景
- **数据类型**：整数、浮点数、布尔、字符串的特点
- **常量**：`const` 和 `iota` 的应用
- **零值**：Go的安全初始化机制

### 最佳实践
1. 局部变量优先使用 `:=` 声明
2. 整数类型推荐使用 `int`，浮点数使用 `float64`  
3. 利用零值特性简化代码
4. 类型转换要注意数据范围和精度

::: tip 练习建议
尝试编写一个简单的计算器程序，练习不同数据类型的声明、运算和转换。这将帮助你更好地理解本章内容。
:::
