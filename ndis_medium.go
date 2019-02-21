/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NDIS_MEDIUM defined in ntddndis.h
// (https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/ntddndis/ne-ntddndis-_ndis_medium)
type NdisMedium uint32

const (
	NdisMedium802_3        NdisMedium = 0
	NdisMedium802_5        NdisMedium = 1
	NdisMediumFddi         NdisMedium = 2
	NdisMediumWan          NdisMedium = 3
	NdisMediumLocalTalk    NdisMedium = 4
	NdisMediumDix          NdisMedium = 5 // defined for convenience, not a real medium
	NdisMediumArcnetRaw    NdisMedium = 6
	NdisMediumArcnet878_2  NdisMedium = 7
	NdisMediumAtm          NdisMedium = 8
	NdisMediumWirelessWan  NdisMedium = 9
	NdisMediumIrda         NdisMedium = 10
	NdisMediumBpc          NdisMedium = 11
	NdisMediumCoWan        NdisMedium = 12
	NdisMedium1394         NdisMedium = 13
	NdisMediumInfiniBand   NdisMedium = 14
	NdisMediumTunnel       NdisMedium = 15
	NdisMediumNative802_11 NdisMedium = 16
	NdisMediumLoopback     NdisMedium = 17
	NdisMediumWiMAX        NdisMedium = 18
	NdisMediumIP           NdisMedium = 19
	NdisMediumMax          NdisMedium = 20
)

func (nm NdisMedium) String() string {
	switch nm {
	case NdisMedium802_3:
		return "NdisMedium802_3"
	case NdisMedium802_5:
		return "NdisMedium802_5"
	case NdisMediumFddi:
		return "NdisMediumFddi"
	case NdisMediumWan:
		return "NdisMediumWan"
	case NdisMediumLocalTalk:
		return "NdisMediumLocalTalk"
	case NdisMediumDix:
		return "NdisMediumDix"
	case NdisMediumArcnetRaw:
		return "NdisMediumArcnetRaw"
	case NdisMediumArcnet878_2:
		return "NdisMediumArcnet878_2"
	case NdisMediumAtm:
		return "NdisMediumAtm"
	case NdisMediumWirelessWan:
		return "NdisMediumWirelessWan"
	case NdisMediumIrda:
		return "NdisMediumIrda"
	case NdisMediumBpc:
		return "NdisMediumBpc"
	case NdisMediumCoWan:
		return "NdisMediumCoWan"
	case NdisMedium1394:
		return "NdisMedium1394"
	case NdisMediumInfiniBand:
		return "NdisMediumInfiniBand"
	case NdisMediumTunnel:
		return "NdisMediumTunnel"
	case NdisMediumNative802_11:
		return "NdisMediumNative802_11"
	case NdisMediumLoopback:
		return "NdisMediumLoopback"
	case NdisMediumWiMAX:
		return "NdisMediumWiMAX"
	case NdisMediumIP:
		return "NdisMediumIP"
	case NdisMediumMax:
		return "NdisMediumMax"
	default:
		return fmt.Sprintf("NdisMedium_UNKNOWN(%d)", nm)
	}
}
