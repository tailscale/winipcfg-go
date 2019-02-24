/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"os"
	"golang.org/x/sys/windows"
	"unsafe"
)

// Corresponds to GetIpInterfaceTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfacetable)
func getWtMibIpinterfaceRows(family AddressFamily) ([]*wtMibIpinterfaceRow, error) {

	var pTable *wtMibIpinterfaceTable = nil

	result := getIpInterfaceTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIpInterfaceTable", windows.Errno(result))
	}

	ipifcs := make([]*wtMibIpinterfaceRow, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIpinterfaceRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		// Dereferencing and rereferencing in order to force copying.
		row := *(*wtMibIpinterfaceRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
		ipifcs[i] = &row
	}

	return ipifcs, nil
}

// Corresponds to GetIpInterfaceEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfaceentry)
func getWtMibIpinterfaceRow(interfaceLuid uint64, family AddressFamily) (*wtMibIpinterfaceRow, error) {

	if family != AF_INET && family != AF_INET6 {
		return nil, fmt.Errorf("argument 'family' has to be either AF_INET or AF_INET6")
	}

	wtrow := wtMibIpinterfaceRow{InterfaceLuid: interfaceLuid, Family: family}

	_ = initializeIpInterfaceEntry(&wtrow)

	wtrow.InterfaceLuid = interfaceLuid
	wtrow.Family = family

	result := getIpInterfaceEntry(&wtrow)

	if result == 0 {
		return &wtrow, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetIpInterfaceEntry", windows.Errno(result))
	}
}

// Corresponds to SetIpInterfaceEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipinterfaceentry)
func (wtipifc *wtMibIpinterfaceRow) set() error {

	result := setIpInterfaceEntry(wtipifc)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.SetIpInterfaceEntry", windows.Errno(result))
	}
}

func (wtipifc *wtMibIpinterfaceRow) toIpInterface() *IpInterface {

	if wtipifc == nil {
		return nil
	}

	return &IpInterface{
		Family:                               wtipifc.Family,
		InterfaceLuid:                        wtipifc.InterfaceLuid,
		InterfaceIndex:                       wtipifc.InterfaceIndex,
		MaxReassemblySize:                    wtipifc.MaxReassemblySize,
		InterfaceIdentifier:                  wtipifc.InterfaceIdentifier,
		MinRouterAdvertisementInterval:       wtipifc.MinRouterAdvertisementInterval,
		MaxRouterAdvertisementInterval:       wtipifc.MaxRouterAdvertisementInterval,
		AdvertisingEnabled:                   uint8ToBool(wtipifc.AdvertisingEnabled),
		ForwardingEnabled:                    uint8ToBool(wtipifc.ForwardingEnabled),
		WeakHostSend:                         uint8ToBool(wtipifc.WeakHostSend),
		WeakHostReceive:                      uint8ToBool(wtipifc.WeakHostReceive),
		UseAutomaticMetric:                   uint8ToBool(wtipifc.UseAutomaticMetric),
		UseNeighborUnreachabilityDetection:   uint8ToBool(wtipifc.UseNeighborUnreachabilityDetection),
		ManagedAddressConfigurationSupported: uint8ToBool(wtipifc.ManagedAddressConfigurationSupported),
		OtherStatefulConfigurationSupported:  uint8ToBool(wtipifc.OtherStatefulConfigurationSupported),
		AdvertiseDefaultRoute:                uint8ToBool(wtipifc.AdvertiseDefaultRoute),
		RouterDiscoveryBehavior:              wtipifc.RouterDiscoveryBehavior,
		DadTransmits:                         wtipifc.DadTransmits,
		BaseReachableTime:                    wtipifc.BaseReachableTime,
		RetransmitTime:                       wtipifc.RetransmitTime,
		PathMtuDiscoveryTimeout:              wtipifc.PathMtuDiscoveryTimeout,
		LinkLocalAddressBehavior:             wtipifc.LinkLocalAddressBehavior,
		LinkLocalAddressTimeout:              wtipifc.LinkLocalAddressTimeout,
		ZoneIndices:                          wtipifc.ZoneIndices,
		SitePrefixLength:                     wtipifc.SitePrefixLength,
		Metric:                               wtipifc.Metric,
		NlMtu:                                wtipifc.NlMtu,
		Connected:                            uint8ToBool(wtipifc.Connected),
		SupportsWakeUpPatterns:               uint8ToBool(wtipifc.SupportsWakeUpPatterns),
		SupportsNeighborDiscovery:            uint8ToBool(wtipifc.SupportsNeighborDiscovery),
		SupportsRouterDiscovery:              uint8ToBool(wtipifc.SupportsRouterDiscovery),
		ReachableTime:                        wtipifc.ReachableTime,
		TransmitOffload:                      *wtipifc.TransmitOffload.toNlInterfaceOffloadRodFlags(),
		ReceiveOffload:                       *wtipifc.ReceiveOffload.toNlInterfaceOffloadRodFlags(),
		DisableDefaultRoutes:                 uint8ToBool(wtipifc.DisableDefaultRoutes),
	}
}
