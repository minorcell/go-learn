# 数组、切片和映射

在Go语言中，数组、切片和映射是三种重要的集合数据类型，用于存储和操作多个相关的数据。

## 本章内容

- 数组：固定长度的元素序列
- 切片：动态数组，更加灵活
- 映射：键值对的集合
- 集合操作和遍历技巧

## 数组 (Array)

数组是固定长度的同类型元素序列，在声明时就确定了大小。

### 数组声明和初始化

```go
package main

import "fmt"

func main() {
    // 声明数组
    var numbers [5]int // 默认零值初始化
    fmt.Printf("默认数组: %v\n", numbers)
    
    // 直接初始化
    scores := [3]int{85, 92, 78}
    fmt.Printf("分数数组: %v\n", scores)
    
    // 让编译器推断长度
    fruits := [...]string{"apple", "banana", "orange"}
    fmt.Printf("水果数组: %v, 长度: %d\n", fruits, len(fruits))
    
    // 指定索引初始化
    weekdays := [7]string{
        0: "周日",
        1: "周一",
        2: "周二",
        6: "周六", // 其他位置为零值
    }
    fmt.Printf("星期数组: %v\n", weekdays)
    
    // 二维数组
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Printf("矩阵:\n")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
}
```

### 数组操作

```go
package main

import "fmt"

func main() {
    temperatures := [7]float64{22.5, 25.0, 23.8, 26.2, 24.1, 21.9, 20.3}
    
    // 访问和修改元素
    fmt.Printf("今天温度: %.1f°C\n", temperatures[0])
    temperatures[0] = 23.0
    fmt.Printf("修正后: %.1f°C\n", temperatures[0])
    
    // 计算平均温度
    sum := 0.0
    for _, temp := range temperatures {
        sum += temp
    }
    average := sum / float64(len(temperatures))
    fmt.Printf("平均温度: %.1f°C\n", average)
    
    // 找最高和最低温度
    max, min := temperatures[0], temperatures[0]
    maxDay, minDay := 0, 0
    
    for i, temp := range temperatures {
        if temp > max {
            max = temp
            maxDay = i
        }
        if temp < min {
            min = temp
            minDay = i
        }
    }
    
    days := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
    fmt.Printf("最高温度: %.1f°C (%s)\n", max, days[maxDay])
    fmt.Printf("最低温度: %.1f°C (%s)\n", min, days[minDay])
    
    // 数组比较（相同类型和长度才能比较）
    week1 := [3]int{1, 2, 3}
    week2 := [3]int{1, 2, 3}
    week3 := [3]int{3, 2, 1}
    
    fmt.Printf("week1 == week2: %t\n", week1 == week2)
    fmt.Printf("week1 == week3: %t\n", week1 == week3)
}
```

## 切片 (Slice)

切片是对数组的一个连续片段的引用，长度可变，使用更加灵活。

### 切片创建和基本操作

```go
package main

import "fmt"

func main() {
    // 直接创建切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("切片: %v, 长度: %d, 容量: %d\n", 
        numbers, len(numbers), cap(numbers))
    
    // 使用 make 创建切片
    scores := make([]int, 3)    // 长度为3，容量为3
    fmt.Printf("make切片: %v, 长度: %d, 容量: %d\n", 
        scores, len(scores), cap(scores))
    
    grades := make([]int, 3, 5) // 长度为3，容量为5
    fmt.Printf("指定容量: %v, 长度: %d, 容量: %d\n", 
        grades, len(grades), cap(grades))
    
    // 从数组创建切片
    arr := [6]int{10, 20, 30, 40, 50, 60}
    slice1 := arr[1:4]  // 索引1到3（不包含4）
    slice2 := arr[:3]   // 从开始到索引2
    slice3 := arr[2:]   // 从索引2到结束
    slice4 := arr[:]    // 整个数组
    
    fmt.Printf("原数组: %v\n", arr)
    fmt.Printf("arr[1:4]: %v\n", slice1)
    fmt.Printf("arr[:3]: %v\n", slice2)
    fmt.Printf("arr[2:]: %v\n", slice3)
    fmt.Printf("arr[:]: %v\n", slice4)
    
    // 修改切片会影响底层数组
    slice1[0] = 99
    fmt.Printf("修改后数组: %v\n", arr)
    fmt.Printf("修改后切片: %v\n", slice1)
}
```

