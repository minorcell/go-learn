---
title: 文件操作与I/O
description: 学习Go语言的文件系统操作、数据格式处理和I/O操作
---

# 文件操作与I/O

文件操作是所有编程语言的基础技能。Go语言提供了丰富的标准库来处理文件系统操作、数据格式解析和I/O流处理，让文件操作变得简单高效。

## 本章内容

- 文件基础操作和路径处理
- 文件内容读写和流式处理
- JSON/CSV/XML等数据格式处理
- 目录遍历和文件系统操作
- 配置文件管理和日志系统

## 文件操作概念

### Go语言I/O体系

Go的I/O系统基于接口设计，核心是`io.Reader`和`io.Writer`：

- **Reader接口**：从数据源读取数据的通用接口
- **Writer接口**：向数据目标写入数据的通用接口
- **组合接口**：ReadWriter、ReadCloser等组合功能
- **缓冲I/O**：bufio包提供缓冲读写功能

### 文件操作优势

| 特性 | 说明 | 优势 |
|------|------|------|
| **接口统一** | 统一的Reader/Writer接口 | 代码复用性高 |
| **错误处理** | 显式错误返回 | 错误处理清晰 |
| **性能优化** | 支持缓冲和并发 | 高效处理大文件 |
| **跨平台** | 统一的文件路径API | 跨平台兼容性好 |

::: tip 设计原则
Go文件操作遵循"简单、显式、高效"的设计理念：
- 使用接口抽象I/O操作
- 显式处理错误和资源管理
- 支持流式处理大文件
:::

## 文件基础操作

### 文件读写基础

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

// 基础文件操作
func basicFileOperations() {
    // 写入文件
    content := "Hello, Go 文件操作!\n学习Go语言文件处理。"
    
    err := os.WriteFile("example.txt", []byte(content), 0644)
    if err != nil {
        fmt.Printf("写入文件失败: %v\n", err)
        return
    }
    fmt.Println("✅ 文件写入成功")
    
    // 读取文件
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("读取文件失败: %v\n", err)
        return
    }
    fmt.Printf("📄 文件内容:\n%s\n", string(data))
    
    // 检查文件是否存在
    if _, err := os.Stat("example.txt"); err == nil {
        fmt.Println("✅ 文件存在")
    } else if os.IsNotExist(err) {
        fmt.Println("❌ 文件不存在")
    }
}

// 使用File对象操作
func fileObjectOperations() {
    // 创建文件
    file, err := os.Create("advanced_example.txt")
    if err != nil {
        fmt.Printf("创建文件失败: %v\n", err)
        return
    }
    defer file.Close()
    
    // 写入多行数据
    lines := []string{
        "第一行数据",
        "第二行数据", 
        "第三行数据",
    }
    
    for i, line := range lines {
        _, err := file.WriteString(fmt.Sprintf("%d: %s\n", i+1, line))
        if err != nil {
            fmt.Printf("写入失败: %v\n", err)
            return
        }
    }
    
    fmt.Println("✅ 高级文件操作完成")
}

// 流式读取处理大文件
func streamReading() {
    file, err := os.Open("advanced_example.txt")
    if err != nil {
        fmt.Printf("打开文件失败: %v\n", err)
        return
    }
    defer file.Close()
    
    // 逐行读取
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("读取失败: %v\n", err)
        return
    }
    
    lines := strings.Split(string(content), "\n")
    fmt.Println("📖 逐行读取结果:")
    for _, line := range lines {
        if line != "" {
            fmt.Printf("  %s\n", line)
        }
    }
}
```

### 缓冲I/O操作

使用bufio包提升大文件处理性能：

```go
import (
    "bufio"
    "fmt"
    "os"
)

