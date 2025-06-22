---
title: 字符串和正则表达式
description: 学习Go语言的字符串处理和正则表达式应用
---

# 字符串和正则表达式

字符串处理是编程中的常见任务。Go语言提供了强大的字符串操作库和正则表达式支持，让我们掌握这些实用技能。

## 本章内容

- 字符串基础操作
- 字符串格式化和模板
- 正则表达式匹配和替换
- 文本解析和处理
- 性能优化技巧

##  字符串基础操作

### 字符串处理函数

```go
package main

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"
    "unicode/utf8"
)

func main() {
    // 字符串基础操作
    basicStringOperations()
    
    // 字符串查找和替换
    stringSearchAndReplace()
    
    // 字符串分割和连接
    stringSplitAndJoin()
    
    // 字符串转换
    stringConversion()
    
    // Unicode处理
    unicodeHandling()
}

// 字符串基础操作
func basicStringOperations() {
    fmt.Println("=== 字符串基础操作 ===")
    
    text := "Hello, Go语言编程!"
    
    // 长度相关
    fmt.Printf("字符串: %s\n", text)
    fmt.Printf("字节长度: %d\n", len(text))
    fmt.Printf("字符长度: %d\n", utf8.RuneCountInString(text))
    
    // 大小写转换
    fmt.Printf("转大写: %s\n", strings.ToUpper(text))
    fmt.Printf("转小写: %s\n", strings.ToLower(text))
    fmt.Printf("标题格式: %s\n", strings.Title(text))
    
    // 前缀和后缀检查
    fmt.Printf("以'Hello'开头: %t\n", strings.HasPrefix(text, "Hello"))
    fmt.Printf("以'编程!'结尾: %t\n", strings.HasSuffix(text, "编程!"))
    
    // 空白处理
    spaceText := "  \t  Go语言  \n  "
    fmt.Printf("原始字符串: '%s'\n", spaceText)
    fmt.Printf("去除空白: '%s'\n", strings.TrimSpace(spaceText))
    fmt.Printf("去除前后字符: '%s'\n", strings.Trim("!!!Go语言!!!", "!"))
    
    fmt.Println()
}

// 字符串查找和替换
func stringSearchAndReplace() {
    fmt.Println("=== 字符串查找和替换 ===")
    
    text := "Go是一门现代化的编程语言，Go语言简洁高效"
    
    // 查找操作
    fmt.Printf("原文: %s\n", text)
    fmt.Printf("包含'Go': %t\n", strings.Contains(text, "Go"))
    fmt.Printf("'Go'首次出现位置: %d\n", strings.Index(text, "Go"))
    fmt.Printf("'Go'最后出现位置: %d\n", strings.LastIndex(text, "Go"))
    fmt.Printf("'Go'出现次数: %d\n", strings.Count(text, "Go"))
    
    // 替换操作
    fmt.Printf("替换第一个'Go': %s\n", strings.Replace(text, "Go", "Golang", 1))
    fmt.Printf("替换所有'Go': %s\n", strings.ReplaceAll(text, "Go", "Golang"))
    
    // 使用Replacer进行多重替换
    replacer := strings.NewReplacer(
        "Go", "Golang",
        "语言", "Language",
        "编程", "Programming",
    )
    fmt.Printf("多重替换: %s\n", replacer.Replace(text))
    
    fmt.Println()
}

// 字符串分割和连接
func stringSplitAndJoin() {
    fmt.Println("=== 字符串分割和连接 ===")
    
    // 分割操作
    csv := "张三,25,北京,工程师"
    fmt.Printf("CSV数据: %s\n", csv)
    
    parts := strings.Split(csv, ",")
    fmt.Printf("分割结果: %v\n", parts)
    
    // 高级分割
    text := "apple;banana;;cherry;grape;"
    fmt.Printf("原始数据: %s\n", text)
    fmt.Printf("普通分割: %v\n", strings.Split(text, ";"))
    fmt.Printf("字段分割(忽略空值): %v\n", strings.FieldsFunc(text, func(c rune) bool {
        return c == ';'
    }))
    
    // 按空白字符分割
    whitespaceText := "Go   语言    编程\t教程\n学习"
    fmt.Printf("按空白分割: %v\n", strings.Fields(whitespaceText))
    
    // 连接操作
    words := []string{"Go", "语言", "编程", "教程"}
    fmt.Printf("数组: %v\n", words)
    fmt.Printf("用空格连接: %s\n", strings.Join(words, " "))
    fmt.Printf("用'-'连接: %s\n", strings.Join(words, "-"))
    
    // 使用Builder高效连接大量字符串
    var builder strings.Builder
    builder.WriteString("构建的字符串: ")
    for i, word := range words {
        if i > 0 {
            builder.WriteString(" -> ")
        }
        builder.WriteString(word)
    }
    fmt.Printf("Builder结果: %s\n", builder.String())
    
    fmt.Println()
}

// 字符串转换
func stringConversion() {
    fmt.Println("=== 字符串转换 ===")
    
    // 数字转字符串
    num := 12345
    float := 123.456
    bool := true
    
    fmt.Printf("整数转字符串: %s\n", strconv.Itoa(num))
    fmt.Printf("浮点转字符串: %s\n", strconv.FormatFloat(float, 'f', 2, 64))
    fmt.Printf("布尔转字符串: %s\n", strconv.FormatBool(bool))
    
    // 字符串转数字
    strNum := "98765"
    strFloat := "987.654"
    strBool := "true"
    
    if parsed, err := strconv.Atoi(strNum); err == nil {
        fmt.Printf("字符串转整数: %d\n", parsed)
    }
    
    if parsed, err := strconv.ParseFloat(strFloat, 64); err == nil {
        fmt.Printf("字符串转浮点: %.3f\n", parsed)
    }
    
    if parsed, err := strconv.ParseBool(strBool); err == nil {
        fmt.Printf("字符串转布尔: %t\n", parsed)
    }
    
    // 进制转换
    fmt.Printf("十进制255转二进制: %s\n", strconv.FormatInt(255, 2))
    fmt.Printf("十进制255转十六进制: %s\n", strconv.FormatInt(255, 16))
    
    if parsed, err := strconv.ParseInt("ff", 16, 64); err == nil {
        fmt.Printf("十六进制'ff'转十进制: %d\n", parsed)
    }
    
    fmt.Println()
}

// Unicode处理
func unicodeHandling() {
    fmt.Println("=== Unicode处理 ===")
    
    text := "Hello,世界!🌍"
    fmt.Printf("混合文本: %s\n", text)
    
    // 遍历字符
    fmt.Println("字符分析:")
    for i, r := range text {
        fmt.Printf("  位置%d: 字符'%c'(U+%04X) - ", i, r, r)
        
        switch {
        case unicode.IsLetter(r):
            fmt.Println("字母")
        case unicode.IsDigit(r):
            fmt.Println("数字")
        case unicode.IsPunct(r):
            fmt.Println("标点")
        case unicode.IsSpace(r):
            fmt.Println("空格")
        case unicode.IsSymbol(r):
            fmt.Println("符号")
        default:
            fmt.Println("其他")
        }
    }
    
    // 字符分类统计
    var letters, digits, symbols, others int
    for _, r := range text {
        switch {
        case unicode.IsLetter(r):
            letters++
        case unicode.IsDigit(r):
            digits++
        case unicode.IsSymbol(r) || unicode.IsPunct(r):
            symbols++
        default:
            others++
        }
    }
    
    fmt.Printf("统计结果 - 字母:%d, 数字:%d, 符号:%d, 其他:%d\n", 
        letters, digits, symbols, others)
    
    fmt.Println()
}
```

