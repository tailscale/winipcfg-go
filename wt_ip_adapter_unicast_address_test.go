/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterUnicastAddressLhSize(t *testing.T) {

	const actualWtIpAdapterPrefixXpSize = unsafe.Sizeof(wtIpAdapterUnicastAddressLh{})

	if actualWtIpAdapterPrefixXpSize != wtIpAdapterUnicastAddressLh_Size {
		t.Errorf("Size of wtIpAdapterUnicastAddressLh is %d, although %d is expected.",
			actualWtIpAdapterPrefixXpSize, wtIpAdapterUnicastAddressLh_Size)
	}
}

func Test_wtIpAdapterUnicastAddressLh_Offsets(t *testing.T) {

	s := wtIpAdapterUnicastAddressLh{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterUnicastAddressLh_Flags_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.Flags offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterUnicastAddressLh_Next_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.Next offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterUnicastAddressLh_Address_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.Address offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_Address_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixOrigin)) - sp

	if offset != wtIpAdapterUnicastAddressLh_PrefixOrigin_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.PrefixOrigin offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_PrefixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SuffixOrigin)) - sp

	if offset != wtIpAdapterUnicastAddressLh_SuffixOrigin_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.SuffixOrigin offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_SuffixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadState)) - sp

	if offset != wtIpAdapterUnicastAddressLh_DadState_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.DadState offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_DadState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != wtIpAdapterUnicastAddressLh_ValidLifetime_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.ValidLifetime offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != wtIpAdapterUnicastAddressLh_PreferredLifetime_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.PreferredLifetime offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.LeaseLifetime)) - sp

	if offset != wtIpAdapterUnicastAddressLh_LeaseLifetime_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.LeaseLifetime offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_LeaseLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OnLinkPrefixLength)) - sp

	if offset != wtIpAdapterUnicastAddressLh_OnLinkPrefixLength_Offset {
		t.Errorf("wtIpAdapterUnicastAddressLh.OnLinkPrefixLength offset is %d although %d is expected", offset,
			wtIpAdapterUnicastAddressLh_OnLinkPrefixLength_Offset)
		return
	}
}
