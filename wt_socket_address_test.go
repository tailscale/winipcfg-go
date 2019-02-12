/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWpSocketAddressSize(t *testing.T) {

	const actualWpSocketAddressSize = unsafe.Sizeof(wtSocketAddress{})

	if actualWpSocketAddressSize != wtSocketAddress_Size {
		t.Errorf("Size of wtSocketAddress is %d, although %d is expected.", actualWpSocketAddressSize, wtSocketAddress_Size)
	}
}

func TestWpSocketAddressOffsets(t *testing.T) {

	s := wtSocketAddress{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.iSockaddrLength)) - sp

	if offset != wtSocketAddress_iSockaddrLength_Offset {
		t.Errorf("wtSocketAddress.iSockaddrLength offset is %d although %d is expected", offset, wtSocketAddress_iSockaddrLength_Offset)
		return
	}
}

func Test_SOCKADDR_Size(t *testing.T) {

	const Actual_SOCKADDR_Size = unsafe.Sizeof(SOCKADDR{})

	if Actual_SOCKADDR_Size != wtSockaddr_Size {
		t.Errorf("Size of SOCKADDR is %d, although %d is expected.", Actual_SOCKADDR_Size, wtSockaddr_Size)
	}
}

func Test_SOCKADDR_Offsets(t *testing.T) {

	s := SOCKADDR{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sa_data)) - sp

	if offset != wtSockaddr_sa_data_Offset {
		t.Errorf("SOCKADDR.sa_data offset is %d although %d is expected", offset, wtSockaddr_sa_data_Offset)
		return
	}
}
