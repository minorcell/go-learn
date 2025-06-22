# 函数

函数是组织代码的基本单位，让程序更加模块化和可重用。Go语言的函数设计简洁而强大。

## 本章内容

- 函数的定义和调用
- 参数传递和返回值
- 匿名函数和闭包
- 变参函数和递归
- 函数作为值传递

## 基本函数

### 函数定义和调用

```go
package main

import "fmt"

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

func main() {
    // 调用函数
    sayHello()
    greet("Alice")
    
    sum := add(3, 5)
    fmt.Printf("3 + 5 = %d\n", sum)
    
    result, err := divide(10, 3)
    if err != nil {
        fmt.Printf("错误: %v\n", err)
    } else {
        fmt.Printf("10 / 3 = %.2f\n", result)
    }
}
```

### 命名返回值

```go
package main

import "fmt"

// 命名返回值
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // 自动返回命名的变量
}

// 更复杂的例子
func analyzeScore(score int) (grade string, passed bool, message string) {
    switch {
    case score >= 90:
        grade = "A"
        passed = true
        message = "优秀"
    case score >= 80:
        grade = "B"
        passed = true
        message = "良好"
    case score >= 60:
        grade = "C"
        passed = true
        message = "及格"
    default:
        grade = "F"
        passed = false
        message = "不及格"
    }
    return
}

func main() {
    sum, product := calculate(4, 5)
    fmt.Printf("和: %d, 积: %d\n", sum, product)
    
    grade, passed, message := analyzeScore(85)
    fmt.Printf("分数分析: %s等级, 通过: %t, 评价: %s\n", 
        grade, passed, message)
}
```

## 参数传递

### 值传递和指针传递

```go
package main

import "fmt"

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

// 切片传递 - 引用类型
func modifySlice(s []int) {
    for i := range s {
        s[i] = s[i] * 2
    }
    fmt.Printf("函数内部切片: %v\n", s)
}

func main() {
    // 值传递示例
    a := 10
    fmt.Printf("修改前 a = %d\n", a)
    modifyValue(a)
    fmt.Printf("修改后 a = %d\n", a) // 不变
    
    // 指针传递示例
    b := 10
    fmt.Printf("\n修改前 b = %d\n", b)
    modifyPointer(&b)
    fmt.Printf("修改后 b = %d\n", b) // 改变了
    
    // 切片传递示例
    slice := []int{1, 2, 3, 4}
    fmt.Printf("\n修改前切片: %v\n", slice)
    modifySlice(slice)
    fmt.Printf("修改后切片: %v\n", slice) // 改变了
}
```

### 变参函数

```go
package main

import "fmt"

// 变参函数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 字符串格式化函数
func formatMessage(template string, args ...interface{}) string {
    return fmt.Sprintf(template, args...)
}

// 混合参数
func processData(operation string, numbers ...int) (string, int) {
    var result int
    switch operation {
    case "sum":
        for _, num := range numbers {
            result += num
        }
    case "product":
        result = 1
        for _, num := range numbers {
            result *= num
        }
    case "max":
        if len(numbers) > 0 {
            result = numbers[0]
            for _, num := range numbers {
                if num > result {
                    result = num
                }
            }
        }
    }
    return operation, result
}

func main() {
    // 调用变参函数
    fmt.Printf("sum() = %d\n", sum())
    fmt.Printf("sum(1, 2, 3) = %d\n", sum(1, 2, 3))
    fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
    
    // 传递切片
    numbers := []int{10, 20, 30}
    fmt.Printf("sum(numbers...) = %d\n", sum(numbers...))
    
    // 格式化消息
    msg := formatMessage("Hello %s, you are %d years old", "Alice", 25)
    fmt.Println(msg)
    
    // 混合参数
    op, result := processData("sum", 1, 2, 3, 4, 5)
    fmt.Printf("%s result: %d\n", op, result)
    
    op, result = processData("max", 15, 7, 23, 9, 12)
    fmt.Printf("%s result: %d\n", op, result)
}
```

