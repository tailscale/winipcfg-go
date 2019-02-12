/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
// Defined in iptypes.h
type IP_ADAPTER_UNICAST_ADDRESS_LH struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *IP_ADAPTER_UNICAST_ADDRESS_LH
	Address wtSocketAddress

	PrefixOrigin IP_PREFIX_ORIGIN
	SuffixOrigin IP_SUFFIX_ORIGIN
	DadState IP_DAD_STATE

	ValidLifetime uint32 // Windows type: ULONG
	PreferredLifetime uint32 // Windows type: ULONG
	LeaseLifetime uint32 // Windows type: ULONG
	OnLinkPrefixLength uint8 // Windows type: UINT8
}
