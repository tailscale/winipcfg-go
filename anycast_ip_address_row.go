/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// Corresponds to MIB_ANYCASTIPADDRESS_ROW defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_anycastipaddress_row).
type AnycastIpAddressRow struct {
	//
	// Key Structure.
	//
	Address        SockaddrInet
	InterfaceLuid  uint64
	InterfaceIndex uint32

	//
	// Read-Only Fields.
	//
	ScopeId uint32
}

// Returns all anycast IP addresses from the system. GetAnycastIpAddressTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddresstable).
func GetAnycastIpAddressRows(family AddressFamily) ([]*AnycastIpAddressRow, error) {

	rows, err := getWtMibAnycastipaddressRows(family)

	if err != nil {
		return nil, err
	}

	length := len(rows)

	addresses := make([]*AnycastIpAddressRow, length, length)

	for idx, row := range rows {

		address, err := row.toAnycastIpAddressRow()

		if err != nil {
			return nil, err
		}

		addresses[idx] = address
	}

	return addresses, nil
}

// Returns anycast IP address specified by the input criteria. Corresponds to GetAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddressentry).
func GetAnycastIpAddressRow(interfaceLuid uint64, ip *net.IP) (*AnycastIpAddressRow, error) {

	row, err := getWtMibAnycastipaddressRowAlt(interfaceLuid, ip)

	if err != nil {
		return nil, err
	}

	return row.toAnycastIpAddressRow()
}

// Adds new anycast IP address to system. Corresponds to CreateAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createanycastipaddressentry)
func (aia *AnycastIpAddressRow) Add() error {

	wtsainet, err := aia.Address.toWtSockaddrInet()

	if err != nil {
		return err
	}

	row := wtMibAnycastipaddressRow{
		Address:        *wtsainet,
		InterfaceLuid:  aia.InterfaceLuid,
		InterfaceIndex: aia.InterfaceIndex,
	}

	return row.add()
}

// Deletes anycast IP address from the system. Corresponds to DeleteAnycastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteanycastipaddressentry).
func (aia *AnycastIpAddressRow) Delete() error {

	wtsainet, err := aia.Address.toWtSockaddrInet()

	if err != nil {
		return err
	}

	row, err := getWtMibAnycastipaddressRow(aia.InterfaceLuid, wtsainet)

	if err == nil {
		return row.delete()
	} else {
		return err
	}
}

func (aia *AnycastIpAddressRow) String() string {
	if aia == nil {
		return "nil"
	} else {
		return fmt.Sprintf(`Address: %s
InterfaceLuid: %d
InterfaceIndex: %d
ScopeId: %d`, aia.Address.String(), aia.InterfaceLuid, aia.InterfaceIndex, aia.ScopeId)
	}
}
