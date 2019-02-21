/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipinterface_table
// MIB_IPINTERFACE_TABLE defined in netioapi.h
type wtMibIpinterfaceTable struct {
	NumEntries uint32 // Windows type: ULONG
	Table      [ANY_SIZE]wtMibIpinterfaceRow
}
