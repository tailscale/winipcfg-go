// +build windows,amd64

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2013-2017 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ole

import (
	"errors"
	"syscall"
	"time"
	"unsafe"
)

// GetVariantDate converts COM Variant Time value to Go time.Time.
func GetVariantDate(value uint64) (time.Time, error) {
	var st syscall.Systemtime
	r, _, _ := procVariantTimeToSystemTime.Call(uintptr(value), uintptr(unsafe.Pointer(&st)))
	if r != 0 {
		return time.Date(int(st.Year), time.Month(st.Month), int(st.Day), int(st.Hour), int(st.Minute), int(st.Second), int(st.Milliseconds/1000), time.UTC), nil
	}
	return time.Now(), errors.New("Could not convert to time, passing current time.")
}
