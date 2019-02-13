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
var interfaceChangeCallbackExample InterfaceChangeCallback = func (ifc *MibIpinterfaceRow, notificationType MibNotificationType) {

	fmt.Printf(`=========================== INTERFACE CHANGED START ===========================
MibNotificationType: %s
MibIpinterfaceRow:
%s
============================ INTERFACE CHANGED END ============================
`, notificationType.String(), toIndentedText(ifc.String(), "  "))

}

func TestRegisterUnregisterInterfaceChangeCallback(t *testing.T) {

	if InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered returns true although nothing is registered.")
		return;
	}

	err := RegisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterInterfaceChangeCallback returned error: %v", err)
		return;
	}

	if !InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered returns false although a callback is registered successfully.")
	}

	err = UnregisterInterfaceChangeCallback(&interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterInterfaceChangeCallback returned error: %v", err)
		return;
	}

	if InterfaceChangeCallbackRegistered(&interfaceChangeCallbackExample) {
		t.Error("InterfaceChangeCallbackRegistered returns true although the callback is unregistered successfully.")
	}
}
