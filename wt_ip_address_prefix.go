/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_ip_address_prefix
// IP_ADDRESS_PREFIX defined in netioapi.h
type wtIpAddressPrefix struct {
	Prefix       wtSockaddrInet
	PrefixLength uint8 // Windows type: UINT8
}

func (addrPfx *wtIpAddressPrefix) String() string {
	return fmt.Sprintf("Prefix: [%s]; PrefixLength: %d", addrPfx.Prefix.String(), addrPfx.PrefixLength)
}
