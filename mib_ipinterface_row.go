/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type MibIpinterfaceRow struct {
	//
	// Key Structure;
	//
	Family         AddressFamily
	InterfaceLuid  uint64
	InterfaceIndex uint32

	//
	// Read-Write fields.
	//

	//
	// Fields currently not exposed.
	//
	MaxReassemblySize              uint32
	InterfaceIdentifier            uint64
	MinRouterAdvertisementInterval uint32
	MaxRouterAdvertisementInterval uint32

	//
	// Fileds currently exposed.
	//
	AdvertisingEnabled                   bool
	ForwardingEnabled                    bool
	WeakHostSend                         bool
	WeakHostReceive                      bool
	UseAutomaticMetric                   bool
	UseNeighborUnreachabilityDetection   bool
	ManagedAddressConfigurationSupported bool
	OtherStatefulConfigurationSupported  bool
	AdvertiseDefaultRoute                bool

	RouterDiscoveryBehavior NlRouterDiscoveryBehavior
	// DupAddrDetectTransmits in RFC 2462.
	DadTransmits      uint32
	BaseReachableTime uint32
	RetransmitTime    uint32
	// Path MTU discovery timeout (in ms).
	PathMtuDiscoveryTimeout uint32

	LinkLocalAddressBehavior NlLinkLocalAddressBehavior
	// In ms.
	LinkLocalAddressTimeout uint32
	// Zone part of a SCOPE_ID.
	ZoneIndices      [ScopeLevelCount]uint32
	SitePrefixLength uint32
	Metric           uint32
	NlMtu            uint32

	//
	// Read Only fields.
	//
	Connected                 bool
	SupportsWakeUpPatterns    bool
	SupportsNeighborDiscovery bool
	SupportsRouterDiscovery   bool

	ReachableTime uint32

	TransmitOffload NlInterfaceOffloadRodFlags
	ReceiveOffload  NlInterfaceOffloadRodFlags

	//
	// Disables using default route on the interface. This flag
	// can be used by VPN clients to restrict Split tunnelling.
	//
	DisableDefaultRoutes bool
}

func (mir *MibIpinterfaceRow) String() string {

	if mir == nil {
		return ""
	}

	return fmt.Sprintf(`Family: %s
InterfaceLuid: %d
InterfaceIndex: %d
MaxReassemblySize: %d
InterfaceIdentifier: %d
MinRouterAdvertisementInterval: %d
MaxRouterAdvertisementInterval: %d
AdvertisingEnabled: %v
ForwardingEnabled: %v
WeakHostSend: %v
WeakHostReceive: %v
UseAutomaticMetric: %v
UseNeighborUnreachabilityDetection: %v
ManagedAddressConfigurationSupported: %v
OtherStatefulConfigurationSupported: %v
AdvertiseDefaultRoute: %v
RouterDiscoveryBehavior: %s
DadTransmits: %d
BaseReachableTime: %d
RetransmitTime: %d
PathMtuDiscoveryTimeout: %d
LinkLocalAddressBehavior: %s
LinkLocalAddressTimeout: %d
ZoneIndices: %v
SitePrefixLength: %d
Metric: %d
NlMtu: %d
Connected: %v
SupportsWakeUpPatterns: %v
SupportsNeighborDiscovery: %v
SupportsRouterDiscovery: %v
ReachableTime: %d
TransmitOffload:
%s
ReceiveOffload:
%s
DisableDefaultRoutes: %v
	`, mir.Family.String(), mir.InterfaceLuid, mir.InterfaceIndex, mir.MaxReassemblySize, mir.InterfaceIdentifier,
	mir.MinRouterAdvertisementInterval, mir.MaxRouterAdvertisementInterval, mir.AdvertisingEnabled,
	mir.ForwardingEnabled, mir.WeakHostSend, mir.WeakHostReceive, mir.UseAutomaticMetric,
	mir.UseNeighborUnreachabilityDetection, mir.ManagedAddressConfigurationSupported,
	mir.OtherStatefulConfigurationSupported, mir.AdvertiseDefaultRoute, mir.RouterDiscoveryBehavior.String(),
	mir.DadTransmits, mir.BaseReachableTime, mir.RetransmitTime, mir.PathMtuDiscoveryTimeout,
	mir.LinkLocalAddressBehavior.String(), mir.LinkLocalAddressTimeout, mir.ZoneIndices, mir.SitePrefixLength,
	mir.Metric, mir.NlMtu, mir.Connected, mir.SupportsWakeUpPatterns, mir.SupportsNeighborDiscovery,
	mir.SupportsRouterDiscovery, mir.ReachableTime, toIndentedText(mir.TransmitOffload.String(), "  "),
	toIndentedText(mir.ReceiveOffload.String(), "  "), mir.DisableDefaultRoutes)
}
