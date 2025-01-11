package main

import (
	"github.com/gin-gonic/gin"
	"ticket/middleware"
	"ticket/routers"
	"ticket/util"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(util.Cors())

	r.GET("/", routers.Hello)
	r.POST("/checkUser", routers.Login)
	//用户注册
	r.POST("/addUser", routers.Register)

	//查询用户
	r.POST("/queryUser", routers.UserInfo)

	//用户列表
	r.POST("/listUser", routers.ListUser)

	//删除用户
	r.POST("/delUser", routers.DelUser)

	//上传小票
	r.POST("/user/upload", routers.Upload)

	//保存小票
	r.POST("/user/saveTicket", routers.SaveTicket)

	//展示小票
	r.POST("/user/showTicket", routers.ShowTicket)

	//小票列表
	r.POST("/user/listTicket", routers.ListTicketByUserId)

	//删除小票
	r.POST("/user/delTicket", routers.DelTicketByTicketId)

	//保存火车票
	r.POST("/user/train/saveTrain", routers.SaveTrain)

	//日志列表
	r.POST("/admin/log/listLog", routers.ListLog)

	//删除日志
	r.POST("/admin/log/delLog", routers.DelLog)

	//------lyz-------//
	//上传歌曲
	r.POST("/user/song/upload", routers.UploadSong)
	//版权查询
	r.POST("/user/song/inquiry", routers.Inquiry)

	r.POST("/user/song/inquiryByStatus", routers.InquiryByStatus)

	//查看交易列表
	r.POST("/user/song/tradeList", routers.GetTradeList)
	//
	r.POST("/user/song/buy", routers.Buy)
	r.POST("/user/song/buyDel", routers.BuyDel)
	//申请版权
	r.POST("/user/song/applyCopyright", routers.ApplyCopyright)

	//更新
	r.POST("/user/song/updateVerifyList", routers.UpdateVerifyList)

	//提交交易
	r.POST("/user/song/addTradeList", routers.AddTradeList)
	_ = r.Run(":8096")
}
