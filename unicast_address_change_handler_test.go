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
var unicastAddressChangeCallbackExample UnicastAddressChangeCallback = func (uar *UnicastAddressData, notificationType MibNotificationType) {

	fmt.Printf(`======================== UNICAST ADDRESS CHANGED START ========================
MibNotificationType: %s
UnicastAddressData:
%s
========================= UNICAST ADDRESS CHANGED END =========================
`, notificationType.String(), toIndentedText(uar.String(), "  "))

}

func TestRegisterUnregisterUnicastAddressChangeCallback(t *testing.T) {

	if UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned true although nothing is registered.")
		return;
	}

	err := RegisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterUnicastAddressChangeCallback() returned error: %v", err)
		return;
	}

	if !UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned false although a callback is registered successfully.")
	}

	err = UnregisterUnicastAddressChangeCallback(&unicastAddressChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterUnicastAddressChangeCallback() returned error: %v", err)
		return;
	}

	if UnicastAddressChangeCallbackRegistered(&unicastAddressChangeCallbackExample) {
		t.Error("UnicastAddressChangeCallbackRegistered() returned true although the callback is unregistered successfully.")
	}
}