func bufferedFileOperations() {
    // 缓冲写入
    file, err := os.Create("buffered_output.txt")
    if err != nil {
        fmt.Printf("创建文件失败: %v\n", err)
        return
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    defer writer.Flush() // 确保缓冲区内容写入文件
    
    // 写入大量数据
    for i := 1; i <= 1000; i++ {
        _, err := writer.WriteString(fmt.Sprintf("行 %d: 这是测试数据\n", i))
        if err != nil {
            fmt.Printf("写入失败: %v\n", err)
            return
        }
    }
    
    fmt.Println("✅ 缓冲写入完成")
    
    // 缓冲读取
    readFile, err := os.Open("buffered_output.txt")
    if err != nil {
        fmt.Printf("打开文件失败: %v\n", err)
        return
    }
    defer readFile.Close()
    
    scanner := bufio.NewScanner(readFile)
    lineCount := 0
    
    for scanner.Scan() {
        lineCount++
        // 只显示前5行和后5行
        if lineCount <= 5 || lineCount > 995 {
            fmt.Printf("第%d行: %s\n", lineCount, scanner.Text())
        } else if lineCount == 6 {
            fmt.Println("... (省略中间行) ...")
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("扫描文件失败: %v\n", err)
        return
    }
    
    fmt.Printf("✅ 总共读取 %d 行\n", lineCount)
}
```

## 路径和目录操作

### 路径处理

```go
import (
    "fmt"
    "path/filepath"
    "os"
)

func pathOperations() {
    // 路径拼接（跨平台）
    path := filepath.Join("data", "users", "profile.json")
    fmt.Printf("拼接路径: %s\n", path)
    
    // 获取路径信息
    dir := filepath.Dir(path)
    base := filepath.Base(path)
    ext := filepath.Ext(path)
    
    fmt.Printf("目录: %s\n", dir)
    fmt.Printf("文件名: %s\n", base)
    fmt.Printf("扩展名: %s\n", ext)
    
    // 绝对路径
    abs, err := filepath.Abs(path)
    if err == nil {
        fmt.Printf("绝对路径: %s\n", abs)
    }
    
    // 清理路径
    cleanPath := filepath.Clean("./data//users/../users/./profile.json")
    fmt.Printf("清理后路径: %s\n", cleanPath)
}

func directoryOperations() {
    // 创建目录
    err := os.MkdirAll("data/users/temp", 0755)
    if err != nil {
        fmt.Printf("创建目录失败: %v\n", err)
        return
    }
    fmt.Println("✅ 目录创建成功")
    
    // 列出目录内容
    entries, err := os.ReadDir("data")
    if err != nil {
        fmt.Printf("读取目录失败: %v\n", err)
        return
    }
    
    fmt.Println("📁 目录内容:")
    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Printf("  📁 %s/\n", entry.Name())
        } else {
            fmt.Printf("  📄 %s\n", entry.Name())
        }
    }
    
    // 删除目录
    err = os.RemoveAll("data/users/temp")
    if err != nil {
        fmt.Printf("删除目录失败: %v\n", err)
    } else {
        fmt.Println("✅ 临时目录已删除")
    }
}
```

### 文件遍历

```go
func walkDirectory() {
    fmt.Println("🚶 遍历当前目录:")
    
    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        // 跳过隐藏文件和目录
        if strings.HasPrefix(info.Name(), ".") {
            if info.IsDir() {
                return filepath.SkipDir
            }
            return nil
        }
        
        if info.IsDir() {
            fmt.Printf("📁 %s/\n", path)
        } else {
            size := info.Size()
            modTime := info.ModTime().Format("2006-01-02 15:04:05")
            fmt.Printf("📄 %s (大小: %d字节, 修改时间: %s)\n", path, size, modTime)
        }
        
        return nil
    })
    
    if err != nil {
        fmt.Printf("遍历目录失败: %v\n", err)
    }
}
```

## 数据格式处理

### JSON处理

```go
import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
    Tags      []string  `json:"tags"`
}

