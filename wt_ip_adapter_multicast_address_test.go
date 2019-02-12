/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterMulticastAddressXpSize(t *testing.T) {

	const actualWtIpAdapterMulticastAddressXpSize = unsafe.Sizeof(wtIpAdapterMulticastAddressXp{})

	if actualWtIpAdapterMulticastAddressXpSize != wtIpAdapterMulticastAddressXp_Size {
		t.Errorf("Size of wtIpAdapterMulticastAddressXp is %d, although %d is expected.",
			actualWtIpAdapterMulticastAddressXpSize, wtIpAdapterMulticastAddressXp_Size)
	}
}

func TestWtIpAdapterMulticastAddressXpOffsets(t *testing.T) {

	s := wtIpAdapterMulticastAddressXp{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterMulticastAddressXp_Flags_Offset {
		t.Errorf("wtIpAdapterMulticastAddressXp.Flags offset is %d although %d is expected", offset,
			wtIpAdapterMulticastAddressXp_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterMulticastAddressXp_Next_Offset {
		t.Errorf("wtIpAdapterMulticastAddressXp.Next offset is %d although %d is expected", offset,
			wtIpAdapterMulticastAddressXp_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterMulticastAddressXp_Address_Offset {
		t.Errorf("wtIpAdapterMulticastAddressXp.Address offset is %d although %d is expected", offset,
			wtIpAdapterMulticastAddressXp_Address_Offset)
		return
	}
}
