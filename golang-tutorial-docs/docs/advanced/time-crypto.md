---
title: 时间处理和加密
description: 学习Go语言的时间操作和密码学应用
---

# 时间处理和加密

时间处理和安全加密是现代应用开发的重要组成部分。Go语言提供了强大的time包和crypto包，让我们深入学习这些功能。

## 本章内容

- 时间基础操作和格式化
- 时间计算和时区处理
- 定时器和周期任务
- 哈希算法和数字签名
- 对称和非对称加密

## 时间处理

### 基础时间操作

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 时间创建和获取
    timeCreationAndGet()
    
    // 时间格式化
    timeFormatting()
    
    // 时间计算
    timeCalculation()
    
    // 时区处理
    timezoneHandling()
}

// 时间创建和获取
func timeCreationAndGet() {
    fmt.Println("=== 时间创建和获取 ===")
    
    // 获取当前时间
    now := time.Now()
    fmt.Printf("当前时间: %v\n", now)
    fmt.Printf("Unix时间戳: %d\n", now.Unix())
    fmt.Printf("Unix毫秒时间戳: %d\n", now.UnixMilli())
    fmt.Printf("Unix纳秒时间戳: %d\n", now.UnixNano())
    
    // 创建特定时间
    specificTime := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
    fmt.Printf("指定时间: %v\n", specificTime)
    
    // 从时间戳创建时间
    timestamp := int64(1705312245)
    fromTimestamp := time.Unix(timestamp, 0)
    fmt.Printf("从时间戳创建: %v\n", fromTimestamp)
    
    // 解析时间字符串
    timeStr := "2024-01-15 10:30:45"
    parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
    if err == nil {
        fmt.Printf("解析时间字符串: %v\n", parsed)
    }
    
    // 获取时间组件
    fmt.Printf("年: %d, 月: %d, 日: %d\n", now.Year(), int(now.Month()), now.Day())
    fmt.Printf("时: %d, 分: %d, 秒: %d\n", now.Hour(), now.Minute(), now.Second())
    fmt.Printf("星期: %s\n", now.Weekday())
    fmt.Printf("年中第几天: %d\n", now.YearDay())
    
    fmt.Println()
}

// 时间格式化
func timeFormatting() {
    fmt.Println("=== 时间格式化 ===")
    
    now := time.Now()
    
    // 常用格式
    fmt.Println("常用时间格式:")
    fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
    fmt.Printf("RFC822: %s\n", now.Format(time.RFC822))
    fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))
    fmt.Printf("Stamp: %s\n", now.Format(time.Stamp))
    
    // 自定义格式 (Go的参考时间: Mon Jan 2 15:04:05 MST 2006)
    fmt.Println("自定义格式:")
    fmt.Printf("日期: %s\n", now.Format("2006-01-02"))
    fmt.Printf("时间: %s\n", now.Format("15:04:05"))
    fmt.Printf("日期时间: %s\n", now.Format("2006-01-02 15:04:05"))
    fmt.Printf("中文格式: %s\n", now.Format("2006年01月02日 15时04分05秒"))
    fmt.Printf("12小时制: %s\n", now.Format("2006-01-02 03:04:05 PM"))
    fmt.Printf("ISO 8601: %s\n", now.Format("2006-01-02T15:04:05Z07:00"))
    
    // 自定义分隔符
    fmt.Printf("斜线分隔: %s\n", now.Format("01/02/2006"))
    fmt.Printf("点分隔: %s\n", now.Format("02.01.2006"))
    fmt.Printf("无分隔符: %s\n", now.Format("20060102150405"))
    
    // 解析不同格式
    timeFormats := []string{
        "2024-01-15",
        "2024/01/15",
        "15-01-2024",
        "2024-01-15 10:30:45",
        "15/01/2024 10:30:45",
    }
    
    layouts := []string{
        "2006-01-02",
        "2006/01/02",
        "02-01-2006",
        "2006-01-02 15:04:05",
        "02/01/2006 15:04:05",
    }
    
    fmt.Println("解析不同格式:")
    for i, timeStr := range timeFormats {
        if parsed, err := time.Parse(layouts[i], timeStr); err == nil {
            fmt.Printf("  %s -> %s\n", timeStr, parsed.Format("2006-01-02 15:04:05"))
        } else {
            fmt.Printf("  %s -> 解析失败: %v\n", timeStr, err)
        }
    }
    
    fmt.Println()
}

