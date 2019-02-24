/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibAnycastipaddressTableSize(t *testing.T) {

	const actualWtMibAnycastipaddressTableSize = unsafe.Sizeof(wtMibAnycastipaddressTable{})

	if actualWtMibAnycastipaddressTableSize != wtMibAnycastipaddressTable_Size {
		t.Errorf("Size of wtMibAnycastipaddressTable is %d, although %d is expected.",
			actualWtMibAnycastipaddressTableSize, wtMibAnycastipaddressTable_Size)
	}
}

func TestWtMibAnycastipaddressTableOffsets(t *testing.T) {

	s := wtMibAnycastipaddressTable{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibAnycastipaddressTable_Table_Offset {
		t.Errorf("wtMibAnycastipaddressTable.Table offset is %d although %d is expected", offset,
			wtMibAnycastipaddressTable_Table_Offset)
		return
	}
}