func jsonOperations() {
    // 创建示例数据
    users := []User{
        {
            ID:        1,
            Name:      "张三",
            Email:     "zhangsan@example.com",
            Age:       25,
            IsActive:  true,
            CreatedAt: time.Now(),
            Tags:      []string{"开发者", "Go语言"},
        },
        {
            ID:        2,
            Name:      "李四",
            Email:     "lisi@example.com",
            Age:       30,
            IsActive:  false,
            CreatedAt: time.Now().Add(-24 * time.Hour),
            Tags:      []string{"设计师", "UI/UX"},
        },
    }
    
    // JSON编码并写入文件
    file, err := os.Create("users.json")
    if err != nil {
        fmt.Printf("创建JSON文件失败: %v\n", err)
        return
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") // 格式化输出
    
    err = encoder.Encode(users)
    if err != nil {
        fmt.Printf("JSON编码失败: %v\n", err)
        return
    }
    
    fmt.Println("✅ JSON文件写入成功")
    
    // 从文件读取JSON
    readFile, err := os.Open("users.json")
    if err != nil {
        fmt.Printf("打开JSON文件失败: %v\n", err)
        return
    }
    defer readFile.Close()
    
    var loadedUsers []User
    decoder := json.NewDecoder(readFile)
    
    err = decoder.Decode(&loadedUsers)
    if err != nil {
        fmt.Printf("JSON解码失败: %v\n", err)
        return
    }
    
    fmt.Printf("📄 读取到 %d 个用户:\n", len(loadedUsers))
    for _, user := range loadedUsers {
        fmt.Printf("  - %s (%s) - 活跃: %t\n", user.Name, user.Email, user.IsActive)
    }
}
```

### CSV处理

```go
import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Product struct {
    ID       int
    Name     string
    Price    float64
    Category string
    InStock  bool
}

func csvOperations() {
    // 创建CSV数据
    products := []Product{
        {1, "Go语言编程", 89.90, "图书", true},
        {2, "MacBook Pro", 12999.00, "电脑", false},
        {3, "无线鼠标", 199.00, "配件", true},
        {4, "机械键盘", 599.00, "配件", true},
    }
    
    // 写入CSV文件
    file, err := os.Create("products.csv")
    if err != nil {
        fmt.Printf("创建CSV文件失败: %v\n", err)
        return
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // 写入表头
    headers := []string{"ID", "名称", "价格", "分类", "库存"}
    err = writer.Write(headers)
    if err != nil {
        fmt.Printf("写入CSV表头失败: %v\n", err)
        return
    }
    
    // 写入数据
    for _, product := range products {
        record := []string{
            strconv.Itoa(product.ID),
            product.Name,
            fmt.Sprintf("%.2f", product.Price),
            product.Category,
            strconv.FormatBool(product.InStock),
        }
        
        err = writer.Write(record)
        if err != nil {
            fmt.Printf("写入CSV数据失败: %v\n", err)
            return
        }
    }
    
    fmt.Println("✅ CSV文件写入成功")
    
    // 读取CSV文件
    readFile, err := os.Open("products.csv")
    if err != nil {
        fmt.Printf("打开CSV文件失败: %v\n", err)
        return
    }
    defer readFile.Close()
    
    reader := csv.NewReader(readFile)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("读取CSV文件失败: %v\n", err)
        return
    }
    
    fmt.Printf("📊 CSV数据 (%d行):\n", len(records))
    for i, record := range records {
        if i == 0 {
            fmt.Printf("表头: %v\n", record)
        } else {
            fmt.Printf("第%d行: %v\n", i, record)
        }
    }
}
```

## 实战项目：配置管理系统

让我们构建一个完整的配置文件管理系统：

```go
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

// 配置结构
type AppConfig struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Logging  LoggingConfig  `json:"logging"`
    Features FeatureConfig  `json:"features"`
}

type ServerConfig struct {
    Host         string `json:"host"`
    Port         int    `json:"port"`
    ReadTimeout  int    `json:"read_timeout"`
    WriteTimeout int    `json:"write_timeout"`
}

type DatabaseConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Database string `json:"database"`
    MaxConns int    `json:"max_connections"`
}

type LoggingConfig struct {
    Level    string `json:"level"`
    File     string `json:"file"`
    MaxSize  int    `json:"max_size_mb"`
    MaxFiles int    `json:"max_files"`
}

type FeatureConfig struct {
    EnableCache   bool     `json:"enable_cache"`
    EnableMetrics bool     `json:"enable_metrics"`
    AllowedHosts  []string `json:"allowed_hosts"`
}

// 配置管理器
type ConfigManager struct {
    configPath string
    config     *AppConfig
}

func NewConfigManager(configPath string) *ConfigManager {
    return &ConfigManager{
        configPath: configPath,
    }
}

// 创建默认配置
func (cm *ConfigManager) CreateDefaultConfig() *AppConfig {
    return &AppConfig{
        Server: ServerConfig{
            Host:         "localhost",
            Port:         8080,
            ReadTimeout:  15,
            WriteTimeout: 15,
        },
        Database: DatabaseConfig{
            Host:     "localhost",
            Port:     5432,
            Username: "app_user",
            Password: "your_password",
            Database: "app_db",
            MaxConns: 10,
        },
        Logging: LoggingConfig{
            Level:    "info",
            File:     "logs/app.log",
            MaxSize:  100,
            MaxFiles: 5,
        },
        Features: FeatureConfig{
            EnableCache:   true,
            EnableMetrics: false,
            AllowedHosts:  []string{"localhost", "127.0.0.1"},
        },
    }
}

