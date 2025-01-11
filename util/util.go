package util

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

//int64 to string
func TranToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//string to int64
func TranToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//struct to string
func TranToStringFromStruct(x interface{}) string {
	b, e := json.Marshal(x)
	if e != nil {
		fmt.Println("err:", e)
	}
	return string(b)
}

func GetResponse() Response {
	return Response{
		Status: http.StatusOK,
		Code:   0,
	}
}

func ReadConfig(configPath, configName, configType string, key ...string) []string {
	config := viper.New()
	config.AddConfigPath(configPath) //设置读取的文件路径
	config.SetConfigName(configName) //设置读取的文件名
	config.SetConfigType(configType) //设置文件的类型

	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//文件读取出来的内容:
	res := make([]string, 0, len(key))
	if len(key) > 0 {
		for _, v := range key {
			res = append(res, config.Get(v).(string))
		}
	}
	return res

}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

//解决中文乱码
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

//删除文件
func DelFileByDst(dst string) {
	dir, _ := ioutil.ReadDir(dst)
	for _, d := range dir {
		_ = os.RemoveAll(path.Join([]string{dst, d.Name()}...))
	}
}

//添加日志
//func GetLog(userId int64, operation string) *database.Log {
//	var userDao database.UserImpl
//	user, _ := userDao.GetUserById(userId)
//	log := &database.Log{
//		Id:         GetSnowflakeId(),
//		UserId:     userId,
//		CreateTime: time.Now(),
//		Operation:  operation,
//		Type:       int(user[0].Type),
//	}
//	return log
//}

//string 转int
func TranToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
