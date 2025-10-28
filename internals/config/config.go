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

// Save writes a config object to disk
func Save(cfg *Config) error {
	path, err := Path()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("failed to write config: %v", err)
	}
	return nil
}

// Load reads the config file and merges environment overrides
func Load() (*Config, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		// If file doesn’t exist, fall back to default config
		return defaultConfig(), nil
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("⚠️ Invalid config format, falling back to defaults.")
		return defaultConfig(), nil
	}

	// Allow environment variable overrides
	if v := os.Getenv("COMMITFEED_PROVIDER"); v != "" {
		cfg.Provider = v
	}
	if v := os.Getenv("COMMITFEED_API_KEY"); v != "" {
		cfg.APIKey = v
	}

	// Default platforms fallback (in case file was missing field)
	if len(cfg.DefaultPlatforms) == 0 {
		cfg.DefaultPlatforms = []string{"linkedin", "twitter"}
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
		fmt.Println("🧩 No config found — launching first-time setup...")
		cfg, err := RunSetupWizard()
		if err != nil {
			fmt.Println("⚠️ Setup canceled or failed — using default configuration.")
			cfg = defaultConfig()
			if saveErr := Save(cfg); saveErr != nil {
				return nil, fmt.Errorf("failed to save default config: %v", saveErr)
			}
		}
		return cfg, nil
	}
	return Load()
}
