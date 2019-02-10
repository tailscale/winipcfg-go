/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import "golang.org/x/sys/windows"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_prefix_xp
// Defined in iptypes.h
type IP_ADAPTER_PREFIX_XP struct {
	Length ULONG
	Flags DWORD
	Next *IP_ADAPTER_PREFIX_XP
	Address windows.SocketAddress
	PrefixLength ULONG
}

// Defined in iptypes.h
type IP_ADAPTER_PREFIX IP_ADAPTER_PREFIX_XP
