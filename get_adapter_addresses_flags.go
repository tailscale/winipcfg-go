/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// Determines the amount of data we want to get by GetInterfaces() function and similar functions which are returning
// Interface struct(s).
type GetAdapterAddressesFlags struct {
	GAA_FLAG_SKIP_UNICAST                bool
	GAA_FLAG_SKIP_ANYCAST                bool
	GAA_FLAG_SKIP_MULTICAST              bool
	GAA_FLAG_SKIP_DNS_SERVER             bool
	GAA_FLAG_INCLUDE_PREFIX              bool
	GAA_FLAG_SKIP_FRIENDLY_NAME          bool
	GAA_FLAG_INCLUDE_WINS_INFO           bool
	GAA_FLAG_INCLUDE_GATEWAYS            bool
	GAA_FLAG_INCLUDE_ALL_INTERFACES      bool
	GAA_FLAG_INCLUDE_ALL_COMPARTMENTS    bool
	GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER bool
	GAA_FLAG_SKIP_DNS_INFO               bool
}

// Returns DEFAULT GetAdapterAddressesFlags struct. In DEFAULT struct all fields are set to false, meaning that nothing
// is explicitly included, nor explicitly skipped.
func DefaultGetAdapterAddressesFlags() *GetAdapterAddressesFlags {
	return &GetAdapterAddressesFlags{
		GAA_FLAG_SKIP_UNICAST:                false,
		GAA_FLAG_SKIP_ANYCAST:                false,
		GAA_FLAG_SKIP_MULTICAST:              false,
		GAA_FLAG_SKIP_DNS_SERVER:             false,
		GAA_FLAG_INCLUDE_PREFIX:              false,
		GAA_FLAG_SKIP_FRIENDLY_NAME:          false,
		GAA_FLAG_INCLUDE_WINS_INFO:           false,
		GAA_FLAG_INCLUDE_GATEWAYS:            false,
		GAA_FLAG_INCLUDE_ALL_INTERFACES:      false,
		GAA_FLAG_INCLUDE_ALL_COMPARTMENTS:    false,
		GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER: false,
		GAA_FLAG_SKIP_DNS_INFO:               false,
	}
}

// Returns MIN GetAdapterAddressesFlags struct. In MIN struct all "skip" fields are set to true (meaning that everything
// that can be explicitly skipped is skipped), and all "include" fields are set to false (meaning that nothing that can
// be explicitly included is included).
func MinGetAdapterAddressesFlags() *GetAdapterAddressesFlags {
	return &GetAdapterAddressesFlags{
		GAA_FLAG_SKIP_UNICAST:                true,
		GAA_FLAG_SKIP_ANYCAST:                true,
		GAA_FLAG_SKIP_MULTICAST:              true,
		GAA_FLAG_SKIP_DNS_SERVER:             true,
		GAA_FLAG_INCLUDE_PREFIX:              false,
		GAA_FLAG_SKIP_FRIENDLY_NAME:          true,
		GAA_FLAG_INCLUDE_WINS_INFO:           false,
		GAA_FLAG_INCLUDE_GATEWAYS:            false,
		GAA_FLAG_INCLUDE_ALL_INTERFACES:      false,
		GAA_FLAG_INCLUDE_ALL_COMPARTMENTS:    false,
		GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER: false,
		GAA_FLAG_SKIP_DNS_INFO:               true,
	}
}

// Returns FULL GetAdapterAddressesFlags struct. In FULL struct all "skip" fields are set to false (meaning that nothing
// that can be explicitly skipped is skipped), and all "include" fields are set to true (meaning that everything that
// can be explicitly included is included).
func FullGetAdapterAddressesFlags() *GetAdapterAddressesFlags {
	return &GetAdapterAddressesFlags{
		GAA_FLAG_SKIP_UNICAST:                false,
		GAA_FLAG_SKIP_ANYCAST:                false,
		GAA_FLAG_SKIP_MULTICAST:              false,
		GAA_FLAG_SKIP_DNS_SERVER:             false,
		GAA_FLAG_INCLUDE_PREFIX:              true,
		GAA_FLAG_SKIP_FRIENDLY_NAME:          false,
		GAA_FLAG_INCLUDE_WINS_INFO:           true,
		GAA_FLAG_INCLUDE_GATEWAYS:            true,
		GAA_FLAG_INCLUDE_ALL_INTERFACES:      true,
		GAA_FLAG_INCLUDE_ALL_COMPARTMENTS:    true,
		GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER: true,
		GAA_FLAG_SKIP_DNS_INFO:               false,
	}
}

