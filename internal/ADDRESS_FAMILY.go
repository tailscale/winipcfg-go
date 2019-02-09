/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// Defined in ws2def.h
type ADDRESS_FAMILY USHORT

// According to https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2
const (
	AF_UNSPEC ADDRESS_FAMILY = 0
	AF_INET ADDRESS_FAMILY = 2
	AF_INET6 ADDRESS_FAMILY = 23
)

func (family ADDRESS_FAMILY) String() string {
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
