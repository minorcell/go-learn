# 接口

接口是Go语言最重要的特性之一，它定义了方法的集合，实现了代码的解耦和多态。Go的接口是隐式实现的，非常灵活和强大。

## 本章内容

- 接口的定义和实现
- 隐式接口实现
- 接口组合和嵌入
- 空接口和类型断言
- 接口的最佳实践

## 接口基础

### 接口定义和实现

```go
package main

import "fmt"

// 定义 Shape 接口
type Shape interface {
    Area() float64
    Perimeter() float64
}

// 定义 Drawable 接口
type Drawable interface {
    Draw() string
}

// 圆形结构体
type Circle struct {
    Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Circle 实现 Drawable 接口
func (c Circle) Draw() string {
    return fmt.Sprintf("绘制一个半径为%.2f的圆形", c.Radius)
}

// 矩形结构体
type Rectangle struct {
    Width, Height float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Rectangle 实现 Drawable 接口
func (r Rectangle) Draw() string {
    return fmt.Sprintf("绘制一个%.2f×%.2f的矩形", r.Width, r.Height)
}

// 使用接口的函数
func printShapeInfo(s Shape) {
    fmt.Printf("面积: %.2f\n", s.Area())
    fmt.Printf("周长: %.2f\n", s.Perimeter())
}

func drawShape(d Drawable) {
    fmt.Println(d.Draw())
}

// 计算总面积
func calculateTotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func main() {
    // 创建形状实例
    circle := Circle{Radius: 5.0}
    rectangle := Rectangle{Width: 10.0, Height: 6.0}
    
    // 使用接口
    fmt.Println("=== 圆形信息 ===")
    printShapeInfo(circle)
    drawShape(circle)
    
    fmt.Println("\n=== 矩形信息 ===")
    printShapeInfo(rectangle)
    drawShape(rectangle)
    
    // 接口切片
    shapes := []Shape{circle, rectangle}
    
    fmt.Printf("\n=== 总计算 ===\n")
    fmt.Printf("总面积: %.2f\n", calculateTotalArea(shapes))
    
    // 遍历形状
    fmt.Println("\n=== 所有形状 ===")
    for i, shape := range shapes {
        fmt.Printf("形状 %d:\n", i+1)
        printShapeInfo(shape)
        
        // 类型断言查看具体类型
        if drawable, ok := shape.(Drawable); ok {
            drawShape(drawable)
        }
        fmt.Println()
    }
}
```

### 接口的多态性

```go
package main

import "fmt"

// 动物接口
type Animal interface {
    Speak() string
    Move() string
}

// 宠物接口
type Pet interface {
    Animal
    Name() string
    Owner() string
}

// 狗
type Dog struct {
    name  string
    owner string
}

func (d Dog) Speak() string {
    return "汪汪！"
}

func (d Dog) Move() string {
    return "跑跑跳跳"
}

func (d Dog) Name() string {
    return d.name
}

func (d Dog) Owner() string {
    return d.owner
}

// 猫
type Cat struct {
    name  string
    owner string
}

func (c Cat) Speak() string {
    return "喵喵~"
}

func (c Cat) Move() string {
    return "优雅地走着"
}

func (c Cat) Name() string {
    return c.name
}

func (c Cat) Owner() string {
    return c.owner
}

// 鸟
type Bird struct {
    species string
}

func (b Bird) Speak() string {
    return "叽叽喳喳"
}

func (b Bird) Move() string {
    return "在天空中飞翔"
}

// 动物园管理员
type Zookeeper struct {
    name string
}

// 管理员可以照顾任何动物
func (z Zookeeper) CareFor(animal Animal) {
    fmt.Printf("%s 正在照顾动物\n", z.name)
    fmt.Printf("动物叫声: %s\n", animal.Speak())
    fmt.Printf("动物行为: %s\n", animal.Move())
}

// 照顾宠物需要更多信息
func (z Zookeeper) CarePet(pet Pet) {
    fmt.Printf("%s 正在照顾宠物 %s (主人: %s)\n", z.name, pet.Name(), pet.Owner())
    z.CareFor(pet) // 宠物也是动物
}

// 训练动物
func trainAnimal(animal Animal) {
    fmt.Printf("训练中... 动物发出: %s\n", animal.Speak())
}

func main() {
    // 创建动物实例
    dog := Dog{name: "旺财", owner: "张三"}
    cat := Cat{name: "咪咪", owner: "李四"}
    bird := Bird{species: "鹦鹉"}
    
    // 创建管理员
    keeper := Zookeeper{name: "王管理员"}
    
    fmt.Println("=== 动物照顾 ===")
    
    // 多态：相同的方法，不同的行为
    animals := []Animal{dog, cat, bird}
    for _, animal := range animals {
        keeper.CareFor(animal)
        fmt.Println()
    }
    
    fmt.Println("=== 宠物特殊照顾 ===")
    
    // 宠物需要特殊照顾
    pets := []Pet{dog, cat}
    for _, pet := range pets {
        keeper.CarePet(pet)
        fmt.Println()
    }
    
    fmt.Println("=== 动物训练 ===")
    
    // 统一的训练接口
    for _, animal := range animals {
        trainAnimal(animal)
    }
}
```

