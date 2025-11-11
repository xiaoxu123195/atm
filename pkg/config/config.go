package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed tools.json
var configFile []byte

// Tool represents an AI development tool configuration
type Tool struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	Description string `json:"description"`
}

// Config represents the application configuration
type Config struct {
	Tools []Tool `json:"tools"`
}

// Load loads the configuration from the embedded JSON file
func Load() (*Config, error) {
	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