// 确保配置目录存在
func (cm *ConfigManager) ensureConfigDir() error {
    dir := filepath.Dir(cm.configPath)
    return os.MkdirAll(dir, 0755)
}

// 保存配置到文件
func (cm *ConfigManager) SaveConfig(config *AppConfig) error {
    if err := cm.ensureConfigDir(); err != nil {
        return fmt.Errorf("创建配置目录失败: %v", err)
    }
    
    file, err := os.Create(cm.configPath)
    if err != nil {
        return fmt.Errorf("创建配置文件失败: %v", err)
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    
    if err := encoder.Encode(config); err != nil {
        return fmt.Errorf("编码配置失败: %v", err)
    }
    
    cm.config = config
    return nil
}

// 从文件加载配置
func (cm *ConfigManager) LoadConfig() (*AppConfig, error) {
    file, err := os.Open(cm.configPath)
    if err != nil {
        if os.IsNotExist(err) {
            // 配置文件不存在，创建默认配置
            fmt.Println("配置文件不存在，创建默认配置...")
            defaultConfig := cm.CreateDefaultConfig()
            if err := cm.SaveConfig(defaultConfig); err != nil {
                return nil, err
            }
            return defaultConfig, nil
        }
        return nil, fmt.Errorf("打开配置文件失败: %v", err)
    }
    defer file.Close()
    
    var config AppConfig
    decoder := json.NewDecoder(file)
    
    if err := decoder.Decode(&config); err != nil {
        return nil, fmt.Errorf("解码配置失败: %v", err)
    }
    
    cm.config = &config
    return &config, nil
}

// 备份配置文件
func (cm *ConfigManager) BackupConfig() error {
    if cm.config == nil {
        return fmt.Errorf("没有加载的配置")
    }
    
    timestamp := time.Now().Format("20060102_150405")
    backupPath := cm.configPath + ".backup." + timestamp
    
    sourceFile, err := os.Open(cm.configPath)
    if err != nil {
        return fmt.Errorf("打开源配置文件失败: %v", err)
    }
    defer sourceFile.Close()
    
    backupFile, err := os.Create(backupPath)
    if err != nil {
        return fmt.Errorf("创建备份文件失败: %v", err)
    }
    defer backupFile.Close()
    
    _, err = backupFile.ReadFrom(sourceFile)
    if err != nil {
        return fmt.Errorf("复制配置文件失败: %v", err)
    }
    
    fmt.Printf("✅ 配置已备份到: %s\n", backupPath)
    return nil
}

// 验证配置
func (cm *ConfigManager) ValidateConfig(config *AppConfig) []string {
    var issues []string
    
    // 验证服务器配置
    if config.Server.Port <= 0 || config.Server.Port > 65535 {
        issues = append(issues, "服务器端口无效")
    }
    
    if config.Server.Host == "" {
        issues = append(issues, "服务器主机地址为空")
    }
    
    // 验证数据库配置
    if config.Database.Username == "" {
        issues = append(issues, "数据库用户名为空")
    }
    
    if config.Database.Port <= 0 || config.Database.Port > 65535 {
        issues = append(issues, "数据库端口无效")
    }
    
    // 验证日志配置
    validLevels := []string{"debug", "info", "warn", "error"}
    levelValid := false
    for _, level := range validLevels {
        if config.Logging.Level == level {
            levelValid = true
            break
        }
    }
    if !levelValid {
        issues = append(issues, "日志级别无效")
    }
    
    return issues
}

// 交互式配置编辑器演示
func demonstrateConfigManager() {
    fmt.Println("🔧 配置管理系统演示")
    fmt.Println("================")
    
    // 创建配置管理器
    manager := NewConfigManager("config/app.json")
    
    // 加载配置
    config, err := manager.LoadConfig()
    if err != nil {
        fmt.Printf("加载配置失败: %v\n", err)
        return
    }
    
    fmt.Println("📄 当前配置:")
    fmt.Printf("服务器: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("数据库: %s:%d/%s\n", config.Database.Host, config.Database.Port, config.Database.Database)
    fmt.Printf("日志级别: %s\n", config.Logging.Level)
    fmt.Printf("缓存启用: %t\n", config.Features.EnableCache)
    
    // 验证配置
    if issues := manager.ValidateConfig(config); len(issues) > 0 {
        fmt.Println("\n⚠️ 配置验证问题:")
        for _, issue := range issues {
            fmt.Printf("  - %s\n", issue)
        }
    } else {
        fmt.Println("\n✅ 配置验证通过")
    }
    
    // 备份配置
    manager.BackupConfig()
}

// 日志管理器
type LogManager struct {
    logFile string
}

func NewLogManager(logFile string) *LogManager {
    return &LogManager{logFile: logFile}
}

func (lm *LogManager) EnsureLogDir() error {
    dir := filepath.Dir(lm.logFile)
    return os.MkdirAll(dir, 0755)
}

func (lm *LogManager) WriteLog(level, message string) error {
    if err := lm.EnsureLogDir(); err != nil {
        return err
    }
    
    file, err := os.OpenFile(lm.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
    
    _, err = file.WriteString(logEntry)
    return err
}

func (lm *LogManager) ReadRecentLogs(lines int) ([]string, error) {
    file, err := os.Open(lm.logFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var allLines []string
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        allLines = append(allLines, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    
    // 返回最后N行
    start := len(allLines) - lines
    if start < 0 {
        start = 0
    }
    
    return allLines[start:], nil
}

func main() {
    // 演示配置管理
    demonstrateConfigManager()
    
    // 演示日志功能
    fmt.Println("\n📝 日志管理演示:")
    logManager := NewLogManager("logs/app.log")
    
    // 写入一些日志
    logManager.WriteLog("INFO", "应用程序启动")
    logManager.WriteLog("DEBUG", "加载配置文件")
    logManager.WriteLog("ERROR", "数据库连接失败")
    logManager.WriteLog("INFO", "重试数据库连接")
    logManager.WriteLog("INFO", "应用程序就绪")
    
    // 读取最近的日志
    recentLogs, err := logManager.ReadRecentLogs(3)
    if err != nil {
        fmt.Printf("读取日志失败: %v\n", err)
    } else {
        fmt.Println("最近的日志记录:")
        for _, log := range recentLogs {
            fmt.Printf("  %s\n", log)
        }
    }
}
```

## 最佳实践

### 1. 资源管理

```go
func safeFileOperation() {
    file, err := os.Open("important.txt")
    if err != nil {
        return
    }
    defer file.Close() // 确保文件关闭
    
    // 使用文件...
}
```

### 2. 错误处理

```go
func robustFileRead(filename string) ([]byte, error) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return nil, fmt.Errorf("文件 %s 不存在", filename)
    }
    
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("读取文件失败: %v", err)
    }
    
    return data, nil
}
```

### 3. 原子操作

```go
func atomicFileWrite(filename string, data []byte) error {
    tempFile := filename + ".tmp"
    
    // 写入临时文件
    err := os.WriteFile(tempFile, data, 0644)
    if err != nil {
        return err
    }
    
    // 原子性重命名
    return os.Rename(tempFile, filename)
}
```

### 4. 大文件处理

```go
func processLargeFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    const maxCapacity = 1024 * 1024 // 1MB buffer
    buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity)
    
    for scanner.Scan() {
        line := scanner.Text()
        // 处理每一行
        processLine(line)
    }
    
    return scanner.Err()
}

func processLine(line string) {
    // 处理单行数据
}
```

## 本章小结

Go文件操作的核心要点：

- **基础操作**：使用os包进行文件读写、创建、删除等操作
- **缓冲I/O**：使用bufio包提升大文件处理性能
- **路径处理**：使用filepath包进行跨平台路径操作
- **数据格式**：处理JSON、CSV等常见数据格式
- **资源管理**：正确使用defer确保资源释放

### 下一步
掌握了文件操作后，我们将学习 [字符串和正则表达式](./strings-regexp.md)，了解文本处理和模式匹配。

::: tip 练习建议
1. 实现一个日志轮转系统
2. 创建配置文件热重载功能
3. 开发文件同步工具
4. 构建数据导入导出工具
::: 