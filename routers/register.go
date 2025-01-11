package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/middleware"
	"ticket/util"
	"time"
)

type RegisterParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Category string `json:"value" binding:"required"`
}

//新用户注册
func Register(c *gin.Context) {
	var json RegisterParam
	resp := util.GetResponse()
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		return
	}

	//验证是否注册过
	if userDao.ExistUser(util.TranToInt64(json.Phone), util.TranToInt64(json.Category)) {
		middleware.Log.Infof("该用户已经注册过,phone[%s],type[%s]", json.Phone, json.Category)
		resp.Message = "该用户已经注册过"
		resp.Data = false
		c.JSON(http.StatusOK, resp)
		return
	}

	// 用户注册
	user := &database.User{
		Id:         util.GetSnowflakeId(),
		Username:   json.Username,
		Password:   json.Password,
		Phone:      util.TranToInt64(json.Phone),
		Type:       util.TranToInt64(json.Category),
		CreateTime: time.Now(),
	}
	err := userDao.AddUser(user)
	if err != nil {
		resp.Code = util.DBError
		resp.Data = false
		resp.Message = "注册失败"
		c.JSON(http.StatusOK, resp)
		middleware.Log.WithError(err).Infof("注册失败,user[%+v]", user)
		return
	}

	_ = logDao.AddLog(GetLog(user.Id, "注册"))

	resp.Message = "注册成功"
	c.JSON(http.StatusOK, resp)
}
