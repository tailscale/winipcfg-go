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

var (
	gatewayIPv4 = net.IPNet{
		IP: net.IPv4zero,
		Mask: net.IPMask(net.IPv4zero),
	}
	gatewayIPv6 = net.IPNet{
		IP: net.IPv6zero,
		Mask: net.IPMask(net.IPv6zero),
	}
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

func interfaceWithLuid(luid uint64, wtiaas []*wtIpAdapterAddresses) *wtIpAdapterAddresses {

	for _, wtiaa := range wtiaas {
		if wtiaa.Luid == luid {
			return wtiaa
		}
	}

	return nil
}

func interfacesFromLuids(luids []uint64, flags *GetAdapterAddressesFlags) ([]*Interface, error) {

	wtiaas, err := getWtIpAdapterAddresses(flags.toGetAdapterAddressesFlagsBytes())

	if err != nil {
		return nil, err
	}

	length := len(luids)

	ifcs := make([]*Interface, length, length)

	for idx, luid := range luids {

		wtiaa := interfaceWithLuid(luid, wtiaas)

		if wtiaa == nil {
			return nil, fmt.Errorf("interfacesFromLuids() - interface with Luid=%d not found", luid)
		}

		ifc, err := wtiaa.toInterface()

		if err != nil {
			return nil, err
		}

		ifcs[idx] = ifc
	}

	return ifcs, nil
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

	wtiaa := interfaceWithLuid(luid, wtiaas)

	if wtiaa == nil {
		return nil, fmt.Errorf("InterfaceFromIndexEx() - interface with specified LUID not found")
	} else {
		return wtiaa.toInterface()
	}
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

// TODO: Check interfaceTable method from 'net' module, interface_windows.go file - it may be useful...

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
	return getRoutes(ifc.Luid, family)
}

// Returns route determined with the input arguments. Corresponds to GetIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardentry2).
// NOTE: If the corresponding route isn't found, the method will return error.
func (ifc *Interface) GetRoute(destination *net.IPNet, nextHop *net.IP) (*Route, error) {
	return getRoute(ifc.Luid, destination, nextHop)
}

// Returns routes which are satisfying defined destination criterion.
func (ifc *Interface) FindRoutes(destination *net.IPNet) ([]*Route, error) {
	return findRoutes(ifc.Luid, destination)
}

// Deletes all interface's routes.
func (ifc *Interface) FlushRoutes() error {

	rows, err := getWtMibIpforwardRow2s(ifc.Luid, AF_UNSPEC)

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

// Adds route to the interface. Corresponds to CreateIpForwardEntry2 function, with added splitDefault feature.
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createipforwardentry2)
func (ifc *Interface) AddRoute(routeData *RouteData, splitDefault bool) error {

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

					// Copying routeData to avoid changing it:
					rd := *routeData

					rd.Destination.Mask = []byte{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // It's now 0::/1

					err := createAndAddWtMibIpforwardRow2(ifc.Luid, &rd)

					if err != nil {
						return err
					}

					rd.Destination.IP = []byte{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // It's now 8000::/1

					return createAndAddWtMibIpforwardRow2(ifc.Luid, &rd)
				}

			} else {
				// IPv4 destination
				if allZeroBytes(dest4) {
					// It is 0.0.0.0/0, so we should split

					// Copying routeData to avoid changing it:
					rd := *routeData

					rd.Destination.Mask = []byte{128, 0, 0, 0} // It's now 0.0.0.0/1

					err := createAndAddWtMibIpforwardRow2(ifc.Luid, &rd)

					if err != nil {
						return err
					}

					rd.Destination.IP = []byte{128, 0, 0, 0} // It's now 128.0.0.0/1

					return createAndAddWtMibIpforwardRow2(ifc.Luid, &rd)
				}
			}
		}
	}

	return createAndAddWtMibIpforwardRow2(ifc.Luid, routeData)
}

// Adds multiple routes to the interface.
func (ifc *Interface) AddRoutes(routesData []*RouteData, splitDefault bool) error {

	for _, rd := range routesData {

		err := ifc.AddRoute(rd, splitDefault)

		if err != nil {
			return err
		}
	}

	return nil
}

// Sets (flush than add) multiple routes to the interface.
func (ifc *Interface) SetRoutes(routesData []*RouteData, splitDefault bool) error {

	err := ifc.FlushRoutes()

	if err != nil {
		return err
	}

	return ifc.AddRoutes(routesData, splitDefault)
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

func (ifc *Interface) GetNetworkAdapterConfiguration() (*NetworkAdapterConfiguration, error) {

	nac, err := getNetworkAdaptersConfigurations(ifc.AdapterName)

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

// The same as DefaultInterfacesEx with 'flags' argument gotten from DefaultGetAdapterAddressesFlags().
func DefaultInterfaces(family AddressFamily) ([]*Interface, error) {
	return DefaultInterfacesEx(family, DefaultGetAdapterAddressesFlags())
}

// Returns all default interfaces (interfaces with 0.0.0.0/0 or ::/0 route, depending on 'family'), ordered by the route
// metric.
func DefaultInterfacesEx(family AddressFamily, flags *GetAdapterAddressesFlags) ([]*Interface, error) {

	if family != AF_INET && family != AF_INET6 {
		return nil, fmt.Errorf("DefaultInterfacesEx() - input argument 'family' has to be either AF_INET or AF_INET6")
	}

	flags.GAA_FLAG_INCLUDE_GATEWAYS = true

	destination := gatewayIPv4

	if family == AF_INET6 {
		destination = gatewayIPv6
	}

	routes, err := findWtMibIpforwardRow2s(0, &destination, family)

	if err != nil {
		return nil, err
	}

	numberOfRoutes := len(routes)

	if numberOfRoutes < 1 {
		return make([]*Interface, 0), nil
	}

	sortWtMibIpforwardRow2sByMetric(routes)

	ifcLuids := make([]uint64, numberOfRoutes, numberOfRoutes)

	ifccounter := 0

	for _, route := range routes {

		added := false

		for i := 0; i < ifccounter; i++ {
			if route.InterfaceLuid == ifcLuids[i] {
				added = true
				break
			}
		}

		if !added {
			ifcLuids[ifccounter] = route.InterfaceLuid
			ifccounter++
		}
	}

	if ifccounter < numberOfRoutes {
		ifcLuids = ifcLuids[:ifccounter]
	}

	return interfacesFromLuids(ifcLuids, flags)
}

// Returns the interface that has 0.0.0.0/0 or ::/0 (depending on 'family').
func DefaultInterface(family AddressFamily) (*Interface, error) {
	return DefaultInterfaceEx(family, DefaultGetAdapterAddressesFlags())
}

// Returns the first item from the slice returned by DefaultInterfacesEx() function, or nil if the returned slice is
// empty.
func DefaultInterfaceEx(family AddressFamily, flags *GetAdapterAddressesFlags) (*Interface, error) {

	ifcs, err := DefaultInterfacesEx(family, flags)

	if err != nil {
		return nil, err
	} else if ifcs == nil || len(ifcs) < 1 {
		return nil, nil
	} else {
		return ifcs[0], nil
	}
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
