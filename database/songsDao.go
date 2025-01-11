package database

import "fmt"

type Songs struct {
	Id     int    `gorm:"column:id" db:"id" json:"id" form:"id"`         //歌曲ID
	Name   string `gorm:"column:name" db:"name" json:"name" form:"name"` //歌曲名称
	Status int    `gorm:"column:status" db:"status" json:"status" form:"status"`
}

func (m *Songs) TableName() string {
	return "songs"
}

type SongsImpl struct {
}

//歌曲信息
func (s *SongsImpl) GetSong() (*Songs, error) {
	db := GetDB()
	song := new(*Songs)
	err := db.Model(&Songs{}).Last(song).Error
	if err != nil {
		return nil, err
	}
	return *song, nil
}

func (s *SongsImpl) GetSongByID(idList ...int64) ([]*Songs, error) {
	fmt.Println("id:", idList)
	db := GetDB()
	song := new([]*Songs)
	err := db.Model(&Songs{}).Where("id = ?", idList).Find(song).Error
	fmt.Println("err:", err)
	fmt.Println("song name in func:", song)
	if err != nil {
		return nil, err
	}
	return *song, nil
}
