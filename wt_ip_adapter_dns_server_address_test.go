/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWpIpAdapterDnsServerAddressXpSize(t *testing.T) {

	const actualWpIpAdapterDnsServerAddressXp = unsafe.Sizeof(wtIpAdapterDnsServerAddressXp{})

	if actualWpIpAdapterDnsServerAddressXp != wtIpAdapterDnsServerAddressXp_Size {
		t.Errorf("Size of wtIpAdapterDnsServerAddressXp is %d, although %d is expected.",
			actualWpIpAdapterDnsServerAddressXp, wtIpAdapterDnsServerAddressXp_Size)
	}
}

func Test_wtIpAdapterDnsServerAddressXp_Offsets(t *testing.T) {

	s := wtIpAdapterDnsServerAddressXp{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Reserved)) - sp

	if offset != wtIpAdapterDnsServerAddressXp_Reserved_Offset {
		t.Errorf("wtIpAdapterDnsServerAddressXp.Reserved offset is %d although %d is expected", offset,
			wtIpAdapterDnsServerAddressXp_Reserved_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterDnsServerAddressXp_Next_Offset {
		t.Errorf("wtIpAdapterDnsServerAddressXp.Next offset is %d although %d is expected", offset,
			wtIpAdapterDnsServerAddressXp_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterDnsServerAddressXp_Address_Offset {
		t.Errorf("wtIpAdapterDnsServerAddressXp.Address offset is %d although %d is expected", offset,
			wtIpAdapterDnsServerAddressXp_Address_Offset)
		return
	}
}
