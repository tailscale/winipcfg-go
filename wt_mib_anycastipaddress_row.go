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

func (wtaia *wtMibAnycastipaddressRow) toAnycastIpAddressRow() (*AnycastIpAddressRow, error) {

	if wtaia == nil {
		return nil, nil
	}

	sainet, err := wtaia.Address.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &AnycastIpAddressRow{
		Address: *sainet,
		InterfaceLuid: wtaia.InterfaceLuid,
		InterfaceIndex: wtaia.InterfaceIndex,
		ScopeId: wtaia.ScopeId,
	}, nil
}
