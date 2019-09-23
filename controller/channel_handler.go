package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
	"strconv"
	"github.com/sslab-instapay/instapay-go-client/service"
					)

func OpenChannelHandler(ctx *gin.Context) {
	//channelName := context.PostForm("ch_name")
	otherAddress := ctx.PostForm("other_addr")
	deposit, _ := strconv.Atoi(ctx.PostForm("deposit"))

	service.SendOpenChannelTransaction(deposit, otherAddress)

	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

// TODO 데모 시나리오 이후 구현
func DepositChannelHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DirectPayChannelHandler(context *gin.Context) {
	//channelId := context.PostForm("ch_id")
	//amount := context.PostForm("amount")

	//conn, err := grpc.Dial(config.EthereumConfig["serverAddr"], grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	//
	//// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func CloseChannelHandler(ctx *gin.Context) {
	channelIdParam := ctx.PostForm("channelId")
	log.Println(channelIdParam)
	channelId, _ := strconv.Atoi(channelIdParam)
	log.Println(channelId)

	service.SendCloseChannelTransaction(int64(channelId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func PaymentToServerChannelHandler(ctx *gin.Context) {
	//
	//otherAddress := ctx.PostForm("addr")
	//amount, err := strconv.Atoi(ctx.PostForm("amount"))
	//if err != nil {
	//	log.Println(err)
	//}
	////
	//myAddress := config.GetAccountConfig()
	////TODO Server의 GRPC 호출
	//
	//connection, err := grpc.Dial(config.EthereumConfig["serverGrpcHost"]+":"+config.EthereumConfig["serverGrpcPort"], grpc.WithInsecure())
	//if err != nil {
	//	log.Println("did not connect: %v", err)
	//}
	//defer connection.Close()
	//client := pb.NewClientClient(connection)
	//
	//clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := client.SayHello(client_context, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func GetChannelListHandler(ctx *gin.Context) {

	channelList, err := repository.GetChannelList()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"channels": channelList,
	})
}
