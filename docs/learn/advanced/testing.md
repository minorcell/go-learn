# 测试：构建坚不可摧的Go程序

> 在专业的软件工程中，代码的正确性、可靠性和可维护性至关重要。测试不是事后的附加品，而是开发过程中不可或缺的一环。Go语言将这一理念深度融入其工具链，提供了简洁、强大且无需任何第三方框架的内置测试支持。
>
> 编写测试能让你充满信心地进行重构、添加新功能，并保证你的软件在各种边界条件下都能如预期般工作。

本文将以"构建自信"的视角，引导你学习Go语言中惯用（idiomatic）的测试方法，特别是强大的**表驱动测试 (Table-Driven Tests)**。

---

## 1. 测试基础

Go的测试工具遵循一套简单的约定：
- 测试代码必须位于 `_test.go` 后缀的文件中。
- 测试函数的命名必须以 `Test` 开头，例如 `TestMyFunction`。
- 测试函数必须接收一个参数：`t *testing.T`。

`*testing.T` 是你的测试"工具箱"，它提供了报告错误、标记失败、打印日志等方法。

让我们为一个简单的 `Add` 函数编写第一个测试。

**代码文件 `math.go`:**
```go
package main

func Add(a, b int) int {
	return a + b
}
```

**测试文件 `math_test.go`:**
```go
package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		// t.Errorf 会标记测试失败，并打印错误信息，但测试会继续执行
		t.Errorf("Add(1, 2) = %d; want %d", got, want)
	}
}
```

在终端中，进入该目录并运行 `go test`：
```sh
$ go test
PASS
ok  	myproject	0.001s
```
如果测试失败，你会看到详细的错误报告。

---

## 2. Go的惯用方式：表驱动测试

如果我们想测试更多的用例（例如，负数、零等），为每个用例都写一个新的`Test...`函数会非常繁琐。

```go
// 不推荐的方式
func TestAdd_Negative(t *testing.T) { /* ... */ }
func TestAdd_Zero(t *testing.T) { /* ... */ }
```

Go社区推崇一种更优雅、更具扩展性的模式：**表驱动测试**。这种模式将所有测试用例组织在一个"表"（通常是一个struct切片）中，然后用一个循环来执行所有测试。

让我们把 `TestAdd` 重构为表驱动测试：

**`math_test.go` (重构后):**
```go
package main

import "testing"

func TestAdd_TableDriven(t *testing.T) {
	// 1. 定义测试用例的"表"
	tests := []struct {
		name string // 测试用例的描述
		a    int    // 输入参数1
		b    int    // 输入参数2
		want int    // 期望的结果
	}{
		// 2. 填充测试用例
		{name: "positive numbers", a: 1, b: 2, want: 3},
		{name: "negative numbers", a: -1, b: -2, want: -3},
		{name: "zero", a: 0, b: 0, want: 0},
		{name: "positive and negative", a: 10, b: -5, want: 5},
	}

	// 3. 遍历表，执行测试
	for _, tt := range tests {
		// t.Run() 可以让你为每个测试用例创建一个独立的子测试
		// 这在测试失败时能提供更清晰的输出
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add() = %d; want %d", got, tt.want)
			}
		})
	}
}
```
**为什么这种模式如此强大？**
- **清晰**: 所有测试的输入和期望输出一目了然。
- **可维护**: 添加新的测试用例只需要在表中增加一行，无需编写新的函数。
- **独立**: 使用 `t.Run` 创建的子测试是独立的。即使一个子测试失败，其他子测试仍会继续执行。
- **可定位**: 如果一个子测试失败，`go test` 会明确报告出失败的子测试名称（例如 `TestAdd_TableDriven/negative_numbers`），让你能快速定位问题。

---

## 3. 其他内置测试工具

Go的 `testing` 包还内置了基准测试和示例功能。

### 基准测试 (Benchmarks)

基准测试用于衡量一段代码的性能。
- 函数名以 `Benchmark` 开头。
- 接收一个 `*testing.B` 类型的参数。
- 测试代码放在一个 `b.N` 循环内。

```go
// math_test.go

func BenchmarkAdd(b *testing.B) {
	// Go会自动调整 b.N 的值，直到测试运行时间足够长，可以得出稳定的平均值
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
```
使用 `go test -bench=.` 来运行基准测试。

### 示例 (Examples)

示例函数既是文档，也是可执行的测试。
- 函数名以 `Example` 开头。
- Go工具会执行示例代码，然后比较其标准输出与函数末尾注释中 `// Output:` 部分的内容。如果两者匹配，测试通过。

```go
// math_test.go

import "fmt"

func ExampleAdd() {
	sum := Add(5, 10)
	fmt.Println(sum)
	// Output: 15
}
```
这是一种绝佳的方式，可以确保你的文档中的代码示例永远不会过时或出错。这些示例也会出现在GoDoc生成的文档页面中。

---

## 总结

- Go的内置测试工具强大且易用，遵循简单的约定。
- **表驱动测试**是Go语言中编写单元测试的惯用（idiomatic）和推荐方式。它使测试更清晰、更易于维护。
- `t.Run()` 用于创建独立的子测试，是实现清晰表驱动测试的关键。
- Go还内置了对**基准测试**和**可执行示例**的支持，提供了一个完整的软件质量保证工具集。

掌握Go的测试方法，是成为一名专业Go开发者的必经之路。它能让你在开发过程中充满自信，交付高质量的软件。 