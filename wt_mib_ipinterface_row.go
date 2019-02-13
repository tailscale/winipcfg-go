/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

func (wt *wtMibIpinterfaceRow) toMibIpinterfaceRow() *MibIpinterfaceRow {

	if wt == nil {
		return nil
	}

	return &MibIpinterfaceRow{
		Family:                               wt.Family,
		InterfaceLuid:                        wt.InterfaceLuid,
		InterfaceIndex:                       wt.InterfaceIndex,
		MaxReassemblySize:                    wt.MaxReassemblySize,
		InterfaceIdentifier:                  wt.InterfaceIdentifier,
		MinRouterAdvertisementInterval:       wt.MinRouterAdvertisementInterval,
		MaxRouterAdvertisementInterval:       wt.MaxRouterAdvertisementInterval,
		AdvertisingEnabled:                   wt.AdvertisingEnabled != 0,
		ForwardingEnabled:                    wt.ForwardingEnabled != 0,
		WeakHostSend:                         wt.WeakHostSend != 0,
		WeakHostReceive:                      wt.WeakHostReceive != 0,
		UseAutomaticMetric:                   wt.UseAutomaticMetric != 0,
		UseNeighborUnreachabilityDetection:   wt.UseNeighborUnreachabilityDetection != 0,
		ManagedAddressConfigurationSupported: wt.ManagedAddressConfigurationSupported != 0,
		OtherStatefulConfigurationSupported:  wt.OtherStatefulConfigurationSupported != 0,
		AdvertiseDefaultRoute:                wt.AdvertiseDefaultRoute != 0,
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
		Connected:                            wt.Connected != 0,
		SupportsWakeUpPatterns:               wt.SupportsWakeUpPatterns != 0,
		SupportsNeighborDiscovery:            wt.SupportsNeighborDiscovery != 0,
		SupportsRouterDiscovery:              wt.SupportsRouterDiscovery != 0,
		ReachableTime:                        wt.ReachableTime,
		TransmitOffload:                      *wt.TransmitOffload.toNlInterfaceOffloadRodFlags(),
		ReceiveOffload:                       *wt.ReceiveOffload.toNlInterfaceOffloadRodFlags(),
		DisableDefaultRoutes:                 wt.DisableDefaultRoutes != 0,
	}
}
