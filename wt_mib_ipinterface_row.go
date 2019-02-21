/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"os"
	"golang.org/x/sys/windows"
)

func getWtMibIpinterfaceRow(interfaceLuid uint64) (*wtMibIpinterfaceRow, error) {

	wtrow := wtMibIpinterfaceRow{InterfaceLuid: interfaceLuid}

	result := getIpInterfaceEntry(&wtrow)

	if result == 0 {
		return &wtrow, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetIpInterfaceEntry", windows.Errno(result))
	}
}

func (wt *wtMibIpinterfaceRow) toInterfaceData() *InterfaceData {

	if wt == nil {
		return nil
	}

	return &InterfaceData{
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
