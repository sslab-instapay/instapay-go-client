package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/instapay-go-client/controller"
)

// 라우터 등록 코드
func RegisterAccountRouter(router *gin.Engine){

	accountRouter := router.Group("account")
	{
		accountRouter.GET("list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "account"})
		})

	}
}

func RegisterChannelRouter(router *gin.Engine){

	channelRouter := router.Group("channels")
	{
		channelRouter.POST("open", controller.OpenChannelHandler)

		//TODO 데모 이후 추가.
		//channelRouter.POST("deposit", func(context *gin.Context) {
		//		//	context.JSON(http.StatusOK, controller.DepositChannelHandler)
		//		//})

		channelRouter.POST("direct", func(context *gin.Context) {
			context.JSON(http.StatusOK, controller.DirectChannelHandler)
		})

		channelRouter.POST("close", controller.CloseChannelHandler)

		channelRouter.POST("server", controller.PaymentToServerChannelHandler)
	}
}





