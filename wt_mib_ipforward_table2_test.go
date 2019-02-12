/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIpforwardTable2Size(t *testing.T) {

	const actualWtMibIpforwardTable2Size = unsafe.Sizeof(wtMibIpforwardTable2{})

	if actualWtMibIpforwardTable2Size != wtMibIpforwardTable2_Size {
		t.Errorf("Size of wtMibIpforwardTable2 is %d, although %d is expected.", actualWtMibIpforwardTable2Size,
			wtMibIpforwardTable2_Size)
	}
}

func TestWtMibIpforwardTable2Offsets(t *testing.T) {

	s := wtMibIpforwardTable2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibIpforwardTable2_Table_Offset {
		t.Errorf("wtMibIpforwardTable2.Table offset is %d although %d is expected", offset,
			wtMibIpforwardTable2_Table_Offset)
		return
	}
}
