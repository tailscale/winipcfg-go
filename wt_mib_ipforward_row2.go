/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_row2
// MIB_IPFORWARD_ROW2 defined in netioapi.h
type wtMibIpforwardRow2 struct {
	//
	// Key Structure.
	//
	InterfaceLuid     uint64 // Windows type: NET_LUID
	InterfaceIndex    uint32 // Windows type: NET_IFINDEX
	DestinationPrefix wtIpAddressPrefix
	NextHop           wtSockaddrInet

	//
	// Read-Write Fields.
	//
	SitePrefixLength  uint8  // Windows type: UCHAR
	ValidLifetime     uint32 // Windows type: ULONG
	PreferredLifetime uint32 // Windows type: ULONG
	Metric            uint32 // Windows type: ULONG
	Protocol          NlRouteProtocol

	Loopback             uint8 // Windows type: BOOLEAN
	AutoconfigureAddress uint8 // Windows type: BOOLEAN
	Publish              uint8 // Windows type: BOOLEAN
	Immortal             uint8 // Windows type: BOOLEAN

	//
	// Read-Only Fields.
	//
	Age    uint32 // Windows type: ULONG
	Origin NlRouteOrigin
}

func getWtMibIpforwardRow2s(family AddressFamily, ifc *Interface) ([]wtMibIpforwardRow2, error) {

	var pTable *wtMibIpforwardTable2 = nil

	result := getIpForwardTable2(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIpForwardTable2", windows.Errno(result))
	}

	rows := make([]wtMibIpforwardRow2, 0)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIpforwardRow2_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {

		row := (*wtMibIpforwardRow2)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))

		if ifc == nil || row.InterfaceLuid == ifc.Luid {
			rows = append(rows, *row)
		}
	}

	return rows, nil
}

func findWtMibIpforwardRow2(destination *net.IPNet, ifc *Interface) (*wtMibIpforwardRow2, error) {

	if destination == nil {
		return nil, fmt.Errorf("findWtMibIpforwardRow2() - input argument 'destination' is nil")
	}

	rows, err := getWtMibIpforwardRow2s(AF_UNSPEC, ifc)

	if err != nil {
		return nil, err
	}

	ones, _ := destination.Mask.Size()

	for _, row := range rows {
		if row.DestinationPrefix.PrefixLength == uint8(ones) && row.DestinationPrefix.Prefix.matches(&destination.IP) {
			return &row, nil
		}
	}

	return nil, nil
}

func addWtMibIpforwardRow2(ifc *Interface, routeData *RouteData) error {

	if ifc == nil || routeData == nil {
		return fmt.Errorf("addWtMibIpforwardRow2() - some of the input arguments is nil")
	}

	wtdest, err := createWtIpAddressPrefix(&routeData.Destination)

	if err != nil {
		return err
	}

	wtsaNextHop, err := createWtSockaddrInet(&routeData.NextHop, 0)

	if err != nil {
		return err
	}

	row := wtMibIpforwardRow2{}

	_ = initializeIpForwardEntry(&row)

	//fmt.Printf("wtMibIpforwardRow2 initialized to:\n%s\n", row.String())

	row.InterfaceLuid = ifc.Luid
	row.InterfaceIndex = ifc.Index
	row.DestinationPrefix = *wtdest
	row.NextHop = *wtsaNextHop
	row.Metric = routeData.Metric

	//fmt.Printf("wtMibIpforwardRow2 to add:\n%s\n", row.String())

	result := createIpForwardEntry2(&row)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.createIpForwardEntry2", windows.Errno(result))
	}
}

func (r *wtMibIpforwardRow2) delete() error {

	result := deleteIpForwardEntry2(r)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.DeleteIpForwardEntry2", windows.Errno(result))
	}
}

func (r *wtMibIpforwardRow2) toRoute() (*Route, error) {

	if r == nil {
		return nil, nil
	}

	iap, err := r.DestinationPrefix.toIpAddressPrefix()

	if err != nil {
		return nil, err
	}

	sainet, err := r.NextHop.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &Route{
		InterfaceLuid:        r.InterfaceLuid,
		InterfaceIndex:       r.InterfaceIndex,
		DestinationPrefix:    *iap,
		NextHop:              *sainet,
		SitePrefixLength:     r.SitePrefixLength,
		ValidLifetime:        r.ValidLifetime,
		PreferredLifetime:    r.PreferredLifetime,
		Metric:               r.Metric,
		Protocol:             r.Protocol,
		Loopback:             uint8ToBool(r.Loopback),
		AutoconfigureAddress: uint8ToBool(r.AutoconfigureAddress),
		Publish:              uint8ToBool(r.Publish),
		Immortal:             uint8ToBool(r.Immortal),
		Age:                  r.Age,
		Origin:               r.Origin,
	}, nil
}

func (r *wtMibIpforwardRow2) extractRouteData() (*RouteData, error) {

	if r == nil {
		return nil, nil
	}

	iap, err := r.DestinationPrefix.toIpAddressPrefix()

	if err != nil {
		return nil, err
	}

	destination, err := iap.toNetIpNet()

	if err != nil {
		return nil, err
	}

	sainet, err := r.NextHop.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &RouteData{
		Destination: *destination,
		NextHop:     sainet.Address,
		Metric:      r.Metric,
	}, nil
}

func (r *wtMibIpforwardRow2) String() string {

	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`InterfaceLuid: %d
InterfaceIndex: %d
DestinationPrefix: %s
NextHop: %s
SitePrefixLength: %d
ValidLifetime: %d
PreferredLifetime: %d
Metric: %d
Protocol: %s
Loopback: %d
AutoconfigureAddress: %d
Publish: %d
Immortal: %d
Age: %d
Origin: %s
`, r.InterfaceLuid, r.InterfaceIndex, r.DestinationPrefix.String(), r.NextHop.String(), r.SitePrefixLength,
		r.ValidLifetime, r.PreferredLifetime, r.Metric, r.Protocol.String(), r.Loopback, r.AutoconfigureAddress,
		r.Publish, r.Immortal, r.Age, r.Origin.String())
}
