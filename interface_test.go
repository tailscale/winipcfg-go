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
	interface_print                              = false
	interface_printIpInterfaces                  = false
	interface_printRoutes                        = false
	interface_printNetworkAdaptersConfigurations = false
	existingLuid                                 = uint64(1689399632855040) // TODO: Set an existing LUID here
	unexistingLuid                               = uint64(42)
	existingIndex                                = uint32(13) // TODO: Set an existing interface index here
	unexistingIndex                              = uint32(42000000)
	existingInterfaceName                        = "LAN" // TODO: Set an existing interface name here
	unexistingInterfaceName                      = "NON-EXISTING-NAME"
)

var (
	unexistentIpAddresToAdd = net.IPNet{
		IP:   net.IP{172, 16, 1, 114},
		Mask: net.IPMask{255, 255, 255, 0},
	}
	unexistentRouteIPv4ToAdd = RouteData{
		Destination: net.IPNet{
			IP:   net.IP{172, 16, 200, 0},
			Mask: net.IPMask{255, 255, 255, 0},
		},
		NextHop: net.IP{172, 16, 1, 2},
		Metric:  0,
	}
	dnsesToSet = []net.IP{
		net.IPv4(8, 8, 8, 8),
		net.IPv4(8, 8, 4, 4),
	}
)

func unicastAddressChangeCallbackExample(notificationType MibNotificationType, interfaceLuid uint64, ip *net.IP) {
	fmt.Printf("UNICAST ADDRESS CHANGED! MibNotificationType: %s; interface LUID: %d; IP: %s\n",
		notificationType.String(), interfaceLuid, ip.String())
}

func routeChangeCallbackExample(notificationType MibNotificationType, route *Route) {
	fmt.Printf("ROUTE CHANGED! MibNotificationType: %s; destination: %s; next hop: %s\n",
		notificationType.String(), route.DestinationPrefix.String(), route.NextHop.String())
}

