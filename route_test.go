/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"testing"
)

const route_print bool = false

func TestGetRoutes(t *testing.T) {

	routes, err := GetRoutes(AF_UNSPEC)

	if err != nil {
		t.Errorf("GetRoutes() returned an error: %v", err)
		return
	}

	if routes == nil || len(routes) < 1 {
		t.Error("GetRoutes() returned nil or empty slice.")
		return
	}

	if route_print {
		for _, route := range routes {
			fmt.Println("========================== ROUTE OUTPUT START ==========================")
			fmt.Println(route)
			fmt.Println("=========================== ROUTE OUTPUT END ===========================")
		}
	}
}
