/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	wtGuid_Size = 16

	wtGuid_Data2_Offset = 4
	wtGuid_Data3_Offset = 6
	wtGuid_Data4_Offset = 8

	wtIn6Addr_Size = 16

	wtInAddr_Size = 4

	wtInAddr_s_b2_Offset = 1
	wtInAddr_s_b3_Offset = 2
	wtInAddr_s_b4_Offset = 3

	wtIpAddressPrefix_Size = 32

	wtIpAddressPrefix_PrefixLength_Offset = 28

	wtMibIpforwardRow2_Size = 104

	wtMibIpforwardRow2_InterfaceIndex_Offset       = 8
	wtMibIpforwardRow2_DestinationPrefix_Offset    = 12
	wtMibIpforwardRow2_NextHop_Offset              = 44
	wtMibIpforwardRow2_SitePrefixLength_Offset     = 72
	wtMibIpforwardRow2_ValidLifetime_Offset        = 76
	wtMibIpforwardRow2_PreferredLifetime_Offset    = 80
	wtMibIpforwardRow2_Metric_Offset               = 84
	wtMibIpforwardRow2_Protocol_Offset             = 88
	wtMibIpforwardRow2_Loopback_Offset             = 92
	wtMibIpforwardRow2_AutoconfigureAddress_Offset = 93
	wtMibIpforwardRow2_Publish_Offset              = 94
	wtMibIpforwardRow2_Immortal_Offset             = 95
	wtMibIpforwardRow2_Age_Offset                  = 96
	wtMibIpforwardRow2_Origin_Offset               = 100

	wtMibIpforwardTable2_Size = 112

	wtMibIpforwardTable2_Table_Offset = 8

	wtMibUnicastipaddressRow_Size = 80

	wtMibUnicastipaddressRow_InterfaceLuid_Offset      = 32
	wtMibUnicastipaddressRow_InterfaceIndex_Offset     = 40
	wtMibUnicastipaddressRow_PrefixOrigin_Offset       = 44
	wtMibUnicastipaddressRow_SuffixOrigin_Offset       = 48
	wtMibUnicastipaddressRow_ValidLifetime_Offset      = 52
	wtMibUnicastipaddressRow_PreferredLifetime_Offset  = 56
	wtMibUnicastipaddressRow_OnLinkPrefixLength_Offset = 60
	wtMibUnicastipaddressRow_SkipAsSource_Offset       = 61
	wtMibUnicastipaddressRow_DadState_Offset           = 64
	wtMibUnicastipaddressRow_ScopeId_Offset            = 68
	wtMibUnicastipaddressRow_CreationTimeStamp_Offset  = 72

	wtSockaddrIn_Size = 16

	wtSockaddrIn_sin_port_Offset = 2
	wtSockaddrIn_sin_addr_Offset = 4
	wtSockaddrIn_sin_zero_Offset = 8

	wtSockaddrIn6Lh_Size = 28

	wtSockaddrIn6Lh_sin6_port_Offset     = 2
	wtSockaddrIn6Lh_sin6_flowinfo_Offset = 4
	wtSockaddrIn6Lh_sin6_addr_Offset     = 8
	wtSockaddrIn6Lh_sin6_scope_id_Offset = 24

	wtSockaddr_Size = 16

	wtSockaddr_sa_data_Offset = 2
)
