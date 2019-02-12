/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterWinsServerAddressLhSize(t *testing.T) {

	const actualWtIpAdapterWinsServerAddressLhSize = unsafe.Sizeof(wtIpAdapterWinsServerAddressLh{})

	if actualWtIpAdapterWinsServerAddressLhSize != wtIpAdapterWinsServerAddressLh_Size {
		t.Errorf("Size of wtIpAdapterWinsServerAddressLh is %d, although %d is expected.",
			actualWtIpAdapterWinsServerAddressLhSize, wtIpAdapterWinsServerAddressLh_Size)
	}
}

func TestWtIpAdapterWinsServerAddressLhOffsets(t *testing.T) {

	s := wtIpAdapterWinsServerAddressLh{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Reserved)) - sp

	if offset != wtIpAdapterWinsServerAddressLh_Reserved_Offset {
		t.Errorf("wtIpAdapterWinsServerAddressLh.Reserved offset is %d although %d is expected", offset,
			wtIpAdapterWinsServerAddressLh_Reserved_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterWinsServerAddressLh_Next_Offset {
		t.Errorf("wtIpAdapterWinsServerAddressLh.Next offset is %d although %d is expected", offset,
			wtIpAdapterWinsServerAddressLh_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterWinsServerAddressLh_Address_Offset {
		t.Errorf("wtIpAdapterWinsServerAddressLh.Address offset is %d although %d is expected", offset,
			wtIpAdapterWinsServerAddressLh_Address_Offset)
		return
	}
}