// 时间计算
func timeCalculation() {
    fmt.Println("=== 时间计算 ===")
    
    now := time.Now()
    fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
    
    // 时间加减
    fmt.Println("时间加减:")
    fmt.Printf("1小时后: %s\n", now.Add(time.Hour).Format("2006-01-02 15:04:05"))
    fmt.Printf("30分钟前: %s\n", now.Add(-30*time.Minute).Format("2006-01-02 15:04:05"))
    fmt.Printf("1天后: %s\n", now.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"))
    fmt.Printf("1个月后: %s\n", now.AddDate(0, 1, 0).Format("2006-01-02 15:04:05"))
    fmt.Printf("1年后: %s\n", now.AddDate(1, 0, 0).Format("2006-01-02 15:04:05"))
    
    // 时间比较
    future := now.Add(2 * time.Hour)
    past := now.Add(-2 * time.Hour)
    
    fmt.Println("时间比较:")
    fmt.Printf("未来时间 > 当前时间: %t\n", future.After(now))
    fmt.Printf("过去时间 < 当前时间: %t\n", past.Before(now))
    fmt.Printf("时间相等: %t\n", now.Equal(now))
    
    // 时间差计算
    fmt.Println("时间差计算:")
    duration1 := future.Sub(now)
    duration2 := now.Sub(past)
    
    fmt.Printf("未来时间 - 当前时间: %v\n", duration1)
    fmt.Printf("当前时间 - 过去时间: %v\n", duration2)
    fmt.Printf("总小时数: %.2f\n", duration1.Hours())
    fmt.Printf("总分钟数: %.2f\n", duration1.Minutes())
    fmt.Printf("总秒数: %.2f\n", duration1.Seconds())
    fmt.Printf("总毫秒数: %d\n", duration1.Milliseconds())
    
    // 时间范围检查
    start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
    check := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
    
    inRange := check.After(start) && check.Before(end)
    fmt.Printf("时间 %s 在范围内: %t\n", check.Format("2006-01-02"), inRange)
    
    fmt.Println()
}

// 时区处理
func timezoneHandling() {
    fmt.Println("=== 时区处理 ===")
    
    now := time.Now()
    
    // 不同时区的当前时间
    locations := map[string]string{
        "UTC":       "UTC",
        "纽约":        "America/New_York",
        "伦敦":        "Europe/London",
        "东京":        "Asia/Tokyo",
        "上海":        "Asia/Shanghai",
        "悉尼":        "Australia/Sydney",
    }
    
    fmt.Println("世界时间:")
    for city, timezone := range locations {
        if loc, err := time.LoadLocation(timezone); err == nil {
            localTime := now.In(loc)
            fmt.Printf("%-6s: %s\n", city, localTime.Format("2006-01-02 15:04:05 MST"))
        }
    }
    
    // 时区转换
    fmt.Println("时区转换:")
    utcTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
    fmt.Printf("UTC时间: %s\n", utcTime.Format("2006-01-02 15:04:05 MST"))
    
    if bjLoc, err := time.LoadLocation("Asia/Shanghai"); err == nil {
        bjTime := utcTime.In(bjLoc)
        fmt.Printf("北京时间: %s\n", bjTime.Format("2006-01-02 15:04:05 MST"))
    }
    
    if nyLoc, err := time.LoadLocation("America/New_York"); err == nil {
        nyTime := utcTime.In(nyLoc)
        fmt.Printf("纽约时间: %s\n", nyTime.Format("2006-01-02 15:04:05 MST"))
    }
    
    // 夏令时处理
    fmt.Println("夏令时示例:")
    if nyLoc, err := time.LoadLocation("America/New_York"); err == nil {
        summer := time.Date(2024, 7, 15, 12, 0, 0, 0, nyLoc)
        winter := time.Date(2024, 1, 15, 12, 0, 0, 0, nyLoc)
        
        fmt.Printf("夏季纽约时间: %s (UTC%s)\n", 
            summer.Format("2006-01-02 15:04:05 MST"), formatOffset(summer))
        fmt.Printf("冬季纽约时间: %s (UTC%s)\n", 
            winter.Format("2006-01-02 15:04:05 MST"), formatOffset(winter))
    }
    
    fmt.Println()
}

// 格式化时区偏移
func formatOffset(t time.Time) string {
    _, offset := t.Zone()
    hours := offset / 3600
    minutes := (offset % 3600) / 60
    
    sign := "+"
    if offset < 0 {
        sign = "-"
        hours = -hours
        minutes = -minutes
    }
    
    return fmt.Sprintf("%s%02d:%02d", sign, hours, minutes)
}
```

### 定时器和周期任务

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    // 基础定时器
    basicTimers()
    
    // 周期任务
    periodicTasks()
    
    // 超时控制
    timeoutControl()
    
    // 任务调度器
    taskScheduler()
}

// 基础定时器
func basicTimers() {
    fmt.Println("=== 基础定时器 ===")
    
    // 延时执行
    fmt.Println("3秒后执行:")
    timer := time.NewTimer(3 * time.Second)
    go func() {
        <-timer.C
        fmt.Println("定时器触发!")
    }()
    
    // 等待定时器
    time.Sleep(3500 * time.Millisecond)
    
    // 周期性定时器
    fmt.Println("每1秒执行一次 (持续5秒):")
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    stopTime := time.Now().Add(5 * time.Second)
    for {
        select {
        case t := <-ticker.C:
            fmt.Printf("  Tick at %s\n", t.Format("15:04:05"))
            if time.Now().After(stopTime) {
                fmt.Println("周期任务完成")
                goto next
            }
        }
    }
    
next:
    // 简单延时
    fmt.Println("使用time.Sleep延时2秒...")
    time.Sleep(2 * time.Second)
    fmt.Println("延时完成")
    
    fmt.Println()
}

// 周期任务
func periodicTasks() {
    fmt.Println("=== 周期任务 ===")
    
    var wg sync.WaitGroup
    
    // 任务1：每2秒记录一次日志
    wg.Add(1)
    go func() {
        defer wg.Done()
        ticker := time.NewTicker(2 * time.Second)
        defer ticker.Stop()
        
        count := 0
        for {
            select {
            case <-ticker.C:
                count++
                fmt.Printf("[日志] 第%d次记录 - %s\n", 
                    count, time.Now().Format("15:04:05"))
                if count >= 3 {
                    return
                }
            }
        }
    }()
    
    // 任务2：每3秒清理一次缓存
    wg.Add(1)
    go func() {
        defer wg.Done()
        ticker := time.NewTicker(3 * time.Second)
        defer ticker.Stop()
        
        count := 0
        for {
            select {
            case <-ticker.C:
                count++
                fmt.Printf("[清理] 清理缓存 #%d - %s\n", 
                    count, time.Now().Format("15:04:05"))
                if count >= 2 {
                    return
                }
            }
        }
    }()
    
    // 等待所有任务完成
    wg.Wait()
    fmt.Println("所有周期任务完成")
    
    fmt.Println()
}

// 超时控制
func timeoutControl() {
    fmt.Println("=== 超时控制 ===")
    
    // 使用context进行超时控制
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // 模拟长时间运行的任务
    result := make(chan string, 1)
    go func() {
        // 模拟工作
        time.Sleep(2 * time.Second)
        result <- "任务完成"
    }()
    
    select {
    case res := <-result:
        fmt.Printf("%s\n", res)
    case <-ctx.Done():
        fmt.Printf("任务超时: %v\n", ctx.Err())
    }
    
    // 模拟超时的任务
    ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel2()
    
    result2 := make(chan string, 1)
    go func() {
        // 模拟较长的工作
        time.Sleep(2 * time.Second)
        result2 <- "长任务完成"
    }()
    
    select {
    case res := <-result2:
        fmt.Printf("%s\n", res)
    case <-ctx2.Done():
        fmt.Printf("长任务超时: %v\n", ctx2.Err())
    }
    
    // 使用time.After进行超时控制
    fmt.Println("使用time.After:")
    start := time.Now()
    select {
    case <-time.After(2 * time.Second):
        fmt.Printf("2秒超时触发 (耗时: %v)\n", time.Since(start))
    }
    
    fmt.Println()
}

// 任务调度器
func taskScheduler() {
    fmt.Println("=== 任务调度器 ===")
    
    scheduler := NewScheduler()
    scheduler.Start()
    defer scheduler.Stop()
    
    // 添加任务
    scheduler.AddTask("task1", 2*time.Second, func() {
        fmt.Printf("[任务1] 执行时间: %s\n", time.Now().Format("15:04:05"))
    })
    
    scheduler.AddTask("task2", 3*time.Second, func() {
        fmt.Printf("[任务2] 执行时间: %s\n", time.Now().Format("15:04:05"))
    })
    
    // 延时任务
    scheduler.AddDelayedTask("delayed", 5*time.Second, func() {
        fmt.Printf("[延时任务] 执行时间: %s\n", time.Now().Format("15:04:05"))
    })
    
    // 运行10秒
    time.Sleep(10 * time.Second)
    
    // 移除任务
    scheduler.RemoveTask("task1")
    fmt.Println("已移除task1")
    
    time.Sleep(5 * time.Second)
    
    fmt.Println()
}

// 简单的任务调度器
type Scheduler struct {
    tasks map[string]*Task
    stop  chan bool
    mutex sync.RWMutex
}

type Task struct {
    name     string
    interval time.Duration
    fn       func()
    ticker   *time.Ticker
    stop     chan bool
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(map[string]*Task),
        stop:  make(chan bool),
    }
}

func (s *Scheduler) Start() {
    go func() {
        <-s.stop
        s.mutex.RLock()
        for _, task := range s.tasks {
            if task.ticker != nil {
                task.ticker.Stop()
            }
            if task.stop != nil {
                close(task.stop)
            }
        }
        s.mutex.RUnlock()
    }()
}

func (s *Scheduler) Stop() {
    close(s.stop)
}

func (s *Scheduler) AddTask(name string, interval time.Duration, fn func()) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    task := &Task{
        name:     name,
        interval: interval,
        fn:       fn,
        ticker:   time.NewTicker(interval),
        stop:     make(chan bool),
    }
    
    s.tasks[name] = task
    
    go func() {
        for {
            select {
            case <-task.ticker.C:
                task.fn()
            case <-task.stop:
                return
            }
        }
    }()
}

func (s *Scheduler) AddDelayedTask(name string, delay time.Duration, fn func()) {
    go func() {
        timer := time.NewTimer(delay)
        defer timer.Stop()
        
        select {
        case <-timer.C:
            fn()
        case <-s.stop:
            return
        }
    }()
}

func (s *Scheduler) RemoveTask(name string) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    if task, exists := s.tasks[name]; exists {
        task.ticker.Stop()
        close(task.stop)
        delete(s.tasks, name)
    }
}
```

## 加密和安全

### 哈希算法

```go
package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "hash"
    "io"
    "strings"
)

func main() {
    // 基础哈希
    basicHashing()
    
    // 文件哈希
    fileHashing()
    
    // 哈希比较
    hashComparison()
    
    // 密码哈希
    passwordHashing()
}

// 基础哈希
func basicHashing() {
    fmt.Println("=== 基础哈希算法 ===")
    
    data := "Hello, Go语言!"
    
    // MD5哈希
    md5Hash := md5.Sum([]byte(data))
    fmt.Printf("MD5:    %x\n", md5Hash)
    
    // SHA1哈希
    sha1Hash := sha1.Sum([]byte(data))
    fmt.Printf("SHA1:   %x\n", sha1Hash)
    
    // SHA256哈希
    sha256Hash := sha256.Sum256([]byte(data))
    fmt.Printf("SHA256: %x\n", sha256Hash)
    
    // SHA512哈希
    sha512Hash := sha512.Sum512([]byte(data))
    fmt.Printf("SHA512: %x\n", sha512Hash)
    
    // 使用hasher接口
    fmt.Println("\n使用hasher接口:")
    hashers := map[string]hash.Hash{
        "MD5":    md5.New(),
        "SHA1":   sha1.New(),
        "SHA256": sha256.New(),
        "SHA512": sha512.New(),
    }
    
    for name, hasher := range hashers {
        hasher.Reset()
        hasher.Write([]byte(data))
        result := hasher.Sum(nil)
        fmt.Printf("%-7s: %x\n", name, result)
    }
    
    fmt.Println()
}

// 文件哈希
func fileHashing() {
    fmt.Println("=== 文件哈希 ===")
    
    // 模拟文件内容
    fileContent := `这是一个测试文件的内容
包含多行文本
用于演示文件哈希计算`
    
    // 计算文件哈希
    fileHash := calculateFileHash(strings.NewReader(fileContent))
    fmt.Printf("文件SHA256哈希: %x\n", fileHash)
    
    // 分块计算大文件哈希
    largeContent := strings.Repeat("大文件内容重复块 ", 1000)
    largeFileHash := calculateLargeFileHash(strings.NewReader(largeContent))
    fmt.Printf("大文件SHA256哈希: %x\n", largeFileHash)
    
    fmt.Println()
}

// 计算文件哈希
func calculateFileHash(reader io.Reader) []byte {
    hasher := sha256.New()
    if _, err := io.Copy(hasher, reader); err != nil {
        fmt.Printf("计算哈希失败: %v\n", err)
        return nil
    }
    return hasher.Sum(nil)
}

// 计算大文件哈希（分块处理）
func calculateLargeFileHash(reader io.Reader) []byte {
    hasher := sha256.New()
    buffer := make([]byte, 1024) // 1KB缓冲区
    
    for {
        n, err := reader.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Printf("读取文件失败: %v\n", err)
            return nil
        }
        hasher.Write(buffer[:n])
    }
    
    return hasher.Sum(nil)
}

// 哈希比较
func hashComparison() {
    fmt.Println("=== 哈希比较 ===")
    
    // 原始数据
    original := "重要的数据内容"
    modified := "重要的数据内容!"
    
    // 计算哈希
    originalHash := sha256.Sum256([]byte(original))
    modifiedHash := sha256.Sum256([]byte(modified))
    
    fmt.Printf("原始数据: %s\n", original)
    fmt.Printf("原始哈希: %x\n", originalHash)
    fmt.Printf("修改数据: %s\n", modified)
    fmt.Printf("修改哈希: %x\n", modifiedHash)
    
    // 比较哈希
    if originalHash == modifiedHash {
        fmt.Println("数据未被修改")
    } else {
        fmt.Println("数据已被修改")
    }
    
    // 数据完整性验证
    fmt.Println("\n数据完整性验证:")
    verifyDataIntegrity(original, originalHash[:])
    verifyDataIntegrity(modified, originalHash[:])
    
    fmt.Println()
}

// 验证数据完整性
func verifyDataIntegrity(data string, expectedHash []byte) {
    actualHash := sha256.Sum256([]byte(data))
    
    if compareHashes(actualHash[:], expectedHash) {
        fmt.Printf("数据 '%s' 完整性验证通过\n", data)
    } else {
        fmt.Printf("数据 '%s' 完整性验证失败\n", data)
    }
}

// 比较哈希值
func compareHashes(hash1, hash2 []byte) bool {
    if len(hash1) != len(hash2) {
        return false
    }
    
    for i := 0; i < len(hash1); i++ {
        if hash1[i] != hash2[i] {
            return false
        }
    }
    
    return true
}

// 密码哈希
func passwordHashing() {
    fmt.Println("=== 密码哈希 ===")
    
    password := "MySecretPassword123!"
    
    // 简单哈希（不安全）
    simpleHash := sha256.Sum256([]byte(password))
    fmt.Printf("简单哈希: %x\n", simpleHash)
    
    // 加盐哈希（推荐）
    salt := "randomsalt123"
    saltedPassword := password + salt
    saltedHash := sha256.Sum256([]byte(saltedPassword))
    fmt.Printf("加盐哈希: %x (盐: %s)\n", saltedHash, salt)
    
    // 多轮哈希
    multiRoundHash := performMultiRoundHash(password, salt, 1000)
    fmt.Printf("多轮哈希: %x (1000轮)\n", multiRoundHash)
    
    // 密码验证
    fmt.Println("\n密码验证:")
    testPasswords := []string{
        "MySecretPassword123!",
        "WrongPassword",
        "MySecretPassword123",
    }
    
    for _, testPwd := range testPasswords {
        if verifyPassword(testPwd, salt, multiRoundHash) {
            fmt.Printf("密码 '%s' 验证通过\n", testPwd)
        } else {
            fmt.Printf("密码 '%s' 验证失败\n", testPwd)
        }
    }
    
    fmt.Println()
}

// 多轮哈希
func performMultiRoundHash(password, salt string, rounds int) []byte {
    result := []byte(password + salt)
    
    for i := 0; i < rounds; i++ {
        hash := sha256.Sum256(result)
        result = hash[:]
    }
    
    return result
}

// 验证密码
func verifyPassword(password, salt string, expectedHash []byte) bool {
    computedHash := performMultiRoundHash(password, salt, 1000)
    return compareHashes(computedHash, expectedHash)
}
```

### 对称加密和非对称加密

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "fmt"
    "io"
)

