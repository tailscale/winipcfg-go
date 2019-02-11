/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

type UnicastAddress struct {
	Interface Interface

	// The rest is from IP_ADAPTER_ADDRESSES_LH

	// TODO: Documentation missing. What is it?
	Length uint32

	// TODO: Documentation missing. What is it?
	Flags uint32

	Address SOCKET_ADDRESS

	PrefixOrigin IP_PREFIX_ORIGIN
	SuffixOrigin IP_SUFFIX_ORIGIN
	DadState IP_DAD_STATE

	ValidLifetime ULONG
	PreferredLifetime ULONG
	LeaseLifetime ULONG
	OnLinkPrefixLength UINT8
}