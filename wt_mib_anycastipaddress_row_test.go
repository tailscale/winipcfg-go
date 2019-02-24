/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibAnycastipaddressRowSize(t *testing.T) {

	const actualWtMibAnycastipaddressRowSize = unsafe.Sizeof(wtMibAnycastipaddressRow{})

	if actualWtMibAnycastipaddressRowSize != wtMibAnycastipaddressRow_Size {
		t.Errorf("Size of wtMibAnycastipaddressRow is %d, although %d is expected.",
			actualWtMibAnycastipaddressRowSize, wtMibAnycastipaddressRow_Size)
	}
}

func TestWtMibAnycastipaddressRowOffsets(t *testing.T) {

	s := wtMibAnycastipaddressRow{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceLuid)) - sp

	if offset != wtMibAnycastipaddressRow_InterfaceLuid_Offset {
		t.Errorf("wtMibAnycastipaddressRow.InterfaceLuid offset is %d although %d is expected", offset,
			wtMibAnycastipaddressRow_InterfaceLuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibAnycastipaddressRow_InterfaceIndex_Offset {
		t.Errorf("wtMibAnycastipaddressRow.InterfaceIndex offset is %d although %d is expected", offset,
			wtMibAnycastipaddressRow_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ScopeId)) - sp

	if offset != wtMibAnycastipaddressRow_ScopeId_Offset {
		t.Errorf("wtMibAnycastipaddressRow.ScopeId offset is %d although %d is expected", offset,
			wtMibAnycastipaddressRow_ScopeId_Offset)
		return
	}
}
