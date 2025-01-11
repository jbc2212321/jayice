package test

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	logrus.WithFields(logrus.Fields{"name": "张三"}).Info("这是日志信息")
}
