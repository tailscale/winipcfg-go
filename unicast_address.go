/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
)

type UnicastAddress struct {
	IpAdapterAddressCommonTypeEx

	PrefixOrigin IpPrefixOrigin
	SuffixOrigin IpSuffixOrigin
	DadState     IpDadState

	ValidLifetime      uint32
	PreferredLifetime  uint32
	LeaseLifetime      uint32
	OnLinkPrefixLength uint8
}

func (ua *UnicastAddress) String() string {
	if ua == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("%s/%d; PrefixOrigin: %s; SuffixOrigin: %s; DadState: %s; ValidLifetime: %d; PreferredLifetime: %d; LeaseLifetime: %d",
			ua.commonTypeExAddressString(), ua.OnLinkPrefixLength, ua.PrefixOrigin.String(), ua.SuffixOrigin.String(),
			ua.DadState.String(), ua.ValidLifetime, ua.PreferredLifetime, ua.LeaseLifetime)
	}
}
