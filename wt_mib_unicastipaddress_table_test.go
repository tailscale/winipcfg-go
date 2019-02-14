/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibUnicastipaddressTableSize(t *testing.T) {

	const actualWtMibUnicastipaddressTableSize = unsafe.Sizeof(wtMibUnicastipaddressTable{})

	if actualWtMibUnicastipaddressTableSize != wtMibUnicastipaddressTable_Size {
		t.Errorf("Size of wtMibUnicastipaddressTable is %d, although %d is expected.",
			actualWtMibUnicastipaddressTableSize, wtMibUnicastipaddressTable_Size)
	}
}

func TestWtMibUnicastipaddressTableOffsets(t *testing.T) {

	s := wtMibUnicastipaddressTable{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibUnicastipaddressTable_Table_Offset {
		t.Errorf("wtMibUnicastipaddressTable.Table offset is %d although %d is expected", offset,
			wtMibUnicastipaddressTable_Table_Offset)
		return
	}
}
