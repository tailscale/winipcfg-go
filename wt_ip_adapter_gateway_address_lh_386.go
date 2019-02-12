/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_gateway_address_lh
// Defined in iptypes.h
type IP_ADAPTER_GATEWAY_ADDRESS_LH struct {
	Length ULONG
	Reserved DWORD
	Next *IP_ADAPTER_GATEWAY_ADDRESS_LH
	Address wtSocketAddress
	// Fixing layout! I've had to add this padding to ensure the same structure size.
	correction [4]uint8
}