## 匿名函数和闭包

### 匿名函数

```go
package main

import "fmt"

func main() {
    // 定义并立即调用匿名函数
    func() {
        fmt.Println("这是一个匿名函数")
    }()
    
    // 带参数的匿名函数
    func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }("Bob")
    
    // 将匿名函数赋值给变量
    greet := func(name string) string {
        return fmt.Sprintf("Hi, %s!", name)
    }
    
    message := greet("Charlie")
    fmt.Println(message)
    
    // 匿名函数作为返回值
    getMultiplier := func(factor int) func(int) int {
        return func(x int) int {
            return x * factor
        }
    }
    
    double := getMultiplier(2)
    triple := getMultiplier(3)
    
    fmt.Printf("double(5) = %d\n", double(5))
    fmt.Printf("triple(5) = %d\n", triple(5))
}
```

### 闭包

```go
package main

import "fmt"

// 计数器闭包
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 累加器闭包
func createAccumulator(initial int) func(int) int {
    total := initial
    return func(value int) int {
        total += value
        return total
    }
}

// 配置闭包
func createFormatter(prefix, suffix string) func(string) string {
    return func(content string) string {
        return prefix + content + suffix
    }
}

func main() {
    // 计数器示例
    counter1 := createCounter()
    counter2 := createCounter()
    
    fmt.Printf("counter1: %d\n", counter1()) // 1
    fmt.Printf("counter1: %d\n", counter1()) // 2
    fmt.Printf("counter2: %d\n", counter2()) // 1
    fmt.Printf("counter1: %d\n", counter1()) // 3
    
    // 累加器示例
    acc := createAccumulator(10)
    fmt.Printf("acc(5) = %d\n", acc(5))   // 15
    fmt.Printf("acc(3) = %d\n", acc(3))   // 18
    fmt.Printf("acc(-2) = %d\n", acc(-2)) // 16
    
    // 格式化器示例
    htmlFormatter := createFormatter("<p>", "</p>")
    markdownFormatter := createFormatter("**", "**")
    
    fmt.Println(htmlFormatter("Hello, World!"))
    fmt.Println(markdownFormatter("Bold Text"))
}
```

## 递归函数

```go
package main

import "fmt"

// 阶乘
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// 二分查找
func binarySearch(arr []int, target, left, right int) int {
    if left > right {
        return -1 // 未找到
    }
    
    mid := (left + right) / 2
    if arr[mid] == target {
        return mid
    } else if arr[mid] > target {
        return binarySearch(arr, target, left, mid-1)
    } else {
        return binarySearch(arr, target, mid+1, right)
    }
}

// 计算目录大小（模拟）
func calculateSize(path string, depth int) int {
    // 模拟文件大小计算
    indent := ""
    for i := 0; i < depth; i++ {
        indent += "  "
    }
    
    if depth > 3 { // 模拟递归深度限制
        fmt.Printf("%s%s (文件: 1KB)\n", indent, path)
        return 1
    }
    
    // 模拟目录包含子项
    fmt.Printf("%s%s (目录)\n", indent, path)
    size := 0
    
    // 模拟子目录和文件
    for i := 1; i <= 2; i++ {
        subPath := fmt.Sprintf("%s/item%d", path, i)
        size += calculateSize(subPath, depth+1)
    }
    
    return size
}

func main() {
    // 阶乘示例
    fmt.Println("阶乘计算:")
    for i := 1; i <= 6; i++ {
        fmt.Printf("%d! = %d\n", i, factorial(i))
    }
    
    // 斐波那契数列
    fmt.Println("\n斐波那契数列:")
    for i := 0; i < 10; i++ {
        fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
    }
    
    // 二分查找
    fmt.Println("\n二分查找:")
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    targets := []int{5, 8, 13, 16}
    
    for _, target := range targets {
        index := binarySearch(arr, target, 0, len(arr)-1)
        if index != -1 {
            fmt.Printf("找到 %d 在索引 %d\n", target, index)
        } else {
            fmt.Printf("未找到 %d\n", target)
        }
    }
    
    // 目录大小计算
    fmt.Println("\n目录结构:")
    totalSize := calculateSize("/home/user", 0)
    fmt.Printf("总大小: %dKB\n", totalSize)
}
```

