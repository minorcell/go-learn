package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/*
04_network_http.go - Go标准库：网络和HTTP
涉及包：
- net: 网络编程基础
- net/http: HTTP客户端和服务器
- net/url: URL解析
- context: 上下文管理
*/

// 响应数据结构
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error,omitempty"`
}

func main() {
	fmt.Println("=== Go标准库：网络和HTTP ===")

	// 1. HTTP客户端基础
	fmt.Println("\n1. HTTP客户端基础：")

	// 简单GET请求
	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("GET请求失败: %v\n", err)
	} else {
		defer response.Body.Close()
		fmt.Printf("状态码: %d\n", response.StatusCode)
		fmt.Printf("Content-Type: %s\n", response.Header.Get("Content-Type"))

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("读取响应体失败: %v\n", err)
		} else {
			fmt.Printf("响应体长度: %d 字节\n", len(body))
		}
	}

	// 2. 带超时的HTTP客户端
	fmt.Println("\n2. 带超时的HTTP客户端：")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err = client.Get("https://httpbin.org/delay/2")
	if err != nil {
		fmt.Printf("超时请求失败: %v\n", err)
	} else {
		defer response.Body.Close()
		fmt.Printf("延迟请求成功，状态码: %d\n", response.StatusCode)
	}

	// 3. POST请求和JSON数据
	fmt.Println("\n3. POST请求和JSON数据：")

	// 准备POST数据
	postData := map[string]interface{}{
		"name":  "Go学习者",
		"email": "learner@example.com",
		"age":   25,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Printf("JSON编码失败: %v\n", err)
	} else {
		// 发送POST请求
		response, err = http.Post(
			"https://httpbin.org/post",
			"application/json",
			strings.NewReader(string(jsonData)),
		)
		if err != nil {
			fmt.Printf("POST请求失败: %v\n", err)
		} else {
			defer response.Body.Close()
			fmt.Printf("POST请求成功，状态码: %d\n", response.StatusCode)

			// 解析响应
			var result map[string]interface{}
			if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
				fmt.Printf("解析响应失败: %v\n", err)
			} else {
				if data, ok := result["json"].(map[string]interface{}); ok {
					fmt.Printf("服务器收到的数据: %+v\n", data)
				}
			}
		}
	}

	// 4. 自定义请求头
	fmt.Println("\n4. 自定义请求头：")

	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
	} else {
		// 设置请求头
		req.Header.Set("User-Agent", "Go-Learning-Client/1.0")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Custom-Header", "学习Go语言")

		response, err = client.Do(req)
		if err != nil {
			fmt.Printf("发送请求失败: %v\n", err)
		} else {
			defer response.Body.Close()
			fmt.Printf("自定义请求头成功，状态码: %d\n", response.StatusCode)

			var result map[string]interface{}
			if err := json.NewDecoder(response.Body).Decode(&result); err == nil {
				if headers, ok := result["headers"].(map[string]interface{}); ok {
					fmt.Printf("请求头: %+v\n", headers)
				}
			}
		}
	}

	// 5. URL参数处理
	fmt.Println("\n5. URL参数处理：")

	baseURL := "https://httpbin.org/get"
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("URL解析失败: %v\n", err)
	} else {
		// 添加查询参数
		q := u.Query()
		q.Set("name", "张三")
		q.Set("age", "25")
		q.Add("hobby", "编程")
		q.Add("hobby", "读书")
		u.RawQuery = q.Encode()

		fmt.Printf("构建的URL: %s\n", u.String())

		response, err = client.Get(u.String())
		if err != nil {
			fmt.Printf("参数请求失败: %v\n", err)
		} else {
			defer response.Body.Close()
			fmt.Printf("参数请求成功，状态码: %d\n", response.StatusCode)
		}
	}

	// 6. 表单数据提交
	fmt.Println("\n6. 表单数据提交：")

	formData := url.Values{}
	formData.Set("username", "testuser")
	formData.Set("password", "testpass")
	formData.Set("remember", "true")

	response, err = http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		fmt.Printf("表单提交失败: %v\n", err)
	} else {
		defer response.Body.Close()
		fmt.Printf("表单提交成功，状态码: %d\n", response.StatusCode)
	}

	// 7. 使用Context控制请求
	fmt.Println("\n7. 使用Context控制请求：")

	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err = http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/5", nil)
	if err != nil {
		fmt.Printf("创建Context请求失败: %v\n", err)
	} else {
		response, err = client.Do(req)
		if err != nil {
			fmt.Printf("Context请求失败 (预期超时): %v\n", err)
		} else {
			defer response.Body.Close()
			fmt.Printf("Context请求成功，状态码: %d\n", response.StatusCode)
		}
	}

	// 8. Cookie处理
	fmt.Println("\n8. Cookie处理：")

	// 创建带cookie jar的客户端
	jar := &http.CookieJar{}
	clientWithCookies := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	// 设置cookie
	response, err = clientWithCookies.Get("https://httpbin.org/cookies/set/session/abc123")
	if err != nil {
		fmt.Printf("设置Cookie失败: %v\n", err)
	} else {
		response.Body.Close()
		fmt.Printf("Cookie设置成功，状态码: %d\n", response.StatusCode)
	}

	// 获取cookie
	response, err = clientWithCookies.Get("https://httpbin.org/cookies")
	if err != nil {
		fmt.Printf("获取Cookie失败: %v\n", err)
	} else {
		defer response.Body.Close()
		fmt.Printf("Cookie获取成功，状态码: %d\n", response.StatusCode)

		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err == nil {
			fmt.Printf("服务器看到的Cookie: %+v\n", result["cookies"])
		}
	}

	// 9. 网络连接基础
	fmt.Println("\n9. 网络连接基础：")

	// TCP连接示例
	conn, err := net.DialTimeout("tcp", "www.google.com:80", 5*time.Second)
	if err != nil {
		fmt.Printf("TCP连接失败: %v\n", err)
	} else {
		defer conn.Close()
		fmt.Printf("TCP连接成功: %s -> %s\n",
			conn.LocalAddr(), conn.RemoteAddr())

		// 发送HTTP请求
		request := "GET / HTTP/1.1\r\nHost: www.google.com\r\nConnection: close\r\n\r\n"
		_, err = conn.Write([]byte(request))
		if err != nil {
			fmt.Printf("发送数据失败: %v\n", err)
		} else {
			// 读取响应头
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Printf("读取响应失败: %v\n", err)
			} else {
				response := string(buffer[:n])
				lines := strings.Split(response, "\r\n")
				if len(lines) > 0 {
					fmt.Printf("HTTP响应状态: %s\n", lines[0])
				}
			}
		}
	}

	// 10. 本地网络信息
	fmt.Println("\n10. 本地网络信息：")

	// 获取本地IP地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("获取网络接口失败: %v\n", err)
	} else {
		fmt.Println("网络接口:")
		for _, iface := range interfaces {
			if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
				addrs, err := iface.Addrs()
				if err != nil {
					continue
				}
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							fmt.Printf("  %s: %s\n", iface.Name, ipnet.IP.String())
						}
					}
				}
			}
		}
	}

	// 11. DNS解析
	fmt.Println("\n11. DNS解析：")

	// 解析域名
	ips, err := net.LookupIP("www.google.com")
	if err != nil {
		fmt.Printf("DNS解析失败: %v\n", err)
	} else {
		fmt.Println("www.google.com 的IP地址:")
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	// 反向DNS解析
	names, err := net.LookupAddr("8.8.8.8")
	if err != nil {
		fmt.Printf("反向DNS解析失败: %v\n", err)
	} else {
		fmt.Printf("8.8.8.8 对应的域名: %v\n", names)
	}

	// 12. HTTP状态码处理
	fmt.Println("\n12. HTTP状态码处理：")

	testURLs := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
	}

	for _, testURL := range testURLs {
		response, err := client.Get(testURL)
		if err != nil {
			fmt.Printf("请求 %s 失败: %v\n", testURL, err)
			continue
		}
		response.Body.Close()

		switch response.StatusCode {
		case 200:
			fmt.Printf("✅ %s - 成功\n", testURL)
		case 404:
			fmt.Printf("❌ %s - 页面未找到\n", testURL)
		case 500:
			fmt.Printf("💥 %s - 服务器错误\n", testURL)
		default:
			fmt.Printf("❓ %s - 状态码: %d\n", testURL, response.StatusCode)
		}
	}

	// 13. 下载文件示例
	fmt.Println("\n13. 文件下载示例：")

	err = downloadFile("https://httpbin.org/json", "downloaded.json")
	if err != nil {
		fmt.Printf("文件下载失败: %v\n", err)
	} else {
		fmt.Println("文件下载成功: downloaded.json")
	}

	fmt.Println("\n网络和HTTP操作演示完成！")
}

// 下载文件函数
func downloadFile(url, filename string) error {
	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 发送请求
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 检查状态码
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，状态码: %d", response.StatusCode)
	}

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 复制响应体到文件
	_, err = io.Copy(file, response.Body)
	return err
}
