/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func TestWtIpAdapterGatewayAddressLhSize(t *testing.T) {

	const actualWtIpAdapterGatewayAddressLhSize = unsafe.Sizeof(wtIpAdapterGatewayAddressLh{})

	if actualWtIpAdapterGatewayAddressLhSize != wtIpAdapterGatewayAddressLh_Size {
		t.Errorf("Size of wtIpAdapterGatewayAddressLh is %d, although %d is expected.",
			actualWtIpAdapterGatewayAddressLhSize, wtIpAdapterGatewayAddressLh_Size)
	}
}

func TestWtIpAdapterGatewayAddressLhOffsets(t *testing.T) {

	s := wtIpAdapterGatewayAddressLh{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Reserved)) - sp

	if offset != wtIpAdapterGatewayAddressLh_Reserved_Offset {
		t.Errorf("wtIpAdapterGatewayAddressLh.Reserved offset is %d although %d is expected", offset,
			wtIpAdapterGatewayAddressLh_Reserved_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterGatewayAddressLh_Next_Offset {
		t.Errorf("wtIpAdapterGatewayAddressLh.Next offset is %d although %d is expected", offset,
			wtIpAdapterGatewayAddressLh_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterGatewayAddressLh_Address_Offset {
		t.Errorf("wtIpAdapterGatewayAddressLh.Address offset is %d although %d is expected", offset,
			wtIpAdapterGatewayAddressLh_Address_Offset)
		return
	}
}
