/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_row2
type MIB_IPFORWARD_ROW2 struct {
	//
	// Key Structure.
	//
	InterfaceLuid     uint64 // Windows type: NET_LUID
	InterfaceIndex    uint32 // Windows type: NET_IFINDEX
	DestinationPrefix IP_ADDRESS_PREFIX
	NextHop           wtSockaddrInet

	//
	// Read-Write Fields.
	//
	SitePrefixLength uint8 // Windows type: UCHAR
	ValidLifetime uint32 // Windows type: ULONG
	PreferredLifetime uint32 // Windows type: ULONG
	Metric uint32 // Windows type: ULONG
	Protocol NlRouteProtocol

	Loopback uint8 // Windows type: BOOLEAN
	AutoconfigureAddress uint8 // Windows type: BOOLEAN
	Publish uint8 // Windows type: BOOLEAN
	Immortal uint8 // Windows type: BOOLEAN

	//
	// Read-Only Fields.
	//
	Age uint32 // Windows type: ULONG
	Origin NlRouteOrigin
}

type PMIB_IPFORWARD_ROW2 *MIB_IPFORWARD_ROW2
