package chatbot

import (
	"io"

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

func stream(c *gin.Context) {

	var rq *chatbot.ReqeustParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := chatbot.CreateStream(rq); err == nil {
		c.Stream(func(w io.Writer) bool {
			if resp, ok := <-res; ok {
				if resp.FinishReason == "<!finish>" {
					c.SSEvent("stop", "finish")
				}
				if resp.FinishReason == "<!error>" {
					c.SSEvent("stop", "error")
				}
				c.SSEvent("message", resp)
				return true
			}
			return false
		})
	} else {
		c.Set("Error", err)
	}

}
