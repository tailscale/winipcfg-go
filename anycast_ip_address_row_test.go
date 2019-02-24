/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const anycastIpAddressRow_print = true

func TestGetAnycastIpAddressRows(t *testing.T) {

	addresses, err := GetAnycastIpAddressRows(AF_UNSPEC)

	if err != nil {
		t.Errorf("GetAnycastIpAddressRows() returned an error: %v", err)
		return
	}

	if anycastIpAddressRow_print {
		for _, address := range addresses {
			fmt.Println("===================== ANYCAST ADDRESS OUTPUT START =====================")
			fmt.Println(address)
			fmt.Println("====================== ANYCAST ADDRESS OUTPUT END ======================")
		}
	}
}