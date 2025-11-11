package app

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/xiaoxu123195/atm/pkg/config"
	"github.com/xiaoxu123195/atm/pkg/i18n"
)

// handleInstall handles the install action
func (a *App) handleInstall() {
	if len(a.uninstalledTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("install.allInstalled")))
		return
	}

	// Create choices
	items := make([]string, len(a.uninstalledTools))
	for i, tool := range a.uninstalledTools {
		items[i] = fmt.Sprintf("%s (%s)", tool.Name, tool.Package)
	}

	// Custom templates for checkbox
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▸ {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: "▸ {{ . | cyan }}",
	}

	// Use a simple select for now (promptui doesn't have built-in multi-select)
	// We'll use a workaround: show menu multiple times or use a different approach
	prompt := promptui.Select{
		Label:     i18n.T("install.selectToInstall") + " " + i18n.T("prompts.useArrowKeys"),
		Items:     append(items, i18n.T("menu.exit")),
		Templates: templates,
		Size:      10,
	}

	var selectedTools []config.Tool

	for {
		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == len(items) { // Exit option
			break
		}

		selectedTools = append(selectedTools, a.uninstalledTools[index])

		// Ask if user wants to select more
		confirmPrompt := promptui.Select{
			Label: "Select more tools?",
			Items: []string{"Yes", "No, start installation"},
		}

		choice, _, _ := confirmPrompt.Run()
		if choice == 1 {
			break
		}
	}

	if len(selectedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("install.noneSelected")))
		return
	}

	// Install selected tools
	for _, tool := range selectedTools {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " " + i18n.T("install.installing", tool.Name)
		s.Start()

		err := a.packageManager.InstallPackage(tool.Package)
		s.Stop()

		if err != nil {
			fmt.Println(color.RedString("✗ " + i18n.T("install.failed", tool.Name, err.Error())))
		} else {
			// Cache version info
			current, _ := a.packageManager.GetPackageVersion(tool.Package)
			latest, _ := a.packageManager.GetLatestVersion(tool.Package)
			a.versionCache[tool.Package] = &VersionInfo{
				CurrentVersion: current,
				LatestVersion:  latest,
			}

			fmt.Println(color.GreenString("✓ " + i18n.T("install.success", tool.Name)))

			// Update cache lists
			a.installedTools = append(a.installedTools, tool)
			a.removeFromUninstalled(tool.Package)
		}
	}
}

// handleQuery handles the query action
func (a *App) handleQuery() {
	if len(a.installedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("query.noneInstalled")))
		return
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + i18n.T("query.checking")
	s.Start()

	// Fetch version info concurrently
	a.fetchVersionsConcurrently(a.installedTools)

	s.Stop()

	fmt.Println(color.CyanString("\n" + i18n.T("query.installedTools") + "\n"))

	for _, tool := range a.installedTools {
		versionInfo := a.versionCache[tool.Package]
		if versionInfo == nil {
			versionInfo = &VersionInfo{"", ""}
		}

		versionText := versionInfo.CurrentVersion
		if versionText == "" {
			versionText = i18n.T("query.unknownVersion")
		} else {
			versionText = "v" + versionText
		}

		updateText := ""
		if versionInfo.CurrentVersion != "" && versionInfo.LatestVersion != "" &&
			versionInfo.CurrentVersion != versionInfo.LatestVersion {
			updateText = color.YellowString(" " + i18n.T("query.updateAvailable", versionInfo.LatestVersion))
		} else if versionInfo.CurrentVersion != "" {
			updateText = color.GreenString(" " + i18n.T("query.upToDate"))
		}

		fmt.Printf("%s %s %s\n",
			color.BlueString("•"),
			color.New(color.Bold).Sprint(tool.Name),
			color.New(color.FgHiBlack).Sprintf("(%s)", tool.Package))
		fmt.Printf("  %s %s%s\n",
			color.New(color.FgHiBlack).Sprint(i18n.T("query.version")),
			versionText,
			updateText)
		fmt.Printf("  %s\n\n",
			color.New(color.FgHiBlack).Sprint(tool.Description))
	}
}

