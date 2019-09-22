package model

import (
	"math/big"
)

type Account struct {
	PublicKeyAddress string
	PrivateKey       string
	Balance          big.Float
}

type ChannelStatus int

const (
	// 0, 1, 2, 3
	IDLE ChannelStatus = iota
	PRE_UPDATE
	POST_UPDATE
	CLOSED
)

type Channel struct {
	ChannelId     int64         `bson:"channelId"`
	ChannelName   string        `bson:"channelName"`
	Status        ChannelStatus `bson:"channelStatus"` //IDLE, PRE_UPDATE, POST_UPDATE
	MyAddress     string        `bson:"myAddress"`
	MyDeposit     float32       `bson:"myDeposit"`
	MyBalance     float32       `bson:"myBalance"`
	LockedBalance float32       `bson:"lockedBalance"`
	OtherAddress  string        `bson:"otherAddress"`
	VersionNumber int           `bson:"versionNumber"`
	OtherIp       int           `bson:"otherIp"`
	OtherPort     int           `bson:"otherPort"`
}
