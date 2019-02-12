/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_prefix_origin
// NL_PREFIX_ORIGIN defined in nldef.h
type NlPrefixOrigin uint32

const (
	IpPrefixOriginOther               NlPrefixOrigin = 0
	IpPrefixOriginManual              NlPrefixOrigin = 1
	IpPrefixOriginWellKnown           NlPrefixOrigin = 2
	IpPrefixOriginDhcp                NlPrefixOrigin = 3
	IpPrefixOriginRouterAdvertisement NlPrefixOrigin = 4
	IpPrefixOriginUnchanged           NlPrefixOrigin = 1 << 4
)

func (o NlPrefixOrigin) String() string {
	switch o {
	case IpPrefixOriginOther:
		return "IpPrefixOriginOther"
	case IpPrefixOriginManual:
		return "IpPrefixOriginManual"
	case IpPrefixOriginWellKnown:
		return "IpPrefixOriginWellKnown"
	case IpPrefixOriginDhcp:
		return "IpPrefixOriginDhcp"
	case IpPrefixOriginRouterAdvertisement:
		return "IpPrefixOriginRouterAdvertisement"
	case IpPrefixOriginUnchanged:
		return "IpPrefixOriginUnchanged"
	default:
		return fmt.Sprintf("NlPrefixOrigin_UNKNOWN(%d)", o)
	}
}

// IP_PREFIX_ORIGIN defined in iptypes.h
type IpPrefixOrigin NlPrefixOrigin

func (o IpPrefixOrigin) String() string {
	return NlPrefixOrigin(o).String()
}