func TestGetInterfaces(t *testing.T) {

	ifcs, err := GetInterfacesEx(FullGetAdapterAddressesFlags())

	if err != nil {
		t.Errorf("GetInterfaces() returned error: %v", err)
	} else if ifcs == nil {
		t.Errorf("GetInterfaces() returned nil.")
	} else if interface_print {
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
	} else if interface_print {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromLUID() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromLUIDNonExisting(t *testing.T) {

	ifc, err := InterfaceFromLUID(unexistingLuid)

	if err == nil {
		t.Errorf("InterfaceFromLUID() returned an interface with LUID=%d, although requested LUID was %d.",
			ifc.Luid, unexistingLuid)
	} else if err.Error() != "InterfaceFromIndexEx() - interface with specified LUID not found" {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
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
	} else if interface_print {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromIndex() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromIndexNonExisting(t *testing.T) {

	ifc, err := InterfaceFromIndex(unexistingIndex)

	if err == nil {
		t.Errorf("InterfaceFromIndex() returned an interface with index=%d, although requested index was %d.",
			ifc.Index, unexistingIndex)
	} else if err.Error() != "InterfaceFromIndexEx() - interface with specified index not found" {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
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
	} else if interface_print {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromFriendlyName() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromFriendlyNameNonExisting(t *testing.T) {

	ifc, err := InterfaceFromFriendlyName(unexistingInterfaceName)

	if err == nil {
		t.Errorf("InterfaceFromFriendlyName() returned an interface with name=%s, although requested name was %s.",
			ifc.FriendlyName, unexistingInterfaceName)
	} else if err.Error() != "InterfaceFromFriendlyNameEx() - interface with specified friendly name not found" {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	}
}

func TestInterfaceFromGUID(t *testing.T) {

	luid := existingLuid

	guid, err := InterfaceLuidToGuid(luid)

	if err != nil {
		t.Errorf("InterfaceLuidToGuid() returned an error: %v. Have you forgot to set existingLuid appropriately?",
			err)
		return
	}

	if guid == nil {
		t.Error("InterfaceLuidToGuid() returned nil.")
		return
	}

	ifc, err := InterfaceFromGUID(guid)

	if err != nil {
		t.Errorf("InterfaceFromGUID() returned an error: %v", err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromGUID() returned nil. Have you forgot to set existingLuid appropriately?")
		return
	}

	if ifc.Luid != luid {
		t.Errorf("LUID mismatch. Expected: %d; actual: %d.", luid, ifc.Luid)
	}
}

func TestInterface_GetData(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.GetIpInterface() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.GetIpInterface() testing cannot be performed.")
		return
	}

	ifcdata, err := ifc.GetIpInterface(AF_INET)

	if err != nil {
		t.Errorf("Interface.GetIpInterface() returned an error: %v", err)
		return
	}

	if interface_printIpInterfaces {
		fmt.Println("====================== INTERFACE DATA OUTPUT START ======================")
		fmt.Println(ifcdata)
		fmt.Println("======================= INTERFACE DATA OUTPUT END =======================")
	}
}

func equalNetIPs(a, b []*net.IPNet) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if netCompare(*a[i], *b[i]) != 0 {
			return false
		}
	}
	return true
}

func ipnet4(ip string, bits int) *net.IPNet {
	return &net.IPNet{
		IP:   net.ParseIP(ip),
		Mask: net.CIDRMask(bits, 32),
	}
}

// each cidr can end in "[4]" to mean To4 form.
func nets(cidrs ...string) (ret []*net.IPNet) {
	for _, s := range cidrs {
		to4 := strings.HasSuffix(s, "[4]")
		if to4 {
			s = strings.TrimSuffix(s, "[4]")
		}
		ip, ipNet, err := net.ParseCIDR(s)
		if err != nil {
			panic(fmt.Sprintf("Bogus CIDR %q in test", s))
		}
		if to4 {
			ip = ip.To4()
		}
		ipNet.IP = ip
		ret = append(ret, ipNet)
	}
	return
}

func TestInterface_DeltaNets(t *testing.T) {
	tests := []struct {
		a, b             []*net.IPNet
		wantAdd, wantDel []*net.IPNet
	}{
		{
			a:       nets("1.2.3.4/24", "1.2.3.4/31", "1.2.3.3/32", "10.0.1.1/32", "100.0.1.1/32"),
			b:       nets("10.0.1.1/32", "100.0.2.1/32", "1.2.3.3/32", "1.2.3.4/24"),
			wantAdd: nets("100.0.2.1/32"),
			wantDel: nets("1.2.3.4/31", "100.0.1.1/32"),
		},
		{
			a:       nets("fe80::99d0:ec2d:b2e7:536b/64", "100.84.36.11/32"),
			b:       nets("100.84.36.11/32"),
			wantDel: nets("fe80::99d0:ec2d:b2e7:536b/64"),
		},
		{
			a:       nets("100.84.36.11/32", "fe80::99d0:ec2d:b2e7:536b/64"),
			b:       nets("100.84.36.11/32"),
			wantDel: nets("fe80::99d0:ec2d:b2e7:536b/64"),
		},
		{
			a:       nets("100.84.36.11/32", "fe80::99d0:ec2d:b2e7:536b/64"),
			b:       nets("100.84.36.11/32[4]"),
			wantDel: nets("fe80::99d0:ec2d:b2e7:536b/64"),
		},
		{
			a: excludeIPv6LinkLocal(nets("100.84.36.11/32", "fe80::99d0:ec2d:b2e7:536b/64")),
			b: nets("100.84.36.11/32"),
		},
		{
			a: []*net.IPNet{
				{
					IP:   net.ParseIP("1.2.3.4"),
					Mask: net.IPMask{0xff, 0xff, 0xff, 0xff},
				},
			},
			b: []*net.IPNet{
				{
					IP:   net.ParseIP("1.2.3.4"),
					Mask: net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
				},
			},
		},
	}
	for i, tt := range tests {
		add, del := deltaNets(tt.a, tt.b)
		if !equalNetIPs(add, tt.wantAdd) {
			t.Errorf("[%d] add:\n  got: %v\n want: %v\n", i, add, tt.wantAdd)
		}
		if !equalNetIPs(del, tt.wantDel) {
			t.Errorf("[%d] del:\n  got: %v\n want: %v\n", i, del, tt.wantDel)
		}
	}
}

func equalRouteDatas(a, b []*RouteData) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if routeDataCompare(a[i], b[i]) != 0 {
			return false
		}
	}
	return true
}

func TestInterface_DeltaRouteData(t *testing.T) {
	var h0 net.IP
	h1 := net.ParseIP("99.99.99.99")
	h2 := net.ParseIP("99.99.9.99")

	a := []*RouteData{
		&RouteData{*ipnet4("1.2.3.4", 32), h0, 1},
		&RouteData{*ipnet4("1.2.3.4", 24), h1, 2},
		&RouteData{*ipnet4("1.2.3.4", 24), h2, 1},
		&RouteData{*ipnet4("1.2.3.5", 32), h0, 1},
	}
	b := []*RouteData{
		&RouteData{*ipnet4("1.2.3.5", 32), h0, 1},
		&RouteData{*ipnet4("1.2.3.4", 24), h1, 2},
		&RouteData{*ipnet4("1.2.3.4", 24), h2, 2},
	}
	add, del := deltaRouteData(a, b)

	expect_add := []*RouteData{
		&RouteData{*ipnet4("1.2.3.4", 24), h2, 2},
	}
	expect_del := []*RouteData{
		&RouteData{*ipnet4("1.2.3.4", 32), h0, 1},
		&RouteData{*ipnet4("1.2.3.4", 24), h2, 1},
	}

	if !equalRouteDatas(expect_add, add) {
		t.Errorf("add:\n  want: %v\n   got: %v\n", expect_add, add)
	}
	if !equalRouteDatas(expect_del, del) {
		t.Errorf("del:\n  want: %v\n   got: %v\n", expect_del, del)
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

	addr, err := ifc.GetUnicastIpAddressRow(&unexistentIpAddresToAdd.IP)

	if err == nil {
		t.Errorf("Unicast address %s already exists. Please set unexistentIpAddresToAdd appropriately.",
			unexistentIpAddresToAdd.IP.String())
		return
	} else if err.Error() != "iphlpapi.GetUnicastIpAddressEntry: Element not found." {
		t.Errorf("Interface.GetUnicastIpAddressRow() returned an error: %v", err)
		return
	}

	cb, err := RegisterUnicastAddressChangeCallback(unicastAddressChangeCallbackExample)

	if err == nil {
		defer cb.Unregister()
	} else {
		t.Errorf("RegisterUnicastAddressChangeCallback() returned an error: %v", err)
	}

	count := len(ifc.UnicastAddresses)

	err = ifc.AddAddresses([]*net.IPNet{&unexistentIpAddresToAdd})

	if err != nil {
		t.Errorf("Interface.AddAddresses() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	ifc, _ = InterfaceFromLUID(ifc.Luid)

	if count+1 != len(ifc.UnicastAddresses) {
		t.Errorf("Number of unicast addresses before adding is %d, while number after adding is %d.", count,
			len(ifc.UnicastAddresses))
	}

	addr, err = ifc.GetUnicastIpAddressRow(&unexistentIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.GetUnicastIpAddressRow() returned an error: %v", err)
	} else if addr == nil {
		t.Errorf("Unicast address %s still doesn't exist, although it's added successfully.",
			unexistentIpAddresToAdd.IP.String())
	}

	err = ifc.DeleteAddress(&unexistentIpAddresToAdd.IP)

	if err != nil {
		t.Errorf("Interface.DeleteAddress() returned an error: %v", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	addr, err = ifc.GetUnicastIpAddressRow(&unexistentIpAddresToAdd.IP)

	if err == nil {
		t.Errorf("Unicast address %s still exists, although it's deleted successfully.",
			unexistentIpAddresToAdd.IP.String())
	} else if err.Error() != "iphlpapi.GetUnicastIpAddressEntry: Element not found." {
		t.Errorf("Interface.GetUnicastIpAddressRow() returned an error: %v", err)
		return
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

	if interface_printRoutes {
		for _, route := range routes {
			fmt.Println("========================== ROUTE OUTPUT START ==========================")
			fmt.Println(route)
			fmt.Println("=========================== ROUTE OUTPUT END ===========================")
		}
	}
}

func TestInterface_AddRoute_DeleteRoute(t *testing.T) {

	findRoute := func(ifc *Interface, dest *net.IPNet) ([]*Route, error) {
		routes, err := ifc.GetRoutes(AF_INET)
		if err != nil {
			return nil, err
		}
		matches := make([]*Route, len(routes))
		i := 0
		ones, _ := dest.Mask.Size()
		for _, route := range routes {
			if route.DestinationPrefix.PrefixLength == uint8(ones) && route.DestinationPrefix.Prefix.Address.Equal(dest.IP) {
				matches[i] = route
				i++
			}
		}
		return matches[:i], nil
	}

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

	_, err = ifc.GetRoute(&unexistentRouteIPv4ToAdd.Destination, &unexistentRouteIPv4ToAdd.NextHop)

	if err == nil {
		t.Error("Interface.GetRoute() returned a route although it isn't added yet. Have you forgot to set unexistentRouteIPv4ToAdd appropriately?")
		return
	} else if err.Error() != "iphlpapi.GetIpForwardEntry2: Element not found." {
		t.Errorf("Interface.GetRoute() returned an error: %v", err)
		return
	}

	routes, err := findRoute(ifc, &unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoutes() returned an error: %v", err)
		return
	}

	if len(routes) != 0 {
		t.Errorf("Interface.FindRoutes() returned %d items although the route isn't added yet. Have you forgot to set unexistentRouteIPv4ToAdd appropriately?",
			len(routes))
		return
	}

	cb, err := RegisterRouteChangeCallback(routeChangeCallbackExample)

	if err == nil {
		defer cb.Unregister()
	} else {
		t.Errorf("RegisterRouteChangeCallback() returned an error: %v", err)
	}

	err = ifc.AddRoute(&unexistentRouteIPv4ToAdd)

	if err != nil {
		t.Errorf("Interface.AddRoute() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	route, err := ifc.GetRoute(&unexistentRouteIPv4ToAdd.Destination, &unexistentRouteIPv4ToAdd.NextHop)

	if err != nil {
		if err.Error() == "iphlpapi.GetIpForwardEntry2: Element not found." {
			t.Error("Interface.GetRoute() returned nil although the route is added successfully.")
		} else {
			t.Errorf("Interface.GetRoute() returned an error: %v", err)
		}
	} else if !route.DestinationPrefix.Prefix.Address.Equal(unexistentRouteIPv4ToAdd.Destination.IP) ||
		!route.NextHop.Address.Equal(route.NextHop.Address) {
		t.Error("Interface.GetRoute() returned a wrong route!")
	}

	routes, err = findRoute(ifc, &unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoutes() returned an error: %v", err)
		return
	}

	if len(routes) != 1 {
		t.Errorf("Interface.FindRoutes() returned %d items although %d is expected.", len(routes), 1)
	} else if !routes[0].DestinationPrefix.Prefix.Address.Equal(unexistentRouteIPv4ToAdd.Destination.IP) {
		t.Errorf("Interface.FindRoutes() returned a wrong route. Dest: %s; expected: %s.",
			routes[0].DestinationPrefix.Prefix.Address.String(), unexistentRouteIPv4ToAdd.Destination.IP.String())
	}

	err = ifc.DeleteRoute(&unexistentRouteIPv4ToAdd.Destination, &unexistentRouteIPv4ToAdd.NextHop)

	if err != nil {
		t.Errorf("Iterface.DeleteRoute() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	_, err = ifc.GetRoute(&unexistentRouteIPv4ToAdd.Destination, &unexistentRouteIPv4ToAdd.NextHop)

	if err == nil {
		t.Error("Interface.GetRoute() returned a route although it is removed successfully.")
	} else if err.Error() != "iphlpapi.GetIpForwardEntry2: Element not found." {
		t.Errorf("Interface.GetRoute() returned an error: %v", err)
	}

	routes, err = findRoute(ifc, &unexistentRouteIPv4ToAdd.Destination)

	if err != nil {
		t.Errorf("Interface.FindRoutes() returned an error: %v", err)
		return
	}

	if len(routes) != 0 {
		t.Errorf("Interface.FindRoutes() returned %d items although the route is deleted successfully.",
			len(routes))
	}
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

	if ifc.DnsServerAddresses != nil {
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

	ifc, _ = InterfaceFromLUID(ifc.Luid)

	if interface_print {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Println(ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}

	if ifc.DnsServerAddresses != nil && len(ifc.DnsServerAddresses) != 0 {
		t.Errorf("DnsServerAddresses contains %d items, although FlushDNS is executed successfully.",
			len(ifc.DnsServerAddresses))
	}

	err = ifc.SetDNS(prevDnses)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v.", err)
	}
}

func TestInterface_AddDNS(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.AddDNS() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.AddDNS() testing cannot be performed.")
		return
	}

	prevDnsesCount := 0

	if ifc.DnsServerAddresses != nil {
		prevDnsesCount = len(ifc.DnsServerAddresses)
	}

	prevDnses := make([]net.IP, prevDnsesCount, prevDnsesCount)

	for i := 0; i < prevDnsesCount; i++ {
		prevDnses[i] = ifc.DnsServerAddresses[i].Address.Address
	}

	expectedDnses := append(prevDnses, dnsesToSet...)

	err = ifc.AddDNS(dnsesToSet)

	if err != nil {
		t.Errorf("Interface.AddDNS() returned an error: %v", err)
		return
	}

	ifc, _ = InterfaceFromLUID(ifc.Luid)

	if interface_print {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Println(ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}

	if expectedDnses == nil {
		if ifc.DnsServerAddresses != nil && len(ifc.DnsServerAddresses) != 0 {
			t.Errorf("expectedDnses is nil, but DnsServerAddresses contains %d items.",
				len(ifc.DnsServerAddresses))
		}
	} else {

		length := len(expectedDnses)

		if ifc.DnsServerAddresses == nil {
			t.Errorf("expectedDnses contains %d items, while DnsServerAddresses is nil.", length)
		} else if len(ifc.DnsServerAddresses) != length {
			t.Errorf("expectedDnses contains %d items, while DnsServerAddresses contains %d.", length,
				len(ifc.DnsServerAddresses))
		} else {
			for idx, dns := range expectedDnses {
				if !dns.Equal(ifc.DnsServerAddresses[idx].Address.Address) {
					t.Errorf("expectedDnses[%d] = %s while DnsServerAddresses[%d].Address.Address = %s.", idx,
						dns.String(), idx, ifc.DnsServerAddresses[idx].Address.Address.String())
				}
			}
		}
	}

	err = ifc.SetDNS(prevDnses)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v.", err)
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

	prevDnsesCount := 0

	if ifc.DnsServerAddresses != nil {
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

	ifc, _ = InterfaceFromLUID(ifc.Luid)

	if interface_print {
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

	err = ifc.SetDNS(prevDnses)

	if err != nil {
		t.Errorf("Interface.SetDNS() returned an error: %v.", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)
}
