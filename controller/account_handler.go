package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
	"strconv"
	"os"
)

func AccountInformationHandler(context *gin.Context)  {
	port, _ := strconv.Atoi(os.Getenv("port"))
	account := config.GetAccountConfig(port)


	context.JSON(http.StatusOK, gin.H{"account": account})
}
