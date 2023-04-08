package aiart

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/artman"
	"tdp-aiart/module/model/artimg"
)

// 图片列表

func list(c *gin.Context) {

	var rq *artimg.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	// rq.UserId = c.GetUint("UserId")

	if lst, err := artimg.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取图片

func detail(c *gin.Context) {

	var rq *artimg.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := artimg.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加图片

func create(c *gin.Context) {

	// 构造参数

	param := &artman.ReqeustParams{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 请求接口

	res, err := artman.TencentAiart(param)

	if err != nil {
		c.Set("Error", err)
		return
	}

	// 存储数据

	userId := c.GetUint("UserId")
	artman.SaveObject(userId, res.ResultImage, param.Payload)

	// 输出数据

	c.Set("Payload", res)

}

// 修改图片

func update(c *gin.Context) {

	var rq *artimg.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artimg.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除图片

func delete(c *gin.Context) {

	var rq *artimg.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artimg.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
