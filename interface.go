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
	Luid                uint64
	Index               uint32
	AdapterName         string
	FriendlyName        string
	UnicastAddresses    []*IpAdapterUnicastAddress
	AnycastAddresses    []*IpAdapterAddressCommonTypeEx
	MulticastAddresses  []*IpAdapterAddressCommonTypeEx
	DnsServerAddresses  []*IpAdapterAddressCommonType
	DnsSuffix           string
	Description         string
	PhysicalAddress     net.HardwareAddr
	Flags               uint32
	Mtu                 uint32
	IfType              IfType
	OperStatus          IfOperStatus
	Ipv6IfIndex         uint32
	ZoneIndices         [16]uint32
	Prefixes            []*IpAdapterPrefix
	TransmitLinkSpeed   uint64
	ReceiveLinkSpeed    uint64
	WinsServerAddresses []*IpAdapterAddressCommonType
	GatewayAddresses    []*IpAdapterAddressCommonType
	Ipv4Metric          uint32
	Ipv6Metric          uint32
	Dhcpv4Server        *SockaddrInet
	CompartmentId       uint32
	NetworkGuid         windows.GUID
	ConnectionType      NetIfConnectionType
	TunnelType          TunnelType
	Dhcpv6Server        *SockaddrInet
	Dhcpv6ClientDuid    []uint8
	Dhcpv6Iaid          uint32
	DnsSuffixes         []string
}

