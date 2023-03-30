package api

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/api/aiart"
	"tdp-aiart/api/config"
	"tdp-aiart/api/passport"
	"tdp-aiart/api/user"

	"tdp-aiart/module/midware"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.OutputHandle())

	{
		aiart.Router(api)
		config.Router(api)
		passport.Router(api)
		user.Router(api)
	}

}
