package database

type Belong struct {
	Num      int    `gorm:"column:num" db:"num" json:"num" form:"num"`                     //编号
	Userid   int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`         //用户id
	Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`         //歌曲id
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"` //用户名
	Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"` //歌曲名
}

func (m *Belong) TableName() string {
	return "belong"
}

type BelongImpl struct {
}

func (I *BelongImpl) AddBelong(belong *Belong) error {
	db := GetDB()
	err := db.Create(&belong).Error
	if err != nil {
		return err
	}
	return nil
}

func (I *BelongImpl) ChangeBelongBySongID(idList int64, userName string, userId int64) error {
	db := GetDB()
	//	err := db.Create(&Belong).Error
	err := db.Model(&Belong{}).Where("songid = ?", idList).Updates(Belong{Username: userName, Userid: userId}).Error
	if err != nil {
		return err
	}
	return nil
}
