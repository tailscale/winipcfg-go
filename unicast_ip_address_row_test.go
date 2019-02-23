/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const unicastAddressData_print = false

func TestGetUnicastAddresses(t *testing.T) {

	addresses, err := GetUnicastAddresses(AF_UNSPEC)

	if err != nil {
		t.Errorf("GetUnicastAddresses() returned an error: %v", err)
		return
	}

	if addresses == nil || len(addresses) < 1 {
		t.Error("GetUnicastAddresses() method returned nil or an empty slice.")
		return
	}

	if unicastAddressData_print {
		for _, address := range addresses {
			fmt.Println("===================== UNICAST ADDRESS OUTPUT START =====================")
			fmt.Println(address)
			fmt.Println("====================== UNICAST ADDRESS OUTPUT END ======================")
		}
	}
}
