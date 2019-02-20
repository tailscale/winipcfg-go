/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2013-2017 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ole

import "unsafe"

type IProvideClassInfo struct {
	IUnknown
}

type IProvideClassInfoVtbl struct {
	IUnknownVtbl
	GetClassInfo uintptr
}

func (v *IProvideClassInfo) VTable() *IProvideClassInfoVtbl {
	return (*IProvideClassInfoVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IProvideClassInfo) GetClassInfo() (cinfo *ITypeInfo, err error) {
	cinfo, err = getClassInfo(v)
	return
}
