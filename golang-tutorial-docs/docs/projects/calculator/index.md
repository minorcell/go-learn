# 📱 计算器项目

一个功能完整的命令行计算器应用，支持基础数学运算、表达式解析和历史记录。

## 🎯 项目概述

这是一个使用Go语言开发的命令行计算器，展示了Go语言在处理字符串解析、数据结构和用户交互方面的能力。

### 核心功能

- ✅ **基础运算**：加、减、乘、除运算
- ✅ **表达式解析**：支持复杂数学表达式
- ✅ **括号支持**：处理运算优先级
- ✅ **历史记录**：保存计算历史
- ✅ **错误处理**：友好的错误提示
- ✅ **交互式界面**：命令行交互模式

### 技术栈

- **语言**：Go 1.19+
- **核心包**：`strconv`、`strings`、`bufio`、`fmt`
- **数据结构**：切片、映射、结构体
- **算法**：表达式解析、栈结构

## 📚 文档导航

| 阶段 | 文档 | 内容 |
|------|------|------|
| 📋 规划阶段 | [产品设计](./product-design.md) | 需求分析、功能规划、用户体验设计 |
| 🏗️ 设计阶段 | [架构设计](./architecture.md) | 系统架构、模块设计、数据流程 |
| 💻 开发阶段 | [代码实现](./implementation.md) | 详细实现、代码解析、测试用例 |

## 🚀 快速开始

```bash
# 克隆项目
git clone https://github.com/minorcell/go-learn.git
cd go-learn/projects/calculator

# 运行计算器
go run main.go

# 编译可执行文件
go build -o calculator main.go
./calculator
```

## 💡 学习目标

通过这个项目，您将学会：

1. **字符串处理**：解析和处理用户输入
2. **数据结构**：栈的实现和应用
3. **错误处理**：优雅的错误处理机制
4. **用户交互**：命令行界面设计
5. **代码组织**：模块化设计思想
6. **测试驱动**：单元测试和集成测试

## 🎮 使用示例

```
$ go run main.go
=== Go计算器 v1.0 ===
输入数学表达式 (输入 'quit' 退出):

> 2 + 3 * 4
结果: 14

> (10 + 5) / 3
结果: 5

> history
计算历史:
1. 2 + 3 * 4 = 14
2. (10 + 5) / 3 = 5

> quit
感谢使用Go计算器!
```

---

<div style="text-align: center; margin-top: 2rem;">
  <h3>🎯 开始学习计算器项目</h3>
  <p>按照文档顺序学习，从产品设计到代码实现</p>
  <a href="./product-design.html" style="display: inline-block; padding: 8px 16px; background: #00ADD8; color: white; text-decoration: none; border-radius: 4px; margin: 0 8px;">产品设计 →</a>
</div> 