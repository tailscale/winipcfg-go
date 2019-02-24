/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

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

func (aia *AnycastIpAddressRow) String() string {
	if aia == nil {
		return "nil"
	} else {
		return fmt.Sprintf(`Address: %s
InterfaceLuid: %d
InterfaceIndex: %d
ScopeId: %d
`, aia.Address.String(), aia.InterfaceLuid, aia.InterfaceIndex, aia.ScopeId)
	}
}
