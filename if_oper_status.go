/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-if_oper_status
// IF_OPER_STATUS defined in ifdef.h
type IfOperStatus uint32

const (
	IfOperStatusUp             IfOperStatus = 1
	IfOperStatusDown           IfOperStatus = 2
	IfOperStatusTesting        IfOperStatus = 3
	IfOperStatusUnknown        IfOperStatus = 4
	IfOperStatusDormant        IfOperStatus = 5
	IfOperStatusNotPresent     IfOperStatus = 6
	IfOperStatusLowerLayerDown IfOperStatus = 7
)

func (s IfOperStatus) String() string {
	switch s {
	case IfOperStatusUp:
		return "IfOperStatusUp"
	case IfOperStatusDown:
		return "IfOperStatusDown"
	case IfOperStatusTesting:
		return "IfOperStatusTesting"
	case IfOperStatusUnknown:
		return "IfOperStatusUnknown"
	case IfOperStatusDormant:
		return "IfOperStatusDormant"
	case IfOperStatusNotPresent:
		return "IfOperStatusNotPresent"
	case IfOperStatusLowerLayerDown:
		return "IfOperStatusLowerLayerDown"
	default:
		return fmt.Sprintf("IfOperStatus_UNKNOWN(%d)", s)
	}
}