## 字符串格式化

### 高级格式化技巧

```go
package main

import (
    "fmt"
    "strings"
    "text/template"
    "time"
)

func main() {
    // 基础格式化
    basicFormatting()
    
    // 模板使用
    templateUsage()
    
    // 自定义格式化
    customFormatting()
}

// 基础格式化
func basicFormatting() {
    fmt.Println("=== 基础格式化 ===")
    
    name := "张三"
    age := 25
    score := 85.6
    passed := true
    
    // 基本格式化
    fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
    fmt.Printf("分数: %.2f, 是否通过: %t\n", score, passed)
    
    // 数字格式化
    num := 1234567
    fmt.Printf("数字: %d\n", num)
    fmt.Printf("带逗号分隔符: %s\n", addCommas(num))
    fmt.Printf("百分比: %.1f%%\n", score)
    fmt.Printf("科学计数法: %e\n", float64(num))
    
    // 字符串格式化
    fmt.Printf("左对齐(宽度20): '%-20s'\n", name)
    fmt.Printf("右对齐(宽度20): '%20s'\n", name)
    fmt.Printf("居中对齐(宽度20): '%s'\n", centerString(name, 20))
    
    // 数字填充
    fmt.Printf("零填充: %08d\n", 123)
    fmt.Printf("空格填充: %8d\n", 123)
    
    // 时间格式化
    now := time.Now()
    fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
    fmt.Printf("日期: %s\n", now.Format("2006年01月02日"))
    fmt.Printf("时间: %s\n", now.Format("15:04:05"))
    
    fmt.Println()
}

// 添加千位分隔符
func addCommas(num int) string {
    str := fmt.Sprintf("%d", num)
    n := len(str)
    if n <= 3 {
        return str
    }
    
    var result strings.Builder
    for i, r := range str {
        if i > 0 && (n-i)%3 == 0 {
            result.WriteString(",")
        }
        result.WriteRune(r)
    }
    return result.String()
}

// 字符串居中
func centerString(s string, width int) string {
    if len(s) >= width {
        return s
    }
    
    padding := width - len(s)
    leftPad := padding / 2
    rightPad := padding - leftPad
    
    return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

// 模板使用
func templateUsage() {
    fmt.Println("=== 模板使用 ===")
    
    // 简单模板
    simpleTemplate := `
