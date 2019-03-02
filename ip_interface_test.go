/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
	"time"
)

const (
	ipInterface_print = false
)

func interfaceChangeCallbackExample(notificationType MibNotificationType, interfaceLuid uint64) {
	fmt.Printf("INTERFACE CHANGED! MibNotificationType: %s; LUID: %d.\n", notificationType.String(),
		interfaceLuid)
}

func TestGetIpInterfaces(t *testing.T) {

	ipifcs, err := GetIpInterfaces(AF_UNSPEC)

	if err != nil {
		t.Errorf("GetIpInterfaces() returned an error: %v", err)
		return
	}

	if ipifcs == nil || len(ipifcs) < 1 {
		t.Error("GetIpInterfaces() returned nil or an empty slice.")
		return
	}

	if ipInterface_print {
		for _, ipifc := range ipifcs {
			fmt.Println("====================== INTERFACE DATA OUTPUT START ======================")
			fmt.Println(ipifc)
			fmt.Println("======================= INTERFACE DATA OUTPUT END =======================")
		}
	}
}

func TestChangeMetric(t *testing.T) {

	ipifc, err := GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
		return
	}

	if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
		return
	}

	cb, err := RegisterInterfaceChangeCallback(interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterInterfaceChangeCallback() returned error: %v", err)
		return
	}

	defer func() {
		err = UnregisterInterfaceChangeCallback(cb)

		if err != nil {
			t.Errorf("UnregisterInterfaceChangeCallback() returned error: %v", err)
		}
	}()

	useAutomaticMetric := ipifc.UseAutomaticMetric
	metric := ipifc.Metric

	newMetric := uint32(100)

	if newMetric == metric {
		newMetric = 200
	}

	ipifc.UseAutomaticMetric = false
	ipifc.Metric = newMetric

	err = ipifc.Set()

	if err != nil {
		t.Errorf("IpInterface.Set() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	ipifc, err = GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
		return
	}

	if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
		return
	}

	if ipifc.Metric != newMetric {
		t.Errorf("Expected metric: %d; actual metric: %d", newMetric, ipifc.Metric)
	}

	if ipifc.UseAutomaticMetric {
		t.Error("UseAutomaticMetric is true although it's set to false.")
	}

	ipifc.UseAutomaticMetric = useAutomaticMetric
	ipifc.Metric = metric

	err = ipifc.Set()

	if err != nil {
		t.Errorf("IpInterface.Set() returned an error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	ipifc, err = GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
		return
	}

	if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
		return
	}

	if ipifc.Metric != metric {
		t.Errorf("Expected metric: %d; actual metric: %d", metric, ipifc.Metric)
	}

	if ipifc.UseAutomaticMetric != useAutomaticMetric {
		t.Errorf("UseAutomaticMetric is %v although %v is expected.", ipifc.UseAutomaticMetric,
			useAutomaticMetric)
	}
}

func TestChangeMtu(t *testing.T) {

	ipifc, err := GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
		return
	}

	if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
		return
	}

	prevMtu := ipifc.NlMtu

	cb, err := RegisterInterfaceChangeCallback(interfaceChangeCallbackExample)

	if err != nil {
		t.Errorf("RegisterInterfaceChangeCallback() returned error: %v", err)
		return
	}

	defer func() {
		err = UnregisterInterfaceChangeCallback(cb)

		if err != nil {
			t.Errorf("UnregisterInterfaceChangeCallback() returned error: %v", err)
		}
	}()

	mtuToSet := prevMtu - 1

	ipifc.NlMtu = mtuToSet

	err = ipifc.Set()

	if err != nil {
		t.Errorf("Interface.Set() returned error: %v", err)
		return
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	ipifc, err = GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
	} else if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
	} else if ipifc.NlMtu != mtuToSet {
		t.Errorf("Interface.NlMtu is %d although %d is expected.", ipifc.NlMtu, mtuToSet)
	}

	ipifc.NlMtu = prevMtu

	err = ipifc.Set()

	if err != nil {
		t.Errorf("Interface.Set() returned error: %v", err)
	}

	// Giving some time to callbacks.
	time.Sleep(500 * time.Millisecond)

	ipifc, err = GetIpInterface(existingLuid, AF_INET)

	if err != nil {
		t.Errorf("GetIpInterface() returned an error: %v", err)
	} else if ipifc == nil {
		t.Error("GetIpInterface() returned nil.")
	} else if ipifc.NlMtu != prevMtu {
		t.Errorf("Interface.NlMtu is %d although %d is expected.", ipifc.NlMtu, prevMtu)
	}
}
