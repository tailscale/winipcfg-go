/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const ipInterface_print = false

func TestGetIpInterfaces(t *testing.T) {

	ipifcs, err := GetIpInterfaces(AF_UNSPEC)

	if err != nil {
		t.Errorf("GetIpInterfaces() returned an error: %v", err)
		return
	}

	if ipifcs == nil || len(ipifcs) < 1 {
		t.Error("GetIpInterfaces() returned nil or an empty slice.")
		return
	}

	if ipInterface_print {
		for _, ipifc := range ipifcs {
			fmt.Println("====================== INTERFACE DATA OUTPUT START ======================")
			fmt.Println(ipifc)
			fmt.Println("======================= INTERFACE DATA OUTPUT END =======================")
		}
	}
}
