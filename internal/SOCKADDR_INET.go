/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import (
	"fmt"
	"net"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ns-ws2def-sockaddr_in
type SOCKADDR_IN struct {
	sin_family ADDRESS_FAMILY
	sin_port USHORT
	sin_addr IN_ADDR
	sin_zero [8]CHAR
}

func (addr *SOCKADDR_IN) String() string {
	return fmt.Sprintf("sin_family: %s; sin_port: %d; IP: %s", addr.sin_family.String(), addr.sin_port, addr.sin_addr.ToIp().String())
}

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2ipdef/ns-ws2ipdef-sockaddr_in6
type SOCKADDR_IN6_LH struct {
	sin6_family ADDRESS_FAMILY // AF_INET6.
	sin6_port USHORT           // Transport level port number.
	sin6_flowinfo ULONG        // IPv6 flow information.
	sin6_addr IN6_ADDR         // IPv6 address.
	sin6_scope_id ULONG        // Set of interfaces for a scope.
}

func (addr *SOCKADDR_IN6_LH) String() string {
	return fmt.Sprintf("sin6_family: %s; sin6_port: %d; sin6_flowinfo: %d; sin6_addr: [%s]; sin6_scope_id: %d",
		addr.sin6_family.String(), addr.sin6_port, addr.sin6_flowinfo, addr.sin6_addr.ToIp().String(), addr.sin6_scope_id)
}

// Defined in ws2ipdef.h
type SOCKADDR_IN6 SOCKADDR_IN6_LH

/*
* According to https://docs.microsoft.com/en-us/windows/desktop/api/ws2ipdef/ns-ws2ipdef-_sockaddr_inet
* SOCKADDR_INET is a usnion of several types, and I'll use the largest among them (SOCKADDR_IN6) instead.
 */
type SOCKADDR_INET SOCKADDR_IN6

func (addr *SOCKADDR_INET) IsIPv4() bool {
	return addr.sin6_family == AF_INET;
}

func (addr *SOCKADDR_INET) IsIPv6() bool {
	return addr.sin6_family == AF_INET6;
}

func (addr *SOCKADDR_INET) ToIPv4() (*SOCKADDR_IN, error) {
	if addr.IsIPv4() {
		return (*SOCKADDR_IN)(unsafe.Pointer(addr)), nil
	} else {
		return nil, fmt.Errorf(
			"Only SOCKADDR_INET values with sin6_family = AF_INET can be converted. In this case sin6_family is %s.",
			addr.sin6_family.String())
	}
}

func (addr *SOCKADDR_INET) String() string {
	if addr.IsIPv4() {
		ipv4, _ := addr.ToIPv4()
		return ipv4.String()
	} else {
		ipv6 := SOCKADDR_IN6_LH(*addr)
		return (&ipv6).String()
	}
}

func Create_SOCKADDR_INET(address net.IP, port uint16) (*SOCKADDR_INET, error) {

	ipv4 := address.To4()

	result := &SOCKADDR_INET{}

	if ipv4 != nil {
		// address is IPv4
		result.FillAs_SOCKADDR_IN(ipv4, port)
		result.sin6_family = AF_INET
		return result, nil
	}

	ipv6 := address.To16()

	if ipv6 == nil {
		return nil, fmt.Errorf("Input parameter doesn't represent a valid IP address.")
	}

	in6_addr, _ := IpTo_IN6_ADDR(ipv6)

	result.sin6_family = AF_INET6
	result.sin6_port = USHORT(port)
	result.sin6_flowinfo = 0
	result.sin6_addr = *in6_addr
	result.sin6_scope_id = 0

	return result, nil
}

func (sin *SOCKADDR_INET) FillAs_SOCKADDR_IN (ipv4 net.IP, port uint16) {

	in_addr, _ := IpTo_IN_ADDR(ipv4)

	sin4 := (*SOCKADDR_IN)(unsafe.Pointer(sin))
	sin4.sin_family = AF_INET
	sin4.sin_addr = *in_addr
	sin4.sin_port = USHORT(port)

	for i := 0; i < 8; i++ {
		sin4.sin_zero[i] = 0
	}
}