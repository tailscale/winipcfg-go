/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

//sys	cancelMibChangeNotify2(NotificationHandle uintptr) (result int32) = iphlpapi.CancelMibChangeNotify2

//sys	getAdaptersAddresses(Family uint32, Flags uint32, Reserved uintptr, AdapterAddresses *wtIpAdapterAddresses, SizePointer *uint32) (result uint32) = iphlpapi.GetAdaptersAddresses

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2
//sys	getIpForwardTable2(family AddressFamily, table *wtMibIpforwardTable2) (result int32) = iphlpapi.GetIpForwardTable2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-freemibtable
//sys	freeMibTable(memory uintptr) = iphlpapi.FreeMibTable

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createipforwardentry2
//sys	createIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.CreateIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipforwardentry2
//sys	setIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.SetIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteipforwardentry2
//sys	deleteIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.DeleteIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyipinterfacechange
//sys	notifyIpInterfaceChange(route *wtMibIpforwardRow2) (result int32) = iphlpapi.NotifyIpInterfaceChange

//setIpForwardEntry2:           iphlpapi.MustFindProc(""),
//deleteIpForwardEntry2:        iphlpapi.MustFindProc(""),
//notifyIpInterfaceChange:      iphlpapi.MustFindProc(""),
//notifyRouteChange2:           iphlpapi.MustFindProc("NotifyRouteChange2"),
//notifyUnicastIpAddressChange: iphlpapi.MustFindProc("NotifyUnicastIpAddressChange"),
//cancelMibChangeNotify2:       iphlpapi.MustFindProc("CancelMibChangeNotify2"),