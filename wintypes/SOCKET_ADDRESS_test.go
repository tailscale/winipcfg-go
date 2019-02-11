/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_SOCKET_ADDRESS_Size(t *testing.T) {

	const Actual_SOCKET_ADDRESS_Size = unsafe.Sizeof(SOCKET_ADDRESS{})

	if Actual_SOCKET_ADDRESS_Size != SOCKET_ADDRESS_Size {
		t.Errorf("Size of SOCKET_ADDRESS is %d, although %d is expected.", Actual_SOCKET_ADDRESS_Size, SOCKET_ADDRESS_Size)
	}
}

func Test_SOCKET_ADDRESS_Offsets(t *testing.T) {

	s := SOCKET_ADDRESS{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.iSockaddrLength)) - sp

	if offset != SOCKET_ADDRESS_iSockaddrLength_Offset {
		t.Errorf("SOCKET_ADDRESS.iSockaddrLength offset is %d although %d is expected", offset, SOCKET_ADDRESS_iSockaddrLength_Offset)
		return
	}
}

func Test_SOCKADDR_Size(t *testing.T) {

	const Actual_SOCKADDR_Size = unsafe.Sizeof(SOCKADDR{})

	if Actual_SOCKADDR_Size != SOCKADDR_Size {
		t.Errorf("Size of SOCKADDR is %d, although %d is expected.", Actual_SOCKADDR_Size, SOCKADDR_Size)
	}
}

func Test_SOCKADDR_Offsets(t *testing.T) {

	s := SOCKADDR{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sa_data)) - sp

	if offset != SOCKADDR_sa_data_Offset {
		t.Errorf("SOCKADDR.sa_data offset is %d although %d is expected", offset, SOCKADDR_sa_data_Offset)
		return
	}
}
