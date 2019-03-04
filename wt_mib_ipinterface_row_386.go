/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipinterface_row
// MIB_IPINTERFACE_ROW defined in netioapi.h
type wtMibIpinterfaceRow struct {
	//
	// Key Structure;
	//
	Family AddressFamily

	offset1 [4]uint8 // Layout correction field

	InterfaceLuid  uint64 // Windows type: NET_LUID
	InterfaceIndex uint32 // Windows type: NET_IFINDEX

	//
	// Read-Write fields.
	//

	//
	// Fields currently not exposed.
	//
	MaxReassemblySize              uint32 // Windows type: ULONG
	InterfaceIdentifier            uint64 // Windows type: ULONG64
	MinRouterAdvertisementInterval uint32 // Windows type: ULONG
	MaxRouterAdvertisementInterval uint32 // Windows type: ULONG

	//
	// Fileds currently exposed.
	//
	AdvertisingEnabled                   uint8 // Windows type: BOOLEAN
	ForwardingEnabled                    uint8 // Windows type: BOOLEAN
	WeakHostSend                         uint8 // Windows type: BOOLEAN
	WeakHostReceive                      uint8 // Windows type: BOOLEAN
	UseAutomaticMetric                   uint8 // Windows type: BOOLEAN
	UseNeighborUnreachabilityDetection   uint8 // Windows type: BOOLEAN
	ManagedAddressConfigurationSupported uint8 // Windows type: BOOLEAN
	OtherStatefulConfigurationSupported  uint8 // Windows type: BOOLEAN
	AdvertiseDefaultRoute                uint8 // Windows type: BOOLEAN

	RouterDiscoveryBehavior NlRouterDiscoveryBehavior
	// DupAddrDetectTransmits in RFC 2462.
	DadTransmits      uint32 // Windows type: ULONG
	BaseReachableTime uint32 // Windows type: ULONG
	RetransmitTime    uint32 // Windows type: ULONG
	// Path MTU discovery timeout (in ms).
	PathMtuDiscoveryTimeout uint32 // Windows type: ULONG

	LinkLocalAddressBehavior NlLinkLocalAddressBehavior
	// In ms.
	LinkLocalAddressTimeout uint32 // Windows type: ULONG
	// Zone part of a SCOPE_ID.
	ZoneIndices      [ScopeLevelCount]uint32 // Windows type: [ScopeLevelCount]ULONG
	SitePrefixLength uint32                  // Windows type: ULONG
	Metric           uint32                  // Windows type: ULONG
	NlMtu            uint32                  // Windows type: ULONG

	//
	// Read Only fields.
	//
	Connected                 uint8 // Windows type: BOOLEAN
	SupportsWakeUpPatterns    uint8 // Windows type: BOOLEAN
	SupportsNeighborDiscovery uint8 // Windows type: BOOLEAN
	SupportsRouterDiscovery   uint8 // Windows type: BOOLEAN

	ReachableTime uint32 // Windows type: ULONG

	TransmitOffload wtNlInterfaceOffloadRodByte
	ReceiveOffload  wtNlInterfaceOffloadRodByte

	//
	// Disables using default route on the interface. This flag
	// can be used by VPN clients to restrict Split tunnelling.
	//
	DisableDefaultRoutes uint8 // Windows type: BOOLEAN
}
