/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

type UnicastAddressData struct {
	Address            SockaddrInet
	InterfaceLuid      uint64
	InterfaceIndex     uint32
	PrefixOrigin       NlPrefixOrigin
	SuffixOrigin       NlSuffixOrigin
	ValidLifetime      uint32
	PreferredLifetime  uint32
	OnLinkPrefixLength uint8
	SkipAsSource       bool
	DadState           NlDadState
	ScopeId            uint32
	CreationTimeStamp  int64
}

func (uad *UnicastAddressData) toWtMibUnicastipaddressRow() (*wtMibUnicastipaddressRow, error) {

	if uad == nil {
		return nil, nil
	}

	wtsai, err := uad.Address.toWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &wtMibUnicastipaddressRow{
		Address:            *wtsai,
		InterfaceLuid:      uad.InterfaceLuid,
		InterfaceIndex:     uad.InterfaceIndex,
		PrefixOrigin:       uad.PrefixOrigin,
		SuffixOrigin:       uad.SuffixOrigin,
		ValidLifetime:      uad.ValidLifetime,
		PreferredLifetime:  uad.PreferredLifetime,
		OnLinkPrefixLength: uad.OnLinkPrefixLength,
		SkipAsSource:       boolToUint8(uad.SkipAsSource),
		DadState:           uad.DadState,
		ScopeId:            uad.ScopeId,
		CreationTimeStamp:  uad.CreationTimeStamp,
	}, nil
}

func GetUnicastAddresses(family AddressFamily) ([]*UnicastAddressData, error) {

	var pTable *wtMibUnicastipaddressTable = nil

	result := getUnicastIpAddressTable(family, unsafe.Pointer(&pTable))

	if pTable != nil {
		defer freeMibTable(unsafe.Pointer(pTable))
	}

	if result != 0 {
		return nil, windows.Errno(result)
	}

	addresses := make([]*UnicastAddressData, pTable.NumEntries, pTable.NumEntries)

	pFirstRow := uintptr(unsafe.Pointer(&pTable.Table[0]))
	rowSize := uintptr(wtMibUnicastipaddressRow_Size) // Should be equal to unsafe.Sizeof(pTable.Table[0])

	for i := uint32(0); i < pTable.NumEntries; i++ {

		wta := (*wtMibUnicastipaddressRow)(unsafe.Pointer(pFirstRow + rowSize*uintptr(i)))

		address, err := wta.toMibUnicastipaddressRow()

		if err != nil {
			return nil, err
		}

		addresses[i] = address
	}

	return addresses, nil
}

//func CreateUnicastAddress(address UnicastAddressData) error {
//
//}

func (uar *UnicastAddressData) String() string {

	if uar == nil {
		return ""
	}

	return fmt.Sprintf(`Address: [%s]/%d
InterfaceLuid: %d
InterfaceIndex: %d
PrefixOrigin: %s
SuffixOrigin: %s
ValidLifetime: %d
PreferredLifetime: %d
SkipAsSource: %v
DadState: %s
ScopeId: %d
CreationTimeStamp: %d
`, uar.Address.String(), uar.OnLinkPrefixLength, uar.InterfaceLuid, uar.InterfaceIndex, uar.PrefixOrigin.String(),
		uar.SuffixOrigin.String(), uar.ValidLifetime, uar.PreferredLifetime, uar.SkipAsSource, uar.DadState.String(),
		uar.ScopeId, uar.CreationTimeStamp)
}
