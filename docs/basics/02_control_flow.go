package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
02_control_flow.go - Go语言基础：控制流
学习内容：
1. if/else 条件语句
2. for 循环语句
3. switch 选择语句
4. defer 延迟执行
5. range 循环
*/

func main() {
	fmt.Println("=== Go语言基础：控制流 ===")

	// 1. if/else 条件语句
	fmt.Println("\n1. if/else 条件语句：")

	age := 18

	// 基本if语句
	if age >= 18 {
		fmt.Println("已成年")
	}

	// if-else语句
	if age < 13 {
		fmt.Println("儿童")
	} else if age < 18 {
		fmt.Println("青少年")
	} else {
		fmt.Println("成年人")
	}

	// if语句中可以包含初始化语句
	if score := 85; score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 2. for 循环语句
	fmt.Println("\n2. for 循环语句：")

	// 基本for循环（类似C语言）
	fmt.Print("基本for循环：")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while风格的for循环
	fmt.Print("while风格for循环：")
	j := 1
	for j <= 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	// 无限循环(带break)
	fmt.Print("无限循环带break：")
	k := 1
	for {
		if k > 3 {
			break
		}
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()

	// 使用continue跳过当前迭代
	fmt.Print("跳过偶数：")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 嵌套循环和标签
	fmt.Println("嵌套循环和标签：")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("跳出外层循环")
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// 3. switch 选择语句
	fmt.Println("\n3. switch 选择语句：")

	// 基本switch语句
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("无效的日期")
	}

	// switch带初始化语句
	switch grade := 85 / 10; grade {
	case 10, 9:
		fmt.Println("等级：A")
	case 8:
		fmt.Println("等级：B")
	case 7, 6:
		fmt.Println("等级：C")
	default:
		fmt.Println("等级：D")
	}

	// 无表达式的switch（相当于if-else链）
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("结冰温度")
	case temperature < 10:
		fmt.Println("很冷")
	case temperature < 20:
		fmt.Println("凉爽")
	case temperature < 30:
		fmt.Println("温暖")
	default:
		fmt.Println("炎热")
	}

	// switch中的fallthrough
	fmt.Println("fallthrough示例：")
	value := 2
	switch value {
	case 1:
		fmt.Println("一")
		fallthrough
	case 2:
		fmt.Println("二")
		fallthrough
	case 3:
		fmt.Println("三")
	default:
		fmt.Println("其他")
	}

	// 4. defer 延迟执行
	fmt.Println("\n4. defer 延迟执行：")

	// defer会在函数返回前执行，执行顺序是栈（LIFO）
	defer fmt.Println("这是defer语句1 (最后执行)")
	defer fmt.Println("这是defer语句2 (倒数第二)")
	defer fmt.Println("这是defer语句3 (倒数第三)")

	fmt.Println("正常执行的语句")

	// defer用于资源清理
	deferExample()

	// 5. range 循环
	fmt.Println("\n5. range 循环：")

	// 遍历字符串
	str := "Hello"
	fmt.Println("遍历字符串：")
	for i, char := range str {
		fmt.Printf("索引: %d, 字符: %c\n", i, char)
	}

	// 只要索引或只要值
	fmt.Print("只要索引：")
	for i := range str {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Print("只要字符：")
	for _, char := range str {
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	// 遍历数组/切片
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("遍历切片：")
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}

	// 遍历映射
	scores := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 78,
	}
	fmt.Println("遍历映射：")
	for name, score := range scores {
		fmt.Printf("%s: %d分\n", name, score)
	}

	// 6. 随机数示例
	fmt.Println("\n6. 随机数示例：")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("掷骰子5次：")
	for i := 0; i < 5; i++ {
		dice := rand.Intn(6) + 1
		fmt.Printf("第%d次: %d\n", i+1, dice)
	}
}

// defer示例函数
func deferExample() {
	fmt.Println("函数开始")
	defer fmt.Println("清理资源 (defer)")
	defer fmt.Println("关闭文件 (defer)")
	fmt.Println("函数主体")
	fmt.Println("函数即将结束")
}
