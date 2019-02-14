/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	windowsGuid_Size = 16

	windowsGuid_Data2_Offset = 4
	windowsGuid_Data3_Offset = 6
	windowsGuid_Data4_Offset = 8

	wtIn6Addr_Size = 16

	wtInAddr_Size = 4

	wtInAddr_s_b2_Offset = 1
	wtInAddr_s_b3_Offset = 2
	wtInAddr_s_b4_Offset = 3

	wtIpAddressPrefix_Size = 32

	wtIpAddressPrefix_PrefixLength_Offset = 28

	wtMibIpinterfaceRow_Size = 168

	wtMibIpinterfaceRow_InterfaceLuid_Offset                        = 8
	wtMibIpinterfaceRow_InterfaceIndex_Offset                       = 16
	wtMibIpinterfaceRow_MaxReassemblySize_Offset                    = 20
	wtMibIpinterfaceRow_InterfaceIdentifier_Offset                  = 24
	wtMibIpinterfaceRow_MinRouterAdvertisementInterval_Offset       = 32
	wtMibIpinterfaceRow_MaxRouterAdvertisementInterval_Offset       = 36
	wtMibIpinterfaceRow_AdvertisingEnabled_Offset                   = 40
	wtMibIpinterfaceRow_ForwardingEnabled_Offset                    = 41
	wtMibIpinterfaceRow_WeakHostSend_Offset                         = 42
	wtMibIpinterfaceRow_WeakHostReceive_Offset                      = 43
	wtMibIpinterfaceRow_UseAutomaticMetric_Offset                   = 44
	wtMibIpinterfaceRow_UseNeighborUnreachabilityDetection_Offset   = 45
	wtMibIpinterfaceRow_ManagedAddressConfigurationSupported_Offset = 46
	wtMibIpinterfaceRow_OtherStatefulConfigurationSupported_Offset  = 47
	wtMibIpinterfaceRow_AdvertiseDefaultRoute_Offset                = 48
	wtMibIpinterfaceRow_RouterDiscoveryBehavior_Offset              = 52
	wtMibIpinterfaceRow_DadTransmits_Offset                         = 56
	wtMibIpinterfaceRow_BaseReachableTime_Offset                    = 60
	wtMibIpinterfaceRow_RetransmitTime_Offset                       = 64
	wtMibIpinterfaceRow_PathMtuDiscoveryTimeout_Offset              = 68
	wtMibIpinterfaceRow_LinkLocalAddressBehavior_Offset             = 72
	wtMibIpinterfaceRow_LinkLocalAddressTimeout_Offset              = 76
	wtMibIpinterfaceRow_ZoneIndices_Offset                          = 80
	wtMibIpinterfaceRow_SitePrefixLength_Offset                     = 144
	wtMibIpinterfaceRow_Metric_Offset                               = 148
	wtMibIpinterfaceRow_NlMtu_Offset                                = 152
	wtMibIpinterfaceRow_Connected_Offset                            = 156
	wtMibIpinterfaceRow_SupportsWakeUpPatterns_Offset               = 157
	wtMibIpinterfaceRow_SupportsNeighborDiscovery_Offset            = 158
	wtMibIpinterfaceRow_SupportsRouterDiscovery_Offset              = 159
	wtMibIpinterfaceRow_ReachableTime_Offset                        = 160
	wtMibIpinterfaceRow_TransmitOffload_Offset                      = 164
	wtMibIpinterfaceRow_ReceiveOffload_Offset                       = 165
	wtMibIpinterfaceRow_DisableDefaultRoutes_Offset                 = 166

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
