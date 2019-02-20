/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
	"testing"
)

// Example callback
var unicastAddressChangeCallbackExample UnicastAddressChangeCallback = func(notificationType MibNotificationType, interfaceLuid uint64, ip *net.IP) {
	fmt.Printf("UNICAST ADDRESS CHANGED! MibNotificationType: %s; interface LUID: %d; IP: %s\n",
		notificationType.String(), interfaceLuid, ip.String())
}

func TestRegisterUnregisterUnicastAddressChangeCallback(t *testing.T) {

	if UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned true although nothing is registered.")
		return
	}

	err := RegisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterUnicastAddressChangeCallback() returned error: %v", err)
		return
	}

	if !UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned false although a callback is registered successfully.")
	}

	err = UnregisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterUnicastAddressChangeCallback() returned error: %v", err)
		return
	}

	if UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned true although the callback is unregistered successfully.")
	}
}
