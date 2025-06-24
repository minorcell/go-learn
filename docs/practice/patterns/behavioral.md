---
title: "行为的定义：Go 对象间的协作协议"
description: "探索 Go 语言如何利用函数、接口和通道等核心特性，以地道且简洁的方式实现策略、观察者、责任链等经典行为型设计模式。"
---

# 行为的定义：Go 对象间的协作协议

行为型设计模式的核心在于**定义对象之间的通信与协作**。在 Go 语言中，我们不依赖复杂的类继承体系，而是通过其简洁的语言特性——尤其是接口、函数和通道——来构建清晰、高效的对象间协作协议。

这些模式在 Go 中往往有更直接、更轻量的实现。

## 1. 策略模式 (Strategy Pattern): 算法即函数

策略模式允许在运行时选择算法的行为。在其他语言中，这通常需要定义一个策略接口和多个实现该接口的类。而在 Go 中，我们可以做得更简单：直接使用函数。

假设我们需要根据不同的策略来格式化输出。

**第一步：定义策略为一个函数类型**

```go
// Strategy 是一个函数类型，它定义了算法的签名
type FormattingStrategy func(string) string
```

**第二步：实现具体的策略函数**

```go
func ToUpperCase(s string) string {
    return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
    return strings.ToLower(s)
}

func Capitalize(s string) string {
    return strings.Title(s)
}
```
每个函数都是一个独立的策略。

**第三步：创建使用策略的"上下文"**

```go
type Formatter struct {
    strategy FormattingStrategy
}

func (f *Formatter) SetStrategy(s FormattingStrategy) {
    f.strategy = s
}

func (f *Formatter) Print(text string) {
    formattedText := f.strategy(text)
    fmt.Println(formattedText)
}
```

**第四步：在运行时切换策略**
```go
formatter := &Formatter{}

formatter.SetStrategy(ToUpperCase)
formatter.Print("Hello, Gopher!") // 输出: HELLO, GOPHER!

formatter.SetStrategy(ToLowerCase)
formatter.Print("Hello, Gopher!") // 输出: hello, gopher!
```
通过将策略定义为函数，我们避免了为每个策略创建新类型的样板代码，使代码更加简洁和灵活。

## 2. 观察者模式 (Observer Pattern): 用通道解耦

观察者模式定义了一种一对多的依赖关系，当一个对象（"主题"）的状态发生改变时，所有依赖于它的对象（"观察者"）都会得到通知并被自动更新。在 Go 中，使用通道（Channel）是实现此模式的绝佳方式，它天生支持并发和解耦。

**第一步：定义主题和事件**
```go
// Event 表示状态变更的事件
type Event struct {
    Data string
}

// Subject 管理订阅者并广播事件
type Subject struct {
    observers []chan Event // 每个观察者是一个接收 Event 的通道
}

// 添加观察者（订阅）
func (s *Subject) Subscribe() chan Event {
    ch := make(chan Event, 1)
    s.observers = append(s.observers, ch)
    return ch
}

// 广播事件
func (s *Subject) Notify(event Event) {
    for _, ch := range s.observers {
        // 非阻塞发送，避免某个观察者阻塞整个广播
        go func(c chan Event) {
            c <- event
        }(ch)
    }
}
```

**第二步：实现观察者**

观察者就是一个监听其订阅通道的 goroutine。

```go
func main() {
    subject := &Subject{}

    // 创建两个观察者
    observer1 := subject.Subscribe()
    observer2 := subject.Subscribe()

    go func() {
        for event := range observer1 {
            fmt.Printf("Observer 1 received: %s\n", event.Data)
        }
    }()

    go func() {
        for event := range observer2 {
            fmt.Printf("Observer 2 received: %s\n", event.Data)
        }
    }()

    // 主题发布事件
    subject.Notify(Event{Data: "State has changed!"})

    time.Sleep(1 * time.Second) // 等待 goroutine 处理
}
```
这种基于通道的实现方式，将主题和观察者彻底解耦。主题不关心观察者是谁，只管向通道发送消息。观察者也不关心主题是谁，只管从通道接收消息。这使得系统非常容易扩展和进行并发处理。

## 3. 责任链模式 (Chain of Responsibility): `http.Handler` 的启示

责任链模式构建了一个对象链，请求沿着链传递，直到链上的某个对象处理它为止。Go 的 `net/http` 包中的中间件（Middleware）模型是该模式的最佳实践范例。

每个中间件都是链上的一个环节，它接收一个 `http.Handler`，并返回一个新的 `http.Handler`。

```go
type Middleware func(http.Handler) http.Handler

// 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request received: %s %s", r.Method, r.URL.Path)
        // 将请求传递给链中的下一个处理器
        next.ServeHTTP(w, r)
        log.Println("Request processing finished.")
    })
}

// 认证中间件
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") != "secret-token" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return // 终止链的执行
        }
        // 认证通过，传递给下一个处理器
        next.ServeHTTP(w, r)
    })
}

// 最终的处理器
func mainHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World from the main handler!"))
}

// 构建责任链
func buildChain(h http.Handler, mws ...Middleware) http.Handler {
    for i := len(mws) - 1; i >= 0; i-- {
        h = mws[i](h)
    }
    return h
}

func main() {
    finalHandler := http.HandlerFunc(mainHandler)
    chain := buildChain(finalHandler, LoggingMiddleware, AuthMiddleware)
    // http.Handle("/", chain)
}
```
在这个模型中，每个中间件都可以决定是处理请求、修改请求/响应，还是将其传递给链中的下一个环节，甚至直接终止请求。这种结构清晰、可组合、易于扩展。

## 结论

Go 语言通过其独特的核心特性，为传统行为型设计模式注入了新的活力。无论是使用函数实现策略模式，利用通道实现观察者模式，还是借鉴 `http` 包构建责任链，Go 的方式都倾向于更简洁、更解耦、更并发友好的解决方案。掌握这些地道的实现方式，是编写高质量 Go 代码的关键一步。
