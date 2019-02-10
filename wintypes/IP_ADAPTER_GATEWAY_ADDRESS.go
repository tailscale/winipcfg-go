/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_gateway_address_lh
// Defined in iptypes.h
type IP_ADAPTER_GATEWAY_ADDRESS_LH struct {
	Length ULONG
	Reserved DWORD
	Next *IP_ADAPTER_GATEWAY_ADDRESS_LH
	Address SOCKET_ADDRESS
}

// Defined in iptypes.h
type IP_ADAPTER_GATEWAY_ADDRESS IP_ADAPTER_GATEWAY_ADDRESS_LH
