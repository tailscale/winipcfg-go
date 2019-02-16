/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterAddressCommonType struct {
	// The interface the address belongs to.
	InterfaceLuid  uint64
	InterfaceIndex uint32

	// TODO: Documentation missing. What is it?
	Length uint32

	// The address.
	Address SockaddrInet
}

func ipAdapterAddressFromLengthAddress(ifc Interface, length uint32, wtsa *wtSocketAddress) (*IpAdapterAddressCommonType,
	error) {

	sainet, err := wtsa.toSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &IpAdapterAddressCommonType{
		InterfaceLuid:  ifc.Luid,
		InterfaceIndex: ifc.Index,
		Length:         length,
		Address:        *sainet}, nil
}

func (a *IpAdapterAddressCommonType) commonTypeAddressString() string {
	if a == nil {
		return ""
	} else {
		return fmt.Sprintf("Length: %d; Address: [%s]", a.Length, a.Address.String())
	}
}

func (a *IpAdapterAddressCommonType) String() string {
	return a.commonTypeAddressString()
}
