/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterDnsSuffix_Size(t *testing.T) {

	const actualWtIpAdapterDnsSuffixSize = unsafe.Sizeof(wtIpAdapterDnsSuffix{})

	if actualWtIpAdapterDnsSuffixSize != wtIpAdapterDnsSuffix_Size {
		t.Errorf("Size of wtIpAdapterDnsSuffix is %d, although %d is expected.", actualWtIpAdapterDnsSuffixSize,
			wtIpAdapterDnsSuffix_Size)
	}
}

func TestWtIpAdapterDnsSuffix_Offsets(t *testing.T) {

	s := wtIpAdapterDnsSuffix{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.String)) - sp

	if offset != wtIpAdapterDnsSuffix_String_Offset {
		t.Errorf("wtIpAdapterDnsSuffix.String offset is %d although %d is expected", offset,
			wtIpAdapterDnsSuffix_String_Offset)
		return
	}
}
