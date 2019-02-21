/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtMibIpinterfaceRowSize(t *testing.T) {

	const actualTestWtMibIpinterfaceRowSize = unsafe.Sizeof(wtMibIpinterfaceRow{})

	if actualTestWtMibIpinterfaceRowSize != wtMibIpinterfaceRow_Size {
		t.Errorf("Size of wtMibIpinterfaceRow is %d, although %d is expected.",
			actualTestWtMibIpinterfaceRowSize, wtMibIpinterfaceRow_Size)
	}
}

func TestWtMibIpinterfaceRowOffsets(t *testing.T) {

	s := wtMibIpinterfaceRow{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.InterfaceLuid)) - sp

	if offset != wtMibIpinterfaceRow_InterfaceLuid_Offset {
		t.Errorf("wtMibIpinterfaceRow.InterfaceLuid offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_InterfaceLuid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp

	if offset != wtMibIpinterfaceRow_InterfaceIndex_Offset {
		t.Errorf("wtMibIpinterfaceRow.InterfaceIndex offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_InterfaceIndex_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.MaxReassemblySize)) - sp

	if offset != wtMibIpinterfaceRow_MaxReassemblySize_Offset {
		t.Errorf("wtMibIpinterfaceRow.MaxReassemblySize offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_MaxReassemblySize_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIdentifier)) - sp

	if offset != wtMibIpinterfaceRow_InterfaceIdentifier_Offset {
		t.Errorf("wtMibIpinterfaceRow.InterfaceIdentifier offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_InterfaceIdentifier_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.MinRouterAdvertisementInterval)) - sp

	if offset != wtMibIpinterfaceRow_MinRouterAdvertisementInterval_Offset {
		t.Errorf("wtMibIpinterfaceRow.MinRouterAdvertisementInterval offset is %d although %d is expected",
			offset, wtMibIpinterfaceRow_MinRouterAdvertisementInterval_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.MaxRouterAdvertisementInterval)) - sp

	if offset != wtMibIpinterfaceRow_MaxRouterAdvertisementInterval_Offset {
		t.Errorf("wtMibIpinterfaceRow.MaxRouterAdvertisementInterval offset is %d although %d is expected",
			offset, wtMibIpinterfaceRow_MaxRouterAdvertisementInterval_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AdvertisingEnabled)) - sp

	if offset != wtMibIpinterfaceRow_AdvertisingEnabled_Offset {
		t.Errorf("wtMibIpinterfaceRow.AdvertisingEnabled offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_AdvertisingEnabled_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ForwardingEnabled)) - sp

	if offset != wtMibIpinterfaceRow_ForwardingEnabled_Offset {
		t.Errorf("wtMibIpinterfaceRow.ForwardingEnabled offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_ForwardingEnabled_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.WeakHostSend)) - sp

	if offset != wtMibIpinterfaceRow_WeakHostSend_Offset {
		t.Errorf("wtMibIpinterfaceRow.WeakHostSend offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_WeakHostSend_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.WeakHostReceive)) - sp

	if offset != wtMibIpinterfaceRow_WeakHostReceive_Offset {
		t.Errorf("wtMibIpinterfaceRow.WeakHostReceive offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_WeakHostReceive_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.UseAutomaticMetric)) - sp

	if offset != wtMibIpinterfaceRow_UseAutomaticMetric_Offset {
		t.Errorf("wtMibIpinterfaceRow.UseAutomaticMetric offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_UseAutomaticMetric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.UseNeighborUnreachabilityDetection)) - sp

	if offset != wtMibIpinterfaceRow_UseNeighborUnreachabilityDetection_Offset {
		t.Errorf("wtMibIpinterfaceRow.UseNeighborUnreachabilityDetection offset is %d although %d is expected",
			offset, wtMibIpinterfaceRow_UseNeighborUnreachabilityDetection_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ManagedAddressConfigurationSupported)) - sp

	if offset != wtMibIpinterfaceRow_ManagedAddressConfigurationSupported_Offset {
		t.Errorf("wtMibIpinterfaceRow.ManagedAddressConfigurationSupported offset is %d although %d is expected",
			offset, wtMibIpinterfaceRow_ManagedAddressConfigurationSupported_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.OtherStatefulConfigurationSupported)) - sp

	if offset != wtMibIpinterfaceRow_OtherStatefulConfigurationSupported_Offset {
		t.Errorf("wtMibIpinterfaceRow.OtherStatefulConfigurationSupported offset is %d although %d is expected",
			offset, wtMibIpinterfaceRow_OtherStatefulConfigurationSupported_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.AdvertiseDefaultRoute)) - sp

	if offset != wtMibIpinterfaceRow_AdvertiseDefaultRoute_Offset {
		t.Errorf("wtMibIpinterfaceRow.AdvertiseDefaultRoute offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_AdvertiseDefaultRoute_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.RouterDiscoveryBehavior)) - sp

	if offset != wtMibIpinterfaceRow_RouterDiscoveryBehavior_Offset {
		t.Errorf("wtMibIpinterfaceRow.RouterDiscoveryBehavior offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_RouterDiscoveryBehavior_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DadTransmits)) - sp

	if offset != wtMibIpinterfaceRow_DadTransmits_Offset {
		t.Errorf("wtMibIpinterfaceRow.DadTransmits offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_DadTransmits_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.BaseReachableTime)) - sp

	if offset != wtMibIpinterfaceRow_BaseReachableTime_Offset {
		t.Errorf("wtMibIpinterfaceRow.BaseReachableTime offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_BaseReachableTime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.RetransmitTime)) - sp

	if offset != wtMibIpinterfaceRow_RetransmitTime_Offset {
		t.Errorf("wtMibIpinterfaceRow.RetransmitTime offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_RetransmitTime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PathMtuDiscoveryTimeout)) - sp

	if offset != wtMibIpinterfaceRow_PathMtuDiscoveryTimeout_Offset {
		t.Errorf("wtMibIpinterfaceRow.PathMtuDiscoveryTimeout offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_PathMtuDiscoveryTimeout_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.LinkLocalAddressBehavior)) - sp

	if offset != wtMibIpinterfaceRow_LinkLocalAddressBehavior_Offset {
		t.Errorf("wtMibIpinterfaceRow.LinkLocalAddressBehavior offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_LinkLocalAddressBehavior_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.LinkLocalAddressTimeout)) - sp

	if offset != wtMibIpinterfaceRow_LinkLocalAddressTimeout_Offset {
		t.Errorf("wtMibIpinterfaceRow.LinkLocalAddressTimeout offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_LinkLocalAddressTimeout_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ZoneIndices)) - sp

	if offset != wtMibIpinterfaceRow_ZoneIndices_Offset {
		t.Errorf("wtMibIpinterfaceRow.ZoneIndices offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_ZoneIndices_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SitePrefixLength)) - sp

	if offset != wtMibIpinterfaceRow_SitePrefixLength_Offset {
		t.Errorf("wtMibIpinterfaceRow.SitePrefixLength offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_SitePrefixLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Metric)) - sp

	if offset != wtMibIpinterfaceRow_Metric_Offset {
		t.Errorf("wtMibIpinterfaceRow.Metric offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_Metric_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.NlMtu)) - sp

	if offset != wtMibIpinterfaceRow_NlMtu_Offset {
		t.Errorf("wtMibIpinterfaceRow.NlMtu offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_NlMtu_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Connected)) - sp

	if offset != wtMibIpinterfaceRow_Connected_Offset {
		t.Errorf("wtMibIpinterfaceRow.Connected offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_Connected_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsWakeUpPatterns)) - sp

	if offset != wtMibIpinterfaceRow_SupportsWakeUpPatterns_Offset {
		t.Errorf("wtMibIpinterfaceRow.SupportsWakeUpPatterns offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_SupportsWakeUpPatterns_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsNeighborDiscovery)) - sp

	if offset != wtMibIpinterfaceRow_SupportsNeighborDiscovery_Offset {
		t.Errorf("wtMibIpinterfaceRow.SupportsNeighborDiscovery offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_SupportsNeighborDiscovery_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsRouterDiscovery)) - sp

	if offset != wtMibIpinterfaceRow_SupportsRouterDiscovery_Offset {
		t.Errorf("wtMibIpinterfaceRow.SupportsRouterDiscovery offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_SupportsRouterDiscovery_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ReachableTime)) - sp

	if offset != wtMibIpinterfaceRow_ReachableTime_Offset {
		t.Errorf("wtMibIpinterfaceRow.ReachableTime offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_ReachableTime_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.TransmitOffload)) - sp

	if offset != wtMibIpinterfaceRow_TransmitOffload_Offset {
		t.Errorf("wtMibIpinterfaceRow.TransmitOffload offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_TransmitOffload_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.ReceiveOffload)) - sp

	if offset != wtMibIpinterfaceRow_ReceiveOffload_Offset {
		t.Errorf("wtMibIpinterfaceRow.ReceiveOffload offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_ReceiveOffload_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DisableDefaultRoutes)) - sp

	if offset != wtMibIpinterfaceRow_DisableDefaultRoutes_Offset {
		t.Errorf("wtMibIpinterfaceRow.DisableDefaultRoutes offset is %d although %d is expected", offset,
			wtMibIpinterfaceRow_DisableDefaultRoutes_Offset)
		return
	}
}
