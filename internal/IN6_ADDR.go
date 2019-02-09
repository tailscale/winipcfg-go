/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"fmt"
	"net"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/in6addr/ns-in6addr-in6_addr
type IN6_ADDR struct {
	Byte [16]UCHAR
}

func (addr *IN6_ADDR) ToIp() net.IP {
	return net.IP{
		byte(addr.Byte[0]),
		byte(addr.Byte[1]),
		byte(addr.Byte[2]),
		byte(addr.Byte[3]),
		byte(addr.Byte[4]),
		byte(addr.Byte[5]),
		byte(addr.Byte[6]),
		byte(addr.Byte[7]),
		byte(addr.Byte[8]),
		byte(addr.Byte[9]),
		byte(addr.Byte[10]),
		byte(addr.Byte[11]),
		byte(addr.Byte[12]),
		byte(addr.Byte[13]),
		byte(addr.Byte[14]),
		byte(addr.Byte[15]), }
}

func IpTo_IN6_ADDR(ip net.IP) (*IN6_ADDR, error) {

	ip6 := ip.To16()

	if ip6 == nil {
		return nil, fmt.Errorf("Input IP isn't a valid IPv6 address.")
	}

	in6_addr := IN6_ADDR{}

	for i := 0; i < 16; i++ {
		in6_addr.Byte[i] = UCHAR(ip6[i])
	}

	return &in6_addr, nil
}
