/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_suffix_origin
// NL_SUFFIX_ORIGIN defined in nldef.h
type NlSuffixOrigin uint32

const (
	IpSuffixOriginOther            NlSuffixOrigin = 0
	IpSuffixOriginManual           NlSuffixOrigin = 1
	IpSuffixOriginWellKnown        NlSuffixOrigin = 2
	IpSuffixOriginDhcp             NlSuffixOrigin = 3
	IpSuffixOriginLinkLayerAddress NlSuffixOrigin = 4
	IpSuffixOriginRandom           NlSuffixOrigin = 5
	IpSuffixOriginUnchanged        NlSuffixOrigin = 1 << 4
)

func (o NlSuffixOrigin) String() string {
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
		return fmt.Sprintf("NlSuffixOrigin_UNKNOWN(%d)", o)
	}
}

// IP_SUFFIX_ORIGIN defined in iptypes.h
type IpSuffixOrigin NlSuffixOrigin

func (o IpSuffixOrigin) String() string {
	return NlSuffixOrigin(o).String()
}
