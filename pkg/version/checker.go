package version

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Checker handles version checking for ATM itself
type Checker struct {
	currentVersion string
	repositoryURL  string
}

// UpdateInfo contains information about available updates
type UpdateInfo struct {
	HasUpdate      bool
	CurrentVersion string
	LatestVersion  string
	RepositoryURL  string
	Error          error
}

// GitHubRelease represents a GitHub release API response
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

// NewChecker creates a new version checker
func NewChecker(currentVersion, repositoryURL string) *Checker {
	return &Checker{
		currentVersion: currentVersion,
		repositoryURL:  repositoryURL,
	}
}

// CheckForUpdates checks if a newer version is available on GitHub
func (c *Checker) CheckForUpdates() *UpdateInfo {
	info := &UpdateInfo{
		CurrentVersion: c.currentVersion,
		RepositoryURL:  c.repositoryURL,
	}

	// Extract owner and repo from URL
	// Expected format: https://github.com/owner/repo
	parts := strings.Split(strings.TrimPrefix(c.repositoryURL, "https://github.com/"), "/")
	if len(parts) < 2 {
		info.Error = fmt.Errorf("invalid repository URL")
		return info
	}

	owner, repo := parts[0], parts[1]
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		info.Error = err
		return info
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		info.Error = fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
		return info
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		info.Error = err
		return info
	}

	// Remove 'v' prefix if present
	latestVersion := strings.TrimPrefix(release.TagName, "v")
	info.LatestVersion = latestVersion

	// Compare versions
	if compareVersions(c.currentVersion, latestVersion) < 0 {
		info.HasUpdate = true
	}

	return info
}

// OpenRepository opens the repository URL in the default browser
func (c *Checker) OpenRepository() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", c.repositoryURL)
	case "darwin":
		cmd = exec.Command("open", c.repositoryURL)
	case "linux":
		cmd = exec.Command("xdg-open", c.repositoryURL)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}

// compareVersions compares two semantic version strings
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersions(v1, v2 string) int {
	// Remove 'v' prefix if present
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")

	// Split versions into parts
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// Compare each part
	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var p1, p2 int

		if i < len(parts1) {
			fmt.Sscanf(parts1[i], "%d", &p1)
		}
		if i < len(parts2) {
			fmt.Sscanf(parts2[i], "%d", &p2)
		}

		if p1 < p2 {
			return -1
		} else if p1 > p2 {
			return 1
		}
	}

	return 0
}
