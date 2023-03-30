package aiart

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/mitchellh/mapstructure"

	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/config"
)

func apiProxy(param *tencent.ReqeustParam) (any, error) {

	// 获取密钥

	kv := config.ValuesOf("Tencent")

	if kv["secretId"] == "" || kv["secretKey"] == "" {
		return nil, errors.New("请先配置腾讯云密钥")
	}

	// 发起请求

	param.Service = "aiart"
	param.Region = "ap-guangzhou"
	param.Version = "2022-12-29"
	param.SecretId = kv["secretId"]
	param.SecretKey = kv["secretKey"]

	return tencent.Request(param)

}

func saveImage(res any) error {

	data := struct {
		RequestId   string
		ResultImage string
	}{}

	if err := mapstructure.Decode(res, &data); err != nil {
		return err
	}

	if data.ResultImage != "" {
		filePath := "D:/test/test.jpg"
		return saveBase64ImageToFile(data.ResultImage, filePath)
	}

	return nil

}

func saveBase64ImageToFile(base64String string, filePath string) error {

	imageBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return err
	}

	return nil

}
