---
title: 时间处理与加密
description: 学习Go语言的时间操作、日期处理和加密安全编程
---

# 时间处理与加密

时间处理和加密是现代应用程序的重要组成部分。Go语言提供了强大的时间处理能力和安全的加密库，让时间操作和安全编程变得简单可靠。

## 本章内容

- 时间基础操作和格式化
- 时区处理和时间计算
- 定时器和周期性任务
- 对称加密和非对称加密
- 密码学哈希和数字签名

## 时间处理概念

### Go语言时间特性

Go语言的time包具有以下特点：

- **精确时间**：纳秒级精度的时间表示
- **时区感知**：内置时区支持和转换
- **Duration类型**：直观的时间间隔表示
- **标准格式**：统一的时间格式化规则

### 时间处理优势

| 特性 | 说明 | 优势 |
|------|------|------|
| **精度高** | 纳秒级时间精度 | 适合高精度计时 |
| **时区完整** | 全球时区数据库 | 国际化应用支持 |
| **API简洁** | 直观的时间操作 | 易于理解和使用 |
| **性能好** | 高效的时间运算 | 适合高频操作 |

::: tip 设计原则
Go时间处理遵循"精确、简洁、可靠"的设计理念：
- 提供高精度的时间表示
- 保持API的一致性和直观性
- 正确处理时区和夏令时
:::

## 时间基础操作

### 时间创建和获取

```go
package main

import (
    "fmt"
    "time"
)

func timeBasics() {
    fmt.Println("⏰ 时间基础操作:")
    
    // 当前时间
    now := time.Now()
    fmt.Printf("当前时间: %s\n", now)
    fmt.Printf("时间戳: %d\n", now.Unix())
    fmt.Printf("纳秒时间戳: %d\n", now.UnixNano())
    
    // 创建特定时间
    birthday := time.Date(1990, 5, 15, 14, 30, 0, 0, time.UTC)
    fmt.Printf("指定时间: %s\n", birthday)
    
    // 从字符串解析时间
    layouts := map[string]string{
        "RFC3339":     "2006-01-02T15:04:05Z07:00",
        "日期时间":       "2006-01-02 15:04:05",
        "仅日期":        "2006-01-02",
        "仅时间":        "15:04:05",
        "中文格式":       "2006年01月02日 15时04分05秒",
    }
    
    timeStrings := map[string]string{
        "RFC3339":     "2023-12-25T15:30:45Z",
        "日期时间":       "2023-12-25 15:30:45",
        "仅日期":        "2023-12-25",
        "仅时间":        "15:30:45",
        "中文格式":       "2023年12月25日 15时30分45秒",
    }
    
    fmt.Println("\n📅 时间解析:")
    for name, layout := range layouts {
        if timeStr, exists := timeStrings[name]; exists {
            if parsedTime, err := time.Parse(layout, timeStr); err == nil {
                fmt.Printf("  %s: %s -> %s\n", name, timeStr, parsedTime.Format("2006-01-02 15:04:05"))
            } else {
                fmt.Printf("  %s: 解析失败 - %v\n", name, err)
            }
        }
    }
}

// 时间格式化
func timeFormatting() {
    fmt.Println("\n🎨 时间格式化:")
    
    now := time.Now()
    
    formats := map[string]string{
        "标准格式":       "2006-01-02 15:04:05",
        "ISO8601":      "2006-01-02T15:04:05Z07:00",
        "RFC822":       "02 Jan 06 15:04 MST",
        "日期":          "2006-01-02",
        "时间":          "15:04:05",
        "年月":          "2006-01",
        "中文长格式":       "2006年01月02日 星期Monday 15时04分05秒",
        "美式日期":        "01/02/2006",
        "12小时制":       "2006-01-02 03:04:05 PM",
        "时区显示":        "2006-01-02 15:04:05 MST",
    }
    
    for name, layout := range formats {
        formatted := now.Format(layout)
        fmt.Printf("  %s: %s\n", name, formatted)
    }
    
    // 自定义格式化函数
    fmt.Printf("\n📝 自定义格式:\n")
    fmt.Printf("  相对时间: %s\n", relativeTime(now))
    fmt.Printf("  友好格式: %s\n", friendlyTime(now))
}

// 相对时间显示
func relativeTime(t time.Time) string {
    now := time.Now()
    diff := now.Sub(t)
    
    if diff < 0 {
        diff = -diff
        if diff < time.Minute {
            return "即将到来"
        } else if diff < time.Hour {
            return fmt.Sprintf("%d分钟后", int(diff.Minutes()))
        } else if diff < 24*time.Hour {
            return fmt.Sprintf("%d小时后", int(diff.Hours()))
        } else {
            return fmt.Sprintf("%d天后", int(diff.Hours()/24))
        }
    }
    
    if diff < time.Minute {
        return "刚刚"
    } else if diff < time.Hour {
        return fmt.Sprintf("%d分钟前", int(diff.Minutes()))
    } else if diff < 24*time.Hour {
        return fmt.Sprintf("%d小时前", int(diff.Hours()))
    } else {
        return fmt.Sprintf("%d天前", int(diff.Hours()/24))
    }
}

// 友好时间格式
func friendlyTime(t time.Time) string {
    hour := t.Hour()
    
    switch {
    case hour < 6:
        return t.Format("凌晨 03:04")
    case hour < 12:
        return t.Format("上午 03:04")
    case hour < 18:
        return t.Format("下午 03:04")
    default:
        return t.Format("晚上 03:04")
    }
}
```

