/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-_nl_router_discovery_behavior
// NL_ROUTER_DISCOVERY_BEHAVIOR defined in nldef.h
type NlRouterDiscoveryBehavior int32

const (
	RouterDiscoveryDisabled  NlRouterDiscoveryBehavior = 0
	RouterDiscoveryEnabled   NlRouterDiscoveryBehavior = 1
	RouterDiscoveryDhcp      NlRouterDiscoveryBehavior = 2
	RouterDiscoveryUnchanged NlRouterDiscoveryBehavior = -1
)

func (rdb NlRouterDiscoveryBehavior) String() string {
	switch rdb {
	case RouterDiscoveryDisabled:
		return "RouterDiscoveryDisabled"
	case RouterDiscoveryEnabled:
		return "RouterDiscoveryEnabled"
	case RouterDiscoveryDhcp:
		return "RouterDiscoveryDhcp"
	case RouterDiscoveryUnchanged:
		return "RouterDiscoveryUnchanged"
	default:
		return fmt.Sprintf("NlRouterDiscoveryBehavior_UNKNOWN(%d)", rdb)
	}
}
