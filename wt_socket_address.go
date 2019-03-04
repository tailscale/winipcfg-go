/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ns-ws2def-_socket_address
// SOCKET_ADDRESS defined in ws2def.h
type wtSocketAddress struct {
	lpSockaddr      *wtSockaddr
	iSockaddrLength int32 // Windows type: INT
}

// https://docs.microsoft.com/en-us/windows/desktop/WinSock/sockaddr-2
// SOCKADDR defined in ws2def.h
type wtSockaddr struct {
	sa_family AddressFamily
	sa_data   [14]uint8 // Windows type: [14]CHAR
}

func (sa *wtSocketAddress) getWtSockaddrInet() (*wtSockaddrInet, error) {

	if sa.lpSockaddr == nil {
		return nil, nil
	}

	if sa.lpSockaddr.sa_family != AF_INET && sa.lpSockaddr.sa_family != AF_INET6 {
		return nil, fmt.Errorf("getWtSockaddrInet() - receiver's argument family has to be AF_INET or AF_INET6")
	}

	return (*wtSockaddrInet)(unsafe.Pointer(sa.lpSockaddr)), nil
}

func (wtsa *wtSocketAddress) toSockaddrInet() (*SockaddrInet, error) {

	if wtsa == nil {
		return nil, nil
	}

	wtsainet, err := wtsa.getWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	sainet, err := wtsainet.toSockaddrInet()

	if err == nil {
		return sainet, nil
	} else {
		return nil, err
	}
}
