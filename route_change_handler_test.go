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
var routeChangeCallbackExample RouteChangeCallback = func (route *Route, notificationType MibNotificationType) {

	fmt.Printf(`============================= ROUTE CHANGED START =============================
MibNotificationType: %s
Route:
%s
============================== ROUTE CHANGED END ==============================
`, notificationType.String(), toIndentedText(route.String(), "  "))

}

func TestRegisterUnregisterRouteChangeCallback(t *testing.T) {

	if RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered returns true although nothing is registered.")
		return;
	}

	err := RegisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterRouteChangeCallback returned error: %v", err)
		return;
	}

	if !RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered returns false although a callback is registered successfully.")
	}

	err = UnregisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterRouteChangeCallback returned error: %v", err)
		return;
	}

	if RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered returns true although the callback is unregistered successfully.")
	}
}
