# 控制流程

掌握了变量和类型后，现在学习如何控制程序的执行流程。控制流程让程序能够根据条件做出决策和重复执行任务。

## 本章内容

- 条件语句：if/else 和 switch
- 循环语句：for 的各种形式
- 流程控制：break、continue、goto
- 标签和跳转控制

## 条件语句

### if/else 语句

```go
package main

import "fmt"

func main() {
    age := 20
    score := 85
    
    // 基本 if 语句
    if age >= 18 {
        fmt.Println("你已经成年了！")
    }
    
    // if/else 语句
    if score >= 90 {
        fmt.Println("成绩优秀！")
    } else if score >= 80 {
        fmt.Println("成绩良好！")
    } else if score >= 60 {
        fmt.Println("成绩及格！")
    } else {
        fmt.Println("需要努力！")
    }
    
    // 带初始化的 if 语句
    if result := score * 1.2; result >= 100 {
        fmt.Printf("加权后分数: %.1f，满分！\n", result)
    } else {
        fmt.Printf("加权后分数: %.1f\n", result)
    }
}
```

### 复杂条件判断

```go
package main

import "fmt"

func main() {
    temperature := 25
    humidity := 80
    hasUmbrella := true
    
    // 逻辑运算符
    if temperature > 30 && humidity > 70 {
        fmt.Println("又热又潮湿，开空调！")
    } else if temperature < 10 || humidity > 90 {
        fmt.Println("天气不适宜外出")
    } else {
        fmt.Println("天气还不错")
    }
    
    // 复杂的组合条件
    canGoOut := (temperature >= 15 && temperature <= 30) || 
                (temperature < 15 && hasUmbrella)
    
    if canGoOut {
        fmt.Println("可以外出")
    } else {
        fmt.Println("建议待在室内")
    }
}
```

### switch 语句

```go
package main

import "fmt"

func main() {
    weekday := 3
    
    // 基本 switch
    switch weekday {
    case 1:
        fmt.Println("星期一 - 新的开始")
    case 2:
        fmt.Println("星期二 - 继续努力")
    case 3:
        fmt.Println("星期三 - 过半了")
    case 4:
        fmt.Println("星期四 - 快结束")
    case 5:
        fmt.Println("星期五 - 感谢上帝")
    case 6, 7:  // 多个值
        fmt.Println("周末 - 休息时间")
    default:
        fmt.Println("无效日期")
    }
    
    // 字符串 switch
    grade := "A"
    switch grade {
    case "A":
        fmt.Println("优秀：90-100分")
    case "B":
        fmt.Println("良好：80-89分")
    case "C":
        fmt.Println("及格：60-79分")
    default:
        fmt.Println("不及格：60分以下")
    }
}
```

### 无表达式的 switch

```go
package main

import "fmt"

func main() {
    score := 85
    age := 20
    
    // 相当于多个 if-else
    switch {
    case score >= 90:
        fmt.Println("成绩优秀")
    case score >= 80:
        fmt.Println("成绩良好")
    case score >= 70:
        fmt.Println("成绩中等")
    case score >= 60:
        fmt.Println("成绩及格")
    default:
        fmt.Println("成绩不及格")
    }
    
    // 复杂条件的 switch
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
}
```

## 循环语句

Go语言只有一种循环关键字 `for`，但它非常灵活。

### 传统的 for 循环

```go
package main

import "fmt"

func main() {
    // 基本 for 循环
    fmt.Println("数数 1 到 5：")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // 计算阶乘
    factorial := 1
    n := 5
    for i := 1; i <= n; i++ {
        factorial *= i
    }
    fmt.Printf("%d! = %d\n", n, factorial)
    
    // 倒计时
    fmt.Println("\n倒计时：")
    for count := 10; count >= 1; count-- {
        fmt.Printf("%d ", count)
    }
    fmt.Println("发射！")
    
    // 步长为 2
    fmt.Println("\n偶数 2 到 10：")
    for i := 2; i <= 10; i += 2 {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}
```

### while 风格的循环

```go
package main

import "fmt"

func main() {
    // 模拟 while 循环
    sum := 0
    i := 1
    for i <= 100 {
        sum += i
        i++
    }
    fmt.Printf("1到100的和: %d\n", sum)
    
    // 猜数字游戏（模拟）
    target := 42
    guess := 1
    attempts := 0
    
    fmt.Println("\n猜数字游戏！目标是42")
    for guess != target {
        attempts++
        if guess < target {
            fmt.Printf("第%d次: %d - 太小了\n", attempts, guess)
            guess += 10
        } else {
            fmt.Printf("第%d次: %d - 太大了\n", attempts, guess)
            guess -= 5
        }
        
        // 防止无限循环
        if attempts > 10 {
            break
        }
    }
    
    if guess == target {
        fmt.Printf("猜中了！用了%d次\n", attempts+1)
    }
}
```

