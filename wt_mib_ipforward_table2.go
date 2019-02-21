/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "unsafe"

const anySize uint32 = 1 // ANY_SIZE defined in Iprtrmib.h

func (mit *wtMibIpforwardTable2) toSlice() []wtMibIpforwardRow2 {

	numberOfRows := uint32(mit.NumEntries)

	rows := make([]wtMibIpforwardRow2, numberOfRows, numberOfRows)

	if numberOfRows == 0 {
		return rows
	}

	cTablePointer := uintptr(unsafe.Pointer(&mit.Table[0]))
	rowSize := unsafe.Sizeof(mit.Table[0])

	for i := uint32(0); i < numberOfRows; i++ {
		rows[i] = *(*wtMibIpforwardRow2)(unsafe.Pointer(cTablePointer + rowSize*uintptr(i)))
	}

	return rows
}
