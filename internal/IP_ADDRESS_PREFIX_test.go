/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_IP_ADDRESS_PREFIX_Size(t *testing.T) {
	const Actual_IP_ADDRESS_PREFIX_Size = unsafe.Sizeof(IP_ADDRESS_PREFIX{})

	if Actual_IP_ADDRESS_PREFIX_Size != IP_ADDRESS_PREFIX_Size {
		t.Errorf("Size of IP_ADDRESS_PREFIX is %d, although %d is expected.", Actual_IP_ADDRESS_PREFIX_Size,
			IP_ADDRESS_PREFIX_Size)
	}
}
