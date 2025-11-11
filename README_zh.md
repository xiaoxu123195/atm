[English](README.md) | [中文](README_zh.md)

# ATM-Go - AI 工具管理器

🔧 使用 Go 语言编写的命令行工具，用于管理 npm 安装的 AI 开发工具。

## 功能特性

- **安装工具**: 安装尚未安装的 AI 开发工具
- **查询工具**: 查看已安装工具的版本和可用更新
- **更新工具**: 将已安装工具更新到最新版本
- **卸载工具**: 从系统中移除已安装的工具
- **多语言支持**: 自动检测系统语言（中文/英文）
- **自动更新检查**: 启动时检查新版本
- **快速轻量**: 单个可执行文件，无需 Node.js
- **跨平台**: 支持 Windows、Linux 和 macOS

## 系统要求

- **npm**: 本工具管理 npm 包，因此需要安装 npm
- 无需其他依赖（运行 exe 文件不需要 Go 运行时）

## 安装

### 快速安装（Windows）

1. 从 [Releases](https://github.com/xiaoxu123195/atm/releases) 下载最新的 `atm.exe`
2. 将 `atm.exe` 移动到 PATH 环境变量包含的目录：
   ```bash
   # 方式 1: 复制到 Windows 系统目录
   copy atm.exe C:\Windows\System32\

   # 方式 2: 复制到 Go bin 目录（如果已安装 Go）
   copy atm.exe %GOPATH%\bin\
   ```
3. 验证安装:
   ```bash
   atm
   ```

### 从源码构建

1. 克隆仓库:
   ```bash
   git clone https://github.com/xiaoxu123195/atm.git
   cd atm-go
   ```

2. 安装依赖:
   ```bash
   go mod download
   ```

3. 构建:
   ```bash
   go build -o bin/atm.exe cmd/atm/main.go
   ```

4. 添加到 PATH 并运行:
   ```bash
   .\bin\atm.exe
   ```

## 使用方法

运行 ATM 工具：
```bash
atm
```

应用程序将显示交互式菜单，你可以：
- 安装新的 AI 工具
- 查询已安装工具并检查更新
- 将现有工具更新到最新版本
- 卸载不再需要的工具

## 支持的工具

目前支持以下 AI 开发工具：

- **Claude Code** - Anthropic 官方 Claude AI 命令行工具
- **Qwen Code** - 通义千问 AI 开发工具
- **Code Buddy** - 腾讯 AI 代码助手
- **Gemini CLI** - Google Gemini AI 命令行界面
- **Auggie** - AI 驱动的代码增强工具
- **Crush** - Charmland 开发工具
- **Codex** - OpenAI Codex 命令行工具
- **iFlow** - iFlow AI 开发命令行工具
- **OpenCode** - 为终端打造的 AI 编码代理
- **Copilot CLI** - GitHub Copilot 命令行工具
- **Kode** - ShareAI Lab 终端 AI 助手

## 配置

工具配置在 `config/tools.json` 文件中。要添加新工具，请编辑配置文件并向 `tools` 数组添加新条目。

## 高级用法

### 禁用版本检查

```bash
# Windows
set ATM_SKIP_VERSION_CHECK=true
atm

# Linux/Mac
ATM_SKIP_VERSION_CHECK=true atm
```

### 强制指定语言

```bash
# 强制英文
set LANG=en_US.UTF-8
atm

# 强制中文
set LANG=zh_CN.UTF-8
atm
```

## 贡献

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

MIT
