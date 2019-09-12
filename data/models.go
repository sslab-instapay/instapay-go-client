package data

type Account struct {
	PublicKeyAddress string
	PrivateKey       string
}

type ChannelStatus int

const (
	READY ChannelStatus = iota
	IDLE
	LOCKED
	SIGNED
	CONFIRMED
	CRASHED
)

type Channel struct {
	ChannelName       string        `json:"channelName"`
	ChannelId         string        `json:"channelId"`
	Status            ChannelStatus `json:"channelStatus"` // IDLE, LOCKED, SIGNED, CONFIRMED, CRASHED
	MyAddress         string        `json:"myAddress"`
	MyDeposit         int           `json:"myDeposit"`
	MyBalance         int           `json:"myBalance"`
	OtherAddress      string        `json:"otherAddress"`
	OtherDeposit      int           `json:"otherDeposit"`
	OtherBalance      int           `json:"otherBalance"`
	VersionNumber     int           `json:"versionNumber"`
	CloseTimeInterval int           `json:"closeTimeInterval"`
	OtherPort         int           `json:"otherPort"`
}

type PendingClose struct {
	ChannelId     string
	MyBalance     int
	OtherBalance  int
	VersionNumber int
	CloseTimeout  int
}

type WaitingDeposit struct {
	ChannelName       string
	ChannelId         string
	MyAddress         string
	MyDeposit         int
	OtherAddress      string
	OtherDeposit      int
	CloseTimeInterval int
}
