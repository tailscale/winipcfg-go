/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"net"
	"os"
	"syscall"
	"unsafe"
)

// I've had to extend the original windows.IpAdapterAddresses because it doesn't contain Luid field.
type IpAdapterAddressesEx struct {
	windows.IpAdapterAddresses
	offset1 [ipAdapterAddressesExOffset1Size]byte
	Luid uint64
	offset2 [ipAdapterAddressesExOffset2Size]byte
}

// Based on function with the same name in 'net' module, in file interface_windows.go
func adapterAddresses() ([]*IpAdapterAddressesEx, error) {
	var b []byte
	size := uint32(15000) // recommended initial size
	for {
		b = make([]byte, size)
		result := getAdaptersAddresses(windows.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0,
			(*IpAdapterAddressesEx)(unsafe.Pointer(&b[0])), &size)
		if result == 0 {
			if size == 0 {
				return nil, nil
			}
			break
		}
		if result != uint32(syscall.ERROR_BUFFER_OVERFLOW) {
			return nil, os.NewSyscallError("getadaptersaddresses", syscall.Errno(result))
		}
		if size <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", syscall.Errno(result))
		}
	}

	var aas []*IpAdapterAddressesEx
	for aa := (*IpAdapterAddressesEx)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.next() {
		aas = append(aas, aa)
	}
	return aas, nil
}

func (aa *IpAdapterAddressesEx) next() *IpAdapterAddressesEx {
	return (*IpAdapterAddressesEx) (unsafe.Pointer(aa.Next))
}

func (aa *IpAdapterAddressesEx) name() string {
	return windows.UTF16ToString((*(*[10000]uint16)(unsafe.Pointer(aa.FriendlyName)))[:])
}

// Created based on interfaceTable method from 'net' module, interface_windows.go file.
func (aa *IpAdapterAddressesEx) toInterfaceEx() *InterfaceEx {
	if aa == nil {
		return nil
	}
	index := aa.IfIndex
	if index == 0 { // ipv6IfIndex is a substitute for ifIndex
		index = aa.Ipv6IfIndex
	}
	ifi := InterfaceEx{
		Interface: net.Interface{
			Index: int(index),
			Name: aa.name(),
		},
		Luid: aa.Luid,
	}
	if aa.OperStatus == windows.IfOperStatusUp {
		ifi.Flags |= net.FlagUp
	}
	// For now we need to infer link-layer service
	// capabilities from media types.
	// TODO: use MIB_IF_ROW2.AccessType now that we no longer support
	// Windows XP.
	switch aa.IfType {
	case windows.IF_TYPE_ETHERNET_CSMACD, windows.IF_TYPE_ISO88025_TOKENRING, windows.IF_TYPE_IEEE80211, windows.IF_TYPE_IEEE1394:
		ifi.Flags |= net.FlagBroadcast | net.FlagMulticast
	case windows.IF_TYPE_PPP, windows.IF_TYPE_TUNNEL:
		ifi.Flags |= net.FlagPointToPoint | net.FlagMulticast
	case windows.IF_TYPE_SOFTWARE_LOOPBACK:
		ifi.Flags |= net.FlagLoopback | net.FlagMulticast
	case windows.IF_TYPE_ATM:
		ifi.Flags |= net.FlagBroadcast | net.FlagPointToPoint | net.FlagMulticast // assume all services available; LANE, point-to-point and point-to-multipoint
	}
	if aa.Mtu == 0xffffffff {
		ifi.MTU = -1
	} else {
		ifi.MTU = int(aa.Mtu)
	}
	if aa.PhysicalAddressLength > 0 {
		ifi.HardwareAddr = make(net.HardwareAddr, aa.PhysicalAddressLength)
		copy(ifi.HardwareAddr, aa.PhysicalAddress[:])
	}
	return &ifi
}
