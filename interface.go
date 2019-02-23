/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"net"
)

// Corresponds to Windows struct IP_ADAPTER_ADDRESSES
// (https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh)
type Interface struct {
	Luid                uint64
	Index               uint32
	AdapterName         string
	FriendlyName        string
	UnicastAddresses    []*UnicastAddress
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

// Returns all available interfaces. Corresponds to GetAdaptersAddresses function
// (https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses)
func GetInterfaces() ([]*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	length := len(wtiaas)

	ifcs := make([]*Interface, length, length)

	for i, wtiaa := range wtiaas {

		ifc, err := wtiaa.toInterface()

		if err != nil {
			return nil, err
		}

		ifcs[i] = ifc
	}

	return ifcs, nil
}

// Returns interface with specified LUID.
func InterfaceFromLUID(luid uint64) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	for _, wtiaa := range wtiaas {
		if wtiaa.Luid == luid {

			ifc, err := wtiaa.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
		}
	}

	return nil, nil
}

// Returns interface at specified index.
func InterfaceFromIndex(index uint32) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	for _, wtiaa := range wtiaas {

		idx := wtiaa.IfIndex

		if idx == 0 {
			idx = wtiaa.Ipv6IfIndex
		}

		if idx == index {

			ifc, err := wtiaa.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
		}
	}

	return nil, nil
}

// Returns interface with specified friendly name.
func InterfaceFromFriendlyName(friendlyName string) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses()

	if err != nil {
		return nil, err
	}

	for _, wtiaa := range wtiaas {
		if wtiaa.getFriendlyName() == friendlyName {

			ifc, err := wtiaa.toInterface()

			if err != nil {
				return nil, err
			}

			return ifc, nil
		}
	}

	return nil, nil
}

// Refreshes the interface by loading all it again from Windows.
func (ifc *Interface) Refresh() error {

	ifcnew, err := InterfaceFromLUID(ifc.Luid)

	if err != nil {
		return err
	}

	if ifcnew == nil {
		return fmt.Errorf("Interface.Refresh() - InterfaceFromLUID() returned nil")
	}

	*ifc = *ifcnew

	return nil
}

// Returns corresponding IpInterface.
// Argument 'family' has to be either AF_INET or AF_INET6.
func (ifc *Interface) GetIpInterface(family AddressFamily) (*IpInterface, error) {

	row, err := getWtMibIpinterfaceRow(ifc.Luid, family)

	if err != nil {
		return nil, err
	} else {
		return row.toIpInterface(), nil
	}
}

// Returns corresponding IfRow. Based on GetIfEntry2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex)
func (ifc *Interface) GetIfRow(level MibIfEntryLevel) (*IfRow, error) {
	return GetIfRow(ifc.Luid, level)
}

func (ifc *Interface) GetMatchingUnicastAddressData(ip *net.IP) (*UnicastAddressData, error) {

	if ifc == nil {
		return nil, fmt.Errorf("Interface.GetMatchingUnicastAddressData() - receiver Interface argument is nil")
	}

	wtas, err := getWtMibUnicastipaddressRows(AF_UNSPEC)

	if err != nil {
		return nil, err
	}

	for _, wta := range wtas {
		if wta.InterfaceLuid == ifc.Luid && wta.Address.matches(ip) {

			address, err := wta.toUnicastAddressData()

			if err == nil {
				return address, nil
			} else {
				return nil, err
			}
		}
	}

	return nil, err
}

// TODO: Check interfaceTable method from 'net' module, interface_windows.go file - it may be useful...

//// Sets up the interface to be totally blank, with no settings. If the user has
//// subsequently edited the interface particulars or added/removed parts using
//// the "Properties" view, this wipes out those changes.
//func (iface *Interface) FlushInterface() error

// Flush removes all, Add adds, Set flushes then adds.

// Deletes all interface's unicast addresses.
func (ifc *Interface) FlushAddresses() error {

	if ifc == nil {
		return fmt.Errorf("Interface.FlushAddresses() - input argument is nil")
	}

	wtas, err := getWtMibUnicastipaddressRows(AF_UNSPEC)

	if err != nil {
		return err
	}

	for _, wta := range wtas {
		if wta.InterfaceLuid == ifc.Luid {

			err = wta.delete()

			if err != nil {
				return err
			}
		}
	}

	_ = ifc.Refresh()

	return nil
}

func (ifc *Interface) AddAddresses(addresses []*net.IPNet) error {

	if ifc == nil {
		return fmt.Errorf("Interface.AddAddresses() - input argument is nil")
	}

	for _, ipnet := range addresses {
		if ipnet != nil {

			err := addWtMibUnicastipaddressRow(ifc, ipnet)

			if err != nil {
				return err
			}
		}
	}

	_ = ifc.Refresh()

	return nil
}

func (ifc *Interface) SetAddresses(addresses []*net.IPNet) error {

	if ifc == nil {
		return fmt.Errorf("Interface.SetAddresses() - receiver Interface argument is nil")
	}

	err := ifc.FlushAddresses()

	if err != nil {
		return err
	}

	err = ifc.AddAddresses(addresses)

	if err != nil {
		return err
	}

	return nil
}

func (ifc *Interface) DeleteAddress(ip *net.IP) error {

	if ifc == nil {
		return fmt.Errorf("Interface.DeleteAddress() - receiver Interface argument is nil")
	}

	addr, err := getMatchingWtMibUnicastipaddressRow(ifc.Luid, ip)

	if err != nil {
		return err
	}

	if addr == nil {
		return fmt.Errorf("address not found")
	}

	err = addr.delete()

	if err == nil {
		_ = ifc.Refresh()
	}

	return err
}