func main() {
    // 对称加密
    symmetricEncryption()
    
    // 非对称加密
    asymmetricEncryption()
    
    // 数字签名
    digitalSignature()
}

// 对称加密
func symmetricEncryption() {
    fmt.Println("=== 对称加密 (AES) ===")
    
    // 准备数据
    plaintext := "这是需要加密的敏感数据！包含中文和English"
    
    // 生成密钥
    key := make([]byte, 32) // AES-256需要32字节密钥
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
        fmt.Printf("生成密钥失败: %v\n", err)
        return
    }
    
    fmt.Printf("原始数据: %s\n", plaintext)
    fmt.Printf("密钥 (Base64): %s\n", base64.StdEncoding.EncodeToString(key))
    
    // 加密
    ciphertext, err := encryptAES([]byte(plaintext), key)
    if err != nil {
        fmt.Printf("加密失败: %v\n", err)
        return
    }
    
    fmt.Printf("加密数据 (Base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))
    
    // 解密
    decrypted, err := decryptAES(ciphertext, key)
    if err != nil {
        fmt.Printf("解密失败: %v\n", err)
        return
    }
    
    fmt.Printf("解密数据: %s\n", string(decrypted))
    
    // 验证
    if string(decrypted) == plaintext {
        fmt.Println("加密解密成功")
    } else {
        fmt.Println("加密解密失败")
    }
    
    fmt.Println()
}

