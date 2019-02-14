/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_row2
// MIB_IPFORWARD_ROW2 defined in netioapi.h
type wtMibIpforwardRow2 struct {
	//
	// Key Structure.
	//
	InterfaceLuid     uint64 // Windows type: NET_LUID
	InterfaceIndex    uint32 // Windows type: NET_IFINDEX
	DestinationPrefix wtIpAddressPrefix
	NextHop           wtSockaddrInet

	//
	// Read-Write Fields.
	//
	SitePrefixLength  uint8  // Windows type: UCHAR
	ValidLifetime     uint32 // Windows type: ULONG
	PreferredLifetime uint32 // Windows type: ULONG
	Metric            uint32 // Windows type: ULONG
	Protocol          NlRouteProtocol

	Loopback             uint8 // Windows type: BOOLEAN
	AutoconfigureAddress uint8 // Windows type: BOOLEAN
	Publish              uint8 // Windows type: BOOLEAN
	Immortal             uint8 // Windows type: BOOLEAN

	//
	// Read-Only Fields.
	//
	Age    uint32 // Windows type: ULONG
	Origin NlRouteOrigin
}

func (wtr *wtMibIpforwardRow2) toRoute() (*Route, error) {

	if wtr == nil {
		return nil, nil
	}

	iap, err := wtr.DestinationPrefix.toIpAddressPrefix()

	if err != nil {
		return nil, err
	}

	sainet, err := wtr.NextHop.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &Route{
		InterfaceLuid:        wtr.InterfaceLuid,
		InterfaceIndex:       wtr.InterfaceIndex,
		DestinationPrefix:    *iap,
		NextHop:              *sainet,
		SitePrefixLength:     wtr.SitePrefixLength,
		ValidLifetime:        wtr.ValidLifetime,
		PreferredLifetime:    wtr.PreferredLifetime,
		Metric:               wtr.Metric,
		Protocol:             wtr.Protocol,
		Loopback:             uint8ToBool(wtr.Loopback),
		AutoconfigureAddress: uint8ToBool(wtr.AutoconfigureAddress),
		Publish:              uint8ToBool(wtr.Publish),
		Immortal:             uint8ToBool(wtr.Immortal),
		Age:                  wtr.Age,
		Origin:               wtr.Origin,
	}, nil
}
