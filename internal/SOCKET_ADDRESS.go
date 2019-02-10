/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ns-ws2def-_socket_address
// Defined in ws2def.h
type SOCKET_ADDRESS struct {
	lpSockaddr *SOCKADDR
	iSockaddrLength INT
}

// https://docs.microsoft.com/en-us/windows/desktop/WinSock/sockaddr-2
// Defined in ws2def.h
type SOCKADDR struct {
	sa_family ADDRESS_FAMILY
	sa_data [14]CHAR
}
