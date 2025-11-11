package i18n

var zhMessages = map[string]string{
	// App
	"app.title":        "AI 工具管理器 (ATM)",
	"app.goodbye":      "再见！",
	"app.initializing": "正在初始化...",
	"app.separator":    "─────────────────────────────────────────────",
	"app.error":        "错误",

	// Menu
	"menu.whatToDo":  "你想做什么？",
	"menu.install":   "安装工具",
	"menu.query":     "查询工具",
	"menu.update":    "更新工具",
	"menu.uninstall": "卸载工具",
	"menu.exit":      "退出",

	// Prompts
	"prompts.useArrowKeys": "(使用方向键选择)",
	"prompts.pressSpace":   "(空格键选择，回车键确认)",
	"prompts.confirm":      "(y/N)",

	// Install
	"install.selectToInstall": "选择要安装的工具：",
	"install.allInstalled":    "所有工具已安装",
	"install.noneSelected":    "未选择任何工具",
	"install.installing":      "正在安装 %s...",
	"install.success":         "成功安装 %s",
	"install.failed":          "安装 %s 失败：%s",

	// Query
	"query.noneInstalled":   "未安装任何工具",
	"query.checking":        "正在检查版本...",
	"query.installedTools":  "已安装的工具：",
	"query.version":         "版本：",
	"query.updateAvailable": "(可更新至：v%s)",
	"query.upToDate":        "(已是最新)",
	"query.unknownVersion":  "未知",

	// Update
	"update.checking":       "正在检查更新...",
	"update.allUpToDate":    "所有工具已是最新版本",
	"update.selectToUpdate": "选择要更新的工具：",
	"update.noneSelected":   "未选择任何工具",
	"update.updating":       "正在更新 %s...",
	"update.success":        "成功更新 %s",
	"update.failed":         "更新 %s 失败：%s",

	// Uninstall
	"uninstall.noneInstalled":     "未安装任何工具",
	"uninstall.selectToUninstall": "选择要卸载的工具：",
	"uninstall.noneSelected":      "未选择任何工具",
	"uninstall.confirm":           "确定要卸载 %d 个工具吗？",
	"uninstall.cancelled":         "已取消卸载",
	"uninstall.uninstalling":      "正在卸载 %s...",
	"uninstall.success":           "成功卸载 %s",
	"uninstall.failed":            "卸载 %s 失败：%s",

	// Version
	"version.checking":             "正在检查更新...",
	"version.updateAvailable":      "ATM 有新版本可用！",
	"version.currentVersion":       "当前版本：v%s",
	"version.latestVersion":        "最新版本：v%s",
	"version.updatePrompt":         "是否打开仓库？",
	"version.openRepository":       "是，打开仓库",
	"version.skipUpdate":           "否，暂时跳过",
	"version.repositoryOpened":     "已在浏览器中打开仓库",
	"version.repositoryOpenFailed": "无法自动打开浏览器，请访问：%s",

	// Config
	"config.loadError": "加载配置失败",
}