### 切片的动态操作

```go
package main

import "fmt"

func main() {
    // 动态添加元素
    var fruits []string
    fmt.Printf("初始切片: %v, 长度: %d\n", fruits, len(fruits))
    
    // append 添加元素
    fruits = append(fruits, "apple")
    fruits = append(fruits, "banana", "orange")
    fmt.Printf("添加后: %v, 长度: %d, 容量: %d\n", 
        fruits, len(fruits), cap(fruits))
    
    // 添加另一个切片
    moreFruits := []string{"grape", "mango"}
    fruits = append(fruits, moreFruits...)
    fmt.Printf("合并后: %v, 长度: %d, 容量: %d\n", 
        fruits, len(fruits), cap(fruits))
    
    // 删除元素（Go没有内置delete函数）
    fmt.Println("\n删除操作：")
    
    // 删除第一个元素
    if len(fruits) > 0 {
        fruits = fruits[1:]
        fmt.Printf("删除第一个: %v\n", fruits)
    }
    
    // 删除最后一个元素
    if len(fruits) > 0 {
        fruits = fruits[:len(fruits)-1]
        fmt.Printf("删除最后一个: %v\n", fruits)
    }
    
    // 删除中间元素（索引2）
    if len(fruits) > 2 {
        index := 2
        fruits = append(fruits[:index], fruits[index+1:]...)
        fmt.Printf("删除索引%d: %v\n", index, fruits)
    }
    
    // 插入元素
    fmt.Println("\n插入操作：")
    insertIndex := 1
    newFruit := "kiwi"
    
    // 在指定位置插入
    fruits = append(fruits[:insertIndex], 
        append([]string{newFruit}, fruits[insertIndex:]...)...)
    fmt.Printf("在索引%d插入%s: %v\n", insertIndex, newFruit, fruits)
}
```

### 切片的复制和扩容

```go
package main

import "fmt"

func main() {
    // 切片复制
    original := []int{1, 2, 3, 4, 5}
    
    // 浅复制（共享底层数组）
    shallow := original
    shallow[0] = 99
    fmt.Printf("原切片: %v\n", original) // 也被修改了
    fmt.Printf("浅复制: %v\n", shallow)
    
    // 深复制
    original = []int{1, 2, 3, 4, 5} // 重置
    deep := make([]int, len(original))
    copy(deep, original)
    deep[0] = 88
    fmt.Printf("原切片: %v\n", original) // 不受影响
    fmt.Printf("深复制: %v\n", deep)
    
    // 部分复制
    source := []int{10, 20, 30, 40, 50, 60}
    dest := make([]int, 3)
    n := copy(dest, source[2:5]) // 复制索引2-4的元素
    fmt.Printf("源切片: %v\n", source)
    fmt.Printf("目标切片: %v, 复制了%d个元素\n", dest, n)
    
    // 切片扩容演示
    fmt.Println("\n扩容演示：")
    nums := make([]int, 0, 2)
    fmt.Printf("初始: 长度=%d, 容量=%d\n", len(nums), cap(nums))
    
    for i := 1; i <= 8; i++ {
        nums = append(nums, i)
        fmt.Printf("添加%d: 长度=%d, 容量=%d, 切片=%v\n", 
            i, len(nums), cap(nums), nums)
    }
    
    // 切片表达式的完整形式
    fmt.Println("\n切片表达式：")
    data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // slice[low:high:max] - 限制容量
    s1 := data[2:5:6] // 从索引2到4，容量限制为6-2=4
    fmt.Printf("data[2:5:6]: %v, 长度:%d, 容量:%d\n", 
        s1, len(s1), cap(s1))
    
    s2 := data[2:5] // 默认容量到底层数组末尾
    fmt.Printf("data[2:5]: %v, 长度:%d, 容量:%d\n", 
        s2, len(s2), cap(s2))
}
```

