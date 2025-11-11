[English](README.md) | [中文](README_zh.md)

# ATM-Go - AI Tools Manager

🔧 A command-line interface for managing npm-installed AI development tools, written in Go.

## Features

- **Install Tools**: Install AI development tools that aren't currently installed
- **Query Tools**: Check installed tools, their versions, and available updates
- **Update Tools**: Update installed tools to their latest versions
- **Uninstall Tools**: Remove installed tools from your system
- **Multi-language Support**: Automatically detects system language (Chinese/English)
- **Auto Update Check**: Checks for newer versions on startup
- **Fast & Lightweight**: Single executable file, no Node.js required
- **Cross-platform**: Works on Windows, Linux, and macOS

## Prerequisites

- **npm**: The tool manages npm packages, so you need npm installed
- No other dependencies required (Go runtime is not needed for the executable)

## Installation

### Quick Install (Windows)

1. Download the latest `atm.exe` from [Releases](https://github.com/xiaoxu123195/atm/releases)
2. Move `atm.exe` to a directory in your PATH:
   ```bash
   # Option 1: Copy to Windows System directory
   copy atm.exe C:\Windows\System32\

   # Option 2: Copy to Go bin directory (if Go is installed)
   copy atm.exe %GOPATH%\bin\
   ```
3. Verify installation:
   ```bash
   atm
   ```

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/xiaoxu123195/atm.git
   cd atm-go
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build:
   ```bash
   go build -o bin/atm.exe cmd/atm/main.go
   ```

4. Add to PATH and run:
   ```bash
   .\bin\atm.exe
   ```

## Usage

Run the ATM tool:
```bash
atm
```

The application will present an interactive menu where you can:
- Install new AI tools
- Query installed tools and check for updates
- Update existing tools to latest versions
- Uninstall tools you no longer need

## Supported Tools

The following AI development tools are supported:

- **Claude Code** - Anthropic's official CLI for Claude AI
- **Qwen Code** - Qwen AI development tools
- **Code Buddy** - Tencent AI code assistant
- **Gemini CLI** - Google Gemini AI command line interface
- **Auggie** - AI-powered code augmentation tool
- **Crush** - Charmland development tool
- **Codex** - OpenAI Codex CLI tool
- **iFlow** - iFlow AI development CLI
- **OpenCode** - AI coding agent, built for the terminal
- **Copilot CLI** - GitHub Copilot CLI
- **Kode** - ShareAI Lab terminal AI assistant

## Configuration

Tools are configured in `config/tools.json`. To add new tools, edit the configuration file and add new entries to the `tools` array.

## Advanced Usage

### Disable Version Check

```bash
# Windows
set ATM_SKIP_VERSION_CHECK=true
atm

# Linux/Mac
ATM_SKIP_VERSION_CHECK=true atm
```

### Force Language

```bash
# Force English
set LANG=en_US.UTF-8
atm

# Force Chinese
set LANG=zh_CN.UTF-8
atm
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT
