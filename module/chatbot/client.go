package chatbot

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatCompletionRequest = openai.ChatCompletionRequest
type ChatCompletionResponse = openai.ChatCompletionResponse

func Create(rq ChatCompletionRequest) (*ChatCompletionResponse, error) {

	client, err := OpenaiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(), rq,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}
