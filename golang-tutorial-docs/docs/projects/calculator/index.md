# 计算器程序

让我们从一个实用的命令行计算器开始实战项目！这个项目将帮你掌握Go语言的基础应用开发。

## 📋 项目需求

### 功能要求
- ✅ 支持基本四则运算（+, -, *, /）
- ✅ 支持小括号优先级
- ✅ 错误处理和输入验证
- ✅ 计算历史记录
- ✅ 交互式命令行界面
- ✅ 帮助信息和退出功能

### 技术要求
- 使用Go标准库
- 良好的代码结构
- 完整的错误处理
- 单元测试覆盖
- 清晰的用户界面

## 🏗️ 项目架构

```
calculator/
├── cmd/
│   └── main.go          # 程序入口
├── internal/
│   ├── calculator/      # 计算器核心逻辑
│   │   ├── calculator.go
│   │   ├── parser.go
│   │   └── history.go
│   └── ui/             # 用户界面
│       └── cli.go
├── pkg/
│   └── mathutils/      # 数学工具函数
│       └── utils.go
├── test/               # 测试文件
│   ├── calculator_test.go
│   └── parser_test.go
├── go.mod
├── go.sum
└── README.md
```

## 💻 实现步骤

### 第一步：项目初始化

首先创建项目目录并初始化Go模块：

```bash
mkdir calculator
cd calculator
go mod init github.com/yourname/calculator
```

### 第二步：核心计算器实现

创建 `internal/calculator/calculator.go`：

```go
package calculator

import (
    "errors"
    "fmt"
    "math"
)

// Calculator 计算器结构体
type Calculator struct {
    history *History
}

// New 创建新的计算器实例
func New() *Calculator {
    return &Calculator{
        history: NewHistory(),
    }
}

// Calculate 执行计算
func (c *Calculator) Calculate(expression string) (float64, error) {
    if expression == "" {
        return 0, errors.New("表达式不能为空")
    }
    
    // 解析表达式
    parser := NewParser(expression)
    result, err := parser.Parse()
    if err != nil {
        return 0, fmt.Errorf("解析错误: %v", err)
    }
    
    // 记录历史
    c.history.Add(expression, result)
    
    return result, nil
}

// GetHistory 获取计算历史
func (c *Calculator) GetHistory() []HistoryEntry {
    return c.history.GetAll()
}

// ClearHistory 清空历史记录
func (c *Calculator) ClearHistory() {
    c.history.Clear()
}

// ValidateExpression 验证表达式是否有效
func (c *Calculator) ValidateExpression(expression string) error {
    parser := NewParser(expression)
    _, err := parser.Parse()
    return err
}
```

### 第三步：表达式解析器

创建 `internal/calculator/parser.go`：

