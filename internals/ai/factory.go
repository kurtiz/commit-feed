package ai

import (
	"fmt"
)

// NewProvider returns the appropriate AI provider implementation
func NewProvider(name string, apiKey string) (Provider, error) {
	switch name {
	case "openai":
		return NewOpenAIProvider(apiKey), nil
	case "gemini":
		return NewGeminiProvider(apiKey), nil
	case "deepseek":
		return NewDeepSeekProvider(apiKey), nil
	case "huggingface", "default", "":
		return NewHuggingFaceProvider(apiKey), nil
	default:
		return nil, fmt.Errorf("unknown AI provider: %s", name)
	}
}
