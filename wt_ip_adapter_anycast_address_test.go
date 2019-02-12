/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtInAdapterAnycastAddressXpSize(t *testing.T) {

	const actualWtInAdapterAnycastAddressXpSize = unsafe.Sizeof(wtIpAdapterAnycastAddressXp{})

	if actualWtInAdapterAnycastAddressXpSize != wtIpAdapterAnycastAddressXp_Size {
		t.Errorf("Size of wtIpAdapterAnycastAddressXp is %d, although %d is expected.",
			actualWtInAdapterAnycastAddressXpSize, wtIpAdapterAnycastAddressXp_Size)
	}
}

func TestWtInAdapterAnycastAddressXpOffsets(t *testing.T) {

	s := wtIpAdapterAnycastAddressXp{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Flags_Offset {
		t.Errorf("wtIpAdapterAnycastAddressXp.Flags offset is %d although %d is expected", offset,
			wtIpAdapterAnycastAddressXp_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Next_Offset {
		t.Errorf("wtIpAdapterAnycastAddressXp.Next offset is %d although %d is expected", offset,
			wtIpAdapterAnycastAddressXp_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Address_Offset {
		t.Errorf("wtIpAdapterAnycastAddressXp.Address offset is %d although %d is expected", offset,
			wtIpAdapterAnycastAddressXp_Address_Offset)
		return
	}
}
