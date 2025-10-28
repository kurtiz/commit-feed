package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kurtiz/commit-feed.git/internals/git"
)

// HuggingFaceProvider represents the Hugging Face API client
type HuggingFaceProvider struct {
	apiKey string
	model  string
}

// NewHuggingFaceProvider creates a new provider (default model can be overridden later)
func NewHuggingFaceProvider(apiKey string) *HuggingFaceProvider {
	return &HuggingFaceProvider{
		apiKey: apiKey,
		model:  "openai/gpt-oss-20b:groq", // default chat-capable model
	}
}

// GeneratePosts builds a prompt and requests AI-generated posts
func (h *HuggingFaceProvider) GeneratePosts(commits []git.Commit) (*GeneratedPosts, error) {
	prompt := buildPrompt(commits)

	// Build OpenAI-style chat request
	payload := map[string]interface{}{
		"model": h.model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"stream": false,
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://router.huggingface.co/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if h.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+h.apiKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("huggingface request failed: %v", err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("huggingface error: %s", string(data))
	}

	// Parse Hugging Face chat completion response
	var parsed struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse huggingface response: %v", err)
	}

	if len(parsed.Choices) == 0 || parsed.Choices[0].Message.Content == "" {
		return nil, fmt.Errorf("no response content returned from model")
	}

	// Extract clean message text
	text := parsed.Choices[0].Message.Content

	// Use existing parser to split LinkedIn/Twitter sections
	return parseResponse(text), nil
}
