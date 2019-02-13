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
