/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// MIB_IF_TABLE2 defined in netioapi.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_table2)
type wtMibIfTable2 struct {
	NumEntries uint32 // Windows type: ULONG

	offset1 [4]byte // Layout correction field

	Table      [anySize]wtMibIfRow2
}
