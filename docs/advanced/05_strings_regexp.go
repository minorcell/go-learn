package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
05_strings_regexp.go - Go标准库：字符串处理和正则表达式
涉及包：
- strings: 字符串操作
- strconv: 字符串转换
- regexp: 正则表达式
- unicode: Unicode处理
- unicode/utf8: UTF-8编码处理
*/

func main() {
	fmt.Println("=== Go标准库：字符串处理和正则表达式 ===")

	// 1. 基本字符串操作
	fmt.Println("\n1. 基本字符串操作：")

	text := "  Hello, Go Programming World!  "
	fmt.Printf("原字符串: '%s'\n", text)

	// 去除空格
	trimmed := strings.TrimSpace(text)
	fmt.Printf("去除空格: '%s'\n", trimmed)

	// 大小写转换
	fmt.Printf("转大写: %s\n", strings.ToUpper(trimmed))
	fmt.Printf("转小写: %s\n", strings.ToLower(trimmed))
	fmt.Printf("标题格式: %s\n", strings.Title(strings.ToLower(trimmed)))

	// 字符串长度和判断
	fmt.Printf("字符串长度: %d\n", len(trimmed))
	fmt.Printf("UTF-8字符数: %d\n", utf8.RuneCountInString(trimmed))
	fmt.Printf("是否以Hello开头: %t\n", strings.HasPrefix(trimmed, "Hello"))
	fmt.Printf("是否以!结尾: %t\n", strings.HasSuffix(trimmed, "!"))
	fmt.Printf("是否包含Go: %t\n", strings.Contains(trimmed, "Go"))

	// 2. 字符串分割和连接
	fmt.Println("\n2. 字符串分割和连接：")

	sentence := "apple,banana,orange,grape"
	fruits := strings.Split(sentence, ",")
	fmt.Printf("分割结果: %v\n", fruits)

	// 连接字符串
	joined := strings.Join(fruits, " | ")
	fmt.Printf("连接结果: %s\n", joined)

	// 字段分割（按空白字符）
	text2 := "  Go  is   awesome  "
	fields := strings.Fields(text2)
	fmt.Printf("字段分割: %v\n", fields)

	// 3. 字符串替换
	fmt.Println("\n3. 字符串替换：")

	original := "I love Java and Java is great"

	// 替换所有
	replaced := strings.ReplaceAll(original, "Java", "Go")
	fmt.Printf("全部替换: %s\n", replaced)

	// 替换指定次数
	partialReplace := strings.Replace(original, "Java", "Go", 1)
	fmt.Printf("替换1次: %s\n", partialReplace)

	// 使用Replacer进行多重替换
	replacer := strings.NewReplacer(
		"Java", "Go",
		"great", "fantastic",
		"love", "adore",
	)
	multiReplace := replacer.Replace(original)
	fmt.Printf("多重替换: %s\n", multiReplace)

	// 4. 字符串查找
	fmt.Println("\n4. 字符串查找：")

	searchText := "Go is a programming language. Go is fast."

	// 查找索引
	index := strings.Index(searchText, "Go")
	fmt.Printf("第一个'Go'的位置: %d\n", index)

	lastIndex := strings.LastIndex(searchText, "Go")
	fmt.Printf("最后一个'Go'的位置: %d\n", lastIndex)

	// 统计出现次数
	count := strings.Count(searchText, "Go")
	fmt.Printf("'Go'出现次数: %d\n", count)

	// 5. 类型转换 (strconv)
	fmt.Println("\n5. 类型转换 (strconv)：")

	// 字符串转数字
	numStr := "123"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("字符串'%s'转整数: %d\n", numStr, num)
	}

	// 字符串转浮点数
	floatStr := "3.14159"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("字符串'%s'转浮点数: %.3f\n", floatStr, floatNum)
	}

	// 字符串转布尔值
	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("字符串'%s'转布尔值: %t\n", boolStr, boolVal)
	}

	// 数字转字符串
	fmt.Printf("整数42转字符串: '%s'\n", strconv.Itoa(42))
	fmt.Printf("浮点数3.14转字符串: '%s'\n", strconv.FormatFloat(3.14, 'f', 2, 64))
	fmt.Printf("布尔值true转字符串: '%s'\n", strconv.FormatBool(true))

	// 不同进制转换
	fmt.Printf("二进制1010转十进制: %s\n", parseAndFormat("1010", 2))
	fmt.Printf("十六进制FF转十进制: %s\n", parseAndFormat("FF", 16))
	fmt.Printf("八进制77转十进制: %s\n", parseAndFormat("77", 8))

	// 6. 正则表达式基础
	fmt.Println("\n6. 正则表达式基础：")

	// 编译正则表达式
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	emails := []string{
		"user@example.com",
		"test.email+tag@domain.co.uk",
		"invalid-email",
		"another@test.org",
	}

	fmt.Println("邮箱验证:")
	for _, email := range emails {
		if emailPattern.MatchString(email) {
			fmt.Printf("✅ %s - 有效邮箱\n", email)
		} else {
			fmt.Printf("❌ %s - 无效邮箱\n", email)
		}
	}

	// 7. 正则表达式查找和提取
	fmt.Println("\n7. 正则表达式查找和提取：")

	text3 := "联系方式：电话 138-1234-5678，邮箱 zhang@company.com，QQ 123456789"

	// 查找电话号码
	phonePattern := regexp.MustCompile(`\d{3}-\d{4}-\d{4}`)
	phone := phonePattern.FindString(text3)
	fmt.Printf("找到电话号码: %s\n", phone)

	// 查找邮箱
	emailPattern2 := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email := emailPattern2.FindString(text3)
	fmt.Printf("找到邮箱: %s\n", email)

	// 查找所有数字
	numberPattern := regexp.MustCompile(`\d+`)
	numbers := numberPattern.FindAllString(text3, -1)
	fmt.Printf("找到所有数字: %v\n", numbers)

	// 8. 正则表达式分组
	fmt.Println("\n8. 正则表达式分组：")

	dateText := "今天是2023年12月25日，明天是2023年12月26日"
	datePattern := regexp.MustCompile(`(\d{4})年(\d{1,2})月(\d{1,2})日`)

	matches := datePattern.FindAllStringSubmatch(dateText, -1)
	fmt.Println("日期提取:")
	for i, match := range matches {
		fmt.Printf("第%d个日期: 完整匹配='%s', 年='%s', 月='%s', 日='%s'\n",
			i+1, match[0], match[1], match[2], match[3])
	}

	// 9. 正则表达式替换
	fmt.Println("\n9. 正则表达式替换：")

	// 隐藏电话号码中间4位
	phoneText := "我的电话是138-1234-5678，请联系"
	phoneHidePattern := regexp.MustCompile(`(\d{3})-(\d{4})-(\d{4})`)
	hiddenPhone := phoneHidePattern.ReplaceAllString(phoneText, "$1-****-$3")
	fmt.Printf("隐藏电话: %s\n", hiddenPhone)

	// 格式化日期
	dateFormatPattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	dateStr := "生日: 1990-05-15"
	formattedDate := dateFormatPattern.ReplaceAllString(dateStr, "$1年$2月$3日")
	fmt.Printf("格式化日期: %s\n", formattedDate)

	// 10. Unicode和中文处理
	fmt.Println("\n10. Unicode和中文处理：")

	chineseText := "Hello, 世界！你好，Go语言！"
	fmt.Printf("原文本: %s\n", chineseText)
	fmt.Printf("字节长度: %d\n", len(chineseText))
	fmt.Printf("字符数量: %d\n", utf8.RuneCountInString(chineseText))

	// 遍历Unicode字符
	fmt.Println("逐字符分析:")
	for i, r := range chineseText {
		fmt.Printf("  位置%d: 字符'%c', Unicode码点U+%04X, 类别: %s\n",
			i, r, r, getUnicodeCategory(r))
	}

	// 提取中文字符
	chinesePattern := regexp.MustCompile(`[\p{Han}]+`)
	chineseWords := chinesePattern.FindAllString(chineseText, -1)
	fmt.Printf("提取的中文: %v\n", chineseWords)

	// 11. 字符串格式化和模板
	fmt.Println("\n11. 字符串格式化：")

	name := "张三"
	age := 25
	score := 95.5

	// 使用Printf格式化
	fmt.Printf("姓名: %s, 年龄: %d, 分数: %.1f\n", name, age, score)

	// 使用Sprintf创建格式化字符串
	info := fmt.Sprintf("学生信息 - 姓名: %s, 年龄: %d岁, 分数: %.1f分", name, age, score)
	fmt.Printf("格式化结果: %s\n", info)

	// 12. 实用字符串函数
	fmt.Println("\n12. 实用字符串函数：")

	// 重复字符串
	repeated := strings.Repeat("Go! ", 3)
	fmt.Printf("重复字符串: %s\n", repeated)

	// 字符串构建器（高效拼接）
	var builder strings.Builder
	words := []string{"Go", "is", "awesome", "for", "programming"}
	for i, word := range words {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(word)
	}
	fmt.Printf("构建器结果: %s\n", builder.String())

	// 13. 文本清理和验证
	fmt.Println("\n13. 文本清理和验证：")

	// 清理用户输入
	userInput := "  <script>alert('hello')</script>User Input  "
	cleaned := cleanUserInput(userInput)
	fmt.Printf("原输入: %s\n", userInput)
	fmt.Printf("清理后: %s\n", cleaned)

	// 验证用户名
	usernames := []string{"user123", "test_user", "admin", "user@domain", "123"}
	fmt.Println("用户名验证:")
	for _, username := range usernames {
		if isValidUsername(username) {
			fmt.Printf("✅ %s - 有效用户名\n", username)
		} else {
			fmt.Printf("❌ %s - 无效用户名\n", username)
		}
	}

	fmt.Println("\n字符串处理演示完成！")
}

