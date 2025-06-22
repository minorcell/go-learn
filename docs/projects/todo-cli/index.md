---
title: TODO CLI 应用
description: 构建一个功能完整的命令行待办事项管理工具
---

# TODO CLI 应用

让我们一步步构建一个功能完整的命令行待办事项管理工具。这个项目将涵盖文件操作、JSON处理、命令行参数解析等核心技能。

## 项目概述

### 功能需求
- 添加新的待办事项
- 列出所有待办事项
- 修改待办事项状态
- 删除待办事项
- 数据持久化存储
- 彩色命令行输出

### 技术要点
- 结构体设计和方法
- JSON文件读写
- 命令行参数处理
- 错误处理机制
- 时间处理和格式化

## 第一步：项目结构设计

### 解释
首先我们需要设计项目的基础结构，包括数据模型、文件组织和核心功能规划。

### 编码
创建项目目录和基础文件：

```go
// main.go - 主程序入口
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

// Todo 待办事项结构体
type Todo struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// TodoList 待办事项列表
type TodoList struct {
    Todos    []Todo `json:"todos"`
    NextID   int    `json:"next_id"`
    FilePath string `json:"-"` // 不序列化到JSON
}

// 创建新的TodoList
func NewTodoList(filePath string) *TodoList {
    return &TodoList{
        Todos:    make([]Todo, 0),
        NextID:   1,
        FilePath: filePath,
    }
}

func main() {
    fmt.Println("TODO CLI 应用启动")
    
    // 创建TodoList实例
    todoList := NewTodoList("todos.json")
    
    // 加载现有数据
    if err := todoList.Load(); err != nil {
        fmt.Printf("创建新的待办事项文件: %s\n", todoList.FilePath)
    } else {
        fmt.Printf("加载现有待办事项: %d 个任务\n", len(todoList.Todos))
    }
    
    // 解析命令行参数
    if len(os.Args) < 2 {
        showHelp()
        return
    }
    
    command := os.Args[1]
    
    switch command {
    case "add", "a":
        handleAdd(todoList, os.Args[2:])
    case "list", "l":
        handleList(todoList, os.Args[2:])
    case "done", "d":
        handleDone(todoList, os.Args[2:])
    case "remove", "r":
        handleRemove(todoList, os.Args[2:])
    case "help", "h":
        showHelp()
    default:
        fmt.Printf("未知命令: %s\n", command)
        showHelp()
    }
}

// 显示帮助信息
func showHelp() {
    fmt.Println(`
TODO CLI 使用帮助

命令格式: todo <command> [arguments]

可用命令:
  add, a     <title> [description]  添加新的待办事项
  list, l    [status]               列出待办事项 (all/pending/done)
  done, d    <id>                   标记待办事项为完成
  remove, r  <id>                   删除待办事项
  help, h                           显示此帮助信息

使用示例:
  todo add "学习Go语言" "完成Go语言基础教程"
  todo list
  todo done 1
  todo remove 2
    `)
}
```

### 运行结果
```bash
$ go run main.go
TODO CLI 应用启动
创建新的待办事项文件: todos.json

TODO CLI 使用帮助

命令格式: todo <command> [arguments]
...
```

## 第二步：数据持久化

### 解释
实现JSON文件的读写功能，确保待办事项数据能够持久化存储。这包括加载现有数据和保存更改。

### 编码

