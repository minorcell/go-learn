# 数组、切片和映射：Go的数据工具箱

> 在上一篇文章中，我们学会了如何使用变量和基本类型来处理单个数据。但现实世界的程序需要管理一系列的数据：一个用户列表、一天的温度记录、一本书的词汇表。Go提供了三种强大的内置数据结构来组织这些集合：数组（Array）、切片（Slice）和映射（Map）。

本文将以一个"工具箱"的视角来探索它们。每一种结构都是一个为特定任务而设计的工具。理解它们的特性和适用场景，是写出高效、清晰的Go代码的关键。

---

## 工具一：数组 (Array) - 固定尺寸的工具盒

想象一个药盒，它有固定数量的格子，比如一周七天，不多不少。**数组**就是这样一个**固定长度**的、存储相同类型元素的容器。

它的长度是其类型的一部分。这意味着 `[4]int` 和 `[5]int` 是完全不同的类型。

### 何时使用数组？

当你明确知道需要存储的元素数量，并且这个数量不会改变时，数组是最佳选择。例如，处理颜色的RGBA值（4个元素）或表示一个棋盘（一个8x8的二维数组）。

```go
package main

import "fmt"

func main() {
	// 声明一个包含4个整数的数组
	// 它的所有元素被自动初始化为零值 0
	var rgba [4]int

	rgba[0] = 255 // Red
	rgba[1] = 128 // Green
	rgba[2] = 0   // Blue
	rgba[3] = 255 // Alpha

	fmt.Printf("颜色值: R=%d, G=%d, B=%d, A=%d\n", rgba[0], rgba[1], rgba[2], rgba[3])

	// 使用数组字面量进行声明和初始化
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("前5个素数:", primes)
	
	// 让编译器自动计算长度
	weekdays := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("一周有 %d 天\n", len(weekdays))
}
```

**核心特点**:
- **固定长度**: 一旦声明，长度不可变。
- **值类型**: 将一个数组赋值给另一个数组，会完整复制所有元素。
- **类型安全**: 长度是类型的一部分，`[4]int`类型的变量不能赋值给`[5]int`。

数组在Go中并不常用，因为它缺乏灵活性。它更像是一个底层构建块，为接下来要介绍的、更有用的工具——切片——提供存储空间。

---

## 工具二：切片 (Slice) - 灵活的动态扳手

如果说数组是固定大小的工具盒，那么**切片（Slice）**就是一把可以根据螺母大小自由伸缩的活动扳手。它是Go中使用最广泛、功能最强大的数据结构。

**切片不是数组，它是对一个数组片段的描述**。它提供了一种动态的、灵活的方式来访问底层数组的连续部分。

### 切片的内部结构：一个三元组

要真正理解切片，你需要了解它的"头部"结构。每个切片变量内部都包含三个信息：

1.  **指针 (Pointer)**: 指向底层数组中该切片所代表的第一个元素。
2.  **长度 (Length)**: 切片中元素的数量。通过 `len()` 函数获取。
3.  **容量 (Capacity)**: 从切片的起始元素到底层数组末尾的元素数量。通过 `cap()` 函数获取。

