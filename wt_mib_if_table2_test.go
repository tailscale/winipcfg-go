/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIfTable2Size(t *testing.T) {

	const actualWtMibIfTable2Size = unsafe.Sizeof(wtMibIfTable2{})

	if actualWtMibIfTable2Size != wtMibIfTable2_Size {
		t.Errorf("Size of wtMibIfTable2 is %d, although %d is expected.", actualWtMibIfTable2Size,
			wtMibIfTable2_Size)
	}
}

func TestWtMibIfTable2Offsets(t *testing.T) {

	s := wtMibIfTable2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibIfTable2_Table_Offset {
		t.Errorf("wtMibIfTable2.Table offset is %d although %d is expected", offset, wtMibIfTable2_Table_Offset)
		return
	}
}
