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

	// 保存生成的图片

	filePath, err := SaveBase64Image("", result)
	if err != nil {
		return 0, errors.New("图片保存失败")
	}

	// 解析输入参数

	param := ImagePayload{}
	err = mapstructure.Decode(payload, &param)
	if err != nil {
		return 0, errors.New("参数解析失败")
	}

	// 保存原始图片

	if param.InputImage != "" {
		inputImagePath := filePath + "-input.png"
		inputImageBase64 := strings.Split(param.InputImage, ",")[1]
		if _, err := SaveBase64Image(inputImagePath, inputImageBase64); err == nil {
			param.InputImage = strings.ReplaceAll(inputImagePath, args.Dataset.Dir, "")
		}
	}

	// 生成记录入库

	outputImage := strings.ReplaceAll(filePath, args.Dataset.Dir, "")

	return artimg.Create(&artimg.CreateParam{
		UserId:     uid,
		Subject:    param,
		OutputFile: outputImage,
		Status:     "success",
	})

}

func SaveBase64Image(filePath, base64Image string) (string, error) {

	if filePath == "" {
		fileName := strings.ReplaceAll(uuid.NewString(), "-", "/")
		filePath = args.Dataset.Dir + "/upload/" + fileName + ".jpg"
	}

	os.MkdirAll(path.Dir(filePath), 0755) // 递归创建目录

	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return "", err
	}

	return filePath, nil

}
