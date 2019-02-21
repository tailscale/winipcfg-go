/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NET_IF_DIRECTION_TYPE defined in ifdef.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-net_if_direction_type)
type NetIfDirectionType uint32

const (
	NET_IF_DIRECTION_SENDRECEIVE NetIfDirectionType = 0
	NET_IF_DIRECTION_SENDONLY    NetIfDirectionType = 1
	NET_IF_DIRECTION_RECEIVEONLY NetIfDirectionType = 2
	NET_IF_DIRECTION_MAXIMUM     NetIfDirectionType = 3
)

func (nidt NetIfDirectionType) String() string {
	switch nidt {
	case NET_IF_DIRECTION_SENDRECEIVE:
		return "NET_IF_DIRECTION_SENDRECEIVE"
	case NET_IF_DIRECTION_SENDONLY:
		return "NET_IF_DIRECTION_SENDONLY"
	case NET_IF_DIRECTION_RECEIVEONLY:
		return "NET_IF_DIRECTION_RECEIVEONLY"
	case NET_IF_DIRECTION_MAXIMUM:
		return "NET_IF_DIRECTION_MAXIMUM"
	default:
		return fmt.Sprintf("NetIfDirectionType_UNKNOWN(%d)", nidt)
	}
}
