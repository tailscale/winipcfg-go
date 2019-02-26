/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ns-nldef-_nl_interface_offload_rod
// NL_INTERFACE_OFFLOAD_ROD defined in nldef.h
// It actually contains flags...
type wtNlInterfaceOffloadRodByte uint8

const (
	nlChecksumSupported         wtNlInterfaceOffloadRodByte = 0x01
	nlOptionsSupported          wtNlInterfaceOffloadRodByte = 0x02
	tlDatagramChecksumSupported wtNlInterfaceOffloadRodByte = 0x04
	tlStreamChecksumSupported   wtNlInterfaceOffloadRodByte = 0x08
	tlStreamOptionsSupported    wtNlInterfaceOffloadRodByte = 0x10
	fastPathCompatible          wtNlInterfaceOffloadRodByte = 0x20
	tlLargeSendOffloadSupported wtNlInterfaceOffloadRodByte = 0x40
	tlGiantSendOffloadSupported wtNlInterfaceOffloadRodByte = 0x80
)

func (wtior wtNlInterfaceOffloadRodByte) toNlInterfaceOffloadRodFlags() *NlInterfaceOffloadRodFlags {
	return &NlInterfaceOffloadRodFlags{
		NlChecksumSupported:         uint8ToBool(uint8(wtior & nlChecksumSupported)),
		NlOptionsSupported:          uint8ToBool(uint8(wtior & nlOptionsSupported)),
		TlDatagramChecksumSupported: uint8ToBool(uint8(wtior & tlDatagramChecksumSupported)),
		TlStreamChecksumSupported:   uint8ToBool(uint8(wtior & tlStreamChecksumSupported)),
		TlStreamOptionsSupported:    uint8ToBool(uint8(wtior & tlStreamOptionsSupported)),
		FastPathCompatible:          uint8ToBool(uint8(wtior & fastPathCompatible)),
		TlLargeSendOffloadSupported: uint8ToBool(uint8(wtior & tlLargeSendOffloadSupported)),
		TlGiantSendOffloadSupported: uint8ToBool(uint8(wtior & tlGiantSendOffloadSupported)),
	}
}
