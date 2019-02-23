/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// Corresponds to MIB_IPFORWARD_ROW2 defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_row2).
type Route struct {
	InterfaceLuid        uint64
	InterfaceIndex       uint32
	DestinationPrefix    IpAddressPrefix
	NextHop              SockaddrInet
	SitePrefixLength     uint8
	ValidLifetime        uint32
	PreferredLifetime    uint32
	Metric               uint32
	Protocol             NlRouteProtocol
	Loopback             bool
	AutoconfigureAddress bool
	Publish              bool
	Immortal             bool
	Age                  uint32
	Origin               NlRouteOrigin
}

func getRoutes(interfaceLuid uint64, family AddressFamily) ([]*Route, error) {

	rows, err := getWtMibIpforwardRow2s(interfaceLuid, family)

	if err != nil {
		return nil, err
	}

	length := len(rows)

	routes := make([]*Route, length, length)

	for idx, row := range rows {

		route, err := row.toRoute()

		if err != nil {
			return nil, err
		}

		routes[idx] = route
	}

	return routes, nil
}

func findRoute(interfaceLuid uint64, destination *net.IPNet) (*Route, error) {

	row, err := findWtMibIpforwardRow2(interfaceLuid, destination)

	if err != nil {
		return nil, err
	}

	if row == nil {
		return nil, nil
	}

	route, err := row.toRoute()

	if err == nil {
		return route, nil
	} else {
		return nil, err
	}
}

func FindRoute(destination *net.IPNet) (*Route, error) {
	return findRoute(0, destination)
}

//func (route *Route) Add() error {
//
//}

// Returns all the routes. Corresponds to GetIpForwardTable2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2).
func GetRoutes(family AddressFamily) ([]*Route, error) {
	return getRoutes(0, family)
}

func (r *Route) String() string {

	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`
InterfaceLuid: %d
InterfaceIndex: %d
DestinationPrefix: %s
NextHop: %s
SitePrefixLength: %d
ValidLifetime: %d
PreferredLifetime: %d
Metric: %d
Protocol: %s
Loopback: %v
AutoconfigureAddress: %v
Publish: %v
Immortal: %v
Age: %d
Origin: %s
`, r.InterfaceLuid, r.InterfaceIndex, r.DestinationPrefix.String(), r.NextHop.String(), r.SitePrefixLength,
		r.ValidLifetime, r.PreferredLifetime, r.Metric, r.Protocol.String(), r.Loopback, r.AutoconfigureAddress,
		r.Publish, r.Immortal, r.Age, r.Origin.String())
}
