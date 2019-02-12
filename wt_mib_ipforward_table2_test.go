/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_MIB_IPFORWARD_TABLE2_Size(t *testing.T) {

	const Actual_MIB_IPFORWARD_TABLE2_Size = unsafe.Sizeof(MIB_IPFORWARD_TABLE2{})

	if Actual_MIB_IPFORWARD_TABLE2_Size != wtMibIpforwardTable2_Size {
		t.Errorf("Size of MIB_IPFORWARD_TABLE2 is %d, although %d is expected.", Actual_MIB_IPFORWARD_TABLE2_Size, wtMibIpforwardTable2_Size)
	}
}

func Test_MIB_IPFORWARD_TABLE2_Offsets(t *testing.T) {

	s := MIB_IPFORWARD_TABLE2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Table)) - sp

	if offset != wtMibIpforwardTable2_Table_Offset {
		t.Errorf("MIB_IPFORWARD_TABLE2.Table offset is %d although %d is expected", offset, wtMibIpforwardTable2_Table_Offset)
		return
	}
}
