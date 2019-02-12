/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NL_ROUTE_ORIGIN defined in nldef.h
type NlRouteOrigin uint32

const (
	NlroManual              NlRouteOrigin = 0
	NlroWellKnown           NlRouteOrigin = 1
	NlroDHCP                NlRouteOrigin = 2
	NlroRouterAdvertisement NlRouteOrigin = 3
	Nlro6to4                NlRouteOrigin = 4
)

func (o NlRouteOrigin) String() string {
	switch o {
	case NlroManual:
		return "NlroManual"
	case NlroWellKnown:
		return "NlroWellKnown"
	case NlroDHCP:
		return "NlroDHCP"
	case NlroRouterAdvertisement:
		return "NlroRouterAdvertisement"
	case Nlro6to4:
		return "Nlro6to4"
	default:
		return fmt.Sprintf("NlRouteOrigin_UNKNOWN(%d)", o)
	}
}
