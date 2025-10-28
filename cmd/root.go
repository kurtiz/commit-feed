/*
Copyright © 2025
Aaron Will Djaba <aaronwilldjaba@outlook.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without subcommands
var rootCmd = &cobra.Command{
	Use:   "commitfeed",
	Short: "Generate AI-powered social media posts from your Git commits.",
	Long: `CommitFeed is a developer tool that reads your Git commit history
and uses AI to automatically generate engaging, platform-ready posts for
LinkedIn, Twitter/X, and other social platforms.

💡 Key features:
  • Reads real Git commits and summarizes your recent work.
  • Uses AI models (Hugging Face, OpenAI, Gemini, etc.) to craft posts.
  • Generates optimized content for LinkedIn and Twitter by default.
  • Lets you preview or publish posts directly from your terminal.
  • Stores API keys securely in ~/.commit-feed/config.json.

Examples:
  # Generate posts for your latest commits
  commitfeed generate

  # Generate posts only for Twitter
  commitfeed generate --platforms=twitter

  # Generate posts for the last 5 commits
  commitfeed generate --range HEAD~5..HEAD

  # Automatically publish generated posts (coming soon)
  commitfeed generate --post
`,
}

// Execute adds all child commands to the root command and handles flag setup
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags and configuration can be added here
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
