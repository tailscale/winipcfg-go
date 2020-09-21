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
// (https://docs.microsoft.com/en-us/windows/win32/api/netioapi/ns-netioapi-mib_ipforward_row2).
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

func getRoutes(family AddressFamily) ([]*Route, error) {

	rows, err := getWtMibIpforwardRow2s(family)

	if err != nil {
		return nil, err
	}

	routes := make([]*Route, len(rows))
	i := 0
	for _, row := range rows {
		route, err := row.toRoute()

		if err != nil {
			return nil, err
		}

		routes[i] = route
		i++
	}

	return routes[:i], nil
}

func getRoute(interfaceLuid uint64, destination *net.IPNet, nextHop *net.IP) (*Route, error) {

	row, err := getWtMibIpforwardRow2Alt(interfaceLuid, destination, nextHop)

	if err == nil {
		return row.toRoute()
	} else {
		return nil, err
	}
}

func (route *Route) copyChangeableFieldsTo(row *wtMibIpforwardRow2) {

	row.SitePrefixLength = route.SitePrefixLength
	row.ValidLifetime = route.ValidLifetime
	row.PreferredLifetime = route.PreferredLifetime
	row.Metric = route.Metric
	row.Protocol = route.Protocol
	row.Loopback = boolToUint8(route.Loopback)
	row.AutoconfigureAddress = boolToUint8(route.AutoconfigureAddress)
	row.Publish = boolToUint8(route.Publish)
	row.Immortal = boolToUint8(route.Immortal)
}

// Returns all the routes. Corresponds to GetIpForwardTable2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2).
func GetRoutes(family AddressFamily) ([]*Route, error) {
	return getRoutes(family)
}

// Adds new route to the system. Similar to Interface.AddRoute() method, but allows setting more options. Additional
// options you can set by using this method are all "changeable" fields of Route struct (see Route.Set() method for more
// details).
func (route *Route) Add() error {

	wtDest, err := route.DestinationPrefix.toWtIpAddressPrefix()

	if err != nil {
		return err
	}

	wtNextHop, err := route.NextHop.toWtSockaddrInet()

	if err != nil {
		return err
	}

	row := getInitializedWtMibIpforwardRow2(route.InterfaceLuid)

	row.DestinationPrefix = *wtDest
	row.NextHop = *wtNextHop

	route.copyChangeableFieldsTo(row)

	return row.add()
}

// Saves (activates) modified Route. Corresponds to SetIpForwardEntry2 function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipforwardentry2).
//
// Note that fields InterfaceLuid, InterfaceIndex, DestinationPrefix and NextHop are used for identifying route to
// change, meaning that they cannot be changed by using this method. Changing some of these fields would cause updating
// some other route. On the other side, fields Age and Origin are read-only, so they also cannot be changed. So fields
// that are "changeable" this way are all between SitePrefixLength and Immortal, inclusive.
// The workflow of using this method is:
// 1) Get Route instance by using any of getter methods (i.e. GetRoutes or any other);
// 2) Change one or more of "changeable" fields enumerated above;
// 3) Calling this method to activate the changes.
func (route *Route) Set() error {

	destination, err := route.DestinationPrefix.toWtIpAddressPrefix()

	if err != nil {
		return err
	}

	nextHop, err := route.NextHop.toWtSockaddrInet()

	if err != nil {
		return err
	}

	old, err := getWtMibIpforwardRow2(route.InterfaceLuid, destination, nextHop)

	if err != nil {
		return err
	}

	route.copyChangeableFieldsTo(old)

	return old.set()
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
Origin: %s`, r.InterfaceLuid, r.InterfaceIndex, r.DestinationPrefix.String(), r.NextHop.String(), r.SitePrefixLength,
		r.ValidLifetime, r.PreferredLifetime, r.Metric, r.Protocol.String(), r.Loopback, r.AutoconfigureAddress,
		r.Publish, r.Immortal, r.Age, r.Origin.String())
}
