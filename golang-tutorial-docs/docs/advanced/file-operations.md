---
title: 文件操作
description: 学习Go语言的文件读写、目录操作和数据处理
---

# 文件操作

文件操作是后端开发的基础技能。Go语言提供了丰富的文件操作API，让我们一起掌握这些重要功能。

## 📖 本章内容

- 基本文件读写操作
- 目录遍历和管理
- 文件信息获取和处理
- JSON/XML/CSV数据处理
- 文件监控和批量操作

## 📁 基本文件操作

### 文件读取

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func main() {
    // 方法1：一次性读取整个文件
    readWholeFile()
    
    // 方法2：逐行读取文件
    readFileLineByLine()
    
    // 方法3：读取固定大小的块
    readFileInChunks()
}

// 一次性读取整个文件
func readWholeFile() {
    fmt.Println("=== 一次性读取整个文件 ===")
    
    // 使用 os.ReadFile (Go 1.16+)
    content, err := os.ReadFile("example.txt")
    if err != nil {
        log.Printf("读取文件失败: %v", err)
        
        // 创建示例文件
        createExampleFile()
        content, err = os.ReadFile("example.txt")
        if err != nil {
            log.Fatal(err)
        }
    }
    
    fmt.Printf("文件内容:\n%s\n", string(content))
    fmt.Printf("文件大小: %d 字节\n\n", len(content))
}

// 逐行读取文件
func readFileLineByLine() {
    fmt.Println("=== 逐行读取文件 ===")
    
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNum := 1
    
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("第%d行: %s\n", lineNum, line)
        lineNum++
    }
    
    if err := scanner.Err(); err != nil {
        log.Printf("读取文件时出错: %v", err)
    }
    fmt.Println()
}

// 读取固定大小的块
func readFileInChunks() {
    fmt.Println("=== 分块读取文件 ===")
    
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    buffer := make([]byte, 32) // 每次读取32字节
    chunkNum := 1
    
    for {
        n, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                fmt.Println("文件读取完成")
                break
            }
            log.Fatal(err)
        }
        
        fmt.Printf("第%d块 (%d字节): %q\n", chunkNum, n, string(buffer[:n]))
        chunkNum++
    }
    fmt.Println()
}

// 创建示例文件
func createExampleFile() {
    content := `Go语言文件操作示例
这是第二行内容
包含中文和English混合内容
数字: 12345
特殊字符: !@#$%^&*()`
    
    err := os.WriteFile("example.txt", []byte(content), 0644)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("已创建示例文件: example.txt")
}
```

### 文件写入

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    // 方法1：一次性写入文件
    writeWholeFile()
    
    // 方法2：逐行写入文件
    writeFileLineByLine()
    
    // 方法3：追加内容到文件
    appendToFile()
    
    // 方法4：使用缓冲写入
    writeWithBuffer()
}

// 一次性写入文件
func writeWholeFile() {
    fmt.Println("=== 一次性写入文件 ===")
    
    content := fmt.Sprintf(`文件写入测试
