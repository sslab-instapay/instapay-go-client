package config_test

import (
	"testing"
	"github.com/sslab-instapay/instapay-go-client/config"
	"fmt"
)

func TestGetBalance(t *testing.T)  {

	service := config.EthereumService{}

	balance := service.GetBalance()

	fmt.Println(balance)

}
