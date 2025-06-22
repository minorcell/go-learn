# 架构设计 - 计算器项目

本文档详细介绍计算器项目的系统架构、模块设计和数据流程。

## 架构概述

计算器采用分层架构设计，将用户界面、业务逻辑和数据处理分离，确保代码的可维护性和可扩展性。

### 设计原则

- **单一职责**：每个模块只负责一个特定功能
- **开闭原则**：对扩展开放，对修改封闭
- **依赖倒置**：高层模块不依赖低层模块
- **接口隔离**：使用小而专一的接口

## 模块设计

### 1. 用户界面层 (UI Layer)

#### UserInterface 用户界面
```go
type UserInterface struct {
    calculator *Calculator
    scanner    *bufio.Scanner
}

// 主要方法
func (ui *UserInterface) Start()
func (ui *UserInterface) displayWelcome()
func (ui *UserInterface) readInput() string
func (ui *UserInterface) displayResult(result float64)
func (ui *UserInterface) displayError(err error)
```

**职责**：
- 处理用户输入输出
- 显示计算结果和错误信息
- 管理程序生命周期

#### CommandParser 命令解析器
```go
type CommandType int

const (
    Calculate CommandType = iota
    ShowHistory
    ShowHelp
    Quit
)

type Command struct {
    Type CommandType
    Expression string
}

func ParseCommand(input string) Command
```

**职责**：
- 解析用户输入的命令类型
- 区分计算表达式和系统命令
- 验证输入格式

### 2. 业务逻辑层 (Business Logic Layer)

#### Calculator 计算引擎
```go
type Calculator struct {
    parser  *ExpressionParser
    history *HistoryManager
}

func (c *Calculator) Calculate(expression string) (float64, error)
func (c *Calculator) GetHistory() []HistoryEntry
func (c *Calculator) ClearHistory()
```

**职责**：
- 协调各个组件完成计算
- 管理计算流程
- 提供对外接口

#### ExpressionParser 表达式解析器
```go
type Token struct {
    Type  TokenType
    Value string
}

type TokenType int

const (
    NUMBER TokenType = iota
    OPERATOR
    LEFT_PAREN
    RIGHT_PAREN
)

type ExpressionParser struct {
    tokens []Token
    pos    int
}

func (p *ExpressionParser) Parse(expression string) (float64, error)
func (p *ExpressionParser) tokenize(expression string) []Token
func (p *ExpressionParser) parseExpression() (float64, error)
func (p *ExpressionParser) parseTerm() (float64, error)
func (p *ExpressionParser) parseFactor() (float64, error)
```

**职责**：
- 词法分析：将表达式分解为token
- 语法分析：按照运算规则解析表达式
- 计算求值：递归计算表达式结果

#### HistoryManager 历史管理器
```go
type HistoryEntry struct {
    Expression string
    Result     float64
    Timestamp  time.Time
}

type HistoryManager struct {
    entries []HistoryEntry
    maxSize int
}

func (h *HistoryManager) Add(expression string, result float64)
func (h *HistoryManager) GetAll() []HistoryEntry
func (h *HistoryManager) Clear()
func (h *HistoryManager) GetLast(n int) []HistoryEntry
```

**职责**：
- 存储计算历史记录
- 管理历史记录容量
- 提供历史查询功能

### 3. 数据层 (Data Layer)

#### 内存存储
- **计算上下文**：当前表达式解析状态
- **运算符栈**：用于处理运算优先级
- **操作数栈**：存储中间计算结果

#### 历史记录存储
- **历史列表**：存储最近的计算记录
- **时间戳**：记录每次计算的时间
- **容量管理**：限制历史记录数量

## 数据流程

### 计算流程

### 表达式解析算法

采用递归下降解析器 (Recursive Descent Parser) 实现：

#### 语法规则 (BNF)
```
Expression := Term (('+' | '-') Term)*
Term       := Factor (('*' | '/') Factor)*
Factor     := Number | '(' Expression ')'
Number     := [0-9]+ ('.' [0-9]+)?
```

## 核心算法

### 1. 词法分析 (Tokenization)