用户信息卡片
================
姓名: {{.Name}}
年龄: {{.Age}}
邮箱: {{.Email}}
注册时间: {{.RegisterTime.Format "2006-01-02"}}
{{if .IsVIP}}🌟 VIP用户{{else}}普通用户{{end}}
================
`
    
    tmpl, err := template.New("user").Parse(simpleTemplate)
    if err != nil {
        fmt.Printf("模板解析错误: %v\n", err)
        return
    }
    
    user := struct {
        Name         string
        Age          int
        Email        string
        RegisterTime time.Time
        IsVIP        bool
    }{
        Name:         "李四",
        Age:          28,
        Email:        "lisi@example.com",
        RegisterTime: time.Now().AddDate(-1, 0, 0),
        IsVIP:        true,
    }
    
    fmt.Println("简单模板输出:")
    tmpl.Execute(fmt.Println, user)
    
    // 复杂模板
    complexTemplate := `
产品报告
========
{{range .Products}}
产品: {{.Name}} (ID: {{.ID}})
价格: ¥{{printf "%.2f" .Price}}
库存: {{.Stock}}{{if lt .Stock 10}} ⚠️ 库存不足{{end}}
评分: {{stars .Rating}}
--------
{{end}}
总计: {{len .Products}} 种产品
`
    
    // 自定义模板函数
    funcMap := template.FuncMap{
        "stars": func(rating float64) string {
            stars := int(rating)
            result := strings.Repeat("⭐", stars)
            if rating-float64(stars) >= 0.5 {
                result += "✨"
            }
            return result
        },
    }
    
    complexTmpl, err := template.New("products").Funcs(funcMap).Parse(complexTemplate)
    if err != nil {
        fmt.Printf("复杂模板解析错误: %v\n", err)
        return
    }
    
    data := struct {
        Products []struct {
            ID     int
            Name   string
            Price  float64
            Stock  int
            Rating float64
        }
    }{
        Products: []struct {
            ID     int
            Name   string
            Price  float64
            Stock  int
            Rating float64
        }{
            {1, "Go编程书籍", 89.99, 15, 4.8},
            {2, "机械键盘", 299.99, 5, 4.5},
            {3, "显示器", 1299.99, 8, 4.2},
        },
    }
    
    fmt.Println("复杂模板输出:")
    complexTmpl.Execute(fmt.Println, data)
    
    fmt.Println()
}

// 自定义格式化
func customFormatting() {
    fmt.Println("=== 自定义格式化 ===")
    
    // 表格格式化
    createTable()
    
    // 进度条
    createProgressBar()
    
    // 树形结构
    createTree()
}

// 创建表格
func createTable() {
    fmt.Println("学生成绩表")
    fmt.Println(strings.Repeat("=", 50))
    
    students := []struct {
        Name    string
        Math    int
        English int
        Science int
    }{
        {"张三", 85, 92, 78},
        {"李四", 92, 88, 95},
        {"王五", 78, 85, 82},
    }
    
    // 表头
    fmt.Printf("%-10s %6s %8s %8s %6s\n", "姓名", "数学", "英语", "科学", "平均")
    fmt.Println(strings.Repeat("-", 50))
    
    // 数据行
    for _, student := range students {
        avg := (student.Math + student.English + student.Science) / 3
        fmt.Printf("%-10s %6d %8d %8d %6d\n", 
            student.Name, student.Math, student.English, student.Science, avg)
    }
    
    fmt.Println()
}

// 创建进度条
func createProgressBar() {
    fmt.Println("任务进度:")
    
    tasks := []struct {
        Name     string
        Progress int
    }{
        {"数据下载", 100},
        {"数据处理", 75},
        {"报告生成", 45},
        {"质量检查", 20},
    }
    
    for _, task := range tasks {
        bar := createBar(task.Progress, 20)
        fmt.Printf("%-12s [%s] %3d%%\n", task.Name, bar, task.Progress)
    }
    
    fmt.Println()
}

// 创建进度条
func createBar(progress, width int) string {
    filled := progress * width / 100
    bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
    return bar
}

// 创建树形结构
func createTree() {
    fmt.Println("项目目录结构:")
    
    tree := []struct {
        Level int
        Name  string
        IsDir bool
    }{
        {0, "myproject", true},
        {1, "src", true},
        {2, "main.go", false},
        {2, "utils", true},
        {3, "helper.go", false},
        {1, "docs", true},
        {2, "README.md", false},
        {1, "tests", true},
        {2, "main_test.go", false},
    }
    
    for _, node := range tree {
        indent := strings.Repeat("  ", node.Level)
        var icon string
        if node.IsDir {
            icon = "📁"
        } else {
            icon = "📄"
        }
        
        prefix := "├── "
        if node.Level == 0 {
            prefix = ""
        }
        
        fmt.Printf("%s%s%s %s\n", indent, prefix, icon, node.Name)
    }
    
    fmt.Println()
}
```

