package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"Go_Ollema_GPT/internal/prompt"
)

func NewOllamaClient(model string) *OllamaClient {
	return &OllamaClient{
		model: model,
	}
}

func (c *OllamaClient) SendPrompt(userPrompt string) (string, error) {
	enhancedPrompt := prompt.EnhanceWithSearch(userPrompt)

	reqBody := OllamaRequest{
		Model:  c.model,
		Prompt: enhancedPrompt,
		Stream: true,
		System: "You are a helpful AI assistant. When providing information, cite sources when available.",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		var response OllamaResponse
		if err := json.Unmarshal(scanner.Bytes(), &response); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %v", err)
		}
		fullResponse.WriteString(response.Response)
	}

	return fullResponse.String(), nil
}