```go
// 加载待办事项数据
func (tl *TodoList) Load() error {
    // 检查文件是否存在
    if _, err := os.Stat(tl.FilePath); os.IsNotExist(err) {
        return err
    }
    
    // 读取文件内容
    data, err := ioutil.ReadFile(tl.FilePath)
    if err != nil {
        return fmt.Errorf("读取文件失败: %v", err)
    }
    
    // 解析JSON数据
    if err := json.Unmarshal(data, tl); err != nil {
        return fmt.Errorf("解析JSON失败: %v", err)
    }
    
    return nil
}

// 保存待办事项数据
func (tl *TodoList) Save() error {
    // 序列化为JSON
    data, err := json.MarshalIndent(tl, "", "  ")
    if err != nil {
        return fmt.Errorf("序列化JSON失败: %v", err)
    }
    
    // 写入文件
    if err := ioutil.WriteFile(tl.FilePath, data, 0644); err != nil {
        return fmt.Errorf("写入文件失败: %v", err)
    }
    
    return nil
}

// 添加新的待办事项
func (tl *TodoList) Add(title, description string) *Todo {
    todo := Todo{
        ID:          tl.NextID,
        Title:       title,
        Description: description,
        Completed:   false,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    tl.Todos = append(tl.Todos, todo)
    tl.NextID++
    
    return &todo
}

// 根据ID查找待办事项
func (tl *TodoList) FindByID(id int) *Todo {
    for i := range tl.Todos {
        if tl.Todos[i].ID == id {
            return &tl.Todos[i]
        }
    }
    return nil
}

// 标记待办事项为完成
func (tl *TodoList) MarkDone(id int) bool {
    todo := tl.FindByID(id)
    if todo == nil {
        return false
    }
    
    todo.Completed = true
    todo.UpdatedAt = time.Now()
    return true
}

// 删除待办事项
func (tl *TodoList) Remove(id int) bool {
    for i, todo := range tl.Todos {
        if todo.ID == id {
            tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
            return true
        }
    }
    return false
}
```

### 运行结果
数据结构设计完成后，我们有了完整的数据模型和基础操作方法。接下来实现具体的命令处理。

## 第三步：添加待办事项

### 解释
实现添加新待办事项的功能，包括参数验证、数据创建和保存操作。

### 编码

```go
// 处理添加命令
func handleAdd(todoList *TodoList, args []string) {
    if len(args) < 1 {
        fmt.Println("请提供待办事项标题")
        fmt.Println("使用方法: todo add <title> [description]")
        return
    }
    
    title := args[0]
    description := ""
    
    // 如果提供了描述
    if len(args) > 1 {
        description = args[1]
    }
    
    // 验证标题
    if strings.TrimSpace(title) == "" {
        fmt.Println("待办事项标题不能为空")
        return
    }
    
    // 添加待办事项
    todo := todoList.Add(title, description)
    
    // 保存到文件
    if err := todoList.Save(); err != nil {
        fmt.Printf("保存失败: %v\n", err)
        return
    }
    
    // 成功反馈
    fmt.Printf("已添加待办事项:\n")
    fmt.Printf("   ID: %d\n", todo.ID)
    fmt.Printf("   标题: %s\n", todo.Title)
    if todo.Description != "" {
        fmt.Printf("   描述: %s\n", todo.Description)
    }
    fmt.Printf("   创建时间: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
}
```

### 运行结果
```bash
$ go run main.go add "学习Go语言" "完成Go语言基础教程"
TODO CLI 应用启动
创建新的待办事项文件: todos.json
已添加待办事项:
   ID: 1
   标题: 学习Go语言
   描述: 完成Go语言基础教程
   创建时间: 2024-01-15 10:30:45

$ go run main.go add "买菜"
TODO CLI 应用启动
加载现有待办事项: 1 个任务
已添加待办事项:
   ID: 2
   标题: 买菜
   创建时间: 2024-01-15 10:31:20
```

## 第四步：列出待办事项

### 解释
实现列表显示功能，支持按状态筛选（全部/待完成/已完成），并提供清晰的彩色输出。

### 编码

