/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"golang.org/x/sys/windows"
	"sync"
	"unsafe"
)

// Defines function that can be used as a callback.
type UnicastAddressChangeCallback func(uar *UnicastAddressData, notificationType MibNotificationType)

var (
	unicastAddressChangeMutex     = sync.Mutex{}
	unicastAddressChangeCallbacks = make([]*UnicastAddressChangeCallback, 0)
	unicastAddressChangeHandle    = uintptr(0)
)

// Registering new UnicastAddressChangeCallback. If this particular callback is already registered, the function will
// silently return.
func RegisterUnicastAddressChangeCallback(callback *UnicastAddressChangeCallback) error {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	if indexOfUnicastAddressChangeCallback(callback) < 0 {

		unicastAddressChangeCallbacks = append(unicastAddressChangeCallbacks, callback)

		return checkUnicastAddressChangeSubscribed()
	}

	return nil
}

// Unregistering UnicastAddressChangeCallback.
func UnregisterUnicastAddressChangeCallback(callback *UnicastAddressChangeCallback) error {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	index := indexOfUnicastAddressChangeCallback(callback)

	if index < 0 {
		// It isn't registered at all, so simply return:
		return nil
	}

	count := len(unicastAddressChangeCallbacks)

	if count == 1 {
		// The last one, so empty the slice:
		unicastAddressChangeCallbacks = make([]*UnicastAddressChangeCallback, 0)

		err := checkUnicastAddressChangeSubscribed()

		if err != nil {
			return err
		}
	} else if index == 0 {
		unicastAddressChangeCallbacks = unicastAddressChangeCallbacks[1:]
	} else if index == count - 1 {
		unicastAddressChangeCallbacks = unicastAddressChangeCallbacks[:index]
	} else {
		unicastAddressChangeCallbacks = append(unicastAddressChangeCallbacks[:index],
			unicastAddressChangeCallbacks[index + 1:]...)
	}

	return nil
}

// For checking if particular handler is already registered.
func UnicastAddressChangeCallbackRegistered(callback *UnicastAddressChangeCallback) bool {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	return indexOfUnicastAddressChangeCallback(callback) >= 0
}

// Unsubscribes all subscribed callbacks, and aborts listening for unicast address changes.
func StopListeningForUnicastAddressChange() {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	unicastAddressChangeCallbacks = make([]*UnicastAddressChangeCallback, 0)

	_ = checkUnicastAddressChangeSubscribed()
}

// Should be called from a locked code!
func indexOfUnicastAddressChangeCallback(callback *UnicastAddressChangeCallback) int {

	for idx, c := range unicastAddressChangeCallbacks {
		if callback == c {
			return idx
		}
	}

	return -1
}

// Should be called from a locked code!
func checkUnicastAddressChangeSubscribed() error {

	if unicastAddressChangeHandle == 0 {
		// We aren't subscribed.
		if len(unicastAddressChangeCallbacks) > 0 {
			// We should subscribe!
			result := notifyUnicastIpAddressChange(AF_UNSPEC, windows.NewCallback(unicastAddressChanged), 0,
				false, unsafe.Pointer(&unicastAddressChangeHandle))

			if result != 0 {
				return windows.Errno(result)
			}
		}
	} else {
		// We are subscribed.
		if len(unicastAddressChangeCallbacks) < 1 {
			// We should unsubscribe!
			result := cancelMibChangeNotify2(unicastAddressChangeHandle)

			if result != 0 {
				return windows.Errno(result)
			}

			unicastAddressChangeHandle = uintptr(0)
		}
	}

	return nil
}

func unicastAddressChanged(callerContext unsafe.Pointer, wtUar *wtMibUnicastipaddressRow, notificationType MibNotificationType) uintptr {

	uar, err := wtUar.toUnicastAddressData()

	if err != nil {
		// TODO: We should at lest implement logging.
		return 0
	}

	// go routine used to avoid blocking OS call.
	go notifyUnicastAddressChangedCallbacks(uar, notificationType)

	return 0
}

func notifyUnicastAddressChangedCallbacks(uar *UnicastAddressData, notificationType MibNotificationType) {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	for _, c := range unicastAddressChangeCallbacks {
		(*c)(uar, notificationType)
	}
}
