/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type UnicastAddress struct {

	// The interface this address belong to.
	Interface Interface

	// The rest is from IP_ADAPTER_ADDRESSES_LH

	// TODO: Documentation missing. What is it?
	Length uint32

	// TODO: Documentation missing. What is it?
	Flags uint32

	// The address
	Address SockaddrInet

	PrefixOrigin IP_PREFIX_ORIGIN
	SuffixOrigin IP_SUFFIX_ORIGIN
	DadState IP_DAD_STATE

	ValidLifetime uint32
	PreferredLifetime uint32
	LeaseLifetime uint32
	OnLinkPrefixLength uint8
}

func toUnicastAddress(ifc Interface, iaua *IP_ADAPTER_UNICAST_ADDRESS_LH) (*UnicastAddress, error) {

	if iaua == nil {
		return nil, nil
	}

	sainet, err := iaua.Address.get_SOCKETADDR_INET()

	if err != nil {
		return nil, err
	}

	sa, err := sainet.toSockAddrInet()

	if err != nil {
		return nil, err
	}

	ua := UnicastAddress{
		Interface: ifc,
		Length: iaua.Length,
		Flags: iaua.Flags,
		Address: *sa,
		PrefixOrigin: iaua.PrefixOrigin,
		ValidLifetime: iaua.ValidLifetime,
		PreferredLifetime: iaua.PreferredLifetime,
		LeaseLifetime: iaua.LeaseLifetime,
		OnLinkPrefixLength: iaua.OnLinkPrefixLength,
	}

	return &ua, nil
}

func (ua *UnicastAddress) String() string {

	if ua == nil {
		return ""
	}

	return fmt.Sprintf("Length: %d; Flags: %d; Address: [%s]/%d; PrefixOrigin: %s; SuffixOrigin: %s; DadState: %s; ValidLifetime: %d; PreferredLifetime: %d; LeaseLifetime: %d",
		ua.Length, ua.Flags, ua.Address.String(), ua.OnLinkPrefixLength, ua.PrefixOrigin.String(),
		ua.SuffixOrigin.String(), ua.DadState.String(), ua.ValidLifetime, ua.PreferredLifetime, ua.LeaseLifetime)
}
