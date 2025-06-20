# Go语言学习项目

## 🎯 项目简介

这是一个全面的Go语言学习项目，按照从基础到进阶再到实战项目的三层结构组织，系统地掌握Go语言编程。

## 📁 项目结构

```
goang/
├── basics/                 # 第一部分：基础语法 (7个模块)
│   ├── 01_variables_types.go   # 变量和数据类型
│   ├── 02_control_flow.go      # 控制流程
│   ├── 03_functions.go         # 函数编程
│   ├── 04_arrays_slices_maps.go # 数组、切片和映射
│   ├── 05_structs_methods.go   # 结构体和方法
│   ├── 06_interfaces.go        # 接口
│   └── 07_error_handling.go    # 错误处理
├── advanced/               # 第二部分：Go标准库包 (6个模块)
│   ├── 01_packages.go          # 包管理基础
│   ├── 02_concurrency.go       # 并发编程
│   ├── 03_file_operations.go   # 文件操作 ✨NEW
│   ├── 04_network_http.go      # 网络和HTTP ✨NEW
│   ├── 05_strings_regexp.go    # 字符串和正则 ✨NEW
│   └── 06_time_crypto.go       # 时间和加密 ✨NEW
├── projects/               # 第三部分：实战项目
│   ├── calculator/             # 命令行计算器
│   │   └── main.go
│   ├── todo-cli/              # TODO管理器
│   │   └── main.go
│   └── web-server/            # Web服务器 ✨NEW
│       └── main.go
├── README.md              # 项目说明
```

## 🚀 解决方案：main函数冲突问题

### 问题说明
之前的版本中，多个文件都声明了`main`函数，导致在同一个包中出现冲突。

### 解决方案
1. **目录分离**：将不同学习模块放在不同目录中
2. **独立运行**：每个文件都可以独立编译和运行
3. **清晰结构**：按学习难度和内容类型分类

## 📚 学习路径

### 第一部分：基础语法 (basics/) - 7个模块

#### 1. 变量和数据类型
```bash
cd basics
go run 01_variables_types.go
```
**学习内容：**
- 变量声明的多种方式
- 基本数据类型（整数、浮点、布尔、字符串）
- 常量定义
- 类型转换
- 指针基础

#### 2. 控制流程
```bash
cd basics
go run 02_control_flow.go
```
**学习内容：**
- if/else 条件语句
- for 循环（基本、while风格、无限循环）
- switch 选择语句
- defer 延迟执行
- range 循环

#### 3. 函数编程
```bash
cd basics
go run 03_functions.go
```
**学习内容：**
- 函数定义和调用
- 参数和返回值（多返回值、命名返回值）
- 可变参数
- 函数作为值
- 闭包和匿名函数
- 递归
- 高阶函数

#### 4. 数组、切片和映射 ✨NEW
```bash
cd basics
go run 04_arrays_slices_maps.go
```
**学习内容：**
- 数组的声明和使用
- 切片的创建和操作
- 映射的基本用法
- 增删改查操作
- 排序和搜索
- 嵌套数据结构
- 实用示例和技巧

#### 5. 结构体和方法 ✨NEW
```bash
cd basics
go run 05_structs_methods.go
```
**学习内容：**
- 结构体定义和实例化
- 方法定义（值接收者vs指针接收者）
- 嵌套结构体
- 构造函数模式
- 结构体组合
- 实际应用示例

#### 6. 接口 ✨NEW
```bash
cd basics
go run 06_interfaces.go
```
**学习内容：**
- 接口定义和实现
- 接口组合
- 空接口 interface{}
- 类型断言和类型选择
- 多态性
- 接口的最佳实践

#### 7. 错误处理 ✨NEW
```bash
cd basics
go run 07_error_handling.go
```
**学习内容：**
- error接口和错误类型
- 错误创建和返回
- 自定义错误类型
- 错误包装和解包
- panic和recover
- 错误处理最佳实践

### 第二部分：Go标准库包 (advanced/) - 6个模块

