/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-if_oper_status
// Defined in ifdef.h
type IF_OPER_STATUS uint32

const (
	IfOperStatusUp             IF_OPER_STATUS = 1
	IfOperStatusDown           IF_OPER_STATUS = 2
	IfOperStatusTesting        IF_OPER_STATUS = 3
	IfOperStatusUnknown        IF_OPER_STATUS = 4
	IfOperStatusDormant        IF_OPER_STATUS = 5
	IfOperStatusNotPresent     IF_OPER_STATUS = 6
	IfOperStatusLowerLayerDown IF_OPER_STATUS = 7
)

func (s IF_OPER_STATUS) String() string {
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
		return fmt.Sprintf("IF_OPER_STATUS_UNKNOWN(%d)", s)
	}
}
