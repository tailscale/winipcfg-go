/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// IP_ADAPTER_WINS_SERVER_ADDRESS defined in iptypes.h
type wtIpAdapterWinsServerAddress wtIpAdapterWinsServerAddressLh

func (wta *wtIpAdapterWinsServerAddressLh) toIpAdapterAddress(ifc Interface) (*IpAdapterAddressCommonType, error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthAddress(ifc, wta.Length, &wta.Address)
	}
}
