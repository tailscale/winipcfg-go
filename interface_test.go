/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

const (
	printInterfaceData      = true
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
	unexistentRouteIPv4ToAdd = RouteData{
		Destination: net.IPNet{
			IP: net.IP{172, 16, 200, 0},
			Mask: net.IPMask{255, 255, 255, 0},
		},
		NextHop: net.IP{172, 16, 1, 2},
		Metric: 0,
	}
	dnsesToSet = []net.IP{
		net.IPv4(8, 8, 8, 8),
		net.IPv4(8, 8, 4, 4),
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

func TestInterface_AddAddresses_DeleteAddress(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so add/delete address testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so add/delete address testing cannot be performed.")
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

	err = ifc.DeleteAddress(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.DeleteAddress() returned an error: %v", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	addr, err = ifc.GetMatchingUnicastAddressData(&unexistingIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.GetMatchingUnicastAddressData() returned an error: %v", err)
	} else if addr != nil {
		t.Errorf("Unicast address %s still exists, although it's deleted successfully.",
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

func TestInterface_AddRoute_DeleteRoute(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so add/delete route testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so add/delete route testing cannot be performed.")
		return
	}

	route, err := ifc.FindRoute(&unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
		return
	}

	if route != nil {
		t.Error("Interface.FindRoute() returned a route although it is not added yet. Have you forgot to set unexistentRouteIPv4ToAdd appropriately?")
		return
	}

	err = RegisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterRouteChangeCallback() returned an error: %v", err)
		return
	}

	defer func() {
		err := UnregisterRouteChangeCallback(&routeChangeCallbackExample)

		if err != nil {
			t.Errorf("UnregisterRouteChangeCallback() returned an error: %v", err)
		}
	}()

	err = ifc.AddRoute(&unexistentRouteIPv4ToAdd, true)

	if err != nil {
		t.Errorf("Interface.AddRoute() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	route, err = ifc.FindRoute(&unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
		return
	}

	if route == nil {
		t.Error("Interface.FindRoute() returned nil although a route is added successfully.")
	}

	err = ifc.DeleteRoute(&unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Iterface.DeleteRoute() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	route, err = ifc.FindRoute(&unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
		return
	}

	if route != nil {
		t.Error("Interface.FindRoute() returned the route although it's deleted successfully.")
	}
}

func TestInterface_AddRoute_DeleteRoute_SplitDefault(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so add/delete route testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so add/delete route testing cannot be performed.")
		return
	}

	routeToAdd := RouteData{
		Destination: net.IPNet{
			IP: net.IP{0, 0, 0, 0},
			Mask: net.IPMask{0, 0, 0, 0},
		},
		NextHop: net.IP{172, 16, 1, 2},
		Metric: 0,
	}

	expect1 := net.IPNet{
		IP: net.IP{0, 0, 0, 0},
		Mask: net.CIDRMask(1, 32),
	}

	expect2 := net.IPNet{
		IP: net.IP{128, 0, 0, 0},
		Mask: net.CIDRMask(1, 32),
	}

	route, err := ifc.FindRoute(&expect1)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
		return
	}

	if route != nil {
		t.Errorf("Route to %s already exists!", expect1.String())
		return
	}

	route, err = ifc.FindRoute(&expect2)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
		return
	}

	if route != nil {
		t.Errorf("Route to %s already exists!", expect2.String())
		return
	}

	err = RegisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterRouteChangeCallback() returned an error: %v", err)
		return
	}

	defer func() {
		err := UnregisterRouteChangeCallback(&routeChangeCallbackExample)

		if err != nil {
			t.Errorf("UnregisterRouteChangeCallback() returned an error: %v", err)
		}
	}()

	err = ifc.AddRoute(&routeToAdd, true)

	if err != nil {
		t.Errorf("Interface.AddRoute() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	route, err = ifc.FindRoute(&expect1)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
	} else if route == nil {
		t.Errorf("Route %s not found although it's added successfully", expect1.String())
	} else {

		err = ifc.DeleteRoute(&expect1)

		if err != nil {
			t.Errorf("Interface.DeleteRoute() returned an error: %v", err)
		}
	}

	route, err = ifc.FindRoute(&expect2)

	if err != nil {
		t.Errorf("Interface.FindRoute() returned an error: %v", err)
	} else if route == nil {
		t.Errorf("Route %s not found although it's added successfully", expect2.String())
	} else {

		err = ifc.DeleteRoute(&expect2)

		if err != nil {
			t.Errorf("Interface.DeleteRoute() returned an error: %v", err)
		}
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)
}

func TestInterface_GetNetworkAdapterConfiguration(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.GetNetworkAdapterConfiguration() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.GetNetworkAdapterConfiguration() testing cannot be performed.")
		return
	}

	nac, err := ifc.GetNetworkAdapterConfiguration()

	if err != nil {
		t.Errorf("Interface.GetNetworkAdapterConfiguration() returned an error: %v", err)
		return
	}

	if nac == nil {
		t.Error("Interface.GetNetworkAdapterConfiguration() returned nil")
		return
	}

	if strings.ToUpper(strings.TrimSpace(ifc.AdapterName)) != strings.ToUpper(strings.TrimSpace(nac.SettingID)) {
		t.Errorf("Interface.GetNetworkAdapterConfiguration() returned NetworkAdapterConfiguration.SettingID = %s, although Interface.AdapterName = %s.",
			nac.SettingID, ifc.AdapterName)
		return
	}

	if printNetworkAdaptersConfigurations {
		fmt.Println("============== NETWORK ADAPTER CONFIGURATION OUTPUT START ==============")
		fmt.Println(nac)
		fmt.Println("=============== NETWORK ADAPTER CONFIGURATION OUTPUT END ===============")
	}
}

