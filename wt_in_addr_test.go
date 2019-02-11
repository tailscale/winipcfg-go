/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

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

func Test_IN_ADDR_Offsets(t *testing.T) {

	s := IN_ADDR{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.s_b2)) - sp

	if offset != IN_ADDR_s_b2_Offset {
		t.Errorf("IN_ADDR.s_b2 offset is %d although %d is expected", offset, IN_ADDR_s_b2_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.s_b3)) - sp

	if offset != IN_ADDR_s_b3_Offset {
		t.Errorf("IN_ADDR.s_b3 offset is %d although %d is expected", offset, IN_ADDR_s_b3_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.s_b4)) - sp

	if offset != IN_ADDR_s_b4_Offset {
		t.Errorf("IN_ADDR.s_b4 offset is %d although %d is expected", offset, IN_ADDR_s_b4_Offset)
		return
	}
}
