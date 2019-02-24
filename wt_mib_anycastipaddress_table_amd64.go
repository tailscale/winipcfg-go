/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// MIB_ANYCASTIPADDRESS_TABLE defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-mib_anycastipaddress_table)
type wtMibAnycastipaddressTable struct {
	NumEntries uint32 // Windows type: ULONG
	Table      [anySize]wtMibAnycastipaddressRow;
}
