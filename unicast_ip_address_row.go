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

// Returns all unicast IP addresses assigned to any interface. Corresponds to GetUnicastIpAddressTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddresstable).
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

// Returns UnicastIpAddressRow struct that matches to provided 'ip' argument, or nil if no such unicast address is found
// on the system.
func GetMatchingUnicastIpAddressRow(ip *net.IP) (*UnicastIpAddressRow, error) {

	if ip == nil {
		return nil, fmt.Errorf("GetMatchingUnicastIpAddressRow() - input ip is nil")
	}

	row, err := getMatchingWtMibUnicastipaddressRow(ip)

	if err != nil {
		return nil, err
	} else if row == nil {
		return nil, nil
	} else {
		return row.toUnicastIpAddressRow()
	}
}

// Saves (activates) modified UnicastIpAddressRow. Corresponds to SetUnicastIpAddressEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setunicastipaddressentry).
//
// Note that fields Address, InterfaceLuid and InterfaceIndex are used for identifying address to change, meaning that
// they cannot be changed by using this method. Changing some of these fields would cause updating some other unicast IP
// address. On the other side, fields DadState, ScopeId and CreationTimeStamp are read-only, so they also cannot be
// changed. So fields that are "changeable" this way are: PrefixOrigin, SuffixOrigin, ValidLifetime, PreferredLifetime,
// OnLinkPrefixLength and SkipAsSource.
// The workflow of using this method is:
// 1) Get UnicastIpAddressRow instance by using any of getter methods (i.e. GetMatchingUnicastIpAddressRow or any other);
// 2) Change one or more of "changeable" fields enumerated above;
// 3) Calling this method to activate the changes.
func (address *UnicastIpAddressRow) Set() error {

	old, err := getWtMibUnicastipaddressRow(address.InterfaceLuid, &address.Address.Address)

	if err != nil {
		return err
	}

	old.PrefixOrigin = address.PrefixOrigin
	old.SuffixOrigin = address.SuffixOrigin
	old.ValidLifetime = address.ValidLifetime
	old.PreferredLifetime = address.PreferredLifetime
	old.OnLinkPrefixLength = address.OnLinkPrefixLength
	old.SkipAsSource = boolToUint8(address.SkipAsSource)

	return old.set()
}

// Deletes unicast IP address from the system.
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
