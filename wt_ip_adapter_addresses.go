/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"net"
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

func (wtiaa *wtIpAdapterAddresses) toInterface() (*Interface, error) {

	ifc := Interface{
		Luid:              wtiaa.Luid,
		Index:             uint32(wtiaa.IfIndex),
		AdapterName:       wtiaa.getAdapterName(),
		FriendlyName:      wtiaa.getFriendlyName(),
		DnsSuffix:         wcharToString(wtiaa.DnsSuffix, 1000),
		Description:       wcharToString(wtiaa.Description, 1000),
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

	var unicastAddresses []*UnicastAddress

	for wtua := wtiaa.FirstUnicastAddress; wtua != nil; wtua = wtua.Next {

		ua, err := wtua.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		unicastAddresses = append(unicastAddresses, ua)
	}

	ifc.UnicastAddresses = unicastAddresses

	var anycastAddresses []*IpAdapterAddressCommonTypeEx

	for wtaa := wtiaa.FirstAnycastAddress; wtaa != nil; wtaa = wtaa.Next {

		ua, err := wtaa.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		anycastAddresses = append(anycastAddresses, ua)
	}

	ifc.AnycastAddresses = anycastAddresses

	var multicastAddresses []*IpAdapterAddressCommonTypeEx

	for wtma := wtiaa.FirstMulticastAddress; wtma != nil; wtma = wtma.Next {

		ma, err := wtma.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		multicastAddresses = append(multicastAddresses, ma)
	}

	ifc.MulticastAddresses = multicastAddresses

	var dnsServerAddresses []*IpAdapterAddressCommonType

	for wtdsa := wtiaa.FirstDnsServerAddress; wtdsa != nil; wtdsa = wtdsa.Next {

		dsa, err := wtdsa.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		dnsServerAddresses = append(dnsServerAddresses, dsa)
	}

	ifc.DnsServerAddresses = dnsServerAddresses

	var prefixes []*IpAdapterPrefix

	for wtp := wtiaa.FirstPrefix; wtp != nil; wtp = wtp.Next {

		p, err := wtp.toIpAdapterPrefix(ifc)

		if err != nil {
			return nil, err
		}

		prefixes = append(prefixes, p)
	}

	ifc.Prefixes = prefixes

	var winsServerAddresses []*IpAdapterAddressCommonType

	for wtwsa := wtiaa.FirstWinsServerAddress; wtwsa != nil; wtwsa = wtwsa.Next {

		wsa, err := wtwsa.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		winsServerAddresses = append(winsServerAddresses, wsa)
	}

	ifc.WinsServerAddresses = winsServerAddresses

	var gatewayAddresses []*IpAdapterAddressCommonType

	for wtga := wtiaa.FirstGatewayAddress; wtga != nil; wtga = wtga.Next {

		wsa, err := wtga.toIpAdapterAddress(ifc)

		if err != nil {
			return nil, err
		}

		gatewayAddresses = append(gatewayAddresses, wsa)
	}

	ifc.GatewayAddresses = gatewayAddresses

	dhcpv4s, err := (&wtiaa.Dhcpv4Server).toSockaddrInet()

	if err != nil {
		return nil, err
	}

	ifc.Dhcpv4Server = dhcpv4s

	dhcpv6s, err := (&wtiaa.Dhcpv6Server).toSockaddrInet()

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
		dnsSuffixes = append(dnsSuffixes, wcharToString(&dnss.String[0], 1000))
	}

	ifc.DnsSuffixes = dnsSuffixes

	return &ifc, nil
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
		return "<nil>"
	} else {
		return charToString(aa.AdapterName, 1000)
	}
}

func (aa *wtIpAdapterAddresses) getFriendlyName() string {
	if aa == nil {
		return "<nil>"
	} else {
		return wcharToString(aa.FriendlyName, 1000)
	}
}