## 接口组合

### 接口嵌入

```go
package main

import "fmt"

// 基础接口
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string) error
}

type Closer interface {
    Close() error
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// 文件结构体
type File struct {
    name     string
    content  string
    isOpen   bool
}

// 实现所有接口方法
func (f *File) Read() string {
    if !f.isOpen {
        return "文件未打开"
    }
    return f.content
}

func (f *File) Write(data string) error {
    if !f.isOpen {
        return fmt.Errorf("文件未打开")
    }
    f.content += data
    return nil
}

func (f *File) Close() error {
    if !f.isOpen {
        return fmt.Errorf("文件已关闭")
    }
    f.isOpen = false
    fmt.Printf("文件 %s 已关闭\n", f.name)
    return nil
}

func (f *File) Open() error {
    if f.isOpen {
        return fmt.Errorf("文件已打开")
    }
    f.isOpen = true
    fmt.Printf("文件 %s 已打开\n", f.name)
    return nil
}

// 网络连接
type NetworkConnection struct {
    address string
    buffer  string
    active  bool
}

func (nc *NetworkConnection) Read() string {
    if !nc.active {
        return "连接未激活"
    }
    data := nc.buffer
    nc.buffer = ""
    return data
}

func (nc *NetworkConnection) Write(data string) error {
    if !nc.active {
        return fmt.Errorf("连接未激活")
    }
    nc.buffer += data
    return nil
}

func (nc *NetworkConnection) Close() error {
    if !nc.active {
        return fmt.Errorf("连接已关闭")
    }
    nc.active = false
    fmt.Printf("网络连接 %s 已关闭\n", nc.address)
    return nil
}

func (nc *NetworkConnection) Connect() error {
    if nc.active {
        return fmt.Errorf("连接已激活")
    }
    nc.active = true
    fmt.Printf("已连接到 %s\n", nc.address)
    return nil
}

// 使用不同接口的函数
func processReader(r Reader) {
    fmt.Printf("读取数据: %s\n", r.Read())
}

func processWriter(w Writer) {
    err := w.Write("Hello, World!")
    if err != nil {
        fmt.Printf("写入失败: %v\n", err)
    } else {
        fmt.Println("数据写入成功")
    }
}

func processReadWriter(rw ReadWriter) {
    fmt.Println("=== ReadWriter 操作 ===")
    processWriter(rw)
    processReader(rw)
}

func processReadWriteCloser(rwc ReadWriteCloser) {
    fmt.Println("=== ReadWriteCloser 操作 ===")
    processWriter(rwc)
    processReader(rwc)
    rwc.Close()
}

func main() {
    // 创建文件
    file := &File{name: "test.txt", content: "初始内容\n", isOpen: false}
    file.Open()
    
    // 创建网络连接
    conn := &NetworkConnection{address: "192.168.1.1:8080", active: false}
    conn.Connect()
    
    fmt.Println("=== 基础接口测试 ===")
    
    // 测试单一接口
    processReader(file)
    processWriter(file)
    processReader(file)
    
    fmt.Println("\n=== 组合接口测试 ===")
    
    // 测试组合接口
    processReadWriter(file)
    
    fmt.Println("\n=== 完整接口测试 ===")
    
    // 测试完整接口
    processReadWriteCloser(conn)
    
    // 重新打开文件测试
    file.Open()
    processReadWriteCloser(file)
}
```

### 接口设计模式

