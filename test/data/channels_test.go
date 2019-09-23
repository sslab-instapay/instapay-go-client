package data_test

import (
	"fmt"
	"testing"
	"log"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/config"
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
	channel, err := repository.GetChannelById(2)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(channel)
}

func TestUpdateChannel(t *testing.T){
	channel, err := repository.GetChannelById(2)
	if err != nil{
		log.Fatal(err)
	}
	
	channel.OtherPort = 3002
	updatedChannel, err := repository.UpdateChannel(channel)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedChannel)
}

func TestInsertChannel(t *testing.T){
	channel := model.Channel{ChannelId: 1, ChannelName: "hoonki", Status: model.IDLE, MyAddress: config.GetAccountConfig(3001).PublicKeyAddress,}

	insertedChannel, err := repository.InsertChannel(channel)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(insertedChannel)
}

func TestGetAllChannelsLockedBalance(t *testing.T){

	lockedBalance, err := repository.GetAllChannelsLockedBalance()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(lockedBalance)
}

func TestGetPaymentDataByPaymentId(t *testing.T){

	paymentData, err := repository.GetPaymentDatasByPaymentId(1)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(paymentData)
}




