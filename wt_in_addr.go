/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/inaddr/ns-inaddr-in_addr
// WtInAddr defined in inaddr.h
type WtInAddr struct {
	s_b1 uint8 // Windows type: UCHAR
	s_b2 uint8 // Windows type: UCHAR
	s_b3 uint8 // Windows type: UCHAR
	s_b4 uint8 // Windows type: UCHAR
}

func NewWtInAddr() *WtInAddr {
	return &WtInAddr{0, 0, 0, 0}
}

func (addr *WtInAddr) toNetIp() net.IP {
	return net.IPv4(byte(addr.s_b1), byte(addr.s_b2), byte(addr.s_b3), byte(addr.s_b4))
}

func netIpToWtInAddr(ip net.IP) (*WtInAddr, error) {

	ip4 := ip.To4()

	if ip4 == nil {
		return nil, fmt.Errorf("Input IP isn't a valid IPv4 address.")
	}

	firstByte := 0

	if len(ip4) == net.IPv6len {
		firstByte = 12
	}

	return &WtInAddr{s_b1: ip4[firstByte], s_b2: ip4[firstByte+1], s_b3: ip4[firstByte+2], s_b4: ip4[firstByte+3]},
		nil
}
