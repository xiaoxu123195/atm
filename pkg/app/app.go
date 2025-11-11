package app

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/xiaoxu123195/atm/pkg/config"
	"github.com/xiaoxu123195/atm/pkg/i18n"
	"github.com/xiaoxu123195/atm/pkg/manager"
	versionpkg "github.com/xiaoxu123195/atm/pkg/version"
)

// VersionInfo stores version information for a tool
type VersionInfo struct {
	CurrentVersion string
	LatestVersion  string
}

// App represents the main application
type App struct {
	version        string
	repositoryURL  string
	config         *config.Config
	packageManager *manager.PackageManager
	versionChecker *versionpkg.Checker

	// Cache
	installedTools   []config.Tool
	uninstalledTools []config.Tool
	versionCache     map[string]*VersionInfo
}

// NewApp creates a new application instance
func NewApp(version, repositoryURL string) *App {
	return &App{
		version:          version,
		repositoryURL:    repositoryURL,
		packageManager:   manager.NewPackageManager(),
		versionChecker:   versionpkg.NewChecker(version, repositoryURL),
		versionCache:     make(map[string]*VersionInfo),
		installedTools:   []config.Tool{},
		uninstalledTools: []config.Tool{},
	}
}

// Run starts the application
func (a *App) Run() error {
	// Display welcome message
	fmt.Println(color.CyanString("\n" + i18n.T("app.title") + "\n"))

	// Load configuration
	if err := a.loadConfig(); err != nil {
		return fmt.Errorf("%s: %w", i18n.T("config.loadError"), err)
	}

	// Initialize tools cache
	a.initializeToolsCache()

	// Check for updates (can be skipped with environment variable)
	if os.Getenv("ATM_SKIP_VERSION_CHECK") != "true" {
		a.checkForUpdates()
	}

	// Main menu loop
	for {
		action, err := a.showMainMenu()
		if err != nil {
			return err
		}

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
			fmt.Println(color.GreenString("\n" + i18n.T("app.goodbye") + "\n"))
			return nil
		}

		fmt.Println("\n" + i18n.T("app.separator") + "\n")
	}
}

// loadConfig loads the configuration file
func (a *App) loadConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.config = cfg
	return nil
}

// initializeToolsCache initializes the tools cache by checking installation status
func (a *App) initializeToolsCache() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + i18n.T("app.initializing")
	s.Start()

	for _, tool := range a.config.Tools {
		installed, _ := a.packageManager.IsPackageInstalled(tool.Package)
		if installed {
			a.installedTools = append(a.installedTools, tool)
		} else {
			a.uninstalledTools = append(a.uninstalledTools, tool)
		}
	}

	s.Stop()
}

// checkForUpdates checks for ATM updates
func (a *App) checkForUpdates() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + i18n.T("version.checking")
	s.Start()

	updateInfo := a.versionChecker.CheckForUpdates()
	s.Stop()

	if updateInfo.Error != nil {
		// Silently ignore errors
		return
	}

	if updateInfo.HasUpdate {
		fmt.Println(color.YellowString("\n" + i18n.T("version.updateAvailable")))
		fmt.Println(color.New(color.FgHiBlack).Sprint(i18n.T("version.currentVersion", updateInfo.CurrentVersion)))
		fmt.Println(color.GreenString(i18n.T("version.latestVersion", updateInfo.LatestVersion)))
		fmt.Println()

		prompt := promptui.Select{
			Label: i18n.T("version.updatePrompt") + " " + i18n.T("prompts.useArrowKeys"),
			Items: []string{
				i18n.T("version.openRepository"),
				i18n.T("version.skipUpdate"),
			},
		}

		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == 0 {
			if err := a.versionChecker.OpenRepository(); err == nil {
				fmt.Println(color.GreenString(i18n.T("version.repositoryOpened")))
			} else {
				fmt.Println(color.YellowString(i18n.T("version.repositoryOpenFailed", updateInfo.RepositoryURL)))
			}
			fmt.Println()
		}
	}
}

// showMainMenu displays the main menu and returns the selected action
func (a *App) showMainMenu() (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▸ {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: "▸ {{ . | cyan }}",
	}

	prompt := promptui.Select{
		Label: i18n.T("menu.whatToDo") + " " + i18n.T("prompts.useArrowKeys"),
		Items: []string{
			i18n.T("menu.install"),
			i18n.T("menu.query"),
			i18n.T("menu.update"),
			i18n.T("menu.uninstall"),
			i18n.T("menu.exit"),
		},
		Templates: templates,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	actions := []string{"install", "query", "update", "uninstall", "exit"}
	return actions[index], nil
}

// fetchVersionsConcurrently fetches version information for multiple tools concurrently
func (a *App) fetchVersionsConcurrently(tools []config.Tool) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5) // Limit to 5 concurrent requests

	for _, tool := range tools {
		if _, exists := a.versionCache[tool.Package]; exists {
			continue // Skip if already cached
		}

		wg.Add(1)
		go func(t config.Tool) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			current, _ := a.packageManager.GetPackageVersion(t.Package)
			latest, _ := a.packageManager.GetLatestVersion(t.Package)

			a.versionCache[t.Package] = &VersionInfo{
				CurrentVersion: current,
				LatestVersion:  latest,
			}
		}(tool)
	}

	wg.Wait()
}
