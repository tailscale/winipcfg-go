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
	if aa == nil {
		return nil
	} else {
		return (*wtIpAdapterAddresses)(unsafe.Pointer(aa.Next))
	}
}

func (aa *wtIpAdapterAddresses) getAdapterName() string {
	if aa == nil {
		return ""
	} else {
		return charToString(aa.AdapterName)
	}
}

func (aa *wtIpAdapterAddresses) getFriendlyName() string {
	if aa == nil {
		return ""
	} else {
		return wcharToString(aa.FriendlyName)
	}
}
