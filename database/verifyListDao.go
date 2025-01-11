package database

import "fmt"

type Verifylist struct {
	ListId   int    `gorm:"column:list_id" db:"list_id" json:"list_id" form:"list_id"`     //审核列表 id
	Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`         //歌曲id
	Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"` //歌曲名
	Userid   int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`         //用户id
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"` //用户名
	Status   int    `gorm:"column:status" db:"status" json:"status" form:"status"`         //申请状态
}

func (m *Verifylist) TableName() string {
	return "verifylist"
}

type VerifylistImpl struct {
}

func (v *VerifylistImpl) AddVerifyList(verify *Verifylist) error {
	db := GetDB()
	err := db.Create(&verify).Error
	if err != nil {
		return err
	}
	return nil
}

func (*VerifylistImpl) GetVerifyListByID(idList ...int64) ([]*Verifylist, error) {
	db := GetDB()
	veris := new([]*Verifylist)
	err := db.Model(&Verifylist{}).Where("userid = ?", idList).Find(veris).Error
	if err != nil {
		return nil, err
	}
	return *veris, nil
}

func (*VerifylistImpl) GetVerifyByStatus(statuList ...int64) ([]*Verifylist, error) {
	db := GetDB()
	veris := new([]*Verifylist)
	fmt.Println("statuList:", statuList)
	err := db.Model(&Verifylist{}).Where("status = ?", statuList).Find(veris).Error
	if err != nil {
		return nil, err
	}
	return *veris, nil
}

func (*VerifylistImpl) UpdateVerifyListById(id, status int) error {
	db := GetDB()
	err := db.Model(&Verifylist{}).Where("list_id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (*VerifylistImpl) ChangeVerifyListBySongID(idList int64, userName string, userId int64) error {
	db := GetDB()
	//	err := db.Create(&Belong).Error
	err := db.Model(&Verifylist{}).Where("songid = ?", idList).Updates(Verifylist{Username: userName, Userid: userId, Status: 1}).Error
	if err != nil {
		return err
	}
	return nil
}
