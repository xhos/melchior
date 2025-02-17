package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"melchior/internal/ollama"
	"os"
	"strings"
)

func main() {
	// Ensure that some input is provided.
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <text>")
		return
	}

	// Combine command-line arguments to form the input.
	inputText := strings.Join(os.Args[1:], " ")

	// Wrap the user input with the revised prompt.
	prompt := `You are Melchior. Your only purpose is to translate English instructions into a single command line.
RULES:
- Output ONLY the command, nothing else
- Do not explain the command
- Do not add comments or suggestions
- Return exactly one line
- No markdown, no formatting

CRITICAL: Return ONLY the command. ANY additional text, explanation, or notes will result in failure.
CRITICAL: When the user input does not imply a command, return "I cannot help with that."
Example 1: "list directory contents"
ls

Example 2: "make directory called test"
mkdir test

Query:
`
	wrappedInputText := prompt + inputText

	// Load configuration from config.toml.
	var cfg ollama.Config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		fmt.Printf("Error loading config.toml (%v). Using defaults.\n", err)
		cfg = ollama.Config{
			Mode:      "ollama",
			Model:     "mistral",
			OllamaURL: "http://localhost:11434",
		}
	}

	// Create the JSON payload using the correct structure.
	reqBody := ollama.ChatRequestBody{
		Model: cfg.Model,
		Messages: []ollama.Message{
			{Role: "user", Content: wrappedInputText},
		},
		Stream: false,
	}

	switch cfg.Mode {
	case "ollama":
		resBody, err := ollama.SendOllamaRequest(reqBody, cfg)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Output the single command returned by the assistant.
		fmt.Println(resBody.Message.Content)
	default:
		fmt.Printf("Unsupported mode: %s\n", cfg.Mode)
	}
}
