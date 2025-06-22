---
title: 字符串处理与正则表达式
description: 学习Go语言的字符串操作、文本处理和正则表达式匹配
---

# 字符串处理与正则表达式

字符串处理是编程中最常见的任务之一。Go语言提供了强大的字符串处理能力和正则表达式支持，让文本处理变得高效且优雅。

## 本章内容

- 字符串基础操作和格式化
- 高级字符串处理技巧
- 正则表达式模式匹配
- 文本解析和数据提取
- 字符串性能优化技巧

## 字符串处理概念

### Go语言字符串特性

Go语言中的字符串具有以下特点：

- **不可变性**：字符串创建后不能修改，确保安全性
- **UTF-8编码**：原生支持Unicode，处理多语言文本
- **高效操作**：strings包提供丰富的操作函数
- **字节序列**：底层是字节数组，可与[]byte互转

### 字符串处理优势

| 特性 | 说明 | 优势 |
|------|------|------|
| **内存安全** | 不可变字符串 | 避免意外修改 |
| **Unicode支持** | 原生UTF-8编码 | 国际化支持好 |
| **性能优化** | 高效的标准库 | 处理速度快 |
| **简洁语法** | 直观的操作方法 | 代码易读易写 |

::: tip 设计原则
Go字符串处理遵循"简洁、高效、安全"的设计理念：
- 提供丰富的标准库函数
- 保持API的一致性和直观性
- 重视性能和内存使用
:::

## 字符串基础操作

### 常用字符串操作

```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

func basicStringOperations() {
    text := "  Hello, Go Programming!  "
    
    // 基础操作
    fmt.Printf("原始字符串: '%s'\n", text)
    fmt.Printf("长度: %d\n", len(text))
    fmt.Printf("去除空格: '%s'\n", strings.TrimSpace(text))
    fmt.Printf("转大写: %s\n", strings.ToUpper(text))
    fmt.Printf("转小写: %s\n", strings.ToLower(text))
    
    // 查找操作
    searchText := "Go Programming"
    fmt.Printf("\n🔍 查找操作:")
    fmt.Printf("包含 'Go': %t\n", strings.Contains(searchText, "Go"))
    fmt.Printf("前缀 'Go': %t\n", strings.HasPrefix(searchText, "Go"))
    fmt.Printf("后缀 'ing': %t\n", strings.HasSuffix(searchText, "ing"))
    fmt.Printf("'Pro'位置: %d\n", strings.Index(searchText, "Pro"))
    
    // 分割和连接
    sentence := "Go,is,awesome,programming,language"
    words := strings.Split(sentence, ",")
    fmt.Printf("\n📝 分割结果: %v\n", words)
    
    joined := strings.Join(words, " ")
    fmt.Printf("连接结果: %s\n", joined)
    
    // 替换操作
    original := "Hello World Hello Universe"
    replaced := strings.Replace(original, "Hello", "Hi", 1)  // 替换第一个
    replacedAll := strings.ReplaceAll(original, "Hello", "Hi") // 替换所有
    
    fmt.Printf("\n🔄 替换操作:")
    fmt.Printf("原文: %s\n", original)
    fmt.Printf("替换第一个: %s\n", replaced)
    fmt.Printf("替换所有: %s\n", replacedAll)
}

// 字符串比较和验证
func stringComparison() {
    fmt.Println("\n📊 字符串比较:")
    
    str1 := "Hello"
    str2 := "hello"
    str3 := "Hello"
    
    // 区分大小写比较
    fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
    fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
    
    // 忽略大小写比较
    fmt.Printf("忽略大小写 '%s' == '%s': %t\n", str1, str2, 
        strings.EqualFold(str1, str2))
    
    // 字符串验证
    testStrings := []string{"123", "abc", "ABC", "Hello123", "  ", ""}
    
    fmt.Println("\n✅ 字符串类型检测:")
    for _, s := range testStrings {
        fmt.Printf("'%s': 数字=%t, 字母=%t, 空白=%t, 空串=%t\n", 
            s, isNumeric(s), isAlpha(s), isWhitespace(s), s == "")
    }
}

// 自定义验证函数
func isNumeric(s string) bool {
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return len(s) > 0
}

func isAlpha(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return len(s) > 0
}

func isWhitespace(s string) bool {
    for _, r := range s {
        if !unicode.IsSpace(r) {
            return false
        }
    }
    return len(s) > 0
}
```

