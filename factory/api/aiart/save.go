package aiart

import (
	"encoding/base64"
	"errors"
	"os"
	"path"
	"strings"

	"tdp-aiart/cmd/args"
	"tdp-aiart/module/model/artimg"
)

func saveObject(payload *ToImageRequest, data *ToImageResponse) (uint, error) {

	outputP, _ := saveOutputImage(data)

	if payload.InputImage != "" {
		inputP := outputP + "-input.png"
		base64 := strings.Split(payload.InputImage, ",")[1]
		if saveBase64ImageToFile(base64, inputP) == nil {
			payload.InputImage = strings.ReplaceAll(inputP, args.Dataset.Dir, "")
		}
	}

	return artimg.Create(&artimg.CreateParam{
		Subject:    payload,
		OutputFile: strings.ReplaceAll(outputP, args.Dataset.Dir, ""),
	})

}

func saveOutputImage(data *ToImageResponse) (string, error) {

	if data.ResultImage == "" {
		return "", errors.New("图片生成失败")
	}

	fileName := strings.ReplaceAll(data.RequestId, "-", "/")
	filePath := args.Dataset.Dir + "/artimg/" + fileName + ".jpg"
	os.MkdirAll(path.Dir(filePath), 0755)

	return filePath, saveBase64ImageToFile(data.ResultImage, filePath)

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
