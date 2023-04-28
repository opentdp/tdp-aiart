package chatbot

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/openai"
)

func message(c *gin.Context) {

	var rq *[]openai.ChatCompletionMessage

	if err := c.ShouldBind(rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := openai.Chat(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}
