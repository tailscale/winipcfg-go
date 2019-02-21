/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NET_IF_ACCESS_TYPE defined in ifdef.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_access_type)
type NetIfAccessType uint32

const (
	NET_IF_ACCESS_LOOPBACK             NetIfAccessType = 1
	NET_IF_ACCESS_BROADCAST            NetIfAccessType = 2
	NET_IF_ACCESS_POINT_TO_POINT       NetIfAccessType = 3
	NET_IF_ACCESS_POINT_TO_MULTI_POINT NetIfAccessType = 4
	NET_IF_ACCESS_MAXIMUM              NetIfAccessType = 5
)

func (niat NetIfAccessType) String() string {
	switch niat {
	case NET_IF_ACCESS_LOOPBACK:
		return "NET_IF_ACCESS_LOOPBACK"
	case NET_IF_ACCESS_BROADCAST:
		return "NET_IF_ACCESS_BROADCAST"
	case NET_IF_ACCESS_POINT_TO_POINT:
		return "NET_IF_ACCESS_POINT_TO_POINT"
	case NET_IF_ACCESS_POINT_TO_MULTI_POINT:
		return "NET_IF_ACCESS_POINT_TO_MULTI_POINT"
	case NET_IF_ACCESS_MAXIMUM:
		return "NET_IF_ACCESS_MAXIMUM"
	default:
		return fmt.Sprintf("NetIfAccessType_UNKNOWN(%d)", niat)
	}
}
