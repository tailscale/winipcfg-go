/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-_nl_link_local_address_behavior
// NL_LINK_LOCAL_ADDRESS_BEHAVIOR defined in nldef.h
type NlLinkLocalAddressBehavior int32

const (
	// Never use link locals.
	LinkLocalAlwaysOff NlLinkLocalAddressBehavior = 0

	// Use link locals only if no other addresses.
	// (default for IPv4).
	// Legacy mapping: IPAutoconfigurationEnabled.
	LinkLocalDelayed   NlLinkLocalAddressBehavior = 1

	// Always use link locals (default for IPv6).
	LinkLocalAlwaysOn  NlLinkLocalAddressBehavior = 2

	LinkLocalUnchanged NlLinkLocalAddressBehavior = -1
)

func (llab NlLinkLocalAddressBehavior) String() string {
	switch llab {
	case LinkLocalAlwaysOff:
		return "LinkLocalAlwaysOff"
	case LinkLocalDelayed:
		return "LinkLocalDelayed"
	case LinkLocalAlwaysOn:
		return "LinkLocalAlwaysOn"
	case LinkLocalUnchanged:
		return "LinkLocalUnchanged"
	default:
		return fmt.Sprintf("NlLinkLocalAddressBehavior_UNKNOWN(%d)", llab)
	}
}
