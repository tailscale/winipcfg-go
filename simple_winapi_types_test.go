/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"testing"
	"unsafe"
)

func TestGuidSize(t *testing.T) {

	const actualGuidSize = unsafe.Sizeof(windows.GUID{})

	if actualGuidSize != windowsGuid_Size {
		t.Errorf("Size of GUID is %d, although %d is expected.", actualGuidSize, windowsGuid_Size)
	}
}

func TestGuidOffsets(t *testing.T) {

	s := windows.GUID{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Data2)) - sp

	if offset != windowsGuid_Data2_Offset {
		t.Errorf("GUID.Data2 offset is %d although %d is expected", offset, windowsGuid_Data2_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Data3)) - sp

	if offset != windowsGuid_Data3_Offset {
		t.Errorf("GUID.Data3 offset is %d although %d is expected", offset, windowsGuid_Data3_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Data4)) - sp

	if offset != windowsGuid_Data4_Offset {
		t.Errorf("GUID.Data4 offset is %d although %d is expected", offset, windowsGuid_Data4_Offset)
		return
	}
}
