/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ne-netioapi-_mib_notification_type
// MIB_NOTIFICATION_TYPE defined in netioapi.h
type MibNotificationType uint32

const (
	//
	// ParameterChange.
	//
	MibParameterNotification MibNotificationType = 0
	//
	// Addition.
	//
	MibAddInstance MibNotificationType = 1
	//
	// Deletion.
	//
	MibDeleteInstance MibNotificationType = 2
	//
	// Initial notification.
	//
	MibInitialNotification MibNotificationType = 3
)

func (mnt MibNotificationType) String() string {
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
		return fmt.Sprintf("MibNotificationType_UNKNOWN(%d)", mnt)
	}
}
