/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
	"testing"
	"time"
)

const (
	printInterfaceData      = false
	existingLuid            = uint64(1689399632855040) // TODO: Set an existing LUID here
	unexistingLuid          = uint64(42)
	existingIndex           = uint32(13) // TODO: Set an existing interface index here
	unexistingIndex         = uint32(42000000)
	existingInterfaceName   = "LAN" // TODO: Set an existing interface name here
	unexistingInterfaceName = "NON-EXISTING-NAME"
	printInterfaceRoutes    = false
)

var (
	unexistingIpAddresToAdd = net.IPNet{
		IP:   net.IP{172, 16, 1, 114},
		Mask: net.IPMask{255, 255, 255, 0},
	}
)

func TestGetInterfaces(t *testing.T) {

	ifcs, err := GetInterfaces()

	if err != nil {
		t.Errorf("GetInterfaces() returned error: %v", err)
	} else if ifcs == nil {
		t.Errorf("GetInterfaces() returned nil.")
	} else if printInterfaceData {
		fmt.Printf("GetInterfaces() returned %d items:\n", len(ifcs))
		for _, ifc := range ifcs {
			fmt.Println("======================== INTERFACE OUTPUT START ========================")
			fmt.Println(ifc)
			fmt.Println("========================= INTERFACE OUTPUT END =========================")
		}
	}
}

func TestInterfaceFromLUIDExisting(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromLUID() returned nil for luid=%d. Have you set existingLuid constant?",
			existingLuid)
	} else if ifc.Luid != existingLuid {
		t.Errorf("InterfaceFromLUID() returned interface with a wrong LUID. Requested: %d; returned: %d.",
			existingLuid, ifc.Luid)
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromLUID() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromLUIDNonExisting(t *testing.T) {

	ifc, err := InterfaceFromLUID(unexistingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromLUID() returned an interface with LUID=%d, although requested LUID was %d.",
			ifc.Luid, unexistingLuid)
	}
}

func TestInterfaceFromIndexExisting(t *testing.T) {

	ifc, err := InterfaceFromIndex(existingIndex)

	if err != nil {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromIndex() returned nil for index=%d. Have you set existingIndex constant?",
			existingIndex)
	} else if uint32(ifc.Index) != existingIndex {
		t.Errorf("InterfaceFromIndex() returned interface with a wrong index. Requested: %d; returned: %d.",
			existingIndex, ifc.Index)
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromIndex() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromIndexNonExisting(t *testing.T) {

	ifc, err := InterfaceFromIndex(unexistingIndex)

	if err != nil {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromIndex() returned an interface with index=%d, although requested index was %d.",
			ifc.Index, unexistingIndex)
	}
}

func TestInterfaceFromFriendlyNameExisting(t *testing.T) {

	ifc, err := InterfaceFromFriendlyName(existingInterfaceName)

	if err != nil {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromFriendlyName() returned nil for name=%s. Have you set existingInterfaceName constant?",
			existingInterfaceName)
	} else if ifc.FriendlyName != existingInterfaceName {
		t.Errorf("InterfaceFromFriendlyName() returned interface with a wrong name. Requested: %s; returned: %s.",
			existingInterfaceName, ifc.FriendlyName)
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromFriendlyName() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromFriendlyNameNonExisting(t *testing.T) {

	ifc, err := InterfaceFromFriendlyName(unexistingInterfaceName)

	if err != nil {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromFriendlyName() returned an interface with name=%s, although requested name was %s.",
			ifc.FriendlyName, unexistingInterfaceName)
	}
}

func TestInterface_AddAddresses_RemoveAddress(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so add/remove address testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so add/remove address testing cannot be performed.")
		return
	}

	addr, err := ifc.GetMatchingUnicastAddressData(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.GetMatchingUnicastAddressData() returned an error: %v", err)
		return
	}

	if addr != nil {
		t.Errorf("Unicast address %s already exists. Please set unexistingIpAddresToAdd appropriately.",
			unexistingIpAddresToAdd.IP.String())
		return
	}

	err = RegisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)
	defer UnregisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterUnicastAddressChangeCallback() returned an error: %v", err)
	}

	count := len(ifc.UnicastAddresses)

	err = ifc.AddAddresses([]*net.IPNet{&unexistingIpAddresToAdd})

	if err != nil {
		t.Errorf("Interface.AddAddresses() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	if count+1 != len(ifc.UnicastAddresses) {
		t.Errorf("Number of unicast addresses before adding is %d, while number after adding is %d.", count,
			len(ifc.UnicastAddresses))
	}

	addr, err = ifc.GetMatchingUnicastAddressData(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.GetMatchingUnicastAddressData() returned an error: %v", err)
	} else if addr == nil {
		t.Errorf("Unicast address %s still doesn't exist, although it's added successfully.",
			unexistingIpAddresToAdd.IP.String())
	}

	err = ifc.RemoveAddress(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.RemoveAddress() returned an error: %v", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	addr, err = ifc.GetMatchingUnicastAddressData(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.GetMatchingUnicastAddressData() returned an error: %v", err)
	} else if addr != nil {
		t.Errorf("Unicast address %s still exists, although it's removed successfully.",
			unexistingIpAddresToAdd.IP.String())
	}
}

func TestInterface_GetRoutes(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.GetRoutes() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.GetRoutes() testing cannot be performed.")
		return
	}

	routes, err := ifc.GetRoutes(AF_UNSPEC)

	if err != nil {
		t.Errorf("Interface.GetRoutes() returned an error: %v", err)
		return
	}

	if routes == nil || len(routes) < 1 {
		t.Error("Interface.GetRoutes() returned nil or empty slice.")
		return
	}

	for _, route := range routes {
		if route.InterfaceLuid != ifc.Luid {
			t.Errorf("Interface.GetRoutes() retuned a route with a wrong LUID. Interface.Luid: %d; Route.InterfaceLuid: %d.",
				ifc.Luid, route.InterfaceLuid)
		}
	}

	if printInterfaceRoutes {
		for _, route := range routes {
			fmt.Println("========================== ROUTE OUTPUT START ==========================")
			fmt.Println(route)
			fmt.Println("=========================== ROUTE OUTPUT END ===========================")
		}
	}
}