![Slice Internals](https://go.dev/blog/images/slices_diagram.png)
*图片来源: The Go Blog*

让我们通过代码来理解：

```go
package main

import "fmt"

func main() {
	// 创建一个包含10个元素的底层数组
	underlyingArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 创建一个切片，它"看"向底层数组的一部分
	// 从索引2开始，到索引5结束（不包括5）
	mySlice := underlyingArray[2:5]

	fmt.Printf("切片内容: %v\n", mySlice)
	fmt.Printf("长度 (len): %d\n", len(mySlice)) // 5 - 2 = 3
	fmt.Printf("容量 (cap): %d\n", cap(mySlice)) // 10 - 2 = 8
	
	// 修改切片的元素
	mySlice[0] = 99
	
	// 底层数组也随之改变，因为它们共享内存
	fmt.Printf("修改后，切片内容: %v\n", mySlice)
	fmt.Printf("修改后，底层数组: %v\n", underlyingArray)
}
```

### 创建和操作切片

通常，我们不直接创建数组，而是使用 `make` 函数来创建切片，它会自动为我们分配和管理底层数组。

```go
// 创建一个长度为5，容量也为5的切片
s1 := make([]string, 5)

// 创建一个长度为0，但容量为5的切片
// 这很高效，因为我们预留了空间，后续添加元素时可能无需重新分配内存
s2 := make([]string, 0, 5)
```

### `append`: 切片的增长魔法

`append` 是一个内置函数，用于向切片末尾添加元素。这是理解切片动态性的关键。

```go
package main

import "fmt"

func main() {
	var fruits []string // 这是一个 nil 切片，长度和容量都是0
	
	fmt.Printf("初始: len=%d, cap=%d, %v\n", len(fruits), cap(fruits), fruits)

	// 1. 添加第一个元素
	// 底层容量不足，append会分配一个新的、更大的底层数组
	fruits = append(fruits, "apple")
	fmt.Printf("添加apple: len=%d, cap=%d, %v\n", len(fruits), cap(fruits), fruits)

	// 2. 继续添加
	// 容量足够，直接在原底层数组上添加
	fruits = append(fruits, "banana", "orange")
	fmt.Printf("添加两种水果: len=%d, cap=%d, %v\n", len(fruits), cap(fruits), fruits)

	// 3. 当容量再次用尽时，会再次分配一个更大的新数组
	// Go的运行时通常会以翻倍的策略来扩展容量，以摊销分配成本
	fruits = append(fruits, "grape", "mango")
	fmt.Printf("容量耗尽后添加: len=%d, cap=%d, %v\n", len(fruits), cap(fruits), fruits)
}
```
**关键点**:
- `append` 返回一个新的切片。**你必须将`append`的结果重新赋值给原切片变量** (`slice = append(slice, ...)`), 因为当容量不足时，它会指向一个全新的底层数组。
- 切片是**引用类型**。函数传递的是切片头的副本，但指针指向同一个底层数组。这意味着在函数内修改切片元素会影响到原始切片。

---

## 工具三：映射 (Map) - 带标签的抽屉柜

当你需要通过一个唯一的"标签"（键）来快速查找一个"物品"（值）时，**映射（Map）** 就是你的不二之选。它是一个无序的**键值对 (key-value)** 集合。

### 何时使用映射？

- 存储用户ID和用户信息的对应关系。
- 统计文本中每个单词出现的次数。
- 配置文件的键值存储。

```go
package main

import "fmt"

func main() {
	// 创建一个映射，键是string类型，值是int类型
	ages := make(map[string]int)

	// 添加键值对
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Println("Bob的年龄是:", ages["Bob"])

	// 删除一个键
	delete(ages, "Charlie")
	fmt.Println("删除Charlie后:", ages)

	// 检查一个键是否存在
	// "comma ok" idiom
	age, ok := ages["David"]
	if !ok {
		fmt.Println("David的年龄未知")
	} else {
		fmt.Println("David的年龄是:", age)
	}

	// 声明并初始化
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"China":  "Beijing",
	}
	fmt.Println("日本的首都是:", capitals["Japan"])
	
	// 遍历映射 (注意：遍历顺序是随机的)
	for country, capital := range capitals {
		fmt.Printf("%s 的首都是 %s\n", country, capital)
	}
}
```

**核心特点**:
- **键值对**: 每个元素都由一个唯一的键和一个值组成。
- **无序性**: 当你遍历一个映射时，Go不保证每次的顺序都一样。
- **引用类型**: 和切片一样，映射也是引用类型。
- **零值是 `nil`**: 一个未初始化的映射是 `nil`。向一个 `nil` 映射写入数据会导致运行时恐慌 (panic)。必须先用 `make` 初始化。

---

## 总结：如何选择你的工具？

| 场景 | 推荐工具 | 理由 |
| :--- | :--- | :--- |
| 需要存储固定数量的元素，如月份、星期 | **数组 (Array)** | 长度固定且在编译期可知，简单高效。 |
| 处理一个可变长度的序列，如用户列表、日志条目 | **切片 (Slice)** | Go的默认选择，灵活、强大，拥有`append`等便捷操作。 |
| 需要通过唯一的标识符来快速查找数据 | **映射 (Map)** | 基于键的快速查找 (`O(1)`)，是构建关联数据的理想选择。 |

掌握这三种数据结构，你就拥有了构建复杂Go程序所需的核心工具。在下一篇文章中，我们将学习**控制流**，看看如何让我们的程序根据不同的条件执行不同的逻辑。 