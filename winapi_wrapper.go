/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses
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
//sys	notifyIpInterfaceChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyIpInterfaceChange

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyunicastipaddresschange
//sys	notifyUnicastIpAddressChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyUnicastIpAddressChange

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyroutechange2
//sys	notifyRouteChange2(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyRouteChange2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-cancelmibchangenotify2
//sys	cancelMibChangeNotify2(NotificationHandle uintptr) (result int32) = iphlpapi.CancelMibChangeNotify2
