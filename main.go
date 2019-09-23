package main

import (
				instapayGrpc "github.com/sslab-instapay/instapay-go-client/grpc"
	pb "github.com/sslab-instapay/instapay-go-client/proto"
	"net"
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-go-client/router"
	"os"
	"strconv"
)

func startGrpcServer(){
	grpcPort, err := strconv.Atoi(os.Getenv("grpc_port"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterClientServer(grpcServer, &instapayGrpc.ClientGrpc{})
	grpcServer.Serve(lis)
}

func startClientWebServer(){
	defaultRouter := gin.Default()
	defaultRouter.LoadHTMLGlob("templates/*")

	router.RegisterChannelRouter(defaultRouter)
	router.RegisterRestAccountRouter(defaultRouter)
	router.RegisterViewRouter(defaultRouter)

	defaultRouter.Run(":7777")
}

func main() {
	// os[1] os[2] 로 전역변수 셋팅.

	os.Setenv("port", "3001")
	os.Setenv("grpc_port", "50001")

	//go config.ListenContractEvent()
	//go startGrpcServer()
	startClientWebServer()

}