---
title: 数组、切片和映射
description: 学习Go语言的集合数据类型：数组、切片和映射的使用方法
---

# 数组、切片和映射

在Go语言中，数组、切片和映射是三种重要的集合数据类型。它们各有特点：数组长度固定、切片动态灵活、映射存储键值对。掌握它们是Go编程的基础。

## 本章内容

- 数组：固定长度的元素序列
- 切片：动态数组，Go中最常用的集合类型
- 映射：键值对的集合，类似其他语言的字典
- 集合操作的最佳实践和性能考虑

## 数组 (Array)

数组是固定长度的同类型元素序列，长度是类型的一部分，在编译时确定。

### 数组的特点

- **固定长度**：一旦声明，长度不可改变
- **连续内存**：元素在内存中连续存储，访问效率高
- **值类型**：赋值和传参时会复制整个数组
- **长度是类型的一部分**：`[3]int` 和 `[5]int` 是不同类型

### 数组声明和初始化

```go
// 基本声明方式
var numbers [5]int                    // 零值初始化：[0 0 0 0 0]
scores := [3]int{85, 92, 78}         // 直接初始化
fruits := [...]string{"apple", "banana", "orange"}  // 自动推断长度

// 指定索引初始化
weekdays := [7]string{
    0: "周日",
    1: "周一", 
    6: "周六",  // 其他位置为空字符串
}

// 二维数组
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

### 数组的基本操作

```go
temperatures := [7]float64{22.5, 25.0, 23.8, 26.2, 24.1, 21.9, 20.3}

// 访问和修改
fmt.Printf("今天温度: %.1f°C\n", temperatures[0])
temperatures[0] = 23.0

// 遍历数组
for i, temp := range temperatures {
    fmt.Printf("第%d天: %.1f°C\n", i+1, temp)
}

// 计算统计值
sum := 0.0
for _, temp := range temperatures {
    sum += temp
}
average := sum / float64(len(temperatures))
```

**数组的使用场景**：
- 长度固定且已知的数据集合
- 对性能要求极高的场景（避免指针间接访问）
- 作为切片的底层存储

::: tip 实际开发建议
在Go中，切片比数组更常用，因为切片更灵活。数组主要用于：长度明确的固定集合、高性能计算、作为切片的底层数组。
:::

## 切片 (Slice)

切片是Go中最重要的数据结构之一，它是对数组的抽象，提供了动态数组的功能。

### 切片的内部结构

切片由三部分组成：
- **指针**：指向底层数组的某个元素
- **长度（length）**：切片中元素的个数
- **容量（capacity）**：从切片第一个元素到底层数组末尾的元素个数

```go
// 切片的创建方式
numbers := []int{1, 2, 3, 4, 5}                    // 字面量
scores := make([]int, 3)                           // make，长度3，容量3
grades := make([]int, 3, 5)                        // make，长度3，容量5

// 从数组创建切片
arr := [6]int{10, 20, 30, 40, 50, 60}
slice := arr[1:4]  // [20 30 40]，不包含索引4

fmt.Printf("长度: %d, 容量: %d\n", len(slice), cap(slice))
```

### 切片操作

#### 添加元素

```go
var fruits []string

// append 添加元素
fruits = append(fruits, "apple")
fruits = append(fruits, "banana", "orange")  // 添加多个

// 合并切片
moreFruits := []string{"grape", "mango"}
fruits = append(fruits, moreFruits...)  // 注意...语法
```

#### 删除元素

Go没有内置的删除函数，需要使用切片操作：

```go
slice := []int{1, 2, 3, 4, 5}

// 删除第一个元素
slice = slice[1:]  // [2 3 4 5]

// 删除最后一个元素  
slice = slice[:len(slice)-1]  // [2 3 4]

// 删除中间元素（索引1）
index := 1
slice = append(slice[:index], slice[index+1:]...)  // [2 4]
```

#### 复制切片

```go
original := []int{1, 2, 3, 4, 5}

// 浅复制（共享底层数组）
shallow := original[:]  
shallow[0] = 99  // 会影响original

// 深复制（独立的底层数组）
deep := make([]int, len(original))
copy(deep, original)
deep[0] = 99  // 不会影响original
```

### 切片的容量扩展

当切片容量不足时，`append` 会自动扩展：

```go
slice := make([]int, 0, 2)  // 长度0，容量2

for i := 0; i < 5; i++ {
    fmt.Printf("添加前: len=%d, cap=%d\n", len(slice), cap(slice))
    slice = append(slice, i)
    fmt.Printf("添加后: len=%d, cap=%d\n", len(slice), cap(slice))
}
```

**扩展规律**：
- 容量小于1024时，通常翻倍增长
- 容量大于1024时，增长约25%
- 扩展时会分配新的底层数组并复制数据

::: warning 注意
频繁的切片扩展会影响性能。如果知道最终大小，建议预分配足够的容量。
:::

## 映射 (Map)

映射是键值对的无序集合，类似其他语言中的哈希表、字典。

### 映射的特点

- **键值对存储**：每个键对应一个值
- **键唯一性**：相同键只能有一个值，重复赋值会覆盖
- **无序性**：遍历顺序不固定
- **引用类型**：赋值传递的是引用，不是副本

### 映射的创建和操作

```go
// 创建方式
var ages map[string]int                          // 声明但未初始化，值为nil
ages = make(map[string]int)                      // 使用make初始化

// 字面量初始化
scores := map[string]int{
    "Alice":   95,
    "Bob":     87,
    "Charlie": 92,
}

// 基本操作
ages["Tom"] = 25      // 添加/修改
age := ages["Tom"]    // 读取
delete(ages, "Tom")   // 删除

