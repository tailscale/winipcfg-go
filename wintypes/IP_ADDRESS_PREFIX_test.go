/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_IP_ADDRESS_PREFIX_Size(t *testing.T) {

	const Actual_IP_ADDRESS_PREFIX_Size = unsafe.Sizeof(IP_ADDRESS_PREFIX{})

	if Actual_IP_ADDRESS_PREFIX_Size != IP_ADDRESS_PREFIX_Size {
		t.Errorf("Size of IP_ADDRESS_PREFIX is %d, although %d is expected.", Actual_IP_ADDRESS_PREFIX_Size, IP_ADDRESS_PREFIX_Size)
	}
}

func Test_IP_ADDRESS_PREFIX_Offsets(t *testing.T) {

	s := IP_ADDRESS_PREFIX{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.PrefixLength)) - sp

	if offset != IP_ADDRESS_PREFIX_PrefixLength_Offset {
		t.Errorf("IP_ADDRESS_PREFIX.PrefixLength offset is %d although %d is expected", offset, IP_ADDRESS_PREFIX_PrefixLength_Offset)
		return
	}
}
