/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

const (
	newAddressPrefixOrigin      = IpPrefixOriginManual
	newAddressSuffixOrigin      = IpSuffixOriginManual
	newAddressValidLifetime     = 4294967295
	newAddressPreferredLifetime = 4294967295
	newAddressSkipAsSource      = false
)

type UnicastAddressData struct {
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

func (address *UnicastAddressData) equivalentTo(other *UnicastAddressData) bool {

	if address == nil || other == nil {
		return false
	}

	return address.InterfaceLuid == other.InterfaceLuid && address.InterfaceIndex == other.InterfaceIndex &&
		address.PrefixOrigin == other.PrefixOrigin && address.SuffixOrigin == other.SuffixOrigin &&
		address.ValidLifetime == other.ValidLifetime && address.PreferredLifetime == other.PreferredLifetime &&
		address.OnLinkPrefixLength == other.OnLinkPrefixLength && address.SkipAsSource == other.SkipAsSource &&
		address.DadState == other.DadState && address.ScopeId == other.ScopeId &&
		address.CreationTimeStamp == other.CreationTimeStamp && address.Address.equivalentTo(other.Address)
}

func createUnicastAddressData(ifc *Interface, ipnet *net.IPNet) (*UnicastAddressData, error) {

	if ifc == nil {
		return nil, fmt.Errorf("createUnicastAddressData() - input Interface is nil")
	}

	if ipnet == nil {
		return nil, fmt.Errorf("createUnicastAddressData() - input IpWithPrefixLength is nil")
	}

	sainet, err := createSockaddrInet(ipnet.IP)

	if err != nil {
		return nil, err
	}

	ones, _ := ipnet.Mask.Size()

	// TODO: Check field values set here.
	return &UnicastAddressData{
		Address:            sainet,
		InterfaceLuid:      ifc.Luid,
		InterfaceIndex:     ifc.Index,
		PrefixOrigin:       newAddressPrefixOrigin,
		SuffixOrigin:       newAddressSuffixOrigin,
		ValidLifetime:      newAddressValidLifetime,
		PreferredLifetime:  newAddressPreferredLifetime,
		OnLinkPrefixLength: uint8(ones),
		SkipAsSource:       newAddressSkipAsSource,
	}, nil
}

func (address *UnicastAddressData) toWtMibUnicastipaddressRow() (*wtMibUnicastipaddressRow, error) {

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

func GetUnicastAddresses(family AddressFamily) ([]*UnicastAddressData, error) {

	wtas, err := getWtMibUnicastipaddressRows(family)

	if err != nil {
		return nil, err
	}

	if wtas == nil {
		return nil, nil
	}

	count := len(wtas)

	addresses := make([]*UnicastAddressData, count, count)

	for idx, wta := range wtas {

		address, err := wta.toUnicastAddressData()

		if err != nil {
			return nil, err
		}

		addresses[idx] = address
	}

	return addresses, nil
}

func (address *UnicastAddressData) Add() error {

	if address == nil {
		return fmt.Errorf("UnicastAddressData.Add() - input argument is nil")
	}

	wta, err := address.toWtMibUnicastipaddressRow()

	if err != nil {
		return err
	}

	err = wta.add()

	if err != nil {
		return err
	}

	// TODO: Not sure if CreateUnicastIpAddressEntry makes any changes to the input MIB_UNICASTIPADDRESS_ROW struct, but if it does the remaining code will back-propagate these changes.
	uachanged, _ := wta.toUnicastAddressData()

	if !address.equivalentTo(uachanged) {
		fmt.Println("Yep, it changes!!!")
	}

	*address = *uachanged

	return nil
}

func (address *UnicastAddressData) Delete() error {

	if address == nil {
		return fmt.Errorf("UnicastAddressData.Delete() - receiver argument is nil")
	}

	wta, err := address.toWtMibUnicastipaddressRow()

	if err != nil {
		return err
	}

	return wta.delete()
}

func (address *UnicastAddressData) String() string {

	if address == nil {
		return ""
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
