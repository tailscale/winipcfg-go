/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"net"
	"os"
	"unsafe"
)

// Uses GetAnycastIpAddressTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddresstable)
func getWtMibAnycastipaddressRows(family AddressFamily) ([]*wtMibAnycastipaddressRow, error) {

	var pTable *wtMibAnycastipaddressTable = nil

	result := getAnycastIpAddressTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetAnycastIpAddressTable", windows.Errno(result))
	}

	addresses := make([]*wtMibAnycastipaddressRow, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibAnycastipaddressRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		// Dereferencing and rereferencing in order to force copying.
		row := *(*wtMibAnycastipaddressRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
		addresses[i] = &row
	}

	return addresses, nil
}

func getWtMibAnycastipaddressRowAlt(interfaceLuid uint64, ip *net.IP) (*wtMibAnycastipaddressRow, error) {

	wtsainet, err := createWtSockaddrInet(ip, 0)

	if err == nil {
		return getWtMibAnycastipaddressRow(interfaceLuid, wtsainet)
	} else {
		return nil, err
	}
}

// Corresponds to GetAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddressentry)
func getWtMibAnycastipaddressRow(interfaceLuid uint64, wtsainet *wtSockaddrInet) (*wtMibAnycastipaddressRow, error) {

	row := &wtMibAnycastipaddressRow{
		Address:       *wtsainet,
		InterfaceLuid: interfaceLuid,
	}

	result := getAnycastIpAddressEntry(row)

	if result == 0 {
		return row, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetAnycastIpAddressEntry", windows.Errno(result))
	}
}

// Uses CreateAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createanycastipaddressentry)
func (wtaia *wtMibAnycastipaddressRow) add() error {

	result := createAnycastIpAddressEntry(wtaia)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.CreateAnycastIpAddressEntry", windows.Errno(result))
	}
}

// Uses DeleteAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteanycastipaddressentry)
func (wtaia *wtMibAnycastipaddressRow) delete() error {

	result := deleteAnycastIpAddressEntry(wtaia)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.DeleteAnycastIpAddressEntry", windows.Errno(result))
	}
}

func (wtaia *wtMibAnycastipaddressRow) toAnycastIpAddressRow() (*AnycastIpAddressRow, error) {

	if wtaia == nil {
		return nil, nil
	}

	sainet, err := wtaia.Address.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &AnycastIpAddressRow{
		Address:        *sainet,
		InterfaceLuid:  wtaia.InterfaceLuid,
		InterfaceIndex: wtaia.InterfaceIndex,
		ScopeId:        wtaia.ScopeId,
	}, nil
}
