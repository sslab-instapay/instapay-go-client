package config

var etheruemConfig = map[string]string{
	/* web3 and ethereum */
	"wsHost":           "localhost", //141.223.121.139
	"wsPort":           "8881",
	"contractAddr":     "0x164e52dD2A8a572f638A1f9EA5C02c2868499953",
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

func GetContract(){
}