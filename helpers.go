/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"unsafe"
)

func wcharToString(wchar *uint16, maxLength uint32) string {

	buffer := make([]uint16, maxLength, maxLength)

	const size = uintptr(2) // unsafe.Sizeof(uint16(0))

	start := unsafe.Pointer(wchar)

	var i uint32

	for i = 0; i < maxLength ; i++ {

		letter := *(*uint16)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))

		if letter == 0 {
			break
		}

		buffer[i] = letter
	}

	if i == 0 {
		return ""
	}

	if i < maxLength {
		buffer = buffer[:i]
	}

	return windows.UTF16ToString(buffer)
}

func charToString(char *uint8, maxLength uint32) string {

	buffer := make([]byte, maxLength, maxLength)

	const size = uintptr(1) // unsafe.Sizeof(uint8(0))

	start := unsafe.Pointer(char)

	var i uint32

	for i = 0; i < maxLength; i++ {

		letter := *(*uint8)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))

		if letter == 0 {
			break
		}

		buffer[i] = byte(letter)
	}

	if i == 0 {
		return ""
	}

	if i < maxLength {
		buffer = buffer[:i]
	}

	return string(buffer)
}

func guidToString(guid *windows.GUID) string {
	if guid == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("{%06X-%04X-%04X-%04X-%012X}", guid.Data1, guid.Data2, guid.Data3, guid.Data4[:2],
			guid.Data4[2:])
	}
}

func toIndentedText(text, indent string) string {

	indented := strings.TrimSpace(text)

	indented = strings.Replace(indented, "\n", fmt.Sprintf("\n%s", indent), -1)

	return indent + indented
}

func uint8ToBool(val uint8) bool {
	return val != 0
}

func boolToUint8(val bool) uint8 {
	if val {
		return 1
	} else {
		return 0
	}
}

func allZeroBytes(bytes []byte) bool {

	for _, b := range bytes {
		if b != 0 {
			return false
		}
	}

	return true
}

func guidsEqual(guid1, guid2 *windows.GUID) bool {

	if guid1 == nil {
		return guid2 == nil
	}

	if guid2 == nil {
		return false
	}

	return guid1.Data1 == guid2.Data1 && guid1.Data2 == guid2.Data2 && guid1.Data3 == guid2.Data3 &&
		guid1.Data4 == guid2.Data4
}

func InterfaceLuidToGuid(luid uint64) (*windows.GUID, error) {

	guid := windows.GUID{}

	result := convertInterfaceLuidToGuid(&luid, &guid)

	if result == 0 {
		return &guid, nil
	} else {
		return nil, os.NewSyscallError("iphlpapi.ConvertInterfaceLuidToGuid", windows.Errno(result))
	}
}

func InterfaceGuidToLuid(guid *windows.GUID) (uint64, error) {

	luid := uint64(0)

	result := convertInterfaceGuidToLuid(guid, &luid)

	if result == 0 {
		return luid, nil
	} else {
		return 0, os.NewSyscallError("iphlpapi.ConvertInterfaceGuidToLuid", windows.Errno(result))
	}
}
