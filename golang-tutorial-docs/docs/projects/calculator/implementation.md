# 💻 代码实现 - 计算器项目

本文档详细介绍计算器项目的完整代码实现，包括逐行解析和运行示例。

## 🚀 完整代码实现

### main.go - 完整实现

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// ===== 数据结构定义 =====

// CommandType 命令类型枚举
type CommandType int

const (
	Calculate CommandType = iota
	ShowHistory
	ShowHelp
	Clear
	Quit
)

// Command 表示用户输入的命令
type Command struct {
	Type       CommandType
	Expression string
}

// TokenType Token类型枚举
type TokenType int

const (
	NUMBER TokenType = iota
	OPERATOR
	LEFT_PAREN
	RIGHT_PAREN
)

// Token 表示表达式中的一个词法单元
type Token struct {
	Type  TokenType
	Value string
}

// HistoryEntry 历史记录条目
type HistoryEntry struct {
	Expression string
	Result     float64
	Timestamp  time.Time
}

// ===== 核心组件 =====

// ExpressionParser 表达式解析器
type ExpressionParser struct {
	tokens []Token
	pos    int
}

// NewExpressionParser 创建新的表达式解析器
func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

// Parse 解析并计算表达式
func (p *ExpressionParser) Parse(expression string) (float64, error) {
	// 重置解析器状态
	p.pos = 0
	
	// 词法分析
	tokens, err := p.tokenize(expression)
	if err != nil {
		return 0, fmt.Errorf("词法分析错误: %v", err)
	}
	
	if len(tokens) == 0 {
		return 0, fmt.Errorf("表达式为空")
	}
	
	p.tokens = tokens
	
	// 语法分析和计算
	result, err := p.parseExpression()
	if err != nil {
		return 0, fmt.Errorf("语法分析错误: %v", err)
	}
	
	// 检查是否还有未处理的token
	if p.pos < len(p.tokens) {
		return 0, fmt.Errorf("表达式语法错误：存在多余的字符")
	}
	
	return result, nil
}

