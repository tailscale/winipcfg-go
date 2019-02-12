/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "unsafe"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_prefix_xp
// IP_ADAPTER_PREFIX_XP defined in iptypes.h
type wtIpAdapterPrefixXp struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *wtIpAdapterPrefixXp
	Address wtSocketAddress
	PrefixLength uint32 // Windows type: ULONG
}

// TODO: IP_ADAPTER_PREFIX and related methods probably can be removed?

// Defined in iptypes.h
type IP_ADAPTER_PREFIX wtIpAdapterPrefixXp

func (pxp *wtIpAdapterPrefixXp) nextCasted() *IP_ADAPTER_PREFIX {
	if pxp == nil {
		return nil
	} else {
		return (*IP_ADAPTER_PREFIX)(unsafe.Pointer(pxp))
	}
}