// handleUpdate handles the update action
func (a *App) handleUpdate() {
	if len(a.installedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("query.noneInstalled")))
		return
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + i18n.T("update.checking")
	s.Start()

	// Fetch version info concurrently
	a.fetchVersionsConcurrently(a.installedTools)

	// Find updatable tools
	var updatableTools []config.Tool
	for _, tool := range a.installedTools {
		versionInfo := a.versionCache[tool.Package]
		if versionInfo != nil &&
			versionInfo.CurrentVersion != "" &&
			versionInfo.LatestVersion != "" &&
			versionInfo.CurrentVersion != versionInfo.LatestVersion {
			updatableTools = append(updatableTools, tool)
		}
	}

	s.Stop()

	if len(updatableTools) == 0 {
		fmt.Println(color.GreenString(i18n.T("update.allUpToDate")))
		return
	}

	// Create choices with version info
	items := make([]string, len(updatableTools))
	for i, tool := range updatableTools {
		versionInfo := a.versionCache[tool.Package]
		items[i] = fmt.Sprintf("%s (v%s → v%s)",
			tool.Name,
			versionInfo.CurrentVersion,
			versionInfo.LatestVersion)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▸ {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: "▸ {{ . | cyan }}",
	}

	prompt := promptui.Select{
		Label:     i18n.T("update.selectToUpdate") + " " + i18n.T("prompts.useArrowKeys"),
		Items:     append(items, i18n.T("menu.exit")),
		Templates: templates,
		Size:      10,
	}

	var selectedTools []config.Tool

	for {
		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == len(items) { // Exit option
			break
		}

		selectedTools = append(selectedTools, updatableTools[index])

		confirmPrompt := promptui.Select{
			Label: "Select more tools?",
			Items: []string{"Yes", "No, start update"},
		}

		choice, _, _ := confirmPrompt.Run()
		if choice == 1 {
			break
		}
	}

	if len(selectedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("update.noneSelected")))
		return
	}

	// Update selected tools
	for _, tool := range selectedTools {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " " + i18n.T("update.updating", tool.Name)
		s.Start()

		err := a.packageManager.UpdatePackage(tool.Package)
		s.Stop()

		if err != nil {
			fmt.Println(color.RedString("✗ " + i18n.T("update.failed", tool.Name, err.Error())))
		} else {
			fmt.Println(color.GreenString("✓ " + i18n.T("update.success", tool.Name)))

			// Update cache with new version
			if versionInfo := a.versionCache[tool.Package]; versionInfo != nil {
				versionInfo.CurrentVersion = versionInfo.LatestVersion
			}
		}
	}
}

// handleUninstall handles the uninstall action
func (a *App) handleUninstall() {
	if len(a.installedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("uninstall.noneInstalled")))
		return
	}

	// Create choices
	items := make([]string, len(a.installedTools))
	for i, tool := range a.installedTools {
		items[i] = fmt.Sprintf("%s (%s)", tool.Name, tool.Package)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▸ {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: "▸ {{ . | cyan }}",
	}

	prompt := promptui.Select{
		Label:     i18n.T("uninstall.selectToUninstall") + " " + i18n.T("prompts.useArrowKeys"),
		Items:     append(items, i18n.T("menu.exit")),
		Templates: templates,
		Size:      10,
	}

	var selectedTools []config.Tool

	for {
		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		if index == len(items) { // Exit option
			break
		}

		selectedTools = append(selectedTools, a.installedTools[index])

		confirmPrompt := promptui.Select{
			Label: "Select more tools?",
			Items: []string{"Yes", "No, proceed to confirm"},
		}

		choice, _, _ := confirmPrompt.Run()
		if choice == 1 {
			break
		}
	}

	if len(selectedTools) == 0 {
		fmt.Println(color.YellowString(i18n.T("uninstall.noneSelected")))
		return
	}

	// Confirm uninstall
	confirmPrompt := promptui.Select{
		Label: fmt.Sprintf("%s %s", i18n.T("uninstall.confirm", len(selectedTools)), i18n.T("prompts.confirm")),
		Items: []string{"No", "Yes"},
	}

	choice, _, err := confirmPrompt.Run()
	if err != nil || choice == 0 {
		fmt.Println(color.YellowString(i18n.T("uninstall.cancelled")))
		return
	}

	// Uninstall selected tools
	for _, tool := range selectedTools {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " " + i18n.T("uninstall.uninstalling", tool.Name)
		s.Start()

		err := a.packageManager.UninstallPackage(tool.Package)
		s.Stop()

		if err != nil {
			fmt.Println(color.RedString("✗ " + i18n.T("uninstall.failed", tool.Name, err.Error())))
		} else {
			fmt.Println(color.GreenString("✓ " + i18n.T("uninstall.success", tool.Name)))

			// Update cache lists
			a.uninstalledTools = append(a.uninstalledTools, tool)
			a.removeFromInstalled(tool.Package)
			delete(a.versionCache, tool.Package)
		}
	}
}

// removeFromInstalled removes a tool from the installed list
func (a *App) removeFromInstalled(packageName string) {
	for i, tool := range a.installedTools {
		if tool.Package == packageName {
			a.installedTools = append(a.installedTools[:i], a.installedTools[i+1:]...)
			break
		}
	}
}

// removeFromUninstalled removes a tool from the uninstalled list
func (a *App) removeFromUninstalled(packageName string) {
	for i, tool := range a.uninstalledTools {
		if tool.Package == packageName {
			a.uninstalledTools = append(a.uninstalledTools[:i], a.uninstalledTools[i+1:]...)
			break
		}
	}
}
