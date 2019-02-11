/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

// Types from https://docs.microsoft.com/en-us/windows/desktop/winprog/windows-data-types
type BYTE uint8
type BOOLEAN BYTE
type CHAR uint8
type UCHAR uint8
type UINT8 uint8
type DWORD uint32
type USHORT uint16
type ULONG DWORD
type ULONG64 uint64
type HANDLE uintptr
type PHANDLE *HANDLE
type PVOID uintptr
type LONGLONG int64
type WCHAR uint16
type INT int32

// Defined in winnt.h, it's a union...
type LARGE_INTEGER LONGLONG

// Defined in guid.h
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]UCHAR
}

// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ns-ifdef-_net_luid_lh
// Defined in ifdef
type NET_LUID_LH ULONG64

// Defined in ifdef.h
type NET_IFINDEX ULONG
type IF_INDEX NET_IFINDEX
type NET_IF_COMPARTMENT_ID uint32
type NET_IF_NETWORK_GUID GUID
type NET_LUID NET_LUID_LH
type IF_LUID NET_LUID

func (b BOOLEAN) String() string {
	if b == 0 {
		return "FALSE"
	} else {
		return "TRUE"
	}
}

// Helper functions

func SystemErrorCode(code uint32) string {
	//https://docs.microsoft.com/en-us/windows/desktop/Debug/system-error-codes--0-499-
	switch code {
	case 0:
		return "ERROR_SUCCESS"
	case 2:
		return "ERROR_FILE_NOT_FOUND"
	case 5:
		return "ERROR_ACCESS_DENIED"
	case 50:
		return "ERROR_NOT_SUPPORTED"
	case 87:
		return "ERROR_INVALID_PARAMETER"
	case 122:
		return "ERROR_INSUFFICIENT_BUFFER"
	case 1168:
		return "ERROR_NOT_FOUND"
	case 5010:
		return "ERROR_OBJECT_ALREADY_EXISTS"
	default:
		return fmt.Sprintf("ERROR CODE: %d", code)
	}
}
