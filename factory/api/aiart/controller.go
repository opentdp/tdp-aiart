package aiart

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/artman"
)

func create(c *gin.Context) {

	// 构造参数

	param := &artman.ReqeustParams{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 请求接口

	res, err := artman.TencentAiart(param)

	if err != nil {
		c.Set("Error", err)
		return
	}

	// 存储数据

	userId := c.GetUint("UserId")
	artman.SaveObject(userId, res.ResultImage, param.Payload)

	// 输出数据

	c.Set("Payload", res)

}
