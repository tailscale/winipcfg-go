/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_MIB_UNICASTIPADDRESS_ROW_Size(t *testing.T) {

	const Actual_MIB_UNICASTIPADDRESS_ROW_Size = unsafe.Sizeof(MIB_UNICASTIPADDRESS_ROW{})

	if Actual_MIB_UNICASTIPADDRESS_ROW_Size != wtMibUnicastipaddressRow_Size {
		t.Errorf("Size of MIB_UNICASTIPADDRESS_ROW is %d, although %d is expected.", Actual_MIB_UNICASTIPADDRESS_ROW_Size, wtMibUnicastipaddressRow_Size)
	}
}

func Test_MIB_UNICASTIPADDRESS_ROW_Offsets(t *testing.T) {

	s := MIB_UNICASTIPADDRESS_ROW{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceLuid)) - sp

	if offset != wtMibUnicastipaddressRow_InterfaceLuid_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.InterfaceLuid offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_InterfaceLuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibUnicastipaddressRow_InterfaceIndex_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.InterfaceIndex offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PrefixOrigin)) - sp

	if offset != wtMibUnicastipaddressRow_PrefixOrigin_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.PrefixOrigin offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_PrefixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SuffixOrigin)) - sp

	if offset != wtMibUnicastipaddressRow_SuffixOrigin_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.SuffixOrigin offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_SuffixOrigin_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != wtMibUnicastipaddressRow_ValidLifetime_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.ValidLifetime offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != wtMibUnicastipaddressRow_PreferredLifetime_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.PreferredLifetime offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OnLinkPrefixLength)) - sp

	if offset != wtMibUnicastipaddressRow_OnLinkPrefixLength_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.OnLinkPrefixLength offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_OnLinkPrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SkipAsSource)) - sp

	if offset != wtMibUnicastipaddressRow_SkipAsSource_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.SkipAsSource offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_SkipAsSource_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadState)) - sp

	if offset != wtMibUnicastipaddressRow_DadState_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.DadState offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_DadState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ScopeId)) - sp

	if offset != wtMibUnicastipaddressRow_ScopeId_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.ScopeId offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_ScopeId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.CreationTimeStamp)) - sp

	if offset != wtMibUnicastipaddressRow_CreationTimeStamp_Offset {
		t.Errorf("MIB_UNICASTIPADDRESS_ROW.CreationTimeStamp offset is %d although %d is expected", offset, wtMibUnicastipaddressRow_CreationTimeStamp_Offset)
		return
	}
}
