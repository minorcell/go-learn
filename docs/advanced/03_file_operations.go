package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
03_file_operations.go - Go标准库：文件操作
涉及包：
- os: 操作系统接口
- io: I/O操作
- bufio: 缓冲I/O
- path/filepath: 文件路径操作
- encoding/json: JSON处理
- encoding/csv: CSV处理
*/

// 示例数据结构
type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

func main() {
	fmt.Println("=== Go标准库：文件操作 ===")

	// 1. 基本文件操作
	fmt.Println("\n1. 基本文件操作：")

	// 创建目录
	testDir := "test_files"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}
	fmt.Printf("创建目录: %s\n", testDir)

	// 创建文件
	filename := filepath.Join(testDir, "example.txt")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 写入内容
	content := "Hello, Go!\n这是一个测试文件。\n"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Printf("创建并写入文件: %s\n", filename)

	// 2. 读取文件的多种方式
	fmt.Println("\n2. 读取文件的多种方式：")

	// 方式1: os.ReadFile (一次性读取整个文件)
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
	} else {
		fmt.Printf("方式1 - os.ReadFile:\n%s", data)
	}

	// 方式2: 使用bufio逐行读取
	fmt.Println("方式2 - bufio逐行读取:")
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			fmt.Printf("  第%d行: %s\n", lineNum, scanner.Text())
			lineNum++
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("读取文件错误: %v\n", err)
		}
	}

	// 方式3: 使用io包
	fmt.Println("方式3 - io.ReadAll:")
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
	} else {
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("读取文件失败: %v\n", err)
		} else {
			fmt.Printf("内容长度: %d 字节\n", len(data))
		}
	}

	// 3. 文件信息和属性
	fmt.Println("\n3. 文件信息和属性：")

	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
	} else {
		fmt.Printf("文件名: %s\n", fileInfo.Name())
		fmt.Printf("文件大小: %d 字节\n", fileInfo.Size())
		fmt.Printf("修改时间: %s\n", fileInfo.ModTime().Format("2006-01-02 15:04:05"))
		fmt.Printf("是否为目录: %t\n", fileInfo.IsDir())
		fmt.Printf("文件权限: %s\n", fileInfo.Mode())
	}

	// 4. 路径操作
	fmt.Println("\n4. 路径操作：")

	fullPath, _ := filepath.Abs(filename)
	fmt.Printf("绝对路径: %s\n", fullPath)
	fmt.Printf("目录: %s\n", filepath.Dir(fullPath))
	fmt.Printf("文件名: %s\n", filepath.Base(fullPath))
	fmt.Printf("扩展名: %s\n", filepath.Ext(fullPath))

	// 路径拼接
	newPath := filepath.Join("data", "files", "test.txt")
	fmt.Printf("路径拼接: %s\n", newPath)

	// 5. JSON文件操作
	fmt.Println("\n5. JSON文件操作：")

	// 创建示例数据
	people := []Person{
		{Name: "张三", Age: 25, Email: "zhangsan@example.com", CreateAt: time.Now()},
		{Name: "李四", Age: 30, Email: "lisi@example.com", CreateAt: time.Now()},
		{Name: "王五", Age: 28, Email: "wangwu@example.com", CreateAt: time.Now()},
	}

	// 写入JSON文件
	jsonFile := filepath.Join(testDir, "people.json")
	file, err = os.Create(jsonFile)
	if err != nil {
		fmt.Printf("创建JSON文件失败: %v\n", err)
	} else {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ") // 格式化输出
		err = encoder.Encode(people)
		if err != nil {
			fmt.Printf("编码JSON失败: %v\n", err)
		} else {
			fmt.Printf("JSON数据写入: %s\n", jsonFile)
		}
	}

	// 读取JSON文件
	var loadedPeople []Person
	file, err = os.Open(jsonFile)
	if err != nil {
		fmt.Printf("打开JSON文件失败: %v\n", err)
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&loadedPeople)
		if err != nil {
			fmt.Printf("解码JSON失败: %v\n", err)
		} else {
			fmt.Printf("从JSON读取到 %d 条记录:\n", len(loadedPeople))
			for i, person := range loadedPeople {
				fmt.Printf("  %d. %s (%d岁) - %s\n", i+1, person.Name, person.Age, person.Email)
			}
		}
	}

	// 6. CSV文件操作
	fmt.Println("\n6. CSV文件操作：")

	// 写入CSV文件
	csvFile := filepath.Join(testDir, "people.csv")
	file, err = os.Create(csvFile)
	if err != nil {
		fmt.Printf("创建CSV文件失败: %v\n", err)
	} else {
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// 写入标题行
		headers := []string{"姓名", "年龄", "邮箱", "创建时间"}
		writer.Write(headers)

		// 写入数据行
		for _, person := range people {
			record := []string{
				person.Name,
				fmt.Sprintf("%d", person.Age),
				person.Email,
				person.CreateAt.Format("2006-01-02 15:04:05"),
			}
			writer.Write(record)
		}
		fmt.Printf("CSV数据写入: %s\n", csvFile)
	}

	// 读取CSV文件
	file, err = os.Open(csvFile)
	if err != nil {
		fmt.Printf("打开CSV文件失败: %v\n", err)
	} else {
		defer file.Close()
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("读取CSV失败: %v\n", err)
		} else {
			fmt.Printf("从CSV读取到 %d 行数据:\n", len(records))
			for i, record := range records {
				fmt.Printf("  第%d行: %v\n", i+1, record)
			}
		}
	}

	// 7. 目录遍历
	fmt.Println("\n7. 目录遍历：")

	fmt.Printf("遍历目录: %s\n", testDir)
	err = filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Printf("  📁 %s/\n", path)
		} else {
			fmt.Printf("  📄 %s (%d bytes)\n", path, info.Size())
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}

	// 8. 文件复制
	fmt.Println("\n8. 文件复制：")

	srcFile := jsonFile
	dstFile := filepath.Join(testDir, "people_backup.json")

	err = copyFile(srcFile, dstFile)
	if err != nil {
		fmt.Printf("文件复制失败: %v\n", err)
	} else {
		fmt.Printf("文件复制成功: %s -> %s\n", srcFile, dstFile)
	}

	// 9. 文件搜索
	fmt.Println("\n9. 文件搜索：")

	// 搜索特定扩展名的文件
	pattern := filepath.Join(testDir, "*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("搜索文件失败: %v\n", err)
	} else {
		fmt.Printf("找到JSON文件:\n")
		for _, match := range matches {
			fmt.Printf("  - %s\n", match)
		}
	}

	// 10. 临时文件
	fmt.Println("\n10. 临时文件：")

	tmpFile, err := os.CreateTemp(testDir, "temp_*.txt")
	if err != nil {
		fmt.Printf("创建临时文件失败: %v\n", err)
	} else {
		defer os.Remove(tmpFile.Name()) // 清理临时文件
		defer tmpFile.Close()

		fmt.Printf("创建临时文件: %s\n", tmpFile.Name())
		tmpFile.WriteString("这是临时文件内容")
	}

	// 11. 文件锁定和原子操作
	fmt.Println("\n11. 原子写入操作：")

	atomicFile := filepath.Join(testDir, "atomic.txt")
	err = atomicWriteFile(atomicFile, "原子写入的内容\n安全可靠")
	if err != nil {
		fmt.Printf("原子写入失败: %v\n", err)
	} else {
		fmt.Printf("原子写入成功: %s\n", atomicFile)
	}

	// 12. 文件监控信息
	fmt.Println("\n12. 文件系统信息：")

	entries, err := os.ReadDir(testDir)
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
	} else {
		fmt.Printf("目录 %s 包含 %d 个条目:\n", testDir, len(entries))
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("  📁 %s/\n", entry.Name())
			} else {
				info, _ := entry.Info()
				fmt.Printf("  📄 %s (%d bytes, %s)\n",
					entry.Name(), info.Size(), info.ModTime().Format("01-02 15:04"))
			}
		}
	}

	fmt.Println("\n文件操作演示完成！")

	// 可选：清理测试文件
	fmt.Print("是否清理测试文件? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		os.RemoveAll(testDir)
		fmt.Println("测试文件已清理")
	}
}

// 文件复制函数
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// 原子写入文件
func atomicWriteFile(filename, content string) error {
	// 创建临时文件
	tmpFile, err := os.CreateTemp(filepath.Dir(filename), ".tmp_"+filepath.Base(filename))
	if err != nil {
		return err
	}

	// 写入内容
	_, err = tmpFile.WriteString(content)
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return err
	}

	// 关闭临时文件
	err = tmpFile.Close()
	if err != nil {
		os.Remove(tmpFile.Name())
		return err
	}

	// 原子性地重命名文件
	return os.Rename(tmpFile.Name(), filename)
}
