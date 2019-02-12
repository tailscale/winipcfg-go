/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAddressPrefixSize(t *testing.T) {

	const actualWtIpAddressPrefixSize = unsafe.Sizeof(wtIpAddressPrefix{})

	if actualWtIpAddressPrefixSize != wtIpAddressPrefix_Size {
		t.Errorf("Size of wtIpAddressPrefix is %d, although %d is expected.", actualWtIpAddressPrefixSize,
			wtIpAddressPrefix_Size)
	}
}

func TestWtIpAddressPrefixOffsets(t *testing.T) {

	s := wtIpAddressPrefix{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.PrefixLength)) - sp

	if offset != wtIpAddressPrefix_PrefixLength_Offset {
		t.Errorf("wtIpAddressPrefix.PrefixLength offset is %d although %d is expected", offset,
			wtIpAddressPrefix_PrefixLength_Offset)
		return
	}
}
