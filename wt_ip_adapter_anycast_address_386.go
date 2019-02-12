/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_anycast_address_xp
// Defined in iptypes.h
type IP_ADAPTER_ANYCAST_ADDRESS_XP struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *IP_ADAPTER_ANYCAST_ADDRESS_XP
	Address wtSocketAddress
	// Fixing layout! I've had to add this padding to ensure the same structure size.
	correction [4]uint8
}
