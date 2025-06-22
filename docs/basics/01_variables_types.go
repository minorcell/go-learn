package main

import "fmt"

/*
01_variables_types.go - Go语言基础：变量和数据类型
学习内容：
1. 包声明和导入
2. 变量声明和初始化
3. 常量
4. 基本数据类型
5. 类型转换
*/

func main() {
	fmt.Println("=== Go语言基础：变量和数据类型 ===")

	// 1. 变量声明的几种方式
	fmt.Println("\n1. 变量声明：")

	// 方式1：var关键字声明
	var name string
	name = "Go语言"
	fmt.Printf("name: %s\n", name)

	// 方式2：声明并初始化
	var age int = 25
	fmt.Printf("age: %d\n", age)

	// 方式3：类型推断
	var city = "北京"
	fmt.Printf("city: %s\n", city)

	// 方式4：短变量声明（只能在函数内使用）
	country := "中国"
	fmt.Printf("country: %s\n", country)

	// 方式5：多变量声明
	var x, y int = 10, 20
	fmt.Printf("x: %d, y: %d\n", x, y)

	a, b, c := 1, 2.5, "hello"
	fmt.Printf("a: %d, b: %.1f, c: %s\n", a, b, c)

	// 2. 常量
	fmt.Println("\n2. 常量：")
	const PI = 3.14159
	const greeting = "Hello, World!"
	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("greeting: %s\n", greeting)

	// 常量组
	const (
		StatusOK            = 200
		StatusNotFound      = 404
		StatusInternalError = 500
	)
	fmt.Printf("HTTP状态码 - OK: %d, NotFound: %d\n", StatusOK, StatusNotFound)

	// 3. 基本数据类型
	fmt.Println("\n3. 基本数据类型：")

	// 整数类型
	var int8Val int8 = 127     // -128 到 127
	var int16Val int16 = 32767 // -32768 到 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807
	var uintVal uint = 42 // 无符号整数

	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d, uint: %d\n",
		int8Val, int16Val, int32Val, int64Val, uintVal)

	// 浮点类型
	var float32Val float32 = 3.14
	var float64Val float64 = 3.141592653589793
	fmt.Printf("float32: %.2f, float64: %.15f\n", float32Val, float64Val)

	// 布尔类型
	var boolVal bool = true
	var isLearning = false
	fmt.Printf("bool值: %t, 正在学习: %t\n", boolVal, isLearning)

	// 字符串类型
	var str1 string = "这是一个字符串"
	str2 := `这是一个
多行字符串`
	fmt.Printf("字符串1: %s\n", str1)
	fmt.Printf("多行字符串: %s\n", str2)

	// byte和rune类型
	var byteVal byte = 'A' // byte是uint8的别名
	var runeVal rune = '中' // rune是int32的别名，表示Unicode字符
	fmt.Printf("byte: %c (%d), rune: %c (%d)\n", byteVal, byteVal, runeVal, runeVal)

	// 4. 零值（默认值）
	fmt.Println("\n4. 零值：")
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	fmt.Printf("int零值: %d, float零值: %.1f, string零值: '%s', bool零值: %t\n",
		zeroInt, zeroFloat, zeroString, zeroBool)

	// 5. 类型转换
	fmt.Println("\n5. 类型转换：")
	var intNum int = 42
	var floatNum float64 = float64(intNum)
	var stringNum string = fmt.Sprintf("%d", intNum)

	fmt.Printf("int转float64: %d -> %.1f\n", intNum, floatNum)
	fmt.Printf("int转string: %d -> %s\n", intNum, stringNum)

	// 字符串长度和字节
	text := "Hello, 世界"
	fmt.Printf("字符串: %s, 字节长度: %d, 字符数量: %d\n",
		text, len(text), len([]rune(text)))

	// 6. 指针基础
	fmt.Println("\n6. 指针基础：")
	var num int = 42
	var ptr *int = &num

	fmt.Printf("num的值: %d\n", num)
	fmt.Printf("num的地址: %p\n", &num)
	fmt.Printf("ptr存储的地址: %p\n", ptr)
	fmt.Printf("ptr指向的值: %d\n", *ptr)

	// 修改指针指向的值
	*ptr = 100
	fmt.Printf("通过指针修改后，num的值: %d\n", num)
}
