package chatbot

import (
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
