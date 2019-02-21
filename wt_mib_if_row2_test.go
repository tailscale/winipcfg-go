/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIfRow2Size(t *testing.T) {

	const actualWtMibIfRow2Size = unsafe.Sizeof(wtMibIfRow2{})

	if actualWtMibIfRow2Size != wtMibIfRow2_Size {
		t.Errorf("Size of wtMibIfRow2 is %d, although %d is expected.", actualWtMibIfRow2Size, wtMibIfRow2_Size)
	}
}

func TestWtMibIfRow2Offsets(t *testing.T) {

	s := wtMibIfRow2{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibIfRow2_InterfaceIndex_Offset {
		t.Errorf("wtMibIfRow2.InterfaceIndex offset is %d although %d is expected", offset,
			wtMibIfRow2_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceGuid)) - sp

	if offset != wtMibIfRow2_InterfaceGuid_Offset {
		t.Errorf("wtMibIfRow2.InterfaceGuid offset is %d although %d is expected", offset,
			wtMibIfRow2_InterfaceGuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Alias)) - sp

	if offset != wtMibIfRow2_Alias_Offset {
		t.Errorf("wtMibIfRow2.Alias offset is %d although %d is expected", offset, wtMibIfRow2_Alias_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Description)) - sp

	if offset != wtMibIfRow2_Description_Offset {
		t.Errorf("wtMibIfRow2.Description offset is %d although %d is expected", offset,
			wtMibIfRow2_Description_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddressLength)) - sp

	if offset != wtMibIfRow2_PhysicalAddressLength_Offset {
		t.Errorf("wtMibIfRow2.PhysicalAddressLength offset is %d although %d is expected", offset,
			wtMibIfRow2_PhysicalAddressLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddress)) - sp

	if offset != wtMibIfRow2_PhysicalAddress_Offset {
		t.Errorf("wtMibIfRow2.PhysicalAddress offset is %d although %d is expected", offset,
			wtMibIfRow2_PhysicalAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PermanentPhysicalAddress)) - sp

	if offset != wtMibIfRow2_PermanentPhysicalAddress_Offset {
		t.Errorf("wtMibIfRow2.PermanentPhysicalAddress offset is %d although %d is expected", offset,
			wtMibIfRow2_PermanentPhysicalAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Mtu)) - sp

	if offset != wtMibIfRow2_Mtu_Offset {
		t.Errorf("wtMibIfRow2.Mtu offset is %d although %d is expected", offset,
			wtMibIfRow2_Mtu_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Type)) - sp

	if offset != wtMibIfRow2_Type_Offset {
		t.Errorf("wtMibIfRow2.Type offset is %d although %d is expected", offset,
			wtMibIfRow2_Type_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TunnelType)) - sp

	if offset != wtMibIfRow2_TunnelType_Offset {
		t.Errorf("wtMibIfRow2.TunnelType offset is %d although %d is expected", offset,
			wtMibIfRow2_TunnelType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.MediaType)) - sp

	if offset != wtMibIfRow2_MediaType_Offset {
		t.Errorf("wtMibIfRow2.MediaType offset is %d although %d is expected", offset,
			wtMibIfRow2_MediaType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalMediumType)) - sp

	if offset != wtMibIfRow2_PhysicalMediumType_Offset {
		t.Errorf("wtMibIfRow2.PhysicalMediumType offset is %d although %d is expected", offset,
			wtMibIfRow2_PhysicalMediumType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AccessType)) - sp

	if offset != wtMibIfRow2_AccessType_Offset {
		t.Errorf("wtMibIfRow2.AccessType offset is %d although %d is expected", offset,
			wtMibIfRow2_AccessType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DirectionType)) - sp

	if offset != wtMibIfRow2_DirectionType_Offset {
		t.Errorf("wtMibIfRow2.DirectionType offset is %d although %d is expected", offset,
			wtMibIfRow2_DirectionType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceAndOperStatusFlags)) - sp

	if offset != wtMibIfRow2_InterfaceAndOperStatusFlags_Offset {
		t.Errorf("wtMibIfRow2.InterfaceAndOperStatusFlags offset is %d although %d is expected", offset,
			wtMibIfRow2_InterfaceAndOperStatusFlags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OperStatus)) - sp

	if offset != wtMibIfRow2_OperStatus_Offset {
		t.Errorf("wtMibIfRow2.OperStatus offset is %d although %d is expected", offset,
			wtMibIfRow2_OperStatus_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AdminStatus)) - sp

	if offset != wtMibIfRow2_AdminStatus_Offset {
		t.Errorf("wtMibIfRow2.AdminStatus offset is %d although %d is expected", offset,
			wtMibIfRow2_AdminStatus_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.MediaConnectState)) - sp

	if offset != wtMibIfRow2_MediaConnectState_Offset {
		t.Errorf("wtMibIfRow2.MediaConnectState offset is %d although %d is expected", offset,
			wtMibIfRow2_MediaConnectState_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NetworkGuid)) - sp

	if offset != wtMibIfRow2_NetworkGuid_Offset {
		t.Errorf("wtMibIfRow2.NetworkGuid offset is %d although %d is expected", offset,
			wtMibIfRow2_NetworkGuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ConnectionType)) - sp

	if offset != wtMibIfRow2_ConnectionType_Offset {
		t.Errorf("wtMibIfRow2.ConnectionType offset is %d although %d is expected", offset,
			wtMibIfRow2_ConnectionType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TransmitLinkSpeed)) - sp

	if offset != wtMibIfRow2_TransmitLinkSpeed_Offset {
		t.Errorf("wtMibIfRow2.TransmitLinkSpeed offset is %d although %d is expected", offset,
			wtMibIfRow2_TransmitLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ReceiveLinkSpeed)) - sp

	if offset != wtMibIfRow2_ReceiveLinkSpeed_Offset {
		t.Errorf("wtMibIfRow2.ReceiveLinkSpeed offset is %d although %d is expected", offset,
			wtMibIfRow2_ReceiveLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InOctets)) - sp

	if offset != wtMibIfRow2_InOctets_Offset {
		t.Errorf("wtMibIfRow2.InOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_InOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InUcastPkts)) - sp

	if offset != wtMibIfRow2_InUcastPkts_Offset {
		t.Errorf("wtMibIfRow2.InUcastPkts offset is %d although %d is expected", offset,
			wtMibIfRow2_InUcastPkts_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InNUcastPkts)) - sp

	if offset != wtMibIfRow2_InNUcastPkts_Offset {
		t.Errorf("wtMibIfRow2.InNUcastPkts offset is %d although %d is expected", offset,
			wtMibIfRow2_InNUcastPkts_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InDiscards)) - sp

	if offset != wtMibIfRow2_InDiscards_Offset {
		t.Errorf("wtMibIfRow2.InDiscards offset is %d although %d is expected", offset,
			wtMibIfRow2_InDiscards_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InErrors)) - sp

	if offset != wtMibIfRow2_InErrors_Offset {
		t.Errorf("wtMibIfRow2.InErrors offset is %d although %d is expected", offset,
			wtMibIfRow2_InErrors_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InUnknownProtos)) - sp

	if offset != wtMibIfRow2_InUnknownProtos_Offset {
		t.Errorf("wtMibIfRow2.InUnknownProtos offset is %d although %d is expected", offset,
			wtMibIfRow2_InUnknownProtos_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InUcastOctets)) - sp

	if offset != wtMibIfRow2_InUcastOctets_Offset {
		t.Errorf("wtMibIfRow2.InUcastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_InUcastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InMulticastOctets)) - sp

	if offset != wtMibIfRow2_InMulticastOctets_Offset {
		t.Errorf("wtMibIfRow2.InMulticastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_InMulticastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InBroadcastOctets)) - sp

	if offset != wtMibIfRow2_InBroadcastOctets_Offset {
		t.Errorf("wtMibIfRow2.InBroadcastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_InBroadcastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutOctets)) - sp

	if offset != wtMibIfRow2_OutOctets_Offset {
		t.Errorf("wtMibIfRow2.OutOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_OutOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutUcastPkts)) - sp

	if offset != wtMibIfRow2_OutUcastPkts_Offset {
		t.Errorf("wtMibIfRow2.OutUcastPkts offset is %d although %d is expected", offset,
			wtMibIfRow2_OutUcastPkts_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutNUcastPkts)) - sp

	if offset != wtMibIfRow2_OutNUcastPkts_Offset {
		t.Errorf("wtMibIfRow2.OutNUcastPkts offset is %d although %d is expected", offset,
			wtMibIfRow2_OutNUcastPkts_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutDiscards)) - sp

	if offset != wtMibIfRow2_OutDiscards_Offset {
		t.Errorf("wtMibIfRow2.OutDiscards offset is %d although %d is expected", offset,
			wtMibIfRow2_OutDiscards_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutErrors)) - sp

	if offset != wtMibIfRow2_OutErrors_Offset {
		t.Errorf("wtMibIfRow2.OutErrors offset is %d although %d is expected", offset,
			wtMibIfRow2_OutErrors_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutUcastOctets)) - sp

	if offset != wtMibIfRow2_OutUcastOctets_Offset {
		t.Errorf("wtMibIfRow2.OutUcastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_OutUcastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutMulticastOctets)) - sp

	if offset != wtMibIfRow2_OutMulticastOctets_Offset {
		t.Errorf("wtMibIfRow2.OutMulticastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_OutMulticastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutBroadcastOctets)) - sp

	if offset != wtMibIfRow2_OutBroadcastOctets_Offset {
		t.Errorf("wtMibIfRow2.OutBroadcastOctets offset is %d although %d is expected", offset,
			wtMibIfRow2_OutBroadcastOctets_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OutQLen)) - sp

	if offset != wtMibIfRow2_OutQLen_Offset {
		t.Errorf("wtMibIfRow2.OutQLen offset is %d although %d is expected", offset, wtMibIfRow2_OutQLen_Offset)
		return
	}
}
