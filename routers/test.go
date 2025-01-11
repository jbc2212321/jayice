package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Hello(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"url":    c.Request.RequestURI,
		"method": c.Request.Method,
		"IP":     c.ClientIP(),
	})
	// 定义一个返回数据的结构体

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
