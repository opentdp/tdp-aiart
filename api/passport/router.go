package passport

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/passport")

	// 匿名接口

	{
		rg.POST("/login", login)
		rg.POST("/register", register)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/profile", profile)
		rg.POST("/profile/update", profileUpdate)
		rg.POST("/avatar/update", avatarUpdate)
		rg.POST("/summary", summary)
	}

}
