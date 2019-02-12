/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_multicast_address_xp
// IP_ADAPTER_MULTICAST_ADDRESS_XP defined in iptypes.h
type wtIpAdapterMulticastAddressXp struct {
	Length ULONG
	Flags DWORD
	Next *wtIpAdapterMulticastAddressXp
	Address wtSocketAddress
	// Fixing layout! I've had to add this padding to ensure the same structure size.
	correction [4]uint8
}