// tokenize 词法分析：将表达式分解为tokens
func (p *ExpressionParser) tokenize(expression string) ([]Token, error) {
	var tokens []Token
	var current strings.Builder
	
	for i, char := range expression {
		switch {
		case unicode.IsDigit(char) || char == '.':
			// 数字字符和小数点
			current.WriteRune(char)
			
		case char == '+' || char == '-' || char == '*' || char == '/':
			// 运算符
			if current.Len() > 0 {
				tokens = append(tokens, Token{NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{OPERATOR, string(char)})
			
		case char == '(':
			// 左括号
			if current.Len() > 0 {
				tokens = append(tokens, Token{NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{LEFT_PAREN, "("})
			
		case char == ')':
			// 右括号
			if current.Len() > 0 {
				tokens = append(tokens, Token{NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{RIGHT_PAREN, ")"})
			
		case char == ' ' || char == '\t':
			// 忽略空白字符
			if current.Len() > 0 {
				tokens = append(tokens, Token{NUMBER, current.String()})
				current.Reset()
			}
			
		default:
			return nil, fmt.Errorf("非法字符 '%c' 在位置 %d", char, i)
		}
	}
	
	// 处理最后一个数字
	if current.Len() > 0 {
		tokens = append(tokens, Token{NUMBER, current.String()})
	}
	
	return tokens, nil
}

// parseExpression 解析表达式（处理 + 和 - 运算符）
func (p *ExpressionParser) parseExpression() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	
	for p.pos < len(p.tokens) {
		token := p.tokens[p.pos]
		
		// 只处理 + 和 - 运算符
		if token.Type != OPERATOR || (token.Value != "+" && token.Value != "-") {
			break
		}
		
		p.pos++ // 消费运算符
		
		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}
		
		if token.Value == "+" {
			left += right
		} else {
			left -= right
		}
	}
	
	return left, nil
}

// parseTerm 解析项（处理 * 和 / 运算符）
func (p *ExpressionParser) parseTerm() (float64, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}
	
	for p.pos < len(p.tokens) {
		token := p.tokens[p.pos]
		
		// 只处理 * 和 / 运算符
		if token.Type != OPERATOR || (token.Value != "*" && token.Value != "/") {
			break
		}
		
		p.pos++ // 消费运算符
		
		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		
		if token.Value == "*" {
			left *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("除数不能为零")
			}
			left /= right
		}
	}
	
	return left, nil
}

// parseFactor 解析因子（处理数字和括号表达式）
func (p *ExpressionParser) parseFactor() (float64, error) {
	if p.pos >= len(p.tokens) {
		return 0, fmt.Errorf("表达式不完整")
	}
	
	token := p.tokens[p.pos]
	
	switch token.Type {
	case NUMBER:
		p.pos++
		value, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return 0, fmt.Errorf("无效的数字: %s", token.Value)
		}
		return value, nil
		
	case LEFT_PAREN:
		p.pos++ // 消费左括号
		result, err := p.parseExpression()
		if err != nil {
			return 0, err
		}
		
		// 期望右括号
		if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != RIGHT_PAREN {
			return 0, fmt.Errorf("缺少右括号")
		}
		p.pos++ // 消费右括号
		
		return result, nil
		
	case OPERATOR:
		// 处理一元负号
		if token.Value == "-" {
			p.pos++
			factor, err := p.parseFactor()
			if err != nil {
				return 0, err
			}
			return -factor, nil
		}
		// 处理一元正号
		if token.Value == "+" {
			p.pos++
			return p.parseFactor()
		}
		return 0, fmt.Errorf("意外的运算符: %s", token.Value)
		
	default:
		return 0, fmt.Errorf("意外的token: %s", token.Value)
	}
}

// HistoryManager 历史记录管理器
type HistoryManager struct {
	entries []HistoryEntry
	maxSize int
}

// NewHistoryManager 创建新的历史管理器
func NewHistoryManager(maxSize int) *HistoryManager {
	return &HistoryManager{
		entries: make([]HistoryEntry, 0),
		maxSize: maxSize,
	}
}

// Add 添加历史记录
func (h *HistoryManager) Add(expression string, result float64) {
	entry := HistoryEntry{
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now(),
	}
	
	h.entries = append(h.entries, entry)
	
	// 保持最大数量限制
	if len(h.entries) > h.maxSize {
		h.entries = h.entries[1:]
	}
}

// GetAll 获取所有历史记录
func (h *HistoryManager) GetAll() []HistoryEntry {
	// 返回副本以防止外部修改
	result := make([]HistoryEntry, len(h.entries))
	copy(result, h.entries)
	return result
}

// Clear 清空历史记录
func (h *HistoryManager) Clear() {
	h.entries = h.entries[:0]
}

// Calculator 计算器主引擎
type Calculator struct {
	parser  *ExpressionParser
	history *HistoryManager
}

// NewCalculator 创建新的计算器
func NewCalculator() *Calculator {
	return &Calculator{
		parser:  NewExpressionParser(),
		history: NewHistoryManager(50), // 最多保存50条历史
	}
}

// Calculate 执行计算
func (c *Calculator) Calculate(expression string) (float64, error) {
	result, err := c.parser.Parse(expression)
	if err != nil {
		return 0, err
	}
	
	// 保存到历史记录
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

// ===== 用户界面 =====

// UserInterface 用户界面管理器
type UserInterface struct {
	calculator *Calculator
	scanner    *bufio.Scanner
}

// NewUserInterface 创建新的用户界面
func NewUserInterface() *UserInterface {
	return &UserInterface{
		calculator: NewCalculator(),
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

// Start 启动用户界面
func (ui *UserInterface) Start() {
	ui.displayWelcome()
	
	for {
		fmt.Print("> ")
		
		if !ui.scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(ui.scanner.Text())
		
		if input == "" {
			continue
		}
		
		command := ui.parseCommand(input)
		
		switch command.Type {
		case Calculate:
			ui.handleCalculate(command.Expression)
		case ShowHistory:
			ui.handleShowHistory()
		case ShowHelp:
			ui.handleShowHelp()
		case Clear:
			ui.handleClear()
		case Quit:
			ui.handleQuit()
			return
		}
	}
}

// displayWelcome 显示欢迎信息
func (ui *UserInterface) displayWelcome() {
	fmt.Println("===== Go计算器 v1.0 =====")
	fmt.Println("支持基础运算: +, -, *, /")
	fmt.Println("支持括号: (), 支持小数")
	fmt.Println("输入 'help' 查看帮助, 'quit' 退出")
	fmt.Println("========================")
}

// parseCommand 解析用户命令
func (ui *UserInterface) parseCommand(input string) Command {
	lower := strings.ToLower(input)
	
	switch {
	case lower == "history" || lower == "h":
		return Command{Type: ShowHistory}
	case lower == "help" || lower == "?":
		return Command{Type: ShowHelp}
	case lower == "clear" || lower == "c":
		return Command{Type: Clear}
	case lower == "quit" || lower == "exit" || lower == "q":
		return Command{Type: Quit}
	default:
		return Command{Type: Calculate, Expression: input}
	}
}

// handleCalculate 处理计算命令
func (ui *UserInterface) handleCalculate(expression string) {
	result, err := ui.calculator.Calculate(expression)
	if err != nil {
		fmt.Printf("❌ 错误: %v\n", err)
		return
	}
	
	// 格式化结果显示
	if result == float64(int64(result)) {
		fmt.Printf("✅ 结果: %.0f\n", result)
	} else {
		fmt.Printf("✅ 结果: %.6g\n", result)
	}
}

// handleShowHistory 处理显示历史命令
func (ui *UserInterface) handleShowHistory() {
	history := ui.calculator.GetHistory()
	
	if len(history) == 0 {
		fmt.Println("📝 暂无计算历史")
		return
	}
	
	fmt.Println("📚 计算历史:")
	for i, entry := range history {
		if entry.Result == float64(int64(entry.Result)) {
			fmt.Printf("%3d. %s = %.0f  [%s]\n", 
				i+1, entry.Expression, entry.Result, 
				entry.Timestamp.Format("15:04:05"))
		} else {
			fmt.Printf("%3d. %s = %.6g  [%s]\n", 
				i+1, entry.Expression, entry.Result, 
				entry.Timestamp.Format("15:04:05"))
		}
	}
}

// handleShowHelp 处理显示帮助命令
func (ui *UserInterface) handleShowHelp() {
	fmt.Println("📖 Go计算器使用帮助:")
	fmt.Println()
	fmt.Println("🧮 支持的运算:")
	fmt.Println("  +  加法    示例: 2 + 3")
	fmt.Println("  -  减法    示例: 5 - 2")
	fmt.Println("  *  乘法    示例: 3 * 4")
	fmt.Println("  /  除法    示例: 8 / 2")
	fmt.Println("  () 括号    示例: (2 + 3) * 4")
	fmt.Println()
	fmt.Println("📝 支持的命令:")
	fmt.Println("  history, h    显示计算历史")
	fmt.Println("  clear, c      清空历史记录")
	fmt.Println("  help, ?       显示帮助信息")
	fmt.Println("  quit, q       退出计算器")
	fmt.Println()
	fmt.Println("💡 示例表达式:")
	fmt.Println("  2 + 3 * 4      → 14")
	fmt.Println("  (10 + 5) / 3   → 5")
	fmt.Println("  -5 + 3         → -2")
	fmt.Println("  3.14 * 2       → 6.28")
}

// handleClear 处理清空历史命令
func (ui *UserInterface) handleClear() {
	ui.calculator.ClearHistory()
	fmt.Println("🗑️  历史记录已清空")
}

// handleQuit 处理退出命令
func (ui *UserInterface) handleQuit() {
	fmt.Println("👋 感谢使用Go计算器！")
}

// ===== 主程序入口 =====

func main() {
	ui := NewUserInterface()
	ui.Start()
}
```

## 🔍 核心算法详解

### 1. 递归下降解析器

计算器使用递归下降解析器实现表达式解析，遵循以下语法规则：

```
Expression := Term (('+' | '-') Term)*
Term       := Factor (('*' | '/') Factor)*
Factor     := Number | '(' Expression ')' | ('-' | '+') Factor
```

#### 运算优先级处理

<MermaidDiagram code="graph TD
    A[parseExpression] --> B[parseTerm]
    B --> C[parseFactor]
    
    A --> D[处理 + -]
    B --> E[处理 * /]
    C --> F[处理数字和括号]
    
    style A fill:#ffcdd2
    style B fill:#f8bbd9
    style C fill:#e1bee7" />

### 2. 词法分析过程

```go
// 输入: "2 + 3 * (4 - 1)"
// 输出tokens:
[
    {NUMBER, "2"},
    {OPERATOR, "+"},
    {NUMBER, "3"},
    {OPERATOR, "*"},
    {LEFT_PAREN, "("},
    {NUMBER, "4"},
    {OPERATOR, "-"},
    {NUMBER, "1"},
    {RIGHT_PAREN, ")"}
]
```

### 3. 语法分析过程

以表达式 `2 + 3 * 4` 为例：

```
parseExpression():
  left = parseTerm() → 2
  遇到 '+' 运算符
  right = parseTerm() → 12 (3 * 4)
  返回: 2 + 12 = 14

parseTerm() for "3 * 4":
  left = parseFactor() → 3
  遇到 '*' 运算符
  right = parseFactor() → 4
  返回: 3 * 4 = 12
```

## 🎮 运行示例

### 编译和运行

```bash
# 编译
go build -o calculator main.go

# 运行
./calculator
```

### 使用示例

```
===== Go计算器 v1.0 =====
支持基础运算: +, -, *, /
支持括号: (), 支持小数
输入 'help' 查看帮助, 'quit' 退出
========================

> 2 + 3
✅ 结果: 5

> 2 + 3 * 4
✅ 结果: 14

> (10 + 5) / 3
✅ 结果: 5

> -5 + 3
✅ 结果: -2

> 3.14 * 2
✅ 结果: 6.28

> 10 / 0
❌ 错误: 语法分析错误: 除数不能为零

> history
📚 计算历史:
  1. 2 + 3 = 5  [14:30:15]
  2. 2 + 3 * 4 = 14  [14:30:20]
  3. (10 + 5) / 3 = 5  [14:30:25]
  4. -5 + 3 = -2  [14:30:30]
  5. 3.14 * 2 = 6.28  [14:30:35]

> help
📖 Go计算器使用帮助:

🧮 支持的运算:
  +  加法    示例: 2 + 3
  -  减法    示例: 5 - 2
  *  乘法    示例: 3 * 4
  /  除法    示例: 8 / 2
  () 括号    示例: (2 + 3) * 4

📝 支持的命令:
  history, h    显示计算历史
  clear, c      清空历史记录
  help, ?       显示帮助信息
  quit, q       退出计算器

💡 示例表达式:
  2 + 3 * 4      → 14
  (10 + 5) / 3   → 5
  -5 + 3         → -2
  3.14 * 2       → 6.28

> quit
👋 感谢使用Go计算器！
```

## 🧪 测试用例

### 创建测试文件 main_test.go

```go
package main

import (
	"testing"
)

func TestExpressionParser(t *testing.T) {
	parser := NewExpressionParser()
	
	testCases := []struct {
		name     string
		input    string
		expected float64
		hasError bool
	}{
		// 基础运算测试
		{"简单加法", "2 + 3", 5, false},
		{"简单减法", "5 - 2", 3, false},
		{"简单乘法", "3 * 4", 12, false},
		{"简单除法", "8 / 2", 4, false},
		
		// 运算优先级测试
		{"乘法优先级", "2 + 3 * 4", 14, false},
		{"除法优先级", "10 - 8 / 2", 6, false},
		{"复合运算", "2 * 3 + 4 * 5", 26, false},
		
		// 括号测试
		{"简单括号", "(2 + 3) * 4", 20, false},
		{"嵌套括号", "((2 + 3) * 4) / 5", 4, false},
		{"复杂括号", "(10 + 5) / (3 * 5)", 1, false},
		
		// 负数测试
		{"负数", "-5", -5, false},
		{"负数运算", "-5 + 3", -2, false},
		{"括号负数", "-(5 + 3)", -8, false},
		
		// 小数测试
		{"小数加法", "3.5 + 2.5", 6, false},
		{"小数乘法", "3.14 * 2", 6.28, false},
		
		// 错误情况测试
		{"除零错误", "10 / 0", 0, true},
		{"语法错误", "2 +", 0, true},
		{"括号不匹配", "(2 + 3", 0, true},
		{"非法字符", "2 + a", 0, true},
		{"空表达式", "", 0, true},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.Parse(tc.input)
			
			if tc.hasError {
				if err == nil {
					t.Errorf("期望出现错误，但没有错误发生")
				}
			} else {
				if err != nil {
					t.Errorf("不期望错误，但发生了错误: %v", err)
				} else if result != tc.expected {
					t.Errorf("期望结果 %v，但得到 %v", tc.expected, result)
				}
			}
		})
	}
}

func TestCalculator(t *testing.T) {
	calc := NewCalculator()
	
	// 测试基础计算
	result, err := calc.Calculate("2 + 3 * 4")
	if err != nil {
		t.Errorf("计算错误: %v", err)
	}
	if result != 14 {
		t.Errorf("期望结果 14，但得到 %v", result)
	}
	
	// 测试历史记录
	history := calc.GetHistory()
	if len(history) != 1 {
		t.Errorf("期望历史记录数量 1，但得到 %d", len(history))
	}
	if history[0].Expression != "2 + 3 * 4" {
		t.Errorf("期望表达式 '2 + 3 * 4'，但得到 '%s'", history[0].Expression)
	}
	if history[0].Result != 14 {
		t.Errorf("期望结果 14，但得到 %v", history[0].Result)
	}
	
	// 测试清空历史
	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("期望历史记录为空，但得到 %d 条记录", len(history))
	}
}

func TestHistoryManager(t *testing.T) {
	hm := NewHistoryManager(3) // 最大容量3
	
	// 添加记录
	hm.Add("1 + 1", 2)
	hm.Add("2 + 2", 4)
	hm.Add("3 + 3", 6)
	
	history := hm.GetAll()
	if len(history) != 3 {
		t.Errorf("期望历史记录数量 3，但得到 %d", len(history))
	}
	
	// 测试容量限制
	hm.Add("4 + 4", 8)
	history = hm.GetAll()
	if len(history) != 3 {
		t.Errorf("期望历史记录数量仍为 3，但得到 %d", len(history))
	}
	
	// 验证最旧的记录被移除
	if history[0].Expression != "2 + 2" {
		t.Errorf("期望最旧记录被移除，但第一个记录是 '%s'", history[0].Expression)
	}
}

// 性能测试
func BenchmarkCalculation(b *testing.B) {
	calc := NewCalculator()
	expression := "1 + 2 * 3 - 4 / 5 + (6 - 7) * 8"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Calculate(expression)
	}
}