### 无限循环

```go
package main

import "fmt"

func main() {
    // 无限循环示例
    counter := 0
    for {
        counter++
        fmt.Printf("循环第%d次\n", counter)
        
        // 设置退出条件
        if counter >= 5 {
            fmt.Println("达到限制，退出")
            break
        }
        
        // 跳过某些情况
        if counter == 3 {
            fmt.Println("跳过第3次的处理")
            continue
        }
        
        fmt.Printf("  → 完成第%d次处理\n", counter)
    }
    
    // 模拟服务器处理
    fmt.Println("\n模拟处理请求：")
    requestCount := 0
    for {
        requestCount++
        fmt.Printf("处理请求 #%d\n", requestCount)
        
        if requestCount >= 3 {
            fmt.Println("服务器关闭")
            break
        }
    }
}
```

### range 循环

```go
package main

import "fmt"

func main() {
    // 遍历切片
    numbers := []int{10, 20, 30, 40, 50}
    
    fmt.Println("遍历切片（索引和值）：")
    for index, value := range numbers {
        fmt.Printf("索引%d: 值%d\n", index, value)
    }
    
    // 只要值
    fmt.Println("\n只获取值：")
    for _, value := range numbers {
        fmt.Printf("%d ", value)
    }
    fmt.Println()
    
    // 只要索引
    fmt.Println("\n只获取索引：")
    for index := range numbers {
        fmt.Printf("索引%d ", index)
    }
    fmt.Println()
    
    // 遍历字符串
    message := "Hello,世界"
    fmt.Printf("\n遍历字符串 '%s'：\n", message)
    for index, char := range message {
        fmt.Printf("位置%d: %c\n", index, char)
    }
    
    // 遍历 map
    colors := map[string]string{
        "red":   "红色",
        "green": "绿色",
        "blue":  "蓝色",
    }
    
    fmt.Println("\n遍历 map：")
    for english, chinese := range colors {
        fmt.Printf("%s -> %s\n", english, chinese)
    }
}
```

## 流程控制

### break 和 continue

```go
package main

import "fmt"

func main() {
    // break 示例：找第一个能被7整除的数
    fmt.Println("找第一个能被7整除的两位数：")
    for i := 10; i < 100; i++ {
        if i%7 == 0 {
            fmt.Printf("找到了：%d\n", i)
            break
        }
    }
    
    // continue 示例：打印奇数
    fmt.Println("\n1到20中的奇数：")
    for i := 1; i <= 20; i++ {
        if i%2 == 0 {
            continue // 跳过偶数
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // 嵌套循环中的 break 和 continue
    fmt.Println("\n九九乘法表（结果 ≤ 30）：")
    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            product := i * j
            if product > 30 {
                break // 跳出内层循环
            }
            if product%5 != 0 {
                continue // 只显示能被5整除的
            }
            fmt.Printf("%d×%d=%d ", i, j, product)
        }
        fmt.Println()
    }
}
```

### 标签和跳转

```go
package main

import "fmt"

func main() {
    // 使用标签跳出多层循环
    fmt.Println("寻找满足条件的数字对：")
    
outer:
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            product := i * j
            fmt.Printf("测试 %d × %d = %d\n", i, j, product)
            
            // 找到目标值就跳出所有循环
            if product == 12 {
                fmt.Printf("找到目标！%d × %d = %d\n", i, j, product)
                break outer
            }
            
            // 值太大就尝试下一组
            if product > 15 {
                fmt.Println("太大了，下一组")
                continue outer
            }
        }
    }
    
    fmt.Println("\n处理二维数据：")
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    target := 6
search:
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[row]); col++ {
            value := matrix[row][col]
            fmt.Printf("检查[%d][%d] = %d\n", row, col, value)
            
            if value == target {
                fmt.Printf("找到 %d 在位置 [%d][%d]\n", target, row, col)
                break search
            }
        }
    }
}
```

## 综合练习

### 学生成绩管理系统

