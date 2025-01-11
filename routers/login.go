package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/middleware"
	"ticket/util"
)

type LoginParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Phone    string `form:"username" json:"phone" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	Category string `json:"value" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginParam
	resp := util.GetResponse()
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Data = false
		c.JSON(http.StatusBadRequest, resp)
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		return
	}
	// 判断用户名密码是否正确
	userId := userDao.CheckUser(util.TranToInt64(json.Phone), util.TranToInt64(json.Category), json.Password)
	if userId == -1 {
		resp.Data = false
		middleware.Log.Infof("手机号或者密码错误,phone:[%s],password:[%s]", json.Phone, json.Password)
		return
	}

	_ = logDao.AddLog(GetLog(userId, "登录"))

	resp.Data = true
	resp.Message = util.TranToString(userId)
	c.JSON(http.StatusOK, resp)
}