func BenchmarkTokenization(b *testing.B) {
	parser := NewExpressionParser()
	expression := "1 + 2 * 3 - 4 / 5 + (6 - 7) * 8"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.tokenize(expression)
	}
}
```

### 运行测试

```bash
# 运行所有测试
go test -v

# 运行性能测试
go test -bench=.

# 运行测试并显示覆盖率
go test -cover
```

## 📊 性能分析

### 测试结果示例

```
=== RUN   TestExpressionParser
=== RUN   TestExpressionParser/简单加法
=== RUN   TestExpressionParser/简单减法
=== RUN   TestExpressionParser/简单乘法
=== RUN   TestExpressionParser/简单除法
=== RUN   TestExpressionParser/乘法优先级
=== RUN   TestExpressionParser/除法优先级
=== RUN   TestExpressionParser/复合运算
=== RUN   TestExpressionParser/简单括号
=== RUN   TestExpressionParser/嵌套括号
=== RUN   TestExpressionParser/复杂括号
=== RUN   TestExpressionParser/负数
=== RUN   TestExpressionParser/负数运算
=== RUN   TestExpressionParser/括号负数
=== RUN   TestExpressionParser/小数加法
=== RUN   TestExpressionParser/小数乘法
=== RUN   TestExpressionParser/除零错误
=== RUN   TestExpressionParser/语法错误
=== RUN   TestExpressionParser/括号不匹配
=== RUN   TestExpressionParser/非法字符
=== RUN   TestExpressionParser/空表达式
--- PASS: TestExpressionParser (0.00s)

