package database

import "time"

type Log struct {
	Id         int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                     //logid
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                 //用户id
	CreateTime time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //执行时间
	Operation  string    `gorm:"column:operation" db:"operation" json:"operation" form:"operation"`         //操作内容
	Type       int       `gorm:"column:type" db:"type" json:"type" form:"type"`                             //用户类型
}

func (l *Log) TableName() string {
	return "log"
}

type LogImpl struct {
}

func (l *LogImpl) AddLog(log *Log) error {
	db := GetDB()
	err := db.Create(&log).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *LogImpl) ListLog() ([]*Log, error) {
	db := GetDB()
	logs := new([]*Log)
	err := db.Model(&Log{}).Find(logs).Error
	if err != nil {
		return nil, err
	}
	return *logs, nil
}

func (l *LogImpl) DelLogById(idList ...int64) error {
	db := GetDB()
	err := db.Model(&Log{}).Where("id = ?", idList).Delete(&Log{}).Error
	if err != nil {
		return err
	}
	return nil
}
