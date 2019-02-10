/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

// Defined in iptypes.h
const (
	MAX_ADAPTER_ADDRESS_LENGTH = 8
	MAX_DHCPV6_DUID_LENGTH = 130
)

// Defined in iptypes.h
type IP_ADAPTER_ADDRESSES IP_ADAPTER_ADDRESSES_LH

func (aa *IP_ADAPTER_ADDRESSES) NextCasted() *IP_ADAPTER_ADDRESSES {
	return (*IP_ADAPTER_ADDRESSES) (unsafe.Pointer(aa.Next))
}

func (aa *IP_ADAPTER_ADDRESSES) Name() string {
	return windows.UTF16ToString((*(*[10000]uint16)(unsafe.Pointer(aa.FriendlyName)))[:])
}