=== RUN   TestCalculator
--- PASS: TestCalculator (0.00s)

=== RUN   TestHistoryManager
--- PASS: TestHistoryManager (0.00s)

BenchmarkCalculation-8         	 1000000	      1203 ns/op
BenchmarkTokenization-8       	 2000000	       856 ns/op

PASS
coverage: 92.3% of statements
ok  	calculator	1.234s
```

## 🚀 扩展功能

### 1. 添加科学计算功能

```go
// 扩展Token类型
const (
	NUMBER TokenType = iota
	OPERATOR
	LEFT_PAREN
	RIGHT_PAREN
	FUNCTION  // 新增：函数类型
)

// 添加数学函数
var mathFunctions = map[string]func(float64) float64{
	"sin":  math.Sin,
	"cos":  math.Cos,
	"tan":  math.Tan,
	"log":  math.Log,
	"sqrt": math.Sqrt,
}
```

### 2. 添加变量存储功能

```go
type VariableManager struct {
	variables map[string]float64
}

func (vm *VariableManager) Set(name string, value float64) {
	vm.variables[name] = value
}

func (vm *VariableManager) Get(name string) (float64, bool) {
	value, exists := vm.variables[name]
	return value, exists
}
```

### 3. 添加表达式保存功能

```go
type ExpressionManager struct {
	expressions map[string]string
}