```go
package calculator

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

// Token 词法单元
type Token struct {
    Type  TokenType
    Value string
}

// TokenType 词法单元类型
type TokenType int

const (
    NUMBER TokenType = iota
    PLUS
    MINUS
    MULTIPLY
    DIVIDE
    LPAREN
    RPAREN
    EOF
)

// Parser 表达式解析器
type Parser struct {
    expression string
    tokens     []Token
    current    int
}

// NewParser 创建新的解析器
func NewParser(expression string) *Parser {
    return &Parser{
        expression: strings.ReplaceAll(expression, " ", ""),
        current:    0,
    }
}

// Parse 解析表达式并计算结果
func (p *Parser) Parse() (float64, error) {
    // 词法分析
    err := p.tokenize()
    if err != nil {
        return 0, err
    }
    
    // 语法分析和计算
    result, err := p.parseExpression()
    if err != nil {
        return 0, err
    }
    
    // 检查是否还有未处理的token
    if p.current < len(p.tokens) && p.tokens[p.current].Type != EOF {
        return 0, errors.New("表达式末尾有多余的字符")
    }
    
    return result, nil
}

// tokenize 词法分析
func (p *Parser) tokenize() error {
    p.tokens = []Token{}
    
    for i, char := range p.expression {
        switch {
        case unicode.IsDigit(char) || char == '.':
            // 解析数字
            numStr := ""
            j := i
            dotCount := 0
            
            for j < len(p.expression) {
                ch := rune(p.expression[j])
                if unicode.IsDigit(ch) {
                    numStr += string(ch)
                } else if ch == '.' {
                    if dotCount > 0 {
                        return fmt.Errorf("数字格式错误：多个小数点 '%s'", numStr)
                    }
                    dotCount++
                    numStr += string(ch)
                } else {
                    break
                }
                j++
            }
            
            if numStr == "." {
                return errors.New("无效的数字格式")
            }
            
            p.tokens = append(p.tokens, Token{NUMBER, numStr})
            
            // 跳过已处理的字符
            for k := i + 1; k < j; k++ {
                // 这些字符已经在上面的循环中处理了
            }
            
        case char == '+':
            p.tokens = append(p.tokens, Token{PLUS, "+"})
        case char == '-':
            p.tokens = append(p.tokens, Token{MINUS, "-"})
        case char == '*':
            p.tokens = append(p.tokens, Token{MULTIPLY, "*"})
        case char == '/':
            p.tokens = append(p.tokens, Token{DIVIDE, "/"})
        case char == '(':
            p.tokens = append(p.tokens, Token{LPAREN, "("})
        case char == ')':
            p.tokens = append(p.tokens, Token{RPAREN, ")"})
        default:
            return fmt.Errorf("无效字符: '%c'", char)
        }
    }
    
    p.tokens = append(p.tokens, Token{EOF, ""})
    return nil
}

// parseExpression 解析表达式 (处理 + 和 -)
func (p *Parser) parseExpression() (float64, error) {
    result, err := p.parseTerm()
    if err != nil {
        return 0, err
    }
    
    for p.current < len(p.tokens) {
        token := p.tokens[p.current]
        if token.Type == PLUS {
            p.current++
            right, err := p.parseTerm()
            if err != nil {
                return 0, err
            }
            result += right
        } else if token.Type == MINUS {
            p.current++
            right, err := p.parseTerm()
            if err != nil {
                return 0, err
            }
            result -= right
        } else {
            break
        }
    }
    
    return result, nil
}

// parseTerm 解析项 (处理 * 和 /)
func (p *Parser) parseTerm() (float64, error) {
    result, err := p.parseFactor()
    if err != nil {
        return 0, err
    }
    
    for p.current < len(p.tokens) {
        token := p.tokens[p.current]
        if token.Type == MULTIPLY {
            p.current++
            right, err := p.parseFactor()
            if err != nil {
                return 0, err
            }
            result *= right
        } else if token.Type == DIVIDE {
            p.current++
            right, err := p.parseFactor()
            if err != nil {
                return 0, err
            }
            if right == 0 {
                return 0, errors.New("除零错误")
            }
            result /= right
        } else {
            break
        }
    }
    
    return result, nil
}

// parseFactor 解析因子 (处理数字和括号)
func (p *Parser) parseFactor() (float64, error) {
    if p.current >= len(p.tokens) {
        return 0, errors.New("意外的表达式结束")
    }
    
    token := p.tokens[p.current]
    
    switch token.Type {
    case NUMBER:
        p.current++
        value, err := strconv.ParseFloat(token.Value, 64)
        if err != nil {
            return 0, fmt.Errorf("无效的数字: %s", token.Value)
        }
        return value, nil
        
    case MINUS:
        // 处理负数
        p.current++
        value, err := p.parseFactor()
        if err != nil {
            return 0, err
        }
        return -value, nil
        
    case PLUS:
        // 处理正数
        p.current++
        return p.parseFactor()
        
    case LPAREN:
        p.current++
        result, err := p.parseExpression()
        if err != nil {
            return 0, err
        }
        
        if p.current >= len(p.tokens) || p.tokens[p.current].Type != RPAREN {
            return 0, errors.New("缺少右括号")
        }
        p.current++
        return result, nil
        
    default:
        return 0, fmt.Errorf("意外的token: %s", token.Value)
    }
}
```

### 第四步：历史记录管理

创建 `internal/calculator/history.go`：

