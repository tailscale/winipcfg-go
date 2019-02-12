/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterPrefix struct {
	Length uint32
	Flags uint32
	Address SockaddrInet
	PrefixLength uint32
}

func ipAdapterPrefixFromWinType(wt *IP_ADAPTER_PREFIX_XP) (*IpAdapterPrefix, error) {

	if wt == nil {
		return nil, nil
	}

	wtsai, err := wt.Address.get_SOCKETADDR_INET()

	if err != nil {
		return nil, err
	}

	sai, err := sockaddrInetFromWinType(wtsai)

	if err != nil {
		return nil, err
	}

	ap := IpAdapterPrefix{
		Length: wt.Length,
		Flags: wt.Flags,
		Address: *sai,
		PrefixLength: wt.PrefixLength,
	}

	return &ap, nil
}

func (ap *IpAdapterPrefix) String() string {
	return fmt.Sprintf("Length: %d; Flags: %d; Address: [%s]/%d", ap.Length, ap.Flags, ap.Address.String(),
		ap.PrefixLength)
}
