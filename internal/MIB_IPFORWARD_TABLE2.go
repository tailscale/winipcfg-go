/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package internal

import "unsafe"

// Defined in Iprtrmib.h
const ANY_SIZE uint32 = 1

func (mit *MIB_IPFORWARD_TABLE2) ToSlice() []MIB_IPFORWARD_ROW2 {

	numberOfRows := uint32(mit.NumEntries)

	rows := make([]MIB_IPFORWARD_ROW2, numberOfRows, numberOfRows)

	if numberOfRows == 0 {
		return rows
	}

	cTablePointer := uintptr(unsafe.Pointer(&mit.Table[0]))
	rowSize := unsafe.Sizeof(mit.Table[0])

	for i := uint32(0); i < numberOfRows; i++ {
		rows[i] = *(*MIB_IPFORWARD_ROW2)(unsafe.Pointer(cTablePointer + rowSize * uintptr(i)))
	}

	return rows
}