```go
package calculator

import (
    "fmt"
    "time"
)

// HistoryEntry 历史记录条目
type HistoryEntry struct {
    ID         int       `json:"id"`
    Expression string    `json:"expression"`
    Result     float64   `json:"result"`
    Timestamp  time.Time `json:"timestamp"`
}

// String 格式化历史记录显示
func (h HistoryEntry) String() string {
    return fmt.Sprintf("%d. %s = %.6g [%s]", 
        h.ID, h.Expression, h.Result, h.Timestamp.Format("15:04:05"))
}

// History 历史记录管理器
type History struct {
    entries []HistoryEntry
    nextID  int
}

// NewHistory 创建新的历史记录管理器
func NewHistory() *History {
    return &History{
        entries: make([]HistoryEntry, 0),
        nextID:  1,
    }
}

// Add 添加历史记录
func (h *History) Add(expression string, result float64) {
    entry := HistoryEntry{
        ID:         h.nextID,
        Expression: expression,
        Result:     result,
        Timestamp:  time.Now(),
    }
    
    h.entries = append(h.entries, entry)
    h.nextID++
    
    // 限制历史记录数量，避免内存占用过多
    const maxHistorySize = 100
    if len(h.entries) > maxHistorySize {
        h.entries = h.entries[len(h.entries)-maxHistorySize:]
    }
}

// GetAll 获取所有历史记录
func (h *History) GetAll() []HistoryEntry {
    // 返回副本，避免外部修改
    result := make([]HistoryEntry, len(h.entries))
    copy(result, h.entries)
    return result
}

// GetLast 获取最近的n条记录
func (h *History) GetLast(n int) []HistoryEntry {
    if n <= 0 {
        return []HistoryEntry{}
    }
    
    start := len(h.entries) - n
    if start < 0 {
        start = 0
    }
    
    result := make([]HistoryEntry, len(h.entries)-start)
    copy(result, h.entries[start:])
    return result
}

// Clear 清空历史记录
func (h *History) Clear() {
    h.entries = make([]HistoryEntry, 0)
    h.nextID = 1
}

// Count 获取历史记录数量
func (h *History) Count() int {
    return len(h.entries)
}
```

### 第五步：命令行界面

创建 `internal/ui/cli.go`：

