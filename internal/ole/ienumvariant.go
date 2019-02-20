/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2013-2017 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ole

import "unsafe"

type IEnumVARIANT struct {
	IUnknown
}

type IEnumVARIANTVtbl struct {
	IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

func (v *IEnumVARIANT) VTable() *IEnumVARIANTVtbl {
	return (*IEnumVARIANTVtbl)(unsafe.Pointer(v.RawVTable))
}
