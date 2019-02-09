/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_MIB_IPFORWARD_ROW2_Size(t *testing.T) {
	const Actual_MIB_IPFORWARD_ROW2_Size = unsafe.Sizeof(MIB_IPFORWARD_ROW2{})

	if Actual_MIB_IPFORWARD_ROW2_Size != MIB_IPFORWARD_ROW2_Size {
		t.Errorf("Size of MIB_IPFORWARD_ROW2 is %d, although %d is expected.", Actual_MIB_IPFORWARD_ROW2_Size,
			MIB_IPFORWARD_ROW2_Size)
	}
}

func Test_MIB_IPFORWARD_ROW2_Offsets(t *testing.T) {
	s := MIB_IPFORWARD_ROW2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != MIB_IPFORWARD_ROW2_InterfaceIndex_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.InterfaceIndex offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DestinationPrefix)) - sp

	if offset != MIB_IPFORWARD_ROW2_DestinationPrefix_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.DestinationPrefix offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_DestinationPrefix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NextHop)) - sp

	if offset != MIB_IPFORWARD_ROW2_NextHop_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.NextHop offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_NextHop_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SitePrefixLength)) - sp

	if offset != MIB_IPFORWARD_ROW2_SitePrefixLength_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.SitePrefixLength offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_SitePrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != MIB_IPFORWARD_ROW2_ValidLifetime_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.ValidLifetime offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != MIB_IPFORWARD_ROW2_PreferredLifetime_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.PreferredLifetime offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Metric)) - sp

	if offset != MIB_IPFORWARD_ROW2_Metric_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Metric offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Protocol)) - sp

	if offset != MIB_IPFORWARD_ROW2_Protocol_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Protocol offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Protocol_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Loopback)) - sp

	if offset != MIB_IPFORWARD_ROW2_Loopback_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Loopback offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Loopback_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AutoconfigureAddress)) - sp

	if offset != MIB_IPFORWARD_ROW2_AutoconfigureAddress_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.AutoconfigureAddress offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_AutoconfigureAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Publish)) - sp

	if offset != MIB_IPFORWARD_ROW2_Publish_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Publish offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Publish_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Immortal)) - sp

	if offset != MIB_IPFORWARD_ROW2_Immortal_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Immortal offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Immortal_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Age)) - sp

	if offset != MIB_IPFORWARD_ROW2_Age_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Age offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Age_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Origin)) - sp

	if offset != MIB_IPFORWARD_ROW2_Origin_Offset {
		t.Errorf("MIB_IPFORWARD_ROW2.Origin offset is %d although %d is expected", offset,
			MIB_IPFORWARD_ROW2_Origin_Offset)
		return
	}
}
