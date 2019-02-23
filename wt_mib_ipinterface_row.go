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
func getWtMibIpinterfaceRows(family AddressFamily) ([]wtMibIpinterfaceRow, error) {

	var pTable *wtMibIpinterfaceTable = nil

	result := getIpInterfaceTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIpInterfaceTable", windows.Errno(result))
	}

	ipifcs := make([]wtMibIpinterfaceRow, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIpinterfaceRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		ipifcs[i] = *(*wtMibIpinterfaceRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
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

func (wt *wtMibIpinterfaceRow) toIpInterface() *IpInterface {

	if wt == nil {
		return nil
	}

	return &IpInterface{
		Family:                               wt.Family,
		InterfaceLuid:                        wt.InterfaceLuid,
		InterfaceIndex:                       wt.InterfaceIndex,
		MaxReassemblySize:                    wt.MaxReassemblySize,
		InterfaceIdentifier:                  wt.InterfaceIdentifier,
		MinRouterAdvertisementInterval:       wt.MinRouterAdvertisementInterval,
		MaxRouterAdvertisementInterval:       wt.MaxRouterAdvertisementInterval,
		AdvertisingEnabled:                   uint8ToBool(wt.AdvertisingEnabled),
		ForwardingEnabled:                    uint8ToBool(wt.ForwardingEnabled),
		WeakHostSend:                         uint8ToBool(wt.WeakHostSend),
		WeakHostReceive:                      uint8ToBool(wt.WeakHostReceive),
		UseAutomaticMetric:                   uint8ToBool(wt.UseAutomaticMetric),
		UseNeighborUnreachabilityDetection:   uint8ToBool(wt.UseNeighborUnreachabilityDetection),
		ManagedAddressConfigurationSupported: uint8ToBool(wt.ManagedAddressConfigurationSupported),
		OtherStatefulConfigurationSupported:  uint8ToBool(wt.OtherStatefulConfigurationSupported),
		AdvertiseDefaultRoute:                uint8ToBool(wt.AdvertiseDefaultRoute),
		RouterDiscoveryBehavior:              wt.RouterDiscoveryBehavior,
		DadTransmits:                         wt.DadTransmits,
		BaseReachableTime:                    wt.BaseReachableTime,
		RetransmitTime:                       wt.RetransmitTime,
		PathMtuDiscoveryTimeout:              wt.PathMtuDiscoveryTimeout,
		LinkLocalAddressBehavior:             wt.LinkLocalAddressBehavior,
		LinkLocalAddressTimeout:              wt.LinkLocalAddressTimeout,
		ZoneIndices:                          wt.ZoneIndices,
		SitePrefixLength:                     wt.SitePrefixLength,
		Metric:                               wt.Metric,
		NlMtu:                                wt.NlMtu,
		Connected:                            uint8ToBool(wt.Connected),
		SupportsWakeUpPatterns:               uint8ToBool(wt.SupportsWakeUpPatterns),
		SupportsNeighborDiscovery:            uint8ToBool(wt.SupportsNeighborDiscovery),
		SupportsRouterDiscovery:              uint8ToBool(wt.SupportsRouterDiscovery),
		ReachableTime:                        wt.ReachableTime,
		TransmitOffload:                      *wt.TransmitOffload.toNlInterfaceOffloadRodFlags(),
		ReceiveOffload:                       *wt.ReceiveOffload.toNlInterfaceOffloadRodFlags(),
		DisableDefaultRoutes:                 uint8ToBool(wt.DisableDefaultRoutes),
	}
}