## 映射 (Map)

映射是键值对的无序集合，类似于其他语言的字典或哈希表。

### 映射的创建和基本操作

```go
package main

import "fmt"

func main() {
    // 创建映射的几种方式
    
    // 1. 使用 make
    scores := make(map[string]int)
    scores["Alice"] = 95
    scores["Bob"] = 87
    scores["Charlie"] = 92
    fmt.Printf("成绩映射: %v\n", scores)
    
    // 2. 映射字面量
    capitals := map[string]string{
        "China":   "Beijing",
        "Japan":   "Tokyo",
        "France":  "Paris",
        "Germany": "Berlin",
    }
    fmt.Printf("首都映射: %v\n", capitals)
    
    // 3. 空映射
    inventory := map[string]int{}
    inventory["apples"] = 50
    inventory["bananas"] = 30
    fmt.Printf("库存映射: %v\n", inventory)
    
    // 访问元素
    fmt.Printf("Alice的分数: %d\n", scores["Alice"])
    fmt.Printf("中国的首都: %s\n", capitals["China"])
    
    // 检查键是否存在
    if score, exists := scores["David"]; exists {
        fmt.Printf("David的分数: %d\n", score)
    } else {
        fmt.Println("David不在分数列表中")
    }
    
    // 修改值
    scores["Alice"] = 98
    fmt.Printf("Alice的新分数: %d\n", scores["Alice"])
    
    // 删除元素
    delete(scores, "Bob")
    fmt.Printf("删除Bob后: %v\n", scores)
    
    // 获取映射长度
    fmt.Printf("映射长度: %d\n", len(scores))
}
```

### 映射的遍历和操作

```go
package main

import "fmt"

func main() {
    // 学生信息映射
    students := map[int]map[string]interface{}{
        1001: {
            "name":  "张三",
            "age":   20,
            "grade": "A",
            "score": 95,
        },
        1002: {
            "name":  "李四",
            "age":   19,
            "grade": "B",
            "score": 87,
        },
        1003: {
            "name":  "王五",
            "age":   21,
            "grade": "A",
            "score": 92,
        },
    }
    
    // 遍历映射
    fmt.Println("学生信息：")
    for id, info := range students {
        fmt.Printf("学号 %d:\n", id)
        for key, value := range info {
            fmt.Printf("  %s: %v\n", key, value)
        }
        fmt.Println()
    }
    
    // 统计信息
    totalScore := 0
    gradeCount := make(map[string]int)
    ageSum := 0
    
    for _, student := range students {
        if score, ok := student["score"].(int); ok {
            totalScore += score
        }
        if grade, ok := student["grade"].(string); ok {
            gradeCount[grade]++
        }
        if age, ok := student["age"].(int); ok {
            ageSum += age
        }
    }
    
    fmt.Printf("统计结果：\n")
    fmt.Printf("平均分: %.1f\n", float64(totalScore)/float64(len(students)))
    fmt.Printf("平均年龄: %.1f\n", float64(ageSum)/float64(len(students)))
    fmt.Printf("等级分布: %v\n", gradeCount)
    
    // 映射的键值集合
    fmt.Println("\n所有学号：")
    for id := range students {
        fmt.Printf("%d ", id)
    }
    fmt.Println()
    
    // 条件查询
    fmt.Println("\nA级学生：")
    for id, student := range students {
        if grade, ok := student["grade"].(string); ok && grade == "A" {
            fmt.Printf("学号 %d: %s\n", id, student["name"])
        }
    }
}
```

