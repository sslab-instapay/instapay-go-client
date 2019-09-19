package data_test

import (
	"fmt"
	"testing"
	"log"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetChannelList(t *testing.T){
	channelList, err := repository.GetChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetClosedChannelList(t *testing.T){
	channelList, err := repository.GetClosedChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetOpenedChannelList(t *testing.T){
	channelList, err := repository.GetOpenedChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetChannelById(t *testing.T){
	objectId, _ := primitive.ObjectIDFromHex("5d7fb9669f65573e75071f97")
	channel, err := repository.GetChannelById(objectId)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(channel)
}






