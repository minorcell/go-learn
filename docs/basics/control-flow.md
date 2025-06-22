---
title: 控制流程
description: 学习Go语言的条件语句、循环语句和流程控制
---

# 控制流程

掌握了变量和类型后，现在学习如何控制程序的执行流程。控制流程是编程的核心，它让程序能够根据条件做出决策、重复执行任务，从而实现复杂的逻辑。

## 本章内容

- 条件语句：if/else 和 switch 的使用
- 循环语句：for 循环的各种形式
- 流程控制：break、continue 和标签
- 实际应用场景和最佳实践

## 条件语句

条件语句让程序能够根据不同的情况执行不同的代码分支。

### if/else 语句

Go语言的 `if` 语句简洁而强大，支持初始化语句：

```go
age := 20
score := 85

// 基本条件判断
if age >= 18 {
    fmt.Println("你已经成年了！")
}

// 多重条件
if score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好") 
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("需要努力")
}

// if语句中的初始化（变量作用域仅在if块内）
if result := score * 1.2; result >= 100 {
    fmt.Printf("加权后满分：%.1f\n", result)
} else {
    fmt.Printf("加权后分数：%.1f\n", result)
}
```

**if语句特点**：
- 条件表达式不需要括号
- 花括号是必需的，即使只有一行代码
- 支持在条件前初始化变量
- 初始化的变量仅在if-else块内有效

### 逻辑运算符

在条件判断中，经常需要组合多个条件：

```go
temperature := 25
humidity := 80
hasUmbrella := true

// 与运算 (&&) - 所有条件都为真
if temperature > 30 && humidity > 70 {
    fmt.Println("又热又潮湿，开空调！")
}

// 或运算 (||) - 任一条件为真
if temperature < 10 || humidity > 90 {
    fmt.Println("天气恶劣，不宜外出")
}

// 复杂组合条件
canGoOut := (temperature >= 15 && temperature <= 30) || 
            (temperature < 15 && hasUmbrella)
```

### switch 语句

当需要根据一个变量的不同值执行不同逻辑时，`switch` 比多个 `if-else` 更清晰：

```go
// 基本switch用法
weekday := 3
switch weekday {
case 1:
    fmt.Println("星期一 - 新的开始")
case 2, 3, 4:  // 多个值匹配
    fmt.Println("工作日 - 继续努力") 
case 5:
    fmt.Println("星期五 - TGIF!")
case 6, 7:
    fmt.Println("周末 - 休息时间")
default:
    fmt.Println("无效日期")
}

// 字符串switch
grade := "A"
switch grade {
case "A":
    fmt.Println("优秀：90-100分")
case "B":
    fmt.Println("良好：80-89分") 
case "C":
    fmt.Println("及格：60-79分")
default:
    fmt.Println("不及格")
}
```

**switch语句特点**：
- 每个case自动break，不会穿透（与C/Java不同）
- 可以匹配多个值：`case 1, 2, 3:`
- 支持任意类型，不仅仅是整数
- `default` 分支是可选的

### 无表达式switch

这是Go语言的特色功能，相当于一系列 `if-else if`：

```go
score := 85
age := 20

switch {
case score >= 90:
    fmt.Println("成绩优秀")
case score >= 80:
    fmt.Println("成绩良好")
case score >= 60:
    fmt.Println("成绩及格")
default:
    fmt.Println("需要努力")
}

// 复杂条件判断
switch {
case age < 13:
    fmt.Println("儿童")
case age >= 13 && age < 20:
    fmt.Println("青少年")
case age >= 20 && age < 60:
    fmt.Println("成年人")
default:
    fmt.Println("老年人")
}
```

::: tip 使用建议
- 简单的值匹配用 `switch`
- 复杂的条件组合用无表达式 `switch`
- 只有2-3个条件时用 `if-else`
:::

## 循环语句

Go语言只有一种循环关键字 `for`，但它足够灵活，可以实现各种循环需求。

### 基本for循环

经典的三段式循环：`初始化; 条件; 后置语句`

```go
// 数数1到5
for i := 1; i <= 5; i++ {
    fmt.Printf("%d ", i)
}

// 计算阶乘
factorial := 1
for i := 1; i <= 5; i++ {
    factorial *= i
}
fmt.Printf("5! = %d\n", factorial)

// 倒计时
for count := 10; count >= 1; count-- {
    fmt.Printf("%d ", count)
}
fmt.Println("发射！")
```

### 条件循环（while风格）

省略初始化和后置语句，只保留条件：

