/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_anycast_address_xp
// IP_ADAPTER_ANYCAST_ADDRESS_XP defined in iptypes.h
type wtIpAdapterAnycastAddressXp struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *wtIpAdapterAnycastAddressXp
	Address wtSocketAddress
}
