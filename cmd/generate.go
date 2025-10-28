/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate social media posts from recent Git commits.",
	Long: `Generate AI-powered social media posts from your recent Git commits.

CommitFeed scans your Git history, summarizes recent changes, and uses AI to craft
platform-ready posts for LinkedIn and Twitter (by default). You can specify which
platforms to generate for, adjust the commit range, or preview the output before posting.

Examples:
  # Generate posts for the latest commits (LinkedIn + Twitter)
  commitfeed generate

  # Generate posts for the last 5 commits
  commitfeed generate --range HEAD~5..HEAD

  # Generate only a Twitter post (dry run, no posting)
  commitfeed generate --platforms=twitter --dry-run`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
