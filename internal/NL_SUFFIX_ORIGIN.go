/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_suffix_origin
// Defined in nldef.h
type NL_SUFFIX_ORIGIN uint32

const (
	IpSuffixOriginOther            NL_SUFFIX_ORIGIN = 0
	IpSuffixOriginManual           NL_SUFFIX_ORIGIN = 1
	IpSuffixOriginWellKnown        NL_SUFFIX_ORIGIN = 2
	IpSuffixOriginDhcp             NL_SUFFIX_ORIGIN = 3
	IpSuffixOriginLinkLayerAddress NL_SUFFIX_ORIGIN = 4
	IpSuffixOriginRandom           NL_SUFFIX_ORIGIN = 5
	IpSuffixOriginUnchanged        NL_SUFFIX_ORIGIN = 1 << 4
)

func (o NL_SUFFIX_ORIGIN) String() string {
	switch o {
	case IpSuffixOriginOther:
		return "IpSuffixOriginOther"
	case IpSuffixOriginManual:
		return "IpSuffixOriginManual"
	case IpSuffixOriginWellKnown:
		return "IpSuffixOriginWellKnown"
	case IpSuffixOriginDhcp:
		return "IpSuffixOriginDhcp"
	case IpSuffixOriginLinkLayerAddress:
		return "IpSuffixOriginLinkLayerAddress"
	case IpSuffixOriginRandom:
		return "IpSuffixOriginRandom"
	case IpSuffixOriginUnchanged:
		return "IpSuffixOriginUnchanged"
	default:
		return fmt.Sprintf("NL_SUFFIX_ORIGIN_UNKNOWN(%d)", o)
	}
}