```go
package main

import (
    "fmt"
    "strings"
)

// 策略模式 - 排序策略接口
type SortStrategy interface {
    Sort([]int) []int
    Name() string
}

// 冒泡排序
type BubbleSort struct{}

func (bs BubbleSort) Sort(arr []int) []int {
    result := make([]int, len(arr))
    copy(result, arr)
    
    n := len(result)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if result[j] > result[j+1] {
                result[j], result[j+1] = result[j+1], result[j]
            }
        }
    }
    return result
}

func (bs BubbleSort) Name() string {
    return "冒泡排序"
}

// 快速排序
type QuickSort struct{}

func (qs QuickSort) Sort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    result := make([]int, len(arr))
    copy(result, arr)
    qs.quickSort(result, 0, len(result)-1)
    return result
}

func (qs QuickSort) quickSort(arr []int, low, high int) {
    if low < high {
        pi := qs.partition(arr, low, high)
        qs.quickSort(arr, low, pi-1)
        qs.quickSort(arr, pi+1, high)
    }
}

func (qs QuickSort) partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

func (qs QuickSort) Name() string {
    return "快速排序"
}

// 排序上下文
type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(arr []int) []int {
    if s.strategy == nil {
        return arr
    }
    return s.strategy.Sort(arr)
}

// 观察者模式
type Observer interface {
    Update(data interface{})
    GetID() string
}

type Subject interface {
    Attach(Observer)
    Detach(Observer)
    Notify(data interface{})
}

// 新闻订阅者
type NewsSubscriber struct {
    id   string
    name string
}

func (ns *NewsSubscriber) Update(data interface{}) {
    if news, ok := data.(string); ok {
        fmt.Printf("[%s] %s 收到新闻: %s\n", ns.id, ns.name, news)
    }
}

func (ns *NewsSubscriber) GetID() string {
    return ns.id
}

// 新闻发布者
type NewsPublisher struct {
    observers []Observer
    news      string
}

func (np *NewsPublisher) Attach(observer Observer) {
    np.observers = append(np.observers, observer)
}

func (np *NewsPublisher) Detach(observer Observer) {
    for i, obs := range np.observers {
        if obs.GetID() == observer.GetID() {
            np.observers = append(np.observers[:i], np.observers[i+1:]...)
            break
        }
    }
}

func (np *NewsPublisher) Notify(data interface{}) {
    for _, observer := range np.observers {
        observer.Update(data)
    }
}

func (np *NewsPublisher) PublishNews(news string) {
    np.news = news
    fmt.Printf("📰 发布新闻: %s\n", news)
    np.Notify(news)
}

// 装饰器模式
type Logger interface {
    Log(message string)
}

// 基础日志器
type BasicLogger struct{}

func (bl BasicLogger) Log(message string) {
    fmt.Printf("LOG: %s\n", message)
}

// 时间戳装饰器
type TimestampDecorator struct {
    logger Logger
}

func (td TimestampDecorator) Log(message string) {
    timestamped := fmt.Sprintf("[2023-12-07 10:30:45] %s", message)
    td.logger.Log(timestamped)
}

// 级别装饰器
type LevelDecorator struct {
    logger Logger
    level  string
}

func (ld LevelDecorator) Log(message string) {
    leveled := fmt.Sprintf("[%s] %s", strings.ToUpper(ld.level), message)
    ld.logger.Log(leveled)
}

func main() {
    // 策略模式演示
    fmt.Println("=== 策略模式 ===")
    
    data := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Printf("原始数据: %v\n", data)
    
    sorter := &Sorter{}
    
    // 使用冒泡排序
    sorter.SetStrategy(BubbleSort{})
    result1 := sorter.Sort(data)
    fmt.Printf("%s 结果: %v\n", sorter.strategy.Name(), result1)
    
    // 使用快速排序
    sorter.SetStrategy(QuickSort{})
    result2 := sorter.Sort(data)
    fmt.Printf("%s 结果: %v\n", sorter.strategy.Name(), result2)
    
    // 观察者模式演示
    fmt.Println("\n=== 观察者模式 ===")
    
    publisher := &NewsPublisher{}
    
    subscriber1 := &NewsSubscriber{id: "001", name: "张三"}
    subscriber2 := &NewsSubscriber{id: "002", name: "李四"}
    subscriber3 := &NewsSubscriber{id: "003", name: "王五"}
    
    publisher.Attach(subscriber1)
    publisher.Attach(subscriber2)
    publisher.Attach(subscriber3)
    
    publisher.PublishNews("Go 1.21 正式发布！")
    
    // 取消订阅
    publisher.Detach(subscriber2)
    fmt.Println("\n李四取消订阅")
    
    publisher.PublishNews("Go 语言教程更新了！")
    
    // 装饰器模式演示
    fmt.Println("\n=== 装饰器模式 ===")
    
    // 基础日志
    basicLogger := BasicLogger{}
    basicLogger.Log("基础日志消息")
    
    // 添加时间戳
    timestampLogger := TimestampDecorator{logger: basicLogger}
    timestampLogger.Log("带时间戳的日志")
    
    // 添加级别
    levelLogger := LevelDecorator{logger: basicLogger, level: "info"}
    levelLogger.Log("带级别的日志")
    
    // 组合装饰器
    compositeLogger := TimestampDecorator{
        logger: LevelDecorator{
            logger: basicLogger,
            level:  "error",
        },
    }
    compositeLogger.Log("组合装饰器日志")
}
```

