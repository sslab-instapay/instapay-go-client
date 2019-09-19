package config

import (
	"github.com/sslab-instapay/instapay-go-client/model"
)

// TODO Enclave에서 포트에 따라 어카운트 주소 불러와야함
func GetAccountConfig(port int) model.Account {
	if port == 3001{

	}else if port == 3002{

	}else if port == 3003{

	}
	return model.Account{
		PublicKeyAddress: "0x12421t2tjgfiq",
		PrivateKey: "0xquotuinq2otnqwg",
		Balance: 10,
	}
}
