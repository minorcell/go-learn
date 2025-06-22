---
title: 变量和数据类型
description: 学习Go语言的变量声明、数据类型和基本概念
---

# 变量和类型

学习任何编程语言，都要从变量和数据类型开始。Go语言的类型系统简洁而强大，让我们一起探索吧！

## 📖 本章内容

- 变量声明的多种方式
- Go语言的基本数据类型
- 常量的定义和使用
- 类型转换和零值概念

## 🔤 变量声明

### 基本声明语法

Go语言提供了多种声明变量的方式：

```go
package main

import "fmt"

func main() {
    // 方式1：声明后赋值
    var name string
    name = "Go语言"
    
    // 方式2：声明时初始化
    var version string = "1.22"
    
    // 方式3：类型推断
    var year = 2024
    
    // 方式4：短变量声明（函数内部）
    language := "Golang"
    
    fmt.Printf("语言: %s, 版本: %s, 年份: %d, 别名: %s\n", 
        name, version, year, language)
}
```

### 多变量声明

```go
package main

import "fmt"

func main() {
    // 同类型多变量
    var a, b, c int = 1, 2, 3
    
    // 不同类型多变量
    var (
        name    string = "Alice"
        age     int    = 25
        height  float64 = 1.68
        married bool   = false
    )
    
    // 短变量声明
    x, y, z := 10, 20, "hello"
    
    fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
    fmt.Printf("姓名:%s, 年龄:%d, 身高:%.2f, 已婚:%t\n", 
        name, age, height, married)
    fmt.Printf("x=%d, y=%d, z=%s\n", x, y, z)
}
```

## 📊 基本数据类型

### 数值类型

```go
package main

import "fmt"

func main() {
    // 整数类型
    var a int8 = 127        // -128 到 127
    var b int16 = 32767     // -32768 到 32767
    var c int32 = 2147483647
    var d int64 = 9223372036854775807
    
    // 无符号整数
    var e uint8 = 255       // 0 到 255
    var f uint16 = 65535    // 0 到 65535
    
    // 浮点数
    var g float32 = 3.14
    var h float64 = 3.141592653589793
    
    // 平台相关的int（推荐使用）
    var count int = 100
    
    fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n", a, b, c, d)
    fmt.Printf("uint8: %d, uint16: %d\n", e, f)
    fmt.Printf("float32: %.2f, float64: %.10f\n", g, h)
    fmt.Printf("int: %d\n", count)
}
```

### 字符串类型

```go
package main

import "fmt"

func main() {
    // 普通字符串
    name := "Go语言"
    greeting := "Hello, World!"
    
    // 原始字符串（反引号）
    multiline := `这是一个
多行字符串
可以包含换行符`
    
    // 字符串拼接
    message := "Hello, " + name + "!"
    
    fmt.Printf("姓名: %s\n", name)
    fmt.Printf("问候: %s\n", greeting)
    fmt.Printf("多行字符串:\n%s\n", multiline)
    fmt.Printf("拼接结果: %s\n", message)
    
    // 字符串长度
    fmt.Printf("name长度: %d字节\n", len(name))
    fmt.Printf("name字符数: %d个\n", len([]rune(name)))
}
```

### 布尔类型

```go
package main

import "fmt"

func main() {
    var isActive bool = true
    var isCompleted bool = false
    
    // 布尔运算
    result1 := isActive && isCompleted  // false
    result2 := isActive || isCompleted  // true
    result3 := !isActive               // false
    
    fmt.Printf("isActive: %t\n", isActive)
    fmt.Printf("isCompleted: %t\n", isCompleted)
    fmt.Printf("AND运算: %t\n", result1)
    fmt.Printf("OR运算: %t\n", result2)
    fmt.Printf("NOT运算: %t\n", result3)
    
    // 实际应用
    age := 20
    hasLicense := true
    canDrive := age >= 18 && hasLicense
    fmt.Printf("年龄%d岁，有驾照:%t，可以开车:%t\n", age, hasLicense, canDrive)
}
```

## 🔒 常量

### 基本常量

```go
package main

import "fmt"

func main() {
    // 单个常量
    const pi = 3.14159
    const greeting = "Hello, Go!"
    
    // 常量组
    const (
        StatusOK = 200
        StatusNotFound = 404
        StatusError = 500
    )
    
    fmt.Printf("圆周率: %.5f\n", pi)
    fmt.Printf("问候语: %s\n", greeting)
    fmt.Printf("状态码: 成功=%d, 未找到=%d, 错误=%d\n", 
        StatusOK, StatusNotFound, StatusError)
    
    // 计算圆面积
    radius := 5.0
    area := pi * radius * radius
    fmt.Printf("半径%.1f的圆面积: %.2f\n", radius, area)
}
```

### iota 枚举器

