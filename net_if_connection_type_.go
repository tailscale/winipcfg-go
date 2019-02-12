/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_connection_type
// NET_IF_CONNECTION_TYPE defined in ifdef.h
type NetIfConnectionType uint32

const (
	NET_IF_CONNECTION_DEDICATED NetIfConnectionType = 1
	NET_IF_CONNECTION_PASSIVE   NetIfConnectionType = 2
	NET_IF_CONNECTION_DEMAND    NetIfConnectionType = 3
	NET_IF_CONNECTION_MAXIMUM   NetIfConnectionType = 4
)

func (t NetIfConnectionType) String() string {
	switch t {
	case NET_IF_CONNECTION_DEDICATED:
		return "NET_IF_CONNECTION_DEDICATED"
	case NET_IF_CONNECTION_PASSIVE:
		return "NET_IF_CONNECTION_PASSIVE"
	case NET_IF_CONNECTION_DEMAND:
		return "NET_IF_CONNECTION_DEMAND"
	case NET_IF_CONNECTION_MAXIMUM:
		return "NET_IF_CONNECTION_MAXIMUM"
	default:
		return fmt.Sprintf("NetIfConnectionType_UNKNOWN(%d)", t)
	}
}
