/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

type IpAddressPrefix struct {
	Prefix       SockaddrInet
	PrefixLength uint8
}

func (ap *IpAddressPrefix) toNetIpNet() (*net.IPNet, error) {

	if ap == nil {
		return nil, nil
	}

	ipv4 := ap.Prefix.Address.To4()

	if ipv4 != nil {
		return &net.IPNet{
			IP:   ipv4,
			Mask: net.CIDRMask(int(ap.PrefixLength), 32),
		}, nil
	}

	ipv6 := ap.Prefix.Address.To16()

	if ipv6 != nil {
		return &net.IPNet{
			IP:   ipv6,
			Mask: net.CIDRMask(int(ap.PrefixLength), 128),
		}, nil
	}

	return nil, fmt.Errorf("IpAddressPrefix.toNetIpNet() - invalid receiver argument")
}

func (ap *IpAddressPrefix) toWtIpAddressPrefix() (*wtIpAddressPrefix, error) {

	if ap == nil {
		return nil, nil
	}

	wtsainet, err := ap.Prefix.toWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &wtIpAddressPrefix{
		Prefix:       *wtsainet,
		PrefixLength: ap.PrefixLength,
	}, nil
}

func (ap *IpAddressPrefix) String() string {
	if ap == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("%s/%d", ap.Prefix.String(), ap.PrefixLength)
	}
}
