package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
	"strconv"
	"os"
	"github.com/sslab-instapay/instapay-go-client/service"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
)

func AccountInformationHandler(context *gin.Context) {
	port, _ := strconv.Atoi(os.Getenv("port"))
	account := config.GetAccountConfig(port)
	balance := service.GetBalance()

	lockedBalance, err := repository.GetAllChannelsLockedBalance()
	if err != nil{
		log.Print(err)
	}

	convertedBalance, _ := balance.Float32()

	totalBalance := convertedBalance - lockedBalance

	context.JSON(http.StatusOK, gin.H{"address": account.PublicKeyAddress, "balance": totalBalance})
}
