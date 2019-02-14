/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

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
		NlChecksumSupported:         uint8ToBool(uint8(wtior)&uint8(nlChecksumSupported)),
		NlOptionsSupported:          uint8ToBool(uint8(wtior)&uint8(nlOptionsSupported)),
		TlDatagramChecksumSupported: uint8ToBool(uint8(wtior)&uint8(tlDatagramChecksumSupported)),
		TlStreamChecksumSupported:   uint8ToBool(uint8(wtior)&uint8(tlStreamChecksumSupported)),
		TlStreamOptionsSupported:    uint8ToBool(uint8(wtior)&uint8(tlStreamOptionsSupported)),
		FastPathCompatible:          uint8ToBool(uint8(wtior)&uint8(fastPathCompatible)),
		TlLargeSendOffloadSupported: uint8ToBool(uint8(wtior)&uint8(tlLargeSendOffloadSupported)),
		TlGiantSendOffloadSupported: uint8ToBool(uint8(wtior)&uint8(tlGiantSendOffloadSupported)),
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

func (iorf *NlInterfaceOffloadRodFlags) String() string {

	if iorf == nil {
		return ""
	}

	return fmt.Sprintf(`NlChecksumSupported: %v        
NlOptionsSupported: %v
TlDatagramChecksumSupported: %v
TlStreamChecksumSupported: %v
TlStreamOptionsSupported: %v
FastPathCompatible: %v
TlLargeSendOffloadSupported: %v
TlGiantSendOffloadSupported: %v`, iorf.NlChecksumSupported, iorf.NlOptionsSupported, iorf.TlDatagramChecksumSupported,
		iorf.TlStreamChecksumSupported, iorf.TlStreamOptionsSupported, iorf.FastPathCompatible,
		iorf.TlLargeSendOffloadSupported, iorf.TlGiantSendOffloadSupported)
}