## 正则表达式

### 正则匹配和处理

```go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    // 基础正则匹配
    basicRegexp()
    
    // 高级正则操作
    advancedRegexp()
    
    // 实用正则示例
    practicalRegexp()
    
    // 正则性能优化
    regexpPerformance()
}

// 基础正则匹配
func basicRegexp() {
    fmt.Println("=== 基础正则匹配 ===")
    
    text := "联系我们: 电话 010-12345678 或邮箱 support@example.com"
    
    // 简单匹配
    matched, _ := regexp.MatchString(`\d{3}-\d{8}`, text)
    fmt.Printf("包含电话号码: %t\n", matched)
    
    // 编译正则表达式
    phoneRegex := regexp.MustCompile(`\d{3}-\d{8}`)
    emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    
    // 查找匹配
    phone := phoneRegex.FindString(text)
    email := emailRegex.FindString(text)
    
    fmt.Printf("找到的电话: %s\n", phone)
    fmt.Printf("找到的邮箱: %s\n", email)
    
    // 查找所有匹配
    multiText := "电话: 010-12345678, 021-87654321, 邮箱: admin@test.com, user@demo.org"
    phones := phoneRegex.FindAllString(multiText, -1)
    emails := emailRegex.FindAllString(multiText, -1)
    
    fmt.Printf("所有电话: %v\n", phones)
    fmt.Printf("所有邮箱: %v\n", emails)
    
    fmt.Println()
}

// 高级正则操作
func advancedRegexp() {
    fmt.Println("=== 高级正则操作 ===")
    
    // 分组捕获
    text := "用户: 张三(ID:1001), 李四(ID:1002), 王五(ID:1003)"
    userRegex := regexp.MustCompile(`(\w+)\(ID:(\d+)\)`)
    
    matches := userRegex.FindAllStringSubmatch(text, -1)
    fmt.Println("用户信息提取:")
    for _, match := range matches {
        fmt.Printf("  用户: %s, ID: %s\n", match[1], match[2])
    }
    
    // 命名分组
    namedRegex := regexp.MustCompile(`(?P<name>\w+)\(ID:(?P<id>\d+)\)`)
    fmt.Println("命名分组提取:")
    for _, match := range namedRegex.FindAllStringSubmatch(text, -1) {
        result := make(map[string]string)
        for i, name := range namedRegex.SubexpNames() {
            if i != 0 && name != "" {
                result[name] = match[i]
            }
        }
        fmt.Printf("  用户: %s, ID: %s\n", result["name"], result["id"])
    }
    
    // 替换操作
    fmt.Println("替换操作:")
    original := "价格: $99.99, $149.99, $79.99"
    priceRegex := regexp.MustCompile(`\$(\d+\.\d+)`)
    
    // 简单替换
    replaced := priceRegex.ReplaceAllString(original, "¥$1")
    fmt.Printf("美元转人民币: %s\n", replaced)
    
    // 函数替换
    converted := priceRegex.ReplaceAllStringFunc(original, func(s string) string {
        price := priceRegex.FindStringSubmatch(s)[1]
        return fmt.Sprintf("¥%.2f", parseFloat(price)*6.5) // 假设汇率6.5
    })
    fmt.Printf("汇率转换: %s\n", converted)
    
    fmt.Println()
}

// 辅助函数：解析浮点数
func parseFloat(s string) float64 {
    var result float64
    fmt.Sscanf(s, "%f", &result)
    return result
}

// 实用正则示例
func practicalRegexp() {
    fmt.Println("=== 实用正则示例 ===")
    
    // 数据验证
    validateData()
    
    // 文本提取
    extractData()
    
    // 日志解析
    parseLog()
}

// 数据验证
func validateData() {
    fmt.Println("--- 数据验证 ---")
    
    // 定义验证规则
    validators := map[string]*regexp.Regexp{
        "phone":    regexp.MustCompile(`^1[3-9]\d{9}$`),
        "email":    regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
        "idcard":   regexp.MustCompile(`^\d{17}[\dXx]$`),
        "password": regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`),
    }
    
    testData := map[string][]string{
        "phone": {"13812345678", "12812345678", "138123456789"},
        "email": {"test@example.com", "invalid.email", "user@domain.co.uk"},
        "idcard": {"11010519491231002X", "123456789", "11010519491231002Y"},
        "password": {"Password123!", "password", "PASSWORD123", "Pass123!"},
    }
    
    for dataType, tests := range testData {
        fmt.Printf("%s验证结果:\n", dataType)
        for _, test := range tests {
            valid := validators[dataType].MatchString(test)
            status := "❌"
            if valid {
                status = "✅"
            }
            fmt.Printf("  %s %s\n", status, test)
        }
        fmt.Println()
    }
}

