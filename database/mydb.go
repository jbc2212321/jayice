package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Self *gorm.DB
}

// 单例
var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetDB(),
	}
}

func openDB(username, password, host, dbname string) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//logrus.Fatalf("数据库连接失败. 数据库名字: %s. 错误信息: %s", name, err)
	} else {
		//logrus.Infof("数据库连接成功, 数据库名字: %s", name)
	}

	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("error")
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func InitDB() *gorm.DB {
	return openDB(readConfig())
}

func GetDB() *gorm.DB {
	return InitDB()
}

func readConfig() (string, string, string, string) {
	config := viper.New()
	config.AddConfigPath("./config") //设置读取的文件路径
	//config.AddConfigPath("../config") //设置读取的文件路径 --test
	config.SetConfigName("db")   //设置读取的文件名
	config.SetConfigType("yaml") //设置文件的类型

	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//文件读取出来的内容:
	username := config.Get("mysql.username").(string)
	password := config.Get("mysql.password").(string)
	host := config.Get("mysql.host").(string)
	dbname := config.Get("mysql.dbname").(string)
	return username, password, host, dbname

}
