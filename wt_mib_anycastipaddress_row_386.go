/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// MIB_ANYCASTIPADDRESS_ROW defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_anycastipaddress_row).
type wtMibAnycastipaddressRow struct {
	//
	// Key Structure.
	//
	Address        wtSockaddrInet

	offset1        [4]byte // Layout correction field

	InterfaceLuid  uint64 // Windows type: NET_LUID
	InterfaceIndex uint32 // Windows type: NET_IFINDEX

	//
	// Read-Only Fields.
	//
	ScopeId uint32 // Windows type: SCOPE_ID
}
