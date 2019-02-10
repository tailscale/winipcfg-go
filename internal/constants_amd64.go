/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

const (
	IP_ADAPTER_ADDRESSES_LH_Size = 448

	IP_ADAPTER_ADDRESSES_LH_IfIndex_Offset = 4
	IP_ADAPTER_ADDRESSES_LH_Next_Offset = 8
	IP_ADAPTER_ADDRESSES_LH_AdapterName_Offset = 16
	IP_ADAPTER_ADDRESSES_LH_FirstUnicastAddress_Offset = 24
	IP_ADAPTER_ADDRESSES_LH_FirstAnycastAddress_Offset = 32
	IP_ADAPTER_ADDRESSES_LH_FirstMulticastAddress_Offset = 40
	IP_ADAPTER_ADDRESSES_LH_FirstDnsServerAddress_Offset = 48
	IP_ADAPTER_ADDRESSES_LH_DnsSuffix_Offset = 56
	IP_ADAPTER_ADDRESSES_LH_Description_Offset = 64
	IP_ADAPTER_ADDRESSES_LH_FriendlyName_Offset = 72
	IP_ADAPTER_ADDRESSES_LH_PhysicalAddress_Offset = 80
	IP_ADAPTER_ADDRESSES_LH_PhysicalAddressLength_Offset = 88
	IP_ADAPTER_ADDRESSES_LH_Flags_Offset = 92
	IP_ADAPTER_ADDRESSES_LH_Mtu_Offset = 96
	IP_ADAPTER_ADDRESSES_LH_IfType_Offset = 100
	IP_ADAPTER_ADDRESSES_LH_OperStatus_Offset = 104
	IP_ADAPTER_ADDRESSES_LH_Ipv6IfIndex_Offset = 108
	IP_ADAPTER_ADDRESSES_LH_ZoneIndices_Offset = 112
	IP_ADAPTER_ADDRESSES_LH_FirstPrefix_Offset = 176
	IP_ADAPTER_ADDRESSES_LH_TransmitLinkSpeed_Offset = 184
	IP_ADAPTER_ADDRESSES_LH_ReceiveLinkSpeed_Offset = 192
	IP_ADAPTER_ADDRESSES_LH_FirstWinsServerAddress_Offset = 200
	IP_ADAPTER_ADDRESSES_LH_FirstGatewayAddress_Offset = 208
	IP_ADAPTER_ADDRESSES_LH_Ipv4Metric_Offset = 216
	IP_ADAPTER_ADDRESSES_LH_Ipv6Metric_Offset = 220
	IP_ADAPTER_ADDRESSES_LH_Luid_Offset = 224
	IP_ADAPTER_ADDRESSES_LH_Dhcpv4Server_Offset = 232
	IP_ADAPTER_ADDRESSES_LH_CompartmentId_Offset = 248
	IP_ADAPTER_ADDRESSES_LH_NetworkGuid_Offset = 252
	IP_ADAPTER_ADDRESSES_LH_ConnectionType_Offset = 268
	IP_ADAPTER_ADDRESSES_LH_TunnelType_Offset = 272
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6Server_Offset = 280
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuid_Offset = 296
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuidLength_Offset = 428
	IP_ADAPTER_ADDRESSES_LH_Dhcpv6Iaid_Offset = 432
	IP_ADAPTER_ADDRESSES_LH_FirstDnsSuffix_Offset = 440

	IP_ADAPTER_ANYCAST_ADDRESS_XP_Size = 32

	IP_ADAPTER_ANYCAST_ADDRESS_XP_Flags_Offset = 4
	IP_ADAPTER_ANYCAST_ADDRESS_XP_Next_Offset = 8
	IP_ADAPTER_ANYCAST_ADDRESS_XP_Address_Offset = 16

	SOCKET_ADDRESS_Size = 16

	SOCKET_ADDRESS_iSockaddrLength_Offset = 8
)
