package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-go-client/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math/big"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"fmt"
)

var EthereumConfig = map[string]string{
	/* web3 and ethereum */
	"rpcHost":          "141.223.121.139",
	"rpcPort":          "8555",
	"wsHost":           "141.223.121.139", //141.223.121.139
	"wsPort":           "8881",
	"contractAddr":     "0x3016947BE73dcb877401Ee33802aC8fA6feE631E", // change to correct address
	"contractSrcPath":  "../contracts/InstaPay.sol",
	"contractInstance": "",
	"web3":             "",
	"event":            "",

	/* grpc configuration */
	"serverGrpcHost": "141.223.121.139",
	"serverGrpcPort": "50004",
	"serverProto":    "",
	"server":         "",
	"myGrpcPort":     "", //process.argv[3]
	"clientProto":    "",
	"receiver":       "",
}

type CreateChannelEvent struct {
	Id       *big.Int
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
}

type CloseChannelEvent struct {
	Id              *big.Int
	OwnerBalance    *big.Int
	ReceiverBalance *big.Int
}

type EjectEvent struct {
	PaymentNum *big.Int
	Stage      int
}

func main() {
	client, err := ethclient.Dial("ws://" + EthereumConfig["wsHost"] + ":" + EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(EthereumConfig["contractAddr"])

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(instapay.ContractABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			var createChannelEvent = CreateChannelEvent{}
			var closeChannelEvent = CloseChannelEvent{}
			var ejectEvent = EjectEvent{}

			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				log.Print("CreateChannel Event Emission")
				fmt.Print(createChannelEvent)
				HandleCreateChannelEvent(createChannelEvent)
			}
			log.Println(err)

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				log.Print("CloseChannel Event Emission")
				HandleCloseChannelEvent(closeChannelEvent)
			}
			log.Println(err)

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {

			}
			log.Println(err)

		}
	}
}

func HandleCreateChannelEvent(event CreateChannelEvent) {

	var channel = model.Channel{ChannelId: event.Id.Int64(), ChannelName: "Random",
		Status: model.IDLE, MyAddress: event.Receiver.String(),
		MyBalance: 0, MyDeposit: 0, OtherAddress: event.Owner.String() }

	repository.InsertChannel(channel)
}

func HandleCloseChannelEvent(event CloseChannelEvent) {
	channel, err := repository.GetChannelById(event.Id.Int64())

	if err != nil {
		log.Println("there is no channel")
		return
	}

	channel.Status = model.CLOSED
	repository.UpdateChannel(channel)
}
