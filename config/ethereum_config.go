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
		"math/big"
	"math"
	)

var EthereumConfig = map[string]string{
	/* web3 and ethereum */
	"wsHost":           "141.223.121.139",
	"wsPort":           "8881",
	"contractAddr":     "0x092d70BB5c1954F5Fa3EBbb282d0416a5e46c818",
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

func ListenContractEvent(){
	client, err := ethclient.Dial("ws://" + EthereumConfig["wsHost"] + ":" + EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(EthereumConfig["contractAddr"])
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
			//event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit);
			//    event EventCloseChannel(uint256 id, uint256 owner_bal, uint256 receiver_bal);
			//    event EventEject(uint256 pn, Stage registered_stage);
		}
	}
}


func GetBalance(address string) big.Float {

	account := common.HexToAddress(address)
	client, err := ethclient.Dial("ws://" + EthereumConfig["wsHost"] + ":" + EthereumConfig["wsPort"])

	if err != nil{
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil{
		log.Fatal(err)
	}

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return *ethValue
}

func SendPaymentTransaction(){

	//client, err := ethclient.Dial(EthereumConfig["wsHort"] + ":" + EthereumConfig["wsPort"])
	//gasLimit := uint64(21000)
	//
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//value := big.NewInt(1000000000000000000)
	//
	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//var data []byte
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
