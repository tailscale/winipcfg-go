/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ns-nldef-_nl_interface_offload_rod
// NL_INTERFACE_OFFLOAD_ROD defined in nldef.h
// It actually contains flags...
type wtNlInterfaceOffloadRod uint8