```go
package main

import "fmt"

func main() {
    // 学生成绩数据
    students := map[string]int{
        "张三": 92,
        "李四": 78,
        "王五": 85,
        "赵六": 67,
        "钱七": 94,
    }
    
    fmt.Println("=== 学生成绩管理系统 ===")
    
    // 统计各分数段
    excellent, good, average, passing, failing := 0, 0, 0, 0, 0
    totalScore := 0
    studentCount := 0
    
    fmt.Println("\n1. 成绩详情：")
    for name, score := range students {
        studentCount++
        totalScore += score
        
        // 分数等级
        var grade string
        switch {
        case score >= 90:
            grade = "优秀"
            excellent++
        case score >= 80:
            grade = "良好" 
            good++
        case score >= 70:
            grade = "中等"
            average++
        case score >= 60:
            grade = "及格"
            passing++
        default:
            grade = "不及格"
            failing++
        }
        
        fmt.Printf("  %s: %d分 (%s)\n", name, score, grade)
    }
    
    // 计算平均分
    avgScore := float64(totalScore) / float64(studentCount)
    fmt.Printf("\n2. 平均分: %.1f\n", avgScore)
    
    // 统计报告
    fmt.Println("\n3. 分布统计：")
    fmt.Printf("  优秀(≥90): %d人\n", excellent)
    fmt.Printf("  良好(80-89): %d人\n", good)
    fmt.Printf("  中等(70-79): %d人\n", average)
    fmt.Printf("  及格(60-69): %d人\n", passing)
    fmt.Printf("  不及格(<60): %d人\n", failing)
    
    // 找最高分和最低分
    maxScore, minScore := 0, 100
    var topStudent, bottomStudent string
    
    for name, score := range students {
        if score > maxScore {
            maxScore = score
            topStudent = name
        }
        if score < minScore {
            minScore = score
            bottomStudent = name
        }
    }
    
    fmt.Printf("\n4. 最高分: %s (%d分)\n", topStudent, maxScore)
    fmt.Printf("   最低分: %s (%d分)\n", bottomStudent, minScore)
    
    // 需要关注的学生
    fmt.Println("\n5. 需要关注的学生：")
    hasLowPerformers := false
    for name, score := range students {
        if float64(score) < avgScore {
            fmt.Printf("  %s: %d分 (低于平均分)\n", name, score)
            hasLowPerformers = true
        }
    }
    
    if !hasLowPerformers {
        fmt.Println("  所有学生都在平均水平以上！")
    }
}
```

### 简单计算器

```go
package main

import "fmt"

func main() {
    fmt.Println("=== 简单计算器 ===")
    
    // 模拟计算任务
    operations := []struct {
        a, b   float64
        op     string
        desc   string
    }{
        {10, 5, "+", "加法"},
        {10, 3, "-", "减法"},
        {6, 7, "*", "乘法"},
        {15, 3, "/", "除法"},
        {10, 0, "/", "除零测试"},
        {2, 8, "^", "无效运算符"},
    }
    
    successCount := 0
    errorCount := 0
    
    for i, calc := range operations {
        fmt.Printf("\n运算 %d: %.1f %s %.1f (%s)\n", 
            i+1, calc.a, calc.op, calc.b, calc.desc)
        
        var result float64
        var valid bool = true
        
        switch calc.op {
        case "+":
            result = calc.a + calc.b
        case "-":
            result = calc.a - calc.b
        case "*":
            result = calc.a * calc.b
        case "/":
            if calc.b == 0 {
                fmt.Println("错误：除数不能为零")
                valid = false
            } else {
                result = calc.a / calc.b
            }
        default:
            fmt.Printf("错误：不支持运算符 '%s'\n", calc.op)
            valid = false
        }
        
        if valid {
            fmt.Printf("结果: %.2f\n", result)
            successCount++
            
            // 结果分类
            switch {
            case result > 50:
                fmt.Println("   (大数值)")
            case result > 10:
                fmt.Println("   (中等数值)")
            case result > 0:
                fmt.Println("   (小正数)")
            case result == 0:
                fmt.Println("   (零)")
            default:
                fmt.Println("   (负数)")
            }
        } else {
            errorCount++
        }
    }
    
    // 运算总结
    fmt.Println("\n=== 运算总结 ===")
    fmt.Printf("成功: %d次\n", successCount)
    fmt.Printf("失败: %d次\n", errorCount)
    fmt.Printf("成功率: %.1f%%\n", 
        float64(successCount)/float64(len(operations))*100)
}
```

##  本章小结

在这一章中，我们学习了：

### 条件语句
- **if/else** - 基本条件判断，支持初始化
- **switch** - 多分支选择，支持无表达式形式
- **逻辑运算符** - &&、||、! 组合复杂条件

### 循环语句  
- **for循环** - Go唯一的循环关键字，功能强大
- **while风格** - 省略初始化和后置语句
- **无限循环** - 使用 for {} 实现
- **range循环** - 遍历数组、切片、字符串、map

### 流程控制
- **break** - 跳出循环
- **continue** - 跳过当前迭代  
- **标签** - 控制多层循环跳转

### 最佳实践
- 优先使用 switch 而不是多层 if-else
- 合理使用 range 遍历集合
- 在复杂嵌套中使用标签提高可读性

## 下一步

掌握了控制流程后，让我们学习 [函数](./functions)，了解如何组织和重用代码！ 