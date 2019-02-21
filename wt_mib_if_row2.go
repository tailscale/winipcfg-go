/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	if_max_string_size         = 256 // IF_MAX_STRING_SIZE defined in ifdef.h
	if_max_phys_address_length = 32  // IF_MAX_PHYS_ADDRESS_LENGTH defined in ifdef.h
)

type interfaceAndOperStatusFlagsByte uint8

const (
	hardwareInterface interfaceAndOperStatusFlagsByte = 1
	filterInterface   interfaceAndOperStatusFlagsByte = 2
	connectorPresent  interfaceAndOperStatusFlagsByte = 4
	notAuthenticated  interfaceAndOperStatusFlagsByte = 8
	notMediaConnected interfaceAndOperStatusFlagsByte = 16
	paused            interfaceAndOperStatusFlagsByte = 32
	lowPower          interfaceAndOperStatusFlagsByte = 64
	endPointInterface interfaceAndOperStatusFlagsByte = 128
)

func (wtior interfaceAndOperStatusFlagsByte) toInterfaceAndOperStatusFlags() *InterfaceAndOperStatusFlags {
	return &InterfaceAndOperStatusFlags{
		HardwareInterface: uint8ToBool(uint8(wtior) & uint8(hardwareInterface)),
		FilterInterface:   uint8ToBool(uint8(wtior) & uint8(filterInterface)),
		ConnectorPresent:  uint8ToBool(uint8(wtior) & uint8(connectorPresent)),
		NotAuthenticated:  uint8ToBool(uint8(wtior) & uint8(notAuthenticated)),
		NotMediaConnected: uint8ToBool(uint8(wtior) & uint8(notMediaConnected)),
		Paused:            uint8ToBool(uint8(wtior) & uint8(paused)),
		LowPower:          uint8ToBool(uint8(wtior) & uint8(lowPower)),
		EndPointInterface: uint8ToBool(uint8(wtior) & uint8(endPointInterface)),
	}
}
