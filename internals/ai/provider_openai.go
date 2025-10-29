package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/kurtiz/commit-feed/internals/git"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
	client *openai.Client
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	client := openai.NewClient(apiKey)
	return &OpenAIProvider{client: client}
}

// GeneratePosts uses OpenAI to generate platform-specific posts.
func (p *OpenAIProvider) GeneratePosts(commits []git.Commit, platforms []string, projectContext string) (*GeneratedPosts, error) {
	prompt := buildPrompt(commits, platforms, projectContext)

	resp, err := p.client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are CommitFeed, a social media assistant for developers."},
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("openai error: %v", err)
	}

	content := resp.Choices[0].Message.Content
	return parseResponse(content), nil
}
