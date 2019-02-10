/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

const (
	GUID_Size = 16

	GUID_Data2_Offset = 4
	GUID_Data3_Offset = 6
	GUID_Data4_Offset = 8

	IN6_ADDR_Size = 16

	IN_ADDR_Size = 4

	IN_ADDR_s_b2_Offset = 1
	IN_ADDR_s_b3_Offset = 2
	IN_ADDR_s_b4_Offset = 3

	IP_ADDRESS_PREFIX_Size = 32

	IP_ADDRESS_PREFIX_PrefixLength_Offset = 28

	MIB_IPFORWARD_ROW2_Size = 104

	MIB_IPFORWARD_ROW2_InterfaceIndex_Offset = 8
	MIB_IPFORWARD_ROW2_DestinationPrefix_Offset = 12
	MIB_IPFORWARD_ROW2_NextHop_Offset = 44
	MIB_IPFORWARD_ROW2_SitePrefixLength_Offset = 72
	MIB_IPFORWARD_ROW2_ValidLifetime_Offset = 76
	MIB_IPFORWARD_ROW2_PreferredLifetime_Offset = 80
	MIB_IPFORWARD_ROW2_Metric_Offset = 84
	MIB_IPFORWARD_ROW2_Protocol_Offset = 88
	MIB_IPFORWARD_ROW2_Loopback_Offset = 92
	MIB_IPFORWARD_ROW2_AutoconfigureAddress_Offset = 93
	MIB_IPFORWARD_ROW2_Publish_Offset = 94
	MIB_IPFORWARD_ROW2_Immortal_Offset = 95
	MIB_IPFORWARD_ROW2_Age_Offset = 96
	MIB_IPFORWARD_ROW2_Origin_Offset = 100

	MIB_IPFORWARD_TABLE2_Size = 112

	MIB_IPFORWARD_TABLE2_Table_Offset = 8

	MIB_UNICASTIPADDRESS_ROW_Size = 80

	MIB_UNICASTIPADDRESS_ROW_InterfaceLuid_Offset = 32
	MIB_UNICASTIPADDRESS_ROW_InterfaceIndex_Offset = 40
	MIB_UNICASTIPADDRESS_ROW_PrefixOrigin_Offset = 44
	MIB_UNICASTIPADDRESS_ROW_SuffixOrigin_Offset = 48
	MIB_UNICASTIPADDRESS_ROW_ValidLifetime_Offset = 52
	MIB_UNICASTIPADDRESS_ROW_PreferredLifetime_Offset = 56
	MIB_UNICASTIPADDRESS_ROW_OnLinkPrefixLength_Offset = 60
	MIB_UNICASTIPADDRESS_ROW_SkipAsSource_Offset = 61
	MIB_UNICASTIPADDRESS_ROW_DadState_Offset = 64
	MIB_UNICASTIPADDRESS_ROW_ScopeId_Offset = 68
	MIB_UNICASTIPADDRESS_ROW_CreationTimeStamp_Offset = 72

	SOCKADDR_IN_Size = 16

	SOCKADDR_IN_sin_port_Offset = 2
	SOCKADDR_IN_sin_addr_Offset = 4
	SOCKADDR_IN_sin_zero_Offset = 8

	SOCKADDR_IN6_LH_Size = 28

	SOCKADDR_IN6_LH_sin6_port_Offset = 2
	SOCKADDR_IN6_LH_sin6_flowinfo_Offset = 4
	SOCKADDR_IN6_LH_sin6_addr_Offset = 8
	SOCKADDR_IN6_LH_sin6_scope_id_Offset = 24

	SOCKADDR_Size = 16

	SOCKADDR_sa_data_Offset = 2
)
