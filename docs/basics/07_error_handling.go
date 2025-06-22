package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
07_error_handling.go - Go语言基础：错误处理
学习内容：
1. 错误类型和error接口
2. 创建和返回错误
3. 错误处理最佳实践
4. 自定义错误类型
5. panic和recover
*/

// 自定义错误类型
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 - 字段: %s, 值: %v, 消息: %s", e.Field, e.Value, e.Message)
}

type DivideByZeroError struct {
	Dividend float64
}

func (e DivideByZeroError) Error() string {
	return fmt.Sprintf("除零错误：无法将 %.2f 除以零", e.Dividend)
}

// 简单除法函数，返回错误
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivideByZeroError{Dividend: a}
	}
	return a / b, nil
}

// 字符串转整数，带错误处理
func parseInteger(s string) (int, error) {
	if s == "" {
		return 0, errors.New("输入字符串不能为空")
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("无法解析字符串 '%s' 为整数: %w", s, err)
	}

	if num < 0 {
		return 0, ValidationError{
			Field:   "number",
			Value:   num,
			Message: "数字不能为负数",
		}
	}

	return num, nil
}

// 用户结构体和验证
type User struct {
	Name  string
	Email string
	Age   int
}

func validateUser(user User) error {
	if user.Name == "" {
		return ValidationError{
			Field:   "name",
			Value:   user.Name,
			Message: "用户名不能为空",
		}
	}

	if len(user.Name) < 2 {
		return ValidationError{
			Field:   "name",
			Value:   user.Name,
			Message: "用户名长度至少为2个字符",
		}
	}

	if user.Email == "" {
		return ValidationError{
			Field:   "email",
			Value:   user.Email,
			Message: "邮箱不能为空",
		}
	}

	if user.Age < 0 || user.Age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "年龄必须在0-150之间",
		}
	}

	return nil
}

// 创建用户
func createUser(name, email string, age int) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	if err := validateUser(user); err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	return &user, nil
}

// 文件操作示例
func readFileContent(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("读取文件 '%s' 失败: %w", filename, err)
	}
	return string(data), nil
}

func writeFileContent(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("写入文件 '%s' 失败: %w", filename, err)
	}
	return nil
}

// 多个操作的链式错误处理
func processData(input string) (result int, err error) {
	// 第一步：解析字符串
	num, err := parseInteger(input)
	if err != nil {
		return 0, fmt.Errorf("数据处理第一步失败: %w", err)
	}

	// 第二步：验证范围
	if num > 1000 {
		return 0, errors.New("数据处理失败：数字超出允许范围(1000)")
	}

	// 第三步：计算
	result = num * 2
	if result < 0 { // 溢出检查
		return 0, errors.New("数据处理失败：计算结果溢出")
	}

	return result, nil
}

// panic和recover示例
func safeDivide(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}

	result = a / b
	return result, nil
}

// 模拟可能panic的函数
func riskyOperation(x int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("从panic中恢复: %v", r)
		}
	}()

	if x < 0 {
		panic("负数不被支持")
	}

	if x == 0 {
		return 0, errors.New("零值处理错误")
	}

	return 100 / x, nil
}

