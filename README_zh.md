<div align="center">

# 🔧 ATM - AI 工具管理器

**强大的 AI 开发工具命令行管理器，使用 Go 语言编写**

[![Go 版本](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![发行版](https://img.shields.io/github/v/release/xiaoxu123195/atm?style=flat&color=blue)](https://github.com/xiaoxu123195/atm/releases)
[![许可证](https://img.shields.io/badge/License-MIT-green.svg?style=flat)](LICENSE)
[![平台](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=flat)]()

[English](README.md) | [中文](README_zh.md)

![演示](https://via.placeholder.com/800x400/1a1a1a/00d9ff?text=ATM+-+AI+Tools+Manager)

</div>

---

## ✨ 特性

- 🚀 **快速轻量** - 单个可执行文件，无需 Node.js
- 📦 **包管理** - 安装、更新、查询和卸载 AI 工具
- 🌍 **多语言** - 自动检测系统语言（中文/英文）
- 🎨 **交互式界面** - 精美的命令行用户界面
- ⚡ **并发操作** - 使用 goroutine 实现快速版本检查
- 🔄 **自动更新检查** - 有新版本时自动通知
- 🔧 **易于配置** - 简单的 JSON 配置文件
- 💻 **跨平台** - 支持 Windows、Linux 和 macOS

## 📋 支持的 AI 工具

ATM 目前支持 **11 个 AI 开发工具**：

| 工具 | 包名 | 描述 |
|------|------|------|
| 🤖 Claude Code | `@anthropic-ai/claude-code` | Anthropic 官方 Claude AI 命令行工具 |
| 🧠 Qwen Code | `@qwen-code/qwen-code` | 通义千问 AI 开发工具 |
| 💼 Code Buddy | `@tencent-ai/codebuddy-code` | 腾讯 AI 代码助手 |
| ✨ Gemini CLI | `@google/gemini-cli` | Google Gemini AI 命令行工具 |
| 🔮 Auggie | `@augmentcode/auggie` | AI 驱动的代码增强工具 |
| 💪 Crush | `@charmland/crush` | Charmland 开发工具 |
| 📝 Codex | `@openai/codex` | OpenAI Codex 命令行工具 |
| 🌊 iFlow | `@iflow-ai/iflow-cli` | iFlow AI 开发命令行工具 |
| 🚀 OpenCode | `opencode-ai` | 为终端打造的 AI 编码代理 |
| 🤝 Copilot CLI | `@github/copilot` | GitHub Copilot 命令行工具 |
| 🎯 Kode | `@shareai-lab/kode` | ShareAI Lab 终端助手 |

## 📥 安装

### 方式 1：下载预编译二进制文件（推荐）

从 [Releases](https://github.com/xiaoxu123195/atm/releases) 下载最新版本

**Windows：**
```bash
# 下载 atm.exe 并移动到系统目录
copy atm.exe C:\Windows\System32\
atm
```

**Linux/macOS：**
```bash
# 下载并安装
sudo cp atm /usr/local/bin/
sudo chmod +x /usr/local/bin/atm
atm
```

### 方式 2：通过 Go 安装

```bash
go install github.com/xiaoxu123195/atm@latest
```

### 方式 3：从源码构建

```bash
# 克隆仓库
git clone https://github.com/xiaoxu123195/atm.git
cd atm

# 安装依赖
go mod download

# 构建
go build -o atm cmd/atm/main.go

# 运行
./atm
```

## 🚀 快速开始

直接运行：
```bash
atm
```

你会看到交互式菜单：

```
AI 工具管理器 (ATM)

你想做什么？ (使用方向键选择)
▸ 安装工具
  查询工具
  更新工具
  卸载工具
  退出
```

### 基本操作

**安装工具：**
1. 选择"安装工具"
2. 从列表中选择工具
3. 等待安装完成

**查询已安装的工具：**
1. 选择"查询工具"
2. 查看所有已安装工具及版本信息

**更新工具：**
1. 选择"更新工具"
2. 选择要更新的工具
3. 确认并等待

**卸载工具：**
1. 选择"卸载工具"
2. 选择要移除的工具
3. 确认卸载

## ⚙️ 配置

### 环境变量

```bash
# 禁用版本检查
export ATM_SKIP_VERSION_CHECK=true

# 强制语言
export LANG=zh_CN.UTF-8  # 中文
export LANG=en_US.UTF-8  # 英文
```

### 添加自定义工具

编辑 `config/tools.json`：

```json
{
  "tools": [
    {
      "name": "你的工具名称",
      "package": "npm-package-name",
      "description": "工具描述"
    }
  ]
}
```

然后重新构建：
```bash
go build -o atm cmd/atm/main.go
```

## 📊 对比

| 特性 | Node.js 版本 | ATM (Go) |
|------|-------------|----------|
| 运行时 | 需要 Node.js | ✅ 无需依赖 |
| 安装 | `npm install -g` | ✅ 单个二进制文件 |
| 启动速度 | ~500ms | ✅ ~50ms（快 10 倍）|
| 大小 | ~10MB + node_modules | ✅ 单个 9.7MB exe |
| 内存占用 | ~50MB | ✅ ~20MB |
| 交叉编译 | ❌ | ✅ 简单 |

## 🔧 系统要求

**运行时：**
- npm（必需 - ATM 管理 npm 包）
- 无需 Go 运行时（已编译为二进制文件）

**开发：**
- Go 1.21+
- npm

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。

1. Fork 本项目
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 📝 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

本项目是对 1e0n 的原始 [ai-tools-manager](https://github.com/1e0n/ai-tools-manager) 项目的 Go 语言重新实现。

## 📮 联系方式

- GitHub: [@xiaoxu123195](https://github.com/xiaoxu123195)
- 项目链接: [https://github.com/xiaoxu123195/atm](https://github.com/xiaoxu123195/atm)

---

<div align="center">

**[⬆ 回到顶部](#-atm---ai-工具管理器)**

用 ❤️ 制作 by xiaoxu123195

</div>
