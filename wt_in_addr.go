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
// IN-ADDR defined in inaddr.h
type wtInAddr struct {
	s_b1 uint8 // Windows type: UCHAR
	s_b2 uint8 // Windows type: UCHAR
	s_b3 uint8 // Windows type: UCHAR
	s_b4 uint8 // Windows type: UCHAR
}

func NewWtInAddr() *wtInAddr {
	return &wtInAddr{0, 0, 0, 0}
}

// Compares two wtInAddr structs for equality. Note that the function will return false if either of structs is nil,
// even if the other is also nil.
func (addr *wtInAddr) equal(other *wtInAddr) bool {
	if addr == nil || other == nil {
		return false
	} else {
		return addr.s_b1 == other.s_b1 && addr.s_b2 == other.s_b2 && addr.s_b3 == other.s_b3 && addr.s_b4 == other.s_b4
	}
}

func (addr *wtInAddr) matches(ip net.IP) bool {

	if addr == nil {
		return false
	}

	ip4 := ip.To4()

	if ip4 == nil {
		return false
	}

	return ip4[0] == addr.s_b1 && ip4[1] == addr.s_b2 && ip4[2] == addr.s_b3 && ip4[3] == addr.s_b4
}

func (addr *wtInAddr) toNetIp() net.IP {
	return net.IPv4(byte(addr.s_b1), byte(addr.s_b2), byte(addr.s_b3), byte(addr.s_b4))
}

func netIpToWtInAddr(ip net.IP) (*wtInAddr, error) {

	ip4 := ip.To4()

	if ip4 == nil {
		return nil, fmt.Errorf("Input IP isn't a valid IPv4 address.")
	}

	firstByte := 0

	if len(ip4) == net.IPv6len {
		firstByte = 12
	}

	return &wtInAddr{s_b1: ip4[firstByte], s_b2: ip4[firstByte+1], s_b3: ip4[firstByte+2], s_b4: ip4[firstByte+3]},
		nil
}
