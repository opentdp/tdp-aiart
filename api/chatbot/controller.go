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
			resp, ok := <-res
			if ok {
				event := "message"
				if resp.FinishReason == "<!finish>" {
					event = "stop"
				}
				if resp.FinishReason == "<!error>" {
					event = "stop"
				}
				c.SSEvent(event, resp)
			}
			return ok
		})
	} else {
		c.Set("Error", err)
	}

}
