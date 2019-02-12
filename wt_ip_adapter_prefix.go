/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "unsafe"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_prefix_xp
// Defined in iptypes.h
type IP_ADAPTER_PREFIX_XP struct {
	Length uint32 // Windows type: ULONG
	Flags uint32 // Windows type: DWORD
	Next *IP_ADAPTER_PREFIX_XP
	Address SOCKET_ADDRESS
	PrefixLength uint32 // Windows type: ULONG
}

// TODO: IP_ADAPTER_PREFIX and related methods probably can be removed?

// Defined in iptypes.h
type IP_ADAPTER_PREFIX IP_ADAPTER_PREFIX_XP

func (pxp *IP_ADAPTER_PREFIX_XP) nextCasted() *IP_ADAPTER_PREFIX {
	if pxp == nil {
		return nil
	} else {
		return (*IP_ADAPTER_PREFIX)(unsafe.Pointer(pxp))
	}
}
