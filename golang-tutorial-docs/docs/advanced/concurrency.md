---
title: 并发编程
description: 学习Go语言的goroutines、channels和并发编程模式
---

# 并发编程

Go语言的并发编程是其最重要的特性之一。通过goroutines和channels，Go让并发编程变得简单而优雅。

## 📖 本章内容

- Goroutines 轻量级协程
- Channels 通道通信
- Select 语句和多路复用
- 并发安全和同步机制
- 常见并发模式和最佳实践

## 🚀 Goroutines 基础

### 什么是 Goroutine

Goroutine 是Go语言的轻量级线程，由Go运行时管理。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("主goroutine开始")
    
    // 启动新的goroutine
    go sayHello("世界")
    go sayHello("Go语言")
    
    // 等待一段时间，让goroutines执行
    time.Sleep(2 * time.Second)
    fmt.Println("主goroutine结束")
}

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i+1)
        time.Sleep(500 * time.Millisecond)
    }
}
```

### Goroutine 池模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Worker 工作者函数
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d 开始处理任务 %d\n", id, job)
        
        // 模拟工作
        time.Sleep(time.Second)
        result := job * 2
        
        fmt.Printf("Worker %d 完成任务 %d，结果: %d\n", id, job, result)
        results <- result
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 5
    
    // 创建通道
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    var wg sync.WaitGroup
    
    // 启动工作者goroutines
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // 发送任务
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 等待所有工作者完成
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // 收集结果
    fmt.Println("\n任务结果:")
    for result := range results {
        fmt.Printf("结果: %d\n", result)
    }
}
```

## 📡 Channels 通道

### 基本通道操作

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建无缓冲通道
    messages := make(chan string)
    
    // 发送数据到通道（在goroutine中）
    go func() {
        messages <- "Hello"
        messages <- "World"
        messages <- "Go"
        close(messages) // 关闭通道
    }()
    
    // 从通道接收数据
    for msg := range messages {
        fmt.Printf("收到消息: %s\n", msg)
    }
    
    // 有缓冲通道
    buffered := make(chan int, 3)
    buffered <- 1
    buffered <- 2
    buffered <- 3
    
    fmt.Printf("缓冲通道长度: %d\n", len(buffered))
    fmt.Printf("缓冲通道容量: %d\n", cap(buffered))
    
    // 读取缓冲通道
    for i := 0; i < 3; i++ {
        value := <-buffered
        fmt.Printf("从缓冲通道读取: %d\n", value)
    }
}
```

### 通道方向

```go
package main

import (
    "fmt"
    "time"
)

// 只发送通道
func sender(ch chan<- string) {
    for i := 0; i < 3; i++ {
        message := fmt.Sprintf("消息 %d", i+1)
        ch <- message
        time.Sleep(500 * time.Millisecond)
    }
    close(ch)
}

// 只接收通道
func receiver(ch <-chan string) {
    for message := range ch {
        fmt.Printf("接收到: %s\n", message)
    }
}

func main() {
    // 双向通道
    ch := make(chan string)
    
    // 启动发送者和接收者
    go sender(ch)
    go receiver(ch)
    
    // 等待完成
    time.Sleep(3 * time.Second)
}
```

## 🔀 Select 语句

### 基本 Select 用法

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    // 第一个goroutine
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "来自c1的消息"
    }()
    
    // 第二个goroutine
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "来自c2的消息"
    }()
    
    // 使用select等待多个通道
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Printf("收到: %s\n", msg1)
        case msg2 := <-c2:
            fmt.Printf("收到: %s\n", msg2)
        }
    }
}
```

