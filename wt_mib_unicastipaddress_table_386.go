/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_unicastipaddress_table
// MIB_UNICASTIPADDRESS_TABLE defined in netioapi.h
type wtMibUnicastipaddressTable struct {
	NumEntries uint32 // Windows type: ULONG

	offset1 [4]uint8 // Layout correction field

	Table [anySize]wtMibUnicastipaddressRow
}
