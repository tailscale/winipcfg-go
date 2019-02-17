/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

// Defined in iptypes.h
const (
	MAX_ADAPTER_ADDRESS_LENGTH = 8
	MAX_DHCPV6_DUID_LENGTH     = 130
)

// IP_ADAPTER_ADDRESSES defined in iptypes.h
type wtIpAdapterAddresses wtIpAdapterAddressesLh

// Based on function with the same name in 'net' module, in file interface_windows.go
func getWtIpAdapterAddresses() ([]*wtIpAdapterAddresses, error) {

	var b []byte

	size := uint32(15000) // recommended initial size

	for {

		b = make([]byte, size)

		result := getAdaptersAddresses(windows.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0,
			(*wtIpAdapterAddresses)(unsafe.Pointer(&b[0])), &size)

		if result == 0 {

			if size == 0 {
				return nil, nil
			}

			break
		}

		if result != uint32(windows.ERROR_BUFFER_OVERFLOW) {
			return nil, os.NewSyscallError("iphlpapi.GetAdaptersAddresses", windows.Errno(result))
		}

		if size <= uint32(len(b)) {
			return nil, os.NewSyscallError("iphlpapi.GetAdaptersAddresses", windows.Errno(result))
		}
	}

	var wtiaas []*wtIpAdapterAddresses

	for wtiaa := (*wtIpAdapterAddresses)(unsafe.Pointer(&b[0])); wtiaa != nil; wtiaa = wtiaa.nextCasted() {
		wtiaas = append(wtiaas, wtiaa)
	}

	return wtiaas, nil
}

func (aa *wtIpAdapterAddresses) nextCasted() *wtIpAdapterAddresses {
	if aa == nil {
		return nil
	} else {
		return (*wtIpAdapterAddresses)(unsafe.Pointer(aa.Next))
	}
}

func (aa *wtIpAdapterAddresses) getAdapterName() string {
	if aa == nil {
		return ""
	} else {
		return charToString(aa.AdapterName)
	}
}

func (aa *wtIpAdapterAddresses) getFriendlyName() string {
	if aa == nil {
		return ""
	} else {
		return wcharToString(aa.FriendlyName)
	}
}
