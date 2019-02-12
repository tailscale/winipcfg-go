/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterUnicastAddress struct {

	// The interface this address belong to.
	Interface Interface

	// The rest is from IP_ADAPTER_ADDRESSES_LH

	// TODO: Documentation missing. What is it?
	Length uint32

	// TODO: Documentation missing. What is it?
	Flags uint32

	// The address
	Address SockaddrInet

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState IpDadState

	ValidLifetime uint32
	PreferredLifetime uint32
	LeaseLifetime uint32
	OnLinkPrefixLength uint8
}

func ipAdapterUnicastAddressFromWinType(ifc Interface, iaua *wtIpAdapterUnicastAddressLh) (*IpAdapterUnicastAddress,
	error) {

	if iaua == nil {
		return nil, nil
	}

	wtsainet, err := iaua.Address.get_SOCKETADDR_INET()

	if err != nil {
		return nil, err
	}

	sainet, err := sockaddrInetFromWinType(wtsainet)

	if err != nil {
		return nil, err
	}

	ua := IpAdapterUnicastAddress{
		Interface: ifc,
		Length: iaua.Length,
		Flags: iaua.Flags,
		Address: *sainet,
		PrefixOrigin: iaua.PrefixOrigin,
		ValidLifetime: iaua.ValidLifetime,
		PreferredLifetime: iaua.PreferredLifetime,
		LeaseLifetime: iaua.LeaseLifetime,
		OnLinkPrefixLength: iaua.OnLinkPrefixLength,
	}

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
