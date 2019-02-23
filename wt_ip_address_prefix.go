/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_ip_address_prefix
// IP_ADDRESS_PREFIX defined in netioapi.h
type wtIpAddressPrefix struct {
	Prefix       wtSockaddrInet
	PrefixLength uint8 // Windows type: UINT8
}

func createWtIpAddressPrefix(ipnet *net.IPNet) (*wtIpAddressPrefix, error) {

	if ipnet == nil {
		return nil, fmt.Errorf("createWtIpAddressPrefix() - input argument 'ipnet' is nil")
	}

	wtsainet, err := createWtSockaddrInet(&ipnet.IP, 0)

	if err != nil {
		return nil, err
	}

	ones, _ := ipnet.Mask.Size()

	return &wtIpAddressPrefix{
		Prefix:       *wtsainet,
		PrefixLength: uint8(ones),
	}, nil
}

func (wtap *wtIpAddressPrefix) toIpAddressPrefix() (*IpAddressPrefix, error) {

	if wtap == nil {
		return nil, nil
	}

	sainet, err := wtap.Prefix.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &IpAddressPrefix{Prefix: *sainet, PrefixLength: wtap.PrefixLength}, nil
}

func (addrPfx *wtIpAddressPrefix) String() string {
	if addrPfx == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("[%s]/%d", addrPfx.Prefix.String(), addrPfx.PrefixLength)
	}
}