```go
package main

import "fmt"

func main() {
    // iota 自动递增
    const (
        Sunday = iota    // 0
        Monday          // 1
        Tuesday         // 2
        Wednesday       // 3
        Thursday        // 4
        Friday          // 5
        Saturday        // 6
    )
    
    // 自定义起始值
    const (
        _ = iota           // 0，忽略
        January            // 1
        February           // 2
        March              // 3
    )
    
    // 表达式中使用iota
    const (
        B = 1 << (10 * iota)  // 1
        KB                    // 1024
        MB                    // 1048576
        GB                    // 1073741824
    )
    
    fmt.Printf("今天是星期%d\n", Wednesday)
    fmt.Printf("现在是%d月\n", March)
    fmt.Printf("存储单位: 1B=%d, 1KB=%d, 1MB=%d, 1GB=%d\n", B, KB, MB, GB)
}
```

## 🔄 类型转换

### 基本类型转换

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 数值类型转换
    var intNum int = 42
    var floatNum float64 = float64(intNum)
    var int32Num int32 = int32(intNum)
    
    fmt.Printf("int: %d\n", intNum)
    fmt.Printf("float64: %.1f\n", floatNum)
    fmt.Printf("int32: %d\n", int32Num)
    
    // 字符串转数字
    str := "123"
    num, err := strconv.Atoi(str)
    if err == nil {
        fmt.Printf("字符串'%s'转数字: %d\n", str, num)
    }
    
    // 数字转字符串
    numStr := strconv.Itoa(intNum)
    fmt.Printf("数字%d转字符串: '%s'\n", intNum, numStr)
    
    // 浮点数转字符串
    floatStr := strconv.FormatFloat(floatNum, 'f', 2, 64)
    fmt.Printf("浮点数%.1f转字符串: '%s'\n", floatNum, floatStr)
}
```

## 🔍 零值

Go语言中，所有类型都有默认的零值：

```go
package main

import "fmt"

func main() {
    // 声明但不初始化的变量会有零值
    var intVal int
    var floatVal float64
    var boolVal bool
    var stringVal string
    
    fmt.Printf("int零值: %d\n", intVal)
    fmt.Printf("float64零值: %.1f\n", floatVal)
    fmt.Printf("bool零值: %t\n", boolVal)
    fmt.Printf("string零值: '%s' (长度: %d)\n", stringVal, len(stringVal))
    
    // 零值的实用性
    var counter int  // 自动初始化为0
    counter++
    fmt.Printf("计数器: %d\n", counter)
    
    var isReady bool // 自动初始化为false
    if !isReady {
        fmt.Println("系统未准备就绪")
    }
}
```

## 🎯 实践练习

让我们通过一个完整的例子来练习所学内容：

```go
package main

import "fmt"

func main() {
    // 个人信息管理
    const title = "=== 个人信息卡 ==="
    
    // 基本信息
    name := "张三"
    age := 28
    height := 175.5  // cm
    weight := 70.2   // kg
    isMarried := false
    city := "北京"
    
    // 计算BMI
    heightInMeters := height / 100.0
    bmi := weight / (heightInMeters * heightInMeters)
    
    // BMI分类
    var bmiCategory string
    switch {
    case bmi < 18.5:
        bmiCategory = "偏瘦"
    case bmi < 24:
        bmiCategory = "正常"
    case bmi < 28:
        bmiCategory = "偏胖"
    default:
        bmiCategory = "肥胖"
    }
    
    // 输出信息
    fmt.Println(title)
    fmt.Printf("姓名: %s\n", name)
    fmt.Printf("年龄: %d岁\n", age)
    fmt.Printf("身高: %.1fcm\n", height)
    fmt.Printf("体重: %.1fkg\n", weight)
    fmt.Printf("BMI: %.1f (%s)\n", bmi, bmiCategory)
    fmt.Printf("婚姻状况: %t\n", isMarried)
    fmt.Printf("居住城市: %s\n", city)
    
    // 年龄相关计算
    const currentYear = 2024
    birthYear := currentYear - age
    fmt.Printf("出生年份: %d年\n", birthYear)
    
    // 生成个性化消息
    greeting := fmt.Sprintf("你好，%s！欢迎来到%s！", name, city)
    fmt.Println(greeting)
}
```

## 📝 本章小结

在这一章中，我们学习了：

### 🔹 变量声明
- `var` 关键字声明
- 类型推断 
- 短变量声明 `:=`
- 多变量声明

### 🔹 数据类型
- **整数类型**：int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint
- **浮点类型**：float32, float64
- **布尔类型**：bool
- **字符串类型**：string

### 🔹 常量
- `const` 关键字
- 常量组
- `iota` 枚举器

### 🔹 重要概念
- 零值机制
- 类型转换
- 字符串与数字互转

## 🎯 下一步

掌握了变量和类型后，让我们继续学习 [控制流程](./control-flow)，了解如何控制程序的执行流程！

<ChapterNav />

<ProgressTracker />
