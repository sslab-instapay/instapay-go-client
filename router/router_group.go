package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	channelRouter := router.Group("channel")
	{
		channelRouter.GET("list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "Channel"})
		})

	}
}


func RegisterViewRouter(router *gin.Engine){

	channelRouter := router.Group("view")
	{
		channelRouter.GET("list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "Channel"})
		})

	}
}




