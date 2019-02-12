/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterAddressCommonTypeEx struct {

	IpAdapterAddressCommonType

	// TODO: Documentation missing. What is it?
	Flags uint32
}

func ipAdapterAddressFromAnycastAddress(ifc Interface, aa *wtIpAdapterAnycastAddressXp) (*IpAdapterAddressCommonTypeEx,
	error) {
	if aa == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthFlagsAddress(ifc, aa.Length, aa.Flags, &aa.Address)
	}
}

func ipAdapterAddressFromMulticastAddress(ifc Interface, aa *wtIpAdapterMulticastAddressXp) (*IpAdapterAddressCommonTypeEx,
	error) {
	if aa == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthFlagsAddress(ifc, aa.Length, aa.Flags, &aa.Address)
	}
}

func ipAdapterAddressFromLengthFlagsAddress(ifc Interface, length uint32, flags uint32, wtsa *wtSocketAddress) (*IpAdapterAddressCommonTypeEx,
	error) {

	act := IpAdapterAddressCommonTypeEx{}

	act.Interface = ifc
	act.Length = length
	act.Flags = flags

	err := act.setAddress(wtsa)

	if err == nil {
		return &act, nil
	} else {
		return nil, err
	}
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
