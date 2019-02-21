/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_multicast_address_xp
// IP_ADAPTER_MULTICAST_ADDRESS_XP defined in iptypes.h
type wtIpAdapterMulticastAddressXp struct {
	Length     uint32 // Windows type: ULONG
	Flags      uint32 // Windows type: DWORD
	Next       *wtIpAdapterMulticastAddressXp
	Address    wtSocketAddress
	correction [4]uint8 // Layout correction field
}