func TestInterface_SetDNS(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.SetDNS() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.SetDNS() testing cannot be performed.")
		return
	}

	err = RegisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterInterfaceChangeCallback() returned an error: %v", err)
		return
	}

	defer func() {
		err := UnregisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

		if err != nil {
			t.Errorf("UnregisterInterfaceChangeCallback() returned an error: %v", err)
		}
	}()

	prevDnsesCount := 0

	if (ifc.DnsServerAddresses != nil) {
		prevDnsesCount = len(ifc.DnsServerAddresses)
	}

	prevDnses := make([]net.IP, prevDnsesCount, prevDnsesCount)

	for i := 0; i < prevDnsesCount; i++ {
		prevDnses[i] = ifc.DnsServerAddresses[i].Address.Address
	}

	err = ifc.SetDNS(dnsesToSet)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	err = ifc.Refresh()

	if err != nil {
		t.Errorf("Interface.Refresh() returned an error: %v", err)
	}

	if err != nil {
		t.Errorf("GetNetworkAdapterConfiguration() returned an error: %v", err)
	} else {

		if printInterfaceData {
			fmt.Println("======================== INTERFACE OUTPUT START ========================")
			fmt.Println(ifc)
			fmt.Println("========================= INTERFACE OUTPUT END =========================")
		}

		if dnsesToSet == nil {
			if ifc.DnsServerAddresses != nil && len(ifc.DnsServerAddresses) != 0 {
				t.Errorf("dnsesToSet is nil, but DnsServerAddresses contains %d items.",
					len(ifc.DnsServerAddresses))
			}
		} else {

			length := len(dnsesToSet)

			if ifc.DnsServerAddresses == nil {
				t.Errorf("dnsesToSet contains %d items, while DnsServerAddresses is nil.", length)
			} else if len(ifc.DnsServerAddresses) != length {
				t.Errorf("dnsesToSet contains %d items, while DnsServerAddresses contains %d.", length,
					len(ifc.DnsServerAddresses))
			} else {
				for idx, dns := range dnsesToSet {
					if !dns.Equal(ifc.DnsServerAddresses[idx].Address.Address) {
						t.Errorf("dnsesToSet[%d] = %s while DnsServerAddresses[%d].Address.Address = %s.", idx,
							dns.String(), idx, ifc.DnsServerAddresses[idx].Address.Address.String())
					}
				}
			}
		}
	}

	err = ifc.SetDNS(prevDnses)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v.", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)
}

func TestInterface_FlushDNS(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.FlushDNS() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.FlushDNS() testing cannot be performed.")
		return
	}

	prevDnsesCount := 0

	if (ifc.DnsServerAddresses != nil) {
		prevDnsesCount = len(ifc.DnsServerAddresses)
	}

	prevDnses := make([]net.IP, prevDnsesCount, prevDnsesCount)

	for i := 0; i < prevDnsesCount; i++ {
		prevDnses[i] = ifc.DnsServerAddresses[i].Address.Address
	}

	err = ifc.FlushDNS()

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v", err)
		return
	}

	err = ifc.Refresh()

	if err != nil {
		t.Errorf("Interface.Refresh() returned an error: %v", err)
	}

	if err != nil {
		t.Errorf("GetNetworkAdapterConfiguration() returned an error: %v", err)
	} else {

		if printInterfaceData {
			fmt.Println("======================== INTERFACE OUTPUT START ========================")
			fmt.Println(ifc)
			fmt.Println("========================= INTERFACE OUTPUT END =========================")
		}

		if ifc.DnsServerAddresses != nil && len(ifc.DnsServerAddresses) != 0 {
			t.Errorf("DnsServerAddresses contains %d items, although FlushDNS is executed successfully.",
				len(ifc.DnsServerAddresses))
		}
	}

	err = ifc.SetDNS(prevDnses)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v.", err)
	}
}
