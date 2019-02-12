/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_unicastipaddress_row
// Defined in netioapi.h
type MIB_UNICASTIPADDRESS_ROW struct {
	//
	// Key Structure.
	//
	Address SOCKADDR_INET
	// Layout correction field
	offset1 [4]byte
	InterfaceLuid NET_LUID
	InterfaceIndex NET_IFINDEX

	//
	// Read-Write Fileds.
	//
	PrefixOrigin NlPrefixOrigin
	SuffixOrigin NL_SUFFIX_ORIGIN
	ValidLifetime ULONG
	PreferredLifetime ULONG
	OnLinkPrefixLength UINT8
	SkipAsSource BOOLEAN

	//
	// Read-Only Fields.
	//
	DadState NlDadState
	ScopeId ULONG
	CreationTimeStamp LARGE_INTEGER
}
