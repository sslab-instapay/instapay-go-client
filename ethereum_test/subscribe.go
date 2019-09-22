package main

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/ethereum/go-ethereum"
	"math/big"
		instapay "github.com/sslab-instapay/instapay-go-client/contract"
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
			createChannelEvent := struct {
				Id       *big.Int
				Owner    common.Address
				Receiver common.Address
				Deposit  *big.Int
			}{}
			closeChannelEvent := struct {
				Id              *big.Int
				OwnerBalance    *big.Int
				ReceiverBalance *big.Int
			}{}

			ejectEvent := struct {
				PaymentNum *big.Int
				Stage      int
			}{}

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				fmt.Printf("Channel ID       : %d\n", closeChannelEvent.Id)
				fmt.Printf("Channel Onwer    : %s\n", closeChannelEvent.OwnerBalance)
				fmt.Printf("Channel Receiver : %s\n", closeChannelEvent.ReceiverBalance)
			}

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {
				fmt.Println("HI")
				fmt.Printf("Channel ID       : %d\n", ejectEvent.PaymentNum.Int64())
				fmt.Printf("Channel Onwer    : %d\n", ejectEvent.Stage)
			}

			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				fmt.Printf("Channel ID       : %d\n", createChannelEvent.Id)
				fmt.Printf("Channel Onwer    : %s\n", createChannelEvent.Owner.Hex())
				fmt.Printf("Channel Receiver : %s\n", createChannelEvent.Receiver.Hex())
				fmt.Printf("Channel Deposit  : %d\n", createChannelEvent.Deposit)
			}

		}
	}
}
