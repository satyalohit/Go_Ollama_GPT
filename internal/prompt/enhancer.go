package prompt

import (
	"fmt"
	"strings"

	"Go_Ollema_GPT/internal/search"
)

func EnhanceWithSearch(prompt string) string {
	results, err := search.SearchDuckDuckGo(prompt)
	if err != nil {
		fmt.Printf("Search error: %v\n", err)
		return prompt
	}

	var enhancedPrompt strings.Builder
	enhancedPrompt.WriteString("Based on the following information and your knowledge, please answer this question: ")
	enhancedPrompt.WriteString(prompt)
	enhancedPrompt.WriteString("\n\nRelevant information from search:\n")

	for i, result := range results {
		if i >= 3 {
			break
		}
		enhancedPrompt.WriteString(fmt.Sprintf("\n%s\n%s\nSource: %s\n", 
			result.Title, result.Content, result.URL))
	}

	return enhancedPrompt.String()
}