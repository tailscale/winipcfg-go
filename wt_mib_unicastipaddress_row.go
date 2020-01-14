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

// Corresponds to GetUnicastIpAddressTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddresstable)
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
		// Dereferencing and rereferencing in order to force copying.
		row := *(*wtMibUnicastipaddressRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
		addresses[i] = &row
	}

	return addresses, nil
}

// Corresponds to GetUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddressentry)
func getWtMibUnicastipaddressRow(interfaceLuid uint64, ip *net.IP) (*wtMibUnicastipaddressRow, error) {

	wtsainet, err := createWtSockaddrInet(ip, 0)

	if err != nil {
		return nil, err
	}

	row := wtMibUnicastipaddressRow{Address: *wtsainet, InterfaceLuid: interfaceLuid}

	result := getUnicastIpAddressEntry(&row)

	if result == 0 {
		return &row, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.GetUnicastIpAddressEntry", windows.Errno(result))
	}
}

func getMatchingWtMibUnicastipaddressRow(ip *net.IP) (*wtMibUnicastipaddressRow, error) {

	wtas, err := getWtMibUnicastipaddressRows(AF_UNSPEC)

	if err != nil {
		return nil, err
	}

	for _, wta := range wtas {
		if wta.Address.matches(ip) {
			return wta, nil
		}
	}

	return nil, nil
}

// Uses InitializeUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-initializeunicastipaddressentry).
func getInitializedWtMibUnicastipaddressRow(interfaceLuid uint64) *wtMibUnicastipaddressRow {

	row := wtMibUnicastipaddressRow{InterfaceLuid: interfaceLuid}

	_ = initializeUnicastIpAddressEntry(&row)

	row.InterfaceLuid = interfaceLuid

	return &row
}

func createAndAddWtMibUnicastipaddressRow(interfaceLuid uint64, ipnet *net.IPNet) error {

	wtsainet, err := createWtSockaddrInet(&ipnet.IP, 0)

	if err != nil {
		return err
	}

	row := getInitializedWtMibUnicastipaddressRow(interfaceLuid)

	row.Address = *wtsainet

	ones, _ := ipnet.Mask.Size()

	row.OnLinkPrefixLength = uint8(ones)

	return row.add()
}

// Uses CreateUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createunicastipaddressentry)
func (row *wtMibUnicastipaddressRow) add() error {

	result := createUnicastIpAddressEntry(row)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.CreateUnicastIpAddressEntry: " + row.Address.String(), windows.Errno(result))
	}
}

// Corresponds to SetUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setunicastipaddressentry)
func (row *wtMibUnicastipaddressRow) set() error {

	result := setUnicastIpAddressEntry(row)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.SetUnicastIpAddressEntry", windows.Errno(result))
	}
}

// Corresponds to DeleteUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteunicastipaddressentry)
func (row *wtMibUnicastipaddressRow) delete() error {

	result := deleteUnicastIpAddressEntry(row)

	if result == 0 {
		return nil
	} else {
		return os.NewSyscallError("iphlpapi.DeleteUnicastIpAddressEntry", windows.Errno(result))
	}
}

func (row *wtMibUnicastipaddressRow) toUnicastIpAddressRow() (*UnicastIpAddressRow, error) {

	if row == nil {
		return nil, nil
	}

	sai, err := row.Address.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &UnicastIpAddressRow{
		Address:            sai,
		InterfaceLuid:      row.InterfaceLuid,
		InterfaceIndex:     row.InterfaceIndex,
		PrefixOrigin:       row.PrefixOrigin,
		SuffixOrigin:       row.SuffixOrigin,
		ValidLifetime:      row.ValidLifetime,
		PreferredLifetime:  row.PreferredLifetime,
		OnLinkPrefixLength: row.OnLinkPrefixLength,
		SkipAsSource:       uint8ToBool(row.SkipAsSource),
		DadState:           row.DadState,
		ScopeId:            row.ScopeId,
		CreationTimeStamp:  row.CreationTimeStamp,
	}, nil
}

func (row *wtMibUnicastipaddressRow) equal(other *wtMibUnicastipaddressRow) bool {

	if row == nil || other == nil {
		return false
	}

	return row.InterfaceLuid == other.InterfaceLuid && row.InterfaceIndex == other.InterfaceIndex &&
		row.PrefixOrigin == other.PrefixOrigin && row.SuffixOrigin == other.SuffixOrigin &&
		row.ValidLifetime == other.ValidLifetime && row.PreferredLifetime == other.PreferredLifetime &&
		row.OnLinkPrefixLength == other.OnLinkPrefixLength && row.SkipAsSource == other.SkipAsSource &&
		row.DadState == other.DadState && row.ScopeId == other.ScopeId &&
		row.CreationTimeStamp == other.CreationTimeStamp && row.Address.equal(&other.Address)
}

func (row *wtMibUnicastipaddressRow) String() string {

	if row == nil {
		return "<nil>"
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
CreationTimeStamp: %d`, row.Address.String(), row.OnLinkPrefixLength, row.InterfaceLuid, row.InterfaceIndex,
		row.PrefixOrigin.String(), row.SuffixOrigin.String(), row.ValidLifetime, row.PreferredLifetime,
		row.SkipAsSource, row.DadState.String(), row.SkipAsSource, row.CreationTimeStamp)
}
