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

type InterfaceChangeCallback struct {
	cb func(notificationType MibNotificationType, interfaceLuid uint64)
}

var (
	interfaceChangeMutex     = sync.Mutex{}
	interfaceChangeCallbacks = make(map[*InterfaceChangeCallback]bool)
	interfaceChangeHandle    = uintptr(0)
)

// Registering new InterfaceChangeCallback. If this particular callback is already registered, the function will
// silently return. Returned InterfaceChangeCallback structure should be used with UnregisterInterfaceChangeCallback
// function to unregister.
func RegisterInterfaceChangeCallback(callback func(notificationType MibNotificationType,
	interfaceLuid uint64)) (*InterfaceChangeCallback, error) {

	cb := &InterfaceChangeCallback{callback}

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	interfaceChangeCallbacks[cb] = true

	if interfaceChangeHandle == 0 {

		result := notifyIpInterfaceChange(AF_UNSPEC, windows.NewCallback(interfaceChanged), 0,
			false, unsafe.Pointer(&interfaceChangeHandle))

		if result != 0 {
			delete(interfaceChangeCallbacks, cb)
			interfaceChangeHandle = 0
			return nil, os.NewSyscallError("iphlpapi.NotifyIpInterfaceChange", windows.Errno(result))
		}
	}

	return cb, nil
}

// Unregistering InterfaceChangeCallback.
func UnregisterInterfaceChangeCallback(callback *InterfaceChangeCallback) error {

	interfaceChangeMutex.Lock()
	defer interfaceChangeMutex.Unlock()

	delete(interfaceChangeCallbacks, callback)

	if len(interfaceChangeCallbacks) < 1 && interfaceChangeHandle != 0 {

		result := cancelMibChangeNotify2(interfaceChangeHandle)

		if result != 0 {
			return os.NewSyscallError("iphlpapi.CancelMibChangeNotify2", windows.Errno(result))
		}

		interfaceChangeHandle = uintptr(0)
	}

	return nil
}

func interfaceChanged(callerContext unsafe.Pointer, wtIfc *wtMibIpinterfaceRow,
	notificationType MibNotificationType) uintptr {

	if wtIfc == nil {
		return 0
	}

	interfaceChangeMutex.Lock()

	for cb := range interfaceChangeCallbacks {
		cb.cb(notificationType, wtIfc.InterfaceLuid)
	}

	interfaceChangeMutex.Unlock()

	return 0
}
