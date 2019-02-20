/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

// Example callback
var interfaceChangeCallbackExample InterfaceChangeCallback = func(notificationType MibNotificationType, interfaceLuid uint64) {
	fmt.Printf("INTERFACE CHANGED! LUID: %d", interfaceLuid)
}

func TestRegisterUnregisterInterfaceChangeCallback(t *testing.T) {

	if InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered() returned true although nothing is registered.")
		return
	}

	err := RegisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterInterfaceChangeCallback() returned error: %v", err)
		return
	}

	if !InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered() returned false although a callback is registered successfully.")
	}

	err = UnregisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterInterfaceChangeCallback() returned error: %v", err)
		return
	}

	if InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered() returned true although the callback is unregistered successfully.")
	}
}