## 空接口和类型断言

### 空接口的使用

```go
package main

import "fmt"

// 空接口可以持有任何类型的值
func processAnyValue(value interface{}) {
    fmt.Printf("接收到值: %v, 类型: %T\n", value, value)
}

// 类型断言
func identifyType(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("这是一个整数: %d\n", v)
    case string:
        fmt.Printf("这是一个字符串: %s\n", v)
    case bool:
        fmt.Printf("这是一个布尔值: %t\n", v)
    case []int:
        fmt.Printf("这是一个整数切片: %v\n", v)
    case map[string]int:
        fmt.Printf("这是一个字符串到整数的映射: %v\n", v)
    case func():
        fmt.Println("这是一个无参数无返回值的函数")
        v() // 调用函数
    default:
        fmt.Printf("未知类型: %T, 值: %v\n", v, v)
    }
}

// 安全的类型断言
func safeTypeAssertion(value interface{}) {
    // 使用 ok 模式进行安全断言
    if str, ok := value.(string); ok {
        fmt.Printf("字符串值: %s (长度: %d)\n", str, len(str))
    } else {
        fmt.Printf("不是字符串类型: %T\n", value)
    }
    
    if num, ok := value.(int); ok {
        fmt.Printf("整数值: %d (平方: %d)\n", num, num*num)
    } else {
        fmt.Printf("不是整数类型: %T\n", value)
    }
}

// 通用容器
type Container struct {
    items []interface{}
}

func (c *Container) Add(item interface{}) {
    c.items = append(c.items, item)
}

func (c *Container) Get(index int) interface{} {
    if index < 0 || index >= len(c.items) {
        return nil
    }
    return c.items[index]
}

func (c *Container) Size() int {
    return len(c.items)
}

func (c *Container) ForEach(fn func(interface{})) {
    for _, item := range c.items {
        fn(item)
    }
}

// JSON 风格的数据处理
func processJSONLike(data interface{}) {
    switch v := data.(type) {
    case map[string]interface{}:
        fmt.Println("处理对象:")
        for key, value := range v {
            fmt.Printf("  %s: ", key)
            processJSONLike(value)
        }
    case []interface{}:
        fmt.Println("处理数组:")
        for i, item := range v {
            fmt.Printf("  [%d]: ", i)
            processJSONLike(item)
        }
    case string:
        fmt.Printf("字符串: \"%s\"\n", v)
    case float64:
        fmt.Printf("数字: %.2f\n", v)
    case bool:
        fmt.Printf("布尔值: %t\n", v)
    case nil:
        fmt.Println("null")
    default:
        fmt.Printf("其他类型: %T = %v\n", v, v)
    }
}

func main() {
    fmt.Println("=== 空接口基础 ===")
    
    // 空接口可以接受任何类型
    processAnyValue(42)
    processAnyValue("Hello")
    processAnyValue(true)
    processAnyValue([]int{1, 2, 3})
    
    fmt.Println("\n=== 类型识别 ===")
    
    values := []interface{}{
        42,
        "Hello, World!",
        true,
        []int{1, 2, 3, 4, 5},
        map[string]int{"apple": 5, "banana": 3},
        func() { fmt.Println("  这是一个匿名函数") },
        3.14159,
    }
    
    for _, value := range values {
        identifyType(value)
    }
    
    fmt.Println("\n=== 安全类型断言 ===")
    
    testValues := []interface{}{
        "Go语言",
        42,
        3.14,
        true,
    }
    
    for _, value := range testValues {
        fmt.Printf("测试值: %v\n", value)
        safeTypeAssertion(value)
        fmt.Println()
    }
    
    fmt.Println("=== 通用容器 ===")
    
    container := &Container{}
    container.Add("字符串")
    container.Add(123)
    container.Add(true)
    container.Add([]int{1, 2, 3})
    
    fmt.Printf("容器大小: %d\n", container.Size())
    
    // 遍历容器
    fmt.Println("容器内容:")
    container.ForEach(func(item interface{}) {
        fmt.Printf("- %v (%T)\n", item, item)
    })
    
    // 获取特定项目
    item := container.Get(1)
    if num, ok := item.(int); ok {
        fmt.Printf("第2个项目是整数: %d\n", num)
    }
    
    fmt.Println("\n=== JSON风格数据 ===")
    
    // 模拟JSON数据结构
    jsonData := map[string]interface{}{
        "name": "张三",
        "age":  30.0,
        "married": true,
        "children": []interface{}{
            map[string]interface{}{
                "name": "小明",
                "age":  8.0,
            },
            map[string]interface{}{
                "name": "小红",
                "age":  6.0,
            },
        },
        "address": map[string]interface{}{
            "street": "长安街",
            "city":   "北京",
            "zipcode": "100000",
        },
        "spouse": nil,
    }
    
    processJSONLike(jsonData)
}
```

