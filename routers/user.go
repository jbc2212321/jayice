package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/middleware"
	"ticket/util"
)

var userDao database.UserImpl

//用户列表
func ListUser(c *gin.Context) {
	resp := util.GetResponse()
	users, err := userDao.ListUser()
	if err != nil {
		resp.Status = util.DBError
		return
	}
	userInfos := make([]*util.UserInfo, 0)
	for _, user := range users {
		userInfo := &util.UserInfo{
			Num:  util.TranToString(user.Id),
			Name: user.Username,
			Tel:  util.TranToString(user.Phone),
			Time: user.CreateTime,
		}
		userInfos = append(userInfos, userInfo)
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	resp.Data = userInfos
	c.JSON(http.StatusOK, resp)
}

type DelUserParam struct {
	UserId string `json:"userId" binding:"required"`
	//Id string `json:"userId" binding:"required"`
}

//删除日志
func DelUser(c *gin.Context) {
	resp := util.GetResponse()
	var json DelUserParam

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	err := userDao.DelUserById(util.TranToInt64(json.UserId))
	if err != nil {
		resp.Code = util.DBError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	resp.Code = util.SUCCESS
	c.JSON(http.StatusOK, resp)
}
