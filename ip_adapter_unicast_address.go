/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterUnicastAddress struct {

	// It extends IpAdapterAddressCommonTypeEx
	IpAdapterAddressCommonTypeEx

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState IpDadState

	ValidLifetime uint32
	PreferredLifetime uint32
	LeaseLifetime uint32
	OnLinkPrefixLength uint8
}

func ipAdapterUnicastAddressFromWinType(ifc Interface, wta *wtIpAdapterUnicastAddressLh) (*IpAdapterUnicastAddress,
	error) {

	if wta == nil {
		return nil, nil
	}

	ua := IpAdapterUnicastAddress{
		PrefixOrigin:       wta.PrefixOrigin,
		SuffixOrigin:       wta.SuffixOrigin,
		DadState:           wta.DadState,
		ValidLifetime:      wta.ValidLifetime,
		PreferredLifetime:  wta.PreferredLifetime,
		LeaseLifetime:      wta.LeaseLifetime,
		OnLinkPrefixLength: wta.OnLinkPrefixLength,
	}

	err := ua.setAddress(&wta.Address)

	if err != nil {
		return nil, err
	}

	ua.Interface = ifc
	ua.Length = wta.Length
	ua.Flags = wta.Flags

	return &ua, nil
}

func (ua *IpAdapterUnicastAddress) String() string {

	if ua == nil {
		return ""
	} else {
		return fmt.Sprintf("%s/%d; PrefixOrigin: %s; SuffixOrigin: %s; DadState: %s; ValidLifetime: %d; PreferredLifetime: %d; LeaseLifetime: %d",
			ua.commonTypeExAddressString(), ua.OnLinkPrefixLength, ua.PrefixOrigin.String(), ua.SuffixOrigin.String(),
			ua.DadState.String(), ua.ValidLifetime, ua.PreferredLifetime, ua.LeaseLifetime)
	}
}
