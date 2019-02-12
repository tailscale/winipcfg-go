/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	wtIpAdapterAddressesLh_Size = 376

	wtIpAdapterAddressesLh_IfIndex_Offset = 4
	wtIpAdapterAddressesLh_Next_Offset = 8
	wtIpAdapterAddressesLh_AdapterName_Offset = 12
	wtIpAdapterAddressesLh_FirstUnicastAddress_Offset = 16
	wtIpAdapterAddressesLh_FirstAnycastAddress_Offset = 20
	wtIpAdapterAddressesLh_FirstMulticastAddress_Offset = 24
	wtIpAdapterAddressesLh_FirstDnsServerAddress_Offset = 28
	wtIpAdapterAddressesLh_DnsSuffix_Offset = 32
	wtIpAdapterAddressesLh_Description_Offset = 36
	wtIpAdapterAddressesLh_FriendlyName_Offset = 40
	wtIpAdapterAddressesLh_PhysicalAddress_Offset = 44
	wtIpAdapterAddressesLh_PhysicalAddressLength_Offset = 52
	wtIpAdapterAddressesLh_Flags_Offset = 56
	wtIpAdapterAddressesLh_Mtu_Offset = 60
	wtIpAdapterAddressesLh_IfType_Offset = 64
	wtIpAdapterAddressesLh_OperStatus_Offset = 68
	wtIpAdapterAddressesLh_Ipv6IfIndex_Offset = 72
	wtIpAdapterAddressesLh_ZoneIndices_Offset = 76
	wtIpAdapterAddressesLh_FirstPrefix_Offset = 140
	wtIpAdapterAddressesLh_TransmitLinkSpeed_Offset = 144
	wtIpAdapterAddressesLh_ReceiveLinkSpeed_Offset = 152
	wtIpAdapterAddressesLh_FirstWinsServerAddress_Offset = 160
	wtIpAdapterAddressesLh_FirstGatewayAddress_Offset = 164
	wtIpAdapterAddressesLh_Ipv4Metric_Offset = 168
	wtIpAdapterAddressesLh_Ipv6Metric_Offset = 172
	wtIpAdapterAddressesLh_Luid_Offset = 176
	wtIpAdapterAddressesLh_Dhcpv4Server_Offset = 184
	wtIpAdapterAddressesLh_CompartmentId_Offset = 192
	wtIpAdapterAddressesLh_NetworkGuid_Offset = 196
	wtIpAdapterAddressesLh_ConnectionType_Offset = 212
	wtIpAdapterAddressesLh_TunnelType_Offset = 216
	wtIpAdapterAddressesLh_Dhcpv6Server_Offset = 220
	wtIpAdapterAddressesLh_Dhcpv6ClientDuid_Offset = 228
	wtIpAdapterAddressesLh_Dhcpv6ClientDuidLength_Offset = 360
	wtIpAdapterAddressesLh_Dhcpv6Iaid_Offset = 364
	wtIpAdapterAddressesLh_FirstDnsSuffix_Offset = 368

	wtIpAdapterAnycastAddressXp_Size = 24

	wtIpAdapterAnycastAddressXp_Flags_Offset = 4
	wtIpAdapterAnycastAddressXp_Next_Offset = 8
	wtIpAdapterAnycastAddressXp_Address_Offset = 12

	wtSocketAddress_Size = 8

	wtSocketAddress_iSockaddrLength_Offset = 4

	wtIpAdapterDnsServerAddressXp_Size = 24

	wtIpAdapterDnsServerAddressXp_Reserved_Offset = 4
	wtIpAdapterDnsServerAddressXp_Next_Offset = 8
	wtIpAdapterDnsServerAddressXp_Address_Offset = 12

	wtIpAdapterDnsSuffix_Size = 516

	wtIpAdapterDnsSuffix_String_Offset = 4

	IP_ADAPTER_GATEWAY_ADDRESS_LH_Size = 24

	IP_ADAPTER_GATEWAY_ADDRESS_LH_Reserved_Offset = 4
	IP_ADAPTER_GATEWAY_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_GATEWAY_ADDRESS_LH_Address_Offset = 12

	IP_ADAPTER_MULTICAST_ADDRESS_XP_Size = 24

	IP_ADAPTER_MULTICAST_ADDRESS_XP_Flags_Offset = 4
	IP_ADAPTER_MULTICAST_ADDRESS_XP_Next_Offset = 8
	IP_ADAPTER_MULTICAST_ADDRESS_XP_Address_Offset = 12

	IP_ADAPTER_PREFIX_XP_Size = 24

	IP_ADAPTER_PREFIX_XP_Flags_Offset = 4
	IP_ADAPTER_PREFIX_XP_Next_Offset = 8
	IP_ADAPTER_PREFIX_XP_Address_Offset = 12
	IP_ADAPTER_PREFIX_XP_PrefixLength_Offset = 20

	IP_ADAPTER_UNICAST_ADDRESS_LH_Size = 48

	IP_ADAPTER_UNICAST_ADDRESS_LH_Flags_Offset = 4
	IP_ADAPTER_UNICAST_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_UNICAST_ADDRESS_LH_Address_Offset = 12
	IP_ADAPTER_UNICAST_ADDRESS_LH_PrefixOrigin_Offset = 20
	IP_ADAPTER_UNICAST_ADDRESS_LH_SuffixOrigin_Offset = 24
	IP_ADAPTER_UNICAST_ADDRESS_LH_DadState_Offset = 28
	IP_ADAPTER_UNICAST_ADDRESS_LH_ValidLifetime_Offset = 32
	IP_ADAPTER_UNICAST_ADDRESS_LH_PreferredLifetime_Offset = 36
	IP_ADAPTER_UNICAST_ADDRESS_LH_LeaseLifetime_Offset = 40
	IP_ADAPTER_UNICAST_ADDRESS_LH_OnLinkPrefixLength_Offset = 44
	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Size = 24

	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Reserved_Offset = 4
	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Address_Offset = 12
)
