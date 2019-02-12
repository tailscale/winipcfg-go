/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// Defined in iptypes.h
const (
	MAX_DNS_SUFFIX_STRING_LENGTH = 256
)

// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_dns_suffix
// IP_ADAPTER_DNS_SUFFIX defined in iptypes.h
type wtIpAdapterDnsSuffix struct {
	Next *wtIpAdapterDnsSuffix
	String [MAX_DNS_SUFFIX_STRING_LENGTH]WCHAR
}