### 时间计算和比较

```go
func timeCalculations() {
    fmt.Println("🧮 时间计算:")
    
    now := time.Now()
    
    // 时间加减
    futureTime := now.Add(2 * time.Hour)
    pastTime := now.Add(-3 * time.Hour)
    
    fmt.Printf("当前时间: %s\n", now.Format("15:04:05"))
    fmt.Printf("2小时后: %s\n", futureTime.Format("15:04:05"))
    fmt.Printf("3小时前: %s\n", pastTime.Format("15:04:05"))
    
    // 时间差计算
    duration := futureTime.Sub(pastTime)
    fmt.Printf("时间差: %v (%v小时)\n", duration, duration.Hours())
    
    // 时间比较
    fmt.Println("\n⚖️ 时间比较:")
    fmt.Printf("现在是否在过去时间之后: %t\n", now.After(pastTime))
    fmt.Printf("现在是否在未来时间之前: %t\n", now.Before(futureTime))
    fmt.Printf("两个时间是否相等: %t\n", now.Equal(now))
    
    // 工作日计算
    workdays := calculateWorkdays(now, futureTime.AddDate(0, 0, 10))
    fmt.Printf("未来10天内工作日: %d天\n", workdays)
    
    // 年龄计算
    birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
    age := calculateAge(birthday, now)
    fmt.Printf("年龄计算: %d岁\n", age)
}

// 计算工作日
func calculateWorkdays(start, end time.Time) int {
    workdays := 0
    current := start
    
    for current.Before(end) || current.Equal(end) {
        weekday := current.Weekday()
        if weekday != time.Saturday && weekday != time.Sunday {
            workdays++
        }
        current = current.AddDate(0, 0, 1)
    }
    
    return workdays
}

// 计算年龄
func calculateAge(birthday, now time.Time) int {
    age := now.Year() - birthday.Year()
    
    // 检查是否还没到生日
    if now.Month() < birthday.Month() || 
       (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
        age--
    }
    
    return age
}
```

### 时区处理

```go
func timezoneHandling() {
    fmt.Println("🌍 时区处理:")
    
    now := time.Now()
    
    // 常用时区
    timezones := map[string]string{
        "北京":    "Asia/Shanghai",
        "东京":    "Asia/Tokyo",
        "纽约":    "America/New_York",
        "伦敦":    "Europe/London",
        "悉尼":    "Australia/Sydney",
        "莫斯科":   "Europe/Moscow",
    }
    
    fmt.Printf("当前时间 (本地): %s\n", now.Format("2006-01-02 15:04:05 MST"))
    
    for city, tzName := range timezones {
        if location, err := time.LoadLocation(tzName); err == nil {
            cityTime := now.In(location)
            fmt.Printf("%s时间: %s\n", city, cityTime.Format("2006-01-02 15:04:05 MST"))
        }
    }
    
    // UTC时间
    utcTime := now.UTC()
    fmt.Printf("UTC时间: %s\n", utcTime.Format("2006-01-02 15:04:05 MST"))
    
    // 时区转换
    fmt.Println("\n🔄 时区转换示例:")
    beijing, _ := time.LoadLocation("Asia/Shanghai")
    newYork, _ := time.LoadLocation("America/New_York")
    
    beijingTime := time.Date(2023, 12, 25, 14, 30, 0, 0, beijing)
    newYorkTime := beijingTime.In(newYork)
    
    fmt.Printf("北京时间: %s\n", beijingTime.Format("2006-01-02 15:04:05 MST"))
    fmt.Printf("对应纽约时间: %s\n", newYorkTime.Format("2006-01-02 15:04:05 MST"))
    fmt.Printf("时差: %.1f小时\n", beijingTime.Sub(newYorkTime).Hours())
}
```

