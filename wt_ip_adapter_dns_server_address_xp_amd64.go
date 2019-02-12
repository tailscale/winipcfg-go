/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_dns_server_address_xp
// IP_ADAPTER_DNS_SERVER_ADDRESS_XP defined in iptypes.h
type wtIpAdapterDnsServerAddressXp struct {
	Length ULONG
	Reserved DWORD
	Next *wtIpAdapterDnsServerAddressXp
	Address wtSocketAddress
}
