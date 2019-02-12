/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-tunnel_type
// TUNNEL_TYPE defined in ifdef.h
type TunnelType uint32

const (
	TUNNEL_TYPE_NONE    TunnelType = 0
	TUNNEL_TYPE_OTHER   TunnelType = 1
	TUNNEL_TYPE_DIRECT  TunnelType = 2
	TUNNEL_TYPE_6TO4    TunnelType = 11
	TUNNEL_TYPE_ISATAP  TunnelType = 13
	TUNNEL_TYPE_TEREDO  TunnelType = 14
	TUNNEL_TYPE_IPHTTPS TunnelType = 15
)

func (t TunnelType) String() string {
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
		return fmt.Sprintf("TunnelType_UNKNOWN(%d)", t)
	}
}
