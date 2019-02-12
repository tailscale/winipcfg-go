/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtInAdapterAddressesLhSize(t *testing.T) {

	const actualWtIn6AddrSize = unsafe.Sizeof(wtIpAdapterAddressesLh{})

	if actualWtIn6AddrSize != wtIpAdapterAddressesLh_Size {
		t.Errorf("Size of wtIpAdapterAddressesLh is %d, although %d is expected.", actualWtIn6AddrSize,
			wtIpAdapterAddressesLh_Size)
	}
}

func TestWtInAdapterAddressesLhOffsets(t *testing.T) {

	s := wtIpAdapterAddressesLh{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.IfIndex)) - sp

	if offset != wtIpAdapterAddressesLh_IfIndex_Offset {
		t.Errorf("wtIpAdapterAddressesLh.IfIndex offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_IfIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterAddressesLh_Next_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Next offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AdapterName)) - sp

	if offset != wtIpAdapterAddressesLh_AdapterName_Offset {
		t.Errorf("wtIpAdapterAddressesLh.AdapterName offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_AdapterName_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstUnicastAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstUnicastAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstUnicastAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstUnicastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstAnycastAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstAnycastAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstAnycastAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstAnycastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstMulticastAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstMulticastAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstMulticastAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstMulticastAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstDnsServerAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstDnsServerAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstDnsServerAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstDnsServerAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DnsSuffix)) - sp

	if offset != wtIpAdapterAddressesLh_DnsSuffix_Offset {
		t.Errorf("wtIpAdapterAddressesLh.DnsSuffix offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_DnsSuffix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Description)) - sp

	if offset != wtIpAdapterAddressesLh_Description_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Description offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Description_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FriendlyName)) - sp

	if offset != wtIpAdapterAddressesLh_FriendlyName_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FriendlyName offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FriendlyName_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddress)) - sp

	if offset != wtIpAdapterAddressesLh_PhysicalAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.PhysicalAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_PhysicalAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PhysicalAddressLength)) - sp

	if offset != wtIpAdapterAddressesLh_PhysicalAddressLength_Offset {
		t.Errorf("wtIpAdapterAddressesLh.PhysicalAddressLength offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_PhysicalAddressLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterAddressesLh_Flags_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Flags offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Mtu)) - sp

	if offset != wtIpAdapterAddressesLh_Mtu_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Mtu offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Mtu_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.IfType)) - sp

	if offset != wtIpAdapterAddressesLh_IfType_Offset {
		t.Errorf("wtIpAdapterAddressesLh.IfType offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_IfType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OperStatus)) - sp

	if offset != wtIpAdapterAddressesLh_OperStatus_Offset {
		t.Errorf("wtIpAdapterAddressesLh.OperStatus offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_OperStatus_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv6IfIndex)) - sp

	if offset != wtIpAdapterAddressesLh_Ipv6IfIndex_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Ipv6IfIndex offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Ipv6IfIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ZoneIndices)) - sp

	if offset != wtIpAdapterAddressesLh_ZoneIndices_Offset {
		t.Errorf("wtIpAdapterAddressesLh.ZoneIndices offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_ZoneIndices_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstPrefix)) - sp

	if offset != wtIpAdapterAddressesLh_FirstPrefix_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstPrefix offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstPrefix_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TransmitLinkSpeed)) - sp

	if offset != wtIpAdapterAddressesLh_TransmitLinkSpeed_Offset {
		t.Errorf("wtIpAdapterAddressesLh.TransmitLinkSpeed offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_TransmitLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ReceiveLinkSpeed)) - sp

	if offset != wtIpAdapterAddressesLh_ReceiveLinkSpeed_Offset {
		t.Errorf("wtIpAdapterAddressesLh.ReceiveLinkSpeed offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_ReceiveLinkSpeed_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstWinsServerAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstWinsServerAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstWinsServerAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstWinsServerAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstGatewayAddress)) - sp

	if offset != wtIpAdapterAddressesLh_FirstGatewayAddress_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstGatewayAddress offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstGatewayAddress_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv4Metric)) - sp

	if offset != wtIpAdapterAddressesLh_Ipv4Metric_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Ipv4Metric offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Ipv4Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Ipv6Metric)) - sp

	if offset != wtIpAdapterAddressesLh_Ipv6Metric_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Ipv6Metric offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Ipv6Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Luid)) - sp

	if offset != wtIpAdapterAddressesLh_Luid_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Luid offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Luid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv4Server)) - sp

	if offset != wtIpAdapterAddressesLh_Dhcpv4Server_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Dhcpv4Server offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Dhcpv4Server_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.CompartmentId)) - sp

	if offset != wtIpAdapterAddressesLh_CompartmentId_Offset {
		t.Errorf("wtIpAdapterAddressesLh.CompartmentId offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_CompartmentId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NetworkGuid)) - sp

	if offset != wtIpAdapterAddressesLh_NetworkGuid_Offset {
		t.Errorf("wtIpAdapterAddressesLh.NetworkGuid offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_NetworkGuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ConnectionType)) - sp

	if offset != wtIpAdapterAddressesLh_ConnectionType_Offset {
		t.Errorf("wtIpAdapterAddressesLh.ConnectionType offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_ConnectionType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TunnelType)) - sp

	if offset != wtIpAdapterAddressesLh_TunnelType_Offset {
		t.Errorf("wtIpAdapterAddressesLh.TunnelType offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_TunnelType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6Server)) - sp

	if offset != wtIpAdapterAddressesLh_Dhcpv6Server_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Dhcpv6Server offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Dhcpv6Server_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6ClientDuid)) - sp

	if offset != wtIpAdapterAddressesLh_Dhcpv6ClientDuid_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Dhcpv6ClientDuid offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Dhcpv6ClientDuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6ClientDuidLength)) - sp

	if offset != wtIpAdapterAddressesLh_Dhcpv6ClientDuidLength_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Dhcpv6ClientDuidLength offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Dhcpv6ClientDuidLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Dhcpv6Iaid)) - sp

	if offset != wtIpAdapterAddressesLh_Dhcpv6Iaid_Offset {
		t.Errorf("wtIpAdapterAddressesLh.Dhcpv6Iaid offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_Dhcpv6Iaid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.FirstDnsSuffix)) - sp

	if offset != wtIpAdapterAddressesLh_FirstDnsSuffix_Offset {
		t.Errorf("wtIpAdapterAddressesLh.FirstDnsSuffix offset is %d although %d is expected", offset,
			wtIpAdapterAddressesLh_FirstDnsSuffix_Offset)
		return
	}
}
