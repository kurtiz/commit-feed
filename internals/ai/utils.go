package ai

import (
	"fmt"
	"strings"

	"github.com/kurtiz/commit-feed.git/internals/git"
)

// buildPrompt creates an AI prompt customized for the target platforms.
func buildPrompt(commits []git.Commit, platforms []string) string {
	var sb strings.Builder

	sb.WriteString(`You are a skilled technical copywriter who creates engaging, platform-appropriate posts for developers and tech audiences.

Your task is to generate short, high-quality social media posts based on the following Git commit messages.
Each commit represents a meaningful code change, bug fix, or feature update.

--- Commit Messages ---
`)

	for _, c := range commits {
		sb.WriteString(fmt.Sprintf("- %s\n", c.Message))
	}

	sb.WriteString("\n--- Platform Guidelines ---\n")

	for _, platform := range platforms {
		switch strings.ToLower(platform) {
		case "linkedin":
			sb.WriteString(`• LinkedIn: Write a friendly and professional summary (2-4 sentences). Explain what changed and why it matters to developers or users.
`)
		case "twitter", "x":
			sb.WriteString(`• Twitter/X: Write a short, catchy summary under 280 characters. Include emojis or hashtags if natural.
`)
		case "mastodon":
			sb.WriteString(`• Mastodon: Write an open-source community-style update with clear tone and hashtags if relevant.
`)
		case "devto", "dev.to":
			sb.WriteString(`• Dev.to: Write a short blog teaser — 2-3 sentences that introduce the update and invite readers to learn more.
`)
		case "reddit":
			sb.WriteString(`• Reddit: Write a conversational summary that would fit in a /r/programming or /r/golang post, no emojis.
`)
		default:
			sb.WriteString(fmt.Sprintf("• %s: Write a concise summary highlighting the main purpose and value of the change.\n", platform))
		}
	}

	sb.WriteString(`

Format your response EXACTLY like this (one line per platform):
LinkedIn: <LinkedIn post>
Twitter: <Twitter post>
Mastodon: <Mastodon post>
... etc.

Be creative but accurate. Focus on clarity, developer value, and readability.
`)

	return sb.String()
}

// parseResponse extracts LinkedIn and Twitter text from LLM response
func parseResponse(text string) *GeneratedPosts {
	posts := &GeneratedPosts{}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.ToLower(line), "linkedin:") {
			posts.LinkedIn = strings.TrimSpace(strings.TrimPrefix(line, "LinkedIn:"))
		} else if strings.HasPrefix(strings.ToLower(line), "twitter:") {
			posts.Twitter = strings.TrimSpace(strings.TrimPrefix(line, "Twitter:"))
		}
	}
	if posts.LinkedIn == "" {
		posts.LinkedIn = text
	}
	if posts.Twitter == "" {
		posts.Twitter = text
	}
	return posts
}
