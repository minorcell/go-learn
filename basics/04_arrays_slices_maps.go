package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
04_arrays_slices_maps.go - Go语言基础：数组、切片和映射
学习内容：
1. 数组的声明和使用
2. 切片的创建和操作
3. 映射的使用
4. 常用操作和技巧
*/

func main() {
	fmt.Println("=== Go语言基础：数组、切片和映射 ===")

	// 1. 数组基础
	fmt.Println("\n1. 数组基础：")

	// 数组声明和初始化
	var arr1 [5]int                // 声明一个长度为5的int数组，零值为0
	arr2 := [5]int{1, 2, 3, 4, 5}  // 声明并初始化
	arr3 := [...]int{10, 20, 30}   // 自动推断长度
	arr4 := [5]int{1: 100, 3: 300} // 指定索引初始化

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v (长度: %d)\n", arr3, len(arr3))
	fmt.Printf("arr4: %v\n", arr4)

	// 数组操作
	arr2[0] = 99
	fmt.Printf("修改后的arr2: %v\n", arr2)

	// 遍历数组
	fmt.Print("遍历arr2: ")
	for i, v := range arr2 {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// 多维数组
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("3x3矩阵:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 2. 切片基础
	fmt.Println("\n2. 切片基础：")

	// 切片的创建
	var slice1 []int               // 声明切片，初始为nil
	slice2 := []int{1, 2, 3, 4, 5} // 直接初始化
	slice3 := make([]int, 5)       // 使用make创建，长度为5
	slice4 := make([]int, 3, 10)   // 长度为3，容量为10

	fmt.Printf("slice1: %v (长度: %d, 容量: %d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v (长度: %d, 容量: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %v (长度: %d, 容量: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v (长度: %d, 容量: %d)\n", slice4, len(slice4), cap(slice4))

	// 从数组创建切片
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice5 := arr[1:4] // 从索引1到3（不包含4）
	slice6 := arr[:3]  // 从开始到索引2
	slice7 := arr[2:]  // 从索引2到结束
	slice8 := arr[:]   // 整个数组

	fmt.Printf("原数组: %v\n", arr)
	fmt.Printf("arr[1:4]: %v\n", slice5)
	fmt.Printf("arr[:3]: %v\n", slice6)
	fmt.Printf("arr[2:]: %v\n", slice7)
	fmt.Printf("arr[:]: %v\n", slice8)

	// 3. 切片操作
	fmt.Println("\n3. 切片操作：")

	numbers := []int{1, 2, 3}
	fmt.Printf("原切片: %v\n", numbers)

	// append操作
	numbers = append(numbers, 4)       // 添加一个元素
	numbers = append(numbers, 5, 6, 7) // 添加多个元素
	fmt.Printf("append后: %v\n", numbers)

	// 合并切片
	more := []int{8, 9, 10}
	numbers = append(numbers, more...) // 展开切片添加
	fmt.Printf("合并后: %v\n", numbers)

	// copy操作
	backup := make([]int, len(numbers))
	copy(backup, numbers)
	fmt.Printf("复制的切片: %v\n", backup)

	// 删除元素（删除索引为2的元素）
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Printf("删除索引%d后: %v\n", index, numbers)

	// 插入元素（在索引2插入100）
	index = 2
	value := 100
	numbers = append(numbers[:index+1], numbers[index:]...)
	numbers[index] = value
	fmt.Printf("在索引%d插入%d后: %v\n", index, value, numbers)

	// 4. 切片排序和搜索
	fmt.Println("\n4. 切片排序和搜索：")

	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("原数据: %v\n", data)

	// 排序
	sort.Ints(data)
	fmt.Printf("排序后: %v\n", data)

	// 搜索
	target := 25
	index = sort.SearchInts(data, target)
	if index < len(data) && data[index] == target {
		fmt.Printf("找到 %d 在索引 %d\n", target, index)
	} else {
		fmt.Printf("未找到 %d\n", target)
	}

	// 字符串切片排序
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("原字符串: %v\n", words)
	sort.Strings(words)
	fmt.Printf("排序后: %v\n", words)

	// 5. 映射基础
	fmt.Println("\n5. 映射基础：")

	// 映射的创建
	var map1 map[string]int      // 声明映射，初始为nil
	map2 := make(map[string]int) // 使用make创建
	map3 := map[string]int{      // 直接初始化
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("map3: %v\n", map3)

	// 映射操作
	if map1 == nil {
		map1 = make(map[string]int)
	}
	map1["go"] = 100
	map1["python"] = 90
	map1["java"] = 85

	fmt.Printf("添加元素后的map1: %v\n", map1)

	// 6. 映射操作详解
	fmt.Println("\n6. 映射操作详解：")

	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"David":   78,
	}

	// 访问元素
	aliceScore := scores["Alice"]
	fmt.Printf("Alice的分数: %d\n", aliceScore)

	// 安全访问（检查键是否存在）
	score, exists := scores["Eve"]
	if exists {
		fmt.Printf("Eve的分数: %d\n", score)
	} else {
		fmt.Println("Eve不在记录中")
	}

	// 修改元素
	scores["Bob"] = 90
	fmt.Printf("Bob更新后的分数: %d\n", scores["Bob"])

	// 删除元素
	delete(scores, "David")
	fmt.Printf("删除David后: %v\n", scores)

	// 遍历映射
	fmt.Println("所有学生的分数:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// 只遍历键
	fmt.Print("所有学生姓名: ")
	for name := range scores {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// 7. 嵌套数据结构
	fmt.Println("\n7. 嵌套数据结构：")

	// 切片的切片（二维切片）
	matrix2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("二维切片:")
	for i, row := range matrix2D {
		fmt.Printf("行%d: %v\n", i, row)
	}

	// 映射的映射
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English": 87,
			"Science": 92,
		},
		"Bob": {
			"Math":    78,
			"English": 84,
			"Science": 89,
		},
	}

	fmt.Println("学生成绩:")
	for student, subjects := range students {
		fmt.Printf("%s的成绩:\n", student)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}

	// 切片的映射
	wordGroups := map[string][]string{
		"fruits":     {"apple", "banana", "orange"},
		"vegetables": {"carrot", "broccoli", "spinach"},
		"colors":     {"red", "blue", "green", "yellow"},
	}

	fmt.Println("单词分组:")
	for category, words := range wordGroups {
		fmt.Printf("%s: %s\n", category, strings.Join(words, ", "))
	}

	// 8. 实用示例
	fmt.Println("\n8. 实用示例：")

	// 统计字符出现次数
	text := "hello world"
	charCount := make(map[rune]int)
	for _, char := range text {
		charCount[char]++
	}

	fmt.Printf("字符串 \"%s\" 中字符出现次数:\n", text)
	for char, count := range charCount {
		if char != ' ' {
			fmt.Printf("  '%c': %d次\n", char, count)
		}
	}

	// 查找重复元素
	nums := []int{1, 2, 3, 2, 4, 3, 5, 1}
	seen := make(map[int]bool)
	duplicates := []int{}

	for _, num := range nums {
		if seen[num] {
			duplicates = append(duplicates, num)
		} else {
			seen[num] = true
		}
	}

	fmt.Printf("数组 %v 中的重复元素: %v\n", nums, duplicates)

	// 分组操作
	people := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	groups := make(map[int][]string)

	for i, person := range people {
		groupID := i % 3 // 分成3组
		groups[groupID] = append(groups[groupID], person)
	}

	fmt.Println("人员分组:")
	for groupID, members := range groups {
		fmt.Printf("第%d组: %v\n", groupID, members)
	}

	// 数据去重
	original := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
	uniqueMap := make(map[string]bool)
	unique := []string{}

	for _, item := range original {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			unique = append(unique, item)
		}
	}

	fmt.Printf("原数据: %v\n", original)
	fmt.Printf("去重后: %v\n", unique)
}
