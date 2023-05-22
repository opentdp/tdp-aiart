package chatbot

import (
	"context"
	"errors"
	"io"
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
		Messages: []ChatMessage{
			{
				Role:    resp.Choices[0].Message.Role,
				Content: resp.Choices[0].Message.Content,
			},
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

	data := make(chan ChatResponse, 100)

	go func() {
		defer stream.Close()
		defer close(data)
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				data <- ChatResponse{
					FinishReason: "<!finish>",
				}
				return
			}
			if err != nil {
				data <- ChatResponse{
					FinishReason: "<!error>",
					Messages: []ChatMessage{
						{
							Content: err.Error(),
						},
					},
				}
				return
			}
			if len(resp.Choices) == 0 {
				data <- ChatResponse{
					FinishReason: "<!finish>",
				}
				return
			}
			data <- ChatResponse{
				Messages: []ChatMessage{
					{
						Index:   resp.Choices[0].Index,
						Content: resp.Choices[0].Delta.Content,
					},
				},
				FinishReason: resp.Choices[0].FinishReason,
			}
		}
	}()

	return data, nil

}
