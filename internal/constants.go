/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

const (
	IN_ADDR_Size = 4
	IN6_ADDR_Size = 16
	IP_ADDRESS_PREFIX_Size = 32
	MIB_IPFORWARD_ROW2_Size = 104
	MIB_IPFORWARD_TABLE2_Size = 112
	MIB_UNICASTIPADDRESS_ROW_Size = 80
	SCOPE_ID_Size = 4
	SOCKADDR_INET_Size = 28
	SOCKADDR_IN_Size = 16
	SOCKADDR_IN6_Size = 28
	SOCKADDR_IN6_LH_Size = 28

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

	MIB_IPFORWARD_TABLE2_Table_Offset = 8
)
