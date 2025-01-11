package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"ticket/database"
	"ticket/util"
	"time"
)

func TestJson(t *testing.T) {
	data, err := ioutil.ReadFile("C:\\Users\\78240\\go\\ticket\\image\\image.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	v := &util.VatInvoice{}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}

	fmt.Println(util.TranToStringFromStruct(v.WordsResult.CommodityName))
	AddTicket(v)
	//fmt.Println(string(v.WordsResult.CommodityAmount))
	//fmt.Println(v.WordsResult.CommodityName)
	//fmt.Println(v.WordsResult.CommodityTax)
	//fmt.Println(v.WordsResult.Password)    //密码区
	//fmt.Println(v.WordsResult.InvoiceType) //发票种类
}

func TranToVatDao(v *util.VatInvoice) *database.VatInvoice {
	vat := &database.VatInvoice{}
	vat.InvoiceCode = v.WordsResult.InvoiceCode
	vat.MachineCode = v.WordsResult.MachineCode
	vat.InvoiceTypeOrg = v.WordsResult.InvoiceTypeOrg
	vat.InvoiceNum = v.WordsResult.InvoiceNum
	vat.InvoiceDate = v.WordsResult.InvoiceDate
	vat.PurchaserName = v.WordsResult.PurchaserName
	vat.PurchaserRegisterNum = v.WordsResult.PurchaserRegisterNum
	vat.PurchaserAddress = v.WordsResult.PurchaserAddress
	vat.PurchaserBank = v.WordsResult.PurchaserBank
	vat.Password = v.WordsResult.Password
	vat.CommodityName = util.TranToStringFromStruct(v.WordsResult.CommodityName)
	vat.CommodityType = util.TranToStringFromStruct(v.WordsResult.CommodityType)
	vat.CommodityUnit = util.TranToStringFromStruct(v.WordsResult.CommodityUnit)
	vat.CommodityNum = util.TranToStringFromStruct(v.WordsResult.CommodityNum)
	vat.CommodityPrice = util.TranToStringFromStruct(v.WordsResult.CommodityPrice)
	vat.CommodityAmount = util.TranToStringFromStruct(v.WordsResult.CommodityAmount)
	vat.CommodityTaxRate = util.TranToStringFromStruct(v.WordsResult.CommodityTaxRate)
	vat.CommodityTax = util.TranToStringFromStruct(v.WordsResult.CommodityTax)
	vat.TotalAmount = v.WordsResult.TotalAmount
	vat.TotalTax = v.WordsResult.TotalTax
	vat.AmountInWords = v.WordsResult.AmountInWords
	vat.SellerName = v.WordsResult.SellerName
	vat.SellerRegisterNum = v.WordsResult.SellerRegisterNum
	vat.SellerAddress = v.WordsResult.SellerAddress
	vat.SellerBank = v.WordsResult.SellerBank
	vat.Remarks = v.WordsResult.Remarks
	vat.Payee = v.WordsResult.Payee
	vat.Checker = v.WordsResult.Checker
	vat.NoteDrawer = v.WordsResult.NoteDrawer
	return vat
}

func AddTicket(v *util.VatInvoice) {
	vat := &database.VatInvoice{
		Id:         util.GetSnowflakeId(),
		CreateTime: time.Now(),
		UserId:     123456,
	}
	vat.InvoiceCode = v.WordsResult.InvoiceCode
	vat.MachineCode = v.WordsResult.MachineCode
	vat.InvoiceTypeOrg = v.WordsResult.InvoiceTypeOrg
	vat.InvoiceNum = v.WordsResult.InvoiceNum
	vat.InvoiceDate = v.WordsResult.InvoiceDate
	vat.PurchaserName = v.WordsResult.PurchaserName
	vat.PurchaserRegisterNum = v.WordsResult.PurchaserRegisterNum
	vat.PurchaserAddress = v.WordsResult.PurchaserAddress
	vat.PurchaserBank = v.WordsResult.PurchaserBank
	vat.Password = v.WordsResult.Password
	vat.CommodityName = util.TranToStringFromStruct(v.WordsResult.CommodityName)
	vat.CommodityType = util.TranToStringFromStruct(v.WordsResult.CommodityType)
	vat.CommodityUnit = util.TranToStringFromStruct(v.WordsResult.CommodityUnit)
	vat.CommodityNum = util.TranToStringFromStruct(v.WordsResult.CommodityNum)
	vat.CommodityPrice = util.TranToStringFromStruct(v.WordsResult.CommodityPrice)
	vat.CommodityAmount = util.TranToStringFromStruct(v.WordsResult.CommodityAmount)
	vat.CommodityTaxRate = util.TranToStringFromStruct(v.WordsResult.CommodityTaxRate)
	vat.CommodityTax = util.TranToStringFromStruct(v.WordsResult.CommodityTax)
	vat.TotalAmount = v.WordsResult.TotalAmount
	vat.TotalTax = v.WordsResult.TotalTax
	vat.AmountInWords = v.WordsResult.AmountInWords
	vat.SellerName = v.WordsResult.SellerName
	vat.SellerRegisterNum = v.WordsResult.SellerRegisterNum
	vat.SellerAddress = v.WordsResult.SellerAddress
	vat.SellerBank = v.WordsResult.SellerBank
	vat.Remarks = v.WordsResult.Remarks
	vat.Payee = v.WordsResult.Payee
	vat.Checker = v.WordsResult.Checker
	vat.NoteDrawer = v.WordsResult.NoteDrawer
	var vatDao database.VatImpl
	b := vatDao.AddVat(vat)
	fmt.Println(b)
}

func TranToStringFromStruct(x []util.CommodityName) string {
	b, _ := json.Marshal(x)
	//fmt.Println(e)
	//fmt.Println(string(b)) // {"id":1,"name":"wxnacy"}
	return string(b)
}

func TranToStringFromStruct2(x interface{}) string {
	//v:=reflect.TypeOf(x)
	b, e := json.Marshal(x)
	//fmt.Println("v:",v)
	fmt.Println("err:", e)
	fmt.Println("b:", string(b)) // {"id":1,"name":"wxnacy"}
	return string(b)
}
