<div align="center">

# 🔧 ATM - AI Tools Manager

**A powerful CLI tool for managing npm-installed AI development tools, written in Go**

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/xiaoxu123195/atm?style=flat&color=blue)](https://github.com/xiaoxu123195/atm/releases)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=flat)]()

[English](README.md) | [中文](README_zh.md)

![Demo](https://via.placeholder.com/800x400/1a1a1a/00d9ff?text=ATM+-+AI+Tools+Manager)

</div>

---

## ✨ Features

- 🚀 **Fast & Lightweight** - Single executable, no Node.js required
- 📦 **Package Management** - Install, update, query, and uninstall AI tools
- 🌍 **Multi-language** - Auto-detects system language (English/Chinese)
- 🎨 **Interactive UI** - Beautiful command-line interface
- ⚡ **Concurrent Operations** - Fast version checking with goroutines
- 🔄 **Auto Update Check** - Notifies when new versions are available
- 🔧 **Easy Configuration** - Simple JSON-based tool configuration
- 💻 **Cross-platform** - Works on Windows, Linux, and macOS

## 📋 Supported AI Tools

ATM currently supports **11 AI development tools**:

| Tool | Package | Description |
|------|---------|-------------|
| 🤖 Claude Code | `@anthropic-ai/claude-code` | Anthropic's official CLI for Claude AI |
| 🧠 Qwen Code | `@qwen-code/qwen-code` | Qwen AI development tools |
| 💼 Code Buddy | `@tencent-ai/codebuddy-code` | Tencent AI code assistant |
| ✨ Gemini CLI | `@google/gemini-cli` | Google Gemini AI CLI |
| 🔮 Auggie | `@augmentcode/auggie` | AI-powered code augmentation |
| 💪 Crush | `@charmland/crush` | Charmland development tool |
| 📝 Codex | `@openai/codex` | OpenAI Codex CLI tool |
| 🌊 iFlow | `@iflow-ai/iflow-cli` | iFlow AI development CLI |
| 🚀 OpenCode | `opencode-ai` | AI coding agent for terminal |
| 🤝 Copilot CLI | `@github/copilot` | GitHub Copilot CLI |
| 🎯 Kode | `@shareai-lab/kode` | ShareAI Lab terminal assistant |

## 📥 Installation

### Option 1: Download Pre-built Binary (Recommended)

Download the latest release from [Releases](https://github.com/xiaoxu123195/atm/releases)

**Windows:**
```bash
# Download atm.exe and move to system directory
copy atm.exe C:\Windows\System32\
atm
```

**Linux/macOS:**
```bash
# Download and install
sudo cp atm /usr/local/bin/
sudo chmod +x /usr/local/bin/atm
atm
```

### Option 2: Install via Go

```bash
go install github.com/xiaoxu123195/atm@latest
```

### Option 3: Build from Source

```bash
# Clone repository
git clone https://github.com/xiaoxu123195/atm.git
cd atm

# Install dependencies
go mod download

# Build
go build -o atm cmd/atm/main.go

# Run
./atm
```

## 🚀 Quick Start

Simply run:
```bash
atm
```

You'll see an interactive menu:

```
AI 工具管理器 (ATM)

你想做什么？ (使用方向键选择)
▸ 安装工具
  查询工具
  更新工具
  卸载工具
  退出
```

### Basic Operations

**Install a tool:**
1. Select "Install Tools"
2. Choose tools from the list
3. Wait for installation

**Query installed tools:**
1. Select "Query Tools"
2. View all installed tools with version info

**Update tools:**
1. Select "Update Tools"
2. Choose tools to update
3. Confirm and wait

**Uninstall tools:**
1. Select "Uninstall Tools"
2. Choose tools to remove
3. Confirm uninstallation

## ⚙️ Configuration

### Environment Variables

```bash
# Disable version check
export ATM_SKIP_VERSION_CHECK=true

# Force language
export LANG=zh_CN.UTF-8  # Chinese
export LANG=en_US.UTF-8  # English
```

### Add Custom Tools

Edit `config/tools.json`:

```json
{
  "tools": [
    {
      "name": "Your Tool Name",
      "package": "npm-package-name",
      "description": "Tool description"
    }
  ]
}
```

Then rebuild:
```bash
go build -o atm cmd/atm/main.go
```

## 🔧 Requirements

**Runtime:**
- npm (required - ATM manages npm packages)
- No Go runtime needed (compiled binary)

**Development:**
- Go 1.21+
- npm

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📮 Contact

- GitHub: [@xiaoxu123195](https://github.com/xiaoxu123195)
- Project Link: [https://github.com/xiaoxu123195/atm](https://github.com/xiaoxu123195/atm)

---

<div align="center">

**[⬆ Back to Top](#-atm---ai-tools-manager)**

Made with ❤️ by xiaoxu123195

</div>