// 提取数据
func extractData() {
    fmt.Println("--- 数据提取 ---")
    
    html := `
    <div class="user-info">
        <span class="name">张三</span>
        <span class="age">25</span>
        <span class="email">zhangsan@example.com</span>
        <img src="/avatar/123.jpg" alt="头像">
    </div>
    <div class="user-info">
        <span class="name">李四</span>
        <span class="age">30</span>
        <span class="email">lisi@example.com</span>
        <img src="/avatar/456.jpg" alt="头像">
    </div>
    `
    
    // 提取用户信息
    nameRegex := regexp.MustCompile(`<span class="name">([^<]+)</span>`)
    ageRegex := regexp.MustCompile(`<span class="age">(\d+)</span>`)
    emailRegex := regexp.MustCompile(`<span class="email">([^<]+)</span>`)
    
    names := nameRegex.FindAllStringSubmatch(html, -1)
    ages := ageRegex.FindAllStringSubmatch(html, -1)
    emails := emailRegex.FindAllStringSubmatch(html, -1)
    
    fmt.Println("从HTML提取的用户信息:")
    for i := 0; i < len(names) && i < len(ages) && i < len(emails); i++ {
        fmt.Printf("  姓名: %s, 年龄: %s, 邮箱: %s\n", 
            names[i][1], ages[i][1], emails[i][1])
    }
    
    fmt.Println()
}