```go
package ui

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    
    "github.com/yourname/calculator/internal/calculator"
)

// CLI 命令行界面
type CLI struct {
    calc     *calculator.Calculator
    scanner  *bufio.Scanner
    running  bool
}

// NewCLI 创建新的命令行界面
func NewCLI() *CLI {
    return &CLI{
        calc:    calculator.New(),
        scanner: bufio.NewScanner(os.Stdin),
        running: true,
    }
}

// Run 运行命令行界面
func (cli *CLI) Run() {
    cli.printWelcome()
    
    for cli.running {
        cli.printPrompt()
        
        if !cli.scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(cli.scanner.Text())
        cli.handleInput(input)
    }
    
    fmt.Println("再见!")
}

// printWelcome 打印欢迎信息
func (cli *CLI) printWelcome() {
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Println("🧮 Go 计算器")
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Println()
    fmt.Println("支持的操作：")
    fmt.Println("  • 基本运算：+, -, *, /")
    fmt.Println("  • 括号支持：(1 + 2) * 3")
    fmt.Println("  • 小数运算：3.14 * 2")
    fmt.Println()
    fmt.Println("特殊命令：")
    fmt.Println("  • help     - 显示帮助信息")
    fmt.Println("  • history  - 查看计算历史")
    fmt.Println("  • clear    - 清空历史记录")
    fmt.Println("  • exit     - 退出程序")
    fmt.Println()
}

// printPrompt 打印输入提示符
func (cli *CLI) printPrompt() {
    fmt.Print("calc> ")
}

// handleInput 处理用户输入
func (cli *CLI) handleInput(input string) {
    if input == "" {
        return
    }
    
    // 处理特殊命令
    switch strings.ToLower(input) {
    case "help", "h":
        cli.showHelp()
        return
    case "history", "hist":
        cli.showHistory()
        return
    case "clear", "c":
        cli.clearHistory()
        return
    case "exit", "quit", "q":
        cli.running = false
        return
    }
    
    // 计算表达式
    result, err := cli.calc.Calculate(input)
    if err != nil {
        fmt.Printf("❌ 错误: %v\n", err)
        return
    }
    
    fmt.Printf("= %.6g\n", result)
}

// showHelp 显示帮助信息
func (cli *CLI) showHelp() {
    fmt.Println()
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Println("📖 帮助信息")
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Println()
    fmt.Println("基本运算符：")
    fmt.Println("  +  加法   例如：3 + 5")
    fmt.Println("  -  减法   例如：10 - 3")
    fmt.Println("  *  乘法   例如：4 * 7")
    fmt.Println("  /  除法   例如：15 / 3")
    fmt.Println()
    fmt.Println("高级功能：")
    fmt.Println("  ()  括号改变优先级   例如：(1 + 2) * 3")
    fmt.Println("  .   小数点         例如：3.14 * 2")
    fmt.Println("  -   负数           例如：-5 + 3")
    fmt.Println()
    fmt.Println("示例表达式：")
    fmt.Println("  • 2 + 3 * 4")
    fmt.Println("  • (10 - 5) / 2.5")
    fmt.Println("  • -3 + 7")
    fmt.Println("  • ((1 + 2) * 3) / 2")
    fmt.Println()
    fmt.Println("命令：")
    fmt.Println("  help     显示此帮助信息")
    fmt.Println("  history  查看计算历史记录")
    fmt.Println("  clear    清空历史记录")
    fmt.Println("  exit     退出计算器")
    fmt.Println()
}

// showHistory 显示历史记录
func (cli *CLI) showHistory() {
    history := cli.calc.GetHistory()
    
    if len(history) == 0 {
        fmt.Println("📝 暂无计算历史")
        return
    }
    
    fmt.Println()
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Printf("📝 计算历史 (共 %d 条记录)\n", len(history))
    fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    fmt.Println()
    
    // 显示最近的10条记录
    start := len(history) - 10
    if start < 0 {
        start = 0
    }
    
    for i := start; i < len(history); i++ {
        fmt.Println(history[i].String())
    }
    
    if len(history) > 10 {
        fmt.Printf("\n只显示最近 10 条记录，总共 %d 条\n", len(history))
    }
    fmt.Println()
}

// clearHistory 清空历史记录
func (cli *CLI) clearHistory() {
    cli.calc.ClearHistory()
    fmt.Println("🗑️  历史记录已清空")
}
```

### 第六步：主程序入口

创建 `cmd/main.go`：

```go
package main

import (
    "github.com/yourname/calculator/internal/ui"
)

func main() {
    cli := ui.NewCLI()
    cli.Run()
}
```

### 第七步：单元测试

创建 `test/calculator_test.go`：

```go
package test

import (
    "testing"
    
    "github.com/yourname/calculator/internal/calculator"
)

func TestCalculator_BasicOperations(t *testing.T) {
    calc := calculator.New()
    
    tests := []struct {
        expression string
        expected   float64
        shouldErr  bool
    }{
        {"2 + 3", 5, false},
        {"10 - 5", 5, false},
        {"4 * 7", 28, false},
        {"15 / 3", 5, false},
        {"2 + 3 * 4", 14, false},
        {"(2 + 3) * 4", 20, false},
        {"10 / 0", 0, true},
        {"", 0, true},
        {"2 +", 0, true},
        {"(2 + 3", 0, true},
        {"2 + 3)", 0, true},
    }
    
    for _, test := range tests {
        result, err := calc.Calculate(test.expression)
        
        if test.shouldErr {
            if err == nil {
                t.Errorf("期望 '%s' 产生错误，但没有", test.expression)
            }
        } else {
            if err != nil {
                t.Errorf("不期望 '%s' 产生错误，但得到: %v", test.expression, err)
            } else if result != test.expected {
                t.Errorf("'%s' 期望结果 %.2f，实际得到 %.2f", 
                    test.expression, test.expected, result)
            }
        }
    }
}

func TestCalculator_History(t *testing.T) {
    calc := calculator.New()
    
    // 添加一些计算
    expressions := []string{"2 + 3", "5 * 4", "10 / 2"}
    for _, expr := range expressions {
        _, err := calc.Calculate(expr)
        if err != nil {
            t.Fatalf("计算 '%s' 时出错: %v", expr, err)
        }
    }
    
    // 检查历史记录
    history := calc.GetHistory()
    if len(history) != len(expressions) {
        t.Errorf("期望历史记录数量 %d，实际得到 %d", len(expressions), len(history))
    }
    
    // 清空历史记录
    calc.ClearHistory()
    history = calc.GetHistory()
    if len(history) != 0 {
        t.Errorf("清空后期望历史记录数量 0，实际得到 %d", len(history))
    }
}
```

