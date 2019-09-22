package config

import (
	"github.com/sslab-instapay/instapay-go-client/model"
)

// TODO Enclave에서 포트에 따라 어카운트 주소 불러와야함
func GetAccountConfig(port string) model.Account {
	if port == "3001"{
		return model.Account{
			PublicKeyAddress: "0x78902c58006916201F65f52f7834e467877f0500",
			PrivateKey: "0xquotuinq2otnqwg",
		}
	}else if port == "3002"{
		return model.Account{
			PublicKeyAddress: "0x0b4161ad4f49781a821C308D672E6c669139843C",
			PrivateKey: "0xquotuinq2otnqwg",
		}

	}else if port == "3003"{
		return model.Account{
			PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
			PrivateKey: "0xquotuinq2otnqwg",
		}
	}
	return model.Account{
		PublicKeyAddress: "0x78902c58006916201F65f52f7834e467877f0500",
		PrivateKey: "0xquotuinq2otnqwg",
	}
}
