/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

const (
	if_max_string_size         = 256 // IF_MAX_STRING_SIZE defined in ifdef.h
	if_max_phys_address_length = 32  // IF_MAX_PHYS_ADDRESS_LENGTH defined in ifdef.h
)

type interfaceAndOperStatusFlagsByte uint8

const (
	hardwareInterface interfaceAndOperStatusFlagsByte = 0x01
	filterInterface   interfaceAndOperStatusFlagsByte = 0x02
	connectorPresent  interfaceAndOperStatusFlagsByte = 0x04
	notAuthenticated  interfaceAndOperStatusFlagsByte = 0x08
	notMediaConnected interfaceAndOperStatusFlagsByte = 0x10
	paused            interfaceAndOperStatusFlagsByte = 0x20
	lowPower          interfaceAndOperStatusFlagsByte = 0x40
	endPointInterface interfaceAndOperStatusFlagsByte = 0x80
)

func (wtior interfaceAndOperStatusFlagsByte) toInterfaceAndOperStatusFlags() *InterfaceAndOperStatusFlags {
	return &InterfaceAndOperStatusFlags{
		HardwareInterface: uint8ToBool(uint8(wtior & hardwareInterface)),
		FilterInterface:   uint8ToBool(uint8(wtior & filterInterface)),
		ConnectorPresent:  uint8ToBool(uint8(wtior & connectorPresent)),
		NotAuthenticated:  uint8ToBool(uint8(wtior & notAuthenticated)),
		NotMediaConnected: uint8ToBool(uint8(wtior & notMediaConnected)),
		Paused:            uint8ToBool(uint8(wtior & paused)),
		LowPower:          uint8ToBool(uint8(wtior & lowPower)),
		EndPointInterface: uint8ToBool(uint8(wtior & endPointInterface)),
	}
}

// When 'guid' is nil corresponds to GetIfTable2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getiftable2ex)
func getWtMibIfRow2s(level MibIfEntryLevel, guid *windows.GUID) ([]*wtMibIfRow2, error) {

	var pTable *wtMibIfTable2 = nil

	result := getIfTable2Ex(level, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIfTable2Ex", windows.Errno(result))
	}

	var rows []*wtMibIfRow2

	if guid == nil {
		rows = make([]*wtMibIfRow2, pTable.NumEntries, pTable.NumEntries)
	}

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIfRow2_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		// Dereferencing and rereferencing in order to force copying.
		row := *(*wtMibIfRow2)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))

		if guid == nil {
			rows[i] = &row
		} else if guidsEqual(guid, &row.InterfaceGuid) {
			rows = append(rows, &row)
		}
	}

	return rows, nil
}

// Corresponds to GetIfEntry2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex)
func getWtMibIfRow(interfaceLuid uint64, level MibIfEntryLevel) (*wtMibIfRow2, error) {

	row := wtMibIfRow2{InterfaceLuid: interfaceLuid}

	result := getIfEntry2Ex(level, &row)

	if result == 0 {
		return &row, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetIfEntry2Ex", windows.Errno(result))
	}
}

func (row *wtMibIfRow2) toIfRow() *IfRow {

	if row == nil {
		return nil
	}

	return &IfRow{
		InterfaceLuid:               row.InterfaceLuid,
		InterfaceIndex:              row.InterfaceIndex,
		InterfaceGuid:               row.InterfaceGuid,
		Alias:                       wcharToString(&row.Alias[0], if_max_string_size+1),
		Description:                 wcharToString(&row.Description[0], if_max_string_size+1),
		PhysicalAddress:             charToString(&row.PhysicalAddress[0], row.PhysicalAddressLength),
		PermanentPhysicalAddress:    charToString(&row.PermanentPhysicalAddress[0], if_max_phys_address_length),
		Mtu:                         row.Mtu,
		Type:                        row.Type,
		TunnelType:                  row.TunnelType,
		MediaType:                   row.MediaType,
		PhysicalMediumType:          row.PhysicalMediumType,
		AccessType:                  row.AccessType,
		DirectionType:               row.DirectionType,
		InterfaceAndOperStatusFlags: *row.InterfaceAndOperStatusFlags.toInterfaceAndOperStatusFlags(),
		OperStatus:                  row.OperStatus,
		AdminStatus:                 row.AdminStatus,
		MediaConnectState:           row.MediaConnectState,
		NetworkGuid:                 row.NetworkGuid,
		ConnectionType:              row.ConnectionType,
		TransmitLinkSpeed:           row.TransmitLinkSpeed,
		ReceiveLinkSpeed:            row.ReceiveLinkSpeed,
		InOctets:                    row.InOctets,
		InUcastPkts:                 row.InUcastPkts,
		InNUcastPkts:                row.InNUcastPkts,
		InDiscards:                  row.InDiscards,
		InErrors:                    row.InErrors,
		InUnknownProtos:             row.InUnknownProtos,
		InUcastOctets:               row.InUcastOctets,
		InMulticastOctets:           row.InMulticastOctets,
		InBroadcastOctets:           row.InBroadcastOctets,
		OutOctets:                   row.OutOctets,
		OutUcastPkts:                row.OutUcastPkts,
		OutNUcastPkts:               row.OutNUcastPkts,
		OutDiscards:                 row.OutDiscards,
		OutErrors:                   row.OutErrors,
		OutUcastOctets:              row.OutUcastOctets,
		OutMulticastOctets:          row.OutMulticastOctets,
		OutBroadcastOctets:          row.OutBroadcastOctets,
		OutQLen:                     row.OutQLen,
	}
}
