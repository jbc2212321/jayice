package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jayice/middleware"
	"jayice/util"
	"net/http"
)

type UserInfoParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	UserId string `json:"userId" binding:"required"`
}

// 查询用户信息
func UserInfo(c *gin.Context) {
	var json UserInfoParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	users, err := userDao.GetUserById(util.TranToInt64(json.UserId))
	if err != nil {
		resp.Status = util.DBError
		return
	}
	resp.Status = util.SUCCESS
	resp.Data = users[0]
	c.JSON(http.StatusOK, resp)
}
