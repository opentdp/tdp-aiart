package artman

import (
	"encoding/base64"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"

	"tdp-aiart/cmd/args"
	"tdp-aiart/module/model/artimg"
)

func SaveObject(uid uint, result string, payload any) (uint, error) {

	if result == "" {
		return 0, errors.New("图片生成失败")
	}

	// 解析输入参数

	param := ImagePayload{}
	if mapstructure.Decode(payload, &param) != nil {
		return 0, errors.New("参数解析失败")
	}

	// 保存生成图片

	outputFile := strings.ReplaceAll(uuid.NewString(), "-", "/") + ".jpg"
	if SaveBase64Image(outputFile, result) != nil {
		return 0, errors.New("图片保存失败")
	}

	// 保存原始图片

	if param.InputImage != "" {
		imageBase64 := strings.Split(param.InputImage, ",")[1]
		imagePath := strings.ReplaceAll(outputFile, ".jpg", "-input.png")
		if SaveBase64Image(imagePath, imageBase64) == nil {
			param.InputImage = imagePath
		}
	}

	// 生成记录入库

	return artimg.Create(&artimg.CreateParam{
		UserId:     uid,
		Subject:    param,
		OutputFile: outputFile,
		Status:     "success",
	})

}

func SaveBase64Image(filePath, base64Image string) error {

	filePath = args.Dataset.Dir + "/upload/" + filePath
	os.MkdirAll(path.Dir(filePath), 0755) // 递归创建目录

	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
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
