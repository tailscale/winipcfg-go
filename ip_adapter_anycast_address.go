/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterAnycastAddress struct {

	// The interface this address belong to.
	Interface Interface

	// The rest is from wtIpAdapterAddressesLh

	// TODO: Documentation missing. What is it?
	Length uint32

	// TODO: Documentation missing. What is it?
	Flags uint32

	// The address
	Address SockaddrInet
}

func ipAdapterAnycastAddressFromWinType(ifc Interface, wtaa *wtIpAdapterAnycastAddressXp) (*IpAdapterAnycastAddress,
	error) {

	if wtaa == nil {
		return nil, nil
	}

	wtsainet, err := wtaa.Address.getWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	sainet, err := sockaddrInetFromWinType(wtsainet)

	if err != nil {
		return nil, err
	}

	aa := IpAdapterAnycastAddress{
		Interface: ifc,
		Length: wtaa.Length,
		Flags: wtaa.Flags,
		Address: *sainet,
	}

	return &aa, nil
}

func (aa *IpAdapterAnycastAddress) String() string {

	if aa == nil {
		return ""
	}

	return fmt.Sprintf("Length: %d; Flags: %d; Address: [%s]", aa.Length, aa.Flags, aa.Address.String())
}
