/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_unicastipaddress_row
// Defined in netioapi.h
type MIB_UNICASTIPADDRESS_ROW struct {
	//
	// Key Structure.
	//
	Address SOCKADDR_INET
	InterfaceLuid NET_LUID
	InterfaceIndex NET_IFINDEX

	//
	// Read-Write Fileds.
	//
	PrefixOrigin NL_PREFIX_ORIGIN
	SuffixOrigin NL_SUFFIX_ORIGIN
	ValidLifetime ULONG
	PreferredLifetime ULONG
	OnLinkPrefixLength UINT8
	SkipAsSource BOOLEAN

	//
	// Read-Only Fields.
	//
	DadState NL_DAD_STATE
	ScopeId SCOPE_ID
	CreationTimeStamp LARGE_INTEGER
}