func (ifc *Interface) GetRoutes(family AddressFamily) ([]*Route, error) {
	return getRoutes(family, ifc)
}

func (ifc *Interface) FindRoute(destination *net.IPNet) (*Route, error) {

	if ifc == nil {
		// Here we need to panic because ifc == nil have another meaning in findRoute function.
		panic("Interface.FindRoute() - receiver argument is nil")
	}

	if destination == nil {
		return nil, fmt.Errorf("Interface.FindRoute() - 'destination' input argument is nil")
	}

	return findRoute(destination, ifc)
}

// splitDefault converts 0.0.0.0/0 into 0.0.0.0/1 and 128.0.0.0/1,
// and ::/0 into ::/1 and 8000::/1.

func (ifc *Interface) FlushRoutes() error {

	if ifc == nil {
		// Here we need to panic because ifc == nil have another meaning in getWtMibIpforwardRow2s function.
		panic("Interface.FlushRoutes() - receiver argument is nil")
	}

	rows, err := getWtMibIpforwardRow2s(AF_UNSPEC, ifc)

	if err != nil {
		return err
	}

	for _, row := range rows {

		err = row.delete()

		if err != nil {
			return err
		}
	}

	return nil
}

// Adds route. Note that routeData can be changed if splitting takes place.
func (ifc *Interface) AddRoute(routeData *RouteData, splitDefault bool) error {

	if routeData == nil {
		return fmt.Errorf("Interface.AddRoute() - input RouteData argument is nil")
	}

	if splitDefault {

		ones, bits := routeData.Destination.Mask.Size()

		if bits < 1 {
			return fmt.Errorf("Interface.AddRoute() - invalid destination (bits = %d)", bits)
		}

		if ones == 0 {
			// Destination prefix length is 0, so it may be splittable
			dest4 := routeData.Destination.IP.To4()

			if dest4 == nil {
				// IPv6 destination
				dest6 := routeData.Destination.IP.To16()

				if dest6 == nil {
					return fmt.Errorf("Interface.AddRoute() - invalid destination (len = %d)",
						len(routeData.Destination.IP))
				}

				if allZeroBytes(dest6) {
					// It is 0::/0, so we should split
					routeData.Destination.Mask = []byte{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // It's now 0::/1

					err := addWtMibIpforwardRow2(ifc, routeData)

					if err != nil {
						return err
					}

					routeData.Destination.IP = []byte{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // It's now 8000::/1

					return addWtMibIpforwardRow2(ifc, routeData)
				}

			} else {
				// IPv4 destination
				if allZeroBytes(dest4) {
					// It is 0.0.0.0/0, so we should split
					routeData.Destination.Mask = []byte{128, 0, 0, 0} // It's now 0.0.0.0/1

					err := addWtMibIpforwardRow2(ifc, routeData)

					if err != nil {
						return err
					}

					routeData.Destination.IP = []byte{128, 0, 0, 0} // It's now 128.0.0.0/1

					return addWtMibIpforwardRow2(ifc, routeData)
				}
			}
		}
	}

	return addWtMibIpforwardRow2(ifc, routeData)
}

func (ifc *Interface) AddRoutes(routesData []*RouteData, splitDefault bool) error {

	for _, rd := range routesData {

		err := ifc.AddRoute(rd, splitDefault)

		if err != nil {
			return err
		}
	}

	return nil
}

func (ifc *Interface) SetRoutes(routesData []*RouteData, splitDefault bool) error {

	err := ifc.FlushRoutes()

	if err != nil {
		return err
	}

	return ifc.AddRoutes(routesData, splitDefault)
}

func (ifc *Interface) DeleteRoute(destination *net.IPNet) error {

	if ifc == nil {
		// Here we need to panic because ifc == nil have another meaning in findWtMibIpforwardRow2 function.
		panic("Interface.DeleteRoute() - receiver argument is nil")
	}

	if destination == nil {
		return fmt.Errorf("Interface.DeleteRoute() - 'destination' input argument is nil")
	}

	row, err := findWtMibIpforwardRow2(destination, ifc)

	if err != nil {
		return err
	}

	if row == nil {
		return fmt.Errorf("Interface.DeleteRoute() - matching route not found")
	} else {
		return row.delete()
	}
}

func (ifc *Interface) GetNetworkAdapterConfiguration() (*NetworkAdapterConfiguration, error) {

	if ifc == nil {
		// Here we need to panic because ifc == nil have another meaning in getNetworkAdaptersConfigurations function.
		panic("Interface.GetNetworkAdapterConfiguration() - receiver argument is nil")
	}

	nac, err := getNetworkAdaptersConfigurations(ifc)

	if err != nil {
		return nil, err
	} else if nac == nil {
		return nil, fmt.Errorf("GetNetworkAdapterConfiguration() - interface not found")
	} else {
		return nac.(*NetworkAdapterConfiguration), nil
	}
}

func (ifc *Interface) FlushDNS() error {
	return setDnses(ifc, nil, false)
}

func (ifc *Interface) AddDNS(dnses []net.IP) error {
	return setDnses(ifc, dnses, true)
}

func (ifc *Interface) SetDNS(dnses []net.IP) error {
	return setDnses(ifc, dnses, false)
}

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