### 接口断言和类型开关

```go
package main

import (
    "fmt"
    "strconv"
)

// 定义一些接口
type Stringer interface {
    String() string
}

type Counter interface {
    Count() int
}

type Resetter interface {
    Reset()
}

// 实现多个接口的结构体
type WordCounter struct {
    words []string
}

func (wc *WordCounter) String() string {
    return fmt.Sprintf("WordCounter with %d words: %v", len(wc.words), wc.words)
}

func (wc *WordCounter) Count() int {
    return len(wc.words)
}

func (wc *WordCounter) Reset() {
    wc.words = nil
}

func (wc *WordCounter) AddWord(word string) {
    wc.words = append(wc.words, word)
}

// 数字计数器
type NumberCounter struct {
    value int
}

func (nc *NumberCounter) String() string {
    return fmt.Sprintf("NumberCounter: %d", nc.value)
}

func (nc *NumberCounter) Count() int {
    return nc.value
}

func (nc *NumberCounter) Reset() {
    nc.value = 0
}

func (nc *NumberCounter) Increment() {
    nc.value++
}

func (nc *NumberCounter) Add(n int) {
    nc.value += n
}

// 接口检测函数
func analyzeInterface(value interface{}) {
    fmt.Printf("\n=== 分析接口: %T ===\n", value)
    
    // 检测 Stringer 接口
    if stringer, ok := value.(Stringer); ok {
        fmt.Printf("实现了 Stringer: %s\n", stringer.String())
    } else {
        fmt.Println("未实现 Stringer")
    }
    
    // 检测 Counter 接口
    if counter, ok := value.(Counter); ok {
        fmt.Printf("实现了 Counter: 计数 = %d\n", counter.Count())
    } else {
        fmt.Println("未实现 Counter")
    }
    
    // 检测 Resetter 接口
    if resetter, ok := value.(Resetter); ok {
        fmt.Println("实现了 Resetter，正在重置...")
        resetter.Reset()
        
        // 重置后再次检查计数
        if counter, ok := value.(Counter); ok {
            fmt.Printf("重置后计数: %d\n", counter.Count())
        }
    } else {
        fmt.Println("未实现 Resetter")
    }
}

// 类型开关处理不同类型
func handleDifferentTypes(value interface{}) {
    switch v := value.(type) {
    case *WordCounter:
        fmt.Printf("单词计数器处理: 添加单词 'hello'\n")
        v.AddWord("hello")
        
    case *NumberCounter:
        fmt.Printf("数字计数器处理: 增加5\n")
        v.Add(5)
        
    case string:
        fmt.Printf("字符串处理: 转换为大写: %s\n", v)
        
    case int:
        fmt.Printf("整数处理: 乘以2 = %d\n", v*2)
        
    case []string:
        fmt.Printf("字符串切片处理: 连接为: %s\n", 
            fmt.Sprintf("[%s]", fmt.Sprintf("%v", v)))
        
    default:
        fmt.Printf("未知类型处理: %T = %v\n", v, v)
    }
}

// 组合接口检测
func checkCombinedInterface(value interface{}) {
    fmt.Printf("\n=== 组合接口检测: %T ===\n", value)
    
    // 检测是否同时实现了多个接口
    canCount := false
    canReset := false
    canString := false
    
    if _, ok := value.(Counter); ok {
        canCount = true
    }
    if _, ok := value.(Resetter); ok {
        canReset = true
    }
    if _, ok := value.(Stringer); ok {
        canString = true
    }
    
    fmt.Printf("计数能力: %t, 重置能力: %t, 字符串化: %t\n", 
        canCount, canReset, canString)
    
    // 如果实现了所有接口，进行完整操作
    if canCount && canReset && canString {
        fmt.Println("实现了所有接口，执行完整操作流程:")
        
        if stringer := value.(Stringer); stringer != nil {
            fmt.Printf("1. 当前状态: %s\n", stringer.String())
        }
        
        if counter := value.(Counter); counter != nil {
            fmt.Printf("2. 当前计数: %d\n", counter.Count())
        }
        
        if resetter := value.(Resetter); resetter != nil {
            fmt.Println("3. 执行重置")
            resetter.Reset()
        }
        
        if stringer := value.(Stringer); stringer != nil {
            fmt.Printf("4. 重置后状态: %s\n", stringer.String())
        }
    }
}

// 动态接口调用
func dynamicInterfaceCall(values []interface{}) {
    fmt.Println("\n=== 动态接口调用 ===")
    
    for i, value := range values {
        fmt.Printf("\n--- 对象 %d ---\n", i+1)
        
        // 根据支持的接口动态调用
        operations := []string{}
        
        if counter, ok := value.(Counter); ok {
            count := counter.Count()
            operations = append(operations, "Count()="+strconv.Itoa(count))
        }
        
        if stringer, ok := value.(Stringer); ok {
            str := stringer.String()
            operations = append(operations, "String()="+str)
        }
        
        if len(operations) > 0 {
            fmt.Printf("支持的操作: %v\n", operations)
        } else {
            fmt.Printf("类型: %T, 值: %v (无支持的接口)\n", value, value)
        }
    }
}

func main() {
    // 创建测试对象
    wc := &WordCounter{words: []string{"hello", "world", "go"}}
    nc := &NumberCounter{value: 10}
    
    fmt.Println("=== 初始状态 ===")
    fmt.Printf("WordCounter: %s\n", wc.String())
    fmt.Printf("NumberCounter: %s\n", nc.String())
    
    // 分析接口实现
    analyzeInterface(wc)
    analyzeInterface(nc)
    analyzeInterface("普通字符串")
    analyzeInterface(42)
    
    // 重新创建对象进行类型处理测试
    wc2 := &WordCounter{words: []string{"go", "lang"}}
    nc2 := &NumberCounter{value: 5}
    
    fmt.Println("\n=== 类型特定处理 ===")
    testValues := []interface{}{wc2, nc2, "hello", 10, []string{"a", "b", "c"}}
    
    for _, value := range testValues {
        handleDifferentTypes(value)
    }
    
    // 组合接口检测
    wc3 := &WordCounter{words: []string{"test"}}
    nc3 := &NumberCounter{value: 15}
    
    checkCombinedInterface(wc3)
    checkCombinedInterface(nc3)
    checkCombinedInterface("字符串不实现接口")
    
    // 动态接口调用
    mixedValues := []interface{}{
        &WordCounter{words: []string{"dynamic", "call"}},
        &NumberCounter{value: 20},
        "plain string",
        123,
        true,
    }
    
    dynamicInterfaceCall(mixedValues)
}
```

## 本章小结

在这一章中，我们深入学习了Go语言的接口系统：

### 接口基础
- **隐式实现** - 无需显式声明实现接口
- **方法集合** - 接口定义方法的集合
- **多态性** - 不同类型实现相同接口
- **接口变量** - 可以持有任何实现该接口的值

### 接口组合
- **接口嵌入** - 通过嵌入组合接口
- **设计模式** - 策略、观察者、装饰器等
- **灵活设计** - 小接口组合成大功能
- **解耦设计** - 依赖接口而非具体类型

### 空接口和断言
- **interface{}** - 可以持有任何类型的值
- **类型断言** - 安全地获取具体类型
- **类型开关** - 根据类型执行不同逻辑
- **通用编程** - 编写类型无关的通用代码

### 最佳实践
- 保持接口小而专注
- 优先定义接口而非结构体
- 在消费者端定义接口
- 使用接口实现松耦合设计
- 合理使用空接口避免类型丢失

### 设计原则
- **单一职责** - 每个接口职责明确
- **接口隔离** - 客户端不依赖不需要的方法
- **依赖倒置** - 依赖抽象而非具体实现
- **组合优于继承** - 通过接口组合实现功能
