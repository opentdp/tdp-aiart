package painter

import (
	"encoding/base64"
	"errors"
	"os"
	"path"
	"strings"
	"time"

	"tdp-aiart/cmd/args"
	"tdp-aiart/helper/strutil"
)

func saveObject(param *ReqeustParam, base64Image string) (*ResponseData, error) {

	if base64Image == "" {
		return nil, errors.New("图片生成失败")
	}

	filePath := time.Now().Format("/2006/0102/1504/05")

	// 保存生成图片

	outputFile := filePath + strutil.Rand(8) + ".jpg"
	if saveBase64Image(outputFile, base64Image) != nil {
		return nil, errors.New("图片保存失败")
	}

	// 保存原始图片

	if param.InputImage != "" {
		imagePath := filePath + strutil.Rand(8) + ".png"
		imageBase64 := strings.Split(param.InputImage, ",")[1]
		if saveBase64Image(imagePath, imageBase64) == nil {
			param.InputImage = imagePath
		}
	}

	// 返回结果

	result := &ResponseData{
		InputFile:  param.InputImage,
		OutputFile: outputFile,
	}

	return result, nil

}

func saveBase64Image(filePath, base64Image string) error {

	filePath = args.Dataset.Dir + "/upload" + filePath
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
