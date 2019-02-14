/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "golang.org/x/sys/windows"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
// IP_ADAPTER_ADDRESSES_LH defined in iptypes.h
type wtIpAdapterAddressesLh struct {
	Length                uint32 // Windows type: ULONG
	IfIndex               uint32 // Windows type: IF_INDEX
	Next                  *wtIpAdapterAddressesLh
	AdapterName           *uint8 // Windows type: *CHAR
	FirstUnicastAddress   *wtIpAdapterUnicastAddressLh
	FirstAnycastAddress   *wtIpAdapterAnycastAddressXp
	FirstMulticastAddress *wtIpAdapterMulticastAddressXp
	FirstDnsServerAddress *wtIpAdapterDnsServerAddressXp
	DnsSuffix             *uint16                           // Windows type: *WCHAR
	Description           *uint16                           // Windows type: *WCHAR
	FriendlyName          *uint16                           // Windows type: *WCHAR
	PhysicalAddress       [MAX_ADAPTER_ADDRESS_LENGTH]uint8 // Windows type: [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	PhysicalAddressLength uint32                            // Windows type: ULONG
	Flags                 uint32                            // Windows type: ULONG
	Mtu                   uint32                            // Windows type: ULONG
	IfType                IfType
	OperStatus            IfOperStatus
	Ipv6IfIndex           uint32     // Windows type: IF_INDEX
	ZoneIndices           [16]uint32 // Windows type: [16]ULONG
	FirstPrefix           *wtIpAdapterPrefixXp

	TransmitLinkSpeed      uint64 // Windows type: ULONG64
	ReceiveLinkSpeed       uint64 // Windows type: ULONG64
	FirstWinsServerAddress *wtIpAdapterWinsServerAddressLh
	FirstGatewayAddress    *wtIpAdapterGatewayAddressLh
	Ipv4Metric             uint32 // Windows type: ULONG
	Ipv6Metric             uint32 // Windows type: ULONG
	Luid                   uint64 // Windows type:  IF_LUID
	Dhcpv4Server           wtSocketAddress
	CompartmentId          uint32       // Windows type: NET_IF_COMPARTMENT_ID
	NetworkGuid            windows.GUID // Windows type: NET_IF_NETWORK_GUID
	ConnectionType         NetIfConnectionType
	TunnelType             TunnelType
	//
	// DHCP v6 Info.
	//
	Dhcpv6Server           wtSocketAddress
	Dhcpv6ClientDuid       [MAX_DHCPV6_DUID_LENGTH]uint8 // Windows type: [MAX_DHCPV6_DUID_LENGTH]BYTE
	Dhcpv6ClientDuidLength uint32                        // Windows type: ULONG
	Dhcpv6Iaid             uint32                        // Windows type: ULONG
	FirstDnsSuffix         *wtIpAdapterDnsSuffix
}
