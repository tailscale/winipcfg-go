/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// IP_ADAPTER_GATEWAY_ADDRESS defined in iptypes.h
type wtIpAdapterGatewayAddress wtIpAdapterGatewayAddressLh

func (wta *wtIpAdapterGatewayAddressLh) toIpAdapterAddress(ifc Interface) (*IpAdapterAddressCommonType, error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthAddress(ifc, wta.Length, &wta.Address)
	}
}