## 定时器和周期任务

### 定时器使用

```go
func timerOperations() {
    fmt.Println("⏲️ 定时器操作:")
    
    // 单次定时器
    fmt.Println("3秒后执行...")
    timer := time.NewTimer(3 * time.Second)
    
    go func() {
        <-timer.C
        fmt.Println("✅ 定时器触发!")
    }()
    
    // 等待定时器完成
    time.Sleep(4 * time.Second)
    
    // 周期性定时器
    fmt.Println("\n📅 周期性定时器 (每2秒):")
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    
    count := 0
    for range ticker.C {
        count++
        fmt.Printf("  第%d次触发: %s\n", count, time.Now().Format("15:04:05"))
        
        if count >= 3 {
            break
        }
    }
    
    // 使用context控制定时器
    fmt.Println("\n🛑 可控制的定时器:")
    controllableTimer()
}

func controllableTimer() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            fmt.Printf("  心跳: %s\n", time.Now().Format("15:04:05"))
        case <-ctx.Done():
            fmt.Println("  定时器停止")
            return
        }
    }
}
```

### 任务调度器

```go
import (
    "context"
    "sync"
    "time"
)

// 任务调度器
type TaskScheduler struct {
    tasks   map[string]*ScheduledTask
    mutex   sync.RWMutex
    ctx     context.Context
    cancel  context.CancelFunc
    running bool
}

type ScheduledTask struct {
    ID       string
    Name     string
    Function func()
    Interval time.Duration
    NextRun  time.Time
    Enabled  bool
    RunCount int
}

func NewTaskScheduler() *TaskScheduler {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &TaskScheduler{
        tasks:  make(map[string]*ScheduledTask),
        ctx:    ctx,
        cancel: cancel,
    }
}

// 添加任务
func (ts *TaskScheduler) AddTask(id, name string, fn func(), interval time.Duration) {
    ts.mutex.Lock()
    defer ts.mutex.Unlock()
    
    task := &ScheduledTask{
        ID:       id,
        Name:     name,
        Function: fn,
        Interval: interval,
        NextRun:  time.Now().Add(interval),
        Enabled:  true,
        RunCount: 0,
    }
    
    ts.tasks[id] = task
    fmt.Printf("✅ 任务已添加: %s (间隔: %v)\n", name, interval)
}

// 启动调度器
func (ts *TaskScheduler) Start() {
    if ts.running {
        return
    }
    
    ts.running = true
    fmt.Println("🚀 任务调度器启动")
    
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            ts.checkAndRunTasks()
        case <-ts.ctx.Done():
            fmt.Println("🛑 任务调度器停止")
            return
        }
    }
}

// 检查并运行任务
func (ts *TaskScheduler) checkAndRunTasks() {
    ts.mutex.Lock()
    defer ts.mutex.Unlock()
    
    now := time.Now()
    
    for _, task := range ts.tasks {
        if task.Enabled && now.After(task.NextRun) {
            go func(t *ScheduledTask) {
                fmt.Printf("⚡ 执行任务: %s\n", t.Name)
                t.Function()
                t.RunCount++
                t.NextRun = now.Add(t.Interval)
            }(task)
        }
    }
}

// 停止调度器
func (ts *TaskScheduler) Stop() {
    ts.cancel()
    ts.running = false
}

// 获取任务状态
func (ts *TaskScheduler) GetTaskStatus() {
    ts.mutex.RLock()
    defer ts.mutex.RUnlock()
    
    fmt.Println("\n📊 任务状态:")
    for _, task := range ts.tasks {
        status := "禁用"
        if task.Enabled {
            status = "启用"
        }
        
        nextRun := "立即"
        if time.Now().Before(task.NextRun) {
            nextRun = relativeTime(task.NextRun)
        }
        
        fmt.Printf("  %s: %s | 运行次数: %d | 下次运行: %s\n", 
            task.Name, status, task.RunCount, nextRun)
    }
}

// 演示调度器
func demonstrateScheduler() {
    scheduler := NewTaskScheduler()
    
    // 添加定时任务
    scheduler.AddTask("heartbeat", "心跳检测", func() {
        fmt.Printf("💓 心跳 - %s\n", time.Now().Format("15:04:05"))
    }, 3*time.Second)
    
    scheduler.AddTask("backup", "数据备份", func() {
        fmt.Printf("💾 执行备份 - %s\n", time.Now().Format("15:04:05"))
    }, 10*time.Second)
    
    scheduler.AddTask("cleanup", "清理日志", func() {
        fmt.Printf("🧹 清理日志 - %s\n", time.Now().Format("15:04:05"))
    }, 15*time.Second)
    
    // 启动调度器
    go scheduler.Start()
    
    // 运行一段时间
    time.Sleep(20 * time.Second)
    
    // 显示状态
    scheduler.GetTaskStatus()
    
    // 停止调度器
    scheduler.Stop()
}
```

