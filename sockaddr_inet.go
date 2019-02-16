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

// Specifies transport address and port.
type SockaddrInet struct {
	// Can be AF_INET (for IPv4) or AF_INET6 (for IPv6).
	Family AddressFamily

	// Port.
	Port uint16

	// IP address. Can be IPv4 or IPv6.
	Address net.IP

	// The rest of the fields are only for IPv6 addresses.

	// IPv6 flow information. NOTE: This field should be used only with IPv4 addresses.
	IPv6FlowInfo uint32

	// Set of interfaces for a scope. NOTE: This field should be used only with IPv4 addresses.
	IPv6ScopeId uint32
}

func (sainet *SockaddrInet) equivalentTo(other *SockaddrInet) bool {

	if sainet == nil || other == nil {
		return false
	}

	switch sainet.Family {
	case AF_INET:
		return other.Family == AF_INET && sainet.Port == other.Port && sainet.Address.Equal(other.Address)
	case AF_INET6:
		return other.Family == AF_INET6 && sainet.Port == other.Port && sainet.IPv6FlowInfo == other.IPv6FlowInfo &&
			sainet.IPv6ScopeId == other.IPv6ScopeId && sainet.Address.Equal(other.Address)
	default:
		// We don't bother to compare invalid addresses.
		return false
	}
}

func createSockaddrInet(ip net.IP) (*SockaddrInet, error) {

	sainet := SockaddrInet{
		Port: 0,
		Address: ip,
		IPv6FlowInfo: 0,
		IPv6ScopeId: 0,
	}

	ip4 := ip.To4()

	if ip4 != nil {

		sainet.Family = AF_INET
		sainet.Address = ip4

		return &sainet, nil
	}

	ip6 := ip.To16()

	if ip6 != nil {

		sainet.Family = AF_INET6
		sainet.Address = ip6

		return &sainet, nil
	}

	return nil, fmt.Errorf("createSockaddrInet() - invalid input IP")
}

func (sainet *SockaddrInet) toWtSockaddrInet() (*wtSockaddrInet, error) {

	if sainet == nil {
		return nil, nil
	}

	wtsainet, err := createWtSockaddrInet(sainet.Address, sainet.Port)

	if err != nil {
		return nil, err
	}

	if wtsainet.sin6_family != sainet.Family {
		switch sainet.Family {
		case AF_INET:
			return nil,
				fmt.Errorf("SockaddrInet.Family value of the input is AF_INET, but it looks that SockaddrInet.Address contains IPv6 address.")
		case AF_INET6:
			return nil,
			fmt.Errorf("SockaddrInet.Family value of the input is AF_INET6, but it looks that SockaddrInet.Address contains IPv4 address.")
		default:
			return nil,
			fmt.Errorf("Input SockaddrInet cannot be converted because its SockaddrInet.Family value %s. Allowed values are AF_INET and AF_INET6.",
				sainet.Family.String())
		}
	}

	if sainet.Family == AF_INET6 {
		wtsainet.sin6_flowinfo = sainet.IPv6FlowInfo
		wtsainet.sin6_scope_id = sainet.IPv6ScopeId
	}

	return wtsainet, nil
}

func (sainet *SockaddrInet) toWtSocketAddress() (*wtSocketAddress, error) {

	if sainet == nil {
		return nil, nil
	}

	wtsainet, err := sainet.toWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	wtsa := wtSocketAddress{lpSockaddr: (*wtSockaddr)(unsafe.Pointer(&wtsainet))}

	if sainet.Family == AF_INET {
		wtsa.iSockaddrLength = wtSockaddrIn_Size
	} else {
		wtsa.iSockaddrLength = wtSockaddrIn6Lh_Size
	}

	return &wtsa, nil
}

func (sainet *SockaddrInet) String() string {

	if sainet == nil {
		return ""
	}

	result := fmt.Sprintf("%s:%d", sainet.Address.String(), sainet.Port)

	if sainet.Family == AF_INET6 {
		result += fmt.Sprintf("; IPv6FlowInfo: %d; IPv6ScopeId: %d", sainet.IPv6FlowInfo, sainet.IPv6ScopeId)
	}

	return result
}