### Select 超时处理

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    // 模拟一个可能很慢的操作
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "操作完成"
    }()
    
    // 使用select实现超时
    select {
    case result := <-ch:
        fmt.Printf("收到结果: %s\n", result)
    case <-time.After(2 * time.Second):
        fmt.Println("操作超时！")
    }
    
    // 非阻塞检查
    select {
    case msg := <-ch:
        fmt.Printf("非阻塞收到: %s\n", msg)
    default:
        fmt.Println("通道中没有数据")
    }
}
```

### Select 实现扇入模式

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// 扇入函数：将多个通道合并为一个
func fanIn(input1, input2 <-chan string) <-chan string {
    output := make(chan string)
    
    go func() {
        for {
            select {
            case s := <-input1:
                output <- s
            case s := <-input2:
                output <- s
            }
        }
    }()
    
    return output
}

// 生成器函数
func generator(name string) <-chan string {
    ch := make(chan string)
    
    go func() {
        for i := 0; ; i++ {
            ch <- fmt.Sprintf("%s: %d", name, i)
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        }
    }()
    
    return ch
}

func main() {
    // 创建两个生成器
    gen1 := generator("生成器1")
    gen2 := generator("生成器2")
    
    // 扇入合并
    merged := fanIn(gen1, gen2)
    
    // 接收合并后的数据
    for i := 0; i < 10; i++ {
        fmt.Println(<-merged)
    }
}
```

## 🔒 并发安全

### Mutex 互斥锁

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// 安全的计数器
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

// 增加计数
func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// 获取值
func (c *SafeCounter) GetValue() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup
    
    // 启动多个goroutines并发增加计数
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            for j := 0; j < 100; j++ {
                counter.Increment()
            }
            
            fmt.Printf("Goroutine %d 完成\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("最终计数: %d\n", counter.GetValue())
}
```

### RWMutex 读写锁

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// 缓存结构
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

// 新建缓存
func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

// 写入数据
func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    fmt.Printf("写入: %s = %s\n", key, value)
    c.data[key] = value
    time.Sleep(100 * time.Millisecond) // 模拟写入耗时
}

// 读取数据
func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    value, exists := c.data[key]
    fmt.Printf("读取: %s = %s (存在: %t)\n", key, value, exists)
    time.Sleep(50 * time.Millisecond) // 模拟读取耗时
    
    return value, exists
}

func main() {
    cache := NewCache()
    var wg sync.WaitGroup
    
    // 启动写入goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", id)
            value := fmt.Sprintf("value%d", id)
            cache.Set(key, value)
        }(i)
    }
    
    // 等待写入完成
    time.Sleep(500 * time.Millisecond)
    
    // 启动多个读取goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", id%3)
            cache.Get(key)
        }(i)
    }
    
    wg.Wait()
}
```

### 原子操作

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var counter int64
    var wg sync.WaitGroup
    
    // 启动多个goroutines进行原子操作
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            for j := 0; j < 1000; j++ {
                atomic.AddInt64(&counter, 1)
            }
        }()
    }
    
    wg.Wait()
    
    fmt.Printf("原子操作最终计数: %d\n", atomic.LoadInt64(&counter))
    
    // 原子操作的其他用法
    var flag int32
    
    // 原子设置
    atomic.StoreInt32(&flag, 1)
    fmt.Printf("标志值: %d\n", atomic.LoadInt32(&flag))
    
    // 原子交换
    old := atomic.SwapInt32(&flag, 2)
    fmt.Printf("交换前: %d, 交换后: %d\n", old, atomic.LoadInt32(&flag))
    
    // 原子比较并交换
    swapped := atomic.CompareAndSwapInt32(&flag, 2, 3)
    fmt.Printf("CAS成功: %t, 当前值: %d\n", swapped, atomic.LoadInt32(&flag))
}
```

## 🎯 常见并发模式

### 生产者-消费者模式

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// 产品结构
type Product struct {
    ID   int
    Name string
}

// 生产者
func producer(products chan<- Product, wg *sync.WaitGroup) {
    defer wg.Done()
    defer close(products)
    
    for i := 1; i <= 10; i++ {
        product := Product{
            ID:   i,
            Name: fmt.Sprintf("产品-%d", i),
        }
        
        fmt.Printf("生产: %s\n", product.Name)
        products <- product
        
        // 随机生产时间
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    }
    
    fmt.Println("生产者完成工作")
}

// 消费者
func consumer(id int, products <-chan Product, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for product := range products {
        fmt.Printf("消费者%d 消费: %s\n", id, product.Name)
        
        // 随机消费时间
        time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
    }
    
    fmt.Printf("消费者%d 完成工作\n", id)
}

func main() {
    // 创建缓冲通道
    products := make(chan Product, 5)
    var wg sync.WaitGroup
    
    // 启动生产者
    wg.Add(1)
    go producer(products, &wg)
    
    // 启动多个消费者
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go consumer(i, products, &wg)
    }
    
    wg.Wait()
    fmt.Println("所有工作完成")
}
```

