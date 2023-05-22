package chatbot

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/chatbot"
)

func create(c *gin.Context) {

	var rq *chatbot.ReqeustParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := chatbot.Create(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