#### 1. 包管理基础
```bash
cd advanced
go run 01_packages.go
```
**学习内容：**
- 包的导入和使用
- 标准库常用包（fmt, math, strings, time, os）
- 包的可见性规则
- init函数
- 自定义类型和方法

#### 2. 并发编程
```bash
cd advanced
go run 02_concurrency.go
```
**学习内容：**
- Goroutines 协程
- Channels 通道（缓冲和非缓冲）
- Select 语句
- sync包（WaitGroup, Mutex, RWMutex）
- 并发模式（生产者-消费者、管道、超时处理）

#### 3. 文件操作 ✨NEW
```bash
cd advanced
go run 03_file_operations.go
```
**涉及标准库：**
- `os`: 操作系统接口
- `io`: I/O操作
- `bufio`: 缓冲I/O
- `path/filepath`: 文件路径操作
- `encoding/json`: JSON处理
- `encoding/csv`: CSV处理

**学习内容：**
- 文件创建、读取、写入
- 目录操作和遍历
- JSON和CSV文件处理
- 文件复制和移动
- 临时文件和原子操作
- 文件权限和属性

#### 4. 网络和HTTP ✨NEW
```bash
cd advanced
go run 04_network_http.go
```
**涉及标准库：**
- `net`: 网络编程基础
- `net/http`: HTTP客户端和服务器
- `net/url`: URL解析
- `context`: 上下文管理

**学习内容：**
- HTTP客户端请求（GET, POST, PUT, DELETE）
- 请求头和响应处理
- Cookie和Session
- 超时和上下文控制
- TCP连接和网络信息
- 文件下载和上传

#### 5. 字符串和正则表达式 ✨NEW
```bash
cd advanced
go run 05_strings_regexp.go
```
**涉及标准库：**
- `strings`: 字符串操作
- `strconv`: 字符串转换
- `regexp`: 正则表达式
- `unicode`: Unicode处理
- `unicode/utf8`: UTF-8编码处理

**学习内容：**
- 字符串分割、连接、替换
- 类型转换（字符串↔数字↔布尔）
- 正则表达式匹配和查找
- Unicode和中文处理
- 文本清理和验证
- 字符串格式化

#### 6. 时间和基础加密 ✨NEW
```bash
cd advanced
go run 06_time_crypto.go
```
**涉及标准库：**
- `time`: 时间和日期处理
- `crypto/md5`: MD5哈希
- `crypto/sha256`: SHA256哈希
- `crypto/rand`: 安全随机数
- `encoding/base64`: Base64编码
- `encoding/hex`: 十六进制编码

**学习内容：**
- 时间格式化和解析
- 时区处理和时间计算
- 定时器和延时
- MD5和SHA256哈希
- Base64编码解码
- 安全随机数生成
- 密码哈希示例

### 第三部分：实战项目 (projects/)

#### 1. 命令行计算器
```bash
cd projects/calculator
go run main.go
```
**功能特性：**
- 基本四则运算
- 历史记录
- 交互式界面
- 帮助系统

#### 2. TODO管理器
```bash
cd projects/todo-cli
go run main.go help
```
**功能特性：**
- 添加/删除任务
- 完成任务
- 任务列表显示
- JSON数据持久化

**使用示例：**
```bash
# 添加任务
go run main.go add "学习Go语言"

# 查看任务列表
go run main.go list

# 完成任务
go run main.go complete 1

# 删除任务
go run main.go delete 2
```

#### 3. Web服务器 ✨NEW
```bash
cd projects/web-server
go run main.go
```
**功能特性：**
- RESTful API
- JSON响应
- HTML模板
- 中间件（日志、CORS）
- 用户管理系统
- 实时状态监控

**API端点：**
- `GET /` - 首页文档
- `GET /api/users` - 获取所有用户
- `GET /api/users/{id}` - 获取指定用户
- `POST /api/users` - 创建新用户
- `PUT /api/users/{id}` - 更新用户
- `DELETE /api/users/{id}` - 删除用户
- `GET /api/status` - 服务器状态
- `GET /time` - 当前时间