### 管道模式

```go
package main

import (
    "fmt"
    "sync"
)

// 生成数字
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// 平方计算
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// 过滤偶数
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

// 扇出：将一个通道分发到多个通道
func fanOut(in <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        out := make(chan int)
        outputs[i] = out
        
        go func() {
            defer close(out)
            for n := range in {
                out <- n * n * n // 立方计算
            }
        }()
    }
    
    return outputs
}

// 扇入：将多个通道合并为一个
func fanIn(inputs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    // 为每个输入通道启动一个goroutine
    multiplex := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            out <- n
        }
    }
    
    wg.Add(len(inputs))
    for _, c := range inputs {
        go multiplex(c)
    }
    
    // 关闭输出通道
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

func main() {
    fmt.Println("=== 基本管道模式 ===")
    
    // 基本管道：生成 -> 平方 -> 过滤偶数
    numbers := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    squared := square(numbers)
    evens := filterEven(squared)
    
    for result := range evens {
        fmt.Printf("结果: %d\n", result)
    }
    
    fmt.Println("\n=== 扇出扇入模式 ===")
    
    // 扇出扇入模式
    input := generator(1, 2, 3, 4, 5)
    
    // 扇出到3个工作者
    workers := fanOut(input, 3)
    
    // 扇入合并结果
    result := fanIn(workers...)
    
    for r := range result {
        fmt.Printf("立方结果: %d\n", r)
    }
}
```

### 限流器模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// 限流器结构
type RateLimiter struct {
    tokens chan struct{}
    ticker *time.Ticker
}

// 创建限流器
func NewRateLimiter(rate int, burst int) *RateLimiter {
    rl := &RateLimiter{
        tokens: make(chan struct{}, burst),
        ticker: time.NewTicker(time.Second / time.Duration(rate)),
    }
    
    // 初始填满tokens
    for i := 0; i < burst; i++ {
        rl.tokens <- struct{}{}
    }
    
    // 定期添加token
    go func() {
        for range rl.ticker.C {
            select {
            case rl.tokens <- struct{}{}:
            default:
                // tokens已满，丢弃
            }
        }
    }()
    
    return rl
}

// 获取令牌
func (rl *RateLimiter) Take() {
    <-rl.tokens
}

// 停止限流器
func (rl *RateLimiter) Stop() {
    rl.ticker.Stop()
}

// 模拟API请求
func makeRequest(id int, limiter *RateLimiter, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("请求%d 等待令牌...\n", id)
    limiter.Take() // 获取令牌
    
    fmt.Printf("请求%d 开始处理\n", id)
    time.Sleep(100 * time.Millisecond) // 模拟处理时间
    fmt.Printf("请求%d 处理完成\n", id)
}

