package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/middleware"
	"ticket/util"
	"time"
)

type ShowTicketParam struct {
	UserId   string `json:"userId" binding:"required"`
	TicketId string `json:"ticketId" binding:"required"`
}

type ListTicketByUserIdParam struct {
	UserId string `json:"userId" binding:"required"`
}

type DelTicketByTicketIdParam struct {
	UserId   string `json:"userId" binding:"required"`
	TicketId string `json:"ticketId" binding:"required"`
}

//展示小票
func ShowTicket(c *gin.Context) {
	//var json database.VatInvoice
	var json ShowTicketParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	images, err := imageDao.GetImageByTicketId(util.TranToInt64(json.TicketId))
	if err != nil {
		resp.Status = util.DBError
		return
	}
	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	type Image struct {
		OcrBinaryData string
		BinaryData    string
	}
	img := Image{
		string(images[0].OcrBinaryData),
		string(images[0].BinaryData),
	}
	resp.Data = img
	//fmt.Println(string(images[0].BinaryData))
	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "展示小票"))

	c.JSON(http.StatusOK, resp)
}

//小票列表
func ListTicketByUserId(c *gin.Context) {
	//var json database.VatInvoice
	var json ListTicketByUserIdParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	vats, err := vatDao.GetVatByUserId(util.TranToInt64(json.UserId))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	trains, err := trainDao.GetTrainByUserId(util.TranToInt64(json.UserId))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	ticketIds := make([]int64, len(vats)+len(trains))

	for _, i := range vats {
		ticketIds = append(ticketIds, i.Id)
	}
	for _, i := range trains {
		ticketIds = append(ticketIds, i.Id)
	}

	//小票列表
	images, err := imageDao.GetImageByTicketId(ticketIds...)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	resp.Status = util.SUCCESS
	resp.Message = "读取成功"

	type ImageList struct {
		TicketId string    `json:"ticket_id"`
		Time     time.Time `json:"create_time"`
		Category string    `json:"type"`
	}
	imageList := make([]ImageList, 0)
	for _, v := range images {
		image := ImageList{
			TicketId: util.TranToString(v.TicketId),
			Time:     v.CreateTime,
		}
		if v.Type == 0 {
			image.Category = "增值税发票"
		} else {
			image.Category = "火车票"
		}
		imageList = append(imageList, image)
	}
	resp.Data = imageList
	//fmt.Println(images[0].TicketId)
	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "查看小票"))

	c.JSON(http.StatusOK, resp)
}

func DelTicketByTicketId(c *gin.Context) {
	var json DelTicketByTicketIdParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	//删除小票
	err := vatDao.DelVatByTicketId(util.TranToInt64(json.TicketId))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	err = imageDao.DelImageByTicketId(util.TranToInt64(json.TicketId))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	err = trainDao.DelTrainByTicketId(util.TranToInt64(json.TicketId))
	if err != nil {
		resp.Status = util.DBError
		return
	}

	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "删除小票"))

	resp.Status = util.SUCCESS
	resp.Message = "删除成功"
	c.JSON(http.StatusOK, resp)
}
