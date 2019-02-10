/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "golang.org/x/sys/windows"

type IP_ADAPTER_UNICAST_ADDRESS_LH struct {
	Length ULONG
	Flags DWORD
	Next *IP_ADAPTER_UNICAST_ADDRESS_LH
	Address windows.SocketAddress

	PrefixOrigin IP_PREFIX_ORIGIN
	SuffixOrigin IP_SUFFIX_ORIGIN
	DadState IP_DAD_STATE

	ValidLifetime ULONG
	PreferredLifetime ULONG
	LeaseLifetime ULONG
	OnLinkPrefixLength UINT8
}
