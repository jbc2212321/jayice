package util

import "time"

type Response struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type VatInvoice struct {
	WordsResultNum int         `json:"words_result_num"`
	WordsResult    WordsResult `json:"words_result"`
	LogID          int64       `json:"log_id"`
}
type CommodityTaxRate struct {
	Row  string `json:"row"`
	Word string `json:"word"`
}
type CommodityAmount struct {
	Row  string `json:"row"`
	Word string `json:"word"`
}
type CommodityTax struct {
	Row  string `json:"row"`
	Word string `json:"word"`
}
type CommodityName struct {
	Row  string `json:"row"`
	Word string `json:"word"`
}

type Common struct {
	Row  string `json:"row"`
	Word string `json:"word"`
}

type WordsResult struct {
	AmountInWords        string             `json:"AmountInWords"`
	InvoiceNumConfirm    string             `json:"InvoiceNumConfirm"`
	CommodityEndDate     []interface{}      `json:"CommodityEndDate"`
	CommodityVehicleType []interface{}      `json:"CommodityVehicleType"`
	CommodityStartDate   []interface{}      `json:"CommodityStartDate"`
	CommodityPrice       []interface{}      `json:"CommodityPrice"`
	NoteDrawer           string             `json:"NoteDrawer"`
	SellerAddress        string             `json:"SellerAddress"`
	CommodityNum         []interface{}      `json:"CommodityNum"`
	SellerRegisterNum    string             `json:"SellerRegisterNum"`
	MachineCode          string             `json:"MachineCode"`
	Remarks              string             `json:"Remarks"`
	SellerBank           string             `json:"SellerBank"`
	CommodityTaxRate     []CommodityTaxRate `json:"CommodityTaxRate"`
	TotalTax             string             `json:"TotalTax"`
	InvoiceCodeConfirm   string             `json:"InvoiceCodeConfirm"`
	CheckCode            string             `json:"CheckCode"`
	InvoiceCode          string             `json:"InvoiceCode"`
	InvoiceDate          string             `json:"InvoiceDate"`
	PurchaserRegisterNum string             `json:"PurchaserRegisterNum"`
	InvoiceTypeOrg       string             `json:"InvoiceTypeOrg"`
	OnlinePay            string             `json:"OnlinePay"`
	Password             string             `json:"Password"`
	Agent                string             `json:"Agent"`
	AmountInFiguers      string             `json:"AmountInFiguers"`
	PurchaserBank        string             `json:"PurchaserBank"`
	Checker              string             `json:"Checker"`
	City                 string             `json:"City"`
	TotalAmount          string             `json:"TotalAmount"`
	CommodityAmount      []CommodityAmount  `json:"CommodityAmount"`
	PurchaserName        string             `json:"PurchaserName"`
	CommodityType        []interface{}      `json:"CommodityType"`
	Province             string             `json:"Province"`
	InvoiceType          string             `json:"InvoiceType"`
	SheetNum             string             `json:"SheetNum"`
	PurchaserAddress     string             `json:"PurchaserAddress"`
	CommodityTax         []CommodityTax     `json:"CommodityTax"`
	CommodityPlateNum    []interface{}      `json:"CommodityPlateNum"`
	CommodityUnit        []interface{}      `json:"CommodityUnit"`
	Payee                string             `json:"Payee"`
	CommodityName        []CommodityName    `json:"CommodityName"`
	SellerName           string             `json:"SellerName"`
	InvoiceNum           string             `json:"InvoiceNum"`
}

//日志
type LogManagement struct {
	Num      string    `json:"num"`
	Id       string    `json:"id"`
	Action   string    `json:"action"`
	Account  string    `json:"account"`
	Time     time.Time `json:"time"`
	UserName string    `json:"username"`
	Phone    string    `json:"phone"`
}

type UserInfo struct {
	Num  string    `json:"num"`
	Time time.Time `json:"time"`
	Name string    `json:"name"`
	Tel  string    `json:"tel"`
}
