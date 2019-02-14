/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAddressPrefix struct {
	Prefix       SockaddrInet
	PrefixLength uint8
}

func (ap *IpAddressPrefix) String() string {
	if ap == nil {
		return ""
	} else {
		return fmt.Sprintf("%s/%d", ap.Prefix.String(), ap.PrefixLength)
	}
}
