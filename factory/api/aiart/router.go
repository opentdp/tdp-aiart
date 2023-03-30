package aiart

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/aiart")

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.POST("/create", create)
	}

}