```go
// 处理列表命令
func handleList(todoList *TodoList, args []string) {
    filter := "all"
    if len(args) > 0 {
        filter = args[0]
    }
    
    // 验证过滤器
    if filter != "all" && filter != "pending" && filter != "done" {
        fmt.Printf("无效的过滤器: %s\n", filter)
        fmt.Println("可用选项: all, pending, done")
        return
    }
    
    // 过滤待办事项
    var filteredTodos []Todo
    for _, todo := range todoList.Todos {
        switch filter {
        case "all":
            filteredTodos = append(filteredTodos, todo)
        case "pending":
            if !todo.Completed {
                filteredTodos = append(filteredTodos, todo)
            }
        case "done":
            if todo.Completed {
                filteredTodos = append(filteredTodos, todo)
            }
        }
    }
    
    // 显示结果
    if len(filteredTodos) == 0 {
        switch filter {
        case "all":
            fmt.Println("📭 暂无待办事项")
        case "pending":
            fmt.Println("没有待完成的事项")
        case "done":
            fmt.Println("😴 没有已完成的事项")
        }
        return
    }
    
    // 显示标题
    switch filter {
    case "all":
        fmt.Printf("所有待办事项 (%d个):\n\n", len(filteredTodos))
    case "pending":
        fmt.Printf("⏳ 待完成事项 (%d个):\n\n", len(filteredTodos))
    case "done":
        fmt.Printf("已完成事项 (%d个):\n\n", len(filteredTodos))
    }
    
    // 显示每个待办事项
    for _, todo := range filteredTodos {
        displayTodo(todo)
        fmt.Println()
    }
    
    // 显示统计信息
    showStatistics(todoList.Todos)
}

// 显示单个待办事项
func displayTodo(todo Todo) {
    // 状态图标
    status := "⏳"
    if todo.Completed {
        status = "✅"
    }
    
    fmt.Printf("%s [%d] %s\n", status, todo.ID, todo.Title)
    
    if todo.Description != "" {
        fmt.Printf("     %s\n", todo.Description)
    }
    
    fmt.Printf("    创建: %s\n", todo.CreatedAt.Format("2006-01-02 15:04"))
    
    if !todo.UpdatedAt.Equal(todo.CreatedAt) {
        fmt.Printf("    更新: %s\n", todo.UpdatedAt.Format("2006-01-02 15:04"))
    }
}

// 显示统计信息
func showStatistics(todos []Todo) {
    total := len(todos)
    completed := 0
    pending := 0
    
    for _, todo := range todos {
        if todo.Completed {
            completed++
        } else {
            pending++
        }
    }
    
    fmt.Println("统计信息:")
    fmt.Printf("   总计: %d 个事项\n", total)
    fmt.Printf("   已完成: %d 个事项\n", completed)
    fmt.Printf("   待完成: %d 个事项\n", pending)
    
    if total > 0 {
        completion := float64(completed) / float64(total) * 100
        fmt.Printf("   完成率: %.1f%%\n", completion)
    }
}
```

### 运行结果
```bash
$ go run main.go list
TODO CLI 应用启动
加载现有待办事项: 2 个任务
所有待办事项 (2个):

⏳ [1] 学习Go语言
     完成Go语言基础教程
    创建: 2024-01-15 10:30

⏳ [2] 买菜
    创建: 2024-01-15 10:31

统计信息:
   总计: 2 个事项
   已完成: 0 个事项
   待完成: 2 个事项
   完成率: 0.0%

$ go run main.go list pending
TODO CLI 应用启动
加载现有待办事项: 2 个任务
⏳ 待完成事项 (2个):

⏳ [1] 学习Go语言
     完成Go语言基础教程
    创建: 2024-01-15 10:30

⏳ [2] 买菜
    创建: 2024-01-15 10:31

统计信息:
   总计: 2 个事项
   已完成: 0 个事项
   待完成: 2 个事项
   完成率: 0.0%
```

## 第五步：标记完成

### 解释
实现标记待办事项为完成的功能，包括ID验证、状态更新和保存操作。

### 编码

```go
// 处理完成命令
func handleDone(todoList *TodoList, args []string) {
    if len(args) < 1 {
        fmt.Println("请提供待办事项ID")
        fmt.Println("使用方法: todo done <id>")
        return
    }
    
    // 解析ID
    id, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Printf("无效的ID: %s\n", args[0])
        return
    }
    
    // 查找待办事项
    todo := todoList.FindByID(id)
    if todo == nil {
        fmt.Printf("找不到ID为 %d 的待办事项\n", id)
        return
    }
    
    // 检查是否已完成
    if todo.Completed {
        fmt.Printf("ℹ️  待办事项 [%d] 已经是完成状态\n", id)
        fmt.Printf("   标题: %s\n", todo.Title)
        return
    }
    
    // 标记为完成
    if todoList.MarkDone(id) {
        // 保存更改
        if err := todoList.Save(); err != nil {
            fmt.Printf("保存失败: %v\n", err)
            return
        }
        
        fmt.Printf("待办事项已标记为完成!\n")
        fmt.Printf("   ID: %d\n", todo.ID)
        fmt.Printf("   标题: %s\n", todo.Title)
        if todo.Description != "" {
            fmt.Printf("   描述: %s\n", todo.Description)
        }
        fmt.Printf("   完成时间: %s\n", todo.UpdatedAt.Format("2006-01-02 15:04:05"))
    }
}
```

