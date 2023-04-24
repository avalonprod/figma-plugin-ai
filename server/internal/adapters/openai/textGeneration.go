package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type textGenerator struct {
	apiKey string
	apiUrl string
}

type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index         int     `json:"index"`
	MessageObject Message `json:"message"`
	FinishReason  string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type textGenirationInput struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func NewTextGenerator(apiKey string, apiUrl string) *textGenerator {
	return &textGenerator{
		apiKey: apiKey,
		apiUrl: apiUrl,
	}
}

func (g *textGenerator) NewMessages(message string) (string, error) {
	// router := gin.Default()
	client := &http.Client{}

	input := textGenirationInput{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			Message{
				Role:    "user",
				Content: message,
			},
		},
	}

	inputBytes, err := json.Marshal(input)
	req, err := http.NewRequest("POST", g.apiUrl, bytes.NewBuffer(inputBytes))

	if err != nil {
		return "", errors.New("error creating request.")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.apiKey))
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	var output ChatCompletionResponse
	json.Unmarshal(body, &output)
	content := output.Choices[0].MessageObject.Content
	return content, nil
}
