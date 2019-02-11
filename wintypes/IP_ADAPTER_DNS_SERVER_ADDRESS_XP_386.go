/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_dns_server_address_xp
// Defined in iptypes.h
type IP_ADAPTER_DNS_SERVER_ADDRESS_XP struct {
	Length ULONG
	Reserved DWORD
	Next *IP_ADAPTER_DNS_SERVER_ADDRESS_XP
	Address SOCKET_ADDRESS
	// Fixing layout! I've had to add this padding to ensure the same structure size.
	correction [4]uint8
}
