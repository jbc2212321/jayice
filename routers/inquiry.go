package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/middleware"
	"ticket/util"
)

//var verifyListDao database.VerifylistImpl

type InquiryParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	UserId string `json:"userId" binding:"required"`
}

type InquiryByStatusParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Status string `json:"status" binding:"required"`
}

func Inquiry(c *gin.Context) {
	var json InquiryParam
	resp := util.GetResponse()
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Data = false
		c.JSON(http.StatusBadRequest, resp)
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		return
	}
	type Varify struct {
		list_id  int
		songid   int
		songname string
		userid   int
		username string
		status   int
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	verifyList, _ := verifyListDao.GetVerifyListByID(util.TranToInt64(json.UserId))
	resp.Data = verifyList

	c.JSON(http.StatusOK, resp)
}

func InquiryByStatus(c *gin.Context) {
	var json InquiryByStatusParam
	resp := util.GetResponse()
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Data = false
		c.JSON(http.StatusBadRequest, resp)
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		return
	}
	type Varify struct {
		list_id  int
		songid   int
		songname string
		userid   int
		username string
		status   int
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	//fmt.Println(json.status)
	verifyList, _ := verifyListDao.GetVerifyByStatus(util.TranToInt64(json.Status))
	fmt.Println(verifyList)
	resp.Data = verifyList

	c.JSON(http.StatusOK, resp)
}
