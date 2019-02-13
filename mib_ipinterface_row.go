/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

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
