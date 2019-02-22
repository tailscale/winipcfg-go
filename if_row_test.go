/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const ifRow_print = true

func TestGetIfRow(t *testing.T) {

	ifc, err := InterfaceFromLUID(existingLuid)

	if err != nil {
		t.Errorf("InterfaceFromLUID() returned an error (%v), so GetIfRow() testing cannot be performed.",
			err)
		return
	}

	if ifc == nil {
		t.Error("InterfaceFromLUID() returned nil, so GetIfRow() testing cannot be performed.")
		return
	}

	ifrow, err := GetIfRow(existingLuid, MibIfEntryNormal)

	if err != nil {
		t.Errorf("GetIfRow() returned an error: %v", err)
		return
	}

	if ifrow == nil {
		t.Error("GetIfRow() returned nil.")
		return
	}

	if ifRow_print {
		fmt.Println("========================== IfRow OUTPUT START ==========================")
		fmt.Println(ifrow)
		fmt.Println("=========================== IfRow OUTPUT END ===========================")
	}
}

func TestGetIfRows(t *testing.T) {

	rows, err := GetIfRows(MibIfEntryNormal)

	if err != nil {
		t.Errorf("GetIfRows() returned an error: %v", err)
	} else if rows == nil || len(rows) < 1 {
		t.Error("GetIfRows() returned nil or an empty slice.")
	} else if ifRow_print {
		for _, row := range rows {
			fmt.Println("========================== IfRow OUTPUT START ==========================")
			fmt.Println(row)
			fmt.Println("=========================== IfRow OUTPUT END ===========================")
		}
	}
}