```go
func (p *ExpressionParser) tokenize(expression string) []Token {
    var tokens []Token
    var current strings.Builder
    
    for i, char := range expression {
        switch {
        case unicode.IsDigit(char) || char == '.':
            current.WriteRune(char)
        case char == '+' || char == '-' || char == '*' || char == '/':
            if current.Len() > 0 {
                tokens = append(tokens, Token{NUMBER, current.String()})
                current.Reset()
            }
            tokens = append(tokens, Token{OPERATOR, string(char)})
        case char == '(':
            tokens = append(tokens, Token{LEFT_PAREN, "("})
        case char == ')':
            tokens = append(tokens, Token{RIGHT_PAREN, ")"})
        case char == ' ':
            // 忽略空格
        default:
            return nil // 非法字符
        }
    }
    
    if current.Len() > 0 {
        tokens = append(tokens, Token{NUMBER, current.String()})
    }
    
    return tokens
}
```

### 2. 运算优先级处理

```go
func (p *ExpressionParser) parseExpression() (float64, error) {
    left, err := p.parseTerm()
    if err != nil {
        return 0, err
    }
    
    for p.pos < len(p.tokens) {
        token := p.tokens[p.pos]
        if token.Type != OPERATOR || (token.Value != "+" && token.Value != "-") {
            break
        }
        
        p.pos++
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
```

### 3. 错误处理策略

## 测试策略

### 单元测试模块

#### 1. ExpressionParser 测试
```go
func TestExpressionParser(t *testing.T) {
    testCases := []struct {
        input    string
        expected float64
        hasError bool
    }{
        {"2 + 3", 5, false},
        {"2 * 3 + 4", 10, false},
        {"(2 + 3) * 4", 20, false},
        {"10 / 0", 0, true},
        {"2 +", 0, true},
    }
    
    parser := NewExpressionParser()
    for _, tc := range testCases {
        result, err := parser.Parse(tc.input)
        // 断言逻辑...
    }
}
```

#### 2. Calculator 集成测试
```go
func TestCalculatorIntegration(t *testing.T) {
    calc := NewCalculator()
    
    // 测试基础计算
    result, err := calc.Calculate("2 + 3 * 4")
    assert.NoError(t, err)
    assert.Equal(t, 14.0, result)
    
    // 测试历史记录
    history := calc.GetHistory()
    assert.Len(t, history, 1)
    assert.Equal(t, "2 + 3 * 4", history[0].Expression)
}
```

### 性能测试

#### 1. 计算性能
```go
func BenchmarkCalculation(b *testing.B) {
    calc := NewCalculator()
    expression := "1 + 2 * 3 - 4 / 5"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        calc.Calculate(expression)
    }
}
```

#### 2. 内存使用
```go
func TestMemoryUsage(t *testing.T) {
    calc := NewCalculator()
    
    // 执行大量计算
    for i := 0; i < 10000; i++ {
        calc.Calculate(fmt.Sprintf("%d + %d", i, i+1))
    }
    
    // 检查内存使用是否在预期范围内
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    assert.Less(t, m.Alloc, uint64(10*1024*1024)) // 小于10MB
}
```

## 扩展性设计

### 1. 插件系统
```go
type Function interface {
    Name() string
    Execute(args []float64) (float64, error)
    ArgCount() int
}

type FunctionRegistry struct {
    functions map[string]Function
}

func (r *FunctionRegistry) Register(fn Function)
func (r *FunctionRegistry) Call(name string, args []float64) (float64, error)
```

### 2. 配置系统
```go
type Config struct {
    Precision    int     `json:"precision"`
    MaxHistory   int     `json:"max_history"`
    AngleUnit    string  `json:"angle_unit"` // "degree" | "radian"
    DecimalSep   string  `json:"decimal_separator"`
}

func LoadConfig(filename string) (*Config, error)
func (c *Config) Save(filename string) error
```

## 性能优化

### 1. 内存优化
- **对象池**：重用Token对象减少GC压力
- **字符串缓存**：缓存常用的数学常量
- **历史容量限制**：防止内存无限增长

### 2. 计算优化
- **表达式缓存**：缓存复杂表达式的解析结果
- **预编译**：对常用表达式进行预编译
- **并发计算**：对独立子表达式并行计算