// 进制转换辅助函数
func parseAndFormat(s string, base int) string {
	if num, err := strconv.ParseInt(s, base, 64); err == nil {
		return strconv.FormatInt(num, 10)
	}
	return "转换失败"
}

// 获取Unicode字符类别
func getUnicodeCategory(r rune) string {
	switch {
	case unicode.IsLetter(r):
		return "字母"
	case unicode.IsDigit(r):
		return "数字"
	case unicode.IsSpace(r):
		return "空格"
	case unicode.IsPunct(r):
		return "标点"
	case unicode.IsSymbol(r):
		return "符号"
	default:
		return "其他"
	}
}

// 清理用户输入
func cleanUserInput(input string) string {
	// 去除首尾空格
	cleaned := strings.TrimSpace(input)

	// 移除HTML标签（简单实现）
	htmlPattern := regexp.MustCompile(`<[^>]*>`)
	cleaned = htmlPattern.ReplaceAllString(cleaned, "")

	// 移除多余空格
	spacePattern := regexp.MustCompile(`\s+`)
	cleaned = spacePattern.ReplaceAllString(cleaned, " ")

	return cleaned
}

// 验证用户名（只允许字母、数字、下划线，长度3-20）
func isValidUsername(username string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	return pattern.MatchString(username)
}
