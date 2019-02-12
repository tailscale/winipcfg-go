/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipforward_table2
type MIB_IPFORWARD_TABLE2 struct {
	NumEntries uint32 // Windows type: ULONG
	Table [ANY_SIZE]MIB_IPFORWARD_ROW2
}
