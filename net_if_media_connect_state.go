/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NET_IF_MEDIA_CONNECT_STATE defined in ifdef.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_media_connect_state)
type NetIfMediaConnectState uint32

const (
	MediaConnectStateUnknown      NetIfMediaConnectState = 0
	MediaConnectStateConnected    NetIfMediaConnectState = 1
	MediaConnectStateDisconnected NetIfMediaConnectState = 2
)

func (nimcs NetIfMediaConnectState) String() string {
	switch nimcs {
	case MediaConnectStateUnknown:
		return "MediaConnectStateUnknown"
	case MediaConnectStateConnected:
		return "MediaConnectStateConnected"
	case MediaConnectStateDisconnected:
		return "MediaConnectStateDisconnected"
	default:
		return fmt.Sprintf("NetIfMediaConnectState_UNKNOWN(%d)", nimcs)
	}
}
