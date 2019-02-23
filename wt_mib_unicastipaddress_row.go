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
func getWtMibUnicastipaddressRows(family AddressFamily) ([]wtMibUnicastipaddressRow, error) {

	var pTable *wtMibUnicastipaddressTable = nil

	result := getUnicastIpAddressTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, os.NewSyscallError("iphlpapi.GetUnicastIpAddressTable", windows.Errno(result))
	}

	addresses := make([]wtMibUnicastipaddressRow, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibUnicastipaddressRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {
		addresses[i] = *(*wtMibUnicastipaddressRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))
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
			return &wta, nil
		}
	}

	return nil, nil
}

func addWtMibUnicastipaddressRow(ifc *Interface, ipnet *net.IPNet) error {

	if ifc == nil || ipnet == nil {
		return fmt.Errorf("addWtMibUnicastipaddressRow() - some of the input arguments is nil")
	}

	wtsa, err := createWtSockaddrInet(&ipnet.IP, 0)

	if err != nil {
		return err
	}

	row := wtMibUnicastipaddressRow{}

	_ = initializeUnicastIpAddressEntry(&row)

	//fmt.Printf("wtMibUnicastipaddressRow initialized to:\n%s\n", row.String())

	ones, _ := ipnet.Mask.Size()

	row.InterfaceLuid = ifc.Luid
	row.InterfaceIndex = ifc.Index
	row.Address = *wtsa
	row.OnLinkPrefixLength = uint8(ones)

	//fmt.Printf("wtMibUnicastipaddressRow to add:\n%s\n", row.String())

	result := createUnicastIpAddressEntry(&row)

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

func (wtua *wtMibUnicastipaddressRow) equal(other *wtMibUnicastipaddressRow) bool {

	if wtua == nil || other == nil {
		return false
	}

	return wtua.InterfaceLuid == other.InterfaceLuid && wtua.InterfaceIndex == other.InterfaceIndex &&
		wtua.PrefixOrigin == other.PrefixOrigin && wtua.SuffixOrigin == other.SuffixOrigin &&
		wtua.ValidLifetime == other.ValidLifetime && wtua.PreferredLifetime == other.PreferredLifetime &&
		wtua.OnLinkPrefixLength == other.OnLinkPrefixLength && wtua.SkipAsSource == other.SkipAsSource &&
		wtua.DadState == other.DadState && wtua.ScopeId == other.ScopeId &&
		wtua.CreationTimeStamp == other.CreationTimeStamp && wtua.Address.equal(&other.Address)
}

func (wtua *wtMibUnicastipaddressRow) String() string {

	if wtua == nil {
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
CreationTimeStamp: %d
`, wtua.Address.String(), wtua.OnLinkPrefixLength, wtua.InterfaceLuid, wtua.InterfaceIndex, wtua.PrefixOrigin.String(),
		wtua.SuffixOrigin.String(), wtua.ValidLifetime, wtua.PreferredLifetime, wtua.SkipAsSource, wtua.DadState.String(),
		wtua.SkipAsSource, wtua.CreationTimeStamp)
}
