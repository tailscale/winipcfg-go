/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

const (
	IP_ADAPTER_ADDRESSES_LH_Size = 376

	IP_ADAPTER_ADDRESSES_LH_IfIndex_Offset = 4
	IP_ADAPTER_ADDRESSES_LH_Next_Offset = 8
	IP_ADAPTER_ADDRESSES_LH_AdapterName_Offset = 12
	IP_ADAPTER_ADDRESSES_LH_FirstUnicastAddress_Offset = 16
	IP_ADAPTER_ADDRESSES_LH_FirstAnycastAddress_Offset = 20
	IP_ADAPTER_ADDRESSES_LH_FirstMulticastAddress_Offset = 24
	IP_ADAPTER_ADDRESSES_LH_FirstDnsServerAddress_Offset = 28
	IP_ADAPTER_ADDRESSES_LH_DnsSuffix_Offset = 32
	IP_ADAPTER_ADDRESSES_LH_Description_Offset = 36
	IP_ADAPTER_ADDRESSES_LH_FriendlyName_Offset = 40
	IP_ADAPTER_ADDRESSES_LH_PhysicalAddress_Offset = 44
	IP_ADAPTER_ADDRESSES_LH_PhysicalAddressLength_Offset = 52
	IP_ADAPTER_ADDRESSES_LH_Flags_Offset = 56
	IP_ADAPTER_ADDRESSES_LH_Mtu_Offset = 60
	IP_ADAPTER_ADDRESSES_LH_IfType_Offset = 64
	IP_ADAPTER_ADDRESSES_LH_OperStatus_Offset = 68
	IP_ADAPTER_ADDRESSES_LH_Ipv6IfIndex_Offset = 72
	IP_ADAPTER_ADDRESSES_LH_ZoneIndices_Offset = 76
	IP_ADAPTER_ADDRESSES_LH_FirstPrefix_Offset = 140
	IP_ADAPTER_ADDRESSES_LH_TransmitLinkSpeed_Offset = 144
	IP_ADAPTER_ADDRESSES_LH_ReceiveLinkSpeed_Offset = 152
	IP_ADAPTER_ADDRESSES_LH_FirstWinsServerAddress_Offset = 160
	IP_ADAPTER_ADDRESSES_LH_FirstGatewayAddress_Offset = 164
	IP_ADAPTER_ADDRESSES_LH_Ipv4Metric_Offset = 168
	IP_ADAPTER_ADDRESSES_LH_Ipv6Metric_Offset = 172
	IP_ADAPTER_ADDRESSES_LH_Luid_Offset = 176
	IP_ADAPTER_ADDRESSES_LH_Dhcpv4Server_Offset = 184
	IP_ADAPTER_ADDRESSES_LH_CompartmentId_Offset = 192
	IP_ADAPTER_ADDRESSES_LH_NetworkGuid_Offset = 196
	IP_ADAPTER_ADDRESSES_LH_ConnectionType_Offset = 212
	IP_ADAPTER_ADDRESSES_LH_TunnelType_Offset = 216
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6Server_Offset = 220
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuid_Offset = 228
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuidLength_Offset = 360
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6Iaid_Offset = 364
	IP_ADAPTER_ADDRESSES_LH_FirstDnsSuffix_Offset = 368

	IP_ADAPTER_ANYCAST_ADDRESS_XP_Size = 24

	IP_ADAPTER_ANYCAST_ADDRESS_XP_Flags_Offset = 4
	IP_ADAPTER_ANYCAST_ADDRESS_XP_Next_Offset = 8
	IP_ADAPTER_ANYCAST_ADDRESS_XP_Address_Offset = 12

	SOCKET_ADDRESS_Size = 8

	SOCKET_ADDRESS_iSockaddrLength_Offset = 4
)
