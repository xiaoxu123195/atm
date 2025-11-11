# ATM-Go 架构设计文档

## 整体架构

```
┌─────────────────────────────────────────────────────────────┐
│                         User Interface                        │
│                    (Interactive CLI Menu)                     │
└────────────────────────┬──────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                      Main Application                        │
│                    (pkg/app/app.go)                          │
│  ┌─────────────────────────────────────────────────────┐    │
│  │  - 加载配置                                          │    │
│  │  - 初始化工具缓存                                    │    │
│  │  - 版本检查                                          │    │
│  │  - 菜单循环                                          │    │
│  └─────────────────────────────────────────────────────┘    │
└───┬──────────────┬──────────────┬──────────────┬────────────┘
    │              │              │              │
    ▼              ▼              ▼              ▼
┌─────────┐  ┌──────────┐  ┌──────────┐  ┌────────────┐
│  i18n   │  │ Package  │  │ Version  │  │   Config   │
│ Module  │  │ Manager  │  │ Checker  │  │   Loader   │
└─────────┘  └──────────┘  └──────────┘  └────────────┘
    │              │              │              │
    ▼              ▼              ▼              ▼
┌─────────┐  ┌──────────┐  ┌──────────┐  ┌────────────┐
│Language │  │   NPM    │  │  GitHub  │  │   JSON     │
│Detection│  │ Commands │  │   API    │  │   Parser   │
└─────────┘  └──────────┘  └──────────┘  └────────────┘
```

## 模块详细设计

### 1. 主程序入口 (cmd/atm/main.go)

**职责：**
- 程序启动入口
- 初始化应用
- 错误捕获和退出处理

**代码结构：**
```go
package main

import (
    "github.com/yourusername/atm-go/pkg/app"
    "github.com/yourusername/atm-go/pkg/i18n"
)

const VERSION = "1.0.0"

func main() {
    // 1. 初始化国际化
    i18n.Init()

    // 2. 创建应用实例
    application := app.NewApp(VERSION)

    // 3. 运行应用
    if err := application.Run(); err != nil {
        // 错误处理
    }
}
```

### 2. 国际化模块 (pkg/i18n)

**文件结构：**
```
pkg/i18n/
├── i18n.go       # 主逻辑和语言检测
├── en.go         # 英文翻译
└── zh.go         # 中文翻译
```

**核心功能：**
```go
// i18n.go
type I18n struct {
    lang     string
    messages map[string]string
}

func Init() {
    // 检测系统语言
    // 加载对应的翻译
}

func T(key string, args ...interface{}) string {
    // 获取翻译文本
    // 支持参数替换
}

// 语言检测逻辑
func detectLanguage() string {
    // 1. 检查环境变量 LANG
    // 2. 检查 LC_ALL
    // 3. 检查 LC_MESSAGES
    // 4. Windows: 调用系统 API
    // 5. 默认返回 "en"
}
```

**翻译键值设计：**
```go
// zh.go
var zhMessages = map[string]string{
    "app.title":              "AI 工具管理器 (ATM)",
    "app.goodbye":            "再见！",
    "menu.whatToDo":          "你想做什么？",
    "menu.install":           "安装工具",
    "menu.query":             "查询工具",
    "menu.update":            "更新工具",
    "menu.uninstall":         "卸载工具",
    "menu.exit":              "退出",
    "install.success":        "安装 %s 成功",
    // ... 更多翻译
}
```

### 3. 包管理器模块 (pkg/manager)

**文件结构：**
```
pkg/manager/
├── package_manager.go    # 包管理主逻辑
└── npm_command.go        # npm 命令封装
```

**核心结构：**
```go
type PackageManager struct {
    // 可选：缓存 npm 全局路径
    npmPath string
}

// 核心方法
func (pm *PackageManager) IsPackageInstalled(packageName string) (bool, error)
func (pm *PackageManager) GetPackageVersion(packageName string) (string, error)
func (pm *PackageManager) GetLatestVersion(packageName string) (string, error)
func (pm *PackageManager) InstallPackage(packageName string) error
func (pm *PackageManager) UpdatePackage(packageName string) error
func (pm *PackageManager) UninstallPackage(packageName string) error

// 辅助方法
func (pm *PackageManager) executeNpmCommand(args ...string) (string, error)
func (pm *PackageManager) extractPackageName(packageWithVersion string) string
```

