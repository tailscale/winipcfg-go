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

		asString, err := oleutil.GetProperty(item, "DNSDomainSuffixSearchOrder")

		if err != nil {
			fmt.Printf("GetProperty() returned an error: %v\n", err)
			continue
		}

		if asString == nil {
			fmt.Println("GetProperty() returned nil.")
		}

		arr := asString.ToArray()

		if arr == nil {
			fmt.Println("ToArray() returned nil.")
			continue
		}

		totalElements, err := arr.TotalElements(0)

		if err != nil {
			fmt.Printf("TotalElements() returned an error: %v\n", err)
			continue
		}

		fmt.Printf("Total elements: %d\n", totalElements)

		if totalElements > 0 {

			fmt.Println("About to do!")
			stringArray := arr.ToValueArray()
			fmt.Println("Did it!")

			if stringArray == nil {
				fmt.Println("ToStringArray() returned nil.")
			} else {
				for _, itm := range stringArray {
					fmt.Println(itm)
				}
			}
		}

		//fmt.Println(asString.ToString())
	}

	return nil
}