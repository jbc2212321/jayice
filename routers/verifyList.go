package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jayice/middleware"
	"jayice/util"
	"net/http"
)

type UpdateVerifyListParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Num    string `json:"num" binding:"required"`
	Status string `json:"status" binding:"required"`
}

// 更新
func UpdateVerifyList(c *gin.Context) {
	var json UpdateVerifyListParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		//fmt.Println("err:", err)
		resp.Status = util.JSONError
		return
	}

	err := verifyListDao.UpdateVerifyListById(util.TranToInt(json.Num), util.TranToInt(json.Status))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	resp.Code = util.SUCCESS
	resp.Message = "保存成功"
	c.JSON(http.StatusOK, resp)
}