**npm 命令映射：**
```go
// 获取已安装包列表
// npm list -g --depth=0 --json

// 获取包版本
// npm list -g <package> --depth=0 --json

// 获取最新版本
// npm view <package> version

// 安装包
// npm install -g <package>

// 更新包
// npm update -g <package>

// 卸载包
// npm uninstall -g <package>
```

**JSON 解析示例：**
```go
type NpmListOutput struct {
    Dependencies map[string]struct {
        Version string `json:"version"`
    } `json:"dependencies"`
}
```

### 4. 版本检查器模块 (pkg/version)

**文件结构：**
```
pkg/version/
└── checker.go
```

**核心功能：**
```go
type Checker struct {
    currentVersion string
    repositoryURL  string
}

type UpdateInfo struct {
    HasUpdate      bool
    CurrentVersion string
    LatestVersion  string
    Error          error
}

func (c *Checker) CheckForUpdates() (*UpdateInfo, error) {
    // 1. 从 GitHub API 获取最新 release
    // GET https://api.github.com/repos/yourusername/atm-go/releases/latest

    // 2. 解析 JSON 获取 tag_name

    // 3. 比较版本号

    // 4. 返回更新信息
}

func (c *Checker) OpenRepository() error {
    // 根据操作系统调用不同命令
    // Windows: start <url>
    // Linux: xdg-open <url>
    // macOS: open <url>
}

// 版本比较
func compareVersions(v1, v2 string) int {
    // 实现语义化版本比较
    // 返回: -1 (v1 < v2), 0 (v1 == v2), 1 (v1 > v2)
}
```

### 5. 主应用逻辑 (pkg/app)

**文件结构：**
```
pkg/app/
├── app.go           # 主应用类
├── config.go        # 配置加载
└── cache.go         # 工具缓存
```

**核心结构：**
```go
type App struct {
    version          string
    config           *Config
    packageManager   *manager.PackageManager
    versionChecker   *version.Checker

    // 缓存
    installedTools   []Tool
    uninstalledTools []Tool
    versionCache     map[string]*VersionInfo
}

type Config struct {
    Tools []Tool `json:"tools"`
}

type Tool struct {
    Name        string `json:"name"`
    Package     string `json:"package"`
    Description string `json:"description"`
}

type VersionInfo struct {
    CurrentVersion string
    LatestVersion  string
}
```

**主流程：**
```go
func (a *App) Run() error {
    // 1. 显示欢迎信息
    fmt.Println(i18n.T("app.title"))

    // 2. 加载配置
    if err := a.loadConfig(); err != nil {
        return err
    }

    // 3. 初始化工具缓存
    a.initializeToolsCache()

    // 4. 检查更新（可跳过）
    if !shouldSkipVersionCheck() {
        a.checkForUpdates()
    }

    // 5. 主菜单循环
    for {
        action := a.showMainMenu()

        switch action {
        case "install":
            a.handleInstall()
        case "query":
            a.handleQuery()
        case "update":
            a.handleUpdate()
        case "uninstall":
            a.handleUninstall()
        case "exit":
            fmt.Println(i18n.T("app.goodbye"))
            return nil
        }
    }
}
```

**菜单处理流程：**
```go
// 安装流程
func (a *App) handleInstall() {
    // 1. 如果所有工具已安装，显示提示并返回
    // 2. 使用 promptui 显示复选框选择未安装的工具
    // 3. 对每个选中的工具：
    //    - 显示 spinner
    //    - 调用 packageManager.InstallPackage()
    //    - 缓存版本信息
    //    - 更新缓存列表（从 uninstalled 移到 installed）
    //    - 显示结果
}

// 查询流程
func (a *App) handleQuery() {
    // 1. 如果没有已安装工具，显示提示并返回
    // 2. 显示 spinner "正在检查..."
    // 3. 并发获取版本信息（使用 goroutine）
    // 4. 格式化输出：
    //    - 工具名称 (包名)
    //    - 当前版本
    //    - 更新状态（有更新 / 已是最新）
    //    - 描述
}

// 更新流程
func (a *App) handleUpdate() {
    // 1. 检查哪些工具有更新
    // 2. 如果没有可更新工具，显示提示并返回
    // 3. 使用 promptui 显示复选框（显示 v1.0 → v2.0）
    // 4. 对每个选中的工具：
    //    - 显示 spinner
    //    - 调用 packageManager.UpdatePackage()
    //    - 更新缓存中的版本信息
    //    - 显示结果
}

// 卸载流程
func (a *App) handleUninstall() {
    // 1. 如果没有已安装工具，显示提示并返回
    // 2. 使用 promptui 显示复选框选择要卸载的工具
    // 3. 二次确认（confirm prompt）
    // 4. 对每个选中的工具：
    //    - 显示 spinner
    //    - 调用 packageManager.UninstallPackage()
    //    - 更新缓存列表（从 installed 移到 uninstalled）
    //    - 从版本缓存中删除
    //    - 显示结果
}
```

