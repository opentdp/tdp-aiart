package tencent

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/config"
)

func apiProxy(c *gin.Context) {

	kv := config.ValuesOf("Tencent")

	if kv["secretId"] == "" || kv["secretKey"] == "" {
		c.Set("Error", "请先配置SecretId和SecretKey")
		return
	}

	// 绑定参数

	param := &tencent.ReqeustParam{
		SecretId:  kv["secretId"],
		SecretKey: kv["secretKey"],
	}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	if res, err := tencent.Request(param); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
