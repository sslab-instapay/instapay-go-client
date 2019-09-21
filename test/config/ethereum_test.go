package config_test

import (
	"testing"
	"github.com/sslab-instapay/instapay-go-client/config"
	"fmt"
)

func TestGetBalance(t *testing.T)  {

	balance := config.GetBalance("0x78902c58006916201F65f52f7834e467877f0500")

	fmt.Println(balance)

}