### 字符串格式化

```go
package main

import (
    "fmt"
    "time"
)

func stringFormatting() {
    name := "张三"
    age := 25
    score := 89.75
    isStudent := true
    now := time.Now()
    
    fmt.Println("📝 字符串格式化:")
    
    // 基础格式化
    fmt.Printf("基础: 我叫%s，今年%d岁\n", name, age)
    fmt.Printf("浮点数: 分数是%.2f\n", score)
    fmt.Printf("布尔值: 是学生？%t\n", isStudent)
    
    // 宽度和对齐
    fmt.Printf("右对齐: '%10s'\n", name)
    fmt.Printf("左对齐: '%-10s'\n", name)
    fmt.Printf("数字补零: %05d\n", age)
    
    // 进制转换
    number := 255
    fmt.Printf("数字转换: 十进制=%d, 八进制=%o, 十六进制=%x, 二进制=%b\n", 
        number, number, number, number)
    
    // 时间格式化
    fmt.Printf("时间: %s\n", now.Format("2006-01-02 15:04:05"))
    
    // 类型和值
    fmt.Printf("类型和值: %T = %v\n", score, score)
    fmt.Printf("Go语法表示: %#v\n", []int{1, 2, 3})
}

// 字符串构建器
func stringBuilding() {
    fmt.Println("\n🔨 高效字符串构建:")
    
    // 使用strings.Builder构建大字符串
    var builder strings.Builder
    
    // 预分配空间提升性能
    builder.Grow(100)
    
    for i := 1; i <= 10; i++ {
        fmt.Fprintf(&builder, "第%d行：这是测试数据\n", i)
    }
    
    result := builder.String()
    fmt.Printf("构建结果 (%d字符):\n%s", len(result), result)
    
    // 重置并重用
    builder.Reset()
    builder.WriteString("重用的构建器\n")
    builder.WriteString("继续添加内容\n")
    
    fmt.Printf("重用结果: %s", builder.String())
}
```

## 正则表达式

### 正则表达式基础

```go
package main

import (
    "fmt"
    "regexp"
)

func regexpBasics() {
    fmt.Println("🔍 正则表达式基础:")
    
    // 编译正则表达式
    emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    emailRegex, err := regexp.Compile(emailPattern)
    if err != nil {
        fmt.Printf("正则表达式编译失败: %v\n", err)
        return
    }
    
    // 测试邮箱地址
    emails := []string{
        "user@example.com",
        "test.email+tag@domain.co.uk",
        "invalid-email",
        "no-at-sign.com",
        "missing@domain",
    }
    
    fmt.Println("邮箱验证结果:")
    for _, email := range emails {
        isValid := emailRegex.MatchString(email)
        status := "❌"
        if isValid {
            status = "✅"
        }
        fmt.Printf("  %s %s\n", status, email)
    }
    
    // 查找匹配
    text := "联系我们: support@company.com 或 info@company.com"
    matches := emailRegex.FindAllString(text, -1)
    
    fmt.Printf("\n在文本中找到的邮箱: %v\n", matches)
}

// 常用正则表达式模式
func commonPatterns() {
    fmt.Println("\n📋 常用正则表达式模式:")
    
    patterns := map[string]string{
        "手机号码":     `^1[3-9]\d{9}$`,
        "身份证号":     `^\d{17}[\dXx]$`,
        "IP地址":      `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`,
        "中文字符":     `[\p{Han}]+`,
        "日期格式":     `^\d{4}-\d{2}-\d{2}$`,
        "时间格式":     `^\d{2}:\d{2}:\d{2}$`,
        "URL链接":     `^https?://[^\s/$.?#].[^\s]*$`,
        "密码强度":     `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`,
    }
    
    testData := map[string][]string{
        "手机号码": {"13812345678", "15987654321", "12345678901", "138123456789"},
        "身份证号": {"110101199003074593", "12345678901234567X", "12345"},
        "IP地址":  {"192.168.1.1", "255.255.255.255", "999.999.999.999", "192.168.1"},
        "中文字符": {"你好世界", "Hello世界", "123", "English"},
        "日期格式": {"2023-12-25", "2023-1-1", "23-12-25", "2023/12/25"},
    }
    
    for name, pattern := range patterns {
        if testCases, exists := testData[name]; exists {
            fmt.Printf("\n%s 模式: %s\n", name, pattern)
            regex, err := regexp.Compile(pattern)
            if err != nil {
                fmt.Printf("  编译失败: %v\n", err)
                continue
            }
            
            for _, test := range testCases {
                isMatch := regex.MatchString(test)
                status := "❌"
                if isMatch {
                    status = "✅"
                }
                fmt.Printf("  %s %s\n", status, test)
            }
        }
    }
}

