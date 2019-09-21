package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-go-client/router"
	)

func main() {
	// os[1] os[2] 로 전역변수 셋팅.
	defaultRouter := gin.Default()
	defaultRouter.LoadHTMLGlob("templates/*")

	router.RegisterChannelRouter(defaultRouter)
	router.RegisterRestAccountRouter(defaultRouter)
	router.RegisterViewRouter(defaultRouter)



	defaultRouter.Run(":7777")

}