## 加密基础

### 对称加密

```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

// AES加密工具
type AESCrypto struct {
    key []byte
}

func NewAESCrypto(key string) (*AESCrypto, error) {
    // 确保密钥长度正确 (16, 24, 或 32 字节)
    keyBytes := []byte(key)
    if len(keyBytes) < 32 {
        // 填充密钥到32字节
        paddedKey := make([]byte, 32)
        copy(paddedKey, keyBytes)
        keyBytes = paddedKey
    } else if len(keyBytes) > 32 {
        keyBytes = keyBytes[:32]
    }
    
    return &AESCrypto{key: keyBytes}, nil
}

// AES-GCM加密
func (ac *AESCrypto) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(ac.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AES-GCM解密
func (ac *AESCrypto) Decrypt(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    
    block, err := aes.NewCipher(ac.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    if len(data) < gcm.NonceSize() {
        return "", fmt.Errorf("密文太短")
    }
    
    nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    
    return string(plaintext), nil
}

func symmetricEncryptionDemo() {
    fmt.Println("🔐 对称加密演示:")
    
    crypto, err := NewAESCrypto("my-secret-key-for-encryption")
    if err != nil {
        fmt.Printf("创建加密器失败: %v\n", err)
        return
    }
    
    // 测试数据
    testData := []string{
        "Hello, World!",
        "这是中文测试数据",
        "Special chars: !@#$%^&*()",
        "JSON数据: {\"name\":\"张三\",\"age\":25}",
    }
    
    for _, plaintext := range testData {
        fmt.Printf("\n原文: %s\n", plaintext)
        
        // 加密
        encrypted, err := crypto.Encrypt(plaintext)
        if err != nil {
            fmt.Printf("加密失败: %v\n", err)
            continue
        }
        fmt.Printf("密文: %s\n", encrypted)
        
        // 解密
        decrypted, err := crypto.Decrypt(encrypted)
        if err != nil {
            fmt.Printf("解密失败: %v\n", err)
            continue
        }
        fmt.Printf("解密: %s\n", decrypted)
        
        // 验证
        if plaintext == decrypted {
            fmt.Println("✅ 加解密验证成功")
        } else {
            fmt.Println("❌ 加解密验证失败")
        }
    }
}
```

### 哈希和数字签名

```go
import (
    "crypto/md5"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// 哈希工具
type HashUtils struct{}

// MD5哈希 (仅用于非安全场合)
func (hu *HashUtils) MD5(data string) string {
    hash := md5.Sum([]byte(data))
    return fmt.Sprintf("%x", hash)
}

// SHA256哈希
func (hu *HashUtils) SHA256(data string) string {
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}

// SHA512哈希
func (hu *HashUtils) SHA512(data string) string {
    hash := sha512.Sum512([]byte(data))
    return fmt.Sprintf("%x", hash)
}

// 密码哈希 (使用bcrypt)
func (hu *HashUtils) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// 验证密码
func (hu *HashUtils) CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func hashingDemo() {
    fmt.Println("🔒 哈希算法演示:")
    
    hashUtils := &HashUtils{}
    testData := "Hello, Go Crypto!"
    
    fmt.Printf("原始数据: %s\n", testData)
    fmt.Printf("MD5:      %s\n", hashUtils.MD5(testData))
    fmt.Printf("SHA256:   %s\n", hashUtils.SHA256(testData))
    fmt.Printf("SHA512:   %s\n", hashUtils.SHA512(testData))
    
    // 密码哈希演示
    fmt.Println("\n🔑 密码哈希演示:")
    passwords := []string{"password123", "mySecretPassword", "中文密码"}
    
    for _, pwd := range passwords {
        hashedPwd, err := hashUtils.HashPassword(pwd)
        if err != nil {
            fmt.Printf("密码哈希失败: %v\n", err)
            continue
        }
        
        fmt.Printf("\n原始密码: %s\n", pwd)
        fmt.Printf("哈希结果: %s\n", hashedPwd)
        
        // 验证密码
        isValid := hashUtils.CheckPassword(pwd, hashedPwd)
        fmt.Printf("验证结果: %t\n", isValid)
        
        // 验证错误密码
        isWrong := hashUtils.CheckPassword("wrongpassword", hashedPwd)
        fmt.Printf("错误密码: %t\n", isWrong)
    }
}
```

