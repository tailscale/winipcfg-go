/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_prefix_origin
// Defined in nldef.h
type NL_PREFIX_ORIGIN uint32

const (
	IpPrefixOriginOther               NL_PREFIX_ORIGIN = 0
	IpPrefixOriginManual              NL_PREFIX_ORIGIN = 1
	IpPrefixOriginWellKnown           NL_PREFIX_ORIGIN = 2
	IpPrefixOriginDhcp                NL_PREFIX_ORIGIN = 3
	IpPrefixOriginRouterAdvertisement NL_PREFIX_ORIGIN = 4
	IpPrefixOriginUnchanged           NL_PREFIX_ORIGIN = 1 << 4
)

func (o NL_PREFIX_ORIGIN) String() string {
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
		return fmt.Sprintf("NL_PREFIX_ORIGIN_UNKNOWN(%d)", o)
	}
}
