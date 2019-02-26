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

	result := interfaceAndOperStatusFlagsByte(0)

	if iaosf.HardwareInterface {
		result |= hardwareInterface
	}

	if iaosf.FilterInterface {
		result |= filterInterface
	}

	if iaosf.ConnectorPresent {
		result |= connectorPresent
	}

	if iaosf.NotAuthenticated {
		result |= notAuthenticated
	}

	if iaosf.NotMediaConnected {
		result |= notMediaConnected
	}

	if iaosf.Paused {
		result |= paused
	}

	if iaosf.LowPower {
		result |= lowPower
	}

	if iaosf.EndPointInterface {
		result |= endPointInterface
	}

	return result
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
EndPointInterface: %v`, iaosf.HardwareInterface, iaosf.FilterInterface, iaosf.ConnectorPresent, iaosf.NotAuthenticated,
		iaosf.NotMediaConnected, iaosf.Paused, iaosf.LowPower, iaosf.EndPointInterface)
}
