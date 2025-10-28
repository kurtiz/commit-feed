package ai

import (
	"fmt"
	"strings"

	"github.com/kurtiz/commit-feed.git/internals/git"
)

// buildPrompt formats commits into a high-quality AI prompt
func buildPrompt(commits []git.Commit) string {
	var sb strings.Builder

	sb.WriteString(`You are a skilled technical copywriter who creates engaging posts for developers and tech audiences.
Your task is to write one short LinkedIn post and one short Twitter/X post based on the following commit messages.

Each commit message represents a meaningful code change or feature update.

Write content that:
- Clearly explains the impact or purpose of the updates.
- Uses a friendly and professional tone.
- Avoids too much technical jargon unless it's relevant.
- Keeps LinkedIn posts engaging (2-4 sentences).
- Keeps Twitter/X posts concise and catchy (max 280 characters, use emojis or hashtags if natural).

Here are the commits:
`)

	for _, c := range commits {
		sb.WriteString(fmt.Sprintf("- %s\n", c.Message))
	}

	sb.WriteString(`
Format your response exactly like this:
LinkedIn: <your LinkedIn post>
Twitter: <your Twitter post>

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
