/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ns-nldef-_nl_interface_offload_rod
// NL_INTERFACE_OFFLOAD_ROD defined in nldef.h
// It actually contains flags...
type wtNlInterfaceOffloadRod uint8

func (wtior wtNlInterfaceOffloadRod) toNlInterfaceOffloadRodFlags() *NlInterfaceOffloadRodFlags {
	return &NlInterfaceOffloadRodFlags{
		NlChecksumSupported:         uint8ToBool(uint8(wtior) & uint8(nlChecksumSupported)),
		NlOptionsSupported:          uint8ToBool(uint8(wtior) & uint8(nlOptionsSupported)),
		TlDatagramChecksumSupported: uint8ToBool(uint8(wtior) & uint8(tlDatagramChecksumSupported)),
		TlStreamChecksumSupported:   uint8ToBool(uint8(wtior) & uint8(tlStreamChecksumSupported)),
		TlStreamOptionsSupported:    uint8ToBool(uint8(wtior) & uint8(tlStreamOptionsSupported)),
		FastPathCompatible:          uint8ToBool(uint8(wtior) & uint8(fastPathCompatible)),
		TlLargeSendOffloadSupported: uint8ToBool(uint8(wtior) & uint8(tlLargeSendOffloadSupported)),
		TlGiantSendOffloadSupported: uint8ToBool(uint8(wtior) & uint8(tlGiantSendOffloadSupported)),
	}
}
