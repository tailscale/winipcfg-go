/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_multicast_address_xp
// Defined in iptypes.h
type IP_ADAPTER_MULTICAST_ADDRESS_XP struct {
	Length ULONG
	Flags DWORD
	Next *IP_ADAPTER_MULTICAST_ADDRESS_XP
	Address wtSocketAddress
}