func main() {
	fmt.Println("=== Go语言基础：错误处理 ===")

	// 1. 基本错误处理
	fmt.Println("\n1. 基本错误处理：")

	// 除法操作
	result1, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result1)
	}

	result2, err := divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", result2)
	}

	// 2. 字符串解析错误处理
	fmt.Println("\n2. 字符串解析错误处理：")

	testStrings := []string{"123", "abc", "", "-5", "999"}

	for _, s := range testStrings {
		num, err := parseInteger(s)
		if err != nil {
			fmt.Printf("解析 '%s' 失败: %v\n", s, err)
		} else {
			fmt.Printf("解析 '%s' 成功: %d\n", s, num)
		}
	}

	// 3. 自定义错误类型
	fmt.Println("\n3. 自定义错误类型：")

	users := []struct {
		name  string
		email string
		age   int
	}{
		{"张三", "zhangsan@example.com", 25},
		{"", "invalid@example.com", 30},    // 名字为空
		{"李四", "", 25},                     // 邮箱为空
		{"王五", "wangwu@example.com", -5},   // 年龄无效
		{"赵六", "zhaoliu@example.com", 200}, // 年龄超出范围
	}

	for i, userData := range users {
		user, err := createUser(userData.name, userData.email, userData.age)
		if err != nil {
			fmt.Printf("用户%d创建失败: %v\n", i+1, err)

			// 类型断言检查具体错误类型
			var validationErr ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("  -> 这是一个验证错误，字段: %s\n", validationErr.Field)
			}
		} else {
			fmt.Printf("用户%d创建成功: %+v\n", i+1, *user)
		}
	}

	// 4. 错误包装和解包
	fmt.Println("\n4. 错误包装和解包：")

	testInputs := []string{"42", "abc", "1500", "-10"}

	for _, input := range testInputs {
		result, err := processData(input)
		if err != nil {
			fmt.Printf("处理 '%s' 失败: %v\n", input, err)

			// 检查是否包含特定错误
			var validationErr ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("  -> 包含验证错误: %s\n", validationErr.Message)
			}
		} else {
			fmt.Printf("处理 '%s' 成功: %d\n", input, result)
		}
	}

	// 5. 文件操作错误处理
	fmt.Println("\n5. 文件操作错误处理：")

	// 尝试读取不存在的文件
	_, err = readFileContent("nonexistent.txt")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)

		// 检查是否是文件不存在错误
		if os.IsNotExist(err) {
			fmt.Println("  -> 文件不存在")
		}
	}

	// 写入文件
	testContent := "这是测试内容\n测试错误处理\n"
	err = writeFileContent("test.txt", testContent)
	if err != nil {
		fmt.Printf("写入文件错误: %v\n", err)
	} else {
		fmt.Println("文件写入成功")

		// 读取刚写入的文件
		content, err := readFileContent("test.txt")
		if err != nil {
			fmt.Printf("读取文件错误: %v\n", err)
		} else {
			fmt.Printf("文件内容: %s", content)
		}

		// 清理测试文件
		os.Remove("test.txt")
	}

	// 6. panic和recover
	fmt.Println("\n6. panic和recover：")

	// 安全除法
	result3, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("安全除法错误: %v\n", err)
	} else {
		fmt.Printf("安全除法结果: %.2f\n", result3)
	}

	result4, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("安全除法错误: %v\n", err)
	} else {
		fmt.Printf("安全除法结果: %.2f\n", result4)
	}

	// 风险操作
	fmt.Println("\n风险操作测试:")
	riskValues := []int{5, 0, -1, 2}

	for _, val := range riskValues {
		result, err := riskyOperation(val)
		if err != nil {
			fmt.Printf("riskyOperation(%d) 错误: %v\n", val, err)
		} else {
			fmt.Printf("riskyOperation(%d) 结果: %d\n", val, result)
		}
	}

	// 7. 错误处理最佳实践总结
	fmt.Println("\n7. 错误处理最佳实践：")

	fmt.Println("✅ 总是检查错误")
	fmt.Println("✅ 提供有意义的错误信息")
	fmt.Println("✅ 使用errors.New()或fmt.Errorf()创建错误")
	fmt.Println("✅ 使用错误包装（%w）保留原始错误")
	fmt.Println("✅ 定义自定义错误类型便于错误分类")
	fmt.Println("✅ 使用errors.Is()和errors.As()进行错误检查")
	fmt.Println("✅ 谨慎使用panic，主要用于不可恢复的错误")
	fmt.Println("✅ 使用defer和recover处理panic")

	// 8. 实用示例：输入验证器
	fmt.Println("\n8. 实用示例：输入验证器")

	inputs := map[string]string{
		"有效数字": "123",
		"无效数字": "abc",
		"空字符串": "",
		"负数":   "-42",
		"过大数字": "99999",
	}

	for desc, input := range inputs {
		if result, err := validateInput(input); err != nil {
			fmt.Printf("%s '%s': ❌ %v\n", desc, input, err)
		} else {
			fmt.Printf("%s '%s': ✅ 验证通过，值: %d\n", desc, input, result)
		}
	}
}

// 输入验证器示例
func validateInput(input string) (int, error) {
	// 第一层：基本验证
	if input == "" {
		return 0, errors.New("输入不能为空")
	}

	// 第二层：格式验证
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("输入格式错误: %w", err)
	}

	// 第三层：业务规则验证
	if num < 0 {
		return 0, ValidationError{
			Field:   "input",
			Value:   num,
			Message: "数值不能为负数",
		}
	}

	if num > 10000 {
		return 0, ValidationError{
			Field:   "input",
			Value:   num,
			Message: "数值不能超过10000",
		}
	}

	return num, nil
}