// AES加密
func encryptAES(plaintext, key []byte) ([]byte, error) {
    // 创建AES加密器
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    // 使用GCM模式
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    // 生成随机nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    // 加密
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}

// AES解密
func decryptAES(ciphertext, key []byte) ([]byte, error) {
    // 创建AES解密器
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    // 使用GCM模式
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    // 提取nonce
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("密文太短")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    
    // 解密
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    
    return plaintext, nil
}

// 非对称加密
func asymmetricEncryption() {
    fmt.Println("=== 非对称加密 (RSA) ===")
    
    // 生成RSA密钥对
    privateKey, publicKey, err := generateRSAKeyPair(2048)
    if err != nil {
        fmt.Printf("生成密钥对失败: %v\n", err)
        return
    }
    
    fmt.Println("RSA密钥对生成成功")
    
    // 准备数据
    plaintext := "RSA加密测试数据"
    fmt.Printf("原始数据: %s\n", plaintext)
    
    // 使用公钥加密
    ciphertext, err := encryptRSA([]byte(plaintext), publicKey)
    if err != nil {
        fmt.Printf("RSA加密失败: %v\n", err)
        return
    }
    
    fmt.Printf("加密数据 (Base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))
    
    // 使用私钥解密
    decrypted, err := decryptRSA(ciphertext, privateKey)
    if err != nil {
        fmt.Printf("RSA解密失败: %v\n", err)
        return
    }
    
    fmt.Printf("解密数据: %s\n", string(decrypted))
    
    // 验证
    if string(decrypted) == plaintext {
        fmt.Println("RSA加密解密成功")
    } else {
        fmt.Println("RSA加密解密失败")
    }
    
    // 导出密钥
    fmt.Println("\n密钥导出:")
    privateKeyPEM, publicKeyPEM := exportKeys(privateKey, publicKey)
    fmt.Printf("私钥 (PEM):\n%s\n", privateKeyPEM)
    fmt.Printf("公钥 (PEM):\n%s\n", publicKeyPEM)
    
    fmt.Println()
}

// 生成RSA密钥对
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, nil, err
    }
    
    return privateKey, &privateKey.PublicKey, nil
}

