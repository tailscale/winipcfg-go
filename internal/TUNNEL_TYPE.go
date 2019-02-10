/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-tunnel_type
// Defined in ifdef.h
type TUNNEL_TYPE int

const (
	TUNNEL_TYPE_NONE    TUNNEL_TYPE = 0
	TUNNEL_TYPE_OTHER   TUNNEL_TYPE = 1
	TUNNEL_TYPE_DIRECT  TUNNEL_TYPE = 2
	TUNNEL_TYPE_6TO4    TUNNEL_TYPE = 11
	TUNNEL_TYPE_ISATAP  TUNNEL_TYPE = 13
	TUNNEL_TYPE_TEREDO  TUNNEL_TYPE = 14
	TUNNEL_TYPE_IPHTTPS TUNNEL_TYPE = 15
)

func (t TUNNEL_TYPE) String() string {
	switch t {
	case TUNNEL_TYPE_NONE:
		return "TUNNEL_TYPE_NONE"
	case TUNNEL_TYPE_OTHER:
		return "TUNNEL_TYPE_OTHER"
	case TUNNEL_TYPE_DIRECT:
		return "TUNNEL_TYPE_DIRECT"
	case TUNNEL_TYPE_6TO4:
		return "TUNNEL_TYPE_6TO4"
	case TUNNEL_TYPE_ISATAP:
		return "TUNNEL_TYPE_ISATAP"
	case TUNNEL_TYPE_TEREDO:
		return "TUNNEL_TYPE_TEREDO"
	case TUNNEL_TYPE_IPHTTPS:
		return "TUNNEL_TYPE_IPHTTPS"
	default:
		return fmt.Sprintf("TUNNEL_TYPE_UNKNOWN(%d)", t)
	}
}
