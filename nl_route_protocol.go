/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_route_protocol
// NL_ROUTE_PROTOCOL defined in nldef.h
type NlRouteProtocol uint32

const (
	RouteProtocolOther   NlRouteProtocol = 1
	RouteProtocolLocal   NlRouteProtocol = 2
	RouteProtocolNetMgmt NlRouteProtocol = 3
	RouteProtocolIcmp    NlRouteProtocol = 4
	RouteProtocolEgp     NlRouteProtocol = 5
	RouteProtocolGgp     NlRouteProtocol = 6
	RouteProtocolHello   NlRouteProtocol = 7
	RouteProtocolRip     NlRouteProtocol = 8
	RouteProtocolIsIs    NlRouteProtocol = 9
	RouteProtocolEsIs    NlRouteProtocol = 10
	RouteProtocolCisco   NlRouteProtocol = 11
	RouteProtocolBbn     NlRouteProtocol = 12
	RouteProtocolOspf    NlRouteProtocol = 13
	RouteProtocolBgp     NlRouteProtocol = 14
	RouteProtocolIdpr    NlRouteProtocol = 15
	RouteProtocolEigrp   NlRouteProtocol = 16
	RouteProtocolDvmrp   NlRouteProtocol = 17
	RouteProtocolRpl     NlRouteProtocol = 18
	RouteProtocolDhcp    NlRouteProtocol = 19

	//
	// Windows-specific definitions.
	//
	NT_AUTOSTATIC        NlRouteProtocol = 10002
	NT_STATIC            NlRouteProtocol = 10006
	NT_STATIC_NON_DOD    NlRouteProtocol = 10007
)

func (protocol NlRouteProtocol) String() string {
	switch protocol {
	case RouteProtocolOther:
		return "RouteProtocolOther"
	case RouteProtocolLocal:
		return "RouteProtocolLocal"
	case RouteProtocolNetMgmt:
		return "RouteProtocolNetMgmt"
	case RouteProtocolIcmp:
		return "RouteProtocolIcmp"
	case RouteProtocolEgp:
		return "RouteProtocolEgp"
	case RouteProtocolGgp:
		return "RouteProtocolGgp"
	case RouteProtocolHello:
		return "RouteProtocolHello"
	case RouteProtocolRip:
		return "RouteProtocolRip"
	case RouteProtocolIsIs:
		return "RouteProtocolIsIs"
	case RouteProtocolEsIs:
		return "RouteProtocolEsIs"
	case RouteProtocolCisco:
		return "RouteProtocolCisco"
	case RouteProtocolBbn:
		return "RouteProtocolBbn"
	case RouteProtocolOspf:
		return "RouteProtocolOspf"
	case RouteProtocolBgp:
		return "RouteProtocolBgp"
	case RouteProtocolIdpr:
		return "RouteProtocolIdpr"
	case RouteProtocolEigrp:
		return "RouteProtocolEigrp"
	case RouteProtocolDvmrp:
		return "RouteProtocolDvmrp"
	case RouteProtocolRpl:
		return "RouteProtocolRpl"
	case RouteProtocolDhcp:
		return "RouteProtocolDhcp"
	case NT_AUTOSTATIC:
		return "NT_AUTOSTATIC"
	case NT_STATIC:
		return "NT_STATIC"
	case NT_STATIC_NON_DOD:
		return "NT_STATIC_NON_DOD"
	default:
		return fmt.Sprintf("NlRouteProtocol_UNKNOWN(%d)", protocol)
	}
}
