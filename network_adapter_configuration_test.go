package winipcfg

import (
	"fmt"
	"testing"
)

const printNetworkAdaptersConfigurations = true

func TestGetNetworkAdaptersConfigurations(t *testing.T) {

	nacs, err := GetNetworkAdaptersConfigurations()

	if err != nil {
		t.Errorf("GetNetworkAdaptersConfigurations() returned an error: %v", err)
		return
	}

	if printNetworkAdaptersConfigurations {
		for _, nac := range nacs {
			fmt.Println("============== NETWORK ADAPTER CONFIGURATION OUTPUT START ==============")
			fmt.Println(nac)
			fmt.Println("=============== NETWORK ADAPTER CONFIGURATION OUTPUT END ===============")
		}
	}
}
