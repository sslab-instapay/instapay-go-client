package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
	"github.com/sslab-instapay/instapay-go-client/service"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"fmt"
	"strconv"
)

func AccountInformationHandler(context *gin.Context) {
	account := config.GetAccountConfig()
	balance, _ := service.GetBalance()

	lockedBalance, err := repository.GetAllChannelsLockedBalance()
	if err != nil {
		log.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		convertedBalance, _ := balance.Int64()
		totalBalance := convertedBalance - lockedBalance
		context.JSON(http.StatusOK, gin.H{"address": account.PublicKeyAddress, "balance": totalBalance})
	}

}

func OnchainPaymentHandler(ctx *gin.Context){
	account := config.GetAccountConfig()
	amount, err := strconv.Atoi(ctx.PostForm("amount"))
	otherAddress := ctx.PostForm("address")
	if err != nil{
		log.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		log.Println(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}
	value := big.NewInt(int64(amount) * 1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}
	toAddress := common.HexToAddress(otherAddress)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Println(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	ctx.JSON(http.StatusOK, gin.H{"message": "payment Success"})
}
