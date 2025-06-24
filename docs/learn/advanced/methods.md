# 方法：为数据赋予行为

> 在前面的章节中，我们学习了如何用 `struct` 来定义数据的“名词”（如 `Person`, `Book`）。但是，这些数据本身是被动的。我们如何让它们拥有自己的“动词”呢？例如，让一个 `Rectangle` 自己会“计算面积”，或者让一个 `BankAccount` 自己能“存款”？
>
> Go语言通过**方法 (Methods)** 来实现这一点。方法是附加到特定类型上的特殊函数。这是Go语言实现“面向对象”思想的核心方式，但它比传统的类（class）更灵活、更直接。

本文将以“为数据赋予行为”的视角，探索如何通过方法让你的数据结构“活”起来。

---

## 1. 什么是方法？

方法就是一个带有**接收者 (receiver)** 的函数。接收者是方法所归属的类型，它出现在 `func` 关键字和方法名之间。

```go
package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area 是一个属于 Rectangle 类型的方法
// (r Rectangle) 就是接收者
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	// 调用方法，就像访问字段一样，使用点号 .
	fmt.Println("矩形的面积是:", rect.Area())
}
```

在 `func (r Rectangle) Area()`, `(r Rectangle)` 部分就是接收者。它声明了 `Area` 这个方法属于 `Rectangle` 类型。`r` 是接收者变量，在方法内部，我们可以用它来访问该 `Rectangle` 实例的字段。

---

## 2. 核心决策：值接收者 vs. 指针接收者

这是使用方法时最关键的决策。选择哪种接收者，决定了你的方法能否修改原始数据。

### 心智模型：复印件 vs. 原始文档

- **值接收者 `(t T)`**: 方法得到的是接收者的一份**复印件 (copy)**。
  - **类比**: 有人给了你一份文件的**复印件**。你可以阅读它，可以在上面涂写画画，但你对复印件做的任何修改，都**不会**影响到那份原始文件。

- **指针接收者 `(t *T)`**: 方法得到的是一个指向接收者的**指针 (pointer)**。
  - **类比**: 有人给了你一份在线文档的**共享链接**。你通过这个链接做的任何编辑，都会直接修改那份**唯一的、原始的**文档。

### 值接收者示例 (Read-Only)

值接收者最适合那些不需要修改接收者状态的"只读"操作。

```go
// (r Rectangle) 是值接收者
// Area() 只需要读取 Width 和 Height，不需要修改它们
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```
调用 `rect.Area()` 时，`Area` 方法内部的 `r` 是 `rect` 的一个完整副本。

### 指针接收者示例 (Modification)

当你需要在一个方法里修改接收者的字段时，**必须**使用指针接收者。

```go
package main

import "fmt"

type Account struct {
	balance float64
}

// Deposit 方法需要修改 balance 字段，所以必须用指针接收者
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		// a 是指向原始 Account 的指针
		// 这里的修改会影响到调用方的原始变量
		a.balance += amount
	}
}

// Balance 方法只是读取，理论上可以用值接收者
// 但为了保持一致性，通常也使用指针接收者
func (a *Account) Balance() float64 {
	return a.balance
}

func main() {
	myAccount := &Account{balance: 1000}

	myAccount.Deposit(500)
	
	fmt.Printf("我的账户余额: %.2f\n", myAccount.Balance()) // 输出: 1500.00
}
```

---

## 3. 如何选择？

遵循以下三个原则来决定使用哪种接收者：

1.  **修改**: 如果方法需要修改接收者的状态，**必须**使用指针接收者 (`*T`)。
2.  **性能**: 如果接收者是一个非常大的结构体，使用指针接收者可以避免在每次方法调用时都进行昂贵的拷贝，从而提升性能。
3.  **一致性**: 这是条非常重要的工程实践建议。**如果一个类型有一个方法使用了指针接收者，那么该类型的所有方法都应该使用指针接收者**，即使某些方法并不需要修改状态。这会让类型的使用方式更加统一和可预测，避免混淆。

> **经验法则**: 当你不确定时，**优先使用指针接收者**。它更通用，也更符合Go语言的工程实践。

---

## 4. 方法与任意类型

Go的另一个强大之处在于，方法可以被定义在任何你自定义的类型上，不限于结构体。

```go
package main

import (
	"fmt"
	"time"
)

// MyDuration 是 time.Duration 的一个自定义类型
type MyDuration time.Duration

// Humanize 方法为 MyDuration 类型赋予了新的行为
func (d MyDuration) Humanize() string {
	// 将底层的 time.Duration 转换为更易读的格式
	duration := time.Duration(d)
	if duration.Minutes() < 1 {
		return fmt.Sprintf("%.0f 秒", duration.Seconds())
	}
	return fmt.Sprintf("%.1f 分钟", duration.Minutes())
}

func main() {
	d := MyDuration(125 * time.Second)
	fmt.Println(d.Humanize()) // 输出: 2.1 分钟
}
```

这展示了Go语言组合思想的强大之处：我们可以通过创建新的类型并为其附加方法，来扩展现有类型的能力，而无需继承。

---

## 总结

- 方法是附加到特定类型（接收者）上的函数。
- **指针接收者 (`*T`)** 能够修改原始数据，并且性能更高，是大多数场景下的首选。
- **值接收者 (`T`)** 操作的是数据的副本，适用于小型的、不需要修改的"只读"场景。
- 保持接收者类型（值或指针）的**一致性**是重要的工程实践。
- 方法让我们可以为任何自定义类型赋予行为，这是Go实现封装和构建清晰API的核心机制。 