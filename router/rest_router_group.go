package router

import (
	"github.com/gin-gonic/gin"
		"github.com/sslab-instapay/instapay-go-client/controller"
)

// 라우터 등록 코드
func RegisterRestAccountRouter(router *gin.Engine){

	accountRouter := router.Group("account")
	{
		accountRouter.GET("list", controller.AccountInformationHandler)
	}
}

func RegisterRestChannelRouter(router *gin.Engine){

	channelRouter := router.Group("channel")
	{
		channelRouter.GET("list", controller.GetChannelListHandler)
	}
}

// TODO 페이먼트 히스토리 추가할 것인지?
func RegisterHistoryChannelRouter(router *gin.Engine){

	channelRouter := router.Group("channel")
	{
		channelRouter.GET("list", controller.GetChannelListHandler)
	}
}





