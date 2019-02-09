/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_dad_state
// Defined in nldef.h
type NL_DAD_STATE uint32

const (
	IpDadStateInvalid    NL_DAD_STATE = 0
	IpDadStateTentative  NL_DAD_STATE = 1
	IpDadStateDuplicate  NL_DAD_STATE = 2
	IpDadStateDeprecated NL_DAD_STATE = 3
	IpDadStatePreferred  NL_DAD_STATE = 4
)

func (s NL_DAD_STATE) String() string {
	switch s {
	case IpDadStateInvalid:
		return "IpDadStateInvalid"
	case IpDadStateTentative:
		return "IpDadStateTentative"
	case IpDadStateDuplicate:
		return "IpDadStateDuplicate"
	case IpDadStateDeprecated:
		return "IpDadStateDeprecated"
	case IpDadStatePreferred:
		return "IpDadStatePreferred"
	default:
		return fmt.Sprintf("NL_DAD_STATE_UNKNOWN(%d)", s)
	}
}
