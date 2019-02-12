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
	lpSockaddr *wtSockaddr
	iSockaddrLength int32 // Windows type: INT
}

// https://docs.microsoft.com/en-us/windows/desktop/WinSock/sockaddr-2
// SOCKADDR defined in ws2def.h
type wtSockaddr struct {
	sa_family AddressFamily
	sa_data   [14]uint8 // Windows type: [14]CHAR
}

func (sa *wtSocketAddress) getWtSockaddrInet() (*wtSockaddrInet, error) {

	if sa == nil {
		return nil, nil
	}

	if sa.lpSockaddr == nil {
		// TODO: The following check is unnecessary. Only for curiosity...
		if sa.iSockaddrLength == 0 {
			return nil, nil
		} else {
			return nil, fmt.Errorf(
				"Should not happen ever! wtSocketAddress.lpSockaddr is nil, although wtSocketAddress.iSockaddrLength is %d (0 expected).",
				sa.iSockaddrLength)
		}
	}

	switch sa.lpSockaddr.sa_family {

	case AF_INET:

		// TODO: Remove this check once it's confirmed that it works OK.
		if sa.iSockaddrLength != wtSockaddrIn_Size {
			return nil,
				fmt.Errorf("wtSocketAddress.lpSockaddr.sa_family is %s, but wtSocketAddress.iSockaddrLength is %d (%d expected).",
					AF_INET.String(), sa.iSockaddrLength, wtSockaddrIn_Size)
		}

		break

	case AF_INET6:

		// TODO: Remove this check once it's confirmed that it works OK.
		if sa.iSockaddrLength != wtSockaddrIn6Lh_Size {
			return nil,
				fmt.Errorf("wtSocketAddress.lpSockaddr.sa_family is %s, but wtSocketAddress.iSockaddrLength is %d (%d expected).",
					AF_INET6.String(), sa.iSockaddrLength, wtSockaddrIn6Lh_Size)
		}

		break

	default:
		return nil, fmt.Errorf("Input argument cannot be converted to wtSockaddrInet because its family is %s.",
			sa.lpSockaddr.sa_family.String())
	}

	return (*wtSockaddrInet)(unsafe.Pointer(sa.lpSockaddr)), nil
}
