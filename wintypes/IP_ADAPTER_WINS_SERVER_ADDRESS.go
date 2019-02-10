/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import "golang.org/x/sys/windows"

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_wins_server_address_lh
// Defined in iptypes.h
type IP_ADAPTER_WINS_SERVER_ADDRESS_LH struct {
	Length ULONG
	Reserved DWORD
	Next *IP_ADAPTER_WINS_SERVER_ADDRESS_LH
	Address windows.SocketAddress
}

// Defined in iptypes.h
type IP_ADAPTER_WINS_SERVER_ADDRESS IP_ADAPTER_WINS_SERVER_ADDRESS_LH
