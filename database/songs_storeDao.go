package database

type SongsStore struct {
	Seq   int    `gorm:"column:seq" db:"seq" json:"seq" form:"seq"`         //歌曲上传的顺序
	Store []byte `gorm:"column:store" db:"store" json:"store" form:"store"` //歌曲本身
	Id    int    `gorm:"column:id" db:"id" json:"id" form:"id"`             //唯一识别id
}

func (m *SongsStore) TableName() string {
	return "songs_store"
}

type SongsStoreImpl struct {
}

func (m *SongsStoreImpl) AddSong(song *SongsStore) error {
	db := GetDB()
	err := db.Create(&song).Error
	if err != nil {
		return err
	}
	return nil
}