创建 `test/parser_test.go`：

```go
package test

import (
    "testing"
    
    "github.com/yourname/calculator/internal/calculator"
)

func TestParser_ComplexExpressions(t *testing.T) {
    tests := []struct {
        expression string
        expected   float64
        shouldErr  bool
    }{
        {"1 + 2 * 3", 7, false},
        {"(1 + 2) * 3", 9, false},
        {"2 * 3 + 4", 10, false},
        {"2 * (3 + 4)", 14, false},
        {"-5 + 3", -2, false},
        {"3 + -5", -2, false},
        {"(-1 + 2) * 3", 3, false},
        {"3.14 * 2", 6.28, false},
        {"10 / 2.5", 4, false},
        {"((1 + 2) * 3) / 2", 4.5, false},
        {"1 + 2 * 3 - 4 / 2", 5, false},
    }
    
    for _, test := range tests {
        parser := calculator.NewParser(test.expression)
        result, err := parser.Parse()
        
        if test.shouldErr {
            if err == nil {
                t.Errorf("期望 '%s' 产生错误，但没有", test.expression)
            }
        } else {
            if err != nil {
                t.Errorf("不期望 '%s' 产生错误，但得到: %v", test.expression, err)
            } else if result != test.expected {
                t.Errorf("'%s' 期望结果 %.6g，实际得到 %.6g", 
                    test.expression, test.expected, result)
            }
        }
    }
}

func TestParser_InvalidExpressions(t *testing.T) {
    invalidExpressions := []string{
        "2 +",
        "+ 2",
        "2 * * 3",
        "(2 + 3",
        "2 + 3)",
        "2..3",
        "2abc",
        "2 + abc",
        "2 / 0",
        "",
        "   ",
    }
    
    for _, expr := range invalidExpressions {
        parser := calculator.NewParser(expr)
        _, err := parser.Parse()
        
        if err == nil {
            t.Errorf("期望 '%s' 产生错误，但没有", expr)
        }
    }
}
```

## 🧪 测试运行

运行测试确保代码正确：

```bash
# 运行所有测试
go test ./test/...

# 运行测试并显示覆盖率
go test -cover ./test/...

# 运行程序
go run cmd/main.go
```

## 🎯 扩展功能

完成基础功能后，可以尝试添加以下扩展：

### 1. 更多数学函数
```go
// 支持 sin, cos, tan, sqrt, pow 等函数
"sin(3.14159/2)"  // = 1
"sqrt(16)"        // = 4
"pow(2, 8)"       // = 256
```

### 2. 变量支持
```go
// 支持变量定义和使用
"x = 10"
"y = 5"
"x + y * 2"  // = 20
```

### 3. 历史记录文件保存
```go
// 将历史记录保存到文件
func (h *History) SaveToFile(filename string) error
func (h *History) LoadFromFile(filename string) error
```

### 4. 表达式验证增强
```go
// 更详细的错误信息和建议
"语法错误：第5个字符处缺少右括号"
"建议：是否想输入 '2 * 3' ？"
```

## 📝 项目总结

### 学到的技能
- ✅ Go项目结构组织
- ✅ 词法分析和语法解析
- ✅ 错误处理机制
- ✅ 单元测试编写
- ✅ 命令行界面开发
- ✅ 包的设计和使用

### 代码质量亮点
- 清晰的职责分离
- 完整的错误处理
- 良好的测试覆盖
- 用户友好的界面
- 可扩展的架构设计

## 🎯 下一步

完成计算器项目后，你可以：

1. **代码审查** - 优化代码结构和性能
2. **功能扩展** - 添加更多数学函数
3. **部署打包** - 学习Go程序的构建和分发
4. **继续学习** - 进入下一个项目 [待办事项CLI](../todo-cli/)

恭喜完成第一个Go实战项目！🎉 