// 日志解析
func parseLog() {
    fmt.Println("--- 日志解析 ---")
    
    logData := `
[2024-01-15 10:30:15] INFO  User login successful - UserID: 1001, IP: 192.168.1.100
[2024-01-15 10:31:22] ERROR Database connection failed - Error: timeout after 30s
[2024-01-15 10:32:05] WARN  Rate limit exceeded - UserID: 1002, IP: 192.168.1.105
[2024-01-15 10:33:18] INFO  User logout - UserID: 1001, IP: 192.168.1.100
[2024-01-15 10:34:30] ERROR Invalid authentication token - UserID: 1003
    `
    
    // 日志解析正则
    logRegex := regexp.MustCompile(`\[([^\]]+)\]\s+(\w+)\s+(.+?)(?:\s+-\s+(.+))?$`)
    
    fmt.Println("日志解析结果:")
    lines := strings.Split(strings.TrimSpace(logData), "\n")
    
    for _, line := range lines {
        if matches := logRegex.FindStringSubmatch(line); matches != nil {
            timestamp := matches[1]
            level := matches[2]
            message := matches[3]
            details := matches[4]
            
            fmt.Printf("时间: %s | 级别: %s | 消息: %s", timestamp, level, message)
            if details != "" {
                fmt.Printf(" | 详情: %s", details)
            }
            fmt.Println()
        }
    }
    
    // 统计日志级别
    levelRegex := regexp.MustCompile(`\]\s+(\w+)\s+`)
    levels := levelRegex.FindAllStringSubmatch(logData, -1)
    
    levelCount := make(map[string]int)
    for _, match := range levels {
        levelCount[match[1]]++
    }
    
    fmt.Println("日志级别统计:")
    for level, count := range levelCount {
        fmt.Printf("  %s: %d\n", level, count)
    }
    
    fmt.Println()
}

// 正则性能优化
func regexpPerformance() {
    fmt.Println("=== 正则性能优化 ===")
    
    text := strings.Repeat("test@example.com, admin@test.org, ", 1000)
    
    // 错误做法：每次都编译
    fmt.Println("性能对比:")
    
    // 预编译（推荐）
    compiledRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    matches1 := compiledRegex.FindAllString(text, -1)
    fmt.Printf("预编译结果: 找到 %d 个邮箱\n", len(matches1))
    
    // 字面量匹配（性能更好）
    if strings.Contains(text, "@") {
        fmt.Println("包含邮箱特征字符")
    }
    
    // 优化技巧
    fmt.Println("优化技巧:")
    fmt.Println("1. 预编译正则表达式")
    fmt.Println("2. 使用字面量匹配替代简单正则")
    fmt.Println("3. 避免贪婪匹配，使用非贪婪匹配")
    fmt.Println("4. 使用字符类而非多选项")
    fmt.Println("5. 将常用匹配放在前面")
    
    fmt.Println()
}
```

## 文本处理实战

### 综合文本处理示例

```go
package main

import (
    "bufio"
    "fmt"
    "regexp"
    "sort"
    "strconv"
    "strings"
)

func main() {
    // 文本分析
    textAnalysis()
    
    // 数据清洗
    dataCleaning()
    
    // 文本转换
    textTransformation()
}

// 文本分析
func textAnalysis() {
    fmt.Println("=== 文本分析 ===")
    
    article := `
Go语言是Google开发的一种静态强类型、编译型语言。Go语言语法与C相近，但功能上有：内存安全，GC（垃圾回收），结构化类型，并发编程。
Go的语法接近C语言，但对于变量的声明有所不同。Go支持垃圾回收功能。Go的并行模型是以东尼·霍尔的通信顺序过程（CSP）为基础，采取类似模型的其他语言包括Occam和Limbo。
与C++相比，Go并不包括如枚举、异常处理、继承、泛型、断言、虚函数等功能，但增加了 切片(Slice) 型、并发、管道、垃圾回收、接口（Interface）等特性的语言级支持。
    `
    
    // 基础统计
    words := extractWords(article)
    sentences := extractSentences(article)
    
    fmt.Printf("文章分析结果:\n")
    fmt.Printf("字符数: %d\n", len(article))
    fmt.Printf("词数: %d\n", len(words))
    fmt.Printf("句子数: %d\n", len(sentences))
    fmt.Printf("平均句长: %.1f 个词\n", float64(len(words))/float64(len(sentences)))
    
    // 词频统计
    wordFreq := countWordFrequency(words)
    fmt.Println("高频词汇 (前10):")
    
    type wordCount struct {
        word  string
        count int
    }
    
    var sortedWords []wordCount
    for word, count := range wordFreq {
        if len(word) > 1 { // 过滤单字符
            sortedWords = append(sortedWords, wordCount{word, count})
        }
    }
    
    sort.Slice(sortedWords, func(i, j int) bool {
        return sortedWords[i].count > sortedWords[j].count
    })
    
    for i, wc := range sortedWords {
        if i >= 10 {
            break
        }
        fmt.Printf("  %s: %d次\n", wc.word, wc.count)
    }
    
    fmt.Println()
}

