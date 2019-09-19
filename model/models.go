package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ChannelId     primitive.ObjectID `bson:"_id"`
	ChannelName   string             `bson:"channelName"`
	Status        ChannelStatus      `bson:"channelStatus"` //IDLE, PRE_UPDATE, POST_UPDATE
	MyAddress     string             `bson:"myAddress"`
	MyDeposit     float64            `bson:"myDeposit"`
	MyBalance     float64            `bson:"myBalance"`
	OtherAddress  string             `bson:"otherAddress"`
	VersionNumber int                `bson:"versionNumber"`
	OtherIp       int                `bson:"otherIp"`
	OtherPort     int                `bson:"otherPort"`
}
