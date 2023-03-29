package api

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/api/config"
	"tdp-aiart/api/passport"
	"tdp-aiart/api/user"

	"tdp-aiart/api/tencent"

	"tdp-aiart/module/midware"
)

func Router(engine *gin.Engine) {

	// application interface

	api := engine.Group("/api")

	api.Use(midware.OutputHandle())

	{
		config.Router(api)
		passport.Router(api)
		user.Router(api)

		tencent.Router(api)
	}

}