// 提取单词
func extractWords(text string) []string {
    // 使用正则表达式提取中英文词汇
    wordRegex := regexp.MustCompile(`[\p{Han}]+|[a-zA-Z]+`)
    return wordRegex.FindAllString(text, -1)
}

// 提取句子
func extractSentences(text string) []string {
    sentences := regexp.MustCompile(`[。！？.!?]+`).Split(text, -1)
    var result []string
    for _, sentence := range sentences {
        sentence = strings.TrimSpace(sentence)
        if len(sentence) > 0 {
            result = append(result, sentence)
        }
    }
    return result
}

// 统计词频
func countWordFrequency(words []string) map[string]int {
    freq := make(map[string]int)
    for _, word := range words {
        word = strings.ToLower(word)
        freq[word]++
    }
    return freq
}

// 数据清洗
func dataCleaning() {
    fmt.Println("=== 数据清洗 ===")
    
    rawData := `
    张三,25,   北京,  工程师
    李四,30,上海,设计师    
    王五,28,广州,产品经理
    ,22,深圳,开发者
    赵六,,杭州,测试工程师
    钱七,35,成都,
    `
    
    fmt.Println("原始数据:")
    fmt.Println(rawData)
    
    // 清洗数据
    cleanedData := cleanCSVData(rawData)
    
    fmt.Println("清洗后数据:")
    for _, record := range cleanedData {
        fmt.Printf("姓名: %-6s 年龄: %-3s 城市: %-6s 职位: %s\n", 
            record[0], record[1], record[2], record[3])
    }
    
    fmt.Println()
}

// 清洗CSV数据
func cleanCSVData(data string) [][]string {
    var cleaned [][]string
    
    scanner := bufio.NewScanner(strings.NewReader(data))
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        
        // 分割字段
        fields := strings.Split(line, ",")
        var cleanedFields []string
        
        for _, field := range fields {
            // 清理空白字符
            cleanedField := strings.TrimSpace(field)
            cleanedFields = append(cleanedFields, cleanedField)
        }
        
        // 验证数据完整性
        if len(cleanedFields) == 4 && 
           cleanedFields[0] != "" && // 姓名不能为空
           cleanedFields[2] != "" && // 城市不能为空
           cleanedFields[3] != "" {  // 职位不能为空
            
            // 验证年龄
            if cleanedFields[1] != "" {
                if age, err := strconv.Atoi(cleanedFields[1]); err != nil || age < 18 || age > 65 {
                    cleanedFields[1] = "未知"
                }
            } else {
                cleanedFields[1] = "未知"
            }
            
            cleaned = append(cleaned, cleanedFields)
        }
    }
    
    return cleaned
}

// 文本转换
func textTransformation() {
    fmt.Println("=== 文本转换 ===")
    
    // markdown转HTML
    markdownToHTML()
    
    // 文本格式化
    textFormatting()
    
    // 代码格式化
    codeFormatting()
}

// Markdown转HTML
func markdownToHTML() {
    fmt.Println("--- Markdown转HTML ---")
    
    markdown := `
# Go语言教程

## 简介
Go是一种**开源**的编程语言，它能让构造简单、*可靠*且高效的软件变得容易。

## 特性
- 静态类型
- 编译型语言
- 并发支持
- 垃圾回收

### 代码示例
` + "```go\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n```" + `

更多信息请访问 [官方网站](https://golang.org)
    `
    
    html := convertMarkdownToHTML(markdown)
    fmt.Println("转换结果:")
    fmt.Println(html)
}

