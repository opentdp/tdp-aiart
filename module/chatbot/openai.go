package chatbot

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/sashabaranov/go-openai"

	"tdp-aiart/model/config"
)

func OpenaiClient() (*openai.Client, error) {

	kv := config.ValuesOf("openai")

	if kv["ApiKey"] == "" {
		return nil, errors.New("请先设置 OpenAI 密钥")
	}

	config := openai.DefaultConfig(kv["ApiKey"])

	if kv["ApiUrl"] != "" {
		config.BaseURL = kv["ApiUrl"]
	}

	if kv["ProxyUrl"] != "" {
		proxyUrl, err := url.Parse(kv["ProxyUrl"])
		if err != nil {
			return nil, err
		}
		config.HTTPClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	return openai.NewClientWithConfig(config), nil

}

func OpenaiChat(rq *ReqeustParam) (*ChatResponse, error) {

	client, err := OpenaiClient()
	if err != nil {
		return nil, err
	}

	// 构造请求参数

	req := openai.ChatCompletionRequest{
		Model:    rq.Model,
		Messages: []openai.ChatCompletionMessage{},
	}

	for _, msg := range rq.Messages {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 调用 OpenAI 接口

	resp, err := client.CreateChatCompletion(
		context.Background(), req,
	)

	if err != nil {
		return nil, err
	}

	// 解析返回的数据

	data := ChatResponse{
		Messages: []ChatMessage{},
	}

	for _, choice := range resp.Choices {
		data.Messages = append(data.Messages, ChatMessage{
			Role:    choice.Message.Role,
			Content: choice.Message.Content,
		})
	}

	return &data, nil

}
