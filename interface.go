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

// The same as GetInterfacesEx() with 'flags' input argument gotten from DefaultGetAdapterAddressesFlags().
func GetInterfaces() ([]*Interface, error) {
	return GetInterfacesEx(DefaultGetAdapterAddressesFlags())
}

// Returns all available interfaces. Corresponds to GetAdaptersAddresses function
// (https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses)
func GetInterfacesEx(flags *GetAdapterAddressesFlags) ([]*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses(flags.toGetAdapterAddressesFlagsBytes())

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

// The same as InterfaceFromLUIDEx() with 'flags' input argument gotten from DefaultGetAdapterAddressesFlags().
func InterfaceFromLUID(luid uint64) (*Interface, error) {
	return InterfaceFromLUIDEx(luid, DefaultGetAdapterAddressesFlags())
}

// Returns interface with specified LUID.
func InterfaceFromLUIDEx(luid uint64, flags *GetAdapterAddressesFlags) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses(flags.toGetAdapterAddressesFlagsBytes())

	if err != nil {
		return nil, err
	}

	for _, wtiaa := range wtiaas {
		if wtiaa.Luid == luid {
			return wtiaa.toInterface()
		}
	}

	return nil, fmt.Errorf("InterfaceFromIndexEx() - interface with specified LUID not found")
}

// The same as InterfaceFromIndexEx() with 'flags' input argument gotten from DefaultGetAdapterAddressesFlags().
func InterfaceFromIndex(index uint32) (*Interface, error) {
	return InterfaceFromIndexEx(index, DefaultGetAdapterAddressesFlags())
}

// Returns interface at specified index.
func InterfaceFromIndexEx(index uint32, flags *GetAdapterAddressesFlags) (*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses(flags.toGetAdapterAddressesFlagsBytes())

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

	return nil, fmt.Errorf("InterfaceFromIndexEx() - interface with specified index not found")
}

// The same as InterfaceFromFriendlyNameEx() with 'flags' input argument gotten from DefaultGetAdapterAddressesFlags().
func InterfaceFromFriendlyName(friendlyName string) (*Interface, error) {
	return InterfaceFromFriendlyNameEx(friendlyName, DefaultGetAdapterAddressesFlags())
}

// Returns interface with specified friendly name.
func InterfaceFromFriendlyNameEx(friendlyName string, flags *GetAdapterAddressesFlags) (*Interface, error) {

	flags.GAA_FLAG_SKIP_FRIENDLY_NAME = false

	wtiaas, err := getWtIpAdapterAddresses(flags.toGetAdapterAddressesFlagsBytes())

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

	return nil, fmt.Errorf("InterfaceFromFriendlyNameEx() - interface with specified friendly name not found")
}

// The same as InterfaceFromGUIDEx() with 'flags' input argument gotten from DefaultGetAdapterAddressesFlags().
func InterfaceFromGUID(guid *windows.GUID) (*Interface, error) {
	return InterfaceFromGUIDEx(guid, DefaultGetAdapterAddressesFlags())
}

// Returns interface with specified GUID. Note that Interface struct doesn't contain interface GUID field.
func InterfaceFromGUIDEx(guid *windows.GUID, flags *GetAdapterAddressesFlags) (*Interface, error) {

	luid, err := InterfaceGuidToLuid(guid)

	if err != nil {
		return nil, err
	}

	return InterfaceFromLUIDEx(luid, flags)
}

// Returns IpInterface struct that corresponds to the interface. Corresponds to GetIpInterfaceEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfaceentry).
// Argument 'family' has to be either AF_INET or AF_INET6.
func (ifc *Interface) GetIpInterface(family AddressFamily) (*IpInterface, error) {
	return GetIpInterface(ifc.Luid, family)
}

// Returns IfRow struct that corresponds to the interface. Based on GetIfEntry2Ex function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex).
func (ifc *Interface) GetIfRow(level MibIfEntryLevel) (*IfRow, error) {
	return GetIfRow(ifc.Luid, level)
}

//// Sets up the interface to be totally blank, with no settings. If the user has
//// subsequently edited the interface particulars or added/removed parts using
//// the "Properties" view, this wipes out those changes.
//func (iface *Interface) FlushInterface() error

// Flush removes all, Add adds, Set flushes then adds.

// Returns UnicastIpAddressRow struct that matches to provided 'ip' argument. Corresponds to GetUnicastIpAddressEntry
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddressentry)
func (ifc *Interface) GetUnicastIpAddressRow(ip *net.IP) (*UnicastIpAddressRow, error) {

	row, err := getWtMibUnicastipaddressRow(ifc.Luid, ip)

	if err == nil {
		return row.toUnicastIpAddressRow()
	} else {
		return nil, err
	}
}

