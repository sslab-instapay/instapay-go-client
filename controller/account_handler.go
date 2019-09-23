package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
	"strconv"
	"os"
	"github.com/sslab-instapay/instapay-go-client/service"
)

func AccountInformationHandler(context *gin.Context)  {
	port, _ := strconv.Atoi(os.Getenv("port"))
	account := config.GetAccountConfig(port)

	account.PrivateKey = ""
	account.Balance = service.GetBalance()


	context.JSON(http.StatusOK, gin.H{"account": account})
}
