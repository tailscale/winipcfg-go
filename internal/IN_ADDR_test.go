/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_IN_ADDR_Size(t *testing.T) {
	const Actual_IN_ADDR_Size = unsafe.Sizeof(IN_ADDR{})

	if Actual_IN_ADDR_Size != IN_ADDR_Size {
		t.Errorf("Size of IN_ADDR is %d, although %d is expected.", Actual_IN_ADDR_Size, IN_ADDR_Size)
	}
}
