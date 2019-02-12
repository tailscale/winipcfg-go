/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtInAddrSize(t *testing.T) {

	const actualWtInAddrSize = unsafe.Sizeof(WtInAddr{})

	if actualWtInAddrSize != wtInAddr_Size {
		t.Errorf("Size of WtInAddr is %d, although %d is expected.", actualWtInAddrSize, wtInAddr_Size)
	}
}

func TestWtInAddrOffsets(t *testing.T) {

	s := WtInAddr{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.s_b2)) - sp

	if offset != wtInAddr_s_b2_Offset {
		t.Errorf("WtInAddr.s_b2 offset is %d although %d is expected", offset, wtInAddr_s_b2_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.s_b3)) - sp

	if offset != wtInAddr_s_b3_Offset {
		t.Errorf("WtInAddr.s_b3 offset is %d although %d is expected", offset, wtInAddr_s_b3_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.s_b4)) - sp

	if offset != wtInAddr_s_b4_Offset {
		t.Errorf("WtInAddr.s_b4 offset is %d although %d is expected", offset, wtInAddr_s_b4_Offset)
		return
	}
}
