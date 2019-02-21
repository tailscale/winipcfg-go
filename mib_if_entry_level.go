/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// MIB_IF_ENTRY_LEVEL defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex)
type MibIfEntryLevel uint32

const (
	MibIfEntryNormal                  MibIfEntryLevel = 0
	MibIfEntryNormalWithoutStatistics MibIfEntryLevel = 2
)

func (lvl MibIfEntryLevel) String() string {
	switch lvl {
	case MibIfEntryNormal:
		return "MibIfEntryNormal"
	case MibIfEntryNormalWithoutStatistics:
		return "MibIfEntryNormalWithoutStatistics"
	default:
		return fmt.Sprintf("MibIfEntryLevel_UNKNOWN(%d)", lvl)
	}
}
