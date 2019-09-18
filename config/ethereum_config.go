package config

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
		"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/ethereum/go-ethereum"
)

var ethereumConfig = map[string]string{
	/* web3 and ethereum */
	"wsHost":           "localhost", //141.223.121.139
	"wsPort":           "8881",
	"contractAddr":     "0x164e52dD2A8a572f638A1f9EA5C02c2868499953",
	"contractSrcPath":  "../contracts/InstaPay.sol",
	"contractInstance": "",
	"web3":             "",
	"event":            "",

	/* grpc configuration */
	"serverGrpcHost": "localhost",
	"serverGrpcPort": "50004",
	"serverProto":    "",
	"server":         "",
	"myGrpcPort":     "", //process.argv[3]
	"clientProto":    "",
	"receiver":       "",
}

func GetContract(){
	client, err := ethclient.Dial(ethereumConfig["wsHort"] + ":" + ethereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(ethereumConfig["contractAddr"])
	// 모든 이벤트 수
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(""))) // ABI
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.Unpack(&event, "EventOpenChannelUD", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			// TODO 이벤트 처리
		}
	}
}
