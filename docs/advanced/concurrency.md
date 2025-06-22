---
title: 并发编程
description: 学习Go语言的goroutines、channels和并发编程模式
---

# 并发编程

并发编程是Go语言的核心特性和最大优势。Go通过goroutines和channels提供了简洁而强大的并发编程模型，让并发编程变得直观和安全。

## 本章内容

- Goroutines 轻量级协程
- Channels 通道通信机制
- Select 语句和多路复用
- 并发安全和同步机制
- 并发设计模式和最佳实践

## 并发编程概念

### 为什么需要并发

并发编程允许程序同时处理多个任务，提高程序性能和响应能力：

- **性能提升**：充分利用多核CPU资源
- **用户体验**：避免阻塞，提高响应速度
- **资源效率**：更好地利用I/O等待时间
- **可扩展性**：支持大规模并发处理

### Go并发模型特点

| 特性 | 说明 | 优势 |
|------|------|------|
| **轻量级** | Goroutine只需2KB内存 | 可创建百万级协程 |
| **通信共享** | 通过channels传递数据 | 避免共享内存竞争 |
| **调度器** | M:N调度模型 | 高效的协程调度 |
| **内置支持** | 语言级别的并发支持 | 简洁的语法和API |

::: tip 并发哲学
Go遵循CSP(Communicating Sequential Processes)模型：
**"不要通过共享内存来通信，通过通信来共享内存"**
:::

## Goroutines 协程

### 基础用法

```go
func main() {
    fmt.Println("主程序开始")
    
    // 启动goroutine
    go sayHello("World")
    go sayHello("Go")
    
    // 等待一段时间让goroutines执行
    time.Sleep(2 * time.Second)
    fmt.Println("主程序结束")
}

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i+1)
        time.Sleep(500 * time.Millisecond)
    }
}
```

### 等待组 WaitGroup

使用`sync.WaitGroup`等待所有goroutines完成：

```go
func main() {
    var wg sync.WaitGroup
    
    tasks := []string{"任务A", "任务B", "任务C"}
    
    for _, task := range tasks {
        wg.Add(1) // 增加计数
        go func(name string) {
            defer wg.Done() // 完成时减少计数
            processTask(name)
        }(task)
    }
    
    wg.Wait() // 等待所有任务完成
    fmt.Println("所有任务完成")
}

func processTask(name string) {
    fmt.Printf("开始执行 %s\n", name)
    time.Sleep(time.Second)
    fmt.Printf("完成执行 %s\n", name)
}
```

## Channels 通道

### 基础通道操作

Channels是goroutines之间通信的管道：

```go
func main() {
    // 创建通道
    messages := make(chan string)
    
    // 发送数据（在goroutine中）
    go func() {
        messages <- "Hello"
        messages <- "World"
        close(messages) // 关闭通道
    }()
    
    // 接收数据
    for msg := range messages {
        fmt.Printf("收到: %s\n", msg)
    }
}
```

### 通道类型和特性

```go
// 无缓冲通道 - 同步通信
unbuffered := make(chan int)

// 有缓冲通道 - 异步通信
buffered := make(chan int, 3)

// 只发送通道
func sender(ch chan<- string) {
    ch <- "message"
}

// 只接收通道
func receiver(ch <-chan string) {
    msg := <-ch
    fmt.Println(msg)
}
```

### Select 多路复用

Select语句让goroutine等待多个通信操作：

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "来自ch1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "来自ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("收到", msg1)
        case msg2 := <-ch2:
            fmt.Println("收到", msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("超时")
            return
        }
    }
}
```

## 并发安全

### 互斥锁 Mutex

当必须共享内存时，使用互斥锁保护数据：

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

### 读写锁 RWMutex

读多写少的场景使用读写锁提高性能：

```go
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    value, ok := sm.data[key]
    return value, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}
