/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// IP_ADAPTER_MULTICAST_ADDRESS defined in iptypes.h
type wtIpAdapterMulticastAddress wtIpAdapterMulticastAddressXp

func (wta *wtIpAdapterMulticastAddressXp) toIpAdapterAddress(ifc Interface) (*IpAdapterAddressCommonTypeEx, error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthFlagsAddress(ifc, wta.Length, wta.Flags, &wta.Address)
	}
}