### 运行结果
```bash
$ go run main.go done 1
TODO CLI 应用启动
加载现有待办事项: 2 个任务
待办事项已标记为完成!
   ID: 1
   标题: 学习Go语言
   描述: 完成Go语言基础教程
   完成时间: 2024-01-15 10:35:22

$ go run main.go done 1
TODO CLI 应用启动
加载现有待办事项: 2 个任务
ℹ️  待办事项 [1] 已经是完成状态
   标题: 学习Go语言

$ go run main.go done 99
TODO CLI 应用启动
加载现有待办事项: 2 个任务
找不到ID为 99 的待办事项
```

## 第六步：删除待办事项

### 解释
实现删除待办事项的功能，包括确认提示和安全删除机制。

### 编码

```go
// 处理删除命令
func handleRemove(todoList *TodoList, args []string) {
    if len(args) < 1 {
        fmt.Println("请提供待办事项ID")
        fmt.Println("使用方法: todo remove <id>")
        return
    }
    
    // 解析ID
    id, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Printf("无效的ID: %s\n", args[0])
        return
    }
    
    // 查找待办事项
    todo := todoList.FindByID(id)
    if todo == nil {
        fmt.Printf("找不到ID为 %d 的待办事项\n", id)
        return
    }
    
    // 显示要删除的事项信息
    fmt.Printf("⚠️  确认删除以下待办事项:\n")
    fmt.Printf("   ID: %d\n", todo.ID)
    fmt.Printf("   标题: %s\n", todo.Title)
    if todo.Description != "" {
        fmt.Printf("   描述: %s\n", todo.Description)
    }
    status := "待完成"
    if todo.Completed {
        status = "已完成"
    }
    fmt.Printf("   状态: %s\n", status)
    fmt.Printf("   创建时间: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
    
    // 确认删除
    fmt.Print("\n🤔 确定要删除吗? (y/N): ")
    var response string
    fmt.Scanln(&response)
    
    response = strings.ToLower(strings.TrimSpace(response))
    if response != "y" && response != "yes" {
        fmt.Println("删除操作已取消")
        return
    }
    
    // 执行删除
    if todoList.Remove(id) {
        // 保存更改
        if err := todoList.Save(); err != nil {
            fmt.Printf("保存失败: %v\n", err)
            return
        }
        
        fmt.Printf(" 待办事项已删除!\n")
        fmt.Printf("   剩余事项: %d 个\n", len(todoList.Todos))
    } else {
        fmt.Printf("删除失败\n")
    }
}
```

### 运行结果
```bash
$ go run main.go remove 2
TODO CLI 应用启动
加载现有待办事项: 2 个任务
⚠️  确认删除以下待办事项:
   ID: 2
   标题: 买菜
   状态: 待完成
   创建时间: 2024-01-15 10:31:20

🤔 确定要删除吗? (y/N): y
 待办事项已删除!
   剩余事项: 1 个

$ go run main.go list
TODO CLI 应用启动
加载现有待办事项: 1 个任务
所有待办事项 (1个):

[1] 学习Go语言
     完成Go语言基础教程
    创建: 2024-01-15 10:30
    更新: 2024-01-15 10:35

统计信息:
   总计: 1 个事项
   已完成: 1 个事项
   待完成: 0 个事项
   完成率: 100.0%
```

## 第七步：增强功能

### 解释
添加一些增强功能，包括搜索、编辑和批量操作，使CLI工具更加实用。

### 编码

