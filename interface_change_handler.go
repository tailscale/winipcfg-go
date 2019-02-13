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
type InterfaceChangeCallback func(ifc *MibIpinterfaceRow, notificationType MibNotificationType)

var (
	interfaceChangeMutex     = sync.Mutex{}
	interfaceChangeCallbacks = make([]*InterfaceChangeCallback, 0)
	interfaceChangeHandle    = uintptr(0)
)

// Registering new InterfaceChangeCallback. If this particular callback is already registered, the function will
// silently return.
func RegisterInterfaceChangeCallback(callback *InterfaceChangeCallback) error {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	if indexOfInterfaceChangeCallback(callback) < 0 {

		interfaceChangeCallbacks = append(interfaceChangeCallbacks, callback)

		return checkInterfaceChangeSubscribed()
	}

	return nil
}

// Unregistering InterfaceChangeCallback.
func UnregisterInterfaceChangeCallback(callback *InterfaceChangeCallback) error {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	index := indexOfInterfaceChangeCallback(callback)

	if index < 0 {
		// It isn't registered at all, so simply return:
		return nil
	}

	count := len(interfaceChangeCallbacks)

	if count == 1 {
		// The last one, so empty the slice:
		interfaceChangeCallbacks = make([]*InterfaceChangeCallback, 0)

		err := checkInterfaceChangeSubscribed()

		if err != nil {
			return err
		}
	} else if index == 0 {
		interfaceChangeCallbacks = interfaceChangeCallbacks[1:]
	} else if index == count - 1 {
		interfaceChangeCallbacks = interfaceChangeCallbacks[:index]
	} else {
		interfaceChangeCallbacks = append(interfaceChangeCallbacks[:index], interfaceChangeCallbacks[index + 1:]...)
	}

	return nil
}

// For checking if particular handler is already registered.
func InterfaceChangeCallbackRegistered(callback *InterfaceChangeCallback) bool {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	return indexOfInterfaceChangeCallback(callback) >= 0
}

// Unsubscribes all subscribed callbacks, and aborts listening for interface changes.
func StopListeningForInterfaceChange() {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	interfaceChangeCallbacks = make([]*InterfaceChangeCallback, 0)

	_ = checkInterfaceChangeSubscribed()
}

// Should be called from a locked code!
func indexOfInterfaceChangeCallback(callback *InterfaceChangeCallback) int {

	for idx, c := range interfaceChangeCallbacks {
		if callback == c {
			return idx
		}
	}

	return -1
}

// Should be called from a locked code!
func checkInterfaceChangeSubscribed() error {

	if interfaceChangeHandle == 0 {
		// We aren't subscribed.
		if len(interfaceChangeCallbacks) > 0 {
			// We should subscribe!
			result := notifyIpInterfaceChange(AF_UNSPEC, windows.NewCallback(interfaceChanged), 0,
				false, unsafe.Pointer(&interfaceChangeHandle))

			if result != 0 {
				return windows.Errno(result)
			}
		}
	} else {
		// We are subscribed.
		if len(interfaceChangeCallbacks) < 1 {
			// We should unsubscribe!
			result := cancelMibChangeNotify2(interfaceChangeHandle)

			if result != 0 {
				return windows.Errno(result)
			}

			interfaceChangeHandle = uintptr(0)
		}
	}

	return nil
}

func interfaceChanged(callerContext unsafe.Pointer, wtIfc *wtMibIpinterfaceRow, notificationType MibNotificationType) uintptr {

	ifc := wtIfc.toMibIpinterfaceRow()

	// go routine used to avoid blocking OS call.
	go notifyInterfaceChangedCallbacks(ifc, notificationType)

	return 0
}

func notifyInterfaceChangedCallbacks(ifc *MibIpinterfaceRow, notificationType MibNotificationType) {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	for _, c := range interfaceChangeCallbacks {
		(*c)(ifc, notificationType)
	}
}
