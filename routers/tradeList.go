package routers

import (
	"fmt"
	"io/ioutil"
	"ticket/middleware"
	//	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/util"
)

type BuyParam struct {
	Num      string `json:"num" binding:"required"`
	Songid   string `json:"songid" binding:"required"`
	Songname string `json:"songname" binding:"required"`
	Owner    string `json:"owner" binding:"required"`
	Ownerid  string `json:"ownerid" binding:"required"`
	Buyerid  string `json:"buyerid" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

var tradeListDao database.TradelistImpl

func GetTradeList(c *gin.Context) {
	resp := util.GetResponse()

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	//fmt.Println(json.status)

	tradeLists, _ := tradeListDao.GetTradeList()
	type TradeInfos struct {
		Num      int    `gorm:"column:num" db:"num" json:"num" form:"num"`
		Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`
		Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"`
		Owner    string `gorm:"column:owner" db:"owner" json:"owner" form:"owner"`
		Buyer    string `gorm:"column:buyer" db:"buyer" json:"buyer" form:"buyer"`
		Ownerid  string `gorm:"column:ownerid" db:"ownerid" json:"ownerid" form:"ownerid"`
		Buyerid  string `gorm:"column:buyerid" db:"buyerid" json:"buyerid" form:"buyerid"`
		Status   int    `gorm:"column:status" db:"status" json:"status" form:"status"`
	}
	tradeInfos := make([]*TradeInfos, 0)
	for _, list := range tradeLists {
		tradeInfo := &TradeInfos{
			Num:      list.Num,
			Songid:   list.Songid,
			Songname: list.Songname,
			Owner:    list.Owner,
			Buyer:    list.Buyer,
			Ownerid:  util.TranToString(list.Ownerid),
			Buyerid:  util.TranToString(list.Buyerid),
			Status:   list.Status,
		}
		tradeInfos = append(tradeInfos, tradeInfo)
	}
	resp.Data = tradeInfos
	//verifyList, _ := verifyListDao.GetVerifyByStatus(util.TranToInt64(json.Status))

	c.JSON(http.StatusOK, resp)
}

var belongDao database.BelongImpl
var songDao database.SongsImpl

func Buy(c *gin.Context) {
	var json BuyParam

	resp := util.GetResponse()
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("data:", string(data))
		resp.Status = util.JSONError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	err := tradeListDao.DelTradeListById(util.TranToInt64(json.Num)) //imageDao.AddImage(img)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	userName, err := userDao.GetUserById(util.TranToInt64(json.Buyerid))
	fmt.Println("buyerName:", userName[0].Username)

	songs, err := songDao.GetSongByID(util.TranToInt64(json.Songid))
	fmt.Println("songName:", songs[0].Name)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	if json.Status == "1" {
		fmt.Println("json status:", 1)
		belong := &database.Belong{
			Userid:   util.TranToInt64(json.Buyerid),
			Username: userName[0].Username,
			Songid:   int(util.TranToInt64(json.Songid)),
			Songname: songs[0].Name,
		}
		err = belongDao.AddBelong(belong)
	} else {
		err = belongDao.ChangeBelongBySongID(util.TranToInt64(json.Songid), userName[0].Username, util.TranToInt64(json.Buyerid))
	}
	if err != nil {
		resp.Status = util.DBError
		return
	}

	if json.Status == "1" {
		verifyList := &database.Verifylist{
			Userid:   util.TranToInt64(json.Buyerid),
			Username: userName[0].Username,
			Songid:   int(util.TranToInt64(json.Songid)),
			Songname: songs[0].Name,
			Status:   1,
		}
		err = verifyListDao.AddVerifyList(verifyList)
	} else {
		err = verifyListDao.ChangeVerifyListBySongID(util.TranToInt64(json.Songid), userName[0].Username, util.TranToInt64(json.Buyerid))
		//	err=belongDao.ChangeBelongBySongID(util.TranToInt64(json.Songid),userName[0].Username,util.TranToInt64(json.Buyerid))
	}
	if err != nil {
		resp.Status = util.DBError
		return
	}
	//fmt.Println(json.status)

	//tradeLists,_:=tradeListDao.GetVerifyList();
	//resp.Data=tradeLists
	//verifyList, _ := verifyListDao.GetVerifyByStatus(util.TranToInt64(json.Status))

	c.JSON(http.StatusOK, resp)
}
func BuyDel(c *gin.Context) {
	var json BuyParam

	resp := util.GetResponse()
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("data:", string(data))
		resp.Status = util.JSONError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	err := tradeListDao.DelTradeListById(util.TranToInt64(json.Num))
	if err != nil {
		resp.Status = util.DBError
		return
	}
}

type AddTradeListParam struct {
	UserId  string `json:"ownerid" binding:"required"`
	BuyerId string `json:"buyerid" binding:"required"`
	SongId  string `json:"songid" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

//保存到trainlist 提交交易
func AddTradeList(c *gin.Context) {
	//var json database.VatInvoice
	var json AddTradeListParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		//fmt.Println("err:", err)
		resp.Status = util.JSONError
		return
	}

	usernames, _ := userDao.GetUserById(util.TranToInt64(json.UserId), util.TranToInt64(json.BuyerId))
	songs, _ := songImpl.GetSongByID(util.TranToInt64(json.SongId))
	tradelist := &database.Tradelist{
		Songid:   util.TranToInt(json.SongId),
		Songname: songs[0].Name,
		Owner:    usernames[0].Username,
		Buyer:    usernames[1].Username,
		Ownerid:  util.TranToInt64(json.UserId),
		Buyerid:  util.TranToInt64(json.BuyerId),
		Status:   util.TranToInt(json.Status),
	}
	err := tradeListDao.AddTradeList(tradelist)
	if err != nil {
		resp.Status = util.DBError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "保存成功"
	c.JSON(http.StatusOK, resp)
}
