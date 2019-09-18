package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/controller"
)


func RegisterRequestChannelRouter(router *gin.Engine){

	channelRouter := router.Group("requests/channels")
	{
		channelRouter.POST("open", controller.OpenChannelHandler)

		//TODO 데모 이후 추가.
		//channelRouter.POST("deposit", func(context *gin.Context) {
		//		//	context.JSON(http.StatusOK, controller.DepositChannelHandler)
		//		//})

		channelRouter.POST("payDirect", func(context *gin.Context) {
			context.JSON(http.StatusOK, controller.DirectChannelHandler)
		})

		channelRouter.POST("close", controller.CloseChannelHandler)

		channelRouter.POST("payServer", controller.PaymentToServerChannelHandler)
	}
}





