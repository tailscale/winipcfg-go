/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// Corresponds to Windows struct MIB_IPINTERFACE_ROW
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipinterface_row).
type IpInterface struct {
	//
	// Key Structure;
	//
	Family         AddressFamily
	InterfaceLuid  uint64
	InterfaceIndex uint32

	//
	// Read-Write fields.
	//

	//
	// Fields currently not exposed.
	//
	MaxReassemblySize              uint32
	InterfaceIdentifier            uint64
	MinRouterAdvertisementInterval uint32
	MaxRouterAdvertisementInterval uint32

	//
	// Fileds currently exposed.
	//
	AdvertisingEnabled                   bool
	ForwardingEnabled                    bool
	WeakHostSend                         bool
	WeakHostReceive                      bool
	UseAutomaticMetric                   bool
	UseNeighborUnreachabilityDetection   bool
	ManagedAddressConfigurationSupported bool
	OtherStatefulConfigurationSupported  bool
	AdvertiseDefaultRoute                bool

	RouterDiscoveryBehavior NlRouterDiscoveryBehavior
	// DupAddrDetectTransmits in RFC 2462.
	DadTransmits      uint32
	BaseReachableTime uint32
	RetransmitTime    uint32
	// Path MTU discovery timeout (in ms).
	PathMtuDiscoveryTimeout uint32

	LinkLocalAddressBehavior NlLinkLocalAddressBehavior
	// In ms.
	LinkLocalAddressTimeout uint32
	// Zone part of a SCOPE_ID.
	ZoneIndices      [ScopeLevelCount]uint32
	SitePrefixLength uint32
	Metric           uint32
	NlMtu            uint32

	//
	// Read Only fields.
	//
	Connected                 bool
	SupportsWakeUpPatterns    bool
	SupportsNeighborDiscovery bool
	SupportsRouterDiscovery   bool

	ReachableTime uint32

	TransmitOffload NlInterfaceOffloadRodFlags
	ReceiveOffload  NlInterfaceOffloadRodFlags

	//
	// Disables using default route on the interface. This flag
	// can be used by VPN clients to restrict Split tunnelling.
	//
	DisableDefaultRoutes bool
}

// Corresponds to GetIpInterfaceTable function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfacetable)
func GetIpInterfaces(family AddressFamily) ([]*IpInterface, error) {

	rows, err := getWtMibIpinterfaceRows(family)

	if err != nil {
		return nil, err
	}

	length := len(rows)

	ipifcs := make([]*IpInterface, length, length)

	for idx, row := range rows {
		ipifcs[idx] = row.toIpInterface()
	}

	return ipifcs, nil
}

// Corresponds to GetIpInterfaceEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfaceentry).
// Argument 'family' has to be either AF_INET or AF_INET6.
func GetIpInterface(interfaceLuid uint64, family AddressFamily) (*IpInterface, error) {

	row, err := getWtMibIpinterfaceRow(interfaceLuid, family)

	if err != nil {
		return nil, err
	} else {
		return row.toIpInterface(), nil
	}
}