func (em *ExpressionManager) Save(name, expression string) {
	em.expressions[name] = expression
}

func (em *ExpressionManager) Load(name string) (string, bool) {
	expression, exists := em.expressions[name]
	return expression, exists
}
```

## 📝 学习要点总结

通过实现这个计算器项目，您学到了：

### Go语言特性
1. **结构体和方法**：面向对象编程基础
2. **接口设计**：抽象和多态概念
3. **错误处理**：Go语言的错误处理模式
4. **包管理**：代码组织和模块化
5. **测试驱动**：单元测试和基准测试

### 算法和数据结构
1. **递归下降解析器**：编译器前端技术
2. **栈结构应用**：表达式求值算法
3. **词法分析**：字符串处理技巧
4. **语法分析**：递归算法应用

### 软件工程实践
1. **分层架构**：代码结构设计
2. **单一职责**：模块化编程思想
3. **错误处理**：健壮性编程
4. **测试覆盖**：质量保证方法

---

## 结语

计算器项目展示了Go语言在系统编程方面的优势，通过完整的项目实现，您掌握了从产品设计到代码实现的全流程开发技能。

<div style="text-align: center; margin-top: 2rem;">
  <a href="./architecture.html" style="display: inline-block; padding: 8px 16px; background: #6c757d; color: white; text-decoration: none; border-radius: 4px; margin: 0 8px;">← 架构设计</a>
  <a href="../todo-cli/" style="display: inline-block; padding: 8px 16px; background: #00ADD8; color: white; text-decoration: none; border-radius: 4px; margin: 0 8px;">下个项目：TODO CLI →</a>
</div> 