当前时间: %s
Go语言版本: 1.21
测试内容包含多行数据`, time.Now().Format("2006-01-02 15:04:05"))
    
    err := os.WriteFile("output.txt", []byte(content), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("已写入文件: output.txt")
    
    // 验证写入结果
    readContent, _ := os.ReadFile("output.txt")
    fmt.Printf("写入内容:\n%s\n\n", string(readContent))
}

// 逐行写入文件
func writeFileLineByLine() {
    fmt.Println("=== 逐行写入文件 ===")
    
    file, err := os.Create("lines.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    lines := []string{
        "第一行数据",
        "第二行: 包含数字 123",
        "第三行: 包含特殊字符 !@#",
        "第四行: English content",
        "第五行: 最后一行",
    }
    
    for i, line := range lines {
        _, err := fmt.Fprintf(file, "%d. %s\n", i+1, line)
        if err != nil {
            log.Fatal(err)
        }
    }
    
    fmt.Println("已逐行写入文件: lines.txt")
    
    // 验证结果
    content, _ := os.ReadFile("lines.txt")
    fmt.Printf("文件内容:\n%s\n", string(content))
}

// 追加内容到文件
func appendToFile() {
    fmt.Println("=== 追加内容到文件 ===")
    
    file, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    appendLines := []string{
        "追加行1: 新增内容",
        "追加行2: " + time.Now().Format("15:04:05"),
        "追加行3: 追加操作完成",
    }
    
    for _, line := range appendLines {
        _, err := fmt.Fprintf(file, "%s\n", line)
        if err != nil {
            log.Fatal(err)
        }
    }
    
    fmt.Println("已追加内容到文件: lines.txt")
    
    // 验证结果
    content, _ := os.ReadFile("lines.txt")
    fmt.Printf("追加后的文件内容:\n%s\n", string(content))
}

// 使用缓冲写入
func writeWithBuffer() {
    fmt.Println("=== 使用缓冲写入 ===")
    
    file, err := os.Create("buffered.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    defer writer.Flush() // 确保缓冲区内容被写入
    
    // 写入大量数据
    for i := 1; i <= 1000; i++ {
        line := fmt.Sprintf("第%d行: 这是缓冲写入测试数据 - %s\n", 
            i, time.Now().Format("15:04:05.000"))
        _, err := writer.WriteString(line)
        if err != nil {
            log.Fatal(err)
        }
        
        // 每100行手动刷新缓冲区
        if i%100 == 0 {
            writer.Flush()
            fmt.Printf("已写入 %d 行\n", i)
        }
    }
    
    fmt.Println("缓冲写入完成: buffered.txt")
    
    // 检查文件大小
    info, _ := os.Stat("buffered.txt")
    fmt.Printf("文件大小: %d 字节\n\n", info.Size())
}
```

## 📂 目录操作

### 目录遍历和管理

```go
package main

import (
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "time"
)

func main() {
    // 创建测试目录结构
    createTestDirectories()
    
    // 目录基本操作
    directoryBasicOps()
    
    // 遍历目录
    walkDirectory()
    
    // 查找特定文件
    findFiles()
    
    // 计算目录大小
    calculateDirSize()
    
    // 清理测试目录
    cleanup()
}

// 创建测试目录结构
func createTestDirectories() {
    fmt.Println("=== 创建测试目录结构 ===")
    
    dirs := []string{
        "testdir",
        "testdir/subdir1",
        "testdir/subdir2",
        "testdir/subdir1/deep",
        "testdir/files",
    }
    
    for _, dir := range dirs {
        err := os.MkdirAll(dir, 0755)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("创建目录: %s\n", dir)
    }
    
    // 创建一些测试文件
    files := map[string]string{
        "testdir/readme.txt":           "这是说明文件",
        "testdir/subdir1/data.json":    `{"name": "test", "value": 123}`,
        "testdir/subdir1/deep/log.txt": "深层目录中的日志文件",
        "testdir/subdir2/config.yaml":  "config:\n  debug: true",
        "testdir/files/image.jpg":      "fake image data",
        "testdir/files/document.pdf":   "fake pdf data",
    }
    
    for path, content := range files {
        err := os.WriteFile(path, []byte(content), 0644)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("创建文件: %s\n", path)
    }
    fmt.Println()
}

// 目录基本操作
func directoryBasicOps() {
    fmt.Println("=== 目录基本操作 ===")
    
    // 获取当前工作目录
    pwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("当前工作目录: %s\n", pwd)
    
    // 检查目录是否存在
    if _, err := os.Stat("testdir"); err == nil {
        fmt.Println("testdir 目录存在")
    } else if os.IsNotExist(err) {
        fmt.Println("testdir 目录不存在")
    }
    
    // 读取目录内容
    entries, err := os.ReadDir("testdir")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("testdir 目录内容:")
    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Printf("  📁 %s/\n", entry.Name())
        } else {
            info, _ := entry.Info()
            fmt.Printf("  📄 %s (%d bytes)\n", entry.Name(), info.Size())
        }
    }
    fmt.Println()
}

// 遍历目录
func walkDirectory() {
    fmt.Println("=== 遍历目录 ===")
    
    err := filepath.WalkDir("testdir", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        
        // 计算缩进级别
        level := len(filepath.SplitList(path)) - 1
        indent := ""
        for i := 0; i < level; i++ {
            indent += "  "
        }
        
        if d.IsDir() {
            fmt.Printf("%s📁 %s/\n", indent, d.Name())
        } else {
            info, _ := d.Info()
            fmt.Printf("%s📄 %s (%d bytes, %s)\n", 
                indent, d.Name(), info.Size(), info.ModTime().Format("15:04:05"))
        }
        
        return nil
    })
    
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println()
}

// 查找特定文件
func findFiles() {
    fmt.Println("=== 查找特定文件 ===")
    
    // 查找所有 .txt 文件
    txtFiles := []string{}
    err := filepath.WalkDir("testdir", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        
        if !d.IsDir() && filepath.Ext(path) == ".txt" {
            txtFiles = append(txtFiles, path)
        }
        
        return nil
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("找到的 .txt 文件:")
    for _, file := range txtFiles {
        info, _ := os.Stat(file)
        fmt.Printf("  %s (%d bytes)\n", file, info.Size())
    }
    
    // 使用 Glob 模式匹配
    matches, err := filepath.Glob("testdir/**/*.json")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("找到的 .json 文件 (使用Glob):")
    for _, match := range matches {
        fmt.Printf("  %s\n", match)
    }
    fmt.Println()
}

// 计算目录大小
func calculateDirSize() {
    fmt.Println("=== 计算目录大小 ===")
    
    var totalSize int64
    fileCount := 0
    dirCount := 0
    
    err := filepath.WalkDir("testdir", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        
        if d.IsDir() {
            dirCount++
        } else {
            info, _ := d.Info()
            totalSize += info.Size()
            fileCount++
        }
        
        return nil
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("目录统计:\n")
    fmt.Printf("  总大小: %d 字节\n", totalSize)
    fmt.Printf("  文件数: %d\n", fileCount)
    fmt.Printf("  目录数: %d\n", dirCount)
    fmt.Println()
}

// 清理测试目录
func cleanup() {
    fmt.Println("=== 清理测试目录 ===")
    
    err := os.RemoveAll("testdir")
    if err != nil {
        log.Fatal(err)
    }
    
    // 清理其他测试文件
    testFiles := []string{"output.txt", "lines.txt", "buffered.txt", "example.txt"}
    for _, file := range testFiles {
        os.Remove(file) // 忽略错误，文件可能不存在
    }
    
    fmt.Println("清理完成")
}
```