```go
// 类似while循环
sum := 0
i := 1
for i <= 100 {
    sum += i
    i++
}
fmt.Printf("1到100的和：%d\n", sum)

// 输入验证循环
var input string
for input != "quit" {
    fmt.Print("请输入命令（输入quit退出）：")
    fmt.Scanln(&input)
    fmt.Printf("你输入了：%s\n", input)
}
```

### 无限循环

省略所有条件，使用 `break` 控制退出：

```go
counter := 0
for {
    counter++
    if counter > 10 {
        break  // 退出循环
    }
    if counter%2 == 0 {
        continue  // 跳过本次循环
    }
    fmt.Printf("奇数：%d\n", counter)
}
```

### range循环

Go语言的特色功能，用于遍历数组、切片、字符串等：

```go
// 遍历数组/切片
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("索引%d：值%d\n", index, value)
}

// 只要值，忽略索引
for _, value := range numbers {
    fmt.Printf("值：%d\n", value)
}

// 只要索引
for index := range numbers {
    fmt.Printf("索引：%d\n", index)
}

// 遍历字符串
for i, char := range "Hello" {
    fmt.Printf("位置%d：字符%c\n", i, char)
}
```

**range循环特点**：
- 返回两个值：索引和元素值
- 可以用 `_` 忽略不需要的值
- 对字符串按UTF-8字符遍历，不是字节

## 流程控制

### break和continue

控制循环的执行流程：

```go
// break：跳出循环
for i := 1; i <= 10; i++ {
    if i == 6 {
        break  // 遇到6就停止
    }
    fmt.Printf("%d ", i)  // 输出：1 2 3 4 5
}

// continue：跳过本次循环
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue  // 跳过偶数
    }
    fmt.Printf("%d ", i)  // 输出：1 3 5 7 9
}
```

### 标签和跳转

在嵌套循环中，可以使用标签精确控制跳转：

```go
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break outer  // 跳出外层循环
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
// 输出：(1,1) (1,2) (1,3) (2,1)
```

::: warning 注意
虽然Go支持 `goto` 语句，但在现代编程中应该避免使用，它会让代码难以理解和维护。
:::

## 实践示例：数字猜谜游戏

让我们用学到的知识实现一个完整的猜数字游戏：

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成1-100的随机数
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    attempts := 0
    maxAttempts := 7
    
    fmt.Println("🎮 欢迎来到猜数字游戏！")
    fmt.Printf("我想了一个1-100的数字，你有%d次机会猜中它！\n", maxAttempts)
    
    for attempts < maxAttempts {
        var guess int
        fmt.Printf("\n第%d次猜测，请输入你的数字：", attempts+1)
        fmt.Scanln(&guess)
        
        attempts++
        
        // 判断猜测结果
        switch {
        case guess == target:
            fmt.Printf("🎉 恭喜你！猜中了！数字就是%d\n", target)
            fmt.Printf("你用了%d次就猜中了，太厉害了！\n", attempts)
            return
        case guess < target:
            fmt.Println("📈 太小了，再试试更大的数字")
        case guess > target:
            fmt.Println("📉 太大了，再试试更小的数字")
        }
        
        // 给出剩余机会提示
        remaining := maxAttempts - attempts
        if remaining > 0 {
            fmt.Printf("还有%d次机会\n", remaining)
        }
    }
    
    fmt.Printf("\n😔 游戏结束！答案是%d，下次再来挑战吧！\n", target)
}
```

## 本章小结

通过本章学习，你应该掌握：

### 核心概念
- **条件语句**：`if-else` 用于简单判断，`switch` 用于多值匹配
- **循环语句**：`for` 的多种形式满足不同需求
- **流程控制**：`break`、`continue` 精确控制程序流程
- **range循环**：Go语言遍历集合类型的优雅方式

### 最佳实践
1. **条件选择**：2-3个分支用 `if-else`
2. **循环选择**：已知次数用计数循环，未知次数用条件循环
3. **可读性优先**：选择最直观表达逻辑的语句形式
4. **避免深层嵌套**：超过3层嵌套时考虑提取函数

### 特色功能
Go语言的控制流有这些特点：
- `if` 语句支持初始化
- `switch` 语句不会穿透，支持多值匹配
- `for` 语句是唯一的循环关键字，但形式灵活
- `range` 提供了遍历集合的统一方式

::: tip 练习建议
尝试实现一个简单的计算器，支持四则运算，练习使用不同的控制流语句。这将帮助你更好地理解本章内容。
:::