package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
		"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
		"strconv"
	"github.com/sslab-instapay/instapay-go-client/service"
)

func OpenChannelHandler(context *gin.Context)  {
	//channelName := context.PostForm("ch_name")
	otherAddress := context.PostForm("other_addr")
	deposit, _ := strconv.Atoi(context.PostForm("deposit"))

	service.SendOpenChannelTransaction(deposit, otherAddress)

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

// TODO 데모 시나리오 이후 구현
func DepositChannelHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DirectPayChannelHandler(context *gin.Context) {
	//channelId := context.PostForm("ch_id")
	//amount := context.PostForm("amount")
	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func CloseChannelHandler(context *gin.Context) {
	channelIdParam := context.PostForm("channelId")
	channelId, _ := strconv.Atoi(channelIdParam)

	service.SendCloseChannelTransaction(int64(channelId))

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func PaymentToServerChannelHandler(context *gin.Context) {

	//otherAddress := context.PostForm("addr")
	//amount, err := strconv.Atoi(context.PostForm("amount"))
	//TODO 주소정보 셋팅
	//address := config.GetAccountConfig(1111)
	//if err != nil {
	//	log.Fatal(err)
	//}

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func GetChannelListHandler(context *gin.Context){

	channelList, err := repository.GetChannelList()
	if err != nil {
		log.Fatal(channelList)
	}

	context.JSON(http.StatusOK, gin.H{
		"channels": channelList,
	})
}