// RSA加密
func encryptRSA(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
    return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
}

// RSA解密
func decryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
}

// 导出密钥
func exportKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, string) {
    // 导出私钥
    privateKeyBytes, _ := x509.MarshalPKCS8PrivateKey(privateKey)
    privateKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "PRIVATE KEY",
        Bytes: privateKeyBytes,
    })
    
    // 导出公钥
    publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: publicKeyBytes,
    })
    
    return string(privateKeyPEM), string(publicKeyPEM)
}

// 数字签名
func digitalSignature() {
    fmt.Println("=== 数字签名 ===")
    
    // 生成密钥对
    privateKey, publicKey, err := generateRSAKeyPair(2048)
    if err != nil {
        fmt.Printf("生成密钥对失败: %v\n", err)
        return
    }
    
    // 准备数据
    message := "这是需要签名的重要文档内容"
    fmt.Printf("原始消息: %s\n", message)
    
    // 创建数字签名
    signature, err := createSignature([]byte(message), privateKey)
    if err != nil {
        fmt.Printf("创建签名失败: %v\n", err)
        return
    }
    
    fmt.Printf("数字签名 (Base64): %s\n", base64.StdEncoding.EncodeToString(signature))
    
    // 验证签名
    valid, err := verifySignature([]byte(message), signature, publicKey)
    if err != nil {
        fmt.Printf("验证签名失败: %v\n", err)
        return
    }
    
    if valid {
        fmt.Println("数字签名验证成功")
    } else {
        fmt.Println("数字签名验证失败")
    }
    
    // 测试篡改数据
    tamperedMessage := "这是被篡改的重要文档内容"
    fmt.Printf("\n篡改消息: %s\n", tamperedMessage)
    
    validTampered, err := verifySignature([]byte(tamperedMessage), signature, publicKey)
    if err != nil {
        fmt.Printf("验证篡改消息失败: %v\n", err)
        return
    }
    
    if validTampered {
        fmt.Println("篡改消息验证通过（不应该发生）")
    } else {
        fmt.Println("篡改消息验证失败（正确）")
    }
    
    fmt.Println()
}

