/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIpforwardRow2Size(t *testing.T) {

	const actualWtMibIpforwardRow2Size = unsafe.Sizeof(wtMibIpforwardRow2{})

	if actualWtMibIpforwardRow2Size != wtMibIpforwardRow2_Size {
		t.Errorf("Size of wtMibIpforwardRow2 is %d, although %d is expected.", actualWtMibIpforwardRow2Size,
			wtMibIpforwardRow2_Size)
	}
}

func TestWtMibIpforwardRow2Offsets(t *testing.T) {

	s := wtMibIpforwardRow2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibIpforwardRow2_InterfaceIndex_Offset {
		t.Errorf("wtMibIpforwardRow2.InterfaceIndex offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DestinationPrefix)) - sp

	if offset != wtMibIpforwardRow2_DestinationPrefix_Offset {
		t.Errorf("wtMibIpforwardRow2.DestinationPrefix offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_DestinationPrefix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NextHop)) - sp

	if offset != wtMibIpforwardRow2_NextHop_Offset {
		t.Errorf("wtMibIpforwardRow2.NextHop offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_NextHop_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SitePrefixLength)) - sp

	if offset != wtMibIpforwardRow2_SitePrefixLength_Offset {
		t.Errorf("wtMibIpforwardRow2.SitePrefixLength offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_SitePrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ValidLifetime)) - sp

	if offset != wtMibIpforwardRow2_ValidLifetime_Offset {
		t.Errorf("wtMibIpforwardRow2.ValidLifetime offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_ValidLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PreferredLifetime)) - sp

	if offset != wtMibIpforwardRow2_PreferredLifetime_Offset {
		t.Errorf("wtMibIpforwardRow2.PreferredLifetime offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_PreferredLifetime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Metric)) - sp

	if offset != wtMibIpforwardRow2_Metric_Offset {
		t.Errorf("wtMibIpforwardRow2.Metric offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Protocol)) - sp

	if offset != wtMibIpforwardRow2_Protocol_Offset {
		t.Errorf("wtMibIpforwardRow2.Protocol offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Protocol_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Loopback)) - sp

	if offset != wtMibIpforwardRow2_Loopback_Offset {
		t.Errorf("wtMibIpforwardRow2.Loopback offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Loopback_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AutoconfigureAddress)) - sp

	if offset != wtMibIpforwardRow2_AutoconfigureAddress_Offset {
		t.Errorf("wtMibIpforwardRow2.AutoconfigureAddress offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_AutoconfigureAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Publish)) - sp

	if offset != wtMibIpforwardRow2_Publish_Offset {
		t.Errorf("wtMibIpforwardRow2.Publish offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Publish_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Immortal)) - sp

	if offset != wtMibIpforwardRow2_Immortal_Offset {
		t.Errorf("wtMibIpforwardRow2.Immortal offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Immortal_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Age)) - sp

	if offset != wtMibIpforwardRow2_Age_Offset {
		t.Errorf("wtMibIpforwardRow2.Age offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Age_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Origin)) - sp

	if offset != wtMibIpforwardRow2_Origin_Offset {
		t.Errorf("wtMibIpforwardRow2.Origin offset is %d although %d is expected", offset,
			wtMibIpforwardRow2_Origin_Offset)
		return
	}
}