// Saves (activates) modified IpInterface. Corresponds to SetIpInterfaceEntry function
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipinterfaceentry).
//
// Note that fields Family, InterfaceLuid and InterfaceIndex are used for identifying address to change, meaning that
// they cannot be changed by using this method. Changing some of these fields would cause updating some other IP
// interface. Fields which are "changeable" by this method are between AdvertisingEnabled and NlMtu, inclusive.
// The workflow of using this method is:
// 1) Get IpInterface instance by using any of getter methods (i.e. GetIpInterface or any other);
// 2) Change one or more of "changeable" fields enumerated above;
// 3) Calling this method to activate the changes.
func (ipifc *IpInterface) Set() error {

	old, err := getWtMibIpinterfaceRow(ipifc.InterfaceLuid, ipifc.Family)

	if err != nil {
		return err
	}

	old.AdvertisingEnabled = boolToUint8(ipifc.AdvertisingEnabled)
	old.ForwardingEnabled = boolToUint8(ipifc.ForwardingEnabled)
	old.WeakHostSend = boolToUint8(ipifc.WeakHostSend)
	old.WeakHostReceive = boolToUint8(ipifc.WeakHostReceive)
	old.UseAutomaticMetric = boolToUint8(ipifc.UseAutomaticMetric)
	old.UseNeighborUnreachabilityDetection = boolToUint8(ipifc.UseNeighborUnreachabilityDetection)
	old.ManagedAddressConfigurationSupported = boolToUint8(ipifc.ManagedAddressConfigurationSupported)
	old.OtherStatefulConfigurationSupported = boolToUint8(ipifc.OtherStatefulConfigurationSupported)
	old.AdvertiseDefaultRoute = boolToUint8(ipifc.AdvertiseDefaultRoute)
	old.RouterDiscoveryBehavior = ipifc.RouterDiscoveryBehavior
	old.DadTransmits = ipifc.DadTransmits
	old.BaseReachableTime = ipifc.BaseReachableTime
	old.RetransmitTime = ipifc.RetransmitTime
	old.PathMtuDiscoveryTimeout = ipifc.PathMtuDiscoveryTimeout
	old.LinkLocalAddressBehavior = ipifc.LinkLocalAddressBehavior
	old.LinkLocalAddressTimeout = ipifc.LinkLocalAddressTimeout
	old.ZoneIndices = ipifc.ZoneIndices
	old.SitePrefixLength = ipifc.SitePrefixLength
	old.Metric = ipifc.Metric
	old.NlMtu = ipifc.NlMtu

	// Patch that fixes SitePrefixLength issue
	// (https://stackoverflow.com/questions/54857292/setipinterfaceentry-returns-error-invalid-parameter?noredirect=1)
	if old.SitePrefixLength > 128 || (old.SitePrefixLength > 32 && old.Family == AF_INET) {
		old.SitePrefixLength = 0
	}

	return old.set()
}

func (mir *IpInterface) String() string {

	if mir == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`Family: %s
InterfaceLuid: %d
InterfaceIndex: %d
MaxReassemblySize: %d
InterfaceIdentifier: %d
MinRouterAdvertisementInterval: %d
MaxRouterAdvertisementInterval: %d
AdvertisingEnabled: %v
ForwardingEnabled: %v
WeakHostSend: %v
WeakHostReceive: %v
UseAutomaticMetric: %v
UseNeighborUnreachabilityDetection: %v
ManagedAddressConfigurationSupported: %v
OtherStatefulConfigurationSupported: %v
AdvertiseDefaultRoute: %v
RouterDiscoveryBehavior: %s
DadTransmits: %d
BaseReachableTime: %d
RetransmitTime: %d
PathMtuDiscoveryTimeout: %d
LinkLocalAddressBehavior: %s
LinkLocalAddressTimeout: %d
ZoneIndices: %v
SitePrefixLength: %d
Metric: %d
NlMtu: %d
Connected: %v
SupportsWakeUpPatterns: %v
SupportsNeighborDiscovery: %v
SupportsRouterDiscovery: %v
ReachableTime: %d
TransmitOffload:
%s
ReceiveOffload:
%s
DisableDefaultRoutes: %v`, mir.Family.String(), mir.InterfaceLuid, mir.InterfaceIndex, mir.MaxReassemblySize,
		mir.InterfaceIdentifier, mir.MinRouterAdvertisementInterval, mir.MaxRouterAdvertisementInterval,
		mir.AdvertisingEnabled, mir.ForwardingEnabled, mir.WeakHostSend, mir.WeakHostReceive, mir.UseAutomaticMetric,
		mir.UseNeighborUnreachabilityDetection, mir.ManagedAddressConfigurationSupported,
		mir.OtherStatefulConfigurationSupported, mir.AdvertiseDefaultRoute, mir.RouterDiscoveryBehavior.String(),
		mir.DadTransmits, mir.BaseReachableTime, mir.RetransmitTime, mir.PathMtuDiscoveryTimeout,
		mir.LinkLocalAddressBehavior.String(), mir.LinkLocalAddressTimeout, mir.ZoneIndices, mir.SitePrefixLength,
		mir.Metric, mir.NlMtu, mir.Connected, mir.SupportsWakeUpPatterns, mir.SupportsNeighborDiscovery,
		mir.SupportsRouterDiscovery, mir.ReachableTime, toIndentedText(mir.TransmitOffload.String(), "  "),
		toIndentedText(mir.ReceiveOffload.String(), "  "), mir.DisableDefaultRoutes)
}
