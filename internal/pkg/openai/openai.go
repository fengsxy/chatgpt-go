package openai

import (
	"context"
	"github.com/lyleshaw/chatgpt-go/pkg/utils/log"
	"github.com/otiai10/openaigo"
	"os"
)

func Chat(message []openaigo.ChatMessage) (string, error) {
	client := openaigo.NewClient(os.Getenv("OPENAI_API_KEY"))
	log.Infof("message: %v", message)
	request := openaigo.ChatCompletionRequestBody{
		Model:       "gpt-3.5-turbo",
		Messages:    message,
		Temperature: 0.7,
		TopP:        1,
		N:           1,
		User:        "user",
	}
	ctx := context.Background()
	response, err := client.Chat(ctx, request)
	if err != nil {
		log.Errorf("error: %v", err)
		return "", err
	}
	log.Infof("response: %v", response)
	return response.Choices[0].Message.Content, err
}
