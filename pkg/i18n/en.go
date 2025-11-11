package i18n

var enMessages = map[string]string{
	// App
	"app.title":       "AI Tools Manager (ATM)",
	"app.goodbye":     "Goodbye!",
	"app.initializing": "Initializing...",
	"app.separator":   "─────────────────────────────────────────────",
	"app.error":       "Error",

	// Menu
	"menu.whatToDo":  "What would you like to do?",
	"menu.install":   "Install Tools",
	"menu.query":     "Query Tools",
	"menu.update":    "Update Tools",
	"menu.uninstall": "Uninstall Tools",
	"menu.exit":      "Exit",

	// Prompts
	"prompts.useArrowKeys": "(Use arrow keys)",
	"prompts.pressSpace":   "(Press space to select, enter to confirm)",
	"prompts.confirm":      "(y/N)",

	// Install
	"install.selectToInstall": "Select tools to install:",
	"install.allInstalled":    "All tools are already installed",
	"install.noneSelected":    "No tools selected",
	"install.installing":      "Installing %s...",
	"install.success":         "Successfully installed %s",
	"install.failed":          "Failed to install %s: %s",

	// Query
	"query.noneInstalled":    "No tools installed",
	"query.checking":         "Checking versions...",
	"query.installedTools":   "Installed Tools:",
	"query.version":          "Version:",
	"query.updateAvailable":  "(Update available: v%s)",
	"query.upToDate":         "(Up to date)",
	"query.unknownVersion":   "Unknown",

	// Update
	"update.checking":      "Checking for updates...",
	"update.allUpToDate":   "All tools are up to date",
	"update.selectToUpdate": "Select tools to update:",
	"update.noneSelected":  "No tools selected",
	"update.updating":      "Updating %s...",
	"update.success":       "Successfully updated %s",
	"update.failed":        "Failed to update %s: %s",

	// Uninstall
	"uninstall.noneInstalled":      "No tools installed",
	"uninstall.selectToUninstall":  "Select tools to uninstall:",
	"uninstall.noneSelected":       "No tools selected",
	"uninstall.confirm":            "Are you sure you want to uninstall %d tool(s)?",
	"uninstall.cancelled":          "Uninstall cancelled",
	"uninstall.uninstalling":       "Uninstalling %s...",
	"uninstall.success":            "Successfully uninstalled %s",
	"uninstall.failed":             "Failed to uninstall %s: %s",

	// Version
	"version.checking":          "Checking for updates...",
	"version.updateAvailable":   "A new version of ATM is available!",
	"version.currentVersion":    "Current version: v%s",
	"version.latestVersion":     "Latest version: v%s",
	"version.updatePrompt":      "Would you like to open the repository?",
	"version.openRepository":    "Yes, open repository",
	"version.skipUpdate":        "No, skip for now",
	"version.repositoryOpened":  "Repository opened in browser",
	"version.repositoryOpenFailed": "Could not open browser automatically. Please visit: %s",

	// Config
	"config.loadError": "Failed to load configuration",
}
