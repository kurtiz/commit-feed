/*
Copyright © 2025 Aaron Will Djaba <aaronwilldjaba@outlook.com>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kurtiz/commit-feed.git/internals/ai"
	"github.com/kurtiz/commit-feed.git/internals/config"
	"github.com/kurtiz/commit-feed.git/internals/git"
)

var (
	rangeFlag     string
	platformsFlag []string
	postFlag      bool // if true, actually post
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate social media posts from recent Git commits.",
	Long: `Generate AI-powered social media posts from your recent Git commits.

CommitFeed scans your Git history, summarizes recent changes, and uses AI to craft
platform-ready posts for LinkedIn and Twitter (by default). You can specify which
platforms to generate for, adjust the commit range, and optionally post them automatically.

Examples:
  # Generate posts for the latest commits (LinkedIn + Twitter)
  commitfeed generate

  # Generate posts for the last 5 commits
  commitfeed generate --range HEAD~5..HEAD

  # Generate and post to both platforms
  commitfeed generate --post

  # Generate and post only to Twitter
  commitfeed generate --platforms=twitter --post`,

	Run: func(cmd *cobra.Command, args []string) {
		// --- 1️⃣ Check Git prerequisites ---
		if !git.IsGitInstalled() {
			fmt.Println("❌ Git is not installed. Please install Git to use CommitFeed.")
			return
		}
		if !git.IsGitRepo() {
			fmt.Println("❌ The current directory is not a Git repository.")
			return
		}

		// --- 2️⃣ Load or create config file ---
		cfg, err := config.EnsureExists()
		if err != nil {
			fmt.Println("❌ Failed to load config:", err)
			os.Exit(1)
		}

		// --- 3️⃣ Determine which platforms to generate for ---
		targetPlatforms := cfg.DefaultPlatforms
		if len(platformsFlag) > 0 {
			targetPlatforms = platformsFlag
		}

		fmt.Printf("📦 Using AI Provider: %s\n", cfg.Provider)
		fmt.Printf("📰 Target Platforms: %v\n\n", targetPlatforms)

		// --- 4️⃣ Fetch commits from Git ---
		commits, err := git.GetCommits(rangeFlag, 0)
		if err != nil {
			fmt.Println("❌ Failed to read commits:", err)
			return
		}
		if len(commits) == 0 {
			fmt.Println("No commits found in the specified range.")
			return
		}

		// --- 5️⃣ Generate posts via AI provider ---
		provider, err := ai.NewProvider(cfg.Provider, cfg.APIKey)
		if err != nil {
			fmt.Println("❌ Error creating AI provider:", err)
			return
		}

		posts, err := provider.GeneratePosts(commits, targetPlatforms)
		if err != nil {
			fmt.Println("❌ Failed to generate posts:", err)
			return
		}

		// --- 6️⃣ Output results ---
		fmt.Println("✅ Generated Posts:")
		for _, p := range targetPlatforms {
			switch p {
			case "linkedin":
				fmt.Printf("🔗 LinkedIn:\n%s\n\n", posts.LinkedIn)
			case "twitter":
				fmt.Printf("🐦 Twitter:\n%s\n\n", posts.Twitter)
			default:
				fmt.Printf("📢 %s:\n%s\n\n", p, posts.LinkedIn)
			}
		}

		// --- 7️⃣ Handle posting ---
		if postFlag {
			fmt.Println("🚀 Posting to selected platforms...")

			// This is where you’ll later plug in your posting logic
			for _, p := range targetPlatforms {
				switch p {
				case "linkedin":
					fmt.Println("🔗 Posted to LinkedIn successfully (placeholder).")
				case "twitter":
					fmt.Println("🐦 Posted to Twitter successfully (placeholder).")
				default:
					fmt.Printf("📢 Skipped unknown platform: %s\n", p)
				}
			}
		} else {
			fmt.Println("💡 Preview only (not posted). Use --post to share automatically.")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&rangeFlag, "range", "r", "HEAD~2..HEAD", "Git commit range to summarize (e.g. HEAD~5..HEAD)")
	generateCmd.Flags().StringSliceVarP(&platformsFlag, "platforms", "t", nil, "Comma-separated list of platforms (e.g. linkedin,twitter)")
	generateCmd.Flags().BoolVarP(&postFlag, "post", "p", false, "Post generated content to selected platforms")
}