// 创建数字签名
func createSignature(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    // 计算消息哈希
    hash := sha256.Sum256(message)
    
    // 使用私钥签名
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
    if err != nil {
        return nil, err
    }
    
    return signature, nil
}

// 验证数字签名
func verifySignature(message, signature []byte, publicKey *rsa.PublicKey) (bool, error) {
    // 计算消息哈希
    hash := sha256.Sum256(message)
    
    // 使用公钥验证签名
    err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
    if err != nil {
        return false, nil // 验证失败不是错误，只是签名无效
    }
    
    return true, nil
}
```

##  本章小结

在这一章中，我们学习了：

### 时间处理
- 时间创建、获取和格式化
- 时间计算和比较操作
- 时区处理和转换
- 定时器和周期任务

### 加密安全
- 哈希算法应用和比较
- 对称加密（AES）实现
- 非对称加密（RSA）应用
- 数字签名和验证

### 实用技巧
- 超时控制和任务调度
- 密码安全存储
- 数据完整性验证
- 密钥管理和导出

### 最佳实践
- 时间处理注意事项
- 加密算法选择
- 安全编程准则
- 性能优化策略

## 下一步

完成了进阶内容的学习后，让我们继续进入 [实战项目](../projects/)，通过实际项目来巩固所学知识！ 