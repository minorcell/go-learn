# 内存模型：并发世界的秩序

> Go 的内存模型定义了在并发环境中，一个 goroutine 对变量的读取何时能观察到另一个 goroutine 对同一变量的写入。这听起来简单，实际却是现代并发编程的核心挑战。

## 为什么需要内存模型？

在单线程程序中，内存操作的顺序是显而易见的：按照程序的书写顺序执行。但在并发环境中，事情变得复杂：

```go
// 这段代码的行为在并发环境中是未定义的
var a, b int

func goroutine1() {
    a = 1
    b = 2
}

func goroutine2() {
    print(b)  // 可能输出 0 或 2
    print(a)  // 可能输出 0 或 1
}
```

问题的根源在于：
- **编译器优化**：可能重排指令顺序
- **CPU 乱序执行**：现代处理器为了性能可能不按顺序执行指令
- **内存层次结构**：缓存、写缓冲区等导致内存操作延迟可见

Go 的内存模型为这种混乱建立了秩序。

## happens-before 关系

Go 内存模型的核心是 "happens-before" 关系。如果事件 A happens-before 事件 B，那么 A 对内存的修改在 B 开始时就是可见的。

### 程序顺序规则

在单个 goroutine 内，happens-before 关系遵循程序顺序：

```go
func singleGoroutine() {
    a := 1        // 事件 1
    b := a + 1    // 事件 2：happens-after 事件 1
    c := b * 2    // 事件 3：happens-after 事件 2
    
    // 在这个 goroutine 内，执行顺序是明确的
    fmt.Println(c)  // 总是输出 4
}
```

但跨 goroutine 的情况就复杂了：

```go
var data int
var flag bool

func writer() {
    data = 42    // 写入数据
    flag = true  // 设置标志
}

func reader() {
    if flag {           // 读取标志
        fmt.Println(data)  // 读取数据 - 不保证能看到 42！
    }
}
```

没有同步机制，`reader` 可能看到 `flag = true` 但 `data` 仍然是 0。

## Channel 的同步语义

Channel 是 Go 提供的主要同步原语，它具有明确的 happens-before 语义：

### 无缓冲 Channel

```go
func demonstrateUnbufferedChannel() {
    ch := make(chan int)
    var data int
    
    go func() {
        data = 42      // 写入数据
        ch <- 1        // 发送到 channel
    }()
    
    <-ch               // 从 channel 接收
    fmt.Println(data)  // 保证能看到 42
}
```

**关键语义**：对无缓冲 channel 的发送操作 happens-before 对应的接收操作完成。

这意味着：发送方在 `ch <- 1` 之前对内存的所有修改，在接收方执行 `<-ch` 后都是可见的。

### 有缓冲 Channel

```go
func demonstrateBufferedChannel() {
    ch := make(chan int, 1)  // 缓冲大小为 1
    var data int
    
    go func() {
        data = 42      // 写入数据
        ch <- 1        // 发送到 channel（不阻塞）
    }()
    
    <-ch               // 从 channel 接收
    fmt.Println(data)  // 仍然保证能看到 42
}
```

**关键语义**：对缓冲 channel 的第 k 个接收操作 happens-before 第 k+C 个发送操作完成（C 是缓冲大小）。

### Channel 关闭的语义

```go
func demonstrateChannelClose() {
    ch := make(chan int)
    var data int
    
    go func() {
        data = 42     // 写入数据
        close(ch)     // 关闭 channel
    }()
    
    <-ch              // 接收到零值和 false
    fmt.Println(data) // 保证能看到 42
}
```

**关键语义**：关闭 channel happens-before 该关闭导致的接收操作返回零值。

## Mutex 的同步语义

`sync.Mutex` 提供了传统的互斥锁语义：

```go
import "sync"

var mu sync.Mutex
var data int

func writeWithMutex() {
    mu.Lock()
    data = 42      // 临界区
    mu.Unlock()
}

func readWithMutex() {
    mu.Lock()
    value := data  // 临界区
    mu.Unlock()
    fmt.Println(value)
}
```

**关键语义**：
- 对 `Mutex.Unlock()` 的调用 happens-before 对同一 `Mutex.Lock()` 的调用返回
- 第 n 次 `Unlock()` happens-before 第 m 次 `Lock()` 返回（其中 m > n）

这保证了互斥访问：同一时间只有一个 goroutine 可以持有锁。

### RWMutex 的复杂语义

```go
var rwmu sync.RWMutex
var sharedData []int

func writer() {
    rwmu.Lock()         // 写锁
    sharedData = append(sharedData, 42)
    rwmu.Unlock()
}

func reader() {
    rwmu.RLock()        // 读锁
    value := len(sharedData)
    rwmu.RUnlock()
    fmt.Println(value)
}
```

**关键语义**：
- `RWMutex.RUnlock()` happens-before 后续的 `RWMutex.Lock()` 返回
- `RWMutex.Unlock()` happens-before 后续的 `RWMutex.RLock()` 返回

## sync.Once 的初始化语义

```go
var once sync.Once
var config *Config

func getConfig() *Config {
    once.Do(func() {
        config = loadConfigFromFile()  // 只执行一次
    })
    return config
}
```

**关键语义**：`once.Do(f)` 中函数 `f` 的完成 happens-before 任何 `once.Do(f)` 调用的返回。

这保证了 `config` 的初始化只发生一次，且所有后续的 `getConfig()` 调用都能看到完全初始化的 `config`。

## 原子操作的语义

`sync/atomic` 包提供了原子操作，它们有特殊的内存语义：

```go
import "sync/atomic"

var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}

func read() int64 {
    return atomic.LoadInt64(&counter)
}
```

