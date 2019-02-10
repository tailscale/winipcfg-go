/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "unsafe"

// Defined in iptypes.h
const (
	MAX_ADAPTER_ADDRESS_LENGTH = 8
	MAX_DHCPV6_DUID_LENGTH = 130
)

// Defined in iptypes.h
type IP_ADAPTER_ADDRESSES IP_ADAPTER_ADDRESSES_LH

type Tuple1 struct {
	NetworkGuid NET_IF_NETWORK_GUID
	ConnectionType NET_IF_CONNECTION_TYPE
}

func extractNetworkGuid(bytes [20]uint8) *NET_IF_NETWORK_GUID {
	return (*NET_IF_NETWORK_GUID)(unsafe.Pointer(&bytes[0]))
}

func extractConnectionType(bytes [20]uint8) *NET_IF_NETWORK_GUID {
	return (*NET_IF_NETWORK_GUID)(unsafe.Pointer(&bytes[16]))
}