```go
// 在main函数的switch语句中添加新命令
func main() {
    // ... 前面的代码保持不变
    
    switch command {
    case "add", "a":
        handleAdd(todoList, os.Args[2:])
    case "list", "l":
        handleList(todoList, os.Args[2:])
    case "done", "d":
        handleDone(todoList, os.Args[2:])
    case "remove", "r":
        handleRemove(todoList, os.Args[2:])
    case "search", "s":
        handleSearch(todoList, os.Args[2:])
    case "edit", "e":
        handleEdit(todoList, os.Args[2:])
    case "clear":
        handleClear(todoList, os.Args[2:])
    case "help", "h":
        showHelp()
    default:
        fmt.Printf("未知命令: %s\n", command)
        showHelp()
    }
}

// 搜索待办事项
func handleSearch(todoList *TodoList, args []string) {
    if len(args) < 1 {
        fmt.Println("请提供搜索关键词")
        fmt.Println("使用方法: todo search <keyword>")
        return
    }
    
    keyword := strings.ToLower(args[0])
    var matches []Todo
    
    for _, todo := range todoList.Todos {
        titleMatch := strings.Contains(strings.ToLower(todo.Title), keyword)
        descMatch := strings.Contains(strings.ToLower(todo.Description), keyword)
        
        if titleMatch || descMatch {
            matches = append(matches, todo)
        }
    }
    
    if len(matches) == 0 {
        fmt.Printf("没有找到包含 '%s' 的待办事项\n", keyword)
        return
    }
    
    fmt.Printf("搜索结果 - 找到 %d 个匹配项:\n\n", len(matches))
    for _, todo := range matches {
        displayTodo(todo)
        fmt.Println()
    }
}

// 编辑待办事项
func handleEdit(todoList *TodoList, args []string) {
    if len(args) < 2 {
        fmt.Println("请提供ID和新标题")
        fmt.Println("使用方法: todo edit <id> <new_title> [new_description]")
        return
    }
    
    id, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Printf("无效的ID: %s\n", args[0])
        return
    }
    
    todo := todoList.FindByID(id)
    if todo == nil {
        fmt.Printf("找不到ID为 %d 的待办事项\n", id)
        return
    }
    
    oldTitle := todo.Title
    oldDescription := todo.Description
    
    todo.Title = args[1]
    if len(args) > 2 {
        todo.Description = args[2]
    }
    todo.UpdatedAt = time.Now()
    
    if err := todoList.Save(); err != nil {
        fmt.Printf("保存失败: %v\n", err)
        return
    }
    
    fmt.Printf("✏️  待办事项已更新:\n")
    fmt.Printf("   ID: %d\n", todo.ID)
    fmt.Printf("   标题: %s → %s\n", oldTitle, todo.Title)
    if oldDescription != todo.Description {
        fmt.Printf("   描述: %s → %s\n", oldDescription, todo.Description)
    }
    fmt.Printf("   更新时间: %s\n", todo.UpdatedAt.Format("2006-01-02 15:04:05"))
}

// 清空已完成的事项
func handleClear(todoList *TodoList, args []string) {
    completedCount := 0
    for _, todo := range todoList.Todos {
        if todo.Completed {
            completedCount++
        }
    }
    
    if completedCount == 0 {
        fmt.Println("ℹ️  没有已完成的待办事项需要清理")
        return
    }
    
    fmt.Printf("⚠️  将删除 %d 个已完成的待办事项\n", completedCount)
    fmt.Print("🤔 确定要清空吗? (y/N): ")
    
    var response string
    fmt.Scanln(&response)
    
    response = strings.ToLower(strings.TrimSpace(response))
    if response != "y" && response != "yes" {
        fmt.Println("清空操作已取消")
        return
    }
    
    // 过滤掉已完成的事项
    var newTodos []Todo
    for _, todo := range todoList.Todos {
        if !todo.Completed {
            newTodos = append(newTodos, todo)
        }
    }
    
    todoList.Todos = newTodos
    
    if err := todoList.Save(); err != nil {
        fmt.Printf("保存失败: %v\n", err)
        return
    }
    
    fmt.Printf("🧹 已清空 %d 个已完成的待办事项\n", completedCount)
    fmt.Printf("   剩余事项: %d 个\n", len(todoList.Todos))
}

// 更新帮助信息
func showHelp() {
    fmt.Println(`
TODO CLI 使用帮助

命令格式: todo <command> [arguments]

基础命令:
  add, a     <title> [description]    添加新的待办事项
  list, l    [status]                 列出待办事项 (all/pending/done)
  done, d    <id>                     标记待办事项为完成
  remove, r  <id>                     删除待办事项

增强命令:
  search, s  <keyword>                搜索待办事项
  edit, e    <id> <title> [desc]      编辑待办事项
  clear                               清空已完成的事项
  help, h                             显示此帮助信息

