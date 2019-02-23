/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
	"testing"
)

const unicastAddressData_print = false

var existingUnicastIpAddress = net.IP{172, 16, 1, 110} // TODO: Ensure that actual existing IP address is set here.

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

func TestGetMatchingUnicastIpAddressRow(t *testing.T) {

	uar, err := GetMatchingUnicastIpAddressRow(&existingUnicastIpAddress)

	if err != nil {
		t.Errorf("GetMatchingUnicastIpAddressRow() returned an error: %v", err)
		return
	}

	if uar == nil {
		t.Errorf("Address %s not found. Have you forgot to set existingUnicastIpAddress appropriately?",
			existingUnicastIpAddress.String())
		return
	}

	if !uar.Address.Address.Equal(existingUnicastIpAddress) {
		t.Errorf("GetMatchingUnicastIpAddressRow() returned UnicastIpAddressRow with IP = %s, while IP = %s is expected.",
			uar.Address.Address.String(), existingUnicastIpAddress.String())
		return
	}

	if unicastAddressData_print {
		fmt.Println("===================== UNICAST ADDRESS OUTPUT START =====================")
		fmt.Println(uar)
		fmt.Println("====================== UNICAST ADDRESS OUTPUT END ======================")
	}
}
