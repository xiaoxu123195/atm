# ATM-Go 开发文档

## 项目概述

ATM-Go 是使用 Go 语言重写的 AI Tools Manager (ATM)，用于管理各种 AI 开发工具的 npm 包安装、更新和卸载。

### 为什么使用 Go？

1. **单文件可执行** - 编译后生成单个 .exe 文件，无需 Node.js 环境
2. **跨平台编译** - 可以轻松为 Windows、Linux、macOS 编译
3. **性能优秀** - 启动速度快，资源占用少
4. **易于分发** - 用户只需下载 exe 文件即可使用

## 项目结构

```
atm-go/
├── cmd/
│   └── atm/
│       └── main.go              # 程序入口
├── pkg/
│   ├── app/
│   │   └── app.go               # 主应用逻辑
│   ├── i18n/
│   │   ├── i18n.go              # 国际化管理器
│   │   ├── zh.go                # 中文翻译
│   │   └── en.go                # 英文翻译
│   ├── manager/
│   │   └── package_manager.go   # NPM 包管理器
│   └── version/
│       └── checker.go           # 版本检查器
├── config/
│   └── tools.json               # AI 工具配置文件
├── build/
│   ├── build.bat                # Windows 构建脚本
│   └── build.sh                 # Linux/Mac 构建脚本
├── go.mod                       # Go 模块定义
├── go.sum                       # 依赖校验和
├── DEV_GUIDE.md                 # 开发文档（本文件）
└── README.md                    # 用户文档
```

## 技术栈

### 核心依赖库

| 库名 | 版本 | 用途 | 对应 Node.js 库 |
|------|------|------|----------------|
| `github.com/manifoldco/promptui` | v0.9.0 | 交互式命令行界面 | inquirer |
| `github.com/fatih/color` | v1.16.0 | 终端彩色输出 | chalk |
| `github.com/briandowns/spinner` | v1.23.0 | 加载动画 | ora |
| `golang.org/x/sys` | latest | 系统调用（语言检测） | - |
| `encoding/json` | 标准库 | JSON 处理 | - |
| `embed` | 标准库 | 内嵌配置文件 | - |

### 安装依赖

```bash
cd D:/go_pro/ai_tools_manager/atm-go
go get github.com/manifoldco/promptui@v0.9.0
go get github.com/fatih/color@v1.16.0
go get github.com/briandowns/spinner@v1.23.0
go mod tidy
```

## 功能模块详解

### 1. 国际化模块 (pkg/i18n)

**功能：** 自动检测系统语言，支持中文和英文

**实现要点：**
- 检测环境变量 `LANG`, `LC_ALL`, `LC_MESSAGES`
- 提供 `T(key string, args ...interface{})` 方法获取翻译文本
- 使用 Go 的 map 存储翻译文本

**示例：**
```go
i18n.T("menu.install")          // "安装工具" 或 "Install Tools"
i18n.T("install.success", name) // "安装 {name} 成功"
```

### 2. 包管理器模块 (pkg/manager)

**功能：** 封装 npm 命令操作

**核心方法：**
- `IsPackageInstalled(packageName string) bool` - 检查包是否已安装
- `GetPackageVersion(packageName string) string` - 获取已安装包的版本
- `GetLatestVersion(packageName string) string` - 获取最新版本
- `InstallPackage(packageName string) error` - 安装包
- `UpdatePackage(packageName string) error` - 更新包
- `UninstallPackage(packageName string) error` - 卸载包

**实现方式：**
使用 `os/exec` 包执行 npm 命令，解析 JSON 输出

### 3. 版本检查器模块 (pkg/version)

**功能：** 检查 ATM 自身是否有新版本

**实现要点：**
- 从 GitHub API 获取最新 release 信息
- 比较当前版本与最新版本
- 提供打开浏览器功能（使用 `open` 命令）

### 4. 主应用模块 (pkg/app)

**功能：** 协调各模块，实现交互式菜单

**核心流程：**
1. 加载配置文件 (tools.json)
2. 初始化工具缓存（检查安装状态）
3. 检查 ATM 版本更新
4. 显示主菜单（安装/查询/更新/卸载/退出）
5. 处理用户选择

**状态缓存：**
- `installedTools []Tool` - 已安装工具列表
- `uninstalledTools []Tool` - 未安装工具列表
- `versionCache map[string]VersionInfo` - 版本信息缓存

## 配置文件格式

### tools.json

```json
{
  "tools": [
    {
      "name": "Claude Code",
      "package": "@anthropic-ai/claude-code",
      "description": "Anthropic's official CLI for Claude AI"
    }
  ]
}
```

## 构建和安装

### 开发模式运行

```bash
cd D:/go_pro/ai_tools_manager/atm-go
go run cmd/atm/main.go
```

### 构建可执行文件

**Windows:**
```bash
cd D:/go_pro/ai_tools_manager/atm-go
go build -o bin/atm.exe cmd/atm/main.go
```

**优化构建（减小文件大小）:**
```bash
go build -ldflags="-s -w" -o bin/atm.exe cmd/atm/main.go
```

