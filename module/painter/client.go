package painter

import (
	"errors"
	"strings"

	"tdp-aiart/module/upload"
)

type ReqeustParam struct {
	Action         string
	Subject        string
	Prompt         string
	NegativePrompt string
	InputImage     string
	Strength       float64
	Styles         []string
	ResultConfig   map[string]any
	LogoAdd        int64
	LogoParam      map[string]any
	Status         string
}

type ResponseData struct {
	InputFile  string
	OutputFile string
}

func Create(rq *ReqeustParam) (*ResponseData, error) {

	base64Image, err := TencentAiart(rq)

	if err != nil {
		return nil, err
	}

	if base64Image == "" {
		return nil, errors.New("图片生成失败")
	}

	return saveObject(rq, base64Image)

}

func saveObject(param *ReqeustParam, base64Image string) (*ResponseData, error) {

	filePath := "/artwork" + upload.TimePathname(7)

	// 保存生成图片

	outputFile := filePath + "o.jpg"
	if upload.SaveBase64Image(outputFile, base64Image) != nil {
		return nil, errors.New("图片保存失败")
	}

	// 保存原始图片

	if param.InputImage != "" {
		imagePath := filePath + "i.png"
		imageBase64 := strings.Split(param.InputImage, ",")[1]
		if upload.SaveBase64Image(imagePath, imageBase64) == nil {
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
