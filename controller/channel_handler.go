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

func CloseChannelHandler(context *gin.Context) {
	channelIdParam := context.PostForm("channelId")
	log.Println(channelIdParam)
	channelId, _ := strconv.Atoi(channelIdParam)
	log.Println(channelId)

	service.SendCloseChannelTransaction(int64(channelId))

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func PaymentToServerChannelHandler(context *gin.Context) {

	//otherAddress := context.PostForm("addr")
	//amount, err := strconv.Atoi(context.PostForm("amount"))
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//myAddress := config.GetAccountConfig()
	//TODO Server의 GRPC 호출


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