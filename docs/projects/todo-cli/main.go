package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
项目案例2：命令行TODO管理器
功能：
1. 添加任务
2. 列出任务
3. 完成任务
4. 删除任务
5. 保存到文件
*/

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TodoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

const dataFile = "todos.json"

func main() {
	fmt.Println("=== Go语言项目案例：TODO管理器 ===")

	todoList := loadTodos()

	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add", "a":
		if len(os.Args) < 3 {
			fmt.Println("错误：请提供任务描述")
			fmt.Println("用法：go run main.go add \"任务描述\"")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		todoList.addTask(description)

	case "list", "l":
		todoList.listTasks()

	case "complete", "c":
		if len(os.Args) < 3 {
			fmt.Println("错误：请提供任务ID")
			fmt.Println("用法：go run main.go complete <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("错误：无效的任务ID")
			return
		}
		todoList.completeTask(id)

	case "delete", "d":
		if len(os.Args) < 3 {
			fmt.Println("错误：请提供任务ID")
			fmt.Println("用法：go run main.go delete <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("错误：无效的任务ID")
			return
		}
		todoList.deleteTask(id)

	case "clear":
		todoList.clearCompleted()

	case "help", "h":
		showUsage()
		return

	default:
		fmt.Printf("错误：未知命令 '%s'\n", command)
		showUsage()
		return
	}

	saveTodos(todoList)
}

// 添加任务
func (tl *TodoList) addTask(description string) {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++

	fmt.Printf("✅ 已添加任务: \"%s\" (ID: %d)\n", description, task.ID)
}

// 列出所有任务
func (tl *TodoList) listTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("📝 暂无任务")
		return
	}

	fmt.Println("📋 任务列表:")
	fmt.Println("=" + strings.Repeat("=", 50))

	for _, task := range tl.Tasks {
		status := "⭕"
		if task.Completed {
			status = "✅"
		}

		timeInfo := task.CreatedAt.Format("2006-01-02 15:04")
		if task.Completed && task.CompletedAt != nil {
			timeInfo += " (完成: " + task.CompletedAt.Format("01-02 15:04") + ")"
		}

		fmt.Printf("%s [%d] %s\n", status, task.ID, task.Description)
		fmt.Printf("    创建时间: %s\n", timeInfo)
		fmt.Println()
	}

	completed := 0
	for _, task := range tl.Tasks {
		if task.Completed {
			completed++
		}
	}

	fmt.Printf("📊 统计: 总共 %d 个任务，已完成 %d 个，待完成 %d 个\n",
		len(tl.Tasks), completed, len(tl.Tasks)-completed)
}

// 完成任务
func (tl *TodoList) completeTask(id int) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			if tl.Tasks[i].Completed {
				fmt.Printf("ℹ️  任务 %d 已经完成了\n", id)
				return
			}

			now := time.Now()
			tl.Tasks[i].Completed = true
			tl.Tasks[i].CompletedAt = &now

			fmt.Printf("🎉 任务 %d 已完成: \"%s\"\n", id, tl.Tasks[i].Description)
			return
		}
	}

	fmt.Printf("❌ 未找到ID为 %d 的任务\n", id)
}

// 删除任务
func (tl *TodoList) deleteTask(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			// 确认删除
			fmt.Printf("确定要删除任务 \"%s\" 吗? (y/N): ", task.Description)

			var response string
			fmt.Scanln(&response)

			if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
				tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
				fmt.Printf("🗑️  已删除任务: \"%s\"\n", task.Description)
			} else {
				fmt.Println("取消删除")
			}
			return
		}
	}

	fmt.Printf("❌ 未找到ID为 %d 的任务\n", id)
}

// 清除已完成的任务
func (tl *TodoList) clearCompleted() {
	initialCount := len(tl.Tasks)

	var remaining []Task
	for _, task := range tl.Tasks {
		if !task.Completed {
			remaining = append(remaining, task)
		}
	}

	tl.Tasks = remaining

	clearedCount := initialCount - len(remaining)
	if clearedCount > 0 {
		fmt.Printf("🧹 已清除 %d 个已完成的任务\n", clearedCount)
	} else {
		fmt.Println("ℹ️  没有已完成的任务需要清除")
	}
}

// 加载TODO数据
func loadTodos() *TodoList {
	todoList := &TodoList{
		Tasks:  []Task{},
		NextID: 1,
	}

	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return todoList
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		fmt.Printf("警告：无法读取数据文件: %s\n", err)
		return todoList
	}

	err = json.Unmarshal(data, todoList)
	if err != nil {
		fmt.Printf("警告：数据文件格式错误: %s\n", err)
		return todoList
	}

	return todoList
}

// 保存TODO数据
func saveTodos(todoList *TodoList) {
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		fmt.Printf("错误：无法保存数据: %s\n", err)
		return
	}

	err = ioutil.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Printf("错误：无法写入文件: %s\n", err)
		return
	}
}

// 显示使用说明
func showUsage() {
	fmt.Println("用法：")
	fmt.Println("  go run main.go <命令> [参数]")
	fmt.Println()
	fmt.Println("命令：")
	fmt.Println("  add <描述>     添加新任务")
	fmt.Println("  list          列出所有任务")
	fmt.Println("  complete <ID> 完成指定任务")
	fmt.Println("  delete <ID>   删除指定任务")
	fmt.Println("  clear         清除所有已完成任务")
	fmt.Println("  help          显示此帮助信息")
	fmt.Println()
	fmt.Println("示例：")
	fmt.Println("  go run main.go add \"学习Go语言\"")
	fmt.Println("  go run main.go list")
	fmt.Println("  go run main.go complete 1")
	fmt.Println("  go run main.go delete 2")
}