**测试示例：**
```bash
# 访问首页
http://localhost:8080

# 获取所有用户
curl http://localhost:8080/api/users

# 创建新用户
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","email":"new@example.com","age":25}'

# 获取服务器状态
curl http://localhost:8080/api/status
```

## ⚡ 快速开始

### 1. 运行基础教程（完整版）
```bash
# 学习变量和类型
cd basics && go run 01_variables_types.go

# 学习控制流程
cd basics && go run 02_control_flow.go

# 学习函数
cd basics && go run 03_functions.go

# 学习数据结构
cd basics && go run 04_arrays_slices_maps.go

# 学习结构体和方法
cd basics && go run 05_structs_methods.go

# 学习接口
cd basics && go run 06_interfaces.go

# 学习错误处理
cd basics && go run 07_error_handling.go
```

### 2. 运行标准库教程
```bash
# 学习包管理
cd advanced && go run 01_packages.go

# 学习并发编程
cd advanced && go run 02_concurrency.go

# 学习文件操作
cd advanced && go run 03_file_operations.go

# 学习网络和HTTP
cd advanced && go run 04_network_http.go

# 学习字符串和正则
cd advanced && go run 05_strings_regexp.go

# 学习时间和加密
cd advanced && go run 06_time_crypto.go
```

### 3. 体验实战项目
```bash
# 体验计算器
cd projects/calculator && go run main.go

# 体验TODO管理器
cd projects/todo-cli && go run main.go list

# 体验Web服务器
cd projects/web-server && go run main.go
```

## 🔧 开发环境要求

- Go 1.16 或更高版本
- 文本编辑器或IDE（推荐VS Code + Go扩展）
- 终端或命令行工具

## 💡 学习建议

1. **按顺序学习**：建议按照基础 → 标准库 → 实战的顺序学习
2. **实践为主**：每个示例都要亲自运行和修改
3. **理解原理**：不仅要知道怎么用，还要理解为什么这样设计
4. **编写代码**：尝试修改示例代码，编写自己的变体
5. **查阅文档**：学会使用Go官方文档和标准库文档

## 🎯 学习目标

### 基础阶段（7个模块）
- [ ] 掌握Go语言基本语法
- [ ] 理解变量、常量、数据类型
- [ ] 熟练使用控制结构
- [ ] 掌握函数编程
- [ ] 掌握数据结构（数组、切片、映射）
- [ ] 理解结构体和方法
- [ ] 掌握接口和多态性
- [ ] 熟练处理错误

### 标准库阶段（6个模块）
- [ ] 理解包的概念和使用
- [ ] 掌握并发编程基础
- [ ] 熟练进行文件操作
- [ ] 掌握网络编程和HTTP
- [ ] 熟练处理字符串和正则表达式
- [ ] 掌握时间处理和基础加密

### 实战阶段
- [ ] 能够编写命令行工具
- [ ] 能够构建Web服务器
- [ ] 理解项目结构组织
- [ ] 掌握错误处理
- [ ] 能够编写实用程序

## 📈 项目特色

### 🔥 新增功能
- ✅ **基础语法扩充**：从3个模块增加到7个模块
- ✅ **标准库详解**：新增4个标准库使用模块
- ✅ **Web服务器项目**：完整的RESTful API示例
- ✅ **更全面的覆盖**：涵盖Go语言核心概念和常用库

### 💪 核心优势
- ✅ **解决冲突**：每个模块独立运行
- ✅ **结构清晰**：三层递进式学习
- ✅ **实用项目**：真实应用场景
- ✅ **中文注释**：详细的中文说明
- ✅ **即学即用**：理论结合实践
- ✅ **标准库全面**：涵盖常用标准库包

## 📖 参考资源

- [Go官方文档](https://golang.org/doc/)
- [Go语言之旅](https://tour.golang.org/)
- [Go标准库文档](https://pkg.go.dev/std)
- [Effective Go](https://golang.org/doc/effective_go.html)

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个学习项目！

---

**祝你学习愉快！** 🎉 