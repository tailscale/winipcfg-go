/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// Corresponds to MIB_UNICASTIPADDRESS_ROW defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_unicastipaddress_row)
type UnicastIpAddressRow struct {
	Address            *SockaddrInet
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

func (address *UnicastIpAddressRow) equal(other *UnicastIpAddressRow) bool {

	if address == nil || other == nil {
		return false
	}

	return address.InterfaceLuid == other.InterfaceLuid && address.InterfaceIndex == other.InterfaceIndex &&
		address.PrefixOrigin == other.PrefixOrigin && address.SuffixOrigin == other.SuffixOrigin &&
		address.ValidLifetime == other.ValidLifetime && address.PreferredLifetime == other.PreferredLifetime &&
		address.OnLinkPrefixLength == other.OnLinkPrefixLength && address.SkipAsSource == other.SkipAsSource &&
		address.DadState == other.DadState && address.ScopeId == other.ScopeId &&
		address.CreationTimeStamp == other.CreationTimeStamp && address.Address.equal(other.Address)
}

func (address *UnicastIpAddressRow) toWtMibUnicastipaddressRow() (*wtMibUnicastipaddressRow, error) {

	if address == nil {
		return nil, nil
	}

	wtsai, err := address.Address.toWtSockaddrInet()

	if err != nil {
		return nil, err
	}

	return &wtMibUnicastipaddressRow{
		Address:            *wtsai,
		InterfaceLuid:      address.InterfaceLuid,
		InterfaceIndex:     address.InterfaceIndex,
		PrefixOrigin:       address.PrefixOrigin,
		SuffixOrigin:       address.SuffixOrigin,
		ValidLifetime:      address.ValidLifetime,
		PreferredLifetime:  address.PreferredLifetime,
		OnLinkPrefixLength: address.OnLinkPrefixLength,
		SkipAsSource:       boolToUint8(address.SkipAsSource),
		DadState:           address.DadState,
		ScopeId:            address.ScopeId,
		CreationTimeStamp:  address.CreationTimeStamp,
	}, nil
}

// Corresponds to GetUnicastIpAddressTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddresstable)
func GetUnicastAddresses(family AddressFamily) ([]*UnicastIpAddressRow, error) {

	wtas, err := getWtMibUnicastipaddressRows(family)

	if err != nil {
		return nil, err
	}

	count := len(wtas)

	addresses := make([]*UnicastIpAddressRow, count, count)

	for idx, wta := range wtas {

		address, err := wta.toUnicastIpAddressRow()

		if err != nil {
			return nil, err
		}

		addresses[idx] = address
	}

	return addresses, nil
}

func GetMatchingUnicastIpAddressRow(ip *net.IP) (*UnicastIpAddressRow, error) {

	if ip == nil {
		return nil, fmt.Errorf("GetMatchingUnicastIpAddressRow() - input ip is nil")
	}

	row, err := getMatchingWtMibUnicastipaddressRow(nil, ip)

	if err != nil {
		return nil, err
	}

	if row == nil {
		return nil, nil
	}

	uad, err := row.toUnicastIpAddressRow()

	if err != nil {
		return nil, err
	}

	return uad, nil
}

func (address *UnicastIpAddressRow) Delete() error {

	if address.Address == nil {
		return fmt.Errorf("UnicastIpAddressRow.Delete() - receiver argument or its Address field is nil")
	}

	wta, err := address.toWtMibUnicastipaddressRow()

	if err != nil {
		return err
	}

	rows, err := getWtMibUnicastipaddressRows(address.Address.Family)

	if err != nil {
		return err
	}

	for _, row := range rows {
		if row.equal(wta) {
			return row.delete()
		}
	}

	return fmt.Errorf("UnicastIpAddressRow.Delete() - address not found")
}

func (address *UnicastIpAddressRow) String() string {

	if address == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`Address: [%s]/%d
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
`, address.Address.String(), address.OnLinkPrefixLength, address.InterfaceLuid, address.InterfaceIndex, address.PrefixOrigin.String(),
		address.SuffixOrigin.String(), address.ValidLifetime, address.PreferredLifetime, address.SkipAsSource, address.DadState.String(),
		address.ScopeId, address.CreationTimeStamp)
}
