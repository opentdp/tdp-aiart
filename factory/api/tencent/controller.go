package tencent

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/helper/tencent"
	"tdp-aiart/module/model/config"
)

func apiProxy(c *gin.Context) {

	secretId, secretKey := fetchSecret()

	if secretId == "" || secretKey == "" {
		c.Set("Error", "请先配置SecretId和SecretKey")
		return
	}

	// 绑定参数

	param := &tencent.ReqeustParam{
		SecretId:  secretId,
		SecretKey: secretKey,
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

// 获取SecretId和SecretKey

func fetchSecret() (string, string) {

	secretId := ""
	secretKey := ""

	rq := &config.FetchAllParam{Module: "Tencent"}
	if res, err := config.FetchAll(rq); err == nil {
		for _, rs := range res {
			if rs.Name == "secretId" {
				secretId = rs.Value
			}
			if rs.Name == "secretKey" {
				secretKey = rs.Value
			}
		}
	}

	return secretId, secretKey

}
