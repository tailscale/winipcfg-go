/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

// Defined in ws2def.h. It's a union there...
type SCOPE_ID struct {
	Value ULONG
}
