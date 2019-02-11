/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"syscall"
	"unsafe"
)

type Interface struct {
	Luid uint64
	Index uint32
	AdapterName string
	FriendlyName string
}

func (iaa *IP_ADAPTER_ADDRESSES) toInterface() *Interface {

	ifc := Interface{
		Luid: iaa.Luid,
		Index: uint32(iaa.IfIndex),
		AdapterName: iaa.getAdapterName(),
		FriendlyName: iaa.getFriendlyName(),
	}

	return &ifc
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

func GetInterfaces() ([]*Interface, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	ifcs := make([]*Interface, len(aa), len(aa))
	for i, ifc := range aa {
		ifcs[i] = ifc.toInterface()
	}
	return ifcs, nil
}

func InterfaceFromLUID(luid uint64) (*Interface, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		if a.Luid == luid {
			return a.toInterface(), nil
		}
	}
	return nil, nil
}

func InterfaceFromIndex(index uint32) (*Interface, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		idx := a.IfIndex
		if idx == 0 {
			idx = a.Ipv6IfIndex
		}
		if idx == index {
			return a.toInterface(), nil
		}
	}
	return nil, nil
}

func InterfaceFromFriendlyName(friendlyName string) (*Interface, error) {
	aa, err := adapterAddresses()
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, nil
	}
	for _, a := range aa {
		if a.getFriendlyName() == friendlyName {
			return a.toInterface(), nil
		}
	}
	return nil, nil
}

// TODO: Check interfaceTable method from 'net' module, interface_windows.go file - it may be useful...

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

func (ifc *Interface) String() string {
	return fmt.Sprintf("Luid: %d; Index: %d; AdapterName: %s; FriendlyName: %s", ifc.Luid, ifc.Index,
		ifc.AdapterName, ifc.FriendlyName)
}
