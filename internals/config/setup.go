package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00C896")).
		Padding(1, 0)
)

// RunSetupWizard launches an interactive TUI to create the config file.
func RunSetupWizard() (*Config, error) {
	fmt.Println(titleStyle.Render("🚀 Welcome to CommitFeed!"))
	fmt.Println("Let's set up your AI provider to generate social posts from your git commits.")

	providerOptions := []string{"huggingface (free default)", "gemini", "openai", "deepseek"}
	var providerChoice string
	var apiKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose an AI provider").
				Options(huh.NewOptions(providerOptions...)...).
				Value(&providerChoice),

			huh.NewInput().
				Title("Enter your API key (press Enter to skip for default)").
				Prompt("> ").
				Value(&apiKey),
		),
	)

	if err := form.Run(); err != nil {
		return nil, fmt.Errorf("setup cancelled: %v", err)
	}

	// Normalize provider string
	provider := "huggingface"
	switch {
	case providerChoice == "gemini":
		provider = "gemini"
	case providerChoice == "openai":
		provider = "openai"
	case providerChoice == "deepseek":
		provider = "deepseek"
	}

	// If non-default, show info on where to get the API key
	if provider != "huggingface" && apiKey == "" {
		fmt.Printf("\n⚠️  You chose %s but didn’t provide an API key.\n", provider)
		switch provider {
		case "gemini":
			fmt.Println("👉 Get your free key at: https://aistudio.google.com/app/apikey")
		case "openai":
			fmt.Println("👉 Get your key at: https://platform.openai.com/account/api-keys")
		case "deepseek":
			fmt.Println("👉 Get your key at: https://platform.deepseek.com/")
		}
		fmt.Println("You can edit it later at ~/.commit-feed/config.json")
	}

	cfg := &Config{
		Provider:         provider,
		APIKey:           apiKey,
		DefaultPlatforms: []string{"linkedin", "twitter"},
	}

	if err := saveConfig(cfg); err != nil {
		return nil, err
	}

	fmt.Println("\n✅ Configuration saved successfully!")
	fmt.Printf("📄 Location: %s\n\n", configPath())
	fmt.Println("Now you can run `commitfeed generate` to create your first social post.")
	return cfg, nil
}

func saveConfig(cfg *Config) error {
	path := configPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, data, 0o644)
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".commit-feed", "config.json")
}