```

## 实战项目：并发任务调度器

让我们创建一个实用的并发任务调度器：

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Task 任务接口
type Task interface {
    Execute(ctx context.Context) error
    GetID() string
    GetPriority() int
}

// EmailTask 邮件任务
type EmailTask struct {
    ID       string
    To       string
    Subject  string
    Priority int
}

func (e EmailTask) Execute(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(500 * time.Millisecond): // 模拟邮件发送
        fmt.Printf("📧 邮件已发送到 %s: %s\n", e.To, e.Subject)
        return nil
    }
}

func (e EmailTask) GetID() string       { return e.ID }
func (e EmailTask) GetPriority() int    { return e.Priority }

// FileTask 文件任务
type FileTask struct {
    ID       string
    FilePath string
    Action   string
    Priority int
}

func (f FileTask) Execute(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(200 * time.Millisecond): // 模拟文件操作
        fmt.Printf("📁 文件操作完成: %s %s\n", f.Action, f.FilePath)
        return nil
    }
}

func (f FileTask) GetID() string       { return f.ID }
func (f FileTask) GetPriority() int    { return f.Priority }

// TaskScheduler 任务调度器
type TaskScheduler struct {
    workers    int
    taskQueue  chan Task
    resultChan chan TaskResult
    ctx        context.Context
    cancel     context.CancelFunc
    wg         sync.WaitGroup
}

type TaskResult struct {
    TaskID string
    Error  error
    Duration time.Duration
}

func NewTaskScheduler(workers int, queueSize int) *TaskScheduler {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &TaskScheduler{
        workers:    workers,
        taskQueue:  make(chan Task, queueSize),
        resultChan: make(chan TaskResult, queueSize),
        ctx:        ctx,
        cancel:     cancel,
    }
}

func (ts *TaskScheduler) Start() {
    fmt.Printf("🚀 启动任务调度器，工作线程数: %d\n", ts.workers)
    
    // 启动工作goroutines
    for i := 0; i < ts.workers; i++ {
        ts.wg.Add(1)
        go ts.worker(i + 1)
    }
    
    // 启动结果处理goroutine
    go ts.resultHandler()
}

func (ts *TaskScheduler) worker(id int) {
    defer ts.wg.Done()
    
    for {
        select {
        case <-ts.ctx.Done():
            fmt.Printf("Worker %d 停止\n", id)
            return
        case task, ok := <-ts.taskQueue:
            if !ok {
                return
            }
            
            start := time.Now()
            fmt.Printf("Worker %d 开始执行任务 %s (优先级: %d)\n", 
                id, task.GetID(), task.GetPriority())
            
            err := task.Execute(ts.ctx)
            duration := time.Since(start)
            
            ts.resultChan <- TaskResult{
                TaskID:   task.GetID(),
                Error:    err,
                Duration: duration,
            }
        }
    }
}

func (ts *TaskScheduler) resultHandler() {
    for result := range ts.resultChan {
        if result.Error != nil {
            fmt.Printf("❌ 任务 %s 执行失败: %v (耗时: %v)\n", 
                result.TaskID, result.Error, result.Duration)
        } else {
            fmt.Printf("✅ 任务 %s 执行成功 (耗时: %v)\n", 
                result.TaskID, result.Duration)
        }
    }
}

func (ts *TaskScheduler) AddTask(task Task) {
    select {
    case ts.taskQueue <- task:
        fmt.Printf("➕ 任务 %s 已加入队列\n", task.GetID())
    case <-ts.ctx.Done():
        fmt.Println("调度器已停止，无法添加任务")
    default:
        fmt.Println("任务队列已满，任务被丢弃")
    }
}

func (ts *TaskScheduler) Stop() {
    fmt.Println("🛑 停止任务调度器...")
    
    close(ts.taskQueue)      // 关闭任务队列
    ts.wg.Wait()             // 等待所有worker完成
    close(ts.resultChan)     // 关闭结果通道
    ts.cancel()              // 取消上下文
    
    fmt.Println("任务调度器已停止")
}

func main() {
    // 创建调度器
    scheduler := NewTaskScheduler(3, 10)
    scheduler.Start()
    
    // 添加不同类型的任务
    tasks := []Task{
        EmailTask{
            ID: "email-1", To: "user1@example.com", 
            Subject: "欢迎邮件", Priority: 1,
        },
        FileTask{
            ID: "file-1", FilePath: "/data/backup.zip", 
            Action: "压缩", Priority: 2,
        },
        EmailTask{
            ID: "email-2", To: "user2@example.com", 
            Subject: "通知邮件", Priority: 1,
        },
        FileTask{
            ID: "file-2", FilePath: "/logs/app.log", 
            Action: "清理", Priority: 3,
        },
        EmailTask{
            ID: "email-urgent", To: "admin@example.com", 
            Subject: "紧急通知", Priority: 0,
        },
    }
    
    // 提交任务
    for _, task := range tasks {
        scheduler.AddTask(task)
        time.Sleep(100 * time.Millisecond) // 模拟任务提交间隔
    }
    
    // 等待一段时间让任务执行
    time.Sleep(3 * time.Second)
    
    // 停止调度器
    scheduler.Stop()
}
```

## 并发模式

### 工作池模式

```go
func WorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- job * 2 // 处理任务
            }
        }()
    }
    
    wg.Wait()
    close(results)
}
```

### 扇出/扇入模式

```go
// 扇出：一个输入分发到多个处理器
func FanOut(input <-chan int, output1, output2 chan<- int) {
    for data := range input {
        select {
        case output1 <- data:
        case output2 <- data:
        }
    }
    close(output1)
    close(output2)
}

// 扇入：多个输入合并到一个输出
func FanIn(input1, input2 <-chan int, output chan<- int) {
    var wg sync.WaitGroup
    
    wg.Add(2)
    go func() {
        defer wg.Done()
        for data := range input1 {
            output <- data
        }
    }()
    
    go func() {
        defer wg.Done()
        for data := range input2 {
            output <- data
        }
    }()
    
    wg.Wait()
    close(output)
}
```

## 最佳实践

### 1. 避免过度并发

```go
// 限制并发数量
semaphore := make(chan struct{}, 10) // 最多10个并发

func processWithLimit() {
    semaphore <- struct{}{} // 获取许可
    defer func() { <-semaphore }() // 释放许可
    
    // 处理逻辑
}
```

### 2. 优雅关闭

```go
func gracefulShutdown() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // 监听系统信号
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-c
        fmt.Println("收到停止信号，开始优雅关闭...")
        cancel()
    }()
    
    // 使用context控制goroutines
}
```

### 3. 错误处理

```go
type Result struct {
    Value int
    Error error
}

func safeProcess(input int) Result {
    // 处理可能出错的操作
    if input < 0 {
        return Result{Error: fmt.Errorf("invalid input: %d", input)}
    }
    return Result{Value: input * 2}
}
```

## 本章小结

Go的并发编程核心要点：

- **Goroutines**：轻量级协程，Go并发的基础
- **Channels**：goroutines间的通信机制，实现"通过通信共享内存"
- **Select**：多通道操作的多路复用
- **同步原语**：必要时使用mutex、WaitGroup等确保安全
- **设计模式**：工作池、扇出扇入等常用并发模式

::: tip 练习建议
1. 实现一个并发的网页爬虫
2. 创建一个生产者-消费者模式的消息队列
3. 构建一个支持并发的缓存系统
4. 实验不同的并发模式，体验其适用场景
:::