package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Config represents the structure of our TOML configuration file.
type Config struct {
	Mode      string `toml:"mode"`
	Model     string `toml:"model"`
	OllamaURL string `toml:"ollama_url"`
}

// Message defines the message structure for request and response.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequestBody defines the structure for the request to Ollama.
type ChatRequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

// ChatResponseBody defines the full structure for the chat response.
type ChatResponseBody struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
	// Include additional fields if needed.
}

// SendOllamaRequest sends a POST to the Ollama server.
func SendOllamaRequest(reqBody ChatRequestBody, cfg Config) (ChatResponseBody, error) {
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return ChatResponseBody{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	url := fmt.Sprintf("%s/api/chat", cfg.OllamaURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return ChatResponseBody{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ChatResponseBody{}, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResponseBody{}, fmt.Errorf("error reading response body: %w", err)
	}

	var resBody ChatResponseBody
	if err := json.Unmarshal(responseBytes, &resBody); err != nil {
		return ChatResponseBody{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resBody, nil
}
