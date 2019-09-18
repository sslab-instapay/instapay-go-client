package data_test

import (
	"github.com/instapay-go-client/repository"
	"fmt"
	"testing"
		"log"
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






