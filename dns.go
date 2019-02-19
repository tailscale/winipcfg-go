/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func test() error {

	// init COM, oh yeah
	err := ole.CoInitialize(0)

	if err != nil {
		return err
	}

	defer ole.CoUninitialize()

	unknown, _ := oleutil.CreateObject("WbemScripting.SWbemLocator")
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface(ole.IID_IDispatch)
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer")
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	// result is a SWBemObjectSet
	resultRaw, _ := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_NetworkAdapterConfiguration")
	result := resultRaw.ToIDispatch()
	defer result.Release()

	countVar, _ := oleutil.GetProperty(result, "Count")
	count := int(countVar.Val)

	for i :=0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", i)
		item := itemRaw.ToIDispatch()
		defer item.Release()

		asString, _ := oleutil.GetProperty(item, "DNSDomainSuffixSearchOrder")

		arr := asString.ToArray()

		tmp := len(*arr)

		for _, itm := range arr.Array {

		}

		fmt.Println(asString.ToString())
	}

	return nil
}