# 变量与类型：构建你的第一个Go程序

> 欢迎来到Go语言的世界。编程的本质是处理数据，而我们与数据的第一次亲密接触，始于理解"变量"和"类型"。本文将引导你编写第一个Go程序，并在此过程中，为你建立一个关于Go如何看待和组织数据的坚实心智模型。

让我们忘掉枯燥的定义列表，从一段最简单的代码开始。

---

## 你的第一个Go程序：`hello.go`

在你的工作区，创建一个名为 `hello.go` 的文件，并输入以下内容：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

- `package main`: 声明这个文件属于 `main` 包。一个可执行程序必须有一个 `main` 包。
- `import "fmt"`: 导入一个名为 `fmt` 的标准库包。`fmt` 提供了格式化输入输出的功能，比如在屏幕上打印文本。
- `func main()`: 定义一个名为 `main` 的函数。这是我们程序的入口，当你运行程序时，`main` 函数里的代码会被执行。
- `fmt.Println(...)`: 调用 `fmt` 包里的 `Println` 函数，将 "Hello, World!" 这个**字符串**打印到控制台。

现在，打开终端，进入文件所在的目录，运行它：

```sh
$ go run hello.go
Hello, World!
```

恭喜！你已经成功运行了你的第一个Go程序。更重要的是，你已经接触到了Go的第一个基本类型：**字符串（string）**。

---

## 变量：给数据一个名字

程序很少只处理静态的文本。我们通常需要存储和操作数据。**变量（Variable）** 就是内存中一个存储数据的命名空间。

让我们修改一下程序，使用变量来存储我们的问候语：

```go
package main

import "fmt"

func main() {
	var greeting string = "Hello, Gopher!"
	fmt.Println(greeting)
}
```

这里发生了几件事：

1.  `var greeting string`: 我们**声明**了一个变量。
    -   `var` 是一个关键字，表示我们要声明一个变量。
    -   `greeting` 是我们给这个变量起的名字。
    -   `string` 是我们告诉Go编译器，这个变量**只能**存储字符串类型的数据。
2.  `= "Hello, Gopher!"`: 我们**初始化**了这个变量，给它赋了一个初始值。

这体现了Go作为一门**静态类型（Static Typing）**语言的核心特征：**变量的类型在编译时就已确定，并且永远不能改变**。你不能把一个数字赋给一个 `string` 类型的变量，这保证了类型安全，减少了运行时错误。

### 类型推断与短变量声明

Go的编译器很聪明。在很多情况下，它能根据你提供的值**推断**出变量的类型。这让我们可以使用一种更简洁的声明方式。

重写上面的代码：

```go
package main

import "fmt"

func main() {
	// Go编译器看到 "Hello, Gopher!" 就知道 greeting 是 string 类型
	var greeting = "Hello, Gopher!" 
	fmt.Println(greeting)
}
```

在函数内部，我们还有一种更常用、更简洁的方式：**短变量声明** `:=`。

```go
package main

import "fmt"

func main() {
	// := 结合了声明和初始化，var 关键字被省略
	// 这是在函数内部最常用的方式
	greeting := "Hello, Gopher!"
	fmt.Println(greeting)
}
```

**注意**: 短变量声明 `:=` 只能在函数内部使用。对于包级别的变量，必须使用 `var` 关键字。

---

## Go的核心类型：构建程序的基石

现在我们已经知道如何声明变量，让我们来认识一下Go语言提供的核心"建筑材料"。

### 1. 布尔型 (bool)

布尔型只有两个值：`true` 或 `false`。它通常用于条件判断。

```go
var isReady bool = true
canProceed := !isReady // "!" 是逻辑非操作符，结果是 false

fmt.Println(canProceed)
```

### 2. 数字类型 (Numeric Types)

Go提供了丰富的数字类型来满足不同的精度和平台需求。

#### 整型 (Integer)

-   `int`: 默认的整数类型。在32位系统上是32位，在64位系统上是64位。**在不确定时，优先使用 `int`**。
-   `int8`, `int16`, `int32`, `int64`: 固定大小的有符号整数。
-   `uint`, `uint8`, `uint16`, `uint32`, `uint64`: 无符号整数，不能表示负数。`byte` 是 `uint8` 的别名。
-   `uintptr`: 一种特殊的无符号整数，用于存储指针的地址，常用于底层编程。

```go
var score int = 100
playerCount := 10

// 不同类型的整数不能直接进行运算
// 下面的代码会编译错误: invalid operation: score + int32(playerCount) (mismatched types int and int32)
// var totalScore = score + int32(playerCount) 

// 必须进行显式类型转换
var totalScore = score + int(playerCount) 
fmt.Println(totalScore)
```
**类型转换**是Go静态类型哲学的重要体现：类型安全由你（开发者）来保证，编译器不会做任何隐式的转换。

#### 浮点型 (Floating-Point)

-   `float32`, `float64`: 32位和64位的浮点数，用于表示小数。**在不确定时，优先使用 `float64`**。

```go
var price float64 = 99.99
taxRate := 0.08
totalPrice := price * (1 + taxRate)

fmt.Println(totalPrice)
```

#### 复数 (Complex)

-   `complex64`, `complex128`: 用于科学计算和工程领域。

### 3. 字符串 (string)

我们已经见过字符串了。Go的字符串是**不可变的（immutable）**，由一系列 `byte` (UTF-8编码) 组成。一旦创建，你不能修改字符串的某个字符。

```go
path := "/usr/local/bin"
// path[0] = 'c' // 这会产生编译错误

// 你可以基于原字符串创建新的字符串
newPath := "C:" + path[1:] // 字符串拼接和切片操作
fmt.Println(newPath)
```

---

## 零值 (Zero Value)

Go有一个重要的概念：**任何变量在声明后，都会被自动初始化为其类型的"零值"**。这可以防止未初始化变量导致的问题。

-   数字类型：`0`
-   布尔类型：`false`
-   字符串类型：`""` (空字符串)
-   指针、接口、切片、映射、通道、函数类型：`nil`

```go
var playerCount int
var price float64
var greeting string
var isReady bool

// 打印它们的值，看看是不是零值
// Printf 是格式化打印函数，%v 代表默认格式，%q 带引号的字符串
fmt.Printf("playerCount: %v, price: %v, greeting: %q, isReady: %v\n", 
	playerCount, price, greeting, isReady)

// 输出:
// playerCount: 0, price: 0, greeting: "", isReady: false
```

这个特性让你的代码更安全、更可预测。

---

## 下一步

你已经掌握了Go语言中最基本，也是最重要的概念：如何通过变量和类型来描述和存储数据。这是构建任何复杂逻辑的起点。

在下一篇文章中，我们将学习如何使用 **数组（Array）、切片（Slice）和映射（Map）** 来组织和管理这些基本数据，从处理单个值升级到处理数据集合。 