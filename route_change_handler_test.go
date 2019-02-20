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
var routeChangeCallbackExample RouteChangeCallback = func(notificationType MibNotificationType, interfaceLuid uint64,
	destination *net.IPNet, nextHop *net.IP) {
	fmt.Printf("ROUTE CHANGED! MibNotificationType: %s; interface LUID: %d; destination: %s; next hop: %s\n",
		notificationType.String(), interfaceLuid, destination.String(), nextHop.String())
}

func TestRegisterUnregisterRouteChangeCallback(t *testing.T) {

	if RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered() returned true although nothing is registered.")
		return
	}

	err := RegisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterRouteChangeCallback() returned error: %v", err)
		return
	}

	if !RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered() returned false although a callback is registered successfully.")
	}

	err = UnregisterRouteChangeCallback(&routeChangeCallbackExample)

	if err != nil {
		t.Errorf("UnregisterRouteChangeCallback() returned error: %v", err)
		return
	}

	if RouteChangeCallbackRegistered(&routeChangeCallbackExample) {
		t.Error("RouteChangeCallbackRegistered() returned true although the callback is unregistered successfully.")
	}
}
