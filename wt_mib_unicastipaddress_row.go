/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os"
	"unsafe"
)

func getWtMibUnicastipaddressRows(family AddressFamily) ([]*wtMibUnicastipaddressRow, error) {

	var pTable *wtMibUnicastipaddressTable = nil

	result := getUnicastIpAddressTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetUnicastIpAddressTable", windows.Errno(result))
	}

	addresses := make([]*wtMibUnicastipaddressRow, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibUnicastipaddressRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		addresses[i] = (*wtMibUnicastipaddressRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
	}

	return addresses, nil
}

func (wtua *wtMibUnicastipaddressRow) toUnicastAddressData() (*UnicastAddressData, error) {

	if wtua == nil {
		return nil, nil
	}

	sai, err := wtua.Address.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &UnicastAddressData{
		Address:            sai,
		InterfaceLuid:      wtua.InterfaceLuid,
		InterfaceIndex:     wtua.InterfaceIndex,
		PrefixOrigin:       wtua.PrefixOrigin,
		SuffixOrigin:       wtua.SuffixOrigin,
		ValidLifetime:      wtua.ValidLifetime,
		PreferredLifetime:  wtua.PreferredLifetime,
		OnLinkPrefixLength: wtua.OnLinkPrefixLength,
		SkipAsSource:       uint8ToBool(wtua.SkipAsSource),
		DadState:           wtua.DadState,
		ScopeId:            wtua.ScopeId,
		CreationTimeStamp:  wtua.CreationTimeStamp,
	}, nil
}

func getMatchingWtMibUnicastipaddressRow(luid uint64, ip *net.IP) (*wtMibUnicastipaddressRow, error) {

	wtas, err := getWtMibUnicastipaddressRows(AF_UNSPEC)

	if err != nil {
		return nil, err
	}

	for _, wta := range wtas {
		if wta.InterfaceLuid == luid && wta.Address.matches(ip) {
			return wta, nil
		}
	}

	return nil, nil
}

func (wtua *wtMibUnicastipaddressRow) add() error {

	if wtua == nil {
		return fmt.Errorf("wtMibUnicastipaddressRow.add() - input argument is nil")
	}

	fmt.Printf("wtMibUnicastipaddressRow add:\n%s\n", wtua.String())

	result := createUnicastIpAddressEntry(wtua)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.CreateUnicastIpAddressEntry", windows.Errno(result))
	}
}

func (wtua *wtMibUnicastipaddressRow) delete() error {

	if wtua == nil {
		return fmt.Errorf("wtMibUnicastipaddressRow.delete() - input argument is nil")
	}

	result := deleteUnicastIpAddressEntry(wtua)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.DeleteUnicastIpAddressEntry", windows.Errno(result))
	}
}

func (wtua *wtMibUnicastipaddressRow) String() string {

	if wtua == nil {
		return ""
	}

	return fmt.Sprintf(`Address: [%s]/%d
InterfaceLuid: %d
InterfaceIndex: %d
PrefixOrigin: %s
SuffixOrigin: %s
ValidLifetime: %d
PreferredLifetime: %d
SkipAsSource: %d
DadState: %s
ScopeId: %d
CreationTimeStamp: %d
`, wtua.Address.String(), wtua.OnLinkPrefixLength, wtua.InterfaceLuid, wtua.InterfaceIndex, wtua.PrefixOrigin.String(),
wtua.SuffixOrigin.String(), wtua.ValidLifetime, wtua.PreferredLifetime, wtua.SkipAsSource, wtua.DadState.String(),
wtua.SkipAsSource, wtua.CreationTimeStamp)
}