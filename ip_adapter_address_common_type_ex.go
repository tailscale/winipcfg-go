/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterAddressCommonTypeEx struct {
	// It extends IpAdapterAddressCommonType
	IpAdapterAddressCommonType

	// TODO: Documentation missing. What is it?
	Flags uint32
}

func ipAdapterAddressFromLengthFlagsAddress(ifc Interface, length uint32, flags uint32, wtsa *wtSocketAddress) (*IpAdapterAddressCommonTypeEx,
	error) {

	sainet, err := wtsa.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	act := IpAdapterAddressCommonTypeEx{Flags: flags}

	act.InterfaceLuid = ifc.Luid
	act.InterfaceIndex = ifc.Index
	act.Length = length
	act.Address = *sainet

	return &act, nil
}

func (a *IpAdapterAddressCommonTypeEx) commonTypeExAddressString() string {

	if a == nil {
		return ""
	} else {
		return fmt.Sprintf("Flags: %d; %s", a.Flags, a.commonTypeAddressString())
	}
}

func (a *IpAdapterAddressCommonTypeEx) String() string {
	return a.commonTypeExAddressString()
}
