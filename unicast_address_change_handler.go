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
type UnicastAddressChangeCallback struct {
	cb func(notificationType MibNotificationType, interfaceLuid uint64, ip *net.IP)
}

var (
	unicastAddressChangeMutex     = sync.Mutex{}
	unicastAddressChangeCallbacks = make(map[*UnicastAddressChangeCallback]bool)
	unicastAddressChangeHandle    = uintptr(0)
)

func RegisterUnicastAddressChangeCallback(
	callback func(notificationType MibNotificationType, interfaceLuid uint64, ip *net.IP)) (*UnicastAddressChangeCallback, error) {

	cb := &UnicastAddressChangeCallback{callback}

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	unicastAddressChangeCallbacks[cb] = true

	if unicastAddressChangeHandle == 0 {

		result := notifyUnicastIpAddressChange(AF_UNSPEC, windows.NewCallback(unicastAddressChanged), 0,
			false, unsafe.Pointer(&unicastAddressChangeHandle))

		if result != 0 {
			delete(unicastAddressChangeCallbacks, cb)
			unicastAddressChangeHandle = 0
			return nil, os.NewSyscallError("iphlpapi.NotifyUnicastIpAddressChange", windows.Errno(result))
		}
	}

	return cb, nil
}

func (callback *UnicastAddressChangeCallback) Unregister() error {

	unicastAddressChangeMutex.Lock()
	defer unicastAddressChangeMutex.Unlock()

	delete(unicastAddressChangeCallbacks, callback)

	if len(unicastAddressChangeCallbacks) < 1 && unicastAddressChangeHandle != 0 {

		result := cancelMibChangeNotify2(unicastAddressChangeHandle)

		if result != 0 {
			return os.NewSyscallError("iphlpapi.CancelMibChangeNotify2", windows.Errno(result))
		}

		unicastAddressChangeHandle = 0
	}

	return nil
}

func unicastAddressChanged(callerContext unsafe.Pointer, wtUar *wtMibUnicastipaddressRow,
	notificationType MibNotificationType) uintptr {

	interfaceLuid := uint64(0)
	var ip net.IP = nil

	if wtUar != nil {

		interfaceLuid = wtUar.InterfaceLuid

		sainet, err := wtUar.Address.toSockaddrInet()

		if err == nil && sainet != nil {
			ip = sainet.Address
		}
	}

	unicastAddressChangeMutex.Lock()

	for cb := range unicastAddressChangeCallbacks {
		cb.cb(notificationType, interfaceLuid, &ip)
	}

	unicastAddressChangeMutex.Unlock()

	return 0
}
