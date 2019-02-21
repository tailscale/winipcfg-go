/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// NDIS_PHYSICAL_MEDIUM defined in ntddndis.h
//(https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_row2)
type NdisPhysicalMedium uint32

const (
	NdisPhysicalMediumUnspecified    NdisPhysicalMedium = 0
	NdisPhysicalMediumWirelessLan    NdisPhysicalMedium = 1
	NdisPhysicalMediumCableModem     NdisPhysicalMedium = 2
	NdisPhysicalMediumPhoneLine      NdisPhysicalMedium = 3
	NdisPhysicalMediumPowerLine      NdisPhysicalMedium = 4
	NdisPhysicalMediumDSL            NdisPhysicalMedium = 5 // includes ADSL and UADSL (G.Lite)
	NdisPhysicalMediumFibreChannel   NdisPhysicalMedium = 6
	NdisPhysicalMedium1394           NdisPhysicalMedium = 7
	NdisPhysicalMediumWirelessWan    NdisPhysicalMedium = 8
	NdisPhysicalMediumNative802_11   NdisPhysicalMedium = 9
	NdisPhysicalMediumBluetooth      NdisPhysicalMedium = 10
	NdisPhysicalMediumInfiniband     NdisPhysicalMedium = 11
	NdisPhysicalMediumWiMax          NdisPhysicalMedium = 12
	NdisPhysicalMediumUWB            NdisPhysicalMedium = 13
	NdisPhysicalMedium802_3          NdisPhysicalMedium = 14
	NdisPhysicalMedium802_5          NdisPhysicalMedium = 15
	NdisPhysicalMediumIrda           NdisPhysicalMedium = 16
	NdisPhysicalMediumWiredWAN       NdisPhysicalMedium = 17
	NdisPhysicalMediumWiredCoWan     NdisPhysicalMedium = 18
	NdisPhysicalMediumOther          NdisPhysicalMedium = 19
	NdisPhysicalMediumNative802_15_4 NdisPhysicalMedium = 20
	NdisPhysicalMediumMax            NdisPhysicalMedium = 21
)

func (npm NdisPhysicalMedium) String() string {
	switch npm {
	case NdisPhysicalMediumUnspecified:
		return "NdisPhysicalMediumUnspecified"
	case NdisPhysicalMediumWirelessLan:
		return "NdisPhysicalMediumWirelessLan"
	case NdisPhysicalMediumCableModem:
		return "NdisPhysicalMediumCableModem"
	case NdisPhysicalMediumPhoneLine:
		return "NdisPhysicalMediumPhoneLine"
	case NdisPhysicalMediumPowerLine:
		return "NdisPhysicalMediumPowerLine"
	case NdisPhysicalMediumDSL:
		return "NdisPhysicalMediumDSL"
	case NdisPhysicalMediumFibreChannel:
		return "NdisPhysicalMediumFibreChannel"
	case NdisPhysicalMedium1394:
		return "NdisPhysicalMedium1394"
	case NdisPhysicalMediumWirelessWan:
		return "NdisPhysicalMediumWirelessWan"
	case NdisPhysicalMediumNative802_11:
		return "NdisPhysicalMediumNative802_11"
	case NdisPhysicalMediumBluetooth:
		return "NdisPhysicalMediumBluetooth"
	case NdisPhysicalMediumInfiniband:
		return "NdisPhysicalMediumInfiniband"
	case NdisPhysicalMediumWiMax:
		return "NdisPhysicalMediumWiMax"
	case NdisPhysicalMediumUWB:
		return "NdisPhysicalMediumUWB"
	case NdisPhysicalMedium802_3:
		return "NdisPhysicalMedium802_3"
	case NdisPhysicalMedium802_5:
		return "NdisPhysicalMedium802_5"
	case NdisPhysicalMediumIrda:
		return "NdisPhysicalMediumIrda"
	case NdisPhysicalMediumWiredWAN:
		return "NdisPhysicalMediumWiredWAN"
	case NdisPhysicalMediumWiredCoWan:
		return "NdisPhysicalMediumWiredCoWan"
	case NdisPhysicalMediumOther:
		return "NdisPhysicalMediumOther"
	case NdisPhysicalMediumNative802_15_4:
		return "NdisPhysicalMediumNative802_15_4"
	case NdisPhysicalMediumMax:
		return "NdisPhysicalMediumMax"
	default:
		return fmt.Sprintf("NdisPhysicalMedium_UNKNOWN(%d)", npm)
	}
}
