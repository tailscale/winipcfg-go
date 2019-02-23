/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// MIB_UNICASTIPADDRESS_ROW defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_unicastipaddress_row)
type wtMibUnicastipaddressRow struct {
	//
	// Key Structure.
	//
	Address        wtSockaddrInet

	offset1        [4]uint8 // Layout correction field

	InterfaceLuid  uint64   // Windows type: NET_LUID
	InterfaceIndex uint32   // Windows type: NET_IFINDEX

	//
	// Read-Write Fileds.
	//
	PrefixOrigin       NlPrefixOrigin
	SuffixOrigin       NlSuffixOrigin
	ValidLifetime      uint32 // Windows type: ULONG
	PreferredLifetime  uint32 // Windows type: ULONG
	OnLinkPrefixLength uint8  // Windows type: UINT8
	SkipAsSource       uint8  // Windows type: BOOLEAN

	//
	// Read-Only Fields.
	//
	DadState          NlDadState
	ScopeId           uint32 // Windows type: ULONG
	CreationTimeStamp int64  // Windows type: LARGE_INTEGER
}