// 分组和捕获
func regexpGroups() {
    fmt.Println("\n🎯 正则表达式分组:")
    
    // 解析URL组件
    urlPattern := `^(https?)://([^/]+)(/.*)$`
    urlRegex := regexp.MustCompile(urlPattern)
    
    urls := []string{
        "https://www.example.com/path/to/page",
        "http://api.service.com/v1/users",
        "https://subdomain.domain.com/",
    }
    
    fmt.Println("URL解析结果:")
    for _, url := range urls {
        matches := urlRegex.FindStringSubmatch(url)
        if len(matches) >= 4 {
            fmt.Printf("  URL: %s\n", url)
            fmt.Printf("    协议: %s\n", matches[1])
            fmt.Printf("    主机: %s\n", matches[2])
            fmt.Printf("    路径: %s\n", matches[3])
            fmt.Println()
        }
    }
    
    // 命名分组
    logPattern := `^(?P<timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(?P<level>\w+)\] (?P<message>.+)$`
    logRegex := regexp.MustCompile(logPattern)
    
    logLine := "2023-12-25 15:30:45 [ERROR] Database connection failed"
    
    if matches := logRegex.FindStringSubmatch(logLine); matches != nil {
        names := logRegex.SubexpNames()
        fmt.Println("日志解析结果:")
        for i, match := range matches {
            if i > 0 && names[i] != "" {
                fmt.Printf("  %s: %s\n", names[i], match)
            }
        }
    }
}
```

### 文本替换和处理

```go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

func textProcessing() {
    fmt.Println("🔄 文本处理和替换:")
    
    // 替换敏感词
    content := "这是一个糟糕的、愚蠢的想法，真是太差劲了！"
    badWords := []string{"糟糕", "愚蠢", "差劲"}
    
    // 构建敏感词正则
    pattern := `(` + strings.Join(badWords, "|") + `)`
    badWordRegex := regexp.MustCompile(pattern)
    
    cleaned := badWordRegex.ReplaceAllStringFunc(content, func(match string) string {
        return strings.Repeat("*", len(match))
    })
    
    fmt.Printf("原文: %s\n", content)
    fmt.Printf("过滤后: %s\n", cleaned)
    
    // 格式化电话号码
    phoneNumbers := []string{
        "13812345678",
        "15987654321",
        "17712345678",
    }
    
    phoneRegex := regexp.MustCompile(`^(\d{3})(\d{4})(\d{4})$`)
    
    fmt.Println("\n📞 电话号码格式化:")
    for _, phone := range phoneNumbers {
        formatted := phoneRegex.ReplaceAllString(phone, "$1-$2-$3")
        fmt.Printf("%s -> %s\n", phone, formatted)
    }
    
    // 提取链接
    htmlContent := `
        访问我们的网站 <a href="https://www.example.com">Example</a>
        或者查看 <a href="https://blog.example.com/post/123">博客文章</a>
        联系邮箱: <a href="mailto:contact@example.com">contact@example.com</a>
    `
    
    linkRegex := regexp.MustCompile(`<a href="([^"]+)"[^>]*>([^<]+)</a>`)
    links := linkRegex.FindAllStringSubmatch(htmlContent, -1)
    
    fmt.Println("\n🔗 提取的链接:")
    for _, link := range links {
        fmt.Printf("  文本: %s\n", link[2])
        fmt.Printf("  链接: %s\n", link[1])
        fmt.Println()
    }
}
```

## 实战项目：文本分析工具

让我们构建一个完整的文本分析工具：

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strings"
    "unicode"
)

// 文本统计结构
type TextStats struct {
    CharCount      int
    WordCount      int
    LineCount      int
    SentenceCount  int
    ParagraphCount int
    WordFrequency  map[string]int
    CommonWords    []WordFreq
}

type WordFreq struct {
    Word  string
    Count int
}

// 文本分析器
type TextAnalyzer struct {
    stopWords map[string]bool
    patterns  map[string]*regexp.Regexp
}

func NewTextAnalyzer() *TextAnalyzer {
    // 中英文停用词
    stopWords := map[string]bool{
        "the": true, "and": true, "or": true, "but": true, "in": true,
        "on": true, "at": true, "to": true, "for": true, "of": true,
        "with": true, "by": true, "是": true, "的": true, "了": true,
        "在": true, "有": true, "和": true, "就": true, "不": true,
        "都": true, "会": true, "说": true, "我": true, "你": true,
    }
    
    // 预编译正则表达式
    patterns := map[string]*regexp.Regexp{
        "word":      regexp.MustCompile(`\b[\p{L}\p{N}]+\b`),
        "sentence":  regexp.MustCompile(`[.!?。！？]+`),
        "paragraph": regexp.MustCompile(`\n\s*\n`),
        "email":     regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`),
        "url":       regexp.MustCompile(`https?://[^\s]+`),
        "phone":     regexp.MustCompile(`1[3-9]\d{9}`),
        "chinese":   regexp.MustCompile(`[\p{Han}]+`),
        "english":   regexp.MustCompile(`[a-zA-Z]+`),
        "number":    regexp.MustCompile(`\d+`),
    }
    
    return &TextAnalyzer{
        stopWords: stopWords,
        patterns:  patterns,
    }
}

