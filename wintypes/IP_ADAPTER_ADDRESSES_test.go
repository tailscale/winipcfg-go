/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_IP_ADAPTER_ADDRESSES_LH_Size(t *testing.T) {

	const Actual_IP_ADAPTER_ADDRESSES_LH_Size = unsafe.Sizeof(IP_ADAPTER_ADDRESSES_LH{})

	if Actual_IP_ADAPTER_ADDRESSES_LH_Size != IP_ADAPTER_ADDRESSES_LH_Size {
		t.Errorf("Size of IP_ADAPTER_ADDRESSES_LH is %d, although %d is expected.", Actual_IP_ADAPTER_ADDRESSES_LH_Size, IP_ADAPTER_ADDRESSES_LH_Size)
	}
}

func Test_IP_ADAPTER_ADDRESSES_LH_Offsets(t *testing.T) {

	s := IP_ADAPTER_ADDRESSES_LH{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.IfIndex)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_IfIndex_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.IfIndex offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_IfIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Next_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Next offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AdapterName)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_AdapterName_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.AdapterName offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_AdapterName_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstUnicastAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstUnicastAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstUnicastAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstUnicastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstAnycastAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstAnycastAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstAnycastAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstAnycastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstMulticastAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstMulticastAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstMulticastAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstMulticastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstDnsServerAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstDnsServerAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstDnsServerAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstDnsServerAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DnsSuffix)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_DnsSuffix_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.DnsSuffix offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_DnsSuffix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Description)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Description_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Description offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Description_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FriendlyName)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FriendlyName_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FriendlyName offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FriendlyName_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_PhysicalAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.PhysicalAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_PhysicalAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddressLength)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_PhysicalAddressLength_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.PhysicalAddressLength offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_PhysicalAddressLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Flags_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Flags offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Mtu)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Mtu_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Mtu offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Mtu_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.IfType)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_IfType_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.IfType offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_IfType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OperStatus)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_OperStatus_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.OperStatus offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_OperStatus_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv6IfIndex)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Ipv6IfIndex_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Ipv6IfIndex offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Ipv6IfIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ZoneIndices)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_ZoneIndices_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.ZoneIndices offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_ZoneIndices_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstPrefix)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstPrefix_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstPrefix offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstPrefix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TransmitLinkSpeed)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_TransmitLinkSpeed_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.TransmitLinkSpeed offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_TransmitLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ReceiveLinkSpeed)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_ReceiveLinkSpeed_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.ReceiveLinkSpeed offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_ReceiveLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstWinsServerAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstWinsServerAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstWinsServerAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstWinsServerAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstGatewayAddress)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstGatewayAddress_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstGatewayAddress offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstGatewayAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv4Metric)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Ipv4Metric_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Ipv4Metric offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Ipv4Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv6Metric)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Ipv6Metric_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Ipv6Metric offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Ipv6Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Luid)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Luid_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Luid offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Luid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv4Server)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Dhcpv4Server_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Dhcpv4Server offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Dhcpv4Server_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.CompartmentId)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_CompartmentId_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.CompartmentId offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_CompartmentId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NetworkGuid)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_NetworkGuid_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.NetworkGuid offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_NetworkGuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ConnectionType)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_ConnectionType_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.ConnectionType offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_ConnectionType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TunnelType)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_TunnelType_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.TunnelType offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_TunnelType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6Server)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Dhcpv6Server_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Dhcpv6Server offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Dhcpv6Server_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6ClientDuid)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuid_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Dhcpv6ClientDuid offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6ClientDuidLength)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuidLength_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Dhcpv6ClientDuidLength offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Dhcpv6ClientDuidLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6Iaid)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_Dhcpv6Iaid_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.Dhcpv6Iaid offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_Dhcpv6Iaid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstDnsSuffix)) - sp

	if offset != IP_ADAPTER_ADDRESSES_LH_FirstDnsSuffix_Offset {
		t.Errorf("IP_ADAPTER_ADDRESSES_LH.FirstDnsSuffix offset is %d although %d is expected", offset, IP_ADAPTER_ADDRESSES_LH_FirstDnsSuffix_Offset)
		return
	}
}
