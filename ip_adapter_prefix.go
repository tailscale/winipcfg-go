/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterPrefix struct {

	// It extends IpAdapterAddressCommonTypeEx
	IpAdapterAddressCommonTypeEx

	// Prefix length.
	PrefixLength uint32
}

func ipAdapterPrefixFromWinType(ifc Interface, wt *wtIpAdapterPrefixXp) (*IpAdapterPrefix, error) {

	if wt == nil {
		return nil, nil
	}

	sainet, err := sockaddrInetFromWtSocketAddress(&wt.Address)

	if err != nil {
		return nil, err
	}

	ap := IpAdapterPrefix{PrefixLength: wt.PrefixLength}

	ap.Interface = ifc
	ap.Length = wt.Length
	ap.Address = *sainet
	ap.Flags = wt.Flags

	return &ap, nil
}

func (ap *IpAdapterPrefix) String() string {
	return fmt.Sprintf("%s/%d", ap.commonTypeExAddressString(), ap.PrefixLength)
}