// 分析文本
func (ta *TextAnalyzer) AnalyzeText(text string) *TextStats {
    stats := &TextStats{
        WordFrequency: make(map[string]int),
    }
    
    // 基础统计
    stats.CharCount = len([]rune(text)) // 正确计算Unicode字符数
    stats.LineCount = strings.Count(text, "\n") + 1
    
    // 词频统计
    words := ta.patterns["word"].FindAllString(text, -1)
    stats.WordCount = len(words)
    
    for _, word := range words {
        word = strings.ToLower(word)
        // 跳过停用词和单字符
        if !ta.stopWords[word] && len(word) > 1 {
            stats.WordFrequency[word]++
        }
    }
    
    // 句子统计
    sentences := ta.patterns["sentence"].FindAllString(text, -1)
    stats.SentenceCount = len(sentences)
    
    // 段落统计
    paragraphs := ta.patterns["paragraph"].Split(text, -1)
    stats.ParagraphCount = len(paragraphs)
    
    // 生成词频排序
    stats.CommonWords = ta.getTopWords(stats.WordFrequency, 10)
    
    return stats
}

// 获取高频词
func (ta *TextAnalyzer) getTopWords(frequency map[string]int, top int) []WordFreq {
    var words []WordFreq
    for word, count := range frequency {
        words = append(words, WordFreq{Word: word, Count: count})
    }
    
    sort.Slice(words, func(i, j int) bool {
        return words[i].Count > words[j].Count
    })
    
    if len(words) > top {
        words = words[:top]
    }
    
    return words
}

// 提取特定内容
func (ta *TextAnalyzer) ExtractContent(text string) map[string][]string {
    results := make(map[string][]string)
    
    extractors := map[string]string{
        "邮箱": "email",
        "网址": "url", 
        "电话": "phone",
        "中文": "chinese",
        "英文": "english",
        "数字": "number",
    }
    
    for name, pattern := range extractors {
        if regex, exists := ta.patterns[pattern]; exists {
            matches := regex.FindAllString(text, -1)
            // 去重
            uniqueMatches := removeDuplicates(matches)
            if len(uniqueMatches) > 0 {
                results[name] = uniqueMatches
            }
        }
    }
    
    return results
}

// 文本质量评估
func (ta *TextAnalyzer) AssessTextQuality(stats *TextStats) map[string]interface{} {
    assessment := make(map[string]interface{})
    
    // 平均词长
    if stats.WordCount > 0 {
        avgWordLength := float64(stats.CharCount) / float64(stats.WordCount)
        assessment["平均词长"] = fmt.Sprintf("%.2f", avgWordLength)
    }
    
    // 句子长度
    if stats.SentenceCount > 0 {
        avgWordsPerSentence := float64(stats.WordCount) / float64(stats.SentenceCount)
        assessment["平均句长"] = fmt.Sprintf("%.2f词", avgWordsPerSentence)
        
        var readability string
        switch {
        case avgWordsPerSentence < 10:
            readability = "简单易读"
        case avgWordsPerSentence < 20:
            readability = "适中"
        case avgWordsPerSentence < 30:
            readability = "稍复杂"
        default:
            readability = "复杂难读"
        }
        assessment["可读性"] = readability
    }
    
    // 词汇丰富度
    if stats.WordCount > 0 {
        uniqueWords := len(stats.WordFrequency)
        diversity := float64(uniqueWords) / float64(stats.WordCount)
        assessment["词汇丰富度"] = fmt.Sprintf("%.2f%%", diversity*100)
    }
    
    return assessment
}