func interfaceFromWtIpAdapterAddresses(wtiaa *wtIpAdapterAddresses) (*Interface, error) {

	ifc := Interface{
		Luid:              wtiaa.Luid,
		Index:             uint32(wtiaa.IfIndex),
		AdapterName:       wtiaa.getAdapterName(),
		FriendlyName:      wtiaa.getFriendlyName(),
		DnsSuffix:         wcharToString(wtiaa.DnsSuffix),
		Description:       wcharToString(wtiaa.Description),
		Flags:             wtiaa.Flags,
		Mtu:               wtiaa.Mtu,
		IfType:            wtiaa.IfType,
		OperStatus:        wtiaa.OperStatus,
		Ipv6IfIndex:       wtiaa.Ipv6IfIndex,
		ZoneIndices:       wtiaa.ZoneIndices,
		TransmitLinkSpeed: wtiaa.TransmitLinkSpeed,
		ReceiveLinkSpeed:  wtiaa.ReceiveLinkSpeed,
		Ipv4Metric:        wtiaa.Ipv4Metric,
		Ipv6Metric:        wtiaa.Ipv6Metric,
		CompartmentId:     wtiaa.CompartmentId,
		NetworkGuid:       wtiaa.NetworkGuid,
		ConnectionType:    wtiaa.ConnectionType,
		TunnelType:        wtiaa.TunnelType,
		Dhcpv6Iaid:        wtiaa.Dhcpv6Iaid,
	}

	if wtiaa.PhysicalAddressLength > 0 {

		ifc.PhysicalAddress = net.HardwareAddr(make([]byte, wtiaa.PhysicalAddressLength, wtiaa.PhysicalAddressLength))

		for i := uint32(0); i < wtiaa.PhysicalAddressLength; i++ {
			ifc.PhysicalAddress[i] = wtiaa.PhysicalAddress[i]
		}
	}

	var unicastAddresses []*IpAdapterUnicastAddress

	for wtua := wtiaa.FirstUnicastAddress; wtua != nil; wtua = wtua.Next {

		ua, err := ipAdapterUnicastAddressFromWinType(ifc, wtua)

		if err != nil {
			return nil, err
		}

		unicastAddresses = append(unicastAddresses, ua)
	}

	ifc.UnicastAddresses = unicastAddresses

	var anycastAddresses []*IpAdapterAddressCommonTypeEx

	for wtaa := wtiaa.FirstAnycastAddress; wtaa != nil; wtaa = wtaa.Next {

		ua, err := ipAdapterAddressFromWtAnycastAddress(ifc, wtaa)

		if err != nil {
			return nil, err
		}

		anycastAddresses = append(anycastAddresses, ua)
	}

	ifc.AnycastAddresses = anycastAddresses

	var multicastAddresses []*IpAdapterAddressCommonTypeEx

	for wtma := wtiaa.FirstMulticastAddress; wtma != nil; wtma = wtma.Next {

		ma, err := ipAdapterAddressFromWtMulticastAddress(ifc, wtma)

		if err != nil {
			return nil, err
		}

		multicastAddresses = append(multicastAddresses, ma)
	}

	ifc.MulticastAddresses = multicastAddresses

	var dnsServerAddresses []*IpAdapterAddressCommonType

	for wtdsa := wtiaa.FirstDnsServerAddress; wtdsa != nil; wtdsa = wtdsa.Next {

		dsa, err := ipAdapterAddressFromWtDnsServerAddress(ifc, wtdsa)

		if err != nil {
			return nil, err
		}

		dnsServerAddresses = append(dnsServerAddresses, dsa)
	}

	ifc.DnsServerAddresses = dnsServerAddresses

	var prefixes []*IpAdapterPrefix

	for wtp := wtiaa.FirstPrefix; wtp != nil; wtp = wtp.Next {

		p, err := ipAdapterPrefixFromWinType(ifc, wtp)

		if err != nil {
			return nil, err
		}

		prefixes = append(prefixes, p)
	}

	ifc.Prefixes = prefixes

	var winsServerAddresses []*IpAdapterAddressCommonType

	for wtwsa := wtiaa.FirstWinsServerAddress; wtwsa != nil; wtwsa = wtwsa.Next {

		wsa, err := ipAdapterAddressFromWtWinsServerAddress(ifc, wtwsa)

		if err != nil {
			return nil, err
		}

		winsServerAddresses = append(winsServerAddresses, wsa)
	}

	ifc.WinsServerAddresses = winsServerAddresses

	var gatewayAddresses []*IpAdapterAddressCommonType

	for wtga := wtiaa.FirstGatewayAddress; wtga != nil; wtga = wtga.Next {

		wsa, err := ipAdapterAddressFromWtGatewayAddress(ifc, wtga)

		if err != nil {
			return nil, err
		}

		gatewayAddresses = append(gatewayAddresses, wsa)
	}

	ifc.GatewayAddresses = gatewayAddresses

	dhcpv4s, err := sockaddrInetFromWtSocketAddress(&wtiaa.Dhcpv4Server)

	if err != nil {
		return nil, err
	}

	ifc.Dhcpv4Server = dhcpv4s

	dhcpv6s, err := sockaddrInetFromWtSocketAddress(&wtiaa.Dhcpv6Server)

	if err != nil {
		return nil, err
	}

	ifc.Dhcpv6Server = dhcpv6s

	if wtiaa.Dhcpv6ClientDuidLength > 0 {

		ifc.Dhcpv6ClientDuid = make([]uint8, wtiaa.Dhcpv6ClientDuidLength, wtiaa.Dhcpv6ClientDuidLength)

		for i := uint32(0); i < wtiaa.Dhcpv6ClientDuidLength; i++ {
			ifc.Dhcpv6ClientDuid[i] = wtiaa.Dhcpv6ClientDuid[i]
		}
	}

	var dnsSuffixes []string

	for dnss := wtiaa.FirstDnsSuffix; dnss != nil; dnss = dnss.Next {
		dnsSuffixes = append(dnsSuffixes, wcharToString(&dnss.String[0]))
	}

	ifc.DnsSuffixes = dnsSuffixes

	return &ifc, nil
}

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

		if result != uint32(syscall.ERROR_BUFFER_OVERFLOW) {
			return nil, os.NewSyscallError("getadaptersaddresses", syscall.Errno(result))
		}

		if size <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", syscall.Errno(result))
		}
	}

	var wtiaas []*wtIpAdapterAddresses

	for wtiaa := (*wtIpAdapterAddresses)(unsafe.Pointer(&b[0])); wtiaa != nil; wtiaa = wtiaa.nextCasted() {
		wtiaas = append(wtiaas, wtiaa)
	}

	return wtiaas, nil
}

func GetInterfaces() ([]*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	if wtiaas == nil {
		return nil, nil
	}

	ifcs := make([]*Interface, len(wtiaas), len(wtiaas))

	for i, wtiaa := range wtiaas {

		ifc, err := interfaceFromWtIpAdapterAddresses(wtiaa)

		if err != nil {
			return nil, err
		}

		ifcs[i] = ifc
	}

	return ifcs, nil
}