// 简单的Markdown转HTML转换器
func convertMarkdownToHTML(markdown string) string {
    lines := strings.Split(markdown, "\n")
    var html strings.Builder
    
    inCodeBlock := false
    
    for _, line := range lines {
        line = strings.TrimSpace(line)
        
        if strings.HasPrefix(line, "```") {
            if inCodeBlock {
                html.WriteString("</pre></code>\n")
                inCodeBlock = false
            } else {
                lang := strings.TrimPrefix(line, "```")
                html.WriteString(fmt.Sprintf("<code><pre class=\"%s\">\n", lang))
                inCodeBlock = true
            }
            continue
        }
        
        if inCodeBlock {
            html.WriteString(line + "\n")
            continue
        }
        
        if line == "" {
            continue
        }
        
        // 标题
        if strings.HasPrefix(line, "### ") {
            html.WriteString(fmt.Sprintf("<h3>%s</h3>\n", strings.TrimPrefix(line, "### ")))
        } else if strings.HasPrefix(line, "## ") {
            html.WriteString(fmt.Sprintf("<h2>%s</h2>\n", strings.TrimPrefix(line, "## ")))
        } else if strings.HasPrefix(line, "# ") {
            html.WriteString(fmt.Sprintf("<h1>%s</h1>\n", strings.TrimPrefix(line, "# ")))
        } else if strings.HasPrefix(line, "- ") {
            html.WriteString(fmt.Sprintf("<li>%s</li>\n", strings.TrimPrefix(line, "- ")))
        } else {
            // 处理内联格式
            line = regexp.MustCompile(`\*\*([^*]+)\*\*`).ReplaceAllString(line, "<strong>$1</strong>")
            line = regexp.MustCompile(`\*([^*]+)\*`).ReplaceAllString(line, "<em>$1</em>")
            line = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(line, "<a href=\"$2\">$1</a>")
            
            html.WriteString(fmt.Sprintf("<p>%s</p>\n", line))
        }
    }
    
    return html.String()
}

// 文本格式化
func textFormatting() {
    fmt.Println("--- 文本格式化 ---")
    
    text := "这是一个很长的句子，需要按照一定的宽度进行换行处理，以便在控制台或者其他固定宽度的显示环境中正确显示。"
    
    wrapped := wrapText(text, 20)
    fmt.Println("文本换行(宽度20):")
    fmt.Println(wrapped)
    
    fmt.Println()
}

// 文本换行
func wrapText(text string, width int) string {
    if len(text) <= width {
        return text
    }
    
    var lines []string
    runes := []rune(text)
    
    for i := 0; i < len(runes); i += width {
        end := i + width
        if end > len(runes) {
            end = len(runes)
        }
        lines = append(lines, string(runes[i:end]))
    }
    
    return strings.Join(lines, "\n")
}

// 代码格式化
func codeFormatting() {
    fmt.Println("--- 代码格式化 ---")
    
    code := `package main
import "fmt"
func main(){fmt.Println("Hello");if true{fmt.Println("World")}}`
    
    formatted := formatGoCode(code)
    fmt.Println("格式化前:")
    fmt.Println(code)
    fmt.Println("\n格式化后:")
    fmt.Println(formatted)
    
    fmt.Println()
}

// 简单的Go代码格式化
func formatGoCode(code string) string {
    // 简单的格式化规则
    code = regexp.MustCompile(`{`).ReplaceAllString(code, " {\n    ")
    code = regexp.MustCompile(`}`).ReplaceAllString(code, "\n}")
    code = regexp.MustCompile(`;`).ReplaceAllString(code, "\n    ")
    code = regexp.MustCompile(`\n\s*\n`).ReplaceAllString(code, "\n")
    
    return strings.TrimSpace(code)
}
```

##  本章小结

在这一章中，我们学习了：

### 字符串操作
- 基础字符串处理函数
- 字符串格式化和模板
- Unicode和字符编码处理
- 性能优化技巧

### 正则表达式
- 基础正则匹配和查找
- 分组捕获和命名分组
- 正则替换和转换
- 性能优化策略

### 文本处理
- 文本分析和统计
- 数据清洗和验证
- 格式转换和处理
- 实际应用案例

### 实用技巧
- 模板系统使用
- 数据验证模式
- 文本格式化工具
- 代码处理技术