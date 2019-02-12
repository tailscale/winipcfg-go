/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
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

func sockaddrInetFromWinType(wtsa *wtSockaddrInet) (*SockaddrInet, error) {

	if wtsa == nil {
		return nil, nil
	}

	if wtsa.isIPv4() {

		wtsa4, _ := wtsa.toWtSockaddrIn()

		sainet := SockaddrInet{
			Family:       AF_INET,
			Port:         wtsa4.sin_port,
			Address:      wtsa4.sin_addr.toNetIp(),
			IPv6FlowInfo: 0,
			IPv6ScopeId:  0,
		}

		return &sainet, nil
	}

	if wtsa.isIPv6() {

		wtsa6, _ := wtsa.toWtSockaddrIn6()

		sainet := SockaddrInet{
			Family:       AF_INET6,
			Port:         wtsa6.sin6_port,
			Address:      wtsa6.sin6_addr.toNetIp(),
			IPv6FlowInfo: wtsa6.sin6_flowinfo,
			IPv6ScopeId:  wtsa6.sin6_scope_id,
		}

		return &sainet, nil
	}

	return nil, fmt.Errorf("Family of the input argument is %s. It has to be either %s or %s",
		wtsa.sin6_family.String(), AF_INET.String(), AF_INET6.String())
}

func sockaddrInetFromWtSocketAddress(wtsa *wtSocketAddress) (*SockaddrInet, error) {

	if wtsa == nil {
		return nil, nil
	}

	wtsainet, err := wtsa.getWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	sainet, err := sockaddrInetFromWinType(wtsainet)

	if err == nil {
		return sainet, nil
	} else {
		return nil, err
	}
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
