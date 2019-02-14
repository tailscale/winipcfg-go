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

func (ap *IpAdapterPrefix) String() string {
	if ap == nil {
		return ""
	} else {
		return fmt.Sprintf("%s/%d", ap.commonTypeExAddressString(), ap.PrefixLength)
	}
}
