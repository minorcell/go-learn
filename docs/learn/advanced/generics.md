# 泛型：编写代码的瑞士军刀

> 在Go 1.18版本之前，开发者在编写需要处理多种数据类型的通用代码时，常常面临一个两难的抉择：要么为每种类型复制粘贴大量重复代码，要么使用空接口 `interface{}` 并放弃类型安全。
>
> **泛型 (Generics)** 的引入，彻底解决了这个痛点。它允许我们编写灵活、可复用、同时又是完全类型安全的代码。我们可以把泛型看作是为我们的代码工具箱打造的一把"瑞士军刀"——一个工具，多种用途。

本文将带你领略泛型的威力，看看它是如何让我们编写出更简洁、更强大的Go代码。

---

## 1. 泛型之前：一个两难的困境

让我们以一个简单的数据结构——**栈 (Stack)** 为例。栈遵循"后进先出" (LIFO) 的原则。

如果我们想实现一个只能操作 `int` 的栈，代码很简单：

```go
type IntStack struct {
	items []int
}
func (s *IntStack) Push(i int) { s.items = append(s.items, i) }
func (s *IntStack) Pop() int   { /* ... */ }
```

但如果我们还需要一个操作 `string` 的栈呢？我们就不得不复制一份几乎完全一样的代码：

```go
type StringStack struct {
	items []string
}
func (s *StringStack) Push(i string) { s.items = append(s.items, i) }
func (s *StringStack) Pop() string   { /* ... */ }
```
这显然是不可持续的。

另一种方法是使用空接口 `interface{}`：

```go
type InterfaceStack struct {
	items []interface{}
}
func (s *InterfaceStack) Push(i interface{}) { s.items = append(s.items, i) }
func (s *InterfaceStack) Pop() interface{}   { /* ... */ }

// 使用时...
var stack InterfaceStack
stack.Push(10)
stack.Push("hello") // 类型不安全！可以把任何类型推进去

// 取出时需要类型断言，非常繁琐且容易出错
val, ok := stack.Pop().(int)
if !ok {
    // 错误处理
}
```
这种方式虽然避免了代码重复，但牺牲了Go语言最重要的特性之一：**类型安全**。我们需要繁琐的类型断言，并且编译器无法在编译时发现类型错误。

---

## 2. 泛型函数：一个简单的开始

泛型通过**类型参数 (Type Parameters)** 来解决这个问题。类型参数就像一个占位符，它代表一个未知的、将在未来调用时被指定的具体类型。

类型参数列表写在中括号 `[]` 内，位于函数名之后。

```go
package main

import "fmt"

// PrintSlice 是一个泛型函数
// [T any] 是类型参数列表。T 是类型参数，any 是它的约束
// any 是一个预定义的约束，表示 T 可以是任何类型
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"a", "b", "c"}

	// 调用泛型函数时，通常不需要显式指定类型参数
	// Go的类型推断会自动识别出 T 是 int 或 string
	PrintSlice(intSlice)
	PrintSlice(stringSlice)
}
```

---

## 3. 类型约束：为泛型设定规则

`any` 约束太宽泛了。如果我们想写一个对数字切片求和的函数，`any` 类型是行不通的，因为它不支持 `+` 操作。

我们需要一个更精确的**约束 (Constraint)**，来告诉编译器类型参数 `T` 必须满足某些条件。在Go中，**约束本身就是一个接口**。

我们可以定义一个接口，列出所有允许的类型。

```go
package main

import "fmt"

// Number 是一个接口，它被用作类型约束
// 它规定了任何满足此约束的类型，必须是 int64 或 float64
type Number interface {
	int64 | float64
}

// SumNumbers 是一个泛型函数，它的类型参数 T 必须满足 Number 约束
func SumNumbers[T Number](nums []T) T {
	var sum T
	for _, v := range nums {
		sum += v // 因为 T 被约束为 Number，所以编译器知道 + 操作是合法的
	}
	return sum
}

func main() {
	intData := []int64{1, 2, 3, 4}
	floatData := []float64{1.1, 2.2, 3.3}

	fmt.Println("Sum of ints:", SumNumbers(intData))
	fmt.Println("Sum of floats:", SumNumbers(floatData))

	// 下面的代码会编译失败，因为 string 不满足 Number 约束
	// stringData := []string{"a", "b"}
	// SumNumbers(stringData)
}
```
Go标准库的 `constraints` 包预定义了一些常用的约束，例如 `constraints.Integer` 和 `constraints.Float`。但自定义接口约束的方式更为通用和强大。

---

## 4. 泛型类型：打造通用数据结构

现在，让我们用泛型来完美地解决最初的栈问题。我们可以定义一个**泛型类型**。

```go
package main

import "fmt"

// Stack[T any] 是一个泛型类型
// T 可以是任何类型
type Stack[T any] struct {
	items []T
}

// Push 方法的接收者是 *Stack[T]，表明它是泛型类型的方法
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // 声明一个 T 类型的零值
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func main() {
	// --- 使用 int 类型的栈 ---
	// 在实例化时，我们指定 T 就是 int
	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)

	val, ok := intStack.Pop()
	if ok {
		// val 的类型就是 int，无需任何类型断言
		fmt.Printf("Popped int: %d (type: %T)\n", val, val)
	}

	// --- 使用 string 类型的栈 ---
	// 在实例化时，我们指定 T 就是 string
	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	
	sVal, sOk := stringStack.Pop()
	if sOk {
		// sVal 的类型就是 string
		fmt.Printf("Popped string: %s (type: %T)\n", sVal, sVal)
	}
	
	// 下面的代码会编译失败，因为 intStack 是 Stack[int] 类型
	// 它不能 Push 一个 string
	// intStack.Push("this will fail")
}
```
通过泛型，我们只用一份代码就实现了一个完全**类型安全**且**可复用**的栈。这就是泛型的威力。

---

## 总结

- 泛型允许我们编写能处理多种类型的函数和数据结构，而无需重复代码。
- 类型参数 `[T ...]` 用于定义泛型。
- 约束（本身是接口）用于限制类型参数所能接受的具体类型范围，保证操作的安全性。
- 泛型类型（如 `Stack[T]`) 让我们能够创建通用的、类型安全的数据结构。

泛型是Go语言一个里程碑式的增强。它使得我们可以编写出更抽象、更灵活、同时又不失Go语言核心优势——简洁与类型安全——的代码。 