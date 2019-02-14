/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

func (wtua *wtMibUnicastipaddressRow) toMibUnicastipaddressRow() (*UnicastAddressData, error) {

	if wtua == nil {
		return nil, nil
	}

	sai, err := wtua.Address.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &UnicastAddressData{
		Address:            *sai,
		InterfaceLuid:      wtua.InterfaceLuid,
		InterfaceIndex:     wtua.InterfaceIndex,
		PrefixOrigin:       wtua.PrefixOrigin,
		SuffixOrigin:       wtua.SuffixOrigin,
		ValidLifetime:      wtua.ValidLifetime,
		PreferredLifetime:  wtua.PreferredLifetime,
		OnLinkPrefixLength: wtua.OnLinkPrefixLength,
		SkipAsSource:       wtua.SkipAsSource != 0,
		DadState:           wtua.DadState,
		ScopeId:            wtua.ScopeId,
		CreationTimeStamp:  wtua.CreationTimeStamp,
	}, nil
}