## 实战项目：安全令牌系统

让我们构建一个完整的安全令牌（JWT）系统：

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "strings"
    "time"
)

// JWT头部
type JWTHeader struct {
    Algorithm string `json:"alg"`
    Type      string `json:"typ"`
}

// JWT负载
type JWTPayload struct {
    Subject   string `json:"sub"`           // 主题
    Issuer    string `json:"iss"`           // 签发者
    Audience  string `json:"aud"`           // 受众
    ExpiresAt int64  `json:"exp"`           // 过期时间
    NotBefore int64  `json:"nbf"`           // 生效时间
    IssuedAt  int64  `json:"iat"`           // 签发时间
    JWTID     string `json:"jti,omitempty"` // JWT ID
    
    // 自定义声明
    Username string   `json:"username,omitempty"`
    Roles    []string `json:"roles,omitempty"`
    Email    string   `json:"email,omitempty"`
}

// JWT令牌管理器
type JWTManager struct {
    secretKey []byte
    issuer    string
}

func NewJWTManager(secretKey, issuer string) *JWTManager {
    return &JWTManager{
        secretKey: []byte(secretKey),
        issuer:    issuer,
    }
}

// 生成JWT令牌
func (jm *JWTManager) GenerateToken(username, email string, roles []string, expiration time.Duration) (string, error) {
    now := time.Now()
    
    // 创建头部
    header := JWTHeader{
        Algorithm: "HS256",
        Type:      "JWT",
    }
    
    // 创建负载
    payload := JWTPayload{
        Subject:   username,
        Issuer:    jm.issuer,
        Audience:  "api-users",
        ExpiresAt: now.Add(expiration).Unix(),
        NotBefore: now.Unix(),
        IssuedAt:  now.Unix(),
        JWTID:     generateJTI(),
        Username:  username,
        Email:     email,
        Roles:     roles,
    }
    
    // 编码头部和负载
    encodedHeader, err := jm.encodeSegment(header)
    if err != nil {
        return "", err
    }
    
    encodedPayload, err := jm.encodeSegment(payload)
    if err != nil {
        return "", err
    }
    
    // 创建签名
    message := encodedHeader + "." + encodedPayload
    signature := jm.createSignature(message)
    
    // 组合完整令牌
    token := message + "." + signature
    return token, nil
}

// 验证JWT令牌
func (jm *JWTManager) ValidateToken(tokenString string) (*JWTPayload, error) {
    parts := strings.Split(tokenString, ".")
    if len(parts) != 3 {
        return nil, fmt.Errorf("无效的JWT格式")
    }
    
    // 验证签名
    message := parts[0] + "." + parts[1]
    expectedSignature := jm.createSignature(message)
    
    if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
        return nil, fmt.Errorf("签名验证失败")
    }
    
    // 解析负载
    payload, err := jm.decodePayload(parts[1])
    if err != nil {
        return nil, err
    }
    
    // 检查时间有效性
    now := time.Now().Unix()
    
    if payload.ExpiresAt < now {
        return nil, fmt.Errorf("令牌已过期")
    }
    
    if payload.NotBefore > now {
        return nil, fmt.Errorf("令牌尚未生效")
    }
    
    return payload, nil
}

// 编码段
func (jm *JWTManager) encodeSegment(data interface{}) (string, error) {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    
    return base64.RawURLEncoding.EncodeToString(jsonBytes), nil
}

