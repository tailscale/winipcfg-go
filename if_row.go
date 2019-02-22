/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
)

// MIB_IF_ROW2 defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_row2)
type IfRow struct {
	//
	// Key structure.  Sorted by preference.
	//
	InterfaceLuid  uint64
	InterfaceIndex uint32

	//
	// Read-Only fields.
	//
	InterfaceGuid            windows.GUID
	Alias                    string
	Description              string
	PhysicalAddress          string
	PermanentPhysicalAddress string

	Mtu                         uint32
	Type                        IfType     // Interface Type.
	TunnelType                  TunnelType // Tunnel Type, if Type = IF_TUNNEL.
	MediaType                   NdisMedium
	PhysicalMediumType          NdisPhysicalMedium
	AccessType                  NetIfAccessType
	DirectionType               NetIfDirectionType
	InterfaceAndOperStatusFlags InterfaceAndOperStatusFlags

	OperStatus        IfOperStatus
	AdminStatus       NetIfAdminStatus
	MediaConnectState NetIfMediaConnectState
	NetworkGuid       windows.GUID
	ConnectionType    NetIfConnectionType

	//
	// Statistics.
	//
	TransmitLinkSpeed uint64
	ReceiveLinkSpeed  uint64

	InOctets           uint64
	InUcastPkts        uint64
	InNUcastPkts       uint64
	InDiscards         uint64
	InErrors           uint64
	InUnknownProtos    uint64
	InUcastOctets      uint64
	InMulticastOctets  uint64
	InBroadcastOctets  uint64
	OutOctets          uint64
	OutUcastPkts       uint64
	OutNUcastPkts      uint64
	OutDiscards        uint64
	OutErrors          uint64
	OutUcastOctets     uint64
	OutMulticastOctets uint64
	OutBroadcastOctets uint64
	OutQLen            uint64
}

// Equivalent to GetIfEntry2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex)
func GetIfRow(interfaceLuid uint64, level MibIfEntryLevel) (*IfRow, error) {

	row, err := getWtMibIfRow(interfaceLuid, level)

	if err != nil {
		return nil, err
	} else if row == nil {
		return nil, nil
	} else {
		return row.toIfRow(), nil
	}
}

// Equivalent to GetIfTable2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getiftable2ex)
func GetIfRows(level MibIfEntryLevel) ([]*IfRow, error) {

	rows, err := getWtMibIfRow2s(level)

	if err != nil {
		return nil, err
	}

	length := len(rows)

	ifrows := make([]*IfRow, length, length)

	for idx, row := range rows {
		ifrows[idx] = row.toIfRow()
	}

	return ifrows, nil
}

func (ifr *IfRow) String() string {

	if ifr == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`InterfaceLuid: %d
InterfaceIndex: %d
InterfaceGuid: %s
Alias: %s
Description: %s
PhysicalAddress: %s
PermanentPhysicalAddress: %s
Mtu: %d
Type: %s
TunnelType: %s
MediaType: %s
PhysicalMediumType: %s
AccessType: %s
DirectionType: %s
InterfaceAndOperStatusFlags:
%s
OperStatus: %s
AdminStatus: %s
MediaConnectState: %s
NetworkGuid: %s
ConnectionType: %s
TransmitLinkSpeed: %d
ReceiveLinkSpeed: %d
InOctets: %d
InUcastPkts: %d
InNUcastPkts: %d
InDiscards: %d
InErrors: %d
InUnknownProtos: %d
InUcastOctets: %d
InMulticastOctets: %d
InBroadcastOctets: %d
OutOctets: %d
OutUcastPkts: %d
OutNUcastPkts: %d
OutDiscards: %d
OutErrors: %d
OutUcastOctets: %d
OutMulticastOctets: %d
OutBroadcastOctets: %d
OutQLen: %d
`, ifr.InterfaceLuid, ifr.InterfaceIndex, guidToString(ifr.InterfaceGuid), ifr.Alias, ifr.Description,
		ifr.PhysicalAddress, ifr.PermanentPhysicalAddress, ifr.Mtu, ifr.Type.String(), ifr.TunnelType.String(),
		ifr.MediaType.String(), ifr.PhysicalMediumType.String(), ifr.AccessType.String(), ifr.DirectionType.String(),
		toIndentedText(ifr.InterfaceAndOperStatusFlags.String(), "    "), ifr.OperStatus.String(),
		ifr.AdminStatus.String(), ifr.MediaConnectState.String(), guidToString(ifr.NetworkGuid),
		ifr.ConnectionType.String(), ifr.TransmitLinkSpeed, ifr.ReceiveLinkSpeed, ifr.InOctets, ifr.InUcastPkts,
		ifr.InNUcastPkts, ifr.InDiscards, ifr.InErrors, ifr.InUnknownProtos, ifr.InUcastOctets, ifr.InMulticastOctets,
		ifr.InBroadcastOctets, ifr.OutOctets, ifr.OutUcastPkts, ifr.OutNUcastPkts, ifr.OutDiscards, ifr.OutErrors,
		ifr.OutUcastOctets, ifr.OutMulticastOctets, ifr.OutBroadcastOctets, ifr.OutQLen)
}
