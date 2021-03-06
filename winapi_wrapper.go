/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// Interface - related functions

// https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses
//sys	getAdaptersAddresses(Family uint32, Flags uint32, Reserved uintptr, AdapterAddresses *wtIpAdapterAddresses, SizePointer *uint32) (result uint32) = iphlpapi.GetAdaptersAddresses

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-initializeipinterfaceentry
//sys	initializeIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) = iphlpapi.InitializeIpInterfaceEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfaceentry
//sys	getIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) = iphlpapi.GetIpInterfaceEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipinterfacetable
//sys	getIpInterfaceTable(Family AddressFamily, Table unsafe.Pointer) (result int32) = iphlpapi.GetIpInterfaceTable

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipinterfaceentry
//sys	setIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) = iphlpapi.SetIpInterfaceEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-freemibtable
//sys	freeMibTable(memory unsafe.Pointer) = iphlpapi.FreeMibTable

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex
//sys	getIfEntry2Ex(Level MibIfEntryLevel, Row *wtMibIfRow2) (result int32) = iphlpapi.GetIfEntry2Ex

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getiftable2ex
//sys	getIfTable2Ex(Level MibIfEntryLevel, Table unsafe.Pointer) (result int32) = iphlpapi.GetIfTable2Ex

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-convertinterfaceluidtoguid
//sys	convertInterfaceLuidToGuid(InterfaceLuid *uint64, InterfaceGuid *windows.GUID) (result int32) = iphlpapi.ConvertInterfaceLuidToGuid

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-convertinterfaceguidtoluid
//sys	convertInterfaceGuidToLuid(InterfaceGuid *windows.GUID, InterfaceLuid *uint64) (result int32) = iphlpapi.ConvertInterfaceGuidToLuid

// Unicast IP address - related functions

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddresstable
//sys	getUnicastIpAddressTable(Family AddressFamily, Table unsafe.Pointer) (result int32) = iphlpapi.GetUnicastIpAddressTable

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getunicastipaddressentry
//sys	getUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) = iphlpapi.GetUnicastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setunicastipaddressentry
//sys	setUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) = iphlpapi.SetUnicastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-initializeunicastipaddressentry
//sys	initializeUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) = iphlpapi.InitializeUnicastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createunicastipaddressentry
//sys	createUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) = iphlpapi.CreateUnicastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteunicastipaddressentry
//sys	deleteUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) = iphlpapi.DeleteUnicastIpAddressEntry

// Anycast IP address - related functions

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddresstable
//sys	getAnycastIpAddressTable(Family AddressFamily, Table unsafe.Pointer) (result int32) = iphlpapi.GetAnycastIpAddressTable

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getanycastipaddressentry
//sys	getAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) = iphlpapi.GetAnycastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createanycastipaddressentry
//sys	createAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) = iphlpapi.CreateAnycastIpAddressEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteanycastipaddressentry
//sys	deleteAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) = iphlpapi.DeleteAnycastIpAddressEntry

// Routing - related functions

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardtable2
//sys	getIpForwardTable2(family AddressFamily, table unsafe.Pointer) (result int32) = iphlpapi.GetIpForwardTable2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getipforwardentry2
//sys	getIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.GetIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-initializeipforwardentry
//sys	initializeIpForwardEntry(route *wtMibIpforwardRow2) (result int32) = iphlpapi.InitializeIpForwardEntry

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-createipforwardentry2
//sys	createIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.CreateIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-setipforwardentry2
//sys	setIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.SetIpForwardEntry2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-deleteipforwardentry2
//sys	deleteIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) = iphlpapi.DeleteIpForwardEntry2

// Notifications - related functions

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyipinterfacechange
//sys	notifyIpInterfaceChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyIpInterfaceChange

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyunicastipaddresschange
//sys	notifyUnicastIpAddressChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyUnicastIpAddressChange

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyroutechange2
//sys	notifyRouteChange2(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) = iphlpapi.NotifyRouteChange2

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-cancelmibchangenotify2
//sys	cancelMibChangeNotify2(NotificationHandle uintptr) (result int32) = iphlpapi.CancelMibChangeNotify2
