package aiart

import (
	"tdp-aiart/helper/tencent"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {

	param := &tencent.ReqeustParam{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	res, err := apiProxy(param)

	if err == nil {
		err = saveImage(res)
	}

	if err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
