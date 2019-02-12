/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtSockaddrInSize(t *testing.T) {

	const actualWtSockaddrInSize = unsafe.Sizeof(wtSockaddrIn{})

	if actualWtSockaddrInSize != wtSockaddrIn_Size {
		t.Errorf("Size of wtSockaddrIn is %d, although %d is expected.", actualWtSockaddrInSize, wtSockaddrIn_Size)
	}
}

func TestWtSockaddrInOffsets(t *testing.T) {

	s := wtSockaddrIn{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sin_port)) - sp

	if offset != wtSockaddrIn_sin_port_Offset {
		t.Errorf("wtSockaddrIn.sin_port offset is %d although %d is expected", offset, wtSockaddrIn_sin_port_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin_addr)) - sp

	if offset != wtSockaddrIn_sin_addr_Offset {
		t.Errorf("wtSockaddrIn.sin_addr offset is %d although %d is expected", offset, wtSockaddrIn_sin_addr_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin_zero)) - sp

	if offset != wtSockaddrIn_sin_zero_Offset {
		t.Errorf("wtSockaddrIn.sin_zero offset is %d although %d is expected", offset, wtSockaddrIn_sin_zero_Offset)
		return
	}
}

func TestWtSockaddrIn6LhSize(t *testing.T) {

	const actualWtSockaddrIn6LhSize = unsafe.Sizeof(wtSockaddrIn6Lh{})

	if actualWtSockaddrIn6LhSize != wtSockaddrIn6Lh_Size {
		t.Errorf("Size of wtSockaddrIn6Lh is %d, although %d is expected.", actualWtSockaddrIn6LhSize,
			wtSockaddrIn6Lh_Size)
	}
}

func TestWtSockaddrIn6LhOffsets(t *testing.T) {

	s := wtSockaddrIn6Lh{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sin6_port)) - sp

	if offset != wtSockaddrIn6Lh_sin6_port_Offset {
		t.Errorf("wtSockaddrIn6Lh.sin6_port offset is %d although %d is expected", offset,
			wtSockaddrIn6Lh_sin6_port_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_flowinfo)) - sp

	if offset != wtSockaddrIn6Lh_sin6_flowinfo_Offset {
		t.Errorf("wtSockaddrIn6Lh.sin6_flowinfo offset is %d although %d is expected", offset,
			wtSockaddrIn6Lh_sin6_flowinfo_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_addr)) - sp

	if offset != wtSockaddrIn6Lh_sin6_addr_Offset {
		t.Errorf("wtSockaddrIn6Lh.sin6_addr offset is %d although %d is expected", offset,
			wtSockaddrIn6Lh_sin6_addr_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_scope_id)) - sp

	if offset != wtSockaddrIn6Lh_sin6_scope_id_Offset {
		t.Errorf("wtSockaddrIn6Lh.sin6_scope_id offset is %d although %d is expected", offset,
			wtSockaddrIn6Lh_sin6_scope_id_Offset)
		return
	}
}
