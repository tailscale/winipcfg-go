/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NET_IF_ADMIN_STATUS defined in ifdef.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-net_if_admin_status)
type NetIfAdminStatus uint32

const (
	NET_IF_ADMIN_STATUS_UP      NetIfAdminStatus = 1
	NET_IF_ADMIN_STATUS_DOWN    NetIfAdminStatus = 2
	NET_IF_ADMIN_STATUS_TESTING NetIfAdminStatus = 3
)

func (nias NetIfAdminStatus) String() string {
	switch nias {
	case NET_IF_ADMIN_STATUS_UP:
		return "NET_IF_ADMIN_STATUS_UP"
	case NET_IF_ADMIN_STATUS_DOWN:
		return "NET_IF_ADMIN_STATUS_DOWN"
	case NET_IF_ADMIN_STATUS_TESTING:
		return "NET_IF_ADMIN_STATUS_TESTING"
	default:
		return fmt.Sprintf("NetIfAdminStatus_UNKNOWN(%d)", nias)
	}
}