使用示例:
  todo add "学习Go语言" "完成Go语言基础教程"
  todo list pending
  todo search "Go"
  todo edit 1 "深入学习Go语言" "包括高级特性"
  todo done 1
  todo clear
    `)
}
```

### 运行结果
```bash
$ go run main.go add "写周报" "总结本周工作内容"
已添加待办事项:
   ID: 2
   标题: 写周报
   描述: 总结本周工作内容
   创建时间: 2024-01-15 11:00:15

$ go run main.go search "Go"
搜索结果 - 找到 1 个匹配项:

[1] 学习Go语言
     完成Go语言基础教程
    创建: 2024-01-15 10:30
    更新: 2024-01-15 10:35

$ go run main.go edit 2 "写技术周报" "总结Go语言学习进展"
✏️  待办事项已更新:
   ID: 2
   标题: 写周报 → 写技术周报
   描述: 总结本周工作内容 → 总结Go语言学习进展
   更新时间: 2024-01-15 11:05:30

$ go run main.go clear
⚠️  将删除 1 个已完成的待办事项
🤔 确定要清空吗? (y/N): y
🧹 已清空 1 个已完成的待办事项
   剩余事项: 1 个
```

## 第八步：构建和分发

### 解释
最后一步是构建可执行文件并创建安装脚本，方便在不同系统上使用。

### 编码

创建构建脚本 `build.sh`：

```bash
#!/bin/bash

echo "🔨 构建 TODO CLI 应用..."

# 设置版本信息
VERSION="1.0.0"
BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 构建信息
LDFLAGS="-X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GitCommit=${GIT_COMMIT}'"

# 为不同平台构建
echo "为不同平台构建..."

# Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/todo-linux-amd64 .
echo "Linux AMD64 构建完成"

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/todo-windows-amd64.exe .
echo "Windows AMD64 构建完成"

# macOS AMD64
GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/todo-darwin-amd64 .
echo "macOS AMD64 构建完成"

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -o dist/todo-darwin-arm64 .
echo "macOS ARM64 构建完成"

echo "所有平台构建完成！"
echo "构建文件位于 dist/ 目录"
```

添加版本信息到 `main.go`：

```go
// 版本信息（构建时注入）
var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

// 在main函数的switch中添加version命令
case "version", "v":
    showVersion()

// 显示版本信息
func showVersion() {
    fmt.Printf("TODO CLI 版本信息:\n")
    fmt.Printf("   版本: %s\n", Version)
    fmt.Printf("   构建时间: %s\n", BuildTime)
    fmt.Printf("   Git提交: %s\n", GitCommit)
    fmt.Printf("   Go版本: %s\n", runtime.Version())
    fmt.Printf("   操作系统: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
```

### 运行结果
```bash
$ chmod +x build.sh
$ ./build.sh
🔨 构建 TODO CLI 应用...
为不同平台构建...
Linux AMD64 构建完成
Windows AMD64 构建完成
macOS AMD64 构建完成
macOS ARM64 构建完成
所有平台构建完成！
构建文件位于 dist/ 目录

$ ./dist/todo-darwin-amd64 version
TODO CLI 版本信息:
   版本: 1.0.0
   构建时间: 2024-01-15 11:15:30
   Git提交: abc1234
   Go版本: go1.21.5
   操作系统: darwin/amd64
```

##  项目总结

### 完成的功能
- 添加、列出、完成、删除待办事项
- 搜索和编辑功能
- 数据持久化存储
- 彩色命令行输出
- 多平台构建支持

### 技术要点
- **结构体设计**: 合理的数据模型和方法组织
- **JSON处理**: 文件读写和数据序列化
- **命令行解析**: 参数验证和命令路由
- **错误处理**: 优雅的错误处理和用户反馈
- **时间处理**: 时间格式化和时间戳管理

### 可扩展功能
- 添加配置文件支持
- 实现任务优先级
- 添加任务标签功能
- 支持任务提醒
- 实现数据导入导出
- 添加任务统计图表

这个项目展示了Go语言在构建实用CLI工具方面的强大能力，涵盖了文件操作、JSON处理、命令行交互等核心技能。 