/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_MIB_UNICASTIPADDRESS_ROW_Size(t *testing.T) {
	const Actual_MIB_UNICASTIPADDRESS_ROW_Size = unsafe.Sizeof(MIB_UNICASTIPADDRESS_ROW{})

	if Actual_MIB_UNICASTIPADDRESS_ROW_Size != MIB_UNICASTIPADDRESS_ROW_Size {
		t.Errorf("Size of MIB_UNICASTIPADDRESS_ROW is %d, although %d is expected.",
			Actual_MIB_UNICASTIPADDRESS_ROW_Size, MIB_UNICASTIPADDRESS_ROW_Size)
	}
}

func Test_MIB_UNICASTIPADDRESS_ROW_Offsets(t *testing.T) {
	s := MIB_UNICASTIPADDRESS_ROW{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceLuid)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_InterfaceLuid_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.InterfaceLuid offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_InterfaceLuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_InterfaceIndex_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.InterfaceIndex offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixOrigin)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_PrefixOrigin_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.PrefixOrigin offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_PrefixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SuffixOrigin)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_SuffixOrigin_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.SuffixOrigin offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_SuffixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_ValidLifetime_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.ValidLifetime offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_PreferredLifetime_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.PreferredLifetime offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OnLinkPrefixLength)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_OnLinkPrefixLength_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.OnLinkPrefixLength offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_OnLinkPrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SkipAsSource)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_SkipAsSource_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.SkipAsSource offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_SkipAsSource_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadState)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_DadState_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.DadState offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_DadState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ScopeId)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_ScopeId_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.ScopeId offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_ScopeId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.CreationTimeStamp)) - sp

	if offset != MIB_UNICASTIPADDRESS_ROW_CreationTimeStamp_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.CreationTimeStamp offset is %d although %d is expected", offset,
			MIB_UNICASTIPADDRESS_ROW_CreationTimeStamp_Offset)
		return
	}
}
