package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Go_Ollema_GPT/internal/api"
	"Go_Ollema_GPT/config"
)

func main() {
	fmt.Println("Welcome to Enhanced Ollama Chat with Web Search!")
	fmt.Println("Type 'exit' to quit the chat")
	fmt.Print("Enter model name (default: llama2): ")

	reader := bufio.NewReader(os.Stdin)
	modelName, _ := reader.ReadString('\n')
	modelName = strings.TrimSpace(modelName)

	if modelName == "" {
		modelName = config.DefaultModel
	}

	fmt.Printf("Using model: %s\n", modelName)
	fmt.Println("You can start chatting now!")

	client := api.NewOllamaClient(modelName)

	for {
		fmt.Print("\nYou: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println("\nSearching for relevant information...")
		response, err := client.SendPrompt(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("\nAssistant: %s\n", response)
	}
}