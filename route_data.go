/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "net"

type RouteData struct {
	Destination net.IPNet
	NextHop     net.IP
	Metric      uint32
}