## 函数作为值

```go
package main

import "fmt"

// 定义函数类型
type MathOperation func(int, int) int
type StringProcessor func(string) string

// 基本数学运算函数
func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int   { return a / b }

// 字符串处理函数
func toUpper(s string) string     { return fmt.Sprintf("UPPER: %s", s) }
func toLower(s string) string     { return fmt.Sprintf("lower: %s", s) }
func addQuotes(s string) string   { return fmt.Sprintf("\"%s\"", s) }

// 接受函数作为参数
func calculate(a, b int, op MathOperation) int {
    return op(a, b)
}

// 批量处理
func processNumbers(numbers []int, operations []MathOperation) {
    for i, op := range operations {
        result := op(numbers[0], numbers[1])
        fmt.Printf("运算 %d: %d\n", i+1, result)
    }
}

// 字符串处理管道
func processString(input string, processors ...StringProcessor) string {
    result := input
    for _, processor := range processors {
        result = processor(result)
    }
    return result
}

// 返回函数
func getOperation(opType string) MathOperation {
    switch opType {
    case "add":
        return add
    case "subtract":
        return subtract
    case "multiply":
        return multiply
    case "divide":
        return divide
    default:
        return nil
    }
}

func main() {
    // 函数作为参数
    fmt.Println("函数作为参数:")
    fmt.Printf("add(10, 5) = %d\n", calculate(10, 5, add))
    fmt.Printf("multiply(10, 5) = %d\n", calculate(10, 5, multiply))
    
    // 函数切片
    fmt.Println("\n批量运算:")
    operations := []MathOperation{add, subtract, multiply, divide}
    processNumbers([]int{20, 4}, operations)
    
    // 字符串处理管道
    fmt.Println("\n字符串处理管道:")
    result := processString("hello world", toUpper, addQuotes)
    fmt.Println("结果:", result)
    
    // 动态获取函数
    fmt.Println("\n动态函数调用:")
    opNames := []string{"add", "multiply", "divide"}
    for _, opName := range opNames {
        op := getOperation(opName)
        if op != nil {
            result := op(15, 3)
            fmt.Printf("%s(15, 3) = %d\n", opName, result)
        }
    }
    
    // 函数映射
    fmt.Println("\n函数映射:")
    opMap := map[string]MathOperation{
        "+": add,
        "-": subtract,
        "*": multiply,
        "/": divide,
    }
    
    expressions := []struct {
        a, b int
        op   string
    }{
        {8, 2, "+"},
        {8, 2, "-"},
        {8, 2, "*"},
        {8, 2, "/"},
    }
    
    for _, expr := range expressions {
        if operation, exists := opMap[expr.op]; exists {
            result := operation(expr.a, expr.b)
            fmt.Printf("%d %s %d = %d\n", expr.a, expr.op, expr.b, result)
        }
    }
}
```

## 实战练习

### 学生管理系统（函数版）

