package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
)

func AccountInformationHandler(context *gin.Context)  {
	account := config.GetAccountConfig(1111)

	context.JSON(http.StatusOK, gin.H{"account": account})
}
