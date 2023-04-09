package artwork

import (
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-aiart/module/model/artwork"
	"tdp-aiart/module/painter"
)

// 图片列表

func list(c *gin.Context) {

	var rq *artwork.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.UserId == 0 || rq.UserId != c.GetUint("UserId") {
		rq.Status = "public"
	}

	if lst, err := artwork.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取图片

func detail(c *gin.Context) {

	var rq *artwork.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := artwork.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加图片

func create(c *gin.Context) {

	// 构造参数

	param := &painter.ReqeustParam{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 请求接口

	res, err := painter.Create(param)

	if err != nil {
		c.Set("Error", err)
		return
	}

	// 存储数据

	rq := &artwork.CreateParam{
		UserId:         c.GetUint("UserId"),
		Subject:        param.Subject,
		Prompt:         param.Prompt,
		NegativePrompt: param.NegativePrompt,
		Styles:         strings.Join(param.Styles, ","),
		Strength:       param.Strength,
		InputFile:      res.InputFile,
		OutputFile:     res.OutputFile,
		Status:         param.Status,
	}

	if id, err := artwork.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id, "OutputFile": res.OutputFile})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改图片

func update(c *gin.Context) {

	var rq *artwork.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artwork.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除图片

func delete(c *gin.Context) {

	var rq *artwork.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artwork.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
