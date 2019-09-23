package service

import (
	"crypto/ecdsa"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strconv"
	instapay "github.com/sslab-instapay/instapay-go-client/contract"
	"github.com/sslab-instapay/instapay-go-client/config"
	"os"
	"context"
	"fmt"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math"
)

func SendOpenChannelTransaction(deposit int, otherAddress string) {

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}

	portNum, _ := strconv.Atoi(os.Getenv("port"))
	// loading instapay contract on the blockchain
	address := common.HexToAddress(config.GetAccountConfig(portNum).PublicKeyAddress)  // change to correct address
	instance, err := instapay.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// loading my public key, nonce and gas price
	privateKey, err := crypto.HexToECDSA(config.GetAccountConfig(portNum).PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// composing a transaction
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(deposit * 10000000000)) // in wei
	auth.GasLimit = uint64(2000000) // in units
	auth.GasPrice = gasPrice

	receiver := common.HexToAddress(otherAddress)

	tx, err := instance.CreateChannel(auth, receiver)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func SendCloseChannelTransaction(channelId int64){

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}

	portNum, _ := strconv.Atoi(os.Getenv("port"))
	// loading instapay contract on the blockchain
	address := common.HexToAddress(config.GetAccountConfig(portNum).PublicKeyAddress)  // change to correct address
	instance, err := instapay.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// loading my public key, nonce and gas price
	privateKey, err := crypto.HexToECDSA(config.GetAccountConfig(portNum).PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// composing a transaction
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(2000000)
	auth.GasPrice = gasPrice

	channel, err := repository.GetChannelById(channelId)

	otherBalance := channel.MyDeposit - channel.MyBalance

	//TODO 밸런스 단위 맞추기 및 소수점 어떻게 볼 것인지 논의
	tx, err := instance.CloseChannel(auth, big.NewInt(channelId), big.NewInt(int64(otherBalance)) , big.NewInt(int64(channel.MyBalance)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func ListenContractEvent() {
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(config.EthereumConfig["contractAddr"])

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
			var createChannelEvent = model.CreateChannelEvent{}
			var closeChannelEvent = model.CloseChannelEvent{}
			var ejectEvent = model.EjectEvent{}

			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				log.Print("CreateChannel Event Emission")
				HandleCreateChannelEvent(createChannelEvent)
			}

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				log.Print("CloseChannel Event Emission")
				HandleCloseChannelEvent(closeChannelEvent)
			}

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {

			}

		}
	}
}

func HandleCreateChannelEvent(event model.CreateChannelEvent) {

	// TODO 상대편의 port와 ip를 요청하는 grpc 서버 콜을 추가해야함.
	portNum, err := strconv.Atoi(os.Getenv("port"))

	if err != nil{
		log.Println(err)
	}

	if event.Receiver.String() == config.GetAccountConfig(portNum).PublicKeyAddress {
		var channel = model.Channel{ChannelId: event.Id.Int64(), ChannelName: "Random",
			Status: model.IDLE, MyAddress: event.Receiver.String(),
			MyBalance: 0, MyDeposit: 0, OtherAddress: event.Owner.String()}

		repository.InsertChannel(channel)
	}else{
		var channel = model.Channel{ChannelId: event.Id.Int64(), ChannelName: "Random",
			Status: model.IDLE, MyAddress: event.Receiver.String(),
			MyBalance: event.Deposit.Int64(), MyDeposit: 0, OtherAddress: event.Owner.String()}
		repository.InsertChannel(channel)
	}



}

func HandleCloseChannelEvent(event model.CloseChannelEvent) {
	channel, err := repository.GetChannelById(event.Id.Int64())

	if err != nil {
		log.Println("there is no channel")
		return
	}

	channel.Status = model.CLOSED
	repository.UpdateChannel(channel)
}

func HandleEjectEvent(event model.EjectEvent) {
	//TODO
}

func GetBalance() big.Float {

	port, _ := strconv.Atoi(os.Getenv("port"))
	account := common.HexToAddress(config.GetAccountConfig(port).PublicKeyAddress)
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])

	if err != nil {
		log.Println(err)
	}

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Println(err)
	}
	log.Println(balance)

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return *ethValue
}

