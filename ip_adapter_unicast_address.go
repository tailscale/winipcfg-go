/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterUnicastAddress struct {

	// It contains everything from IpAdapterAnycastAddress
	IpAdapterAnycastAddress

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState IpDadState

	ValidLifetime uint32
	PreferredLifetime uint32
	LeaseLifetime uint32
	OnLinkPrefixLength uint8
}

func ipAdapterUnicastAddressFromWinType(ifc Interface, wtua *wtIpAdapterUnicastAddressLh) (*IpAdapterUnicastAddress,
	error) {

	if wtua == nil {
		return nil, nil
	}

	wtsainet, err := wtua.Address.getWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	sainet, err := sockaddrInetFromWinType(wtsainet)

	if err != nil {
		return nil, err
	}

	ua := IpAdapterUnicastAddress{
		PrefixOrigin:       wtua.PrefixOrigin,
		ValidLifetime:      wtua.ValidLifetime,
		PreferredLifetime:  wtua.PreferredLifetime,
		LeaseLifetime:      wtua.LeaseLifetime,
		OnLinkPrefixLength: wtua.OnLinkPrefixLength,
	}

	ua.Interface = ifc
	ua.Length = wtua.Length
	ua.Flags = wtua.Flags
	ua.Address = *sainet

	return &ua, nil
}

func (ua *IpAdapterUnicastAddress) String() string {

	if ua == nil {
		return ""
	}

	return fmt.Sprintf("Length: %d; Flags: %d; Address: [%s]/%d; PrefixOrigin: %s; SuffixOrigin: %s; DadState: %s; ValidLifetime: %d; PreferredLifetime: %d; LeaseLifetime: %d",
		ua.Length, ua.Flags, ua.Address.String(), ua.OnLinkPrefixLength, ua.PrefixOrigin.String(),
		ua.SuffixOrigin.String(), ua.DadState.String(), ua.ValidLifetime, ua.PreferredLifetime, ua.LeaseLifetime)
}
