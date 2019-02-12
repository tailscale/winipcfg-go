/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIn6AddrSize(t *testing.T) {
	const actualWtIn6AddrSize = unsafe.Sizeof(wtIn6Addr{})

	if actualWtIn6AddrSize != wtIn6Addr_Size {
		t.Errorf("Size of wtIn6Addr is %d, although %d is expected.", actualWtIn6AddrSize, wtIn6Addr_Size)
	}
}
