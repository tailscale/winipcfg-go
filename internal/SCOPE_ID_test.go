/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_SCOPE_ID_Size(t *testing.T) {
	const Actual_SCOPE_ID_Size = unsafe.Sizeof(SCOPE_ID{})

	if Actual_SCOPE_ID_Size != SCOPE_ID_Size {
		t.Errorf("Size of SCOPE_ID is %d, although %d is expected.", Actual_SCOPE_ID_Size, SCOPE_ID_Size)
	}
}
