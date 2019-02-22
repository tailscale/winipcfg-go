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
	hardwareInterface interfaceAndOperStatusFlagsByte = 1
	filterInterface   interfaceAndOperStatusFlagsByte = 2
	connectorPresent  interfaceAndOperStatusFlagsByte = 4
	notAuthenticated  interfaceAndOperStatusFlagsByte = 8
	notMediaConnected interfaceAndOperStatusFlagsByte = 16
	paused            interfaceAndOperStatusFlagsByte = 32
	lowPower          interfaceAndOperStatusFlagsByte = 64
	endPointInterface interfaceAndOperStatusFlagsByte = 128
)

func (wtior interfaceAndOperStatusFlagsByte) toInterfaceAndOperStatusFlags() *InterfaceAndOperStatusFlags {
	return &InterfaceAndOperStatusFlags{
		HardwareInterface: uint8ToBool(uint8(wtior) & uint8(hardwareInterface)),
		FilterInterface:   uint8ToBool(uint8(wtior) & uint8(filterInterface)),
		ConnectorPresent:  uint8ToBool(uint8(wtior) & uint8(connectorPresent)),
		NotAuthenticated:  uint8ToBool(uint8(wtior) & uint8(notAuthenticated)),
		NotMediaConnected: uint8ToBool(uint8(wtior) & uint8(notMediaConnected)),
		Paused:            uint8ToBool(uint8(wtior) & uint8(paused)),
		LowPower:          uint8ToBool(uint8(wtior) & uint8(lowPower)),
		EndPointInterface: uint8ToBool(uint8(wtior) & uint8(endPointInterface)),
	}
}

// Equivalent to GetIfTable2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getiftable2ex)
func getWtMibIfRow2s(level MibIfEntryLevel) ([]wtMibIfRow2, error) {

	var pTable *wtMibIfTable2 = nil

	result := getIfTable2Ex(level, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetIfTable2Ex", windows.Errno(result))
	}

	rows := make([]wtMibIfRow2, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibIfRow2_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		rows[i] = *(*wtMibIfRow2)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
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
