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

func ipAdapterAddressFromWtAnycastAddress(ifc Interface, aa *wtIpAdapterAnycastAddressXp) (*IpAdapterAddressCommonTypeEx,
	error) {
	if aa == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthFlagsAddress(ifc, aa.Length, aa.Flags, &aa.Address)
	}
}

func ipAdapterAddressFromWtMulticastAddress(ifc Interface, aa *wtIpAdapterMulticastAddressXp) (*IpAdapterAddressCommonTypeEx,
	error) {
	if aa == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthFlagsAddress(ifc, aa.Length, aa.Flags, &aa.Address)
	}
}

func ipAdapterAddressFromLengthFlagsAddress(ifc Interface, length uint32, flags uint32, wtsa *wtSocketAddress) (*IpAdapterAddressCommonTypeEx,
	error) {

	sainet, err := sockaddrInetFromWtSocketAddress(wtsa)

	if err != nil {
		return nil, err
	}

	act := IpAdapterAddressCommonTypeEx{Flags: flags}

	act.Interface = ifc
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
