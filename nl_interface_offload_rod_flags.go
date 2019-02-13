/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

type nlInterfaceOffloadRod uint8

const (
	nlChecksumSupported         nlInterfaceOffloadRod = 1
	nlOptionsSupported          nlInterfaceOffloadRod = 2
	tlDatagramChecksumSupported nlInterfaceOffloadRod = 4
	tlStreamChecksumSupported   nlInterfaceOffloadRod = 8
	tlStreamOptionsSupported    nlInterfaceOffloadRod = 16
	fastPathCompatible          nlInterfaceOffloadRod = 32
	tlLargeSendOffloadSupported nlInterfaceOffloadRod = 64
	tlGiantSendOffloadSupported nlInterfaceOffloadRod = 128
)

type NlInterfaceOffloadRodFlags struct {
	NlChecksumSupported         bool
	NlOptionsSupported          bool
	TlDatagramChecksumSupported bool
	TlStreamChecksumSupported   bool
	TlStreamOptionsSupported    bool
	FastPathCompatible          bool
	TlLargeSendOffloadSupported bool
	TlGiantSendOffloadSupported bool
}

func (wtior wtNlInterfaceOffloadRod) toNlInterfaceOffloadRodFlags() *NlInterfaceOffloadRodFlags {
	return &NlInterfaceOffloadRodFlags{
		NlChecksumSupported:         uint8(wtior)&uint8(nlChecksumSupported) != 0,
		NlOptionsSupported:          uint8(wtior)&uint8(nlOptionsSupported) != 0,
		TlDatagramChecksumSupported: uint8(wtior)&uint8(tlDatagramChecksumSupported) != 0,
		TlStreamChecksumSupported:   uint8(wtior)&uint8(tlStreamChecksumSupported) != 0,
		TlStreamOptionsSupported:    uint8(wtior)&uint8(tlStreamOptionsSupported) != 0,
		FastPathCompatible:          uint8(wtior)&uint8(fastPathCompatible) != 0,
		TlLargeSendOffloadSupported: uint8(wtior)&uint8(tlLargeSendOffloadSupported) != 0,
		TlGiantSendOffloadSupported: uint8(wtior)&uint8(tlGiantSendOffloadSupported) != 0,
	}
}

func (ior *NlInterfaceOffloadRodFlags) toWtNlInterfaceOffloadRod() wtNlInterfaceOffloadRod {

	uint8Val := uint8(0)

	if ior.NlChecksumSupported {
		uint8Val |= uint8(nlChecksumSupported)
	}

	if ior.NlOptionsSupported {
		uint8Val |= uint8(nlOptionsSupported)
	}

	if ior.TlDatagramChecksumSupported {
		uint8Val |= uint8(tlDatagramChecksumSupported)
	}

	if ior.TlStreamChecksumSupported {
		uint8Val |= uint8(tlStreamChecksumSupported)
	}

	if ior.TlStreamOptionsSupported {
		uint8Val |= uint8(tlStreamOptionsSupported)
	}

	if ior.FastPathCompatible {
		uint8Val |= uint8(fastPathCompatible)
	}

	if ior.TlLargeSendOffloadSupported {
		uint8Val |= uint8(tlLargeSendOffloadSupported)
	}

	if ior.TlGiantSendOffloadSupported {
		uint8Val |= uint8(tlGiantSendOffloadSupported)
	}

	return wtNlInterfaceOffloadRod(uint8Val)
}