// 生成文本报告
func (ta *TextAnalyzer) GenerateReport(text string) {
    fmt.Println("📊 文本分析报告")
    fmt.Println("================")
    
    // 基础统计
    stats := ta.AnalyzeText(text)
    
    fmt.Printf("📝 基础统计:\n")
    fmt.Printf("  字符数: %d\n", stats.CharCount)
    fmt.Printf("  词汇数: %d\n", stats.WordCount)
    fmt.Printf("  句子数: %d\n", stats.SentenceCount)
    fmt.Printf("  段落数: %d\n", stats.ParagraphCount)
    fmt.Printf("  行数: %d\n", stats.LineCount)
    
    // 高频词
    fmt.Printf("\n🔥 高频词汇 (前10):\n")
    for i, word := range stats.CommonWords {
        fmt.Printf("  %d. %s (%d次)\n", i+1, word.Word, word.Count)
    }
    
    // 提取内容
    extracted := ta.ExtractContent(text)
    if len(extracted) > 0 {
        fmt.Printf("\n🔍 提取内容:\n")
        for category, items := range extracted {
            fmt.Printf("  %s: %v\n", category, items)
        }
    }
    
    // 质量评估
    quality := ta.AssessTextQuality(stats)
    fmt.Printf("\n📈 文本质量:\n")
    for metric, value := range quality {
        fmt.Printf("  %s: %v\n", metric, value)
    }
}

// 交互式文本分析器
type InteractiveAnalyzer struct {
    analyzer *TextAnalyzer
    scanner  *bufio.Scanner
}

func NewInteractiveAnalyzer() *InteractiveAnalyzer {
    return &InteractiveAnalyzer{
        analyzer: NewTextAnalyzer(),
        scanner:  bufio.NewScanner(os.Stdin),
    }
}

func (ia *InteractiveAnalyzer) Start() {
    fmt.Println("🔍 交互式文本分析工具")
    fmt.Println("====================")
    
    for {
        ia.showMenu()
        choice := ia.getInput()
        
        switch choice {
        case "1":
            ia.analyzeTextInput()
        case "2":
            ia.analyzeFile()
        case "3":
            ia.batchAnalyze()
        case "4":
            ia.showPatterns()
        case "5":
            fmt.Println("👋 退出程序")
            return
        default:
            fmt.Println("❌ 无效选择，请重试")
        }
        
        fmt.Println()
    }
}

func (ia *InteractiveAnalyzer) showMenu() {
    fmt.Println("📋 选择分析模式:")
    fmt.Println("1. 分析输入文本")
    fmt.Println("2. 分析文件")
    fmt.Println("3. 批量分析")
    fmt.Println("4. 显示支持的模式")
    fmt.Println("5. 退出")
    fmt.Print("请选择 (1-5): ")
}

func (ia *InteractiveAnalyzer) getInput() string {
    ia.scanner.Scan()
    return strings.TrimSpace(ia.scanner.Text())
}

func (ia *InteractiveAnalyzer) analyzeTextInput() {
    fmt.Println("请输入要分析的文本 (输入空行结束):")
    
    var lines []string
    for ia.scanner.Scan() {
        line := ia.scanner.Text()
        if line == "" {
            break
        }
        lines = append(lines, line)
    }
    
    if len(lines) == 0 {
        fmt.Println("❌ 没有输入文本")
        return
    }
    
    text := strings.Join(lines, "\n")
    ia.analyzer.GenerateReport(text)
}

func (ia *InteractiveAnalyzer) analyzeFile() {
    fmt.Print("请输入文件路径: ")
    filepath := ia.getInput()
    
    content, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Printf("❌ 读取文件失败: %v\n", err)
        return
    }
    
    fmt.Printf("📄 分析文件: %s\n", filepath)
    ia.analyzer.GenerateReport(string(content))
}

