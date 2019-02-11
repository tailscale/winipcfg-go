/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_IP_ADAPTER_DNS_SUFFIX_Size(t *testing.T) {

	const Actual_IP_ADAPTER_DNS_SUFFIX_Size = unsafe.Sizeof(IP_ADAPTER_DNS_SUFFIX{})

	if Actual_IP_ADAPTER_DNS_SUFFIX_Size != IP_ADAPTER_DNS_SUFFIX_Size {
		t.Errorf("Size of IP_ADAPTER_DNS_SUFFIX is %d, although %d is expected.", Actual_IP_ADAPTER_DNS_SUFFIX_Size, IP_ADAPTER_DNS_SUFFIX_Size)
	}
}

func Test_IP_ADAPTER_DNS_SUFFIX_Offsets(t *testing.T) {

	s := IP_ADAPTER_DNS_SUFFIX{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.String)) - sp

	if offset != IP_ADAPTER_DNS_SUFFIX_String_Offset {
		t.Errorf("IP_ADAPTER_DNS_SUFFIX.String offset is %d although %d is expected", offset, IP_ADAPTER_DNS_SUFFIX_String_Offset)
		return
	}
}
