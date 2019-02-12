/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
// IP_ADAPTER_UNICAST_ADDRESS_LH defined in iptypes.h
type wtIpAdapterUnicastAddressLh struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *wtIpAdapterUnicastAddressLh
	Address wtSocketAddress

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState IpDadState

	ValidLifetime uint32 // Windows type: ULONG
	PreferredLifetime uint32 // Windows type: ULONG
	LeaseLifetime uint32 // Windows type: ULONG
	OnLinkPrefixLength uint8 // Windows type: UINT8
}