## 📊 数据格式处理

### JSON 数据处理

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
)

// 用户结构体
type User struct {
    ID       int       `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Age      int       `json:"age"`
    IsActive bool      `json:"is_active"`
    Created  time.Time `json:"created_at"`
    Profile  Profile   `json:"profile"`
    Tags     []string  `json:"tags"`
}

// 用户资料结构体
type Profile struct {
    Bio     string `json:"bio"`
    Website string `json:"website,omitempty"`
    Company string `json:"company,omitempty"`
}

func main() {
    // JSON编码和解码
    jsonEncodeAndDecode()
    
    // 读写JSON文件
    jsonFileOperations()
    
    // 处理动态JSON
    handleDynamicJSON()
    
    // JSON流处理
    jsonStreamProcessing()
}

// JSON编码和解码
func jsonEncodeAndDecode() {
    fmt.Println("=== JSON编码和解码 ===")
    
    // 创建示例用户
    user := User{
        ID:       1,
        Name:     "张三",
        Email:    "zhangsan@example.com",
        Age:      28,
        IsActive: true,
        Created:  time.Now(),
        Profile: Profile{
            Bio:     "Go语言开发者",
            Website: "https://example.com",
            Company: "科技公司",
        },
        Tags: []string{"golang", "backend", "microservices"},
    }
    
    // 编码为JSON
    jsonData, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("JSON编码结果:\n%s\n\n", string(jsonData))
    
    // 解码JSON
    var decodedUser User
    err = json.Unmarshal(jsonData, &decodedUser)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("解码后的用户信息:\n")
    fmt.Printf("姓名: %s\n", decodedUser.Name)
    fmt.Printf("邮箱: %s\n", decodedUser.Email)
    fmt.Printf("年龄: %d\n", decodedUser.Age)
    fmt.Printf("创建时间: %s\n", decodedUser.Created.Format("2006-01-02 15:04:05"))
    fmt.Printf("标签: %v\n\n", decodedUser.Tags)
}

// 读写JSON文件
func jsonFileOperations() {
    fmt.Println("=== 读写JSON文件 ===")
    
    // 创建多个用户
    users := []User{
        {
            ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 28, IsActive: true,
            Created: time.Now(),
            Profile: Profile{Bio: "Go开发者", Company: "A公司"},
            Tags:    []string{"golang", "backend"},
        },
        {
            ID: 2, Name: "李四", Email: "lisi@example.com", Age: 32, IsActive: false,
            Created: time.Now().Add(-24 * time.Hour),
            Profile: Profile{Bio: "前端开发者", Company: "B公司"},
            Tags:    []string{"javascript", "react"},
        },
        {
            ID: 3, Name: "王五", Email: "wangwu@example.com", Age: 25, IsActive: true,
            Created: time.Now().Add(-48 * time.Hour),
            Profile: Profile{Bio: "全栈开发者", Website: "https://wangwu.dev"},
            Tags:    []string{"golang", "javascript", "python"},
        },
    }
    
    // 写入JSON文件
    jsonData, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    
    err = os.WriteFile("users.json", jsonData, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("已写入用户数据到 users.json")
    
    // 读取JSON文件
    fileData, err := os.ReadFile("users.json")
    if err != nil {
        log.Fatal(err)
    }
    
    var loadedUsers []User
    err = json.Unmarshal(fileData, &loadedUsers)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("从文件加载了 %d 个用户:\n", len(loadedUsers))
    for _, user := range loadedUsers {
        fmt.Printf("- %s (%s) - 活跃: %t\n", user.Name, user.Email, user.IsActive)
    }
    fmt.Println()
}

// 处理动态JSON
func handleDynamicJSON() {
    fmt.Println("=== 处理动态JSON ===")
    
    // 模拟接收到的动态JSON数据
    dynamicJSON := `{
        "event": "user_login",
        "timestamp": "2024-01-15T10:30:00Z",
        "user_id": 123,
        "metadata": {
            "ip": "192.168.1.100",
            "user_agent": "Mozilla/5.0...",
            "platform": "web"
        },
        "properties": {
            "login_method": "email",
            "remember_me": true,
            "session_duration": 3600
        }
    }`
    
    // 使用 map[string]interface{} 处理动态JSON
    var data map[string]interface{}
    err := json.Unmarshal([]byte(dynamicJSON), &data)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("动态JSON解析结果:")
    fmt.Printf("事件类型: %s\n", data["event"])
    fmt.Printf("用户ID: %.0f\n", data["user_id"])
    
    // 处理嵌套对象
    if metadata, ok := data["metadata"].(map[string]interface{}); ok {
        fmt.Printf("IP地址: %s\n", metadata["ip"])
        fmt.Printf("平台: %s\n", metadata["platform"])
    }
    
    // 处理不同类型的值
    if properties, ok := data["properties"].(map[string]interface{}); ok {
        for key, value := range properties {
            fmt.Printf("属性 %s: %v (类型: %T)\n", key, value, value)
        }
    }
    fmt.Println()
}

// JSON流处理
func jsonStreamProcessing() {
    fmt.Println("=== JSON流处理 ===")
    
    // 创建大量数据进行流处理演示
    file, err := os.Create("stream_data.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    
    // 写入JSON数组的开始
    file.WriteString("[\n")
    
    // 流式写入多个JSON对象
    for i := 1; i <= 5; i++ {
        user := User{
            ID:       i,
            Name:     fmt.Sprintf("用户%d", i),
            Email:    fmt.Sprintf("user%d@example.com", i),
            Age:      20 + i*2,
            IsActive: i%2 == 1,
            Created:  time.Now().Add(time.Duration(-i) * time.Hour),
            Profile: Profile{
                Bio: fmt.Sprintf("这是用户%d的简介", i),
            },
            Tags: []string{fmt.Sprintf("tag%d", i)},
        }
        
        if i > 1 {
            file.WriteString(",\n")
        }
        
        // 使用encoder写入，但不包含数组括号
        userData, _ := json.MarshalIndent(user, "  ", "  ")
        file.WriteString("  " + string(userData))
    }
    
    file.WriteString("\n]")
    
    fmt.Println("已创建流数据文件: stream_data.json")
    
    // 流式读取JSON数据
    file, err = os.Open("stream_data.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    
    // 读取数组开始标记
    token, err := decoder.Token()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("开始标记: %v\n", token)
    
    userCount := 0
    // 逐个读取用户对象
    for decoder.More() {
        var user User
        err := decoder.Decode(&user)
        if err != nil {
            log.Fatal(err)
        }
        
        userCount++
        fmt.Printf("流式读取用户%d: %s (%s)\n", 
            userCount, user.Name, user.Email)
    }
    
    // 读取数组结束标记
    token, err = decoder.Token()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("结束标记: %v\n", token)
    
    // 清理文件
    os.Remove("users.json")
    os.Remove("stream_data.json")
    fmt.Println()
}
```

## 📝 CSV 和 XML 处理

### CSV 文件处理

```go
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// 学生结构体
type Student struct {
    ID     int
    Name   string
    Age    int
    Grade  string
    Score  float64
    Active bool
}

func main() {
    // CSV读写操作
    csvOperations()
    
    // 处理大型CSV文件
    processBigCSV()
    
    // CSV数据转换
    csvDataConversion()
}

// CSV读写操作
func csvOperations() {
    fmt.Println("=== CSV读写操作 ===")
    
    // 创建学生数据
    students := []Student{
        {1, "张三", 20, "A", 85.5, true},
        {2, "李四", 19, "B", 92.0, true},
        {3, "王五", 21, "A", 78.5, false},
        {4, "赵六", 20, "C", 88.0, true},
        {5, "钱七", 22, "B", 95.5, true},
    }
    
    // 写入CSV文件
    writeCSV(students)
    
    // 读取CSV文件
    loadedStudents := readCSV()
    
    // 显示读取结果
    fmt.Println("从CSV文件读取的学生数据:")
    for _, student := range loadedStudents {
        fmt.Printf("ID:%d, 姓名:%s, 年龄:%d, 等级:%s, 分数:%.1f, 活跃:%t\n",
            student.ID, student.Name, student.Age, student.Grade, student.Score, student.Active)
    }
    fmt.Println()
}

// 写入CSV文件
func writeCSV(students []Student) {
    file, err := os.Create("students.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // 写入表头
    header := []string{"ID", "姓名", "年龄", "等级", "分数", "活跃状态"}
    err = writer.Write(header)
    if err != nil {
        log.Fatal(err)
    }
    
    // 写入数据行
    for _, student := range students {
        record := []string{
            strconv.Itoa(student.ID),
            student.Name,
            strconv.Itoa(student.Age),
            student.Grade,
            fmt.Sprintf("%.1f", student.Score),
            strconv.FormatBool(student.Active),
        }
        
        err = writer.Write(record)
        if err != nil {
            log.Fatal(err)
        }
    }
    
    fmt.Println("已写入学生数据到 students.csv")
}

// 读取CSV文件
func readCSV() []Student {
    file, err := os.Open("students.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    
    var students []Student
    
    // 跳过表头，从第二行开始处理
    for i, record := range records[1:] {
        if len(record) != 6 {
            fmt.Printf("第%d行数据格式错误: %v\n", i+2, record)
            continue
        }
        
        id, err := strconv.Atoi(record[0])
        if err != nil {
            fmt.Printf("第%d行ID转换错误: %v\n", i+2, err)
            continue
        }
        
        age, err := strconv.Atoi(record[2])
        if err != nil {
            fmt.Printf("第%d行年龄转换错误: %v\n", i+2, err)
            continue
        }
        
        score, err := strconv.ParseFloat(record[4], 64)
        if err != nil {
            fmt.Printf("第%d行分数转换错误: %v\n", i+2, err)
            continue
        }
        
        active, err := strconv.ParseBool(record[5])
        if err != nil {
            fmt.Printf("第%d行活跃状态转换错误: %v\n", i+2, err)
            continue
        }
        
        student := Student{
            ID:     id,
            Name:   record[1],
            Age:    age,
            Grade:  record[3],
            Score:  score,
            Active: active,
        }
        
        students = append(students, student)
    }
    
    return students
}

// 处理大型CSV文件
func processBigCSV() {
    fmt.Println("=== 处理大型CSV文件 ===")
    
    // 创建大型CSV文件用于演示
    createBigCSV()
    
    // 流式处理大文件
    file, err := os.Open("big_data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    
    var totalScore float64
    var count int
    var highScoreCount int
    
    // 读取表头
    header, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CSV表头: %v\n", header)
    
    // 逐行处理
    for {
        record, err := reader.Read()
        if err != nil {
            break // 文件结束
        }
        
        count++
        
        // 解析分数
        if len(record) >= 3 {
            score, err := strconv.ParseFloat(record[2], 64)
            if err == nil {
                totalScore += score
                if score >= 90 {
                    highScoreCount++
                }
            }
        }
        
        // 每处理1000行显示进度
        if count%1000 == 0 {
            fmt.Printf("已处理 %d 行...\n", count)
        }
    }
    
    // 统计结果
    fmt.Printf("处理完成:\n")
    fmt.Printf("  总行数: %d\n", count)
    fmt.Printf("  平均分数: %.2f\n", totalScore/float64(count))
    fmt.Printf("  高分(>=90)人数: %d\n", highScoreCount)
    fmt.Printf("  高分比例: %.2f%%\n", float64(highScoreCount)/float64(count)*100)
    
    // 清理文件
    os.Remove("big_data.csv")
    fmt.Println()
}

// 创建大型CSV文件
func createBigCSV() {
    file, err := os.Create("big_data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // 写入表头
    writer.Write([]string{"ID", "姓名", "分数"})
    
    // 写入5000行数据
    for i := 1; i <= 5000; i++ {
        record := []string{
            strconv.Itoa(i),
            fmt.Sprintf("学生%d", i),
            fmt.Sprintf("%.1f", 60.0+float64(i%40)), // 分数在60-100之间
        }
        writer.Write(record)
    }
    
    fmt.Println("已创建大型CSV文件: big_data.csv (5000行)")
}

// CSV数据转换
func csvDataConversion() {
    fmt.Println("=== CSV数据转换 ===")
    
    // 读取原始CSV
    students := readCSV()
    
    // 数据统计和转换
    gradeStats := make(map[string]int)
    var totalScore float64
    activeCount := 0
    
    for _, student := range students {
        gradeStats[student.Grade]++
        totalScore += student.Score
        if student.Active {
            activeCount++
        }
    }
    
    // 创建统计报告CSV
    reportFile, err := os.Create("report.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer reportFile.Close()
    
    writer := csv.NewWriter(reportFile)
    defer writer.Flush()
    
    // 写入统计报告
    writer.Write([]string{"统计项", "数值"})
    writer.Write([]string{"总学生数", strconv.Itoa(len(students))})
    writer.Write([]string{"平均分数", fmt.Sprintf("%.2f", totalScore/float64(len(students)))})
    writer.Write([]string{"活跃学生数", strconv.Itoa(activeCount)})
    writer.Write([]string{"活跃比例", fmt.Sprintf("%.2f%%", float64(activeCount)/float64(len(students))*100)})
    
    // 等级分布
    writer.Write([]string{"", ""}) // 空行
    writer.Write([]string{"等级分布", ""})
    for grade, count := range gradeStats {
        writer.Write([]string{fmt.Sprintf("等级%s", grade), strconv.Itoa(count)})
    }
    
    fmt.Println("已生成统计报告: report.csv")
    
    // 显示统计结果
    fmt.Printf("数据统计结果:\n")
    fmt.Printf("  总学生数: %d\n", len(students))
    fmt.Printf("  平均分数: %.2f\n", totalScore/float64(len(students)))
    fmt.Printf("  活跃学生: %d (%.1f%%)\n", activeCount, 
        float64(activeCount)/float64(len(students))*100)
    fmt.Printf("  等级分布: %v\n", gradeStats)
    
    // 清理文件
    os.Remove("students.csv")
    os.Remove("report.csv")
    fmt.Println()
}
```

## 📝 本章小结

在这一章中，我们学习了：

### 🔹 基本文件操作
- 文件读取的多种方式
- 文件写入和追加操作
- 缓冲读写提升性能

### 🔹 目录管理
- 目录创建和删除
- 目录遍历和搜索
- 文件信息获取

### 🔹 数据格式处理
- JSON编码解码和文件操作
- CSV文件的读写和处理
- 大文件的流式处理

### 🔹 实用技巧
- 错误处理最佳实践
- 性能优化方法
- 数据转换和统计

## 🎯 下一步

掌握了文件操作后，让我们继续学习 [网络编程](./network-http)，探索网络通信和HTTP服务开发！ 