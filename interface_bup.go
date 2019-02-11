/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os"
	"syscall"
	"unsafe"
)

type InterfaceEx struct {
	net.Interface
	Luid uint64
	// Plus private members if required
}

// Created based on interfaceTable method from 'net' module, interface_windows.go file.
func toInterfaceEx(aa *IP_ADAPTER_ADDRESSES) *InterfaceEx {
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
			Name: aa.Name(),
		},
		// TODO: Casting to uint64 won't be needed once we "flatten" Windows types.
		Luid: uint64(aa.Luid),
	}
	if aa.OperStatus == IfOperStatusUp {
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
		for i := uint32(0); i < uint32(aa.PhysicalAddressLength); i++ {
			ifi.HardwareAddr[i] = byte(aa.PhysicalAddress[i])
		}
	}
	return &ifi
}

// Based on function with the same name in 'net' module, in file interface_windows.go
func adapterAddresses() ([]*IP_ADAPTER_ADDRESSES, error) {
	var b []byte
	size := uint32(15000) // recommended initial size
	for {
		b = make([]byte, size)
		result := getAdaptersAddresses(windows.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0,
			(*IP_ADAPTER_ADDRESSES)(unsafe.Pointer(&b[0])), &size)
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

	var aas []*IP_ADAPTER_ADDRESSES
	for aa := (*IP_ADAPTER_ADDRESSES)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.NextCasted() {
		aas = append(aas, aa)
	}
	return aas, nil
}

func GetInterfaces() ([]*InterfaceEx, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	ifcs := make([]*InterfaceEx, len(aa), len(aa))
	for i, ifc := range aa {
		ifcs[i] = toInterfaceEx(ifc)
	}
	return ifcs, nil
}

func InterfaceFromLUID(luid uint64) (*InterfaceEx, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		// TODO: Casting to uint64 won't be needed once we "flatten" Windows types.
		if uint64(a.Luid) == luid {
			return toInterfaceEx(a), nil
		}
	}
	return nil, nil
}

func InterfaceFromIndex(index uint32) (*InterfaceEx, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		idx := uint32(a.IfIndex)
		if idx == 0 {
			idx = uint32(a.Ipv6IfIndex)
		}
		if idx == index {
			return toInterfaceEx(a), nil
		}
	}
	return nil, nil
}

func InterfaceFromName(name string) (*InterfaceEx, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		if a.Name() == name {
			return toInterfaceEx(a), nil
		}
	}
	return nil, nil
}

//// Sets up the interface to be totally blank, with no settings. If the user has
//// subsequently edited the interface particulars or added/removed parts using
//// the "Properties" view, this wipes out those changes.
//func (iface *Interface) FlushInterface() error
//
//// Flush removes all, Add adds, Set flushes then adds.
//func (iface *Interface) FlushAddresses() error
//func (iface *Interface) AddAddresses(addresses []net.IP) error
//func (iface *Interface) SetAddresses(addresses []net.IP) error
//
//// splitDefault converts 0.0.0.0/0 into 0.0.0.0/1 and 128.0.0.0/1,
//// and ::/0 into ::/1 and 8000::/1.
//func (iface *Interface) FlushRoutes() error
//func (iface *Interface) AddRoutes(routes []net.IPNet, splitDefault bool) error
//func (iface *Interface) SetRoutes(routes []net.IPNet, splitDefault bool) error
//
//func (iface *Interface) FlushDNS() error
//func (iface *Interface) AddDNS(dnses []net.IP) error
//func (iface *Interface) SetDNS(dnses []net.IP) error
//
//// These make sure we don't leak through another interface's resolver.
//func (iface *Interface) ForceDNSPriority() (windows.HANDLE, error)
//func UnforceDNSPriority(handle windows.HANDLE) error
//
//func (iface *Interface) func SetMTU(mtu uint16) error
//
//// If metric is zero, then UseAutomaticMetric=true; otherwise
//// UseAutomaticMetric=false and the metric is set for the interface.
//func (iface *Interface) func SetMetric(metric uint32) error
//
//// Calls callback with a default interface if the route to 0.0.0.0/0 changes,
//// or if the default interface's MTU changes.
//func RegisterDefaultInterfaceNotifier(callback func(*Interface)) (windows.HANDLE, error)
//func UnregisterDefaultInterfaceNotifier(handle windows.HANDLE) error
//
//// Returns the interface that has 0.0.0.0/0.
//func DefaultInterface() (*Interface, error)

func (ifc *InterfaceEx) String() string {
	return fmt.Sprintf("Luid: %d; Index: %d; MTU: %d; Name: %s; HardwareAddr: %s", ifc.Luid, ifc.Index, ifc.MTU,
		ifc.Name, ifc.HardwareAddr)
}
