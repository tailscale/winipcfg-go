/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_row2
type MIB_IPFORWARD_ROW2 struct {
	//
	// Key Structure.
	//
	InterfaceLuid     NET_LUID
	InterfaceIndex    NET_IFINDEX
	DestinationPrefix IP_ADDRESS_PREFIX
	NextHop           SOCKADDR_INET

	//
	// Read-Write Fields.
	//
	SitePrefixLength UCHAR
	ValidLifetime ULONG
	PreferredLifetime ULONG
	Metric ULONG
	Protocol NL_ROUTE_PROTOCOL

	Loopback BOOLEAN
	AutoconfigureAddress BOOLEAN
	Publish BOOLEAN
	Immortal BOOLEAN

	//
	// Read-Only Fields.
	//
	Age ULONG
	Origin NL_ROUTE_ORIGIN
}

type PMIB_IPFORWARD_ROW2 *MIB_IPFORWARD_ROW2

func (row *MIB_IPFORWARD_ROW2) String() string {
	return fmt.Sprintf("InterfaceLuid: %+v\nInterfaceIndex: %d\nDestinationPrefix: %s\nNextHop: %s\nSitePrefixLength: %d\nValidLifetime: %d\nPreferredLifetime: %d\nMetric: %d\nProtocol: %s\nLoopback: %s\nAutoconfigureAddress: %s\nPublish: %s\nImmortal: %s\nAge: %d\nOrigin: %s",
		row.InterfaceLuid, row.InterfaceIndex, row.DestinationPrefix.String(), row.NextHop.String(), row.SitePrefixLength,
		row.ValidLifetime, row.PreferredLifetime, row.Metric, row.Protocol.String(), row.Loopback.String(),
		row.AutoconfigureAddress.String(), row.Publish.String(), row.Immortal.String(), row.Age, row.Origin.String())
}
