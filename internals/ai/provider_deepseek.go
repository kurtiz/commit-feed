package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kurtiz/commit-feed.git/internals/git"
)

type DeepSeekProvider struct {
	apiKey string
}

func NewDeepSeekProvider(apiKey string) *DeepSeekProvider {
	if apiKey == "" {
		apiKey = os.Getenv("DEEPSEEK_API_KEY")
	}
	return &DeepSeekProvider{apiKey: apiKey}
}

func (d *DeepSeekProvider) GeneratePosts(commits []git.Commit) (*GeneratedPosts, error) {
	body := map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{"role": "system", "content": "You are CommitFeed, summarizing commits into social posts."},
			{"role": "user", "content": buildPrompt(commits)},
		},
	}
	data, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+d.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("deepseek request error: %v", err)
	}
	defer resp.Body.Close()

	var parsed struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	json.NewDecoder(resp.Body).Decode(&parsed)

	if len(parsed.Choices) == 0 {
		return nil, fmt.Errorf("no response from deepseek")
	}

	return parseResponse(parsed.Choices[0].Message.Content), nil
}
