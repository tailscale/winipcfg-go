/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_wins_server_address_lh
// IP_ADAPTER_WINS_SERVER_ADDRESS_LH defined in iptypes.h
type wtIpAdapterWinsServerAddressLh struct {
	Length   uint32 // Windows type: ULONG
	Reserved uint32 // Windows type: DWORD
	Next     *wtIpAdapterWinsServerAddressLh
	Address  wtSocketAddress

	offset1 [4]uint8 // Layout correction field
}
