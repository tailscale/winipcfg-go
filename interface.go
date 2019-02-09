/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

type InterfaceEx struct {
	net.Interface
	Luid uint64
	// Plus private members if required
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
		ifcs[i] = ifc.toInterfaceEx()
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
		if a.Luid == luid {
			return a.toInterfaceEx(), nil
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
		idx := a.IpAdapterAddresses.IfIndex
		if idx == 0 {
			idx = a.IpAdapterAddresses.Ipv6IfIndex
		}
		if idx == index {
			return a.toInterfaceEx(), nil
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
		if a.name() == name {
			return a.toInterfaceEx(), nil
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
