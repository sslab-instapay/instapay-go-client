package model

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	PublicKeyAddress string
	PrivateKey       string
	Balance          big.Float
}

type ChannelStatus string

const (
	// 0, 1, 2, 3
	IDLE        ChannelStatus = "IDLE"
	PRE_UPDATE                = "PRE_UPDATE"
	POST_UPDATE               = "POST_UPDATE"
	CLOSED                    = "CLOSED"
)

type PaymentData struct {
	PaymentNumber int64 `bson:"paymentNumber"`
	ChannelId     int64 `bson:"channelId"`
	Amount        int64 `bson:"amount"`
}

type Channel struct {
	ChannelId     int64         `bson:"channelId"`
	ChannelName   string        `bson:"channelName"`
	Status        ChannelStatus `bson:"channelStatus"`
	MyAddress     string        `bson:"myAddress"`
	MyDeposit     int64         `bson:"myDeposit"`
	MyBalance     int64         `bson:"myBalance"`
	LockedBalance int64         `bson:"lockedBalance"`
	OtherAddress  string        `bson:"otherAddress"`
	VersionNumber int           `bson:"versionNumber"`
	OtherIp       int           `bson:"otherIp"`
	OtherPort     int           `bson:"otherPort"`
}

type CreateChannelEvent struct {
	Id       *big.Int
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
}

type CloseChannelEvent struct {
	Id          *big.Int
	Ownerbal    *big.Int
	Receiverbal *big.Int
}

type EjectEvent struct {
	Pn              *big.Int
	Registeredstage int
}
