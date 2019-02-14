/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_table2
// MIB_IPFORWARD_TABLE2 defined in netioapi.h
type wtMibIpforwardTable2 struct {
	NumEntries uint32 // Windows type: ULONG
	// In 32-bit builds we have to artificially add an offset, in order to get the same size of the struct.
	offset [4]byte
	Table  [ANY_SIZE]wtMibIpforwardRow2
}
