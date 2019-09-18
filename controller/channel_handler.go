package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func OpenChannelHandler(context *gin.Context)  {
	//channelName := context.PostForm("ch_name")
	//myAddress := context.PostForm("my_addr")
	//otherAddress := context.PostForm("other_addr")
	//deposit := context.PostForm("deposit")

	// 여기서 채널 open

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DepositChannelHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DirectChannelHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func CloseChannelHandler(context *gin.Context) {
	context.PostForm("")
	//<td>{{ $channel.ChannelId }}</td>
	//<td>{{ $channel.MyAddress }}</td>
	//<td>{{ $channel.OtherAddress }}</td>
	//<td>{{ $channel.Status }}</td>

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func PaymentToServerChannelHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}