package service

import (
	"crypto/ecdsa"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-go-client/contract"
	serverPb "github.com/sslab-instapay/instapay-go-client/proto/server"
	"github.com/sslab-instapay/instapay-go-client/config"
	"context"
	"fmt"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math"
	"google.golang.org/grpc"
	"time"
)

func SendOpenChannelTransaction(deposit int, otherAddress string) (string, error) {

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Println(err)
		return "", err
	}

	// loading instapay contract on the blockchain
	address := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress) // change to correct address
	instance, err := instapay.NewContract(address, client)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// loading my public key, nonce and gas price
	privateKey, err := crypto.HexToECDSA(config.GetAccountConfig().PrivateKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return "", err
	}

	// composing a transaction
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(deposit * 10000000000)) // in wei
	auth.GasLimit = uint64(2000000)                       // in units
	auth.GasPrice = gasPrice

	receiver := common.HexToAddress(otherAddress)

	tx, err := instance.CreateChannel(auth, receiver)
	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func SendCloseChannelTransaction(channelId int64) {

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Println(err)
	}

	// loading instapay contract on the blockchain
	address := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress) // change to correct address
	instance, err := instapay.NewContract(address, client)
	if err != nil {
		log.Println(err)
	}

	// loading my public key, nonce and gas price
	privateKey, err := crypto.HexToECDSA(config.GetAccountConfig().PrivateKey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}

	// composing a transaction
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(2000000)
	auth.GasPrice = gasPrice

	channel, err := repository.GetChannelById(channelId)
	//TODO other Balance MyBalance 계산
	otherBalance := channel.MyDeposit - channel.MyBalance
	tx, err := instance.CloseChannel(auth, big.NewInt(channelId), big.NewInt(int64(otherBalance)), big.NewInt(int64(channel.MyBalance)))
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func ListenContractEvent() {
	log.Println("---Start Listen Contract Event---")
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Println(err)
	}
	contractAddress := common.HexToAddress(config.EthereumConfig["contractAddr"])

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(instapay.ContractABI)))
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case vLog := <-logs:
			var createChannelEvent = model.CreateChannelEvent{}
			var closeChannelEvent = model.CloseChannelEvent{}
			var ejectEvent = model.EjectEvent{}

			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				log.Println("CreateChannel Event Emission")
				fmt.Printf("Channel ID       : %d\n", createChannelEvent.Id)
				fmt.Printf("Channel Onwer    : %s\n", createChannelEvent.Owner.Hex())
				fmt.Printf("Channel Receiver : %s\n", createChannelEvent.Receiver.Hex())
				fmt.Printf("Channel Deposit  : %d\n", createChannelEvent.Deposit)
				HandleCreateChannelEvent(createChannelEvent)
			}

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				log.Print("CloseChannel Event Emission")
				fmt.Printf("Channel ID       : %d\n", closeChannelEvent.Id)
				fmt.Printf("Owner Balance    : %d\n", closeChannelEvent.Ownerbal)
				fmt.Printf("Receiver Balance : %d\n", closeChannelEvent.Receiverbal)
				HandleCloseChannelEvent(closeChannelEvent)
			}

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {
				fmt.Printf("Payment Number   : %d\n", ejectEvent.Pn)
				fmt.Printf("Stage            : %d\n", ejectEvent.Registeredstage)
			}

		}
	}
}

func HandleCreateChannelEvent(event model.CreateChannelEvent) error{

	// 내가 리시버 즉 IN 채널
	log.Println("----- Handle Create Channel Event ----")
	if event.Receiver.String() == config.GetAccountConfig().PublicKeyAddress {
		var channel = model.Channel{ChannelId: event.Id.Int64(), Type: model.IN,
			Status: model.IDLE, MyAddress: event.Receiver.String(),
			MyBalance: 0, MyDeposit: 0, OtherDeposit: event.Deposit.Int64(), OtherAddress: event.Owner.String()}
		repository.InsertChannel(channel)
	} else if event.Owner.String() == config.GetAccountConfig().PublicKeyAddress {
		// 아웃 채널
		var channel = model.Channel{ChannelId: event.Id.Int64(), Type: model.OUT,
			Status: model.IDLE, MyAddress: event.Owner.String(),
			MyBalance: event.Deposit.Int64(), MyDeposit: event.Deposit.Int64(), OtherDeposit: 0, OtherAddress: event.Receiver.String()}
		repository.InsertChannel(channel)
	}

	myAddress := config.GetAccountConfig().PublicKeyAddress
	connection, err := grpc.Dial(config.EthereumConfig["serverGrpcHost"] + ":" + config.EthereumConfig["serverGrpcPort"], grpc.WithInsecure())
	if err != nil {
		log.Println("GRPC Connection Error")
		log.Println(err)
		return err
	}
	defer connection.Close()
	client := serverPb.NewServerClient(connection)

	clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CommunicationInfoRequest(clientContext, &serverPb.Address{Addr: myAddress})
	if err != nil {
		log.Println("GRPC Request Error")
		log.Println(err)
		return err
	}

	channel, err := repository.GetChannelById(event.Id.Int64())
	if err != nil{
		log.Println(err)
		return err
	}

	log.Println(r.Port)
	log.Println(r.IPAddress)
	channel.OtherPort = int(r.Port)
	channel.OtherIp = r.IPAddress
	_, err = repository.UpdateChannel(channel)
	if err != nil{
		log.Println(err)
		return err
	}
	log.Println("----- Handle Create Channel Event END ----")
	return nil
}

func HandleCloseChannelEvent(event model.CloseChannelEvent) {
	channel, err := repository.GetChannelById(event.Id.Int64())

	if err != nil {
		log.Println("there is no channel")
	}

	channel.Status = model.CLOSED
	repository.UpdateChannel(channel)
}

func HandleEjectEvent(event model.EjectEvent) {
	//TODO
}

func GetBalance() big.Float {

	account := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress)
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