func (ia *InteractiveAnalyzer) batchAnalyze() {
    fmt.Println("📁 批量分析演示:")
    
    sampleTexts := map[string]string{
        "技术文档": `
            Go语言是Google开发的开源编程语言。它具有简洁的语法、高效的性能和强大的并发支持。
            Go语言特别适合开发网络服务、分布式系统和云原生应用。
            主要特性包括：垃圾回收、静态类型、编译型语言、丰富的标准库。
        `,
        "新闻报道": `
            今日科技新闻：人工智能技术在各个领域的应用越来越广泛。
            从自动驾驶汽车到智能语音助手，AI正在改变我们的生活方式。
            专家预测，未来十年将是AI技术发展的黄金时期。
            联系方式：news@tech.com 或访问 https://tech.com/ai-news
        `,
        "商务邮件": `
            尊敬的客户您好，
            
            感谢您对我们产品的关注。我们的技术团队已经完成了新版本的开发。
            如有任何问题，请联系我们：support@company.com 或致电 400-123-4567。
            
            访问我们的官网：https://www.company.com 了解更多信息。
            
            祝好！
        `,
    }
    
    for title, text := range sampleTexts {
        fmt.Printf("\n=== %s ===\n", title)
        ia.analyzer.GenerateReport(text)
    }
}

func (ia *InteractiveAnalyzer) showPatterns() {
    fmt.Println("🎯 支持的匹配模式:")
    
    patterns := map[string]string{
        "邮箱地址": `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`,
        "网址链接": `https?://[^\s]+`,
        "手机号码": `1[3-9]\d{9}`,
        "中文字符": `[\p{Han}]+`,
        "英文单词": `[a-zA-Z]+`,
        "数字":   `\d+`,
    }
    
    examples := map[string]string{
        "邮箱地址": "user@example.com",
        "网址链接": "https://www.example.com",
        "手机号码": "13812345678",
        "中文字符": "你好世界",
        "英文单词": "Hello",
        "数字":   "12345",
    }
    
    for name, pattern := range patterns {
        example := examples[name]
        fmt.Printf("  %s: %s\n", name, pattern)
        fmt.Printf("    示例: %s\n", example)
        fmt.Println()
    }
}

// 工具函数
func removeDuplicates(slice []string) []string {
    keys := make(map[string]bool)
    var result []string
    
    for _, item := range slice {
        if !keys[item] {
            keys[item] = true
            result = append(result, item)
        }
    }
    
    return result
}

func main() {
    // 创建交互式分析器
    analyzer := NewInteractiveAnalyzer()
    
    // 启动交互式分析
    analyzer.Start()
}
```

## 性能优化技巧

### 1. 字符串连接优化

```go
// 低效方式 - 频繁创建新字符串
func inefficientConcat(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s // 每次都创建新字符串
    }
    return result
}

// 高效方式 - 使用strings.Builder
func efficientConcat(strs []string) string {
    var builder strings.Builder
    builder.Grow(estimateSize(strs)) // 预分配空间
    
    for _, s := range strs {
        builder.WriteString(s)
    }
    
    return builder.String()
}

func estimateSize(strs []string) int {
    total := 0
    for _, s := range strs {
        total += len(s)
    }
    return total
}
```

### 2. 正则表达式优化

```go
// 预编译正则表达式，避免重复编译
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func validateEmails(emails []string) []bool {
    results := make([]bool, len(emails))
    for i, email := range emails {
        results[i] = emailRegex.MatchString(email)
    }
    return results
}
```

### 3. Unicode处理

```go
func safeStringLength(s string) int {
    return len([]rune(s)) // 正确计算Unicode字符数
}

func isChineseChar(r rune) bool {
    return unicode.Is(unicode.Scripts["Han"], r)
}
```

## 本章小结

Go字符串处理的核心要点：

- **基础操作**：掌握strings包的常用函数和字符串操作
- **格式化输出**：使用fmt包进行灵活的字符串格式化
- **正则表达式**：利用regexp包进行模式匹配和文本处理
- **性能优化**：合理使用strings.Builder和预编译正则表达式
- **Unicode支持**：正确处理多语言文本和字符编码

::: tip 练习建议
1. 实现一个日志解析器
2. 开发文本内容过滤工具
3. 创建简单的模板引擎
4. 构建数据验证器
:::