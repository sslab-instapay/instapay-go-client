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
	instapay "github.com/sslab-instapay/instapay-go-client/contract"
	"math"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"strconv"
	"os"
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

func ListenContractEvent() {
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

func HandleCreateChannelEvent(event CreateChannelEvent) {

	// TODO 상대편의 port와 ip를 요청하는 grpc 서버 콜을 추가해야함.
	var channel = model.Channel{ChannelId: event.Id.Int64(), ChannelName: "Random",
		Status: model.IDLE, MyAddress: event.Receiver.String(),
		MyBalance: 0, MyDeposit: 0, OtherAddress: event.Owner.String()}

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

func HandleEjectEvent(event EjectEvent) {
	//TODO
}

func GetBalance() big.Float {
	port, _ := strconv.Atoi(os.Getenv("port"))
	account := common.HexToAddress(GetAccountConfig(port).PublicKeyAddress)
	client, err := ethclient.Dial("ws://" + EthereumConfig["wsHost"] + ":" + EthereumConfig["wsPort"])

	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return *ethValue
}

func SendOpenChannelTransaction() {

	//account := GetAccountConfig(service.Port)
	//privateKey := common.HexToAddress(account.PrivateKey)
	//
	//client, err := ethclient.Dial(EthereumConfig["wsHort"] + ":" + EthereumConfig["wsPort"])
	//gasLimit := uint64(21000)
	//
	//privKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey := privKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//auth := bind.NewKeyedTransactor(privKey)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	//
	//address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	//instance, err := store.NewStore(address, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))
	//tx, err := instance.SetItem(auth, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b2 7ad307bdfb97a1dca51d0f3035dcee3e870
	//result, err := instance.Items(nil, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(result[:])) // "bar"
}
