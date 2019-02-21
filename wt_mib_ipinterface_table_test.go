/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIpinterfaceTableSize(t *testing.T) {

	const actualWtMibIpinterfaceTableSize = unsafe.Sizeof(wtMibIpinterfaceTable{})

	if actualWtMibIpinterfaceTableSize != wtMibIpinterfaceTable_Size {
		t.Errorf("Size of wtMibIpinterfaceTable is %d, although %d is expected.", actualWtMibIpinterfaceTableSize,
			wtMibIpinterfaceTable_Size)
	}
}

func TestWtMibIpinterfaceTableOffsets(t *testing.T) {

	s := wtMibIpinterfaceTable{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibIpinterfaceTable_Table_Offset {
		t.Errorf("wtMibIpinterfaceTable.Table offset is %d although %d is expected", offset,
			wtMibIpinterfaceTable_Table_Offset)
		return
	}
}
