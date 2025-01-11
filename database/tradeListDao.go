package database

type Tradelist struct {
	Num      int    `gorm:"column:num" db:"num" json:"num" form:"num"`                     //交易编号
	Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`         //歌曲id 与songs中的id对应
	Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"` //歌曲名
	Owner    string `gorm:"column:owner" db:"owner" json:"owner" form:"owner"`             //所有者
	Buyer    string `gorm:"column:buyer" db:"buyer" json:"buyer" form:"buyer"`             //购买者
	Ownerid  int64  `gorm:"column:ownerid" db:"ownerid" json:"ownerid" form:"ownerid"`     //所有者id
	Buyerid  int64  `gorm:"column:buyerid" db:"buyerid" json:"buyerid" form:"buyerid"`     //购买者id
	Status   int    `gorm:"column:status" db:"status" json:"status" form:"status"`         //交易状态
}

func (m *Tradelist) TableName() string {
	return "tradelist"
}

type TradelistImpl struct {
}

func (*TradelistImpl) GetTradeList() ([]*Tradelist, error) {
	db := GetDB()
	tradelist := new([]*Tradelist)
	//fmt.Println("statuList:", statuList)
	err := db.Model(&Tradelist{}).Where("1=1").Find(tradelist).Error
	if err != nil {
		return nil, err
	}
	return *tradelist, nil
}

func (*TradelistImpl) DelTradeListById(idList ...int64) error {
	db := GetDB()
	err := db.Model(&Tradelist{}).Where("num = ?", idList).Delete(&Tradelist{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TradelistImpl) AddTradeList(tradelist *Tradelist) error {
	db := GetDB()
	err := db.Create(&tradelist).Error
	if err != nil {
		return err
	}
	return nil
}