func InterfaceFromLUID(luid uint64) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	if wtiaas == nil {
		return nil, nil
	}

	for _, wtiaa := range wtiaas {
		if wtiaa.Luid == luid {

			ifc, err := interfaceFromWtIpAdapterAddresses(wtiaa)

			if err != nil {
				return nil, err
			}

			return ifc, nil
		}
	}

	return nil, nil
}

func InterfaceFromIndex(index uint32) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	if wtiaas == nil {
		return nil, nil
	}

	for _, wtiaa := range wtiaas {

		idx := wtiaa.IfIndex

		if idx == 0 {
			idx = wtiaa.Ipv6IfIndex
		}

		if idx == index {

			ifc, err := interfaceFromWtIpAdapterAddresses(wtiaa)

			if err != nil {
				return nil, err
			}

			return ifc, nil
		}
	}

	return nil, nil
}

func InterfaceFromFriendlyName(friendlyName string) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	if wtiaas == nil {
		return nil, nil
	}

	for _, wtiaa := range wtiaas {
		if wtiaa.getFriendlyName() == friendlyName {

			ifc, err := interfaceFromWtIpAdapterAddresses(wtiaa)

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
		`Luid: %d
Index: %d
AdapterName: %s
FriendlyName: %s
UnicastAddresses:
`, ifc.Luid, ifc.Index, ifc.AdapterName, ifc.FriendlyName)

	for _, ua := range ifc.UnicastAddresses {
		result += fmt.Sprintf("\t%s\n", ua.String())
	}

	result += "AnycastAddresses:\n"

	for _, aa := range ifc.AnycastAddresses {
		result += fmt.Sprintf("\t%s\n", aa.String())
	}

	result += "MulticastAddresses:\n"

	for _, ma := range ifc.MulticastAddresses {
		result += fmt.Sprintf("\t%s\n", ma.String())
	}

	result += "DnsServerAddresses:\n"

	for _, dsa := range ifc.DnsServerAddresses {
		result += fmt.Sprintf("\t%s\n", dsa.String())
	}

	result += fmt.Sprintf(`DnsSuffix: %s
Description: %s
PhysicalAddress: %s
Flags: %d
Mtu: %d
IfType: %s
OperStatus: %s
Ipv6IfIndex: %d
ZoneIndices: %v
Prefixes:
`, ifc.DnsSuffix, ifc.Description, ifc.PhysicalAddress.String(), ifc.Flags, ifc.Mtu, ifc.IfType.String(),
ifc.OperStatus.String(), ifc.Ipv6IfIndex, ifc.ZoneIndices)

	for _, p := range ifc.Prefixes {
		result += fmt.Sprintf("\t%s\n", p.String())
	}

	result += fmt.Sprintf("TransmitLinkSpeed: %d\nReceiveLinkSpeed: %d\nWinsServerAddresses:\n",
		ifc.TransmitLinkSpeed, ifc.ReceiveLinkSpeed)

	for _, wsa := range ifc.WinsServerAddresses {
		result += fmt.Sprintf("\t%s\n", wsa.String())
	}

	result += "GatewayAddresses:\n"

	for _, ga := range ifc.GatewayAddresses {
		result += fmt.Sprintf("\t%s\n", ga.String())
	}

	result += fmt.Sprintf(`Ipv4Metric: %d
Ipv6Metric: %d
Dhcpv4Server: %s
CompartmentId: %d
NetworkGuid: %v
ConnectionType: %s
TunnelType: %s
Dhcpv6Server: %s
Dhcpv6ClientDuid: %v
Dhcpv6Iaid: %d
`, ifc.Ipv4Metric, ifc.Ipv6Metric, ifc.Dhcpv4Server.String(), ifc.CompartmentId, guidToString(ifc.NetworkGuid),
ifc.ConnectionType.String(), ifc.TunnelType.String(), ifc.Dhcpv6Server.String(), ifc.Dhcpv6ClientDuid, ifc.Dhcpv6Iaid)

	result += "DnsSuffixes:\n"

	for _, dnss := range ifc.DnsSuffixes {
		result += fmt.Sprintf("\t%s\n", dnss)
	}

	return result
}
