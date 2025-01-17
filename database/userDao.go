package database

import (
	"database/sql"
	"fmt"
	"jayice/util"
	"time"
)

type User struct {
	Id         int64          `gorm:"column:id;type:int(11);primary_key" json:"id"`                      // 主键;雪花算法生成
	Username   string         `gorm:"column:username;type:varchar(20);NOT NULL" json:"username"`         // 用户名
	Password   string         `gorm:"column:password;type:varchar(20);NOT NULL" json:"password"`         // 密码
	Phone      int64          `gorm:"column:phone;type:int(11)" json:"phone"`                            // 手机号
	Type       int64          `gorm:"column:type;type:int(11);NOT NULL" json:"type"`                     // 用户类型;0-用户 1-管理员
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`      // 创建日期
	IsDelete   int            `gorm:"column:is_delete;type:int(11);default:0;NOT NULL" json:"is_delete"` // 逻辑删除;1-删除
	Ext        sql.NullString `gorm:"column:ext;type:varchar(255)" json:"ext"`                           // 预留
}

func (m *User) TableName() string {
	return "user"
}

type UserImpl struct {
}

// 验证是否存在用户
func (u *UserImpl) ExistUser(phone, category int64) bool {
	db := GetDB()
	var user User
	err := db.Model(&User{}).Where("phone = ? and type = ? and is_delete=0", phone, category).First(&user).Error
	if err != nil {
		//err==gorm.ErrRecordNotFound
		return false
	}
	fmt.Println(user)
	return true
}

func (u *UserImpl) AddUser(user *User) error {
	db := GetDB()
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// 用户登录
func (u *UserImpl) CheckUser(phone, category int64, password string) int64 {
	db := GetDB()
	var user User
	queryMap := make(map[string]string, 0)
	queryMap["phone"] = util.TranToString(phone)
	queryMap["type"] = util.TranToString(category)
	queryMap["is_delete"] = "0"
	queryMap["password"] = password
	err := db.Model(&User{}).Where(queryMap).First(&user).Error
	if err != nil {
		//err==gorm.ErrRecordNotFound
		return -1
	}
	fmt.Println(user)
	return user.Id
}

// 用户信息
func (u *UserImpl) GetUserById(idList ...int64) ([]*User, error) {
	db := GetDB()
	users := new([]*User)
	err := db.Model(&User{}).Where("id in ?", idList).Find(users).Error
	if err != nil {
		return nil, err
	}
	return *users, nil
}

// 用户列表
func (u *UserImpl) ListUser() ([]*User, error) {
	db := GetDB()
	users := new([]*User)

	queryMap := make(map[string]string, 0)
	queryMap["type"] = string(0)
	queryMap["is_delete"] = "0"

	err := db.Model(&User{}).Where(queryMap).Find(users).Error
	if err != nil {
		return nil, err
	}
	return *users, nil
}

// 删除用户
func (u *UserImpl) DelUserById(idList ...int64) error {
	db := GetDB()
	err := db.Model(&User{}).Where("id = ?", idList).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}
