# 搭建Go语言开发环境

就像装修房子需要准备工具一样，写Go代码也需要先搭建好开发环境。别担心，整个过程很简单，跟着做就行！

## 第一步：检查你的电脑

Go语言支持几乎所有主流系统：
- **Windows** 7及以上版本
- **macOS** 10.11及以上版本
- **Linux** 各种发行版

```bash
# 检查你的系统信息
# Windows: 按Win+R，输入winver
# macOS: 点击苹果菜单 > 关于本机
# Linux: 运行 lsb_release -a
```

## 第二步：下载并安装Go

### Windows用户

**方法：下载安装包（推荐）**
1. 打开 [Go官网下载页](https://golang.org/dl/)
2. 下载 `go1.22.x.windows-amd64.msi` 文件
3. 双击安装包，一路点"下一步"就行
4. 安装完成！

**安装示意**：
```
下载安装包 → 双击运行 → 同意协议 → 选择路径 → 安装完成
     ↓            ↓         ↓         ↓         ↓
  几十MB文件    启动向导    点同意    默认即可   可以用了
```

### macOS用户

**方法一：官方安装包（推荐新手）**
1. 下载 `go1.22.x.darwin-amd64.pkg`
2. 双击安装包，按提示操作
3. 输入系统密码确认安装

**方法二：Homebrew（推荐有经验用户）**
```bash
# 如果没装Homebrew，先装它：
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# 然后安装Go：
brew install go
```

### Linux用户

**方法一：官方二进制包（推荐）**
```bash
# 1. 下载Go安装包
wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz

# 2. 解压到系统目录（需要管理员权限）
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# 3. 添加到环境变量
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**方法二：包管理器（更简单）**
```bash
# Ubuntu/Debian系统
sudo apt update
sudo apt install golang-go

# CentOS/RHEL系统  
sudo yum install golang

# Arch Linux
sudo pacman -S go
```

## 第三步：验证安装

打开终端（命令行），输入：
```bash
go version
```

如果看到类似这样的输出，说明安装成功了：
```
go version go1.22.0 linux/amd64
```

**如果出现"command not found"**：
```
问题：找不到go命令
原因：环境变量没设置好
解决：重新打开终端，或者重启电脑
```

## 第四步：配置开发环境

### 理解几个重要概念

**GOROOT** - Go语言的"家"
```
就像：
- Windows程序装在 C:\Program Files\
- Go语言装在某个文件夹里，这就是GOROOT
- 一般不需要你手动设置
```

**GOPATH** - 你的"工作空间"（旧方式，了解即可）
```
以前：所有Go项目都必须放在GOPATH目录下
现在：有了Go Modules，可以在任何地方创建项目
```

**Go Modules** - 现代项目管理方式（重要！）
```
就像：
- 每个项目都有自己的"小仓库"
- 自动管理依赖包，不会冲突
- 这是现在的标准做法
```

### 设置国内代理（中国用户必做）

由于网络原因，国内下载Go包可能很慢，设置代理能解决这个问题：

```bash
# 设置代理服务器
go env -w GOPROXY=https://goproxy.cn,direct

# 设置校验数据库
go env -w GOSUMDB=sum.golang.google.cn
```

**代理的作用**：
```
没有代理：你的电脑 → (龟速/失败) → 国外服务器
设置代理：你的电脑 → (飞速) → 国内镜像 → 国外服务器
```

### 启用模块支持

```bash
# 确保启用Go Modules（新版本默认开启）
go env -w GO111MODULE=on
```

## 第五步：选择开发工具

### VS Code（强烈推荐）

**为什么推荐VS Code？**
- 免费开源，微软出品
- 插件丰富，Go支持很好
- 轻量级，启动快
- 跨平台，Windows/Mac/Linux都能用

**安装步骤**：
1. 访问 [VS Code官网](https://code.visualstudio.com/)
2. 下载并安装（一路下一步）
3. 安装Go扩展：
   - 打开VS Code
   - 按 `Ctrl+Shift+X`（Mac用`Cmd+Shift+X`）
   - 搜索"Go"
   - 安装Google出的Go扩展

**配置Go工具**：
1. 按 `Ctrl+Shift+P`（Mac用`Cmd+Shift+P`）
2. 输入"Go: Install/Update Tools"
3. 全选，点击"OK"
4. 等待安装完成

### 其他选择

**GoLand** - JetBrains出品的专业IDE
```
优点：功能强大，智能提示好，调试方便
缺点：收费软件（学生可免费）
适合：有经验的开发者，大型项目
```

**Vim/Neovim** - 命令行编辑器
```
优点：轻量，效率高，装逼利器
缺点：学习曲线陡峭
适合：Linux高手，命令行爱好者
```

## 第六步：创建第一个项目

### 创建项目目录
```bash
# 创建一个新文件夹
mkdir hello-go
cd hello-go

# 初始化Go模块（重要！）
go mod init hello-go
```

**这一步做了什么？**
```
1. 创建了go.mod文件
2. 告诉Go这是一个独立项目
3. 设定了模块名为"hello-go"
4. 准备好了依赖管理
```

### 编写第一个程序

创建 `main.go` 文件：
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
    fmt.Println("我的第一个Go程序!")
}
```

**代码解释**：
```go
package main     // 告诉Go这是主程序包
import "fmt"     // 导入格式化输出包（像借用工具）
func main() {    // 主函数，程序从这里开始执行
    // 在这里写你的代码
}
```

### 运行程序

**方法一：直接运行**
```bash
go run main.go
```

**方法二：先编译，再运行**
```bash
# 编译生成可执行文件
go build

# 运行可执行文件
./hello-go      # Linux/macOS
hello-go.exe    # Windows
```

**两种方法的区别**：
```
go run：  适合开发测试，每次都重新编译
go build：适合生产部署，生成独立的可执行文件
```

## 常用命令速查

| 命令 | 作用 | 使用场景 |
|------|------|----------|
| `go run main.go` | 直接运行Go文件 | 开发测试 |
| `go build` | 编译生成可执行文件 | 发布程序 |
| `go mod init 项目名` | 初始化模块 | 新建项目 |
| `go mod tidy` | 整理依赖包 | 清理无用依赖 |
| `go get 包名` | 下载依赖包 | 添加第三方库 |
| `go test` | 运行测试 | 测试代码 |
| `go fmt main.go` | 格式化代码 | 统一代码风格 |
| `go version` | 查看Go版本 | 检查安装 |

## 常见问题解决

### 问题1：找不到go命令
```
错误信息：command not found: go
原因：环境变量PATH没设置
解决：
1. 重新打开终端
2. 检查安装路径是否正确
3. 手动添加到PATH环境变量
```

### 问题2：网络连接问题
```
错误信息：timeout、connection refused等
原因：网络问题，无法下载包
解决：设置国内代理（见上面步骤）
```

### 问题3：权限问题
```
错误信息：permission denied
原因：没有写入权限
解决：
- Linux/macOS: 使用sudo或修改文件夹权限
- Windows: 以管理员身份运行
```

### 问题4：VS Code Go扩展无法安装工具
```
原因：网络问题或代理设置
解决：
1. 先设置go env代理
2. 重启VS Code
3. 重新安装Go Tools
```

## 验证环境是否就绪

运行这个检查脚本：
```bash
# 检查Go版本
echo "=== Go版本 ==="
go version

# 检查环境变量
echo "=== 环境变量 ==="
go env GOROOT
go env GOPROXY

# 检查模块支持
echo "=== 创建测试项目 ==="
mkdir go-test && cd go-test
go mod init test
echo 'package main
import "fmt"
func main() { fmt.Println("环境OK!") }' > main.go
go run main.go

# 清理
cd .. && rm -rf go-test
echo "=== 环境检查完成 ==="
```

## 推荐学习资源

**在线工具**：
- [Go Playground](https://play.golang.org/) - 在线运行Go代码
- [Go语言之旅](https://tour.golang.org/) - 官方交互式教程

**文档资源**：
- [Go官方文档](https://golang.org/doc/) - 官方权威文档
- [Go标准库](https://pkg.go.dev/std) - 标准库API参考

**社区论坛**：
- [Go语言中文网](https://studygolang.com/) - 中文社区
- [Stack Overflow](https://stackoverflow.com/questions/tagged/go) - 问题求助

---

**小贴士**: 开发环境就像工具箱，一开始可能觉得复杂，但一旦配置好就能大大提高效率。遇到问题不要慌，Google/百度一下，Go社区很活跃，大部分问题都能找到解决方案！ 