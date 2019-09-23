package config_test

import (
	"testing"
		"fmt"
	"github.com/sslab-instapay/instapay-go-client/service"
)

func TestGetBalance(t *testing.T)  {

	balance := service.GetBalance()

	fmt.Println(balance)

}
