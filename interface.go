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

type Interface struct {
	Luid uint64
	Index uint32
	AdapterName string
	FriendlyName string
	UnicastAddress []*UnicastAddress
	DnsSuffix string
	Description string
	PhysicalAddress net.HardwareAddr
	Flags uint32
	Mtu uint32
	IfType IFTYPE
	OperStatus IF_OPER_STATUS
	Ipv6IfIndex uint32
	ZoneIndices [16]uint32
}

func (iaa *IP_ADAPTER_ADDRESSES) toInterface() (*Interface, error) {

	ifc := Interface{
		Luid: iaa.Luid,
		Index: uint32(iaa.IfIndex),
		AdapterName: iaa.getAdapterName(),
		FriendlyName: iaa.getFriendlyName(),
		DnsSuffix: wcharToString(iaa.DnsSuffix),
		Description: wcharToString(iaa.Description),
		Flags: iaa.Flags,
		Mtu: iaa.Mtu,
		IfType: iaa.IfType,
		OperStatus: iaa.OperStatus,
		Ipv6IfIndex: iaa.Ipv6IfIndex,
		ZoneIndices: iaa.ZoneIndices,
	}

	if iaa.PhysicalAddressLength > 0 {

		ifc.PhysicalAddress = net.HardwareAddr(make([]byte, iaa.PhysicalAddressLength, iaa.PhysicalAddressLength))

		for i := uint32(0); i < iaa.PhysicalAddressLength; i++ {
			ifc.PhysicalAddress[i] = iaa.PhysicalAddress[i]
		}
	}

	uap := iaa.FirstUnicastAddress

	var unicastAddresses []*UnicastAddress

	for ; uap != nil; uap = uap.Next {

		ua, err := toUnicastAddress(ifc, uap)

		if err != nil {
			return nil, err
		}

		unicastAddresses = append(unicastAddresses, ua)
	}

	ifc.UnicastAddress = unicastAddresses

	return &ifc, nil
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

	iaas, err := adapterAddresses()

	if err != nil {
		return nil, err
	}

	if iaas == nil {
		return nil, nil
	}

	ifcs := make([]*Interface, len(iaas), len(iaas))

	for i, iaa := range iaas {

		ifc, err := iaa.toInterface()

		if err != nil {
			return nil, err
		}

		ifcs[i] = ifc
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

			ifc, err := a.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
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

			ifc, err := a.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
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

			ifc, err := a.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
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

	result := fmt.Sprintf(
		`
======================== INTERFACE OUTPUT START ========================
Luid: %d
Index: %d
AdapterName: %s
FriendlyName: %s
Unicast addresses:
`, ifc.Luid, ifc.Index, ifc.AdapterName, ifc.FriendlyName)

	for _, ifc := range ifc.UnicastAddress {
		result += fmt.Sprintf("\t%s\n", ifc.String())
	}

	result += fmt.Sprintf(`DnsSuffix: %s
Description: %s
PhysicalAddress: %s
Flags: %d
MTU: %d
IfType: %s
OperStatus: %s
Ipv6IfIndex: %d
ZoneIndices: %v
========================= INTERFACE OUTPUT END =========================
`, ifc.DnsSuffix, ifc.Description, ifc.PhysicalAddress.String(), ifc.Flags, ifc.Mtu, ifc.IfType.String(),
ifc.OperStatus.String(), ifc.Ipv6IfIndex, ifc.ZoneIndices)

	return result
}