```go
package main

import "fmt"

// 学生结构
type Student struct {
    Name  string
    Score int
}

// 成绩统计
type Statistics struct {
    Average   float64
    Max       int
    Min       int
    PassCount int
    Total     int
}

// 函数类型定义
type StudentFilter func(Student) bool
type ScoreCalculator func([]Student) float64

// 创建学生
func createStudent(name string, score int) Student {
    return Student{Name: name, Score: score}
}

// 添加学生
func addStudent(students []Student, student Student) []Student {
    return append(students, student)
}

// 过滤学生
func filterStudents(students []Student, filter StudentFilter) []Student {
    var result []Student
    for _, student := range students {
        if filter(student) {
            result = append(result, student)
        }
    }
    return result
}

// 计算统计信息
func calculateStatistics(students []Student) Statistics {
    if len(students) == 0 {
        return Statistics{}
    }
    
    total := 0
    max := students[0].Score
    min := students[0].Score
    passCount := 0
    
    for _, student := range students {
        total += student.Score
        if student.Score > max {
            max = student.Score
        }
        if student.Score < min {
            min = student.Score
        }
        if student.Score >= 60 {
            passCount++
        }
    }
    
    return Statistics{
        Average:   float64(total) / float64(len(students)),
        Max:       max,
        Min:       min,
        PassCount: passCount,
        Total:     len(students),
    }
}

// 打印学生列表
func printStudents(students []Student, title string) {
    fmt.Printf("\n=== %s ===\n", title)
    if len(students) == 0 {
        fmt.Println("没有学生")
        return
    }
    
    for i, student := range students {
        fmt.Printf("%d. %s: %d分\n", i+1, student.Name, student.Score)
    }
}

// 打印统计信息
func printStatistics(stats Statistics) {
    fmt.Printf("\n=== 统计信息 ===\n")
    fmt.Printf("总人数: %d\n", stats.Total)
    fmt.Printf("平均分: %.1f\n", stats.Average)
    fmt.Printf("最高分: %d\n", stats.Max)
    fmt.Printf("最低分: %d\n", stats.Min)
    fmt.Printf("及格人数: %d\n", stats.PassCount)
    fmt.Printf("及格率: %.1f%%\n", 
        float64(stats.PassCount)/float64(stats.Total)*100)
}

func main() {
    // 创建学生数据
    var students []Student
    
    studentData := []struct {
        name  string
        score int
    }{
        {"张三", 92},
        {"李四", 78},
        {"王五", 85},
        {"赵六", 67},
        {"钱七", 94},
        {"孙八", 56},
        {"周九", 88},
    }
    
    // 添加学生
    for _, data := range studentData {
        student := createStudent(data.name, data.score)
        students = addStudent(students, student)
    }
    
    // 打印所有学生
    printStudents(students, "所有学生")
    
    // 定义过滤器函数
    excellentFilter := func(s Student) bool { return s.Score >= 90 }
    passingFilter := func(s Student) bool { return s.Score >= 60 }
    failingFilter := func(s Student) bool { return s.Score < 60 }
    
    // 过滤并显示不同类型的学生
    excellent := filterStudents(students, excellentFilter)
    printStudents(excellent, "优秀学生 (≥90分)")
    
    passing := filterStudents(students, passingFilter)
    printStudents(passing, "及格学生 (≥60分)")
    
    failing := filterStudents(students, failingFilter)
    printStudents(failing, "不及格学生 (<60分)")
    
    // 计算并显示统计信息
    stats := calculateStatistics(students)
    printStatistics(stats)
    
    // 特定范围查询
    goodFilter := func(s Student) bool { 
        return s.Score >= 80 && s.Score < 90 
    }
    good := filterStudents(students, goodFilter)
    printStudents(good, "良好学生 (80-89分)")
}
```

## 本章小结

在这一章中，我们学习了：

### 函数基础
- **函数定义** - func 关键字，参数和返回值
- **函数调用** - 参数传递，返回值接收
- **命名返回值** - 简化返回语句

### 参数传递
- **值传递** - 基本类型的拷贝传递
- **指针传递** - 通过指针修改原始值
- **变参函数** - 接受可变数量的参数

### 高级特性
- **匿名函数** - 函数字面量
- **闭包** - 函数捕获外部变量
- **递归** - 函数调用自身
- **函数作为值** - 函数的一等公民地位

### 最佳实践
- 函数名要见名知意
- 保持函数功能单一
- 适当使用命名返回值
- 合理运用闭包和高阶函数
