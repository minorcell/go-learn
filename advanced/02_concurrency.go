package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
02_concurrency.go - Go语言进阶：并发编程
学习内容：
1. Goroutines 协程
2. Channels 通道
3. Select 语句
4. sync包：WaitGroup, Mutex, RWMutex
5. 并发模式和最佳实践
*/

func main() {
	fmt.Println("=== Go语言进阶：并发编程 ===")

	// 1. 基本的Goroutine
	fmt.Println("\n1. 基本的Goroutine：")

	// 启动一个goroutine
	go sayHello("World")
	go sayHello("Go")
	go sayHello("并发")

	// 主goroutine稍等一下，让其他goroutine执行
	time.Sleep(100 * time.Millisecond)

	// 2. 使用WaitGroup等待goroutine完成
	fmt.Println("\n2. 使用WaitGroup：")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待的goroutine数量
		go func(id int) {
			defer wg.Done() // 完成时调用Done
			fmt.Printf("Goroutine %d 开始执行\n", id)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Printf("Goroutine %d 执行完成\n", id)
		}(i)
	}

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有goroutine已完成")

	// 3. 通道基础
	fmt.Println("\n3. 通道基础：")

	// 创建一个字符串通道
	ch := make(chan string)

	// 在goroutine中发送数据
	go func() {
		ch <- "Hello"
		ch <- "Channel"
		ch <- "Communication"
		close(ch) // 关闭通道
	}()

	// 接收数据
	for msg := range ch {
		fmt.Printf("接收到: %s\n", msg)
	}

	// 4. 带缓冲的通道
	fmt.Println("\n4. 带缓冲的通道：")

	bufferedCh := make(chan int, 3) // 缓冲区大小为3

	// 发送数据（不会阻塞，因为有缓冲区）
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	fmt.Printf("通道长度: %d, 容量: %d\n", len(bufferedCh), cap(bufferedCh))

	// 接收数据
	for i := 0; i < 3; i++ {
		value := <-bufferedCh
		fmt.Printf("接收到: %d\n", value)
	}

	// 5. Select语句
	fmt.Println("\n5. Select语句：")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "来自通道1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "来自通道2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("通道1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("通道2: %s\n", msg2)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("超时")
		}
	}

	// 6. Mutex互斥锁
	fmt.Println("\n6. Mutex互斥锁：")

	var counter int
	var mutex sync.Mutex
	var wg2 sync.WaitGroup

	// 启动多个goroutine同时修改counter
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	wg2.Wait()
	fmt.Printf("最终计数: %d\n", counter)

	// 7. 读写锁RWMutex
	fmt.Println("\n7. 读写锁RWMutex：")

	var data = map[string]int{"a": 1, "b": 2, "c": 3}
	var rwMutex sync.RWMutex
	var wg3 sync.WaitGroup

	// 读操作（可以并发）
	for i := 0; i < 3; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			rwMutex.RLock()
			fmt.Printf("读取者 %d: %v\n", id, data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.RUnlock()
		}(i)
	}

	// 写操作（独占）
	wg3.Add(1)
	go func() {
		defer wg3.Done()
		time.Sleep(20 * time.Millisecond)
		rwMutex.Lock()
		data["d"] = 4
		fmt.Println("写入者: 添加了 d=4")
		rwMutex.Unlock()
	}()

	wg3.Wait()

	// 8. 生产者-消费者模式
	fmt.Println("\n8. 生产者-消费者模式：")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("结果: %d\n", result)
	}

	// 9. 管道模式
	fmt.Println("\n9. 管道模式：")

	// 创建管道：数字生成器 -> 平方计算器 -> 结果收集器
	numbers := generateNumbers(1, 5)
	squares := calculateSquares(numbers)

	fmt.Println("平方结果:")
	for square := range squares {
		fmt.Printf("%d ", square)
	}
	fmt.Println()

	// 10. 超时处理
	fmt.Println("\n10. 超时处理：")

	slowCh := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		slowCh <- "慢操作完成"
	}()

	select {
	case result := <-slowCh:
		fmt.Printf("收到结果: %s\n", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("操作超时")
	}

	// 11. 上下文取消
	fmt.Println("\n11. 取消操作示例：")

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("任务被取消")
				return
			default:
				fmt.Println("工作中...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(150 * time.Millisecond)
	done <- true
	time.Sleep(10 * time.Millisecond)
}

// 简单的goroutine函数
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello %s! (%d)\n", name, i+1)
		time.Sleep(10 * time.Millisecond)
	}
}

// Worker函数用于生产者-消费者模式
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d 开始任务 %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		results <- job * 2
		fmt.Printf("Worker %d 完成任务 %d\n", id, job)
	}
}

// 数字生成器
func generateNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

// 平方计算器
func calculateSquares(numbers <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range numbers {
			ch <- num * num
		}
	}()
	return ch
}
