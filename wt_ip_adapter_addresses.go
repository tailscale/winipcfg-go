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

// IP_ADAPTER_ADDRESSES defined in iptypes.h
type wtIpAdapterAddresses wtIpAdapterAddressesLh

func (aa *wtIpAdapterAddresses) nextCasted() *wtIpAdapterAddresses {
	return (*wtIpAdapterAddresses) (unsafe.Pointer(aa.Next))
}

func (aa *wtIpAdapterAddresses) getAdapterName() string {
	return charToString(aa.AdapterName)
}

func (aa *wtIpAdapterAddresses) getFriendlyName() string {
	return wcharToString(aa.FriendlyName)
}
