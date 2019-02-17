/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

type Route struct {
	InterfaceLuid        uint64
	InterfaceIndex       uint32
	DestinationPrefix    IpAddressPrefix
	NextHop              SockaddrInet
	SitePrefixLength     uint8
	ValidLifetime        uint32
	PreferredLifetime    uint32
	Metric               uint32
	Protocol             NlRouteProtocol
	Loopback             bool
	AutoconfigureAddress bool
	Publish              bool
	Immortal             bool
	Age                  uint32
	Origin               NlRouteOrigin
}

func getRoutes(family AddressFamily, ifc *Interface) ([]*Route, error) {

	var pTable *wtMibIpforwardTable2 = nil

	result := getIpForwardTable2(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIpForwardTable2", windows.Errno(result))
	}

	routes := make([]*Route, 0)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIpforwardRow2_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {

		wtr := (*wtMibIpforwardRow2)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))

		if ifc == nil || wtr.InterfaceLuid == ifc.Luid {

			route, err := wtr.toRoute()

			if err != nil {
				return nil, err
			}

			routes = append(routes, route)
		}
	}

	return routes, nil
}

func GetRoutes(family AddressFamily) ([]*Route, error) {
	return getRoutes(family, nil)
}

func (r *Route) String() string {

	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`
InterfaceLuid: %d
InterfaceIndex: %d
DestinationPrefix: %s
NextHop: %s
SitePrefixLength: %d
ValidLifetime: %d
PreferredLifetime: %d
Metric: %d
Protocol: %s
Loopback: %v
AutoconfigureAddress: %v
Publish: %v
Immortal: %v
Age: %d
Origin: %s
`, r.InterfaceLuid, r.InterfaceIndex, r.DestinationPrefix.String(), r.NextHop.String(), r.SitePrefixLength,
		r.ValidLifetime, r.PreferredLifetime, r.Metric, r.Protocol.String(), r.Loopback, r.AutoconfigureAddress,
		r.Publish, r.Immortal, r.Age, r.Origin.String())
}
