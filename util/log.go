package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func initLogrus() error {
	// 设置为json格式的日志
	log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("./gin_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("创建日志文件/打开日志文件失败")
		return err
	}
	// 设置log默认文件输出
	log.Out = file
	gin.SetMode(gin.ReleaseMode)
	// gin框架自己记录的日志也会输出
	gin.DefaultWriter = log.Out
	// 设置日志级别
	log.Level = logrus.InfoLevel
	return nil
}

func init() {
	err := initLogrus()
	if err != nil {
		fmt.Println(err)
		return
	}
}
