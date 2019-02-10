/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_GUID_Size(t *testing.T) {

	const Actual_GUID_Size = unsafe.Sizeof(GUID{})

	if Actual_GUID_Size != GUID_Size {
		t.Errorf("Size of GUID is %d, although %d is expected.", Actual_GUID_Size, GUID_Size)
	}
}

func Test_GUID_Offsets(t *testing.T) {

	s := GUID{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Data2)) - sp

	if offset != GUID_Data2_Offset {
		t.Errorf("GUID.Data2 offset is %d although %d is expected", offset, GUID_Data2_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Data3)) - sp

	if offset != GUID_Data3_Offset {
		t.Errorf("GUID.Data3 offset is %d although %d is expected", offset, GUID_Data3_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Data4)) - sp

	if offset != GUID_Data4_Offset {
		t.Errorf("GUID.Data4 offset is %d although %d is expected", offset, GUID_Data4_Offset)
		return
	}
}
