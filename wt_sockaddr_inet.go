/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ns-ws2def-sockaddr_in
// SOCKADDR_IN defined in ws2def.h
type wtSockaddrIn struct {
	sin_family AddressFamily
	sin_port   uint16 // USHORT flattened to uint16
	sin_addr   WtInAddr
	sin_zero   [8]uint8 // Windows type: [8]CHAR
}

// wtSockaddrIn constructor. Creates an empty wtSockaddrIn struct.
func NewWtSockaddrIn() *wtSockaddrIn {
	return &wtSockaddrIn{
		sin_family: AF_INET,
		sin_port: 0,
		sin_addr: *NewWtInAddr(),
		sin_zero: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}}
}

func (addr *wtSockaddrIn) String() string {
	return fmt.Sprintf("sin_family: %s; sin_port: %d; IP: %s", addr.sin_family.String(), addr.sin_port,
		addr.sin_addr.toNetIp().String())
}

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2ipdef/ns-ws2ipdef-sockaddr_in6
// SOCKADDR_IN6_LH defined in ws2ipdef.h
type wtSockaddrIn6Lh struct {
	// AF_INET6.
	sin6_family AddressFamily
	// Transport level port number.
	sin6_port uint16 // USHORT flattened to uint16
	// IPv6 flow information.
	sin6_flowinfo uint32 // ULONG flattened to uint32
	// IPv6 address.
	sin6_addr IN6_ADDR
	// Set of interfaces for a scope.
	sin6_scope_id uint32 // ULONG flattened to uint32
}

func (addr *wtSockaddrIn6Lh) String() string {
	return fmt.Sprintf("sin6_family: %s; sin6_port: %d; sin6_flowinfo: %d; sin6_addr: [%s]; sin6_scope_id: %d",
		addr.sin6_family.String(), addr.sin6_port, addr.sin6_flowinfo, addr.sin6_addr.toNetIp().String(),
		addr.sin6_scope_id)
}

// SOCKADDR_IN6 defined in ws2ipdef.h
type wtSockaddrIn6 wtSockaddrIn6Lh

/*
* According to https://docs.microsoft.com/en-us/windows/desktop/api/ws2ipdef/ns-ws2ipdef-_sockaddr_inet
* SOCKADDR_INET is a usnion of several types, and I'll use the largest among them (SOCKADDR_IN6) instead.
 */
type wtSockaddrInet wtSockaddrIn6

func (addr *wtSockaddrInet) isIPv4() bool {
	return addr.sin6_family == AF_INET;
}

func (addr *wtSockaddrInet) isIPv6() bool {
	return addr.sin6_family == AF_INET6;
}

func (addr *wtSockaddrInet) toWtSockaddrIn() (*wtSockaddrIn, error) {

	if addr == nil {
		return nil, nil
	}

	if addr.isIPv4() {
		return (*wtSockaddrIn)(unsafe.Pointer(addr)), nil
	} else {
		return nil,
			fmt.Errorf("Only wtSockaddrInet values with sin6_family = %s can be converted to wtSockaddrIn. In this case sin6_family is %s.",
				AF_INET.String(), addr.sin6_family.String())
	}
}

func (addr *wtSockaddrInet) toWtSockaddrIn6() (*wtSockaddrIn6, error) {

	if addr == nil {
		return nil, nil
	}

	if addr.isIPv6() {
		return (*wtSockaddrIn6)(unsafe.Pointer(addr)), nil
	} else {
		return nil,
			fmt.Errorf("Only wtSockaddrInet values with sin6_family = %s can be converted to wtSockaddrIn6. In this case sin6_family is %s.",
				AF_INET6.String(), addr.sin6_family.String())
	}
}

func (addr *wtSockaddrInet) String() string {
	if addr.isIPv4() {
		ipv4, _ := addr.toWtSockaddrIn()
		return ipv4.String()
	} else {
		ipv6 := wtSockaddrIn6Lh(*addr)
		return (&ipv6).String()
	}
}

func createWtSockaddrInet(address net.IP, port uint16) (*wtSockaddrInet, error) {

	ipv4 := address.To4()

	result := &wtSockaddrInet{}

	if ipv4 != nil {
		// address is IPv4
		result.fillAsWtSockaddrIn(ipv4, port)
		result.sin6_family = AF_INET
		return result, nil
	}

	ipv6 := address.To16()

	if ipv6 == nil {
		return nil, fmt.Errorf("Input parameter doesn't represent a valid IP address.")
	}

	in6_addr, _ := netIpToWtIn6Addr(ipv6)

	result.sin6_family = AF_INET6
	result.sin6_port = port
	result.sin6_flowinfo = 0
	result.sin6_addr = *in6_addr
	result.sin6_scope_id = 0

	return result, nil
}

func (sin *wtSockaddrInet) fillAsWtSockaddrIn(ipv4 net.IP, port uint16) {

	in_addr, _ := netIpToWtInAddr(ipv4)

	sin4 := (*wtSockaddrIn)(unsafe.Pointer(sin))
	sin4.sin_family = AF_INET
	sin4.sin_addr = *in_addr
	sin4.sin_port = port

	for i := 0; i < 8; i++ {
		sin4.sin_zero[i] = 0
	}
}
