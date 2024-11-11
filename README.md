# 🤖 Enhanced Ollama GPT Chat with Web Search

A powerful Go-based chat application that combines Ollama's AI capabilities with real-time web search to provide informed and up-to-date responses.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)


## ✨ Features

- 🔍 Real-time web search integration with DuckDuckGo
- 🤖 AI-powered responses using Ollama
- 📚 Source citations for information verification
- 🌊 Stream-based response handling
- 💻 User-friendly command-line interface
- 🔄 Real-time information updates

## 🏗️ Project Structure
├── cmd/
│ └── main.go # Application entry point
├── internal/
│ ├── api/ # Ollama API interaction
│ │ ├── ollama.go
│ │ └── types.go
│ ├── search/ # Web search functionality
│ │ └── duckduckgo.go
│ └── prompt/ # Prompt enhancement logic
│ └── enhancer.go
├── config/
│ └── config.go # Configuration constants
└── README.md


## 🚀 Quick Start

### Prerequisites

1. Go 1.21 or later
2. Ollama installed
3. Internet connection

### Installation

1. **Clone the repository**

git clone https://github.com/yourusername/Go_Ollema_GPT.git


2. **Install dependencies**

go mod tidy

3. **Install Ollama**

brew install ollama  # For macOS

ollama serve

5. **Pull the default model**

ollama pull llama2

### Running the Application

go run cmd/main.go

## 💡 Usage

1. Start the application
2. Enter your preferred model (default: llama2)
3. Type your questions
4. Get AI responses with real-time web search results
5. Type 'exit' to quit

Example interaction:
Welcome to Enhanced Ollama Chat with Web Search!
Enter model name (default: llama2): 

You: What are the latest developments in AI?
Searching for relevant information...