// Deletes all interface's unicast IP addresses.
func (ifc *Interface) FlushAddresses() error {

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

	return nil
}

// Adds new unicast IP address to the interface. Corresponds to CreateUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createunicastipaddressentry).
func (ifc *Interface) AddAddress(address *net.IPNet) error {
	return createAndAddWtMibUnicastipaddressRow(ifc.Luid, address)
}

// Adds multiple new unicast IP addresses to the interface. Corresponds to CreateUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createunicastipaddressentry).
func (ifc *Interface) AddAddresses(addresses []*net.IPNet) error {

	for _, ipnet := range addresses {
		if ipnet != nil {

			err := createAndAddWtMibUnicastipaddressRow(ifc.Luid, ipnet)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Sets interface's unicast IP addresses.
func (ifc *Interface) SetAddresses(addresses []*net.IPNet) error {

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

// Deletes interface's unicast IP address. Corresponds to DeleteUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteunicastipaddressentry).
func (ifc *Interface) DeleteAddress(ip *net.IP) error {

	addr, err := getWtMibUnicastipaddressRow(ifc.Luid, ip)

	if err != nil {
		return err
	}

	return addr.delete()
}

// Returns all the interface's routes. Corresponds to GetIpForwardTable2 function, but filtered by interface.
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2)
func (ifc *Interface) GetRoutes(family AddressFamily) ([]*Route, error) {
	routes, err := getRoutes(family)
	if err != nil {
		return nil, err
	}
	matches := make([]*Route, len(routes))
	i := 0
	for _, route := range routes {
		if route.InterfaceLuid == ifc.Luid {
			matches[i] = route
			i++
		}
	}
	return matches[:i], nil
}

// Returns route determined with the input arguments. Corresponds to GetIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardentry2).
// NOTE: If the corresponding route isn't found, the method will return error.
func (ifc *Interface) GetRoute(destination *net.IPNet, nextHop *net.IP) (*Route, error) {
	return getRoute(ifc.Luid, destination, nextHop)
}

// Deletes all interface's routes.
func (ifc *Interface) FlushRoutes() error {

	rows, err := getWtMibIpforwardRow2s(AF_UNSPEC)

	if err != nil {
		return err
	}

	for _, row := range rows {
		if row.InterfaceLuid != ifc.Luid {
			continue
		}
		err = row.delete()

		if err != nil {
			return err
		}
	}

	return nil
}

// Adds route to the interface. Corresponds to CreateIpForwardEntry2 function, with added splitDefault feature.
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createipforwardentry2)
func (ifc *Interface) AddRoute(routeData *RouteData) error {
	return createAndAddWtMibIpforwardRow2(ifc.Luid, routeData)
}

// Adds multiple routes to the interface.
func (ifc *Interface) AddRoutes(routesData []*RouteData) error {

	for _, rd := range routesData {

		err := ifc.AddRoute(rd)

		if err != nil {
			return err
		}
	}

	return nil
}

// Sets (flush than add) multiple routes to the interface.
func (ifc *Interface) SetRoutes(routesData []*RouteData) error {

	err := ifc.FlushRoutes()

	if err != nil {
		return err
	}

	return ifc.AddRoutes(routesData)
}

// Deletes a route that matches the criteria. Corresponds to DeleteIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteipforwardentry2).
func (ifc *Interface) DeleteRoute(destination *net.IPNet, nextHop *net.IP) error {

	row, err := getWtMibIpforwardRow2Alt(ifc.Luid, destination, nextHop)

	if err == nil {
		return row.delete()
	} else {
		return err
	}
}

func (ifc *Interface) FlushDNS() error {
	return runNetsh(flushDnsCmds(ifc))
}

func (ifc *Interface) AddDNS(dnses []net.IP) error {
	return runNetsh(addDnsCmds(ifc, dnses))
}

func (ifc *Interface) SetDNS(dnses []net.IP) error {
	return runNetsh(append(flushDnsCmds(ifc), addDnsCmds(ifc, dnses)...))
}

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
`, ifc.Ipv4Metric, ifc.Ipv6Metric, ifc.Dhcpv4Server.String(), ifc.CompartmentId, guidToString(&ifc.NetworkGuid),
		ifc.ConnectionType.String(), ifc.TunnelType.String(), ifc.Dhcpv6Server.String(), ifc.Dhcpv6ClientDuid, ifc.Dhcpv6Iaid)

	result += "DnsSuffixes:\n"

	for _, dnss := range ifc.DnsSuffixes {
		result += fmt.Sprintf("\t%s\n", dnss)
	}

	return result
}
