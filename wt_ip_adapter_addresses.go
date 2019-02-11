/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
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

func (aa *IP_ADAPTER_ADDRESSES) getAdapterName() string {
	return charToString(aa.AdapterName)
}

func (aa *IP_ADAPTER_ADDRESSES) getFriendlyName() string {
	return wcharToString(aa.FriendlyName)
}
