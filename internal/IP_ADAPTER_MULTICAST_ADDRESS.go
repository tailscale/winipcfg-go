/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "golang.org/x/sys/windows"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_multicast_address_xp
// Defined in iptypes.h
type IP_ADAPTER_MULTICAST_ADDRESS_XP struct {
	Length ULONG
	Flags DWORD
	Next *IP_ADAPTER_MULTICAST_ADDRESS_XP
	Address windows.SocketAddress
}

// Defined in iptypes.h
type IP_ADAPTER_MULTICAST_ADDRESS IP_ADAPTER_MULTICAST_ADDRESS_XP
