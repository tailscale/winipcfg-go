/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

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

func (iorf *NlInterfaceOffloadRodFlags) toWtNlInterfaceOffloadRod() wtNlInterfaceOffloadRodByte {

	result := wtNlInterfaceOffloadRodByte(0)

	if iorf.NlChecksumSupported {
		result |= nlChecksumSupported
	}

	if iorf.NlOptionsSupported {
		result |= nlOptionsSupported
	}

	if iorf.TlDatagramChecksumSupported {
		result |= tlDatagramChecksumSupported
	}

	if iorf.TlStreamChecksumSupported {
		result |= tlStreamChecksumSupported
	}

	if iorf.TlStreamOptionsSupported {
		result |= tlStreamOptionsSupported
	}

	if iorf.FastPathCompatible {
		result |= fastPathCompatible
	}

	if iorf.TlLargeSendOffloadSupported {
		result |= tlLargeSendOffloadSupported
	}

	if iorf.TlGiantSendOffloadSupported {
		result |= tlGiantSendOffloadSupported
	}

	return result
}

func (iorf *NlInterfaceOffloadRodFlags) String() string {

	if iorf == nil {
		return "<nil>"
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
