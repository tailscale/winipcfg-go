/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

func TestGetInterfaces(t *testing.T) {
	ifcs, err := GetInterfaces()

	if err != nil {
		t.Errorf("GetInterfaces() returned error: %v", err)
	} else if ifcs == nil {
		t.Errorf("GetInterfaces() returned nil.")
	} else {
		fmt.Printf("GetInterfaces() returned %d items:\n", len(ifcs))
		for _, ifc := range ifcs {
			fmt.Println(ifc)
		}
	}
}

// TODO: Set an existing LUID here:
const existingLuid uint64 = 1689399632855040

func TestInterfaceFromLUIDExisting(t *testing.T) {
	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromLUID() returned nil for luid=%d. Have you set existingLuid constant?",
			existingLuid)
	} else if ifc.Luid != existingLuid {
		t.Errorf("InterfaceFromLUID() returned interface with a wrong LUID. Requested: %d; returned: %d.",
			existingLuid, ifc.Luid)
	} else {
		fmt.Printf("InterfaceFromLUID() returned corresponding interface:\n%s\n", ifc)
	}
}

const nonExistingLuid uint64 = 42

func TestInterfaceFromLUIDNonExisting(t *testing.T) {
	ifc, err := InterfaceFromLUID(nonExistingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromLUID() returned an interface with LUID=%d, although requested LUID was %d.",
			ifc.Luid, nonExistingLuid)
	}
}

// TODO: Set an existing interface index here:
const existingIndex uint32 = 13

func TestInterfaceFromIndexExisting(t *testing.T) {
	ifc, err := InterfaceFromIndex(existingIndex)

	if err != nil {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromIndex() returned nil for index=%d. Have you set existingIndex constant?",
			existingIndex)
	} else if uint32(ifc.Index) != existingIndex {
		t.Errorf("InterfaceFromIndex() returned interface with a wrong index. Requested: %d; returned: %d.",
			existingIndex, ifc.Index)
	} else {
		fmt.Printf("InterfaceFromIndex() returned corresponding interface:\n%s\n", ifc)
	}
}

const nonExistingIndex uint32 = 42000000

func TestInterfaceFromIndexNonExisting(t *testing.T) {
	ifc, err := InterfaceFromIndex(nonExistingIndex)

	if err != nil {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromIndex() returned an interface with index=%d, although requested index was %d.",
			ifc.Index, nonExistingIndex)
	}
}

// TODO: Set an existing interface name here:
const existingName string = "LAN"

func TestInterfaceFromFriendlyNameExisting(t *testing.T) {
	ifc, err := InterfaceFromFriendlyName(existingName)

	if err != nil {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	} else if ifc == nil {
		t.Errorf("InterfaceFromFriendlyName() returned nil for name=%s. Have you set existingName constant?",
			existingName)
	} else if ifc.FriendlyName != existingName {
		t.Errorf("InterfaceFromFriendlyName() returned interface with a wrong name. Requested: %s; returned: %s.",
			existingName, ifc.FriendlyName)
	} else {
		fmt.Printf("InterfaceFromFriendlyName() returned corresponding interface:\n%s\n", ifc)
	}
}

const nonExistingName string = "NON-EXISTING-NAME"

func TestInterfaceFromFriendlyNameNonExisting(t *testing.T) {
	ifc, err := InterfaceFromFriendlyName(nonExistingName)

	if err != nil {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromFriendlyName() returned an interface with name=%s, although requested name was %s.",
			ifc.FriendlyName, nonExistingName)
	}
}
