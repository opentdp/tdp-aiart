package painter

import (
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/opentdp/go-helper/tencent"

	"tdp-aiart/model/config"
)

func TencentAiart(rq *ReqeustParam) (string, error) {

	// 获取密钥

	kv := config.ValuesOf("tencent")

	if kv["SecretId"] == "" || kv["SecretKey"] == "" {
		return "", errors.New("请先配置腾讯云密钥")
	}

	// 构造参数

	var payload []byte

	if rq.InputImage == "" {
		payload, _ = json.Marshal(TecnetTextPayload{
			Prompt:         rq.Prompt,
			NegativePrompt: rq.NegativePrompt,
			Styles:         rq.Styles,
			ResultConfig:   rq.ResultConfig,
			LogoAdd:        rq.LogoAdd,
			LogoParam:      rq.LogoParam,
		})
	} else {
		payload, _ = json.Marshal(TecnetImagePayload{
			Prompt:         rq.Prompt,
			NegativePrompt: rq.NegativePrompt,
			InputImage:     rq.InputImage,
			Strength:       rq.Strength,
			Styles:         rq.Styles,
			ResultConfig:   rq.ResultConfig,
			LogoAdd:        rq.LogoAdd,
			LogoParam:      rq.LogoParam,
		})
	}

	// 发起请求

	param := &tencent.ReqeustParam{
		Service:   "aiart",
		Region:    "ap-guangzhou",
		Version:   "2022-12-29",
		SecretId:  kv["SecretId"],
		SecretKey: kv["SecretKey"],
		Action:    rq.Action,
		Payload:   payload,
	}

	resp, err := tencent.Request(param)

	if err != nil {
		return "", err
	}

	// 解析结果

	result := &TencentResponse{}
	err = mapstructure.Decode(resp, &result)

	if err != nil {
		return "", err
	}

	// 返回图片

	return result.ResultImage, err

}

type TecnetImagePayload struct {
	Prompt         string `json:",omitempty"`
	NegativePrompt string `json:",omitempty"`
	InputImage     string
	Strength       float64        `json:",omitempty"`
	Styles         []string       `json:",omitempty"`
	ResultConfig   map[string]any `json:",omitempty"`
	LogoAdd        int64
	LogoParam      map[string]any `json:",omitempty"`
}

type TecnetTextPayload struct {
	Prompt         string
	NegativePrompt string
	Styles         []string       `json:",omitempty"`
	ResultConfig   map[string]any `json:",omitempty"`
	LogoAdd        int64
	LogoParam      map[string]any `json:",omitempty"`
}

type TencentResponse struct {
	RequestId   string
	ResultImage string
}
