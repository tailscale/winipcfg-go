/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

type NL_ROUTE_PROTOCOL uint32

// According to https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_route_protocol (defined in nldef.h)
const (
	RouteProtocolOther   NL_ROUTE_PROTOCOL = 1
	RouteProtocolLocal   NL_ROUTE_PROTOCOL = 2
	RouteProtocolNetMgmt NL_ROUTE_PROTOCOL = 3
	RouteProtocolIcmp    NL_ROUTE_PROTOCOL = 4
	RouteProtocolEgp     NL_ROUTE_PROTOCOL = 5
	RouteProtocolGgp     NL_ROUTE_PROTOCOL = 6
	RouteProtocolHello   NL_ROUTE_PROTOCOL = 7
	RouteProtocolRip     NL_ROUTE_PROTOCOL = 8
	RouteProtocolIsIs    NL_ROUTE_PROTOCOL = 9
	RouteProtocolEsIs    NL_ROUTE_PROTOCOL = 10
	RouteProtocolCisco   NL_ROUTE_PROTOCOL = 11
	RouteProtocolBbn     NL_ROUTE_PROTOCOL = 12
	RouteProtocolOspf    NL_ROUTE_PROTOCOL = 13
	RouteProtocolBgp     NL_ROUTE_PROTOCOL = 14
	RouteProtocolIdpr    NL_ROUTE_PROTOCOL = 15
	RouteProtocolEigrp   NL_ROUTE_PROTOCOL = 16
	RouteProtocolDvmrp   NL_ROUTE_PROTOCOL = 17
	RouteProtocolRpl     NL_ROUTE_PROTOCOL = 18
	RouteProtocolDhcp    NL_ROUTE_PROTOCOL = 19

	//
	// Windows-specific definitions.
	//
	NT_AUTOSTATIC        NL_ROUTE_PROTOCOL = 10002
	NT_STATIC            NL_ROUTE_PROTOCOL = 10006
	NT_STATIC_NON_DOD    NL_ROUTE_PROTOCOL = 10007
)

func (protocol NL_ROUTE_PROTOCOL) String() string {
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
		return fmt.Sprintf("NL_ROUTE_PROTOCOL_UNKNOWN(%d)", protocol)
	}
}
