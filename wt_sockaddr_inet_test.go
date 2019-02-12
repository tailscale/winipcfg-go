/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_SOCKADDR_IN_Size(t *testing.T) {

	const Actual_SOCKADDR_IN_Size = unsafe.Sizeof(SOCKADDR_IN{})

	if Actual_SOCKADDR_IN_Size != wtSockaddrIn_Size {
		t.Errorf("Size of SOCKADDR_IN is %d, although %d is expected.", Actual_SOCKADDR_IN_Size, wtSockaddrIn_Size)
	}
}

func Test_SOCKADDR_IN_Offsets(t *testing.T) {

	s := SOCKADDR_IN{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sin_port)) - sp

	if offset != wtSockaddrIn_sin_port_Offset {
		t.Errorf("SOCKADDR_IN.sin_port offset is %d although %d is expected", offset, wtSockaddrIn_sin_port_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin_addr)) - sp

	if offset != wtSockaddrIn_sin_addr_Offset {
		t.Errorf("SOCKADDR_IN.sin_addr offset is %d although %d is expected", offset, wtSockaddrIn_sin_addr_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin_zero)) - sp

	if offset != wtSockaddrIn_sin_zero_Offset {
		t.Errorf("SOCKADDR_IN.sin_zero offset is %d although %d is expected", offset, wtSockaddrIn_sin_zero_Offset)
		return
	}
}

func Test_SOCKADDR_IN6_LH_Size(t *testing.T) {

	const Actual_SOCKADDR_IN6_LH_Size = unsafe.Sizeof(SOCKADDR_IN6_LH{})

	if Actual_SOCKADDR_IN6_LH_Size != wtSockaddrIn6Lh_Size {
		t.Errorf("Size of SOCKADDR_IN6_LH is %d, although %d is expected.", Actual_SOCKADDR_IN6_LH_Size, wtSockaddrIn6Lh_Size)
	}
}

func Test_SOCKADDR_IN6_LH_Offsets(t *testing.T) {

	s := SOCKADDR_IN6_LH{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sin6_port)) - sp

	if offset != wtSockaddrIn6Lh_sin6_port_Offset {
		t.Errorf("SOCKADDR_IN6_LH.sin6_port offset is %d although %d is expected", offset, wtSockaddrIn6Lh_sin6_port_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_flowinfo)) - sp

	if offset != wtSockaddrIn6Lh_sin6_flowinfo_Offset {
		t.Errorf("SOCKADDR_IN6_LH.sin6_flowinfo offset is %d although %d is expected", offset, wtSockaddrIn6Lh_sin6_flowinfo_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_addr)) - sp

	if offset != wtSockaddrIn6Lh_sin6_addr_Offset {
		t.Errorf("SOCKADDR_IN6_LH.sin6_addr offset is %d although %d is expected", offset, wtSockaddrIn6Lh_sin6_addr_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sin6_scope_id)) - sp

	if offset != wtSockaddrIn6Lh_sin6_scope_id_Offset {
		t.Errorf("SOCKADDR_IN6_LH.sin6_scope_id offset is %d although %d is expected", offset, wtSockaddrIn6Lh_sin6_scope_id_Offset)
		return
	}
}
