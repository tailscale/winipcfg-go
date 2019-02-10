/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_connection_type
// Defined in ifdef.h
type NET_IF_CONNECTION_TYPE uint32

const (
	NET_IF_CONNECTION_DEDICATED NET_IF_CONNECTION_TYPE = 1
	NET_IF_CONNECTION_PASSIVE   NET_IF_CONNECTION_TYPE = 2
	NET_IF_CONNECTION_DEMAND    NET_IF_CONNECTION_TYPE = 3
	NET_IF_CONNECTION_MAXIMUM   NET_IF_CONNECTION_TYPE = 4
)

func (t NET_IF_CONNECTION_TYPE) String() string {
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
		return fmt.Sprintf("NET_IF_CONNECTION_TYPE_UNKNOWN(%d)", t)
	}
}