**交叉编译（在 Windows 上编译 Linux 版本）:**
```bash
set GOOS=linux
set GOARCH=amd64
go build -o bin/atm-linux cmd/atm/main.go
```

### 全局安装

**方法 1: 复制到 PATH 目录**
```bash
# 将 atm.exe 复制到 C:\Windows\System32
copy bin\atm.exe C:\Windows\System32\

# 或复制到 Go 的 bin 目录（通常已在 PATH 中）
copy bin\atm.exe %GOPATH%\bin\
```

**方法 2: 添加自定义目录到 PATH**
1. 创建目录如 `C:\Tools`
2. 将 `atm.exe` 复制到该目录
3. 将 `C:\Tools` 添加到系统环境变量 PATH

**验证安装：**
```bash
atm  # 在任意目录执行
```

## 开发计划

### 阶段 1: 基础框架（1-2天）
- [x] 创建项目结构
- [ ] 实现国际化模块
- [ ] 实现配置文件加载
- [ ] 创建基本的 main 函数

### 阶段 2: 核心功能（2-3天）
- [ ] 实现包管理器模块
- [ ] 实现工具缓存机制
- [ ] 实现安装功能
- [ ] 实现查询功能
- [ ] 实现更新功能
- [ ] 实现卸载功能

### 阶段 3: 增强功能（1天）
- [ ] 实现版本检查器
- [ ] 添加错误处理
- [ ] 优化用户体验

### 阶段 4: 测试和发布（1天）
- [ ] 功能测试
- [ ] 编写构建脚本
- [ ] 编写用户文档
- [ ] 打包发布

## 编码规范

### 命名约定
- **包名**: 小写，简短，如 `manager`, `i18n`
- **文件名**: 小写下划线，如 `package_manager.go`
- **类型名**: 大驼峰（导出），如 `PackageManager`
- **函数名**: 大驼峰（导出），小驼峰（私有），如 `InstallPackage`, `parseVersion`
- **常量**: 大驼峰或全大写，如 `DefaultTimeout`, `VERSION`

### 错误处理
```go
// 好的做法
if err != nil {
    return fmt.Errorf("failed to install package: %w", err)
}

// 避免
if err != nil {
    panic(err)  // 不要使用 panic
}
```

### 日志输出
使用 `color` 包进行彩色输出：
```go
color.Green("✓ 安装成功")
color.Red("✗ 安装失败: %s", err)
color.Yellow("⚠ 警告信息")
color.Cyan("ℹ 提示信息")
```

## 调试技巧

### 1. 调试 npm 命令
```go
// 打印实际执行的命令
fmt.Printf("Executing: %s\n", cmd.String())
```

### 2. 查看 JSON 解析
```go
// 打印原始 JSON
fmt.Printf("Raw JSON: %s\n", string(output))
```

### 3. 使用 Go 调试器
```bash
# 安装 delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 调试运行
dlv debug cmd/atm/main.go
```

## 常见问题

### Q1: 执行 npm 命令失败？
**A:** 确保系统已安装 Node.js 和 npm，且在 PATH 中

### Q2: 中文显示乱码？
**A:** Windows 终端需要设置为 UTF-8 编码：
```bash
chcp 65001
```

### Q3: 编译后文件太大？
**A:** 使用 `-ldflags="-s -w"` 去除调试信息：
```bash
go build -ldflags="-s -w" -o bin/atm.exe cmd/atm/main.go
```

### Q4: 如何嵌入配置文件？
**A:** 使用 Go 1.16+ 的 embed 功能：
```go
//go:embed config/tools.json
var toolsConfig []byte
```

## 测试清单

- [ ] 安装单个工具
- [ ] 安装多个工具
- [ ] 查询已安装工具
- [ ] 查询时显示版本信息
- [ ] 更新工具
- [ ] 卸载工具
- [ ] 中文界面显示正常
- [ ] 英文界面显示正常
- [ ] 版本检查功能
- [ ] 错误处理（网络错误、npm 错误等）

## 性能优化

1. **并发检查版本**: 使用 goroutine 并发检查多个工具的版本
2. **缓存机制**: 缓存已检查的版本信息，避免重复请求
3. **配置预加载**: 使用 embed 嵌入配置文件，减少 IO 操作

## 发布流程

1. **版本号管理**: 在 `cmd/atm/main.go` 中定义 `VERSION` 常量
2. **构建多平台版本**:
   ```bash
   # Windows
   GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o release/atm-windows-amd64.exe cmd/atm/main.go

   # Linux
   GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o release/atm-linux-amd64 cmd/atm/main.go

   # macOS
   GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o release/atm-darwin-amd64 cmd/atm/main.go
   ```
3. **打包发布**: 创建 GitHub Release 并上传二进制文件

## 参考资源

- [Go 官方文档](https://golang.org/doc/)
- [promptui 文档](https://github.com/manifoldco/promptui)
- [color 文档](https://github.com/fatih/color)
- [原 Node.js 版本](../ai-tools-manager/)

## 联系方式

如有问题，请提交 Issue 或 Pull Request。
