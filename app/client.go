package main

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type ChatGptClient struct {
	client      *openai.Client
	model       string
	n           int
	temperature float32
}

type ChatGptClientConfiguration struct {
	authToken   string
	model       string
	n           int
	temperature float32
}

func (c *ChatGptClient) send(messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {

	return c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       c.model,
			N:           c.n,
			Temperature: c.temperature,
			Messages:    messages,
		},
	)
}

func (c *ChatGptClient) convert(messages ...string) []openai.ChatCompletionMessage {

	converted := make([]openai.ChatCompletionMessage, len(messages), len(messages))

	for i, message := range messages {
		converted[i] = openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		}
	}

	return converted
}

func GetClient(configuration *ChatGptClientConfiguration) *ChatGptClient {
	return &ChatGptClient{
		openai.NewClient(configuration.authToken),
		configuration.model,
		configuration.n,
		configuration.temperature,
	}
}
