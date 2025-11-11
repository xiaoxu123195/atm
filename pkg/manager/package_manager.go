package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// PackageManager handles npm package operations
type PackageManager struct{}

// NewPackageManager creates a new PackageManager instance
func NewPackageManager() *PackageManager {
	return &PackageManager{}
}

// NpmListOutput represents the output of npm list command
type NpmListOutput struct {
	Dependencies map[string]struct {
		Version string `json:"version"`
	} `json:"dependencies"`
}

// IsPackageInstalled checks if a package is installed globally
func (pm *PackageManager) IsPackageInstalled(packageName string) (bool, error) {
	cleanName := pm.extractPackageName(packageName)
	cmd := exec.Command("npm", "list", "-g", cleanName, "--depth=0", "--json")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// npm list returns non-zero exit code if package not found, but that's ok
	_ = cmd.Run()

	var result NpmListOutput
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		// If JSON parsing fails, package is not installed
		return false, nil
	}

	_, exists := result.Dependencies[cleanName]
	return exists, nil
}

// GetPackageVersion gets the currently installed version of a package
func (pm *PackageManager) GetPackageVersion(packageName string) (string, error) {
	cleanName := pm.extractPackageName(packageName)
	cmd := exec.Command("npm", "list", "-g", cleanName, "--depth=0", "--json")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	_ = cmd.Run()

	var result NpmListOutput
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return "", fmt.Errorf("failed to parse npm output")
	}

	if dep, exists := result.Dependencies[cleanName]; exists {
		return dep.Version, nil
	}

	return "", fmt.Errorf("package not found")
}

// GetLatestVersion gets the latest available version from npm registry
func (pm *PackageManager) GetLatestVersion(packageName string) (string, error) {
	cleanName := pm.extractPackageName(packageName)
	cmd := exec.Command("npm", "view", cleanName, "version")

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get latest version: %s", errOut.String())
	}

	version := strings.TrimSpace(out.String())
	return version, nil
}

// InstallPackage installs a package globally
func (pm *PackageManager) InstallPackage(packageName string) error {
	cmd := exec.Command("npm", "install", "-g", packageName)

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installation failed: %s", errOut.String())
	}

	return nil
}

// UpdatePackage updates a package to the latest version
func (pm *PackageManager) UpdatePackage(packageName string) error {
	cmd := exec.Command("npm", "update", "-g", packageName)

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("update failed: %s", errOut.String())
	}

	return nil
}

// UninstallPackage removes a package from the system
func (pm *PackageManager) UninstallPackage(packageName string) error {
	cleanName := pm.extractPackageName(packageName)
	cmd := exec.Command("npm", "uninstall", "-g", cleanName)

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("uninstallation failed: %s", errOut.String())
	}

	return nil
}

// extractPackageName extracts the package name from a package string that may include version
// Handles scoped packages like @scope/package and @scope/package@version
func (pm *PackageManager) extractPackageName(packageWithVersion string) string {
	// Handle scoped packages (starting with @)
	if strings.HasPrefix(packageWithVersion, "@") {
		parts := strings.Split(packageWithVersion, "@")
		if len(parts) == 2 {
			// @scope/package (no version)
			return packageWithVersion
		} else if len(parts) >= 3 {
			// @scope/package@version
			return "@" + parts[1]
		}
	} else {
		// Regular packages: package or package@version
		parts := strings.Split(packageWithVersion, "@")
		return parts[0]
	}

	return packageWithVersion
}
