package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jayice/database"
	"jayice/middleware"
	"jayice/util"
	"net/http"
	"time"
)

var logDao database.LogImpl

func GetLog(userId int64, operation string) *database.Log {
	var userDao database.UserImpl
	user, _ := userDao.GetUserById(userId)
	log := &database.Log{
		Id:         util.GetSnowflakeId(),
		UserId:     userId,
		CreateTime: time.Now(),
		Operation:  operation,
		Type:       int(user[0].Type),
	}
	return log
}

// 日志列表
func ListLog(c *gin.Context) {
	resp := util.GetResponse()
	logs, err := logDao.ListLog()
	if err != nil {
		resp.Status = util.DBError
		return
	}
	logsManagements := make([]*util.LogManagement, 0)
	for _, log := range logs {
		logsManagement := &util.LogManagement{
			Num:    util.TranToString(log.Id),
			Id:     util.TranToString(log.UserId),
			Action: log.Operation,
			Time:   log.CreateTime,
		}
		user, _ := userDao.GetUserById(log.UserId)
		logsManagement.UserName = user[0].Username
		logsManagement.Phone = util.TranToString(user[0].Phone)
		if user[0].Type == 0 {
			logsManagement.Account = "用户"
		} else {
			logsManagement.Account = "管理员"
		}
		logsManagements = append(logsManagements, logsManagement)
	}
	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	resp.Data = logsManagements
	c.JSON(http.StatusOK, resp)
}

type DelLogParam struct {
	//UserId   string `json:"userId" binding:"required"`
	Id string `json:"num" binding:"required"`
}

// 删除日志
func DelLog(c *gin.Context) {
	resp := util.GetResponse()
	var json DelLogParam

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	err := logDao.DelLogById(util.TranToInt64(json.Id))
	if err != nil {
		resp.Code = util.DBError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	resp.Code = util.SUCCESS
	c.JSON(http.StatusOK, resp)
}
