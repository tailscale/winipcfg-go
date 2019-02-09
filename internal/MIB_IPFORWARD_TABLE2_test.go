/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_MIB_IPFORWARD_TABLE2_Size(t *testing.T) {
	const Actual_MIB_IPFORWARD_TABLE2_Size = unsafe.Sizeof(MIB_IPFORWARD_TABLE2{})

	if Actual_MIB_IPFORWARD_TABLE2_Size != MIB_IPFORWARD_TABLE2_Size {
		t.Errorf("Size of MIB_IPFORWARD_TABLE2 is %d, although %d is expected.", Actual_MIB_IPFORWARD_TABLE2_Size,
			MIB_IPFORWARD_TABLE2_Size)
	}
}

func Test_MIB_IPFORWARD_TABLE2_Offsets(t *testing.T) {
	s := MIB_IPFORWARD_TABLE2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != MIB_IPFORWARD_TABLE2_Table_Offset {
		t.Errorf("MIB_IPFORWARD_TABLE2.Table offset is %d although %d is expected", offset,
			MIB_IPFORWARD_TABLE2_Table_Offset)
		return
	}
}
