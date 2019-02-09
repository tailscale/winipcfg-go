/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"fmt"
	"net"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/inaddr/ns-inaddr-in_addr
type IN_ADDR struct {
	s_b1 UCHAR
	s_b2 UCHAR
	s_b3 UCHAR
	s_b4 UCHAR
}

func (addr *IN_ADDR) ToIp() net.IP {
	return net.IPv4(byte(addr.s_b1), byte(addr.s_b2), byte(addr.s_b3), byte(addr.s_b4))
}

func IpTo_IN_ADDR(ip net.IP) (*IN_ADDR, error) {

	ip4 := ip.To4()

	if ip4 == nil {
		return nil, fmt.Errorf("Input IP isn't a valid IPv4 address.")
	}

	firstByte := 0

	if len(ip4) == net.IPv6len {
		firstByte = 12
	}

	return &IN_ADDR{s_b1: UCHAR(ip4[firstByte]), s_b2: UCHAR(ip4[firstByte + 1]), s_b3: UCHAR(ip4[firstByte + 2]), s_b4: UCHAR(ip4[firstByte + 3])}, nil
}