func (gaa *GetAdapterAddressesFlags) toGetAdapterAddressesFlagsBytes() getAdapterAddressesFlagsBytes {

	result := getAdapterAddressesFlagsBytes(0)

	if gaa.GAA_FLAG_SKIP_UNICAST {
		result |= gaa_flag_skip_unicast
	}

	if gaa.GAA_FLAG_SKIP_ANYCAST {
		result |= gaa_flag_skip_anycast
	}

	if gaa.GAA_FLAG_SKIP_MULTICAST {
		result |= gaa_flag_skip_multicast
	}

	if gaa.GAA_FLAG_SKIP_DNS_SERVER {
		result |= gaa_flag_skip_dns_server
	}

	if gaa.GAA_FLAG_INCLUDE_PREFIX {
		result |= gaa_flag_include_prefix
	}

	if gaa.GAA_FLAG_SKIP_FRIENDLY_NAME {
		result |= gaa_flag_skip_friendly_name
	}

	if gaa.GAA_FLAG_INCLUDE_WINS_INFO {
		result |= gaa_flag_include_wins_info
	}

	if gaa.GAA_FLAG_INCLUDE_GATEWAYS {
		result |= gaa_flag_include_gateways
	}

	if gaa.GAA_FLAG_INCLUDE_ALL_INTERFACES {
		result |= gaa_flag_include_all_interfaces
	}

	if gaa.GAA_FLAG_INCLUDE_ALL_COMPARTMENTS {
		result |= gaa_flag_include_all_compartments
	}

	if gaa.GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER {
		result |= gaa_flag_include_tunnel_bindingorder
	}

	if gaa.GAA_FLAG_SKIP_DNS_INFO {
		result |= gaa_flag_skip_dns_info
	}

	return result
}

func (gaa *GetAdapterAddressesFlags) String() string {

	if gaa == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`GAA_FLAG_SKIP_UNICAST: %v
GAA_FLAG_SKIP_ANYCAST: %v
GAA_FLAG_SKIP_MULTICAST: %v
GAA_FLAG_SKIP_DNS_SERVER: %v
GAA_FLAG_INCLUDE_PREFIX: %v
GAA_FLAG_SKIP_FRIENDLY_NAME: %v
GAA_FLAG_INCLUDE_WINS_INFO: %v
GAA_FLAG_INCLUDE_GATEWAYS: %v
GAA_FLAG_INCLUDE_ALL_INTERFACES: %v
GAA_FLAG_INCLUDE_ALL_COMPARTMENTS: %v
GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER: %v
GAA_FLAG_SKIP_DNS_INFO: %v`, gaa.GAA_FLAG_SKIP_UNICAST, gaa.GAA_FLAG_SKIP_ANYCAST, gaa.GAA_FLAG_SKIP_MULTICAST,
		gaa.GAA_FLAG_SKIP_DNS_SERVER, gaa.GAA_FLAG_INCLUDE_PREFIX, gaa.GAA_FLAG_SKIP_FRIENDLY_NAME,
		gaa.GAA_FLAG_INCLUDE_WINS_INFO, gaa.GAA_FLAG_INCLUDE_GATEWAYS, gaa.GAA_FLAG_INCLUDE_ALL_INTERFACES,
		gaa.GAA_FLAG_INCLUDE_ALL_COMPARTMENTS, gaa.GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER, gaa.GAA_FLAG_SKIP_DNS_INFO)
}
