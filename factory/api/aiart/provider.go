package aiart

import (
	"encoding/json"
	"errors"

	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/config"

	"github.com/mitchellh/mapstructure"
)

func apiProxy(rq *ReqeustParams) (*ToImageResponse, error) {

	// 获取密钥

	kv := config.ValuesOf("Tencent")

	if kv["secretId"] == "" || kv["secretKey"] == "" {
		return nil, errors.New("请先配置腾讯云密钥")
	}

	// 发起请求

	payload, _ := json.Marshal(rq.Payload)

	param := &tencent.ReqeustParam{
		Service:   "aiart",
		Region:    "ap-guangzhou",
		Version:   "2022-12-29",
		SecretId:  kv["secretId"],
		SecretKey: kv["secretKey"],
		Action:    rq.Action,
		Payload:   payload,
	}

	resp, err := tencent.Request(param)

	// 解析参数

	output := &ToImageResponse{}

	if err == nil {
		err = mapstructure.Decode(resp, &output)
	}

	return output, err

}

type ReqeustParams struct {
	Action  string
	Payload ToImageRequest
}

type ToImageRequest struct {
	InputImage     string
	InputUrl       string
	Prompt         string
	NegativePrompt string
	Styles         []string
	ResultConfig   any
	LogoAdd        int64
	Strength       float64
}

type ToImageResponse struct {
	RequestId   string
	ResultImage string
}
