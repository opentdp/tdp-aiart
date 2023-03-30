package aiart

import (
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {

	// 构造参数

	param := &ReqeustParams{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 请求接口

	res, err := apiProxy(param)

	if err != nil {
		c.Set("Error", err)
		return
	}

	// 存储数据

	saveObject(&param.Payload, res)

	// 输出数据

	c.Set("Payload", res)

}
