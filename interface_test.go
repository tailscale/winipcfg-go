/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const (
	printInterfaceData   = false
	existingLuid         = uint64(1689399632855040) // TODO: Set an existing LUID here
	nonExistingLuid      = uint64(42)
	existingIndex        = uint32(13) // TODO: Set an existing interface index here
	nonExistingIndex     = uint32(42000000)
	existingName         = "LAN" // TODO: Set an existing interface name here
	nonExistingName      = "NON-EXISTING-NAME"
	printInterfaceRoutes = true
)

func TestGetInterfaces(t *testing.T) {
	ifcs, err := GetInterfaces()

	if err != nil {
		t.Errorf("GetInterfaces() returned error: %v", err)
	} else if ifcs == nil {
		t.Errorf("GetInterfaces() returned nil.")
	} else if printInterfaceData {
		fmt.Printf("GetInterfaces() returned %d items:\n", len(ifcs))
		for _, ifc := range ifcs {
			fmt.Println("======================== INTERFACE OUTPUT START ========================")
			fmt.Println(ifc)
			fmt.Println("========================= INTERFACE OUTPUT END =========================")
		}
	}
}

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
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromLUID() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromLUIDNonExisting(t *testing.T) {
	ifc, err := InterfaceFromLUID(nonExistingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromLUID() returned an interface with LUID=%d, although requested LUID was %d.",
			ifc.Luid, nonExistingLuid)
	}
}

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
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromIndex() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromIndexNonExisting(t *testing.T) {
	ifc, err := InterfaceFromIndex(nonExistingIndex)

	if err != nil {
		t.Errorf("InterfaceFromIndex() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromIndex() returned an interface with index=%d, although requested index was %d.",
			ifc.Index, nonExistingIndex)
	}
}

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
	} else if printInterfaceData {
		fmt.Println("======================== INTERFACE OUTPUT START ========================")
		fmt.Printf("InterfaceFromFriendlyName() returned corresponding interface:\n%s\n", ifc)
		fmt.Println("========================= INTERFACE OUTPUT END =========================")
	}
}

func TestInterfaceFromFriendlyNameNonExisting(t *testing.T) {
	ifc, err := InterfaceFromFriendlyName(nonExistingName)

	if err != nil {
		t.Errorf("InterfaceFromFriendlyName() returned error: %v", err)
	} else if ifc != nil {
		t.Errorf("InterfaceFromFriendlyName() returned an interface with name=%s, although requested name was %s.",
			ifc.FriendlyName, nonExistingName)
	}
}

func TestInterface_GetRoutes(t *testing.T) {
	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so Interface.GetRoutes() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so Interface.GetRoutes() testing cannot be performed.")
		return
	}

	routes, err := ifc.GetRoutes(AF_UNSPEC)

	if err != nil {
		t.Errorf("Interface.GetRoutes() returned an error: %v", err)
		return
	}

	if routes == nil || len(routes) < 1 {
		t.Error("Interface.GetRoutes() returned nil or empty slice.")
		return
	}

	for _, route := range routes {
		if route.InterfaceLuid != ifc.Luid {
			t.Errorf("Interface.GetRoutes() retuned a route with a wrong LUID. Interface.Luid: %d; Route.InterfaceLuid: %d.",
				ifc.Luid, route.InterfaceLuid)
		}
	}

	if printInterfaceRoutes {
		for _, route := range routes {
			fmt.Println("========================== ROUTE OUTPUT START ==========================")
			fmt.Println(route)
			fmt.Println("=========================== ROUTE OUTPUT END ===========================")
		}
	}
}
