package api

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/api/artwork"
	"tdp-aiart/api/config"
	"tdp-aiart/api/passport"
	"tdp-aiart/api/upgrade"
	"tdp-aiart/api/user"

	"tdp-aiart/module/midware"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.OutputHandle())
	api.Use(midware.JwtGuard())

	{
		artwork.Router(api)
		config.Router(api)
		passport.Router(api)
		upgrade.Router(api)
		user.Router(api)
	}

}
