package artman

import (
	"errors"

	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/config"

	"github.com/mitchellh/mapstructure"
)

func TencentAiart(rq *ReqeustParams) (*TencentAiartResponse, error) {

	// 获取密钥

	kv := config.ValuesOf("Tencent")

	if kv["SecretId"] == "" || kv["SecretKey"] == "" {
		return nil, errors.New("请先配置腾讯云密钥")
	}

	// 发起请求

	param := &tencent.ReqeustParam{
		Service:   "aiart",
		Region:    "ap-guangzhou",
		Version:   "2022-12-29",
		SecretId:  kv["SecretId"],
		SecretKey: kv["SecretKey"],
		Action:    rq.Action,
		Payload:   rq.Payload,
	}

	resp, err := tencent.Request(param)

	// 解析参数

	output := &TencentAiartResponse{}

	if err == nil {
		err = mapstructure.Decode(resp, &output)
	}

	return output, err

}

type TencentAiartResponse struct {
	RequestId   string
	ResultImage string
}
