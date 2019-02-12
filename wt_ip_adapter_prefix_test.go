/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterPrefixXpSize(t *testing.T) {

	const actualWtIpAdapterPrefixXpSize = unsafe.Sizeof(wtIpAdapterPrefixXp{})

	if actualWtIpAdapterPrefixXpSize != wtIpAdapterPrefixXp_Size {
		t.Errorf("Size of wtIpAdapterPrefixXp is %d, although %d is expected.", actualWtIpAdapterPrefixXpSize,
			wtIpAdapterPrefixXp_Size)
	}
}

func TestWtIpAdapterPrefixXpOffsets(t *testing.T) {

	s := wtIpAdapterPrefixXp{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterPrefixXp_Flags_Offset {
		t.Errorf("wtIpAdapterPrefixXp.Flags offset is %d although %d is expected", offset,
			wtIpAdapterPrefixXp_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterPrefixXp_Next_Offset {
		t.Errorf("wtIpAdapterPrefixXp.Next offset is %d although %d is expected", offset,
			wtIpAdapterPrefixXp_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterPrefixXp_Address_Offset {
		t.Errorf("wtIpAdapterPrefixXp.Address offset is %d although %d is expected", offset,
			wtIpAdapterPrefixXp_Address_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixLength)) - sp

	if offset != wtIpAdapterPrefixXp_PrefixLength_Offset {
		t.Errorf("wtIpAdapterPrefixXp.PrefixLength offset is %d although %d is expected", offset,
			wtIpAdapterPrefixXp_PrefixLength_Offset)
		return
	}
}
