package artwork

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/artwork")

	{
		rg.POST("/list", list)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/create", create)
		rg.POST("/detail", detail)
		rg.POST("/update", update)
		rg.POST("/delete", delete)
	}

}
