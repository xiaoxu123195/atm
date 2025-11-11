package i18n

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

var (
	currentLang string
	messages    map[string]string
)

// Init initializes the i18n system by detecting the system language
func Init() {
	currentLang = detectLanguage()
	loadMessages()
}

// T returns the translated string for the given key with optional formatting
func T(key string, args ...interface{}) string {
	if msg, ok := messages[key]; ok {
		if len(args) > 0 {
			return fmt.Sprintf(msg, args...)
		}
		return msg
	}
	// Fallback to key if translation not found
	return key
}

// GetLanguage returns the current language code
func GetLanguage() string {
	return currentLang
}

// detectLanguage detects the system language from environment variables
func detectLanguage() string {
	// Check environment variables in order
	envVars := []string{"LANG", "LC_ALL", "LC_MESSAGES"}

	for _, envVar := range envVars {
		if lang := os.Getenv(envVar); lang != "" {
			if strings.HasPrefix(strings.ToLower(lang), "zh") {
				return "zh"
			}
			return "en"
		}
	}

	// Windows-specific: Get system default locale
	if runtime.GOOS == "windows" {
		if lang := getWindowsLanguage(); lang != "" {
			return lang
		}
	}

	// Default to English
	return "en"
}

// getWindowsLanguage gets the system language on Windows
func getWindowsLanguage() string {
	if runtime.GOOS != "windows" {
		return ""
	}

	// Try to get Windows locale
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getUserDefaultLocaleName := kernel32.NewProc("GetUserDefaultLocaleName")

	buffer := make([]uint16, 85) // LOCALE_NAME_MAX_LENGTH
	ret, _, _ := getUserDefaultLocaleName.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
	)

	if ret > 0 {
		localeName := syscall.UTF16ToString(buffer)
		if strings.HasPrefix(strings.ToLower(localeName), "zh") {
			return "zh"
		}
	}

	return "en"
}

// loadMessages loads the appropriate message map based on current language
func loadMessages() {
	switch currentLang {
	case "zh":
		messages = zhMessages
	default:
		messages = enMessages
	}
}

// SetLanguage manually sets the language (useful for testing)
func SetLanguage(lang string) {
	currentLang = lang
	loadMessages()
}
