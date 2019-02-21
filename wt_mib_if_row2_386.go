/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "golang.org/x/sys/windows"

// MIB_IF_ROW2 defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_row2)
type wtMibIfRow2 struct {
	//
	// Key structure.  Sorted by preference.
	//
	InterfaceLuid  uint64 // Windows type: NET_LUID
	InterfaceIndex uint32 // Windows type: NET_IFINDEX

	//
	// Read-Only fields.
	//
	InterfaceGuid            windows.GUID                      // Windows type: GUID
	Alias                    [if_max_string_size + 1]uint16    // Windows type: WCHAR
	Description              [if_max_string_size + 1]uint16    // Windows type: WCHAR
	PhysicalAddressLength    uint32                            // Windows type: ULONG
	PhysicalAddress          [if_max_phys_address_length]uint8 // Windows type: UCHAR
	PermanentPhysicalAddress [if_max_phys_address_length]uint8 // Windows type: UCHAR

	Mtu                         uint32     // Windows type: ULONG
	Type                        IfType     // Interface Type.
	TunnelType                  TunnelType // Tunnel Type, if Type = IF_TUNNEL.
	MediaType                   NdisMedium
	PhysicalMediumType          NdisPhysicalMedium
	AccessType                  NetIfAccessType
	DirectionType               NetIfDirectionType
	InterfaceAndOperStatusFlags interfaceAndOperStatusFlagsByte

	OperStatus        IfOperStatus
	AdminStatus       NetIfAdminStatus
	MediaConnectState NetIfMediaConnectState
	NetworkGuid       windows.GUID // Windows type: NET_IF_NETWORK_GUID
	ConnectionType    NetIfConnectionType

	offset1 [4]byte // Layout correction field

	//
	// Statistics.
	//
	TransmitLinkSpeed uint64 // Windows type: ULONG64
	ReceiveLinkSpeed  uint64 // Windows type: ULONG64

	InOctets           uint64 // Windows type: ULONG64
	InUcastPkts        uint64 // Windows type: ULONG64
	InNUcastPkts       uint64 // Windows type: ULONG64
	InDiscards         uint64 // Windows type: ULONG64
	InErrors           uint64 // Windows type: ULONG64
	InUnknownProtos    uint64 // Windows type: ULONG64
	InUcastOctets      uint64 // Windows type: ULONG64
	InMulticastOctets  uint64 // Windows type: ULONG64
	InBroadcastOctets  uint64 // Windows type: ULONG64
	OutOctets          uint64 // Windows type: ULONG64
	OutUcastPkts       uint64 // Windows type: ULONG64
	OutNUcastPkts      uint64 // Windows type: ULONG64
	OutDiscards        uint64 // Windows type: ULONG64
	OutErrors          uint64 // Windows type: ULONG64
	OutUcastOctets     uint64 // Windows type: ULONG64
	OutMulticastOctets uint64 // Windows type: ULONG64
	OutBroadcastOctets uint64 // Windows type: ULONG64
	OutQLen            uint64 // Windows type: ULONG64
}
