/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"testing"
	"unsafe"
)

func Test_IpAdapterAddresses_Size(t *testing.T) {
	const ActualIpAdapterAddressesSize = uint32(unsafe.Sizeof(windows.IpAdapterAddresses{}))

	if ActualIpAdapterAddressesSize != expectedIpAdapterAddressesSize {
		t.Errorf("Size of windows.IpAdapterAddresses is %d, although %d is expected.",
			ActualIpAdapterAddressesSize, expectedIpAdapterAddressesSize)
	}
}

func Test_IpAdapterAddressesEx_Size(t *testing.T) {
	const ActualIpAdapterAddressesExSize = uint32(unsafe.Sizeof(IpAdapterAddressesEx{}))

	if ActualIpAdapterAddressesExSize != ipAdapterAddressesExSize {
		t.Errorf("Size of IpAdapterAddressesEx is %d, although %d is expected.", ActualIpAdapterAddressesExSize,
			ipAdapterAddressesExSize)
	}
}
