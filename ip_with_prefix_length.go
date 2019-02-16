/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

type IpWithPrefixLength struct {
	IP           net.IP
	PrefixLength uint8
}

func (ippl *IpWithPrefixLength) String() string {
	if ippl == nil {
		return ""
	} else {
		return fmt.Sprintf("%s/%d", ippl.IP.String(), ippl.PrefixLength)
	}
}
