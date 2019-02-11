/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
// Defined in iptypes.h
type IP_ADAPTER_ADDRESSES_LH struct {
	Length ULONG
	IfIndex IF_INDEX
	Next *IP_ADAPTER_ADDRESSES_LH
	AdapterName *CHAR
	FirstUnicastAddress *IP_ADAPTER_UNICAST_ADDRESS_LH
	FirstAnycastAddress *IP_ADAPTER_ANYCAST_ADDRESS_XP
	FirstMulticastAddress *IP_ADAPTER_MULTICAST_ADDRESS_XP
	FirstDnsServerAddress *IP_ADAPTER_DNS_SERVER_ADDRESS_XP
	DnsSuffix *WCHAR
	Description *WCHAR
	FriendlyName *WCHAR
	PhysicalAddress [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	PhysicalAddressLength ULONG
	Flags ULONG
	Mtu ULONG
	IfType IFTYPE
	OperStatus IF_OPER_STATUS
	Ipv6IfIndex IF_INDEX
	ZoneIndices [16]ULONG
	FirstPrefix *IP_ADAPTER_PREFIX_XP

	TransmitLinkSpeed ULONG64
	ReceiveLinkSpeed ULONG64
	FirstWinsServerAddress *IP_ADAPTER_WINS_SERVER_ADDRESS_LH
	FirstGatewayAddress *IP_ADAPTER_GATEWAY_ADDRESS_LH
	Ipv4Metric ULONG
	Ipv6Metric ULONG
	Luid IF_LUID
	Dhcpv4Server SOCKET_ADDRESS
	CompartmentId NET_IF_COMPARTMENT_ID
	NetworkGuid NET_IF_NETWORK_GUID
	ConnectionType NET_IF_CONNECTION_TYPE
	TunnelType TUNNEL_TYPE
	//
	// DHCP v6 Info.
	//
	Dhcpv6Server SOCKET_ADDRESS
	Dhcpv6ClientDuid [MAX_DHCPV6_DUID_LENGTH]BYTE
	Dhcpv6ClientDuidLength ULONG
	Dhcpv6Iaid ULONG
	FirstDnsSuffix *IP_ADAPTER_DNS_SUFFIX
}
