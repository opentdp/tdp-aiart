package passport

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/strutil"
	"github.com/spf13/cast"

	"tdp-aiart/cmd/args"
	"tdp-aiart/model/config"
	"tdp-aiart/model/passport"
	"tdp-aiart/model/user"
)

// 注册用户

func register(c *gin.Context) {

	var rq *user.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	// 是否禁止注册
	if config.ValueOf("Registrable") != "true" {
		c.Set("Error", "抱歉，已关闭注册功能")
		return
	}

	// 校验用户信息
	if err := user.CheckUserinfo(rq.Username, rq.Password, rq.Email); err != nil {
		c.Set("Error", err)
		return
	}

	// 绘图默认配额
	if val := config.ValueOf("ArtworkQuota"); val != "" {
		rq.ArtworkQuota = cast.ToUint(val)
	}

	rq.Level = 0 //防止逃逸
	rq.AppKey = strutil.Rand(32)
	rq.StoreKey = args.Dataset.Secret

	if id, err := user.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq *passport.LoginParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.IpAddress = c.ClientIP()
	rq.UserAgent = c.Request.UserAgent()

	if res, err := passport.Login(rq); err == nil {
		c.Set("Message", "登录成功")
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 获取资料

func profile(c *gin.Context) {

	rq := &user.FetchParam{
		Id: c.GetUint("UserId"),
	}

	if res, err := user.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func profileUpdate(c *gin.Context) {

	var rq *passport.ProfileUpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = c.GetUint("UserId")
	rq.AppKey = "" //禁止修改

	if err := passport.ProfileUpdate(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改头像

func avatarUpdate(c *gin.Context) {

	var rq *passport.AvatarUpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if url, err := passport.AvatarUpdate(rq); err == nil {
		c.Set("Message", "修改成功")
		c.Set("Payload", gin.H{"Avatar": url})
	} else {
		c.Set("Error", err)
	}

}

// 统计信息

func summary(c *gin.Context) {

	userId := c.GetUint("UserId")
	res := passport.Summary(userId)

	c.Set("Payload", res)

}
