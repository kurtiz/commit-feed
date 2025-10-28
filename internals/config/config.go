package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config defines user settings stored in ~/.commit-feed/config.json
type Config struct {
	Provider         string   `json:"provider"`
	APIKey           string   `json:"api_key"`
	DefaultPlatforms []string `json:"default_platforms"`
}

// Path returns the full config file path (~/.commit-feed/config.json)
func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get home dir: %w", err)
	}
	return filepath.Join(home, ".commit-feed", "config.json"), nil
}

// defaultConfig returns a basic default setup (used if user skips setup)
func defaultConfig() *Config {
	return &Config{
		Provider:         "huggingface",
		APIKey:           "",
		DefaultPlatforms: []string{"linkedin", "twitter"},
	}
}

// Load reads the config file and merges environment overrides
func Load() (*Config, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config format: %v", err)
	}

	// Allow environment variable overrides
	if v := os.Getenv("COMMITFEED_PROVIDER"); v != "" {
		cfg.Provider = v
	}
	if v := os.Getenv("COMMITFEED_API_KEY"); v != "" {
		cfg.APIKey = v
	}
	return &cfg, nil
}

// EnsureExists loads config if present, otherwise runs the setup wizard
func EnsureExists() (*Config, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("🧩 No config found — launching first-time setup...\n")
		return RunSetupWizard()
	}
	return Load()
}