func main() {
    // 创建限流器：每秒2个请求，突发最多5个
    limiter := NewRateLimiter(2, 5)
    defer limiter.Stop()
    
    var wg sync.WaitGroup
    
    // 模拟10个并发请求
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go makeRequest(i, limiter, &wg)
    }
    
    wg.Wait()
    fmt.Println("所有请求完成")
}
```

## 🎯 实践练习

让我们创建一个完整的并发Web爬虫：

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"
)

// 爬虫结果
type CrawlResult struct {
    URL    string
    Status int
    Size   int
    Error  error
}

// Web爬虫
type WebCrawler struct {
    maxWorkers   int
    maxRetries   int
    timeout      time.Duration
    rateLimiter  chan struct{}
}

// 创建爬虫
func NewWebCrawler(maxWorkers, maxRetries int, timeout time.Duration) *WebCrawler {
    return &WebCrawler{
        maxWorkers:  maxWorkers,
        maxRetries:  maxRetries,
        timeout:     timeout,
        rateLimiter: make(chan struct{}, maxWorkers),
    }
}

// 爬取单个URL
func (wc *WebCrawler) crawlURL(url string) CrawlResult {
    // 限流
    wc.rateLimiter <- struct{}{}
    defer func() { <-wc.rateLimiter }()
    
    client := &http.Client{
        Timeout: wc.timeout,
    }
    
    var lastErr error
    for attempt := 0; attempt <= wc.maxRetries; attempt++ {
        if attempt > 0 {
            fmt.Printf("重试 %s (第%d次)\n", url, attempt)
            time.Sleep(time.Duration(attempt) * time.Second)
        }
        
        resp, err := client.Get(url)
        if err != nil {
            lastErr = err
            continue
        }
        
        defer resp.Body.Close()
        
        // 读取响应体大小
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            lastErr = err
            continue
        }
        
        return CrawlResult{
            URL:    url,
            Status: resp.StatusCode,
            Size:   len(body),
            Error:  nil,
        }
    }
    
    return CrawlResult{
        URL:   url,
        Error: lastErr,
    }
}

// 并发爬取多个URL
func (wc *WebCrawler) CrawlURLs(urls []string) []CrawlResult {
    urlChan := make(chan string, len(urls))
    resultChan := make(chan CrawlResult, len(urls))
    
    var wg sync.WaitGroup
    
    // 启动工作者goroutines
    for i := 0; i < wc.maxWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            for url := range urlChan {
                fmt.Printf("Worker %d 爬取: %s\n", workerID, url)
                result := wc.crawlURL(url)
                resultChan <- result
            }
        }(i)
    }
    
    // 发送URLs到通道
    go func() {
        for _, url := range urls {
            urlChan <- url
        }
        close(urlChan)
    }()
    
    // 等待所有工作者完成
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    // 收集结果
    var results []CrawlResult
    for result := range resultChan {
        results = append(results, result)
    }
    
    return results
}

func main() {
    // 要爬取的URL列表
    urls := []string{
        "https://httpbin.org/delay/1",
        "https://httpbin.org/delay/2",
        "https://httpbin.org/status/200",
        "https://httpbin.org/status/404",
        "https://httpbin.org/json",
        "https://httpbin.org/xml",
        "https://httpbin.org/html",
        "https://httpbin.org/robots.txt",
    }
    
    // 创建爬虫：3个工作者，最多重试2次，超时10秒
    crawler := NewWebCrawler(3, 2, 10*time.Second)
    
    fmt.Printf("开始爬取 %d 个URL...\n", len(urls))
    start := time.Now()
    
    // 执行爬取
    results := crawler.CrawlURLs(urls)
    
    duration := time.Since(start)
    
    // 输出结果
    fmt.Printf("\n爬取完成，耗时: %v\n", duration)
    fmt.Println("结果统计:")
    
    successCount := 0
    for _, result := range results {
        if result.Error != nil {
            fmt.Printf("❌ %s: %v\n", result.URL, result.Error)
        } else {
            fmt.Printf("✅ %s: %d (%d bytes)\n", 
                result.URL, result.Status, result.Size)
            successCount++
        }
    }
    
    fmt.Printf("\n成功: %d/%d\n", successCount, len(results))
}
```

## 📝 本章小结

在这一章中，我们学习了：

### 🔹 Goroutines
- 轻量级协程的创建和使用
- Goroutine池模式
- 生命周期管理

### 🔹 Channels
- 通道的创建和操作
- 缓冲和无缓冲通道
- 通道方向和关闭

### 🔹 Select语句
- 多通道选择
- 超时处理
- 非阻塞操作

### 🔹 并发安全
- Mutex和RWMutex
- 原子操作
- 数据竞争防护

### 🔹 并发模式
- 扇入扇出模式
- 生产者消费者模式
- 管道模式和限流器

## 🎯 下一步

掌握了并发编程后，让我们继续学习 [文件操作](./file-operations)，了解如何处理文件和数据！ 