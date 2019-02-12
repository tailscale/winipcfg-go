/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_dad_state
// NL_DAD_STATE defined in nldef.h
type NlDadState uint32

const (
	IpDadStateInvalid    NlDadState = 0
	IpDadStateTentative  NlDadState = 1
	IpDadStateDuplicate  NlDadState = 2
	IpDadStateDeprecated NlDadState = 3
	IpDadStatePreferred  NlDadState = 4
)

func (s NlDadState) String() string {
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
		return fmt.Sprintf("NlDadState_UNKNOWN(%d)", s)
	}
}

// IP_DAD_STATE defined in iptypes.h
type IpDadState NlDadState

func (s IpDadState) String() string {
	return NlDadState(s).String()
}
