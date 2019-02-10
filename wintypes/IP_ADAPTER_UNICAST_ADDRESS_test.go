/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_IP_ADAPTER_UNICAST_ADDRESS_LH_Size(t *testing.T) {

	const Actual_IP_ADAPTER_UNICAST_ADDRESS_LH_Size = unsafe.Sizeof(IP_ADAPTER_UNICAST_ADDRESS_LH{})

	if Actual_IP_ADAPTER_UNICAST_ADDRESS_LH_Size != IP_ADAPTER_UNICAST_ADDRESS_LH_Size {
		t.Errorf("Size of IP_ADAPTER_UNICAST_ADDRESS_LH is %d, although %d is expected.", Actual_IP_ADAPTER_UNICAST_ADDRESS_LH_Size, IP_ADAPTER_UNICAST_ADDRESS_LH_Size)
	}
}

func Test_IP_ADAPTER_UNICAST_ADDRESS_LH_Offsets(t *testing.T) {

	s := IP_ADAPTER_UNICAST_ADDRESS_LH{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_Flags_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.Flags offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_Next_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.Next offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_Address_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.Address offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_Address_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixOrigin)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_PrefixOrigin_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.PrefixOrigin offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_PrefixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SuffixOrigin)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_SuffixOrigin_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.SuffixOrigin offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_SuffixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadState)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_DadState_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.DadState offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_DadState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_ValidLifetime_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.ValidLifetime offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_PreferredLifetime_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.PreferredLifetime offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.LeaseLifetime)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_LeaseLifetime_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.LeaseLifetime offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_LeaseLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OnLinkPrefixLength)) - sp

	if offset != IP_ADAPTER_UNICAST_ADDRESS_LH_OnLinkPrefixLength_Offset {
		t.Errorf("IP_ADAPTER_UNICAST_ADDRESS_LH.OnLinkPrefixLength offset is %d although %d is expected", offset, IP_ADAPTER_UNICAST_ADDRESS_LH_OnLinkPrefixLength_Offset)
		return
	}
}
