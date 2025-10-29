package ai

import (
	"context"
	"fmt"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"github.com/kurtiz/commit-feed/internals/git"
	"google.golang.org/api/option"
)

type GeminiProvider struct {
	client *genai.Client
}

func NewGeminiProvider(apiKey string) *GeminiProvider {
	if apiKey == "" {
		apiKey = os.Getenv("GEMINI_API_KEY")
	}
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	return &GeminiProvider{client: client}
}

func (g *GeminiProvider) GeneratePosts(commits []git.Commit, platforms []string, projectContext string) (*GeneratedPosts, error) {
	prompt := buildPrompt(commits, platforms, projectContext)

	resp, err := g.client.GenerativeModel("gemini-1.5-flash").
		GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini error: %v", err)
	}

	output := resp.Candidates[0].Content.Parts[0].(genai.Text)
	return parseResponse(string(output)), nil
}
