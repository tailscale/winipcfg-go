/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

/*
* Although the original type is union (https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ns-ifdef-_net_luid_lh),
* here I'll create structure instead, with one (the largest) field.
 */
type NET_LUID_LH struct {
	Value ULONG64
}

// Defined in ifdef.h
type NET_LUID NET_LUID_LH
