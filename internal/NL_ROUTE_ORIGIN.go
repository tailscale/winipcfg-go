/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

type NL_ROUTE_ORIGIN uint32

const (
	NlroManual              NL_ROUTE_ORIGIN = 0
	NlroWellKnown           NL_ROUTE_ORIGIN = 1
	NlroDHCP                NL_ROUTE_ORIGIN = 2
	NlroRouterAdvertisement NL_ROUTE_ORIGIN = 3
	Nlro6to4                NL_ROUTE_ORIGIN = 4
)

func (o NL_ROUTE_ORIGIN) String() string {
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
		return fmt.Sprintf("NL_ROUTE_ORIGIN_UNKNOWN(%d)", o)
	}
}
