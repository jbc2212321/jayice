package test

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
	"ticket/util"
)

func TestConfig(t *testing.T) {
	config := viper.New()
	config.AddConfigPath("../config") //设置读取的文件路径
	config.SetConfigName("db")        //设置读取的文件名
	config.SetConfigType("yaml")      //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//打印文件读取出来的内容:
	fmt.Println(config.Get("mysql"))
	fmt.Println(config.Get("mysql.dbname").(string))
	username := config.Get("mysql.username").(string)
	password := config.Get("mysql.password").(string)
	//dbname:=config.Get("mysql.dbname").(string)
	fmt.Println(username, password)
}

func TestGetConfig(t *testing.T) {
	res := util.ReadConfig("../config", util.DBConfig, util.YAML, "mysql.dbname", "mysql.username")
	fmt.Println(res)
}
