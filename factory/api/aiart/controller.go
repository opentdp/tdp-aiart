package aiart

import (
	"strings"
	"tdp-aiart/cmd/args"
	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/artimg"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func create(c *gin.Context) {

	// 构造参数

	param := &tencent.ReqeustParam{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 读取参数

	payload := ToImageRequest{}

	if err := mapstructure.Decode(param.Payload, &payload); err != nil {
		c.Set("Error", err)
		return
	}

	// 请求接口

	res, err := apiProxy(param)

	if err != nil {
		c.Set("Error", err)
		return
	}

	// 保存图片

	outputP, _ := saveOutputImage(res)

	if payload.InputImage != "" {
		inputP := outputP + "-input.png"
		base64 := strings.Split(payload.InputImage, ",")[1]
		if saveBase64ImageToFile(base64, inputP) == nil {
			payload.InputImage = strings.ReplaceAll(inputP, args.Dataset.Dir, "")
		}
	}

	artimg.Create(&artimg.CreateParam{
		Subject:    payload,
		OutputFile: strings.ReplaceAll(outputP, args.Dataset.Dir, ""),
	})

	// 输出数据

	c.Set("Payload", res)

}

type ToImageRequest struct {
	InputImage     string
	InputUrl       string
	Prompt         string
	NegativePrompt string
	Styles         []string
	ResultConfig   any
	LogoAdd        int64
	LogoParam      any
	Strength       float64
}