// 检查键是否存在
if age, exists := ages["Tom"]; exists {
    fmt.Printf("Tom's age: %d\n", age)
} else {
    fmt.Println("Tom not found")
}
```

### 映射的遍历

```go
scores := map[string]int{
    "语文": 85,
    "数学": 92,
    "英语": 78,
}

// 遍历键值对
for subject, score := range scores {
    fmt.Printf("%s: %d分\n", subject, score)
}

// 只遍历键
for subject := range scores {
    fmt.Printf("科目: %s\n", subject)
}

// 只遍历值
for _, score := range scores {
    fmt.Printf("分数: %d\n", score)
}
```

### 映射的高级操作

```go
// 统计字符出现次数
text := "hello world"
charCount := make(map[rune]int)

for _, char := range text {
    charCount[char]++  // 零值特性：int的零值是0
}

for char, count := range charCount {
    fmt.Printf("'%c': %d次\n", char, count)
}

// 按键排序遍历
import "sort"

var keys []string
for k := range scores {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Printf("%s: %d\n", k, scores[k])
}
```

## 实践示例：学生成绩管理系统

让我们用数组、切片和映射实现一个完整的成绩管理系统：

```go
package main

import (
    "fmt"
    "sort"
)

// 学生结构体
type Student struct {
    Name   string
    Scores map[string]int  // 各科成绩
}

func main() {
    // 使用切片存储学生列表
    students := []Student{
        {
            Name: "张三",
            Scores: map[string]int{
                "语文": 85, "数学": 92, "英语": 78,
            },
        },
        {
            Name: "李四", 
            Scores: map[string]int{
                "语文": 76, "数学": 88, "英语": 91,
            },
        },
        {
            Name: "王五",
            Scores: map[string]int{
                "语文": 93, "数学": 79, "英语": 85,
            },
        },
    }
    
    // 1. 计算每个学生的总分和平均分
    fmt.Println("📊 学生成绩统计")
    for _, student := range students {
        total := 0
        count := 0
        for subject, score := range student.Scores {
            total += score
            count++
        }
        average := float64(total) / float64(count)
        fmt.Printf("%s: 总分%d，平均%.1f分\n", student.Name, total, average)
    }
    
    // 2. 统计各科平均分
    fmt.Println("\n📈 各科平均分")
    subjectTotals := make(map[string]int)
    subjectCounts := make(map[string]int)
    
    for _, student := range students {
        for subject, score := range student.Scores {
            subjectTotals[subject] += score
            subjectCounts[subject]++
        }
    }
    
    // 按科目名排序输出
    var subjects []string
    for subject := range subjectTotals {
        subjects = append(subjects, subject)
    }
    sort.Strings(subjects)
    
    for _, subject := range subjects {
        average := float64(subjectTotals[subject]) / float64(subjectCounts[subject])
        fmt.Printf("%s: %.1f分\n", subject, average)
    }
    
    // 3. 找出各科最高分
    fmt.Println("\n🏆 各科最高分")
    subjectMax := make(map[string]int)
    subjectToppers := make(map[string]string)
    
    for _, student := range students {
        for subject, score := range student.Scores {
            if score > subjectMax[subject] {
                subjectMax[subject] = score
                subjectToppers[subject] = student.Name
            }
        }
    }
    
    for _, subject := range subjects {
        fmt.Printf("%s: %s（%d分）\n", 
            subject, subjectToppers[subject], subjectMax[subject])
    }
    
    // 4. 成绩分布统计
    fmt.Println("\n📋 成绩分布")
    gradeRanges := map[string]int{
        "优秀(90-100)": 0,
        "良好(80-89)":  0,
        "中等(70-79)":  0,
        "及格(60-69)":  0,
        "不及格(<60)":   0,
    }
    
    for _, student := range students {
        for _, score := range student.Scores {
            switch {
            case score >= 90:
                gradeRanges["优秀(90-100)"]++
            case score >= 80:
                gradeRanges["良好(80-89)"]++
            case score >= 70:
                gradeRanges["中等(70-79)"]++
            case score >= 60:
                gradeRanges["及格(60-69)"]++
            default:
                gradeRanges["不及格(<60)"]++
            }
        }
    }
    
    for grade, count := range gradeRanges {
        if count > 0 {
            fmt.Printf("%s: %d人次\n", grade, count)
        }
    }
}
```

## 本章小结

通过本章学习，你应该掌握：

### 核心概念
- **数组**：固定长度，值类型，长度是类型的一部分
- **切片**：动态长度，引用类型，Go中最常用的集合
- **映射**：键值对集合，无序，键唯一

### 使用场景对比

| 类型 | 使用场景 | 优点 | 缺点 |
|------|----------|------|------|
| 数组 | 固定长度数据，高性能场景 | 内存连续，访问快速 | 长度固定，不够灵活 |
| 切片 | 动态数据集合，日常开发 | 灵活动态，功能丰富 | 有一定内存开销 |
| 映射 | 键值对存储，快速查找 | 查找效率高，使用简单 | 无序，内存开销较大 |

### 最佳实践
1. **优先使用切片**：除非明确需要固定长度，否则选择切片
2. **预分配容量**：知道大致大小时，使用 `make([]T, 0, capacity)` 预分配
3. **检查映射键存在性**：使用 `value, ok := map[key]` 检查
4. **注意共享底层数组**：切片操作可能影响其他切片

### 性能考虑
- 切片扩展有成本，预分配容量可提升性能
- 映射的查找、插入、删除都是 O(1) 平均时间复杂度
- 大切片的复制和传递成本较低（只复制切片头部）

::: tip 练习建议
尝试实现一个简单的通讯录程序，练习切片存储联系人、映射分类管理、数组处理固定字段等操作。
:::