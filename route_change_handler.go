/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"os"
	"sync"
	"unsafe"
)

type RouteChangeCallback struct {
	cb func(notificationType MibNotificationType, route *Route)
}

var (
	routeChangeMutex     = sync.Mutex{}
	routeChangeCallbacks = make(map[*RouteChangeCallback]bool)
	routeChangeHandle    = uintptr(0)
)

func RegisterRouteChangeCallback(cb func(notificationType MibNotificationType, route *Route)) (*RouteChangeCallback, error) {
	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()
	s := &RouteChangeCallback{cb}
	routeChangeCallbacks[s] = true
	if routeChangeHandle == 0 {
		result := notifyRouteChange2(AF_UNSPEC, windows.NewCallback(routeChanged), 0, false, unsafe.Pointer(&routeChangeHandle))
		if result != 0 {
			delete(routeChangeCallbacks, s)
			return nil, os.NewSyscallError("iphlpapi.NotifyRouteChange2", windows.Errno(result))
		}
	}
	return s, nil
}

func UnregisterRouteChangeCallback(cb *RouteChangeCallback) error {
	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()
	delete(routeChangeCallbacks, cb)
	if len(routeChangeCallbacks) == 0 && routeChangeHandle != 0 {
		result := cancelMibChangeNotify2(routeChangeHandle)
		if result != 0 {
			return os.NewSyscallError("iphlpapi.CancelMibChangeNotify2", windows.Errno(result))
		}
		routeChangeHandle = uintptr(0)
	}
	return nil
}

func routeChanged(callerContext unsafe.Pointer, wtr *wtMibIpforwardRow2, notificationType MibNotificationType) uintptr {
	route, err := wtr.toRoute()
	if route == nil || err != nil {
		return 0
	}
	routeChangeMutex.Lock()
	for cb := range routeChangeCallbacks {
		cb.cb(notificationType, route)
	}
	routeChangeMutex.Unlock()
	return 0
}
