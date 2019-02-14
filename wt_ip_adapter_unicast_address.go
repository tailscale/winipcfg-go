/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
// IP_ADAPTER_UNICAST_ADDRESS_LH defined in iptypes.h
type wtIpAdapterUnicastAddressLh struct {
	Length  uint32 // Windows type: ULONG
	Flags   uint32 // Windows type: DWORD
	Next    *wtIpAdapterUnicastAddressLh
	Address wtSocketAddress

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState     IpDadState

	ValidLifetime      uint32 // Windows type: ULONG
	PreferredLifetime  uint32 // Windows type: ULONG
	LeaseLifetime      uint32 // Windows type: ULONG
	OnLinkPrefixLength uint8  // Windows type: UINT8
}

func (wta *wtIpAdapterUnicastAddressLh) toIpAdapterAddress(ifc Interface) (*UnicastAddress, error) {

	if wta == nil {
		return nil, nil
	}

	sainet, err := (&wta.Address).toSockaddrInet()

	if err != nil {
		return nil, err
	}

	ua := UnicastAddress{
		PrefixOrigin:       wta.PrefixOrigin,
		SuffixOrigin:       wta.SuffixOrigin,
		DadState:           wta.DadState,
		ValidLifetime:      wta.ValidLifetime,
		PreferredLifetime:  wta.PreferredLifetime,
		LeaseLifetime:      wta.LeaseLifetime,
		OnLinkPrefixLength: wta.OnLinkPrefixLength,
	}

	ua.InterfaceLuid = ifc.Luid
	ua.InterfaceIndex = ifc.Index
	ua.Length = wta.Length
	ua.Address = *sainet
	ua.Flags = wta.Flags

	return &ua, nil
}