### 映射的高级用法

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 单词计数器
    text := "hello world hello go world go programming"
    words := []string{"hello", "world", "hello", "go", "world", "go", "programming"}
    
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    
    fmt.Printf("单词计数: %v\n", wordCount)
    
    // 按值排序（需要转换为切片）
    type wordPair struct {
        word  string
        count int
    }
    
    var pairs []wordPair
    for word, count := range wordCount {
        pairs = append(pairs, wordPair{word, count})
    }
    
    // 按计数降序排序
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].count > pairs[j].count
    })
    
    fmt.Println("\n按频率排序：")
    for _, pair := range pairs {
        fmt.Printf("%s: %d次\n", pair.word, pair.count)
    }
    
    // 反向映射
    fmt.Println("\n反向映射示例：")
    colorCodes := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    
    // 创建反向映射
    codeColors := make(map[string]string)
    for color, code := range colorCodes {
        codeColors[code] = color
    }
    
    fmt.Printf("颜色 -> 代码: %v\n", colorCodes)
    fmt.Printf("代码 -> 颜色: %v\n", codeColors)
    
    // 映射嵌套
    fmt.Println("\n映射嵌套示例：")
    menu := map[string]map[string]float64{
        "breakfast": {
            "pancakes": 8.50,
            "eggs":     6.00,
            "coffee":   3.50,
        },
        "lunch": {
            "sandwich": 12.00,
            "salad":    10.50,
            "soup":     8.00,
        },
        "dinner": {
            "steak":  25.00,
            "pasta":  18.00,
            "wine":   15.00,
        },
    }
    
    for meal, dishes := range menu {
        fmt.Printf("%s菜单:\n", meal)
        total := 0.0
        for dish, price := range dishes {
            fmt.Printf("  %s: $%.2f\n", dish, price)
            total += price
        }
        fmt.Printf("  小计: $%.2f\n\n", total)
    }
}
```

## 实战练习

### 学生成绩分析系统

```go
package main

import (
    "fmt"
    "sort"
)

// 学生结构
type Student struct {
    ID     int
    Name   string
    Scores map[string]int // 科目 -> 分数
}

// 成绩分析器
type GradeAnalyzer struct {
    Students []Student
    Subjects []string
}

// 添加学生
func (ga *GradeAnalyzer) AddStudent(student Student) {
    ga.Students = append(ga.Students, student)
    
    // 更新科目列表
    for subject := range student.Scores {
        found := false
        for _, s := range ga.Subjects {
            if s == subject {
                found = true
                break
            }
        }
        if !found {
            ga.Subjects = append(ga.Subjects, subject)
        }
    }
}

// 计算学生总分
func (ga *GradeAnalyzer) GetStudentTotal(studentID int) int {
    for _, student := range ga.Students {
        if student.ID == studentID {
            total := 0
            for _, score := range student.Scores {
                total += score
            }
            return total
        }
    }
    return 0
}

// 计算学生平均分
func (ga *GradeAnalyzer) GetStudentAverage(studentID int) float64 {
    for _, student := range ga.Students {
        if student.ID == studentID {
            if len(student.Scores) == 0 {
                return 0
            }
            total := 0
            for _, score := range student.Scores {
                total += score
            }
            return float64(total) / float64(len(student.Scores))
        }
    }
    return 0
}

// 按总分排名
func (ga *GradeAnalyzer) GetRankingByTotal() []struct {
    Student Student
    Total   int
    Average float64
} {
    type studentStat struct {
        Student Student
        Total   int
        Average float64
    }
    
    var stats []studentStat
    for _, student := range ga.Students {
        total := ga.GetStudentTotal(student.ID)
        average := ga.GetStudentAverage(student.ID)
        stats = append(stats, studentStat{student, total, average})
    }
    
    // 按总分降序排序
    sort.Slice(stats, func(i, j int) bool {
        return stats[i].Total > stats[j].Total
    })
    
    return stats
}

// 科目分析
func (ga *GradeAnalyzer) GetSubjectAnalysis(subject string) map[string]interface{} {
    var scores []int
    studentCount := 0
    
    for _, student := range ga.Students {
        if score, exists := student.Scores[subject]; exists {
            scores = append(scores, score)
            studentCount++
        }
    }
    
    if len(scores) == 0 {
        return nil
    }
    
    // 计算统计信息
    total := 0
    max := scores[0]
    min := scores[0]
    
    for _, score := range scores {
        total += score
        if score > max {
            max = score
        }
        if score < min {
            min = score
        }
    }
    
    average := float64(total) / float64(len(scores))
    
    // 计算及格率
    passCount := 0
    for _, score := range scores {
        if score >= 60 {
            passCount++
        }
    }
    passRate := float64(passCount) / float64(len(scores)) * 100
    
    return map[string]interface{}{
        "subject":    subject,
        "count":      len(scores),
        "average":    average,
        "max":        max,
        "min":        min,
        "passRate":   passRate,
        "scores":     scores,
    }
}

