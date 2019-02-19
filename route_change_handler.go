/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"net"
	"os"
	"sync"
	"unsafe"
)

// Defines function that can be used as a callback.
type RouteChangeCallback func(notificationType MibNotificationType, interfaceLuid uint64, destination *net.IPNet,
	nextHop *net.IP)

var (
	routeChangeMutex     = sync.Mutex{}
	routeChangeCallbacks = make([]*RouteChangeCallback, 0)
	routeChangeHandle    = uintptr(0)
)

// Registering new RouteChangeCallback. If this particular callback is already registered, the function will
// silently return.
func RegisterRouteChangeCallback(callback *RouteChangeCallback) error {

	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()

	if indexOfRouteChangeCallback(callback) < 0 {

		routeChangeCallbacks = append(routeChangeCallbacks, callback)

		return checkRouteChangeSubscribed()
	}

	return nil
}

// Unregistering RouteChangeCallback.
func UnregisterRouteChangeCallback(callback *RouteChangeCallback) error {

	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()

	index := indexOfRouteChangeCallback(callback)

	if index < 0 {
		// It isn't registered at all, so simply return:
		return nil
	}

	count := len(routeChangeCallbacks)

	if count == 1 {
		// The last one, so empty the slice:
		routeChangeCallbacks = make([]*RouteChangeCallback, 0)

		err := checkRouteChangeSubscribed()

		if err != nil {
			return err
		}
	} else if index == 0 {
		routeChangeCallbacks = routeChangeCallbacks[1:]
	} else if index == count-1 {
		routeChangeCallbacks = routeChangeCallbacks[:index]
	} else {
		routeChangeCallbacks = append(routeChangeCallbacks[:index], routeChangeCallbacks[index+1:]...)
	}

	return nil
}

// For checking if particular handler is already registered.
func RouteChangeCallbackRegistered(callback *RouteChangeCallback) bool {

	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()

	return indexOfRouteChangeCallback(callback) >= 0
}

// Unsubscribes all subscribed callbacks, and aborts listening for route changes.
func StopListeningForRouteChange() {

	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()

	routeChangeCallbacks = make([]*RouteChangeCallback, 0)

	_ = checkRouteChangeSubscribed()
}

// Should be called from a locked code!
func indexOfRouteChangeCallback(callback *RouteChangeCallback) int {

	for idx, c := range routeChangeCallbacks {
		if callback == c {
			return idx
		}
	}

	return -1
}

// Should be called from a locked code!
func checkRouteChangeSubscribed() error {

	if routeChangeHandle == 0 {
		// We aren't subscribed.
		if len(routeChangeCallbacks) > 0 {
			// We should subscribe!
			result := notifyRouteChange2(AF_UNSPEC, windows.NewCallback(routeChanged), 0,
				false, unsafe.Pointer(&routeChangeHandle))

			if result != 0 {
				return os.NewSyscallError("iphlpapi.NotifyRouteChange2", windows.Errno(result))
			}
		}
	} else {
		// We are subscribed.
		if len(routeChangeCallbacks) < 1 {
			// We should unsubscribe!
			result := cancelMibChangeNotify2(routeChangeHandle)

			if result != 0 {
				return os.NewSyscallError("iphlpapi.CancelMibChangeNotify2", windows.Errno(result))
			}

			routeChangeHandle = uintptr(0)
		}
	}

	return nil
}

func routeChanged(callerContext unsafe.Pointer, wtr *wtMibIpforwardRow2, notificationType MibNotificationType) uintptr {

	routeData, err := wtr.extractRouteData()

	if err != nil {
		// TODO: At lest we should add some logging here.
		return 0
	}

	// go routine used to avoid blocking OS call.
	go notifyRouteChangedCallbacks(notificationType, wtr.InterfaceLuid, &routeData.Destination, &routeData.NextHop)

	return 0
}

func notifyRouteChangedCallbacks(notificationType MibNotificationType, interfaceLuid uint64, destination *net.IPNet,
	nextHop *net.IP) {

	routeChangeMutex.Lock()
	defer routeChangeMutex.Unlock()

	for _, c := range routeChangeCallbacks {
		(*c)(notificationType, interfaceLuid, destination, nextHop)
	}
}
