/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

const (
	expectedStringLength = 1000
)

func wcharToString(wchar *uint16) string {

	buffer := make([]uint16, expectedStringLength)

	start := unsafe.Pointer(wchar)
	size := unsafe.Sizeof(uint16(0))

	for i := 0; ; i++ {

		letter := *(*uint16)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))

		if letter == 0 {

			if i < expectedStringLength {
				buffer[i] = 0
			}

			break
		}

		if i < expectedStringLength {
			buffer[i] = letter
		} else {
			buffer = append(buffer, letter)
		}
	}

	return windows.UTF16ToString(buffer)
}

func charToString(char *uint8) string {

	buffer := make([]byte, expectedStringLength)

	start := unsafe.Pointer(char)
	size := unsafe.Sizeof(uint8(0))

	var i int

	for i = 0; ; i++ {

		letter := *(*uint8)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))

		if letter == 0 {
			break
		}

		if i < expectedStringLength {
			buffer[i] = byte(letter)
		} else {
			buffer = append(buffer, byte(letter))
		}
	}

	if i < len(buffer) {
		buffer = buffer[:i]
	}

	return string(buffer)
}