// 打印详细报告
func (ga *GradeAnalyzer) PrintDetailedReport() {
    fmt.Println("=== 学生成绩分析报告 ===\n")
    
    // 学生排名
    fmt.Println("1. 学生总分排名：")
    rankings := ga.GetRankingByTotal()
    for i, stat := range rankings {
        fmt.Printf("%d. %s (ID: %d) - 总分: %d, 平均: %.1f\n",
            i+1, stat.Student.Name, stat.Student.ID, 
            stat.Total, stat.Average)
    }
    
    // 科目分析
    fmt.Println("\n2. 科目分析：")
    for _, subject := range ga.Subjects {
        analysis := ga.GetSubjectAnalysis(subject)
        if analysis != nil {
            fmt.Printf("\n%s:\n", subject)
            fmt.Printf("  参与人数: %d\n", analysis["count"])
            fmt.Printf("  平均分: %.1f\n", analysis["average"])
            fmt.Printf("  最高分: %d\n", analysis["max"])
            fmt.Printf("  最低分: %d\n", analysis["min"])
            fmt.Printf("  及格率: %.1f%%\n", analysis["passRate"])
        }
    }
    
    // 个人详情
    fmt.Println("\n3. 学生详细成绩：")
    for _, student := range ga.Students {
        fmt.Printf("\n%s (ID: %d):\n", student.Name, student.ID)
        total := 0
        for subject, score := range student.Scores {
            fmt.Printf("  %s: %d分\n", subject, score)
            total += score
        }
        average := float64(total) / float64(len(student.Scores))
        fmt.Printf("  总分: %d, 平均: %.1f\n", total, average)
    }
}

func main() {
    // 创建分析器
    analyzer := &GradeAnalyzer{}
    
    // 添加学生数据
    students := []Student{
        {
            ID:   1001,
            Name: "张三",
            Scores: map[string]int{
                "数学": 95,
                "语文": 88,
                "英语": 92,
                "物理": 89,
            },
        },
        {
            ID:   1002,
            Name: "李四",
            Scores: map[string]int{
                "数学": 78,
                "语文": 85,
                "英语": 80,
                "物理": 72,
            },
        },
        {
            ID:   1003,
            Name: "王五",
            Scores: map[string]int{
                "数学": 92,
                "语文": 90,
                "英语": 88,
                "物理": 94,
            },
        },
        {
            ID:   1004,
            Name: "赵六",
            Scores: map[string]int{
                "数学": 67,
                "语文": 75,
                "英语": 70,
                "物理": 65,
            },
        },
    }
    
    // 添加所有学生
    for _, student := range students {
        analyzer.AddStudent(student)
    }
    
    // 打印详细报告
    analyzer.PrintDetailedReport()
}
```

##  本章小结

在这一章中，我们学习了Go语言的三种集合类型：

### 数组 (Array)
- **固定长度** - 声明时确定大小，不可改变
- **值类型** - 赋值和传参是复制整个数组
- **性能好** - 内存连续，访问效率高
- **使用场景** - 长度固定的小规模数据

### 切片 (Slice)  
- **动态长度** - 可以动态增长和缩减
- **引用类型** - 底层引用数组
- **灵活操作** - append、copy等丰富操作
- **使用场景** - 大多数情况下的首选

### 映射 (Map)
- **键值对** - 无序的键值对集合  
- **引用类型** - 零值为nil
- **快速查找** - 基于哈希表实现
- **使用场景** - 需要快速查找的关联数据

### 最佳实践
- 优先使用切片而不是数组
- 映射要先初始化再使用
- 注意引用类型的共享特性
- 合理使用容量避免频繁扩容