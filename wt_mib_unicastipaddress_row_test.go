/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibUnicastipaddressRowSize(t *testing.T) {

	const actualWtMibUnicastipaddressRowSize = unsafe.Sizeof(wtMibUnicastipaddressRow{})

	if actualWtMibUnicastipaddressRowSize != wtMibUnicastipaddressRow_Size {
		t.Errorf("Size of wtMibUnicastipaddressRow is %d, although %d is expected.",
			actualWtMibUnicastipaddressRowSize, wtMibUnicastipaddressRow_Size)
	}
}

func TestWtMibUnicastipaddressRowOffsets(t *testing.T) {

	s := wtMibUnicastipaddressRow{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceLuid)) - sp

	if offset != wtMibUnicastipaddressRow_InterfaceLuid_Offset {
		t.Errorf("wtMibUnicastipaddressRow.InterfaceLuid offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_InterfaceLuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibUnicastipaddressRow_InterfaceIndex_Offset {
		t.Errorf("wtMibUnicastipaddressRow.InterfaceIndex offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixOrigin)) - sp

	if offset != wtMibUnicastipaddressRow_PrefixOrigin_Offset {
		t.Errorf("wtMibUnicastipaddressRow.PrefixOrigin offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_PrefixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SuffixOrigin)) - sp

	if offset != wtMibUnicastipaddressRow_SuffixOrigin_Offset {
		t.Errorf("wtMibUnicastipaddressRow.SuffixOrigin offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_SuffixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != wtMibUnicastipaddressRow_ValidLifetime_Offset {
		t.Errorf("wtMibUnicastipaddressRow.ValidLifetime offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != wtMibUnicastipaddressRow_PreferredLifetime_Offset {
		t.Errorf("wtMibUnicastipaddressRow.PreferredLifetime offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OnLinkPrefixLength)) - sp

	if offset != wtMibUnicastipaddressRow_OnLinkPrefixLength_Offset {
		t.Errorf("wtMibUnicastipaddressRow.OnLinkPrefixLength offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_OnLinkPrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SkipAsSource)) - sp

	if offset != wtMibUnicastipaddressRow_SkipAsSource_Offset {
		t.Errorf("wtMibUnicastipaddressRow.SkipAsSource offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_SkipAsSource_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadState)) - sp

	if offset != wtMibUnicastipaddressRow_DadState_Offset {
		t.Errorf("wtMibUnicastipaddressRow.DadState offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_DadState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ScopeId)) - sp

	if offset != wtMibUnicastipaddressRow_ScopeId_Offset {
		t.Errorf("wtMibUnicastipaddressRow.ScopeId offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_ScopeId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.CreationTimeStamp)) - sp

	if offset != wtMibUnicastipaddressRow_CreationTimeStamp_Offset {
		t.Errorf("wtMibUnicastipaddressRow.CreationTimeStamp offset is %d although %d is expected", offset,
			wtMibUnicastipaddressRow_CreationTimeStamp_Offset)
		return
	}
}
