package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

/*
06_time_crypto.go - Go标准库：时间处理和基础加密
涉及包：
- time: 时间和日期处理
- crypto/md5: MD5哈希
- crypto/sha256: SHA256哈希
- crypto/rand: 安全随机数
- encoding/base64: Base64编码
- encoding/hex: 十六进制编码
*/

func main() {
	fmt.Println("=== Go标准库：时间处理和基础加密 ===")

	// 1. 基本时间操作
	fmt.Println("\n1. 基本时间操作：")

	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix时间戳: %d\n", now.Unix())
	fmt.Printf("纳秒时间戳: %d\n", now.UnixNano())

	// 时间的各个组成部分
	fmt.Printf("年份: %d\n", now.Year())
	fmt.Printf("月份: %s (%d)\n", now.Month(), int(now.Month()))
	fmt.Printf("日期: %d\n", now.Day())
	fmt.Printf("星期: %s\n", now.Weekday())
	fmt.Printf("小时: %d\n", now.Hour())
	fmt.Printf("分钟: %d\n", now.Minute())
	fmt.Printf("秒数: %d\n", now.Second())

	// 2. 时间格式化
	fmt.Println("\n2. 时间格式化：")

	// Go的时间格式化使用特定的参考时间: Mon Jan 2 15:04:05 MST 2006
	formats := map[string]string{
		"标准格式":    "2006-01-02 15:04:05",
		"ISO8601": "2006-01-02T15:04:05Z07:00",
		"RFC3339": time.RFC3339,
		"日期":      "2006年01月02日",
		"时间":      "15:04:05",
		"12小时制":   "2006-01-02 03:04:05 PM",
		"简短格式":    "06/01/02",
	}

	for name, format := range formats {
		fmt.Printf("%s: %s\n", name, now.Format(format))
	}

	// 3. 时间解析
	fmt.Println("\n3. 时间解析：")

	timeStrings := []string{
		"2023-12-25 15:30:45",
		"2023/12/25 15:30:45",
		"Dec 25, 2023 3:30:45 PM",
		"2023-12-25T15:30:45Z",
	}

	timeFormats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"Jan 2, 2006 3:04:05 PM",
		"2006-01-02T15:04:05Z",
	}

	for i, timeStr := range timeStrings {
		if parsedTime, err := time.Parse(timeFormats[i], timeStr); err != nil {
			fmt.Printf("解析失败 '%s': %v\n", timeStr, err)
		} else {
			fmt.Printf("解析成功 '%s': %s\n", timeStr, parsedTime.Format("2006-01-02 15:04:05"))
		}
	}

	// 4. 时间计算
	fmt.Println("\n4. 时间计算：")

	// 时间加减
	future := now.Add(24 * time.Hour)
	past := now.Add(-7 * 24 * time.Hour)

	fmt.Printf("24小时后: %s\n", future.Format("2006-01-02 15:04:05"))
	fmt.Printf("7天前: %s\n", past.Format("2006-01-02 15:04:05"))

	// 时间差计算
	duration := future.Sub(now)
	fmt.Printf("时间差: %s\n", duration)
	fmt.Printf("时间差(小时): %.1f\n", duration.Hours())
	fmt.Printf("时间差(分钟): %.0f\n", duration.Minutes())

	// 5. 特定时间创建
	fmt.Println("\n5. 特定时间创建：")

	// 创建特定时间
	birthday := time.Date(1990, time.May, 15, 14, 30, 0, 0, time.UTC)
	fmt.Printf("生日: %s\n", birthday.Format("2006年01月02日 15:04:05"))

	// 计算年龄
	age := now.Sub(birthday)
	years := int(age.Hours() / (24 * 365.25))
	fmt.Printf("年龄: 大约%d岁\n", years)

	// 6. 时区处理
	fmt.Println("\n6. 时区处理：")

	// 加载不同时区
	locations := []string{
		"UTC",
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Asia/Shanghai",
	}

	for _, locName := range locations {
		if loc, err := time.LoadLocation(locName); err != nil {
			fmt.Printf("加载时区失败 %s: %v\n", locName, err)
		} else {
			localTime := now.In(loc)
			fmt.Printf("%s: %s\n", locName, localTime.Format("2006-01-02 15:04:05 MST"))
		}
	}

	// 7. 定时器和延时
	fmt.Println("\n7. 定时器演示：")

	// 延时执行
	fmt.Println("开始等待...")
	time.Sleep(1 * time.Second)
	fmt.Println("等待1秒完成")

	// 使用Timer
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("定时器触发（2秒后）")
	}()

	// 等待定时器完成
	time.Sleep(3 * time.Second)

	// 8. MD5哈希
	fmt.Println("\n8. MD5哈希：")

	data := "Hello, Go Programming!"
	md5Hash := md5.Sum([]byte(data))
	md5Hex := hex.EncodeToString(md5Hash[:])
	fmt.Printf("原文: %s\n", data)
	fmt.Printf("MD5: %s\n", md5Hex)

	// 验证MD5
	verifyData := "Hello, Go Programming!"
	verifyMD5 := md5.Sum([]byte(verifyData))
	verifyMD5Hex := hex.EncodeToString(verifyMD5[:])
	fmt.Printf("验证MD5: %s (匹配: %t)\n", verifyMD5Hex, md5Hex == verifyMD5Hex)

	// 9. SHA256哈希
	fmt.Println("\n9. SHA256哈希：")

	sha256Hash := sha256.Sum256([]byte(data))
	sha256Hex := hex.EncodeToString(sha256Hash[:])
	fmt.Printf("SHA256: %s\n", sha256Hex)

	// 10. Base64编码
	fmt.Println("\n10. Base64编码：")

	originalText := "Go语言学习笔记"

	// 标准Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(originalText))
	fmt.Printf("原文: %s\n", originalText)
	fmt.Printf("Base64编码: %s\n", encoded)

	// Base64解码
	if decoded, err := base64.StdEncoding.DecodeString(encoded); err != nil {
		fmt.Printf("Base64解码失败: %v\n", err)
	} else {
		fmt.Printf("Base64解码: %s\n", string(decoded))
	}

	// URL安全的Base64编码
	urlEncoded := base64.URLEncoding.EncodeToString([]byte(originalText))
	fmt.Printf("URL安全Base64: %s\n", urlEncoded)

	// 11. 随机数生成
	fmt.Println("\n11. 安全随机数生成：")

	// 生成随机字节
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		fmt.Printf("生成随机数失败: %v\n", err)
	} else {
		fmt.Printf("随机字节(hex): %s\n", hex.EncodeToString(randomBytes))
		fmt.Printf("随机字节(base64): %s\n", base64.StdEncoding.EncodeToString(randomBytes))
	}

	// 12. 实用函数：生成随机字符串
	fmt.Println("\n12. 生成随机字符串：")

	randomString := generateRandomString(10)
	fmt.Printf("随机字符串(10位): %s\n", randomString)

	randomToken := generateRandomToken(32)
	fmt.Printf("随机令牌(32字节): %s\n", randomToken)

	// 13. 密码哈希示例
	fmt.Println("\n13. 简单密码哈希：")

	password := "mySecretPassword123"
	salt := generateRandomString(16)

	hashedPassword := hashPassword(password, salt)
	fmt.Printf("原密码: %s\n", password)
	fmt.Printf("盐值: %s\n", salt)
	fmt.Printf("哈希后: %s\n", hashedPassword)

	// 验证密码
	isValid := verifyPassword(password, salt, hashedPassword)
	fmt.Printf("密码验证: %t\n", isValid)

	// 14. 时间性能测试
	fmt.Println("\n14. 时间性能测试：")

	// 测试函数执行时间
	start := time.Now()

	// 模拟一些工作
	total := 0
	for i := 0; i < 1000000; i++ {
		total += i
	}

	duration = time.Since(start)
	fmt.Printf("计算完成，结果: %d\n", total)
	fmt.Printf("执行时间: %s\n", duration)
	fmt.Printf("执行时间(纳秒): %d\n", duration.Nanoseconds())

	// 15. 时间格式化实用示例
	fmt.Println("\n15. 时间格式化实用示例：")

	// 文件名时间戳
	filename := fmt.Sprintf("backup_%s.sql", now.Format("20060102_150405"))
	fmt.Printf("备份文件名: %s\n", filename)

	// 日志时间戳
	logTime := fmt.Sprintf("[%s]", now.Format("2006-01-02 15:04:05.000"))
	fmt.Printf("日志时间戳: %s 用户登录\n", logTime)

	// 相对时间显示
	past2 := now.Add(-2 * time.Hour)
	relativeTime := getRelativeTime(past2, now)
	fmt.Printf("相对时间: %s\n", relativeTime)

	fmt.Println("\n时间处理和加密演示完成！")
}

// 生成随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		randomBytes := make([]byte, 1)
		rand.Read(randomBytes)
		b[i] = charset[randomBytes[0]%byte(len(charset))]
	}
	return string(b)
}

// 生成随机令牌
func generateRandomToken(length int) string {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

// 简单密码哈希
func hashPassword(password, salt string) string {
	combined := password + salt
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

// 验证密码
func verifyPassword(password, salt, hashedPassword string) bool {
	return hashPassword(password, salt) == hashedPassword
}

// 获取相对时间描述
func getRelativeTime(past, now time.Time) string {
	duration := now.Sub(past)

	switch {
	case duration < time.Minute:
		return "刚刚"
	case duration < time.Hour:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d分钟前", minutes)
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		return fmt.Sprintf("%d小时前", hours)
	case duration < 7*24*time.Hour:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d天前", days)
	default:
		return past.Format("2006-01-02")
	}
}
