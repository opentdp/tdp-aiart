package openai

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/sashabaranov/go-openai"

	"tdp-aiart/model/config"
)

func Client() (*openai.Client, error) {

	kv := config.ValuesOf("openai")

	if kv["ApiKey"] == "" {
		return nil, errors.New("请先设置 OpenAI 密钥")
	}

	config := openai.DefaultConfig(kv["ApiKey"])

	if kv["proxyUrl"] == "" {
		proxyUrl, err := url.Parse(kv["proxyUrl"])
		if err != nil {
			return nil, err
		}
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}

	return openai.NewClientWithConfig(config), nil

}

func Chat(rq *[]ChatCompletionMessage) (*ChatCompletionMessage, error) {

	client, err := Client()
	if err != nil {
		return nil, err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: *rq,
		},
	)
	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message, nil

}
