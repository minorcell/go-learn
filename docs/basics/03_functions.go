package main

import (
	"errors"
	"fmt"
)

/*
03_functions.go - Go语言基础：函数
学习内容：
1. 函数定义和调用
2. 参数和返回值
3. 可变参数
4. 函数作为值
5. 闭包
6. 递归
7. 匿名函数
*/

func main() {
	fmt.Println("=== Go语言基础：函数 ===")

	// 1. 基本函数调用
	fmt.Println("\n1. 基本函数调用：")
	greet("Go语言")

	// 2. 带返回值的函数
	fmt.Println("\n2. 带返回值的函数：")
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// 3. 多返回值函数
	fmt.Println("\n3. 多返回值函数：")
	quotient, remainder := divideInt(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 命名返回值
	area, perimeter := rectangleInfo(5, 3)
	fmt.Printf("矩形面积: %d, 周长: %d\n", area, perimeter)

	// 4. 错误处理
	fmt.Println("\n4. 错误处理：")
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %s\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %s\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}

	// 5. 可变参数
	fmt.Println("\n5. 可变参数：")
	total1 := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total1)

	numbers := []int{10, 20, 30}
	total2 := sumAll(numbers...) // 展开切片
	fmt.Printf("10+20+30 = %d\n", total2)

	// 6. 函数作为值
	fmt.Println("\n6. 函数作为值：")
	var operation func(int, int) int
	operation = add
	fmt.Printf("使用函数变量: %d\n", operation(5, 3))

	// 函数作为参数
	fmt.Println("函数作为参数：")
	applyOperation(10, 5, add, "加法")
	applyOperation(10, 5, multiply, "乘法")

	// 7. 匿名函数
	fmt.Println("\n7. 匿名函数：")
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("5的平方: %d\n", square(5))

	// 立即执行匿名函数
	avgResult := func(a, b float64) float64 {
		return (a + b) / 2
	}(10, 20)
	fmt.Printf("10和20的平均值: %.1f\n", avgResult)

	// 8. 闭包
	fmt.Println("\n8. 闭包：")
	counter := createCounter()
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())

	// 另一个闭包示例
	multiplier := createMultiplier(3)
	fmt.Printf("3 × 4 = %d\n", multiplier(4))
	fmt.Printf("3 × 7 = %d\n", multiplier(7))

	// 9. 递归
	fmt.Println("\n9. 递归：")
	fmt.Printf("5的阶乘: %d\n", factorial(5))
	fmt.Printf("斐波那契数列第8项: %d\n", fibonacci(8))

	// 10. 高阶函数
	fmt.Println("\n10. 高阶函数：")
	nums := []int{1, 2, 3, 4, 5}

	// 映射函数
	squares := mapFunc(nums, func(x int) int { return x * x })
	fmt.Printf("平方: %v\n", squares)

	// 过滤函数
	evens := filterFunc(nums, func(x int) bool { return x%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)

	// 归约函数
	sumResult := reduceFunc(nums, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("求和: %d\n", sumResult)
}

// 基本函数，无返回值
func greet(name string) {
	fmt.Printf("你好, %s!\n", name)
}

// 带返回值的函数
func add(a, b int) int {
	return a + b
}

// 乘法函数
func multiply(a, b int) int {
	return a * b
}

// 多返回值函数
func divideInt(a, b int) (int, int) {
	return a / b, a % b
}

// 命名返回值函数
func rectangleInfo(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回
}

// 错误处理函数
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 可变参数函数
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 接受函数作为参数
func applyOperation(a, b int, op func(int, int) int, opName string) {
	result := op(a, b)
	fmt.Printf("%s: %d\n", opName, result)
}

// 返回闭包的函数
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 带参数的闭包工厂
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// 递归函数 - 阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 递归函数 - 斐波那契
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 映射函数（类似map）
func mapFunc(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 过滤函数（类似filter）
func filterFunc(slice []int, fn func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// 归约函数（类似reduce）
func reduceFunc(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}
