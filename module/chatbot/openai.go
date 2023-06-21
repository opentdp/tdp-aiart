package chatbot

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

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

func OpenaiModels() ([]string, error) {

	client, err := OpenaiClient()
	if err != nil {
		return nil, err
	}

	// 调用 OpenAI 接口

	resp, err := client.ListModels(
		context.Background(),
	)

	if err != nil {
		return nil, err
	}

	// 解析返回的数据

	models := []string{}

	for _, model := range resp.Models {
		if strings.Contains(model.ID, "gpt") {
			models = append(models, model.ID)
		}
	}

	return models, nil

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
		Message: ChatMessage{
			Role:    resp.Choices[0].Message.Role,
			Content: resp.Choices[0].Message.Content,
		},
	}

	return &data, nil

}

func OpenaiChatStream(rq *ReqeustParam) (chan ChatResponse, error) {

	client, err := OpenaiClient()
	if err != nil {
		return nil, err
	}

	// 构造请求参数

	req := openai.ChatCompletionRequest{
		Model:    rq.Model,
		Messages: []openai.ChatCompletionMessage{},
		Stream:   true,
	}

	for _, msg := range rq.Messages {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 调用 OpenAI 接口

	stream, err := client.CreateChatCompletionStream(
		context.Background(), req,
	)

	if err != nil {
		return nil, err
	}

	// 解析返回的数据

	data := make(chan ChatResponse, 32*1024)

	go func() {
		defer stream.Close()
		defer close(data)
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				data <- ChatResponse{
					Reason: "<!finish>",
				}
				return
			}
			if err != nil {
				data <- ChatResponse{
					Reason: "<!error>",
					Message: ChatMessage{
						Content: err.Error(),
					},
				}
				return
			}
			if len(resp.Choices) == 0 {
				data <- ChatResponse{
					Reason: "<!finish>",
				}
				return
			}
			data <- ChatResponse{
				Index:  resp.Choices[0].Index,
				Reason: string(resp.Choices[0].FinishReason),
				Message: ChatMessage{
					Role:    "assistant",
					Content: resp.Choices[0].Delta.Content,
				},
			}
		}
	}()

	return data, nil

}
