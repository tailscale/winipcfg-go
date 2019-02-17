/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
)

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

func getRoutes(family AddressFamily, ifc *Interface) ([]*Route, error) {

	rows, err := getWtMibIpforwardRow2s(family, ifc)

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

func GetRoutes(family AddressFamily) ([]*Route, error) {
	return getRoutes(family, nil)
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
