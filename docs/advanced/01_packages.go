package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
01_packages.go - Go语言进阶：包管理
学习内容：
1. 包的定义和导入
2. 包的可见性规则
3. init函数
4. 标准库常用包
5. 自定义包
6. 模块管理
*/

func main() {
	fmt.Println("=== Go语言进阶：包管理 ===")

	// 1. 标准库包的使用
	fmt.Println("\n1. 标准库包的使用：")

	// fmt包 - 格式化输入输出
	name := "张三"
	age := 25
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
	fmt.Fprintf(os.Stdout, "使用Fprintf输出: %s\n", "Hello World")

	// math包 - 数学函数
	fmt.Printf("圆周率: %.6f\n", math.Pi)
	fmt.Printf("平方根: %.2f\n", math.Sqrt(16))
	fmt.Printf("最大值: %.0f\n", math.Max(10, 20))
	fmt.Printf("正弦值: %.4f\n", math.Sin(math.Pi/2))

	// strings包 - 字符串操作
	text := "Hello, Go Language!"
	fmt.Printf("原字符串: %s\n", text)
	fmt.Printf("转大写: %s\n", strings.ToUpper(text))
	fmt.Printf("转小写: %s\n", strings.ToLower(text))
	fmt.Printf("包含'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("替换: %s\n", strings.Replace(text, "Go", "Golang", 1))

	words := strings.Split(text, " ")
	fmt.Printf("分割结果: %v\n", words)
	joined := strings.Join(words, "-")
	fmt.Printf("连接结果: %s\n", joined)

	// time包 - 时间处理
	fmt.Println("\n2. 时间处理：")
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("时间戳: %d\n", now.Unix())

	// 时间解析
	timeStr := "2024-01-15 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("时间解析错误: %s\n", err)
	} else {
		fmt.Printf("解析的时间: %s\n", parsedTime.Format("2006年01月02日 15:04:05"))
	}

	// 时间计算
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("明天: %s\n", tomorrow.Format("2006-01-02"))

	nextHour := now.Add(time.Hour)
	fmt.Printf("一小时后: %s\n", nextHour.Format("15:04:05"))

	// 3. 文件路径处理
	fmt.Println("\n3. 文件路径处理：")

	path := "/Users/john/documents/project/main.go"
	fmt.Printf("原路径: %s\n", path)
	fmt.Printf("目录: %s\n", filepath.Dir(path))
	fmt.Printf("文件名: %s\n", filepath.Base(path))
	fmt.Printf("扩展名: %s\n", filepath.Ext(path))

	// 路径连接
	newPath := filepath.Join("/home", "user", "documents", "file.txt")
	fmt.Printf("连接路径: %s\n", newPath)

	// 4. 操作系统接口
	fmt.Println("\n4. 操作系统接口：")

	// 环境变量
	fmt.Printf("HOME目录: %s\n", os.Getenv("HOME"))
	fmt.Printf("当前用户: %s\n", os.Getenv("USER"))

	// 命令行参数
	fmt.Printf("程序名: %s\n", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Printf("参数: %v\n", os.Args[1:])
	} else {
		fmt.Println("没有额外的命令行参数")
	}

	// 5. 自定义类型和方法
	fmt.Println("\n5. 自定义类型和方法：")

	// 自定义类型
	type Temperature float64

	// 为自定义类型定义方法
	const (
		Celsius    Temperature = 0
		Fahrenheit Temperature = 1
	)

	// 方法定义
	celsiusToFahrenheit := func(c Temperature) Temperature {
		return c*9/5 + 32
	}

	fahrenheitToCelsius := func(f Temperature) Temperature {
		return (f - 32) * 5 / 9
	}

	temp := Temperature(25)
	fmt.Printf("摄氏度: %.1f°C\n", temp)
	fmt.Printf("华氏度: %.1f°F\n", celsiusToFahrenheit(temp))

	tempF := Temperature(77)
	fmt.Printf("华氏度: %.1f°F\n", tempF)
	fmt.Printf("摄氏度: %.1f°C\n", fahrenheitToCelsius(tempF))

	// 6. 包的可见性示例
	fmt.Println("\n6. 包的可见性规则：")
	fmt.Println("大写字母开头的标识符是公开的（可导出）")
	fmt.Println("小写字母开头的标识符是私有的（不可导出）")

	// 公开的函数/变量/类型：fmt.Println, math.Pi, time.Now
	// 私有的在其他包中无法访问

	// 7. 常用包功能展示
	fmt.Println("\n7. 常用包功能：")

	// 数学计算
	numbers := []float64{1, 2, 3, 4, 5}
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	fmt.Printf("平均值: %.2f\n", avg)
	fmt.Printf("标准差示例: %.2f\n", math.Sqrt(calculateVariance(numbers, avg)))

	// 字符串处理工具函数
	fmt.Println("\n字符串处理工具：")
	demoStringProcessing()

	// 时间工具函数
	fmt.Println("\n时间工具：")
	demoTimeProcessing()
}

// 计算方差的辅助函数
func calculateVariance(numbers []float64, mean float64) float64 {
	var variance float64
	for _, num := range numbers {
		variance += math.Pow(num-mean, 2)
	}
	return variance / float64(len(numbers))
}

// 字符串处理示例
func demoStringProcessing() {
	text := "  Hello, World! This is Go programming.  "
	fmt.Printf("原文本: '%s'\n", text)
	fmt.Printf("去除空格: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("单词数量: %d\n", len(strings.Fields(text)))
	fmt.Printf("字符长度: %d\n", len(text))
	fmt.Printf("是否以Hello开头: %t\n", strings.HasPrefix(strings.TrimSpace(text), "Hello"))
	fmt.Printf("是否以.结尾: %t\n", strings.HasSuffix(strings.TrimSpace(text), "."))
}

// 时间处理示例
func demoTimeProcessing() {
	now := time.Now()

	// 不同的时间格式
	formats := map[string]string{
		"标准格式":    "2006-01-02 15:04:05",
		"日期格式":    "2006-01-02",
		"时间格式":    "15:04:05",
		"中文格式":    "2006年01月02日",
		"12小时格式":  "2006-01-02 03:04:05 PM",
		"RFC3339": time.RFC3339,
	}

	for name, format := range formats {
		fmt.Printf("%s: %s\n", name, now.Format(format))
	}

	// 时间计算
	fmt.Println("\n时间计算示例：")
	fmt.Printf("一周前: %s\n", now.AddDate(0, 0, -7).Format("2006-01-02"))
	fmt.Printf("一个月后: %s\n", now.AddDate(0, 1, 0).Format("2006-01-02"))
	fmt.Printf("一年后: %s\n", now.AddDate(1, 0, 0).Format("2006-01-02"))
}

// init函数示例
func init() {
	fmt.Println("这是init函数，在main函数之前执行")
}
