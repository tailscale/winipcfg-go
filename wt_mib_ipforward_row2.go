/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
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

// Uses GetIpForwardTable2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2).
func getWtMibIpforwardRow2s(family AddressFamily) ([]*wtMibIpforwardRow2, error) {

	var pTable *wtMibIpforwardTable2 = nil

	result := getIpForwardTable2(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIpForwardTable2", windows.Errno(result))
	}

	rows := make([]*wtMibIpforwardRow2, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIpforwardRow2_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		// Dereferencing and rereferencing in order to force copying.

		row := *(*wtMibIpforwardRow2)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
		rows[i] = &row
	}

	return rows, nil
}

// Uses InitializeIpForwardEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-initializeipforwardentry).
func getInitializedWtMibIpforwardRow2(interfaceLuid uint64) *wtMibIpforwardRow2 {

	row := wtMibIpforwardRow2{InterfaceLuid: interfaceLuid}

	_ = initializeIpForwardEntry(&row)

	row.InterfaceLuid = interfaceLuid

	return &row
}

// Alternative version (with different input arguments) of getWtMibIpforwardRow2.
func getWtMibIpforwardRow2Alt(interfaceLuid uint64, destination *net.IPNet, nextHop *net.IP) (*wtMibIpforwardRow2, error) {

	wtDest, err := createWtIpAddressPrefix(destination)

	if err != nil {
		return nil, err
	}

	wtNextHop, err := createWtSockaddrInet(nextHop, 0)

	if err != nil {
		return nil, err
	}

	return getWtMibIpforwardRow2(interfaceLuid, wtDest, wtNextHop)
}

// Uses GetIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardentry2).
func getWtMibIpforwardRow2(interfaceLuid uint64, destination *wtIpAddressPrefix, nextHop *wtSockaddrInet) (*wtMibIpforwardRow2, error) {

	row := getInitializedWtMibIpforwardRow2(interfaceLuid)

	row.DestinationPrefix = *destination
	row.NextHop = *nextHop

	result := getIpForwardEntry2(row)

	if result == 0 {
		return row, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetIpForwardEntry2", windows.Errno(result))
	}
}

func createAndAddWtMibIpforwardRow2(interfaceLuid uint64, routeData *RouteData) error {

	wtdest, err := createWtIpAddressPrefix(&routeData.Destination)

	if err != nil {
		return err
	}

	wtsaNextHop, err := createWtSockaddrInet(&routeData.NextHop, 0)

	if err != nil {
		return err
	}

	row := getInitializedWtMibIpforwardRow2(interfaceLuid)

	row.DestinationPrefix = *wtdest
	row.NextHop = *wtsaNextHop
	row.Metric = routeData.Metric

	log.Printf("Add forward row: %v\n", routeData)
	return row.add()
}

// Uses CreateIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createipforwardentry2).
func (r *wtMibIpforwardRow2) add() error {

	result := createIpForwardEntry2(r)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.CreateIpForwardEntry2", windows.Errno(result))
	}
}

// Uses SetIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipforwardentry2).
func (r *wtMibIpforwardRow2) set() error {

	result := setIpForwardEntry2(r)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.SetIpForwardEntry2", windows.Errno(result))
	}
}

// Uses DeleteIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteipforwardentry2).
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
Origin: %s`, r.InterfaceLuid, r.InterfaceIndex, r.DestinationPrefix.String(), r.NextHop.String(), r.SitePrefixLength,
		r.ValidLifetime, r.PreferredLifetime, r.Metric, r.Protocol.String(), r.Loopback, r.AutoconfigureAddress,
		r.Publish, r.Immortal, r.Age, r.Origin.String())
}

type planetSorter struct {
	planets []*wtMibIpforwardRow2
	by      func(p1, p2 *wtMibIpforwardRow2) bool // Closure used in the Less method.
}
