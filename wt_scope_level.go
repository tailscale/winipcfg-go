/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ne-ws2def-scope_level
// SCOPE_LEVEL defined in ws2def.h
type wtScopeLevel uint32

const (
	ScopeLevelInterface    wtScopeLevel = 1
	ScopeLevelLink         wtScopeLevel = 2
	ScopeLevelSubnet       wtScopeLevel = 3
	ScopeLevelAdmin        wtScopeLevel = 4
	ScopeLevelSite         wtScopeLevel = 5
	ScopeLevelOrganization wtScopeLevel = 8
	ScopeLevelGlobal       wtScopeLevel = 14
	ScopeLevelCount        wtScopeLevel = 16
)

func (sl wtScopeLevel) String() string {
	switch sl {
	case ScopeLevelInterface:
		return "ScopeLevelInterface"
	case ScopeLevelLink:
		return "ScopeLevelLink"
	case ScopeLevelSubnet:
		return "ScopeLevelSubnet"
	case ScopeLevelAdmin:
		return "ScopeLevelAdmin"
	case ScopeLevelSite:
		return "ScopeLevelSite"
	case ScopeLevelOrganization:
		return "ScopeLevelOrganization"
	case ScopeLevelGlobal:
		return "ScopeLevelGlobal"
	case ScopeLevelCount:
		return "ScopeLevelCount"
	default:
		return fmt.Sprintf("wtScopeLevel_UNKNOWN(%d)", sl)
	}
}