## 数据流图

```
用户输入
   │
   ▼
┌──────────────┐
│ Main Menu    │
│ (promptui)   │
└──────┬───────┘
       │
       ├─ Install ──────┐
       ├─ Query ────────┼──────┐
       ├─ Update ───────┼──────┼─────┐
       ├─ Uninstall ────┼──────┼─────┼────┐
       └─ Exit          │      │     │    │
                        ▼      ▼     ▼    ▼
                    ┌──────────────────────────┐
                    │   Package Manager        │
                    │  - executeNpmCommand()   │
                    └────────┬─────────────────┘
                             │
                             ▼
                    ┌──────────────┐
                    │  npm binary  │
                    └────────┬─────┘
                             │
                             ▼
                    ┌──────────────┐
                    │   Output     │
                    │  (JSON/Text) │
                    └────────┬─────┘
                             │
                             ▼
                    ┌──────────────┐
                    │    Parse     │
                    └────────┬─────┘
                             │
                             ▼
                    ┌──────────────┐
                    │  Update UI   │
                    │   & Cache    │
                    └──────────────┘
```

## 配置文件设计

**config/tools.json:**
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

**加载方式（使用 embed）：**
```go
//go:embed config/tools.json
var configFile []byte

func loadConfig() (*Config, error) {
    var config Config
    if err := json.Unmarshal(configFile, &config); err != nil {
        return nil, err
    }
    return &config, nil
}
```

## 错误处理策略

1. **网络错误**: 静默处理，不影响主流程（如版本检查失败）
2. **npm 错误**: 显示用户友好的错误信息
3. **配置错误**: 启动时检查，失败则退出
4. **用户取消**: 正常流程，返回主菜单

## 性能优化

### 1. 并发获取版本信息
```go
func (a *App) fetchVersionsConcurrently(tools []Tool) {
    var wg sync.WaitGroup
    semaphore := make(chan struct{}, 5) // 限制并发数

    for _, tool := range tools {
        wg.Add(1)
        go func(t Tool) {
            defer wg.Done()
            semaphore <- struct{}{} // 获取信号量
            defer func() { <-semaphore }() // 释放信号量

            current := a.packageManager.GetPackageVersion(t.Package)
            latest := a.packageManager.GetLatestVersion(t.Package)
            a.versionCache[t.Package] = &VersionInfo{current, latest}
        }(tool)
    }

    wg.Wait()
}
```

### 2. 缓存机制
- 程序运行期间缓存工具安装状态
- 缓存版本信息，避免重复查询
- 操作后立即更新缓存

### 3. 惰性加载
- 只在需要时才获取版本信息
- 初始化时只检查安装状态（快速）

## 测试策略

### 单元测试
```go
// pkg/manager/package_manager_test.go
func TestExtractPackageName(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"@scope/package", "@scope/package"},
        {"@scope/package@1.0.0", "@scope/package"},
        {"package@1.0.0", "package"},
    }

    for _, tt := range tests {
        result := extractPackageName(tt.input)
        if result != tt.expected {
            t.Errorf("extractPackageName(%s) = %s; want %s",
                tt.input, result, tt.expected)
        }
    }
}
```

### 集成测试
- 测试完整的安装/卸载流程
- 使用模拟的 npm 命令

## 发布清单

- [ ] 所有功能测试通过
- [ ] 编写完整文档
- [ ] 为 Windows/Linux/macOS 构建二进制文件
- [ ] 创建 GitHub Release
- [ ] 更新 README 中的下载链接
- [ ] 添加版本号和更新日志

## 未来扩展

1. **插件系统**: 支持自定义工具配置
2. **批量操作**: 一键安装所有工具
3. **配置同步**: 导出/导入配置
4. **自动更新**: 自动下载并替换新版本
5. **工具推荐**: 根据使用情况推荐工具
