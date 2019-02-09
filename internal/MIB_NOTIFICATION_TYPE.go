/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ne-netioapi-_mib_notification_type
type MIB_NOTIFICATION_TYPE uint32

const (
	//
	// ParameterChange.
	//
	MibParameterNotification MIB_NOTIFICATION_TYPE = 0
	//
	// Addition.
	//
	MibAddInstance           MIB_NOTIFICATION_TYPE = 1
	//
	// Deletion.
	//
	MibDeleteInstance        MIB_NOTIFICATION_TYPE = 2
	//
	// Initial notification.
	//
	MibInitialNotification   MIB_NOTIFICATION_TYPE = 3
)

func (mnt MIB_NOTIFICATION_TYPE) String() string {
	switch mnt {
	case MibParameterNotification:
		return "MibParameterNotification"
	case MibAddInstance:
		return "MibAddInstance"
	case MibDeleteInstance:
		return "MibDeleteInstance"
	case MibInitialNotification:
		return "MibInitialNotification"
	default:
		return fmt.Sprintf("MIB_NOTIFICATION_TYPE_UNKNOWN(%d)", mnt)
	}
}
