/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type InterfaceAndOperStatusFlags struct {
	HardwareInterface bool
	FilterInterface   bool
	ConnectorPresent  bool
	NotAuthenticated  bool
	NotMediaConnected bool
	Paused            bool
	LowPower          bool
	EndPointInterface bool
}

func (iaosf *InterfaceAndOperStatusFlags) toInterfaceAndOperStatusFlagsByte() interfaceAndOperStatusFlagsByte {

	if iaosf == nil {
		panic("toInterfaceAndOperStatusFlagsByte() - receiver argument is nil")
	}

	result := uint8(0)

	if iaosf.HardwareInterface {
		result |= uint8(hardwareInterface)
	}

	if iaosf.FilterInterface {
		result |= uint8(filterInterface)
	}

	if iaosf.ConnectorPresent {
		result |= uint8(connectorPresent)
	}

	if iaosf.NotAuthenticated {
		result |= uint8(notAuthenticated)
	}

	if iaosf.NotMediaConnected {
		result |= uint8(notMediaConnected)
	}

	if iaosf.Paused {
		result |= uint8(paused)
	}

	if iaosf.LowPower {
		result |= uint8(lowPower)
	}

	if iaosf.EndPointInterface {
		result |= uint8(endPointInterface)
	}

	return interfaceAndOperStatusFlagsByte(result)
}

func (iaosf *InterfaceAndOperStatusFlags) String() string {

	if iaosf == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`HardwareInterface: %v
FilterInterface: %v
ConnectorPresent: %v
NotAuthenticated: %v
NotMediaConnected: %v
Paused: %v
LowPower: %v
EndPointInterface: %v
`, iaosf.HardwareInterface, iaosf.FilterInterface, iaosf.ConnectorPresent, iaosf.NotAuthenticated,
		iaosf.NotMediaConnected, iaosf.Paused, iaosf.LowPower, iaosf.EndPointInterface)
}
