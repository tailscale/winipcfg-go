/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_IP_ADAPTER_GATEWAY_ADDRESS_LH_Size(t *testing.T) {

	const Actual_IP_ADAPTER_GATEWAY_ADDRESS_LH_Size = unsafe.Sizeof(IP_ADAPTER_GATEWAY_ADDRESS_LH{})

	if Actual_IP_ADAPTER_GATEWAY_ADDRESS_LH_Size != IP_ADAPTER_GATEWAY_ADDRESS_LH_Size {
		t.Errorf("Size of IP_ADAPTER_GATEWAY_ADDRESS_LH is %d, although %d is expected.", Actual_IP_ADAPTER_GATEWAY_ADDRESS_LH_Size, IP_ADAPTER_GATEWAY_ADDRESS_LH_Size)
	}
}

func Test_IP_ADAPTER_GATEWAY_ADDRESS_LH_Offsets(t *testing.T) {

	s := IP_ADAPTER_GATEWAY_ADDRESS_LH{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Reserved)) - sp

	if offset != IP_ADAPTER_GATEWAY_ADDRESS_LH_Reserved_Offset {
		t.Errorf("IP_ADAPTER_GATEWAY_ADDRESS_LH.Reserved offset is %d although %d is expected", offset, IP_ADAPTER_GATEWAY_ADDRESS_LH_Reserved_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != IP_ADAPTER_GATEWAY_ADDRESS_LH_Next_Offset {
		t.Errorf("IP_ADAPTER_GATEWAY_ADDRESS_LH.Next offset is %d although %d is expected", offset, IP_ADAPTER_GATEWAY_ADDRESS_LH_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != IP_ADAPTER_GATEWAY_ADDRESS_LH_Address_Offset {
		t.Errorf("IP_ADAPTER_GATEWAY_ADDRESS_LH.Address offset is %d although %d is expected", offset, IP_ADAPTER_GATEWAY_ADDRESS_LH_Address_Offset)
		return
	}
}
