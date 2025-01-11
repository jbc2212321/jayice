package database

import "time"

type Image struct {
	Id            int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                                     //图片id
	TicketId      int64     `gorm:"column:ticket_id" db:"ticket_id" json:"ticket_id" form:"ticket_id"`                         //小票id
	Type          int       `gorm:"column:type" db:"type" json:"type" form:"type"`                                             //小票类型
	BinaryData    []byte    `gorm:"column:binary_data" db:"binary_data" json:"binary_data" form:"binary_data"`                 //图片2进制文件
	OcrBinaryData []byte    `gorm:"column:ocr_binary_data" db:"ocr_binary_data" json:"ocr_binary_data" form:"ocr_binary_data"` //识别后文件
	CreateTime    time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                 //创建日期
}

func (m *Image) TableName() string {
	return "image"
}

type ImageImpl struct {
}

func (I *ImageImpl) AddImage(img *Image) error {
	db := GetDB()
	err := db.Create(&img).Error
	if err != nil {
		return err
	}
	return nil
}

func (I *ImageImpl) GetImageByTicketId(idList ...int64) ([]*Image, error) {
	db := GetDB()
	images := new([]*Image)
	err := db.Model(&Image{}).Where("ticket_id in ?", idList).Find(images).Error
	if err != nil {
		return nil, err
	}
	return *images, nil
}

func (I *ImageImpl) ListImages() ([]*Image, error) {
	db := GetDB()
	images := new([]*Image)
	err := db.Model(&Image{}).Find(images).Error
	if err != nil {
		return nil, err
	}
	return *images, nil
}

func (I *ImageImpl) DelImageByTicketId(idList ...int64) error {
	db := GetDB()
	err := db.Model(&Image{}).Where("ticket_id = ?", idList).Delete(&Image{}).Error
	if err != nil {
		return err
	}
	return nil
}
