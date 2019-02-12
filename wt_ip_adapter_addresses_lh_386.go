/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
// Defined in iptypes.h
type IP_ADAPTER_ADDRESSES_LH struct {
	Length ULONG
	IfIndex uint32 // Windows type: IF_INDEX
	Next *IP_ADAPTER_ADDRESSES_LH
	AdapterName *uint8 //*CHAR flattened to *uint8
	FirstUnicastAddress *wtIpAdapterUnicastAddressLh
	FirstAnycastAddress *IP_ADAPTER_ANYCAST_ADDRESS_XP
	FirstMulticastAddress *wtIpAdapterMulticastAddressXp
	FirstDnsServerAddress *wtIpAdapterDnsServerAddressXp
	DnsSuffix *uint16 // Windows type: *WCHAR
	Description *uint16 // Windows type: *WCHAR
	FriendlyName *uint16 // Windows type: *WCHAR
	PhysicalAddress [MAX_ADAPTER_ADDRESS_LENGTH]uint8 // Windows type: [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	PhysicalAddressLength uint32 // Windows type: ULONG
	Flags uint32 // Windows type: ULONG
	Mtu uint32 // Windows type: ULONG
	IfType IfType
	OperStatus IfOperStatus
	Ipv6IfIndex uint32 // Windows type: IF_INDEX
	ZoneIndices [16]uint32 // Windows type: [16]ULONG
	FirstPrefix *wtIpAdapterPrefixXp

	TransmitLinkSpeed ULONG64
	ReceiveLinkSpeed ULONG64
	FirstWinsServerAddress *wtIpAdapterWinsServerAddressLh
	FirstGatewayAddress *wtIpAdapterGatewayAddressLh
	Ipv4Metric ULONG
	Ipv6Metric ULONG
	Luid uint64 // Windows type:  IF_LUID
	Dhcpv4Server wtSocketAddress
	CompartmentId NET_IF_COMPARTMENT_ID
	NetworkGuid NET_IF_NETWORK_GUID
	ConnectionType NetIfConnectionType
	TunnelType TUNNEL_TYPE
	//
	// DHCP v6 Info.
	//
	Dhcpv6Server wtSocketAddress
	Dhcpv6ClientDuid [MAX_DHCPV6_DUID_LENGTH]BYTE
	Dhcpv6ClientDuidLength ULONG
	Dhcpv6Iaid ULONG
	FirstDnsSuffix *wtIpAdapterDnsSuffix
	// Fixing layout! I've had to add this padding to ensure the same structure size.
	correction [4]uint8
}
