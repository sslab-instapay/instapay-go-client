package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	PublicKeyAddress string
	PrivateKey       string
	Balance          float64
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
	MyDeposit     int                `bson:"myDeposit"`
	MyBalance     int                `bson:"myBalance"`
	OtherAddress  string             `bson:"otherAddress"`
	VersionNumber int                `bson:"versionNumber"`
	OtherIp       int                `bson:"otherIp"`
	OtherPort     int                `bson:"otherPort"`
}
