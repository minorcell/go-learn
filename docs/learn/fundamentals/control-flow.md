# 控制流：指挥程序的逻辑

> 我们已经学会了如何定义数据，但程序真正的威力在于它的动态性——根据不同的情况执行不同的指令。**控制流（Control Flow）** 就是我们用来指挥程序执行顺序的工具。它决定了代码是顺序执行、进行选择，还是重复操作。

本文将引导你像一个算法设计师一样思考。我们将把控制流语句看作是构建程序逻辑流程图的指令，让你的代码从静态的文本，变成一个能够做出决策、执行任务的动态过程。

---

## `if/else`：代码的十字路口

最基本的决策工具是 `if` 语句。你可以把它想象成程序执行路径上的一个十字路口。当程序走到这里，它会检查一个"路况"（一个布尔条件），然后决定走哪条路。

### 简单的 `if`

如果只有一个条件需要关心，我们就使用一个简单的 `if`。

**流程图视角**:
> 开始 -> 检查条件 -> (如果为真) -> 执行特定代码 -> 继续主流程
>             ^
>             |
>             (如果为假) -> 直接跳到主流程

```go
package main

import "fmt"

func main() {
	temperature := 30

	// 条件：temperature > 28
	if temperature > 28 {
		fmt.Println("天气炎热，建议开空调。") // 条件为真时执行
	}

	fmt.Println("祝您一天愉快。") // 无论条件如何，都会执行
}
```

### `if-else`：二选一的路径

当存在两种互斥的可能性时，我们使用 `if-else` 结构。

**流程图视角**:
> 开始 -> 检查条件 -> (如果为真) -> 执行A代码 -> 继续主流程
>             ^
>             |
>             (如果为假) -> 执行B代码 -> 继续主流程

```go
package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("您已成年，可以进入。")
	} else {
		fmt.Println("您未成年，禁止入内。")
	}
}
```

### `if-else if-else`：多重选择

当有多个条件需要依次判断时，就形成了 `if-else if-else` 梯子。

**流程图视角**:
> 开始 -> 检查条件1? -> (真) -> A -> 结束
>       |
>       (假)
>       v
>       检查条件2? -> (真) -> B -> 结束
>       |
>       (假)
>       v
>       ... (更多条件)
>       |
>       v
>       执行默认代码 -> 结束

```go
package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
```

**Go的特色：带初始化语句的`if`**

Go允许在 `if` 的条件判断之前，执行一个简短的初始化语句。这在处理函数返回值时特别有用，它可以将变量的作用域限制在 `if-else` 块内，让代码更整洁。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("test.txt")
	// if err != nil { ... }
	// 像上面这样，err变量会泄露到main函数作用域

	// 使用带初始化的if语句
	if file, err := os.Open("non-existent-file.txt"); err != nil {
		fmt.Println("错误:", err)
		// 这里的 file 和 err 变量只在 if/else 块内可见
	} else {
		fmt.Println("文件打开成功:", file.Name())
		file.Close()
	}
	
	// 在这里访问 err 会导致编译错误
	// fmt.Println(err) // undefined: err
}
```

---

## `switch`：高效的调度中心

当你需要根据一个表达式的多种可能值来执行不同操作时，`if-else if` 梯子会显得很冗长。这时，`switch` 就像一个高效的调度中心，能清晰地将表达式的值"调度"到对应的处理分支。

**流程图视角**:
> 开始 -> 评估表达式 -> (值等于case A?) -> 执行A代码 -> break -> 结束
>                      -> (值等于case B?) -> 执行B代码 -> break -> 结束
>                      -> (...)
>                      -> (无匹配?) -> 执行default代码 -> 结束

Go的 `switch` 非常强大和灵活：

- **自动 `break`**: 与C/C++/Java不同，Go的 `case` 默认在执行完毕后自动终止，你不需要手动写 `break`。这避免了意外的"贯穿"（fallthrough）错误。如果确实需要贯穿，可以使用 `fallthrough` 关键字。
- **表达式作为 `case`**: `case` 后面可以跟任意能产生同类型值的表达式，不限于常量。
- **无表达式 `switch`**: `switch` 后面可以不跟任何表达式，此时它等价于 `switch true`，可以让你写出更清晰的 `if-else if` 逻辑链。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 基本的 switch
	day := "Friday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("这是工作日。")
	case "Saturday", "Sunday":
		fmt.Println("这是周末！")
	default:
		fmt.Println("无效的星期。")
	}

	// 2. 无表达式的 switch (更清晰的 if-else if)
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("上午好！")
	case hour < 18:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
```

---

## `for`：唯一的循环大师

在许多语言中，你会看到 `while`、`do-while`、`for`、`foreach` 等多种循环语句。Go语言秉承"少即是多"的哲学，只提供了一种循环结构：**`for` 循环**。

但别担心，这个 `for` 循环非常灵活，足以优雅地实现所有你需要的循环模式。

### 1. 经典三段式 `for` 循环 (类比C语言的for)

这是最常见的形式，包含初始化、条件和后置语句。

**流程图视角**:
> 开始 -> 初始化 -> 检查条件? -> (假) -> 结束
>           ^         |
>           |         (真)
>           |         v
>        后置语句 <- 执行循环体

```go
// 打印 0 到 4
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

### 2. 条件 `for` 循环 (类比while)

如果你省略初始化和后置语句，它就变成了一个 `while` 循环。

```go
n := 0
for n < 5 { // 只有条件
	fmt.Println(n)
	n++
}
```

### 3. 无限 `for` 循环 (死循环)

如果连条件也省略，你就得到了一个无限循环。通常与 `break` 或 `return` 结合使用。

```go
for {
	fmt.Println("这是一个无限循环，按 Ctrl+C 退出。")
	// 在真实应用中，这里会有退出条件，例如：
	// if someCondition {
	//     break 
	// }
	time.Sleep(1 * time.Second)
}
```

### 4. `for-range`：遍历数据结构

`for` 循环与 `range` 关键字结合，可以方便地遍历切片、映射、字符串、数组和通道。

```go
// 遍历切片
items := []string{"a", "b", "c"}
for index, value := range items {
	fmt.Printf("索引: %d, 值: %s\n", index, value)
}

// 遍历映射 (注意顺序是随机的)
ages := map[string]int{"Alice": 30, "Bob": 25}
for name, age := range ages {
	fmt.Printf("%s 的年龄是 %d\n", name, age)
}
```

### `break` 和 `continue`

- `break`: 完全跳出当前循环。
- `continue`: 停止当前这次迭代，并立即开始下一次迭代。

---

## 总结

Go的控制流设计体现了其核心哲学：**清晰、简洁、够用就好**。通过 `if`、`switch` 和一个统一的 `for` 循环，你可以构建出任何复杂的程序逻辑。

理解了数据结构和控制流，你就掌握了编程的两大支柱。接下来，我们将学习如何通过**函数**来组织和重用我们的代码，构建更大、更模块化的程序。

## 下一步

理解了 Go 的控制流设计哲学后，让我们探索[函数](/learn/fundamentals/functions)，看看 Go 如何让函数成为程序设计的强大工具。

 