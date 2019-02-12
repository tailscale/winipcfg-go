/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// Defined in ws2def.h as AddressFamily
type AddressFamily uint16 // Windows type: USHORT

// According to https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2
const (
	AF_UNSPEC AddressFamily = 0
	AF_INET   AddressFamily = 2
	AF_INET6  AddressFamily = 23
)

func (family AddressFamily) String() string {
	switch family {
	case AF_UNSPEC:
		return "AF_UNSPEC"
	case AF_INET:
		return "AF_INET"
	case AF_INET6:
		return "AF_INET6"
	default:
		return fmt.Sprintf("ADDRESS_FAMILY_UNKNOWN(%d)", family)
	}
}