// 解析负载
func (jm *JWTManager) decodePayload(segment string) (*JWTPayload, error) {
    jsonBytes, err := base64.RawURLEncoding.DecodeString(segment)
    if err != nil {
        return nil, err
    }
    
    var payload JWTPayload
    if err := json.Unmarshal(jsonBytes, &payload); err != nil {
        return nil, err
    }
    
    return &payload, nil
}

// 创建签名
func (jm *JWTManager) createSignature(message string) string {
    mac := hmac.New(sha256.New, jm.secretKey)
    mac.Write([]byte(message))
    return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// 刷新令牌
func (jm *JWTManager) RefreshToken(tokenString string) (string, error) {
    payload, err := jm.ValidateToken(tokenString)
    if err != nil {
        return "", err
    }
    
    // 检查是否可以刷新 (距离过期还有一定时间)
    now := time.Now().Unix()
    if payload.ExpiresAt-now > 300 { // 5分钟
        return "", fmt.Errorf("令牌刷新过早")
    }
    
    // 生成新令牌
    return jm.GenerateToken(payload.Username, payload.Email, payload.Roles, time.Hour)
}

// 用户认证系统
type AuthSystem struct {
    jwtManager *JWTManager
    users      map[string]*User // 简单的内存用户存储
}

type User struct {
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash"`
    Roles        []string  `json:"roles"`
    CreatedAt    time.Time `json:"created_at"`
    LastLogin    time.Time `json:"last_login"`
    IsActive     bool      `json:"is_active"`
}

func NewAuthSystem(secretKey string) *AuthSystem {
    return &AuthSystem{
        jwtManager: NewJWTManager(secretKey, "go-auth-system"),
        users:      make(map[string]*User),
    }
}

// 注册用户
func (as *AuthSystem) RegisterUser(username, email, password string, roles []string) error {
    if _, exists := as.users[username]; exists {
        return fmt.Errorf("用户已存在")
    }
    
    // 哈希密码
    hashUtils := &HashUtils{}
    passwordHash, err := hashUtils.HashPassword(password)
    if err != nil {
        return err
    }
    
    user := &User{
        Username:     username,
        Email:        email,
        PasswordHash: passwordHash,
        Roles:        roles,
        CreatedAt:    time.Now(),
        IsActive:     true,
    }
    
    as.users[username] = user
    fmt.Printf("✅ 用户注册成功: %s\n", username)
    return nil
}

// 用户登录
func (as *AuthSystem) Login(username, password string) (string, error) {
    user, exists := as.users[username]
    if !exists {
        return "", fmt.Errorf("用户不存在")
    }
    
    if !user.IsActive {
        return "", fmt.Errorf("用户已被禁用")
    }
    
    // 验证密码
    hashUtils := &HashUtils{}
    if !hashUtils.CheckPassword(password, user.PasswordHash) {
        return "", fmt.Errorf("密码错误")
    }
    
    // 更新最后登录时间
    user.LastLogin = time.Now()
    
    // 生成JWT令牌
    token, err := as.jwtManager.GenerateToken(user.Username, user.Email, user.Roles, time.Hour*24)
    if err != nil {
        return "", err
    }
    
    fmt.Printf("✅ 用户登录成功: %s\n", username)
    return token, nil
}

// 验证令牌
func (as *AuthSystem) ValidateToken(tokenString string) (*JWTPayload, error) {
    return as.jwtManager.ValidateToken(tokenString)
}

// 检查权限
func (as *AuthSystem) CheckPermission(tokenString string, requiredRole string) bool {
    payload, err := as.ValidateToken(tokenString)
    if err != nil {
        return false
    }
    
    for _, role := range payload.Roles {
        if role == requiredRole || role == "admin" {
            return true
        }
    }
    
    return false
}

// 演示认证系统
func demonstrateAuthSystem() {
    fmt.Println("🔐 认证系统演示")
    fmt.Println("==============")
    
    // 创建认证系统
    authSystem := NewAuthSystem("my-super-secret-jwt-key")
    
    // 注册用户
    users := []struct {
        username string
        email    string
        password string
        roles    []string
    }{
        {"admin", "admin@example.com", "admin123", []string{"admin", "user"}},
        {"alice", "alice@example.com", "alice123", []string{"user", "editor"}},
        {"bob", "bob@example.com", "bob123", []string{"user"}},
    }
    
    fmt.Println("👥 注册用户:")
    for _, u := range users {
        authSystem.RegisterUser(u.username, u.email, u.password, u.roles)
    }
    
    // 登录演示
    fmt.Println("\n🔑 登录演示:")
    token, err := authSystem.Login("alice", "alice123")
    if err != nil {
        fmt.Printf("登录失败: %v\n", err)
        return
    }
    
    fmt.Printf("登录令牌: %s...\n", token[:50])
    
    // 令牌验证
    fmt.Println("\n✅ 令牌验证:")
    payload, err := authSystem.ValidateToken(token)
    if err != nil {
        fmt.Printf("令牌验证失败: %v\n", err)
        return
    }
    
    fmt.Printf("用户: %s\n", payload.Username)
    fmt.Printf("邮箱: %s\n", payload.Email)
    fmt.Printf("角色: %v\n", payload.Roles)
    fmt.Printf("过期时间: %s\n", time.Unix(payload.ExpiresAt, 0).Format("2006-01-02 15:04:05"))
    
    // 权限检查
    fmt.Println("\n🛡️ 权限检查:")
    permissions := []string{"user", "editor", "admin"}
    
    for _, perm := range permissions {
        hasPermission := authSystem.CheckPermission(token, perm)
        status := "❌"
        if hasPermission {
            status = "✅"
        }
        fmt.Printf("  %s 权限: %s\n", perm, status)
    }
    
    // 错误登录演示
    fmt.Println("\n❌ 错误登录演示:")
    if _, err := authSystem.Login("alice", "wrongpassword"); err != nil {
        fmt.Printf("预期错误: %v\n", err)
    }
    
    if _, err := authSystem.Login("nonexistent", "password"); err != nil {
        fmt.Printf("预期错误: %v\n", err)
    }
}

// 工具函数
func generateJTI() string {
    return fmt.Sprintf("jwt_%d", time.Now().UnixNano())
}

func main() {
    // 时间处理演示
    timeBasics()
    timeFormatting()
    timeCalculations()
    timezoneHandling()
    
    // 定时器演示
    timerOperations()
    demonstrateScheduler()
    
    // 加密演示
    symmetricEncryptionDemo()
    hashingDemo()
    
    // 认证系统演示
    demonstrateAuthSystem()
}
```

## 最佳实践

### 1. 时间处理最佳实践

```go
// 总是使用UTC进行存储和计算
func storeTime() time.Time {
    return time.Now().UTC()
}

// 显示时根据用户时区转换
func displayTimeForUser(t time.Time, userTimezone string) string {
    location, _ := time.LoadLocation(userTimezone)
    return t.In(location).Format("2006-01-02 15:04:05")
}

// 使用常量定义时间间隔
const (
    DefaultTimeout = 30 * time.Second
    CacheExpiry    = 5 * time.Minute
    SessionTimeout = 24 * time.Hour
)
```

### 2. 加密安全最佳实践

```go
// 使用强随机密钥
func generateSecureKey() ([]byte, error) {
    key := make([]byte, 32) // 256位密钥
    _, err := rand.Read(key)
    return key, err
}

// 安全比较哈希值
func secureCompare(a, b []byte) bool {
    return hmac.Equal(a, b)
}

// 密钥派生
func deriveKey(password, salt []byte) []byte {
    return pbkdf2.Key(password, salt, 10000, 32, sha256.New)
}
```

### 3. 令牌管理最佳实践

```go
// 短期访问令牌 + 长期刷新令牌
type TokenPair struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    int    `json:"expires_in"`
}

// 安全的令牌存储
type SecureTokenStore struct {
    tokens map[string]time.Time // token -> expiry
    mutex  sync.RWMutex
}

func (sts *SecureTokenStore) IsTokenValid(token string) bool {
    sts.mutex.RLock()
    defer sts.mutex.RUnlock()
    
    expiry, exists := sts.tokens[token]
    return exists && time.Now().Before(expiry)
}
```

## 本章小结

Go语言时间处理和加密的核心要点：

- **时间操作**：掌握time包的时间创建、格式化、计算和比较
- **时区处理**：正确处理时区转换和国际化时间显示
- **定时任务**：使用Timer和Ticker实现任务调度
- **加密安全**：掌握对称加密、哈希算法和数字签名
- **令牌系统**：实现安全的JWT认证和授权系统

### 下一步
完成进阶内容学习后，可以开始学习 [实战项目](../projects/)，将所学知识应用到具体项目中。

::: tip 练习建议
1. 实现一个完整的用户认证系统
2. 开发定时任务调度器
3. 创建文件加密工具
4. 构建日志分析和监控系统
:::