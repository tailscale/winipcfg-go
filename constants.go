/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	wtIn6Addr_Size = 16

	wtInAddr_Size = 4

	wtInAddr_s_b2_Offset = 1
	wtInAddr_s_b3_Offset = 2
	wtInAddr_s_b4_Offset = 3

	wtIpAddressPrefix_Size = 32

	wtIpAddressPrefix_PrefixLength_Offset = 28

	wtMibIfRow2_Size = 1352

	wtMibIfRow2_InterfaceIndex_Offset              = 8
	wtMibIfRow2_InterfaceGuid_Offset               = 12
	wtMibIfRow2_Alias_Offset                       = 28
	wtMibIfRow2_Description_Offset                 = 542
	wtMibIfRow2_PhysicalAddressLength_Offset       = 1056
	wtMibIfRow2_PhysicalAddress_Offset             = 1060
	wtMibIfRow2_PermanentPhysicalAddress_Offset    = 1092
	wtMibIfRow2_Mtu_Offset                         = 1124
	wtMibIfRow2_Type_Offset                        = 1128
	wtMibIfRow2_TunnelType_Offset                  = 1132
	wtMibIfRow2_MediaType_Offset                   = 1136
	wtMibIfRow2_PhysicalMediumType_Offset          = 1140
	wtMibIfRow2_AccessType_Offset                  = 1144
	wtMibIfRow2_DirectionType_Offset               = 1148
	wtMibIfRow2_InterfaceAndOperStatusFlags_Offset = 1152
	wtMibIfRow2_OperStatus_Offset                  = 1156
	wtMibIfRow2_AdminStatus_Offset                 = 1160
	wtMibIfRow2_MediaConnectState_Offset           = 1164
	wtMibIfRow2_NetworkGuid_Offset                 = 1168
	wtMibIfRow2_ConnectionType_Offset              = 1184
	wtMibIfRow2_TransmitLinkSpeed_Offset           = 1192
	wtMibIfRow2_ReceiveLinkSpeed_Offset            = 1200
	wtMibIfRow2_InOctets_Offset                    = 1208
	wtMibIfRow2_InUcastPkts_Offset                 = 1216
	wtMibIfRow2_InNUcastPkts_Offset                = 1224
	wtMibIfRow2_InDiscards_Offset                  = 1232
	wtMibIfRow2_InErrors_Offset                    = 1240
	wtMibIfRow2_InUnknownProtos_Offset             = 1248
	wtMibIfRow2_InUcastOctets_Offset               = 1256
	wtMibIfRow2_InMulticastOctets_Offset           = 1264
	wtMibIfRow2_InBroadcastOctets_Offset           = 1272
	wtMibIfRow2_OutOctets_Offset                   = 1280
	wtMibIfRow2_OutUcastPkts_Offset                = 1288
	wtMibIfRow2_OutNUcastPkts_Offset               = 1296
	wtMibIfRow2_OutDiscards_Offset                 = 1304
	wtMibIfRow2_OutErrors_Offset                   = 1312
	wtMibIfRow2_OutUcastOctets_Offset              = 1320
	wtMibIfRow2_OutMulticastOctets_Offset          = 1328
	wtMibIfRow2_OutBroadcastOctets_Offset          = 1336
	wtMibIfRow2_OutQLen_Offset                     = 1344

	wtMibIfTable2_Size = 1360

	wtMibIfTable2_Table_Offset = 8

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

	wtMibIpinterfaceTable_Size = 176

	wtMibIpinterfaceTable_Table_Offset = 8

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

	wtMibUnicastipaddressTable_Size = 88

	wtMibUnicastipaddressTable_Table_Offset = 8

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
