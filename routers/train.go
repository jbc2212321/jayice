package routers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"ticket/database"
	"ticket/middleware"
	"ticket/util"
	"time"
)

var trainDao database.TrainImpl

type SaveTrainParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	UserId    string         `json:"userId" binding:"required"`
	ImageName string         `json:"imageName" binding:"required"`
	TicketImg string         `json:"ticketImg" binding:"required"`
	Train     database.Train `json:"train" binding:"required"`
}

//上传火车票
func SaveTrain(c *gin.Context) {
	var json SaveTrainParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	train := &json.Train
	train.Id = util.GetSnowflakeId()
	train.UserId = util.TranToInt64(json.UserId)
	train.CreateTime = time.Now()
	err := trainDao.AddTrain(train)
	if err != nil {
		resp.Code = util.DBError
		return
	}

	//保存到image表中
	imgPath := util.ImagePath + json.ImageName
	fileBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		fmt.Println(err)
	}
	ticketImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(fileBytes)

	img := &database.Image{
		Id:            util.GetSnowflakeId(),
		TicketId:      train.Id,
		BinaryData:    []byte(ticketImage),
		OcrBinaryData: []byte(json.TicketImg),
		Type:          1,
		CreateTime:    time.Now(),
	}

	err = imageDao.AddImage(img)
	if err != nil {
		resp.Status = util.DBError
		return
	}

	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "保存车票"))

	resp.Code = util.SUCCESS
	resp.Message = "上传文件成功"
	c.JSON(http.StatusOK, resp)
}
