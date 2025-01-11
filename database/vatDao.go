package database

import (
	"time"
)

type VatInvoice struct {
	Id                   int64     `gorm:"column:id;primary_key" json:"id"`                         // 发票id
	UserId               int64     `gorm:"column:user_id;NOT NULL" json:"user_id"`                  // 提交人id
	CreateTime           time.Time `gorm:"column:create_time;NOT NULL" json:"create_time"`          // 创建时间
	InvoiceCode          string    `gorm:"column:InvoiceCode" json:"InvoiceCode"`                   // 发票代码
	MachineCode          string    `gorm:"column:MachineCode" json:"MachineCode"`                   // 机器编号
	InvoiceTypeOrg       string    `gorm:"column:InvoiceTypeOrg" json:"InvoiceTypeOrg"`             // 发票名称
	InvoiceNum           string    `gorm:"column:InvoiceNum" json:"InvoiceNum"`                     // No:
	InvoiceDate          string    `gorm:"column:InvoiceDate" json:"InvoiceDate"`                   // 开票日期
	PurchaserName        string    `gorm:"column:PurchaserName" json:"PurchaserName"`               // 名称
	PurchaserRegisterNum string    `gorm:"column:PurchaserRegisterNum" json:"PurchaserRegisterNum"` // 纳税人识别号
	PurchaserAddress     string    `gorm:"column:PurchaserAddress" json:"PurchaserAddress"`         // 地址、电话(PurchaserAddress)
	PurchaserBank        string    `gorm:"column:PurchaserBank" json:"PurchaserBank"`               // 购买方开户行及账号(PurchaserBank)
	Password             string    `gorm:"column:Password" json:"Password"`                         // 密码区
	CommodityName        string    `gorm:"column:CommodityName" json:"CommodityName"`               // 货物或应税劳务、服务名称
	CommodityType        string    `gorm:"column:CommodityType" json:"CommodityType"`               // 型号
	CommodityUnit        string    `gorm:"column:CommodityUnit" json:"CommodityUnit"`               // 单位
	CommodityNum         string    `gorm:"column:CommodityNum" json:"CommodityNum"`                 // 数量
	CommodityPrice       string    `gorm:"column:CommodityPrice" json:"CommodityPrice"`             // 单价
	CommodityAmount      string    `gorm:"column:CommodityAmount" json:"CommodityAmount"`           // 金额
	CommodityTaxRate     string    `gorm:"column:CommodityTaxRate" json:"CommodityTaxRate"`         // 税率
	CommodityTax         string    `gorm:"column:CommodityTax" json:"CommodityTax"`                 // 税额
	TotalAmount          string    `gorm:"column:TotalAmount" json:"TotalAmount"`                   // 金额合计
	TotalTax             string    `gorm:"column:TotalTax" json:"TotalTax"`                         // 税额合计
	AmountInWords        string    `gorm:"column:AmountInWords" json:"AmountInWords"`               // 价税大写
	AmountInFiguers      string    `gorm:"column:AmountInFiguers" json:"AmountInFiguers"`           // 价税小写
	SellerName           string    `gorm:"column:SellerName" json:"SellerName"`                     // 名称(SellerName)
	SellerRegisterNum    string    `gorm:"column:SellerRegisterNum" json:"SellerRegisterNum"`       // 纳税人识别号(SellerRegisterNum)
	SellerAddress        string    `gorm:"column:SellerAddress" json:"SellerAddress"`               // 地址、电话(SellerAddress)
	SellerBank           string    `gorm:"column:SellerBank" json:"SellerBank"`                     // 销售方开户行及账号
	Remarks              string    `gorm:"column:Remarks" json:"Remarks"`                           // 备注
	Payee                string    `gorm:"column:Payee" json:"Payee"`                               // 收款人
	Checker              string    `gorm:"column:Checker" json:"Checker"`                           // 复核人
	NoteDrawer           string    `gorm:"column:NoteDrawer" json:"NoteDrawer"`                     // 开票人
}

func (m *VatInvoice) TableName() string {
	return "VatInvoice"
}

type VatImpl struct {
}

func (v *VatImpl) AddVat(vat *VatInvoice) error {
	db := GetDB()
	err := db.Create(&vat).Error
	if err != nil {
		return err
	}
	return nil
}

func (v *VatImpl) GetVatByUserId(idList ...int64) ([]*VatInvoice, error) {
	db := GetDB()
	vats := new([]*VatInvoice)
	err := db.Model(&VatInvoice{}).Where("user_id in ?", idList).Find(vats).Error
	if err != nil {
		return nil, err
	}
	return *vats, nil
}

func (v *VatImpl) DelVatByTicketId(idList ...int64) error {
	db := GetDB()
	err := db.Model(&VatInvoice{}).Where("id = ?", idList).Delete(&VatInvoice{}).Error
	if err != nil {
		return err
	}
	return nil
}
