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

var imageDao database.ImageImpl
var vatDao database.VatImpl

type SaveTicketParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	//Username string `json:"username" binding:"required"`
	//Password string `json:"password" binding:"required"`
	UserId     string              `json:"userId" binding:"required"`
	ImageName  string              `json:"imageName" binding:"required"`
	TicketImg  string              `json:"ticketImg" binding:"required"`
	VatInvoice database.VatInvoice `json:"vatInvoice" binding:"required"`
}

//保存识别后的小票
func SaveTicket(c *gin.Context) {
	//var json database.VatInvoice
	var json SaveTicketParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		//fmt.Println("err:", err)
		resp.Status = util.JSONError
		return
	}
	if json.ImageName == "" {
		resp.Status = util.IMAGEError
		return
	}
	fmt.Println("图片名称:", json.ImageName)
	//保存到vatInvoice表中
	vat := &json.VatInvoice
	vat.Id = util.GetSnowflakeId()
	vat.UserId = util.TranToInt64(json.UserId)
	vat.CreateTime = time.Now()
	err := vatDao.AddVat(vat)
	if err != nil {
		resp.Status = util.DBError
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
		TicketId:      vat.Id,
		BinaryData:    []byte(ticketImage),
		OcrBinaryData: []byte(json.TicketImg),
		Type:          0,
		CreateTime:    time.Now(),
	}

	err = imageDao.AddImage(img)
	if err != nil {
		resp.Status = util.DBError
		return
	}

	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "保存发票"))

	resp.Status = util.SUCCESS
	resp.Message = "保存成功"
	c.JSON(http.StatusOK, resp)
}