**关键语义**：原子操作具有顺序一致性语义，相当于有一个全局的操作顺序。

但要注意，原子操作只保证单个操作的原子性，不保证组合操作的原子性：

```go
// 这不是原子的！
func nonAtomicIncrement() {
    old := atomic.LoadInt64(&counter)  // 读取
    atomic.StoreInt64(&counter, old+1) // 写入 - 中间可能被其他 goroutine 修改
}
```

## 实际应用中的模式

### 发布-订阅模式

```go
type Publisher struct {
    mu          sync.RWMutex
    subscribers []chan<- int
}

func (p *Publisher) Subscribe() <-chan int {
    ch := make(chan int, 1)
    
    p.mu.Lock()
    p.subscribers = append(p.subscribers, ch)
    p.mu.Unlock()
    
    return ch
}

func (p *Publisher) Publish(value int) {
    p.mu.RLock()
    for _, ch := range p.subscribers {
        select {
        case ch <- value:
        default: // 非阻塞发送
        }
    }
    p.mu.RUnlock()
}
```

这个模式利用了 RWMutex 的语义：`RUnlock()` happens-before 后续的 `Lock()`，确保发布时看到的订阅者列表是一致的。

### 懒加载单例

```go
type Singleton struct {
    data string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        // 复杂的初始化逻辑
        instance = &Singleton{
            data: loadExpensiveData(),
        }
    })
    return instance
}
```

`sync.Once` 的语义保证初始化只发生一次，且所有 goroutine 都能看到完全初始化的实例。

### 优雅关闭模式

```go
type Server struct {
    quit    chan struct{}
    workers sync.WaitGroup
}

func (s *Server) Start() {
    for i := 0; i < 10; i++ {
        s.workers.Add(1)
        go s.worker()
    }
}

func (s *Server) worker() {
    defer s.workers.Done()
    
    for {
        select {
        case <-s.quit:
            return  // 收到关闭信号
        default:
            // 处理工作
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func (s *Server) Stop() {
    close(s.quit)    // 通知所有 worker 停止
    s.workers.Wait() // 等待所有 worker 完成
}
```

这个模式利用了 channel 关闭的语义和 `WaitGroup` 的同步性质。

## 常见的内存模型陷阱

### 双重检查锁定

```go
// ❌ 错误的双重检查锁定
var instance *Singleton
var mu sync.Mutex

func GetInstanceWrong() *Singleton {
    if instance == nil {  // 第一次检查（无锁）
        mu.Lock()
        if instance == nil {  // 第二次检查（有锁）
            instance = &Singleton{}  // 可能被其他 goroutine 看到部分初始化的对象
        }
        mu.Unlock()
    }
    return instance
}

// ✅ 正确的方式：使用 sync.Once
func GetInstanceCorrect() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

问题在于，即使在锁内初始化，其他 goroutine 可能看到 `instance != nil` 但对象还没有完全初始化。

### 数据竞争的检测

Go 提供了竞争检测器来帮助发现数据竞争：

```bash
go run -race main.go
go test -race
go build -race
```

```go
// 这段代码有数据竞争
var counter int

func race() {
    go func() {
        for i := 0; i < 1000; i++ {
            counter++  // 写入
        }
    }()
    
    go func() {
        for i := 0; i < 1000; i++ {
            fmt.Println(counter)  // 读取
        }
    }()
}
```

竞争检测器会报告这种未同步的并发访问。

## 性能考量

不同同步机制有不同的性能特征：

```go
// 性能从高到低
func benchmarkSynchronization() {
    // 1. 原子操作 - 最快
    atomic.AddInt64(&counter, 1)
    
    // 2. 无缓冲 channel - 中等
    ch <- 1
    <-ch
    
    // 3. Mutex - 较慢
    mu.Lock()
    // 临界区
    mu.Unlock()
    
    // 4. 有缓冲 channel - 更慢（但更灵活）
    bufferedCh <- 1
}
```

选择同步机制时要考虑：
- **性能需求**：原子操作最快，但功能有限
- **使用场景**：channel 适合通信，mutex 适合保护状态
- **复杂度**：简单场景用原子操作，复杂场景用 channel

## 内存模型的最佳实践

### 1. 优先使用 Channel

```go
// ✅ 推荐：使用 channel 通信
func coordinateWithChannel() {
    done := make(chan bool)
    
    go func() {
        doWork()
        done <- true
    }()
    
    <-done  // 等待工作完成
}
```

### 2. 必要时使用 Mutex

```go
// ✅ 保护共享状态时使用 mutex
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}
```

### 3. 避免数据竞争

```go
// ❌ 数据竞争
func dataRace() {
    var x int
    go func() { x = 1 }()
    go func() { x = 2 }()
}

// ✅ 使用同步
func noDataRace() {
    var x int
    var mu sync.Mutex
    
    go func() {
        mu.Lock()
        x = 1
        mu.Unlock()
    }()
    
    go func() {
        mu.Lock()
        x = 2
        mu.Unlock()
    }()
}
```

### 4. 理解语义差异

```go
// select 的非阻塞语义
select {
case ch <- value:
    // 发送成功
default:
    // channel 满了，立即执行这里
}

// 对比阻塞发送
ch <- value  // 会等待直到可以发送
```

## 下一步

内存模型是 Go 并发编程的理论基础，接下来让我们了解[编译器](/learn/concepts/compiler)，看看 Go 如何将我们的代码转换为高效的机器代码。

记住：内存模型不是用来背诵的规则，而是用来理解并发程序行为的工具。当您使用正确的同步原语时，大多数时候不需要担心这些底层细节。但在调试复杂的并发问题时，理解这些语义能帮您找到问题的根源。
