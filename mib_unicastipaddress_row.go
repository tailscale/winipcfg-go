/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type MibUnicastipaddressRow struct {
	Address            SockaddrInet
	InterfaceLuid      uint64
	InterfaceIndex     uint32
	PrefixOrigin       NlPrefixOrigin
	SuffixOrigin       NlSuffixOrigin
	ValidLifetime      uint32
	PreferredLifetime  uint32
	OnLinkPrefixLength uint8
	SkipAsSource       bool
	DadState           NlDadState
	ScopeId            uint32
	CreationTimeStamp  int64
}

func (uar *MibUnicastipaddressRow) String() string {

	if uar == nil {
		return ""
	}

	return fmt.Sprintf(`Address: %s/%d
InterfaceLuid: %d
InterfaceIndex: %d
PrefixOrigin: %s
SuffixOrigin: %s
ValidLifetime: %d
PreferredLifetime: %d
SkipAsSource: %v
DadState: %s
ScopeId: %d
CreationTimeStamp: %d
`, uar.Address.String(), uar.OnLinkPrefixLength, uar.InterfaceLuid, uar.InterfaceIndex, uar.PrefixOrigin.String(),
		uar.SuffixOrigin.String(), uar.ValidLifetime, uar.PreferredLifetime, uar.SkipAsSource, uar.DadState.String(),
		uar.ScopeId, uar.CreationTimeStamp)
}
