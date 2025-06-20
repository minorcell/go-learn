package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
项目案例1：命令行计算器
功能：
1. 基本四则运算
2. 历史记录
3. 连续计算
4. 退出命令
*/

type Calculator struct {
	history []string
}

func main() {
	fmt.Println("=== Go语言项目案例：命令行计算器 ===")
	fmt.Println("支持的操作：+, -, *, /")
	fmt.Println("输入格式：数字1 操作符 数字2")
	fmt.Println("特殊命令：history (查看历史), clear (清除历史), quit (退出)")
	fmt.Println("================================================")

	calc := &Calculator{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("计算器> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		// 处理特殊命令
		switch input {
		case "quit", "exit", "q":
			fmt.Println("再见！")
			return
		case "history", "h":
			calc.showHistory()
			continue
		case "clear", "c":
			calc.clearHistory()
			fmt.Println("历史记录已清除")
			continue
		case "help":
			calc.showHelp()
			continue
		}

		// 处理计算
		result, err := calc.calculate(input)
		if err != nil {
			fmt.Printf("错误: %s\n", err)
		} else {
			fmt.Printf("结果: %.2f\n", result)
			calc.addToHistory(fmt.Sprintf("%s = %.2f", input, result))
		}
	}
}

// 执行计算
func (c *Calculator) calculate(input string) (float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("输入格式错误，请使用：数字1 操作符 数字2")
	}

	// 解析第一个数字
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("无效的第一个数字: %s", parts[0])
	}

	// 获取操作符
	operator := parts[1]

	// 解析第二个数字
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("无效的第二个数字: %s", parts[2])
	}

	// 执行计算
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("不支持的操作符: %s", operator)
	}
}

// 添加到历史记录
func (c *Calculator) addToHistory(record string) {
	c.history = append(c.history, record)
	// 限制历史记录数量
	if len(c.history) > 10 {
		c.history = c.history[1:]
	}
}

// 显示历史记录
func (c *Calculator) showHistory() {
	if len(c.history) == 0 {
		fmt.Println("暂无历史记录")
		return
	}

	fmt.Println("计算历史：")
	for i, record := range c.history {
		fmt.Printf("%d. %s\n", i+1, record)
	}
}

// 清除历史记录
func (c *Calculator) clearHistory() {
	c.history = []string{}
}

// 显示帮助信息
func (c *Calculator) showHelp() {
	fmt.Println("使用说明：")
	fmt.Println("1. 基本计算：输入 '数字1 操作符 数字2'")
	fmt.Println("   例如：10 + 5, 20.5 * 3, 100 / 4")
	fmt.Println("2. 支持的操作符：+ - * /")
	fmt.Println("3. 特殊命令：")
	fmt.Println("   history 或 h  - 查看计算历史")
	fmt.Println("   clear 或 c    - 清除历史记录")
	fmt.Println("   help          - 显示此帮助")
	fmt.Println("   quit 或 exit  - 退出计算器")
}
