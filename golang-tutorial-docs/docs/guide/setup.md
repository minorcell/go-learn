# 环境搭建

本章将指导你安装和配置Go语言开发环境。

## 🖥️ 系统要求

Go语言支持主流操作系统：
- **Windows** 7 或更高版本
- **macOS** 10.11 或更高版本  
- **Linux** 内核 2.6.23 或更高版本

## 📥 安装Go语言

### Windows安装

1. 访问 [Go官网](https://golang.org/dl/)
2. 下载 `go1.22.x.windows-amd64.msi`
3. 双击运行安装程序
4. 按默认选项安装即可

### macOS安装

**方法一：官方安装包**
1. 下载 `go1.22.x.darwin-amd64.pkg`
2. 双击安装包，按提示安装

**方法二：Homebrew**
```bash
brew install go
```

### Linux安装

**方法一：官方二进制包**
```bash
# 下载安装包
wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz

# 解压到/usr/local
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# 添加到PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**方法二：包管理器**
```bash
# Ubuntu/Debian
sudo apt install golang-go

# CentOS/RHEL
sudo yum install golang

# Arch Linux
sudo pacman -S go
```

## ✅ 验证安装

```bash
go version
```

应该看到类似输出：
```
go version go1.22.0 linux/amd64
```

## ⚙️ 环境配置

### 1. GOPATH和GOROOT

**GOROOT** - Go安装目录（通常自动设置）
```bash
go env GOROOT
```

**GOPATH** - 工作空间目录（Go 1.11后可选）
```bash
# 设置工作空间
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### 2. Go模块（推荐）

从Go 1.11开始，推荐使用Go Modules：
```bash
# 查看模块设置
go env GOPROXY
go env GOMOD

# 启用模块支持
go env -w GO111MODULE=on
```

### 3. 代理设置（中国用户）

设置代理加速模块下载：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```

## 🛠️ 开发工具

### 1. VS Code（推荐）

**安装VS Code**
- 访问 [VS Code官网](https://code.visualstudio.com/)
- 下载并安装

**安装Go扩展**
1. 打开VS Code
2. 按 `Ctrl+Shift+X` 打开扩展面板
3. 搜索"Go"
4. 安装Google的Go扩展

**配置Go扩展**
1. 按 `Ctrl+Shift+P` 打开命令面板
2. 输入"Go: Install/Update Tools"
3. 选择所有工具并安装

### 2. GoLand

JetBrains出品的专业Go IDE：
- 强大的代码补全和调试功能
- 内置版本控制支持
- 30天免费试用

### 3. Vim/Neovim

配置vim-go插件：
```vim
" 在.vimrc中添加
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
```

## 🏗️ 创建第一个项目

### 1. 初始化模块
```bash
# 创建项目目录
mkdir hello-go
cd hello-go

# 初始化Go模块
go mod init hello-go
```

### 2. 编写代码
创建 `main.go` 文件：
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### 3. 运行程序
```bash
# 直接运行
go run main.go

# 编译后运行
go build
./hello-go  # Linux/macOS
# 或 hello-go.exe  # Windows
```

## 🔧 常用命令

| 命令 | 用途 |
|------|------|
| `go run` | 编译并运行Go程序 |
| `go build` | 编译程序生成可执行文件 |
| `go mod init` | 初始化新模块 |
| `go mod tidy` | 添加缺失模块，删除无用模块 |
| `go get` | 下载依赖包 |
| `go test` | 运行测试 |
| `go fmt` | 格式化代码 |
| `go vet` | 静态分析工具 |

## 🔍 故障排除

### 问题1：command not found
**原因**：PATH环境变量未设置
**解决**：确保将Go的bin目录添加到PATH

### 问题2：网络连接问题
**原因**：无法访问Go模块代理
**解决**：设置国内代理镜像

### 问题3：权限问题
**原因**：没有写入权限
**解决**：使用sudo或更改目录权限

## 📚 学习资源

- [Go官方文档](https://golang.org/doc/)
- [Go语言之旅](https://tour.golang.org/)
- [Go标准库](https://pkg.go.dev/std)
- [Go播客](https://changelog.com/gotime)

## 🎉 下一步

环境搭建完成后，查看 [学习路线](./roadmap) 制定学习计划，然后开始 [基础语法](/basics/) 的学习！ 