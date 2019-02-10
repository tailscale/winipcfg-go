/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import (
	"testing"
	"unsafe"
)

func Test_GUID_Size(t *testing.T) {

	const Actual_GUID_Size = unsafe.Sizeof(GUID{})

	if Actual_GUID_Size != GUID_Size {
		t.Errorf("Size of GUID is %d, although %d is expected.", Actual_GUID_Size, GUID_Size)
	}
}

func Test_NET_IF_NETWORK_GUID_Size(t *testing.T) {

	const Actual_NET_IF_NETWORK_GUID_Size = unsafe.Sizeof(NET_IF_NETWORK_GUID{})

	if Actual_NET_IF_NETWORK_GUID_Size != NET_IF_NETWORK_GUID_Size {
		t.Errorf("Size of NET_IF_NETWORK_GUID is %d, although %d is expected.", Actual_NET_IF_NETWORK_GUID_Size, NET_IF_NETWORK_GUID_Size)
	}
}
