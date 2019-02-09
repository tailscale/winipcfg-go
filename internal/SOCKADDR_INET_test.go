/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"testing"
	"unsafe"
)

func Test_SOCKADDR_IN_Size(t *testing.T) {
	const Actual_SOCKADDR_IN_Size = unsafe.Sizeof(SOCKADDR_IN{})

	if Actual_SOCKADDR_IN_Size != SOCKADDR_IN_Size {
		t.Errorf("Size of SOCKADDR_IN is %d, although %d is expected.", Actual_SOCKADDR_IN_Size, SOCKADDR_IN_Size)
	}
}

func Test_SOCKADDR_IN6_LH_Size(t *testing.T) {
	const Actual_SOCKADDR_IN6_LH_Size = unsafe.Sizeof(SOCKADDR_IN6_LH{})

	if Actual_SOCKADDR_IN6_LH_Size != SOCKADDR_IN6_LH_Size {
		t.Errorf("Size of SOCKADDR_IN6_LH is %d, although %d is expected.", Actual_SOCKADDR_IN6_LH_Size,
			SOCKADDR_IN6_LH_Size)
	}
}

func Test_SOCKADDR_IN6_Size(t *testing.T) {
	const Actual_SOCKADDR_IN6_Size = unsafe.Sizeof(SOCKADDR_IN6{})

	if Actual_SOCKADDR_IN6_Size != SOCKADDR_IN6_Size {
		t.Errorf("Size of SOCKADDR_IN6 is %d, although %d is expected.", Actual_SOCKADDR_IN6_Size,
			SOCKADDR_IN6_Size)
	}
}

func Test_SOCKADDR_INET_Size(t *testing.T) {
	const Actual_SOCKADDR_INET_Size = unsafe.Sizeof(SOCKADDR_INET{})

	if Actual_SOCKADDR_INET_Size != SOCKADDR_INET_Size {
		t.Errorf("Size of SOCKADDR_INET is %d, although %d is expected.", Actual_SOCKADDR_INET_Size,
			SOCKADDR_INET_Size)
	}
}
