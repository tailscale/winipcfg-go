/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package wintypes

import "fmt"

// Defined in ipifcons.h
type IFTYPE ULONG

const (
	IF_TYPE_OTHER                            IFTYPE = 1 // None of the below
	IF_TYPE_REGULAR_1822                     IFTYPE = 2
	IF_TYPE_HDH_1822                         IFTYPE = 3
	IF_TYPE_DDN_X25                          IFTYPE = 4
	IF_TYPE_RFC877_X25                       IFTYPE = 5
	IF_TYPE_ETHERNET_CSMACD                  IFTYPE = 6
	IF_TYPE_IS088023_CSMACD                  IFTYPE = 7
	IF_TYPE_ISO88024_TOKENBUS                IFTYPE = 8
	IF_TYPE_ISO88025_TOKENRING               IFTYPE = 9
	IF_TYPE_ISO88026_MAN                     IFTYPE = 10
	IF_TYPE_STARLAN                          IFTYPE = 11
	IF_TYPE_PROTEON_10MBIT                   IFTYPE = 12
	IF_TYPE_PROTEON_80MBIT                   IFTYPE = 13
	IF_TYPE_HYPERCHANNEL                     IFTYPE = 14
	IF_TYPE_FDDI                             IFTYPE = 15
	IF_TYPE_LAP_B                            IFTYPE = 16
	IF_TYPE_SDLC                             IFTYPE = 17
	IF_TYPE_DS1                              IFTYPE = 18 // DS1-MIB
	IF_TYPE_E1                               IFTYPE = 19 // Obsolete; see DS1-MIB
	IF_TYPE_BASIC_ISDN                       IFTYPE = 20
	IF_TYPE_PRIMARY_ISDN                     IFTYPE = 21
	IF_TYPE_PROP_POINT2POINT_SERIAL          IFTYPE = 22 // proprietary serial
	IF_TYPE_PPP                              IFTYPE = 23
	IF_TYPE_SOFTWARE_LOOPBACK                IFTYPE = 24
	IF_TYPE_EON                              IFTYPE = 25 // CLNP over IP
	IF_TYPE_ETHERNET_3MBIT                   IFTYPE = 26
	IF_TYPE_NSIP                             IFTYPE = 27 // XNS over IP
	IF_TYPE_SLIP                             IFTYPE = 28 // Generic Slip
	IF_TYPE_ULTRA                            IFTYPE = 29 // ULTRA Technologies
	IF_TYPE_DS3                              IFTYPE = 30 // DS3-MIB
	IF_TYPE_SIP                              IFTYPE = 31 // SMDS, coffee
	IF_TYPE_FRAMERELAY                       IFTYPE = 32 // DTE only
	IF_TYPE_RS232                            IFTYPE = 33
	IF_TYPE_PARA                             IFTYPE = 34 // Parallel port
	IF_TYPE_ARCNET                           IFTYPE = 35
	IF_TYPE_ARCNET_PLUS                      IFTYPE = 36
	IF_TYPE_ATM                              IFTYPE = 37 // ATM cells
	IF_TYPE_MIO_X25                          IFTYPE = 38
	IF_TYPE_SONET                            IFTYPE = 39 // SONET or SDH
	IF_TYPE_X25_PLE                          IFTYPE = 40
	IF_TYPE_ISO88022_LLC                     IFTYPE = 41
	IF_TYPE_LOCALTALK                        IFTYPE = 42
	IF_TYPE_SMDS_DXI                         IFTYPE = 43
	IF_TYPE_FRAMERELAY_SERVICE               IFTYPE = 44 // FRNETSERV-MIB
	IF_TYPE_V35                              IFTYPE = 45
	IF_TYPE_HSSI                             IFTYPE = 46
	IF_TYPE_HIPPI                            IFTYPE = 47
	IF_TYPE_MODEM                            IFTYPE = 48 // Generic Modem
	IF_TYPE_AAL5                             IFTYPE = 49 // AAL5 over ATM
	IF_TYPE_SONET_PATH                       IFTYPE = 50
	IF_TYPE_SONET_VT                         IFTYPE = 51
	IF_TYPE_SMDS_ICIP                        IFTYPE = 52 // SMDS InterCarrier Interface
	IF_TYPE_PROP_VIRTUAL                     IFTYPE = 53 // Proprietary virtual/wintypes
	IF_TYPE_PROP_MULTIPLEXOR                 IFTYPE = 54 // Proprietary multiplexing
	IF_TYPE_IEEE80212                        IFTYPE = 55 // 100BaseVG
	IF_TYPE_FIBRECHANNEL                     IFTYPE = 56
	IF_TYPE_HIPPIINTERFACE                   IFTYPE = 57
	IF_TYPE_FRAMERELAY_INTERCONNECT          IFTYPE = 58 // Obsolete, use 32 or 44
	IF_TYPE_AFLANE_8023                      IFTYPE = 59 // ATM Emulated LAN for 802.3
	IF_TYPE_AFLANE_8025                      IFTYPE = 60 // ATM Emulated LAN for 802.5
	IF_TYPE_CCTEMUL                          IFTYPE = 61 // ATM Emulated circuit
	IF_TYPE_FASTETHER                        IFTYPE = 62 // Fast Ethernet (100BaseT)
	IF_TYPE_ISDN                             IFTYPE = 63 // ISDN and X.25
	IF_TYPE_V11                              IFTYPE = 64 // CCITT V.11/X.21
	IF_TYPE_V36                              IFTYPE = 65 // CCITT V.36
	IF_TYPE_G703_64K                         IFTYPE = 66 // CCITT G703 at 64Kbps
	IF_TYPE_G703_2MB                         IFTYPE = 67 // Obsolete; see DS1-MIB
	IF_TYPE_QLLC                             IFTYPE = 68 // SNA QLLC
	IF_TYPE_FASTETHER_FX                     IFTYPE = 69 // Fast Ethernet (100BaseFX)
	IF_TYPE_CHANNEL                          IFTYPE = 70
	IF_TYPE_IEEE80211                        IFTYPE = 71  // Radio spread spectrum
	IF_TYPE_IBM370PARCHAN                    IFTYPE = 72  // IBM System 360/370 OEMI Channel
	IF_TYPE_ESCON                            IFTYPE = 73  // IBM Enterprise Systems Connection
	IF_TYPE_DLSW                             IFTYPE = 74  // Data Link Switching
	IF_TYPE_ISDN_S                           IFTYPE = 75  // ISDN S/T interface
	IF_TYPE_ISDN_U                           IFTYPE = 76  // ISDN U interface
	IF_TYPE_LAP_D                            IFTYPE = 77  // Link Access Protocol D
	IF_TYPE_IPSWITCH                         IFTYPE = 78  // IP Switching Objects
	IF_TYPE_RSRB                             IFTYPE = 79  // Remote Source Route Bridging
	IF_TYPE_ATM_LOGICAL                      IFTYPE = 80  // ATM Logical Port
	IF_TYPE_DS0                              IFTYPE = 81  // Digital Signal Level 0
	IF_TYPE_DS0_BUNDLE                       IFTYPE = 82  // Group of ds0s on the same ds1
	IF_TYPE_BSC                              IFTYPE = 83  // Bisynchronous Protocol
	IF_TYPE_ASYNC                            IFTYPE = 84  // Asynchronous Protocol
	IF_TYPE_CNR                              IFTYPE = 85  // Combat Net Radio
	IF_TYPE_ISO88025R_DTR                    IFTYPE = 86  // ISO 802.5r DTR
	IF_TYPE_EPLRS                            IFTYPE = 87  // Ext Pos Loc Report Sys
	IF_TYPE_ARAP                             IFTYPE = 88  // Appletalk Remote Access Protocol
	IF_TYPE_PROP_CNLS                        IFTYPE = 89  // Proprietary Connectionless Proto
	IF_TYPE_HOSTPAD                          IFTYPE = 90  // CCITT-ITU X.29 PAD Protocol
	IF_TYPE_TERMPAD                          IFTYPE = 91  // CCITT-ITU X.3 PAD Facility
	IF_TYPE_FRAMERELAY_MPI                   IFTYPE = 92  // Multiproto Interconnect over FR
	IF_TYPE_X213                             IFTYPE = 93  // CCITT-ITU X213
	IF_TYPE_ADSL                             IFTYPE = 94  // Asymmetric Digital Subscrbr Loop
	IF_TYPE_RADSL                            IFTYPE = 95  // Rate-Adapt Digital Subscrbr Loop
	IF_TYPE_SDSL                             IFTYPE = 96  // Symmetric Digital Subscriber Loop
	IF_TYPE_VDSL                             IFTYPE = 97  // Very H-Speed Digital Subscrb Loop
	IF_TYPE_ISO88025_CRFPRINT                IFTYPE = 98  // ISO 802.5 CRFP
	IF_TYPE_MYRINET                          IFTYPE = 99  // Myricom Myrinet
	IF_TYPE_VOICE_EM                         IFTYPE = 100 // Voice recEive and transMit
	IF_TYPE_VOICE_FXO                        IFTYPE = 101 // Voice Foreign Exchange Office
	IF_TYPE_VOICE_FXS                        IFTYPE = 102 // Voice Foreign Exchange Station
	IF_TYPE_VOICE_ENCAP                      IFTYPE = 103 // Voice encapsulation
	IF_TYPE_VOICE_OVERIP                     IFTYPE = 104 // Voice over IP encapsulation
	IF_TYPE_ATM_DXI                          IFTYPE = 105 // ATM DXI
	IF_TYPE_ATM_FUNI                         IFTYPE = 106 // ATM FUNI
	IF_TYPE_ATM_IMA                          IFTYPE = 107 // ATM IMA
	IF_TYPE_PPPMULTILINKBUNDLE               IFTYPE = 108 // PPP Multilink Bundle
	IF_TYPE_IPOVER_CDLC                      IFTYPE = 109 // IBM ipOverCdlc
	IF_TYPE_IPOVER_CLAW                      IFTYPE = 110 // IBM Common Link Access to Workstn
	IF_TYPE_STACKTOSTACK                     IFTYPE = 111 // IBM stackToStack
	IF_TYPE_VIRTUALIPADDRESS                 IFTYPE = 112 // IBM VIPA
	IF_TYPE_MPC                              IFTYPE = 113 // IBM multi-proto channel support
	IF_TYPE_IPOVER_ATM                       IFTYPE = 114 // IBM ipOverAtm
	IF_TYPE_ISO88025_FIBER                   IFTYPE = 115 // ISO 802.5j Fiber Token Ring
	IF_TYPE_TDLC                             IFTYPE = 116 // IBM twinaxial data link control
	IF_TYPE_GIGABITETHERNET                  IFTYPE = 117
	IF_TYPE_HDLC                             IFTYPE = 118
	IF_TYPE_LAP_F                            IFTYPE = 119
	IF_TYPE_V37                              IFTYPE = 120
	IF_TYPE_X25_MLP                          IFTYPE = 121 // Multi-Link Protocol
	IF_TYPE_X25_HUNTGROUP                    IFTYPE = 122 // X.25 Hunt Group
	IF_TYPE_TRANSPHDLC                       IFTYPE = 123
	IF_TYPE_INTERLEAVE                       IFTYPE = 124 // Interleave channel
	IF_TYPE_FAST                             IFTYPE = 125 // Fast channel
	IF_TYPE_IP                               IFTYPE = 126 // IP (for APPN HPR in IP networks)
	IF_TYPE_DOCSCABLE_MACLAYER               IFTYPE = 127 // CATV Mac Layer
	IF_TYPE_DOCSCABLE_DOWNSTREAM             IFTYPE = 128 // CATV Downstream interface
	IF_TYPE_DOCSCABLE_UPSTREAM               IFTYPE = 129 // CATV Upstream interface
	IF_TYPE_A12MPPSWITCH                     IFTYPE = 130 // Avalon Parallel Processor
	IF_TYPE_TUNNEL                           IFTYPE = 131 // Encapsulation interface
	IF_TYPE_COFFEE                           IFTYPE = 132 // Coffee pot
	IF_TYPE_CES                              IFTYPE = 133 // Circuit Emulation Service
	IF_TYPE_ATM_SUBINTERFACE                 IFTYPE = 134 // ATM Sub Interface
	IF_TYPE_L2_VLAN                          IFTYPE = 135 // Layer 2 Virtual LAN using 802.1Q
	IF_TYPE_L3_IPVLAN                        IFTYPE = 136 // Layer 3 Virtual LAN using IP
	IF_TYPE_L3_IPXVLAN                       IFTYPE = 137 // Layer 3 Virtual LAN using IPX
	IF_TYPE_DIGITALPOWERLINE                 IFTYPE = 138 // IP over Power Lines
	IF_TYPE_MEDIAMAILOVERIP                  IFTYPE = 139 // Multimedia Mail over IP
	IF_TYPE_DTM                              IFTYPE = 140 // Dynamic syncronous Transfer Mode
	IF_TYPE_DCN                              IFTYPE = 141 // Data Communications Network
	IF_TYPE_IPFORWARD                        IFTYPE = 142 // IP Forwarding Interface
	IF_TYPE_MSDSL                            IFTYPE = 143 // Multi-rate Symmetric DSL
	IF_TYPE_IEEE1394                         IFTYPE = 144 // IEEE1394 High Perf Serial Bus
	IF_TYPE_IF_GSN                           IFTYPE = 145
	IF_TYPE_DVBRCC_MACLAYER                  IFTYPE = 146
	IF_TYPE_DVBRCC_DOWNSTREAM                IFTYPE = 147
	IF_TYPE_DVBRCC_UPSTREAM                  IFTYPE = 148
	IF_TYPE_ATM_VIRTUAL                      IFTYPE = 149
	IF_TYPE_MPLS_TUNNEL                      IFTYPE = 150
	IF_TYPE_SRP                              IFTYPE = 151
	IF_TYPE_VOICEOVERATM                     IFTYPE = 152
	IF_TYPE_VOICEOVERFRAMERELAY              IFTYPE = 153
	IF_TYPE_IDSL                             IFTYPE = 154
	IF_TYPE_COMPOSITELINK                    IFTYPE = 155
	IF_TYPE_SS7_SIGLINK                      IFTYPE = 156
	IF_TYPE_PROP_WIRELESS_P2P                IFTYPE = 157
	IF_TYPE_FR_FORWARD                       IFTYPE = 158
	IF_TYPE_RFC1483                          IFTYPE = 159
	IF_TYPE_USB                              IFTYPE = 160
	IF_TYPE_IEEE8023AD_LAG                   IFTYPE = 161
	IF_TYPE_BGP_POLICY_ACCOUNTING            IFTYPE = 162
	IF_TYPE_FRF16_MFR_BUNDLE                 IFTYPE = 163
	IF_TYPE_H323_GATEKEEPER                  IFTYPE = 164
	IF_TYPE_H323_PROXY                       IFTYPE = 165
	IF_TYPE_MPLS                             IFTYPE = 166
	IF_TYPE_MF_SIGLINK                       IFTYPE = 167
	IF_TYPE_HDSL2                            IFTYPE = 168
	IF_TYPE_SHDSL                            IFTYPE = 169
	IF_TYPE_DS1_FDL                          IFTYPE = 170
	IF_TYPE_POS                              IFTYPE = 171
	IF_TYPE_DVB_ASI_IN                       IFTYPE = 172
	IF_TYPE_DVB_ASI_OUT                      IFTYPE = 173
	IF_TYPE_PLC                              IFTYPE = 174
	IF_TYPE_NFAS                             IFTYPE = 175
	IF_TYPE_TR008                            IFTYPE = 176
	IF_TYPE_GR303_RDT                        IFTYPE = 177
	IF_TYPE_GR303_IDT                        IFTYPE = 178
	IF_TYPE_ISUP                             IFTYPE = 179
	IF_TYPE_PROP_DOCS_WIRELESS_MACLAYER      IFTYPE = 180
	IF_TYPE_PROP_DOCS_WIRELESS_DOWNSTREAM    IFTYPE = 181
	IF_TYPE_PROP_DOCS_WIRELESS_UPSTREAM      IFTYPE = 182
	IF_TYPE_HIPERLAN2                        IFTYPE = 183
	IF_TYPE_PROP_BWA_P2MP                    IFTYPE = 184
	IF_TYPE_SONET_OVERHEAD_CHANNEL           IFTYPE = 185
	IF_TYPE_DIGITAL_WRAPPER_OVERHEAD_CHANNEL IFTYPE = 186
	IF_TYPE_AAL2                             IFTYPE = 187
	IF_TYPE_RADIO_MAC                        IFTYPE = 188
	IF_TYPE_ATM_RADIO                        IFTYPE = 189
	IF_TYPE_IMT                              IFTYPE = 190
	IF_TYPE_MVL                              IFTYPE = 191
	IF_TYPE_REACH_DSL                        IFTYPE = 192
	IF_TYPE_FR_DLCI_ENDPT                    IFTYPE = 193
	IF_TYPE_ATM_VCI_ENDPT                    IFTYPE = 194
	IF_TYPE_OPTICAL_CHANNEL                  IFTYPE = 195
	IF_TYPE_OPTICAL_TRANSPORT                IFTYPE = 196
	IF_TYPE_IEEE80216_WMAN                   IFTYPE = 237
	IF_TYPE_WWANPP                           IFTYPE = 243 // WWAN devices based on GSM technology
	IF_TYPE_WWANPP2                          IFTYPE = 244 // WWAN devices based on CDMA technology
	IF_TYPE_IEEE802154                       IFTYPE = 259 // IEEE 802.15.4 WPAN interface
	IF_TYPE_XBOX_WIRELESS                    IFTYPE = 281
)

func (t IFTYPE) String() string {
	switch t {
	case IF_TYPE_OTHER:
		return "IF_TYPE_OTHER"
	case IF_TYPE_REGULAR_1822:
		return "IF_TYPE_REGULAR_1822"
	case IF_TYPE_HDH_1822:
		return "IF_TYPE_HDH_1822"
	case IF_TYPE_DDN_X25:
		return "IF_TYPE_DDN_X25"
	case IF_TYPE_RFC877_X25:
		return "IF_TYPE_RFC877_X25"
	case IF_TYPE_ETHERNET_CSMACD:
		return "IF_TYPE_ETHERNET_CSMACD"
	case IF_TYPE_IS088023_CSMACD:
		return "IF_TYPE_IS088023_CSMACD"
	case IF_TYPE_ISO88024_TOKENBUS:
		return "IF_TYPE_ISO88024_TOKENBUS"
	case IF_TYPE_ISO88025_TOKENRING:
		return "IF_TYPE_ISO88025_TOKENRING"
	case IF_TYPE_ISO88026_MAN:
		return "IF_TYPE_ISO88026_MAN"
	case IF_TYPE_STARLAN:
		return "IF_TYPE_STARLAN"
	case IF_TYPE_PROTEON_10MBIT:
		return "IF_TYPE_PROTEON_10MBIT"
	case IF_TYPE_PROTEON_80MBIT:
		return "IF_TYPE_PROTEON_80MBIT"
	case IF_TYPE_HYPERCHANNEL:
		return "IF_TYPE_HYPERCHANNEL"
	case IF_TYPE_FDDI:
		return "IF_TYPE_FDDI"
	case IF_TYPE_LAP_B:
		return "IF_TYPE_LAP_B"
	case IF_TYPE_SDLC:
		return "IF_TYPE_SDLC"
	case IF_TYPE_DS1:
		return "IF_TYPE_DS1"
	case IF_TYPE_E1:
		return "IF_TYPE_E1"
	case IF_TYPE_BASIC_ISDN:
		return "IF_TYPE_BASIC_ISDN"
	case IF_TYPE_PRIMARY_ISDN:
		return "IF_TYPE_PRIMARY_ISDN"
	case IF_TYPE_PROP_POINT2POINT_SERIAL:
		return "IF_TYPE_PROP_POINT2POINT_SERIAL"
	case IF_TYPE_PPP:
		return "IF_TYPE_PPP"
	case IF_TYPE_SOFTWARE_LOOPBACK:
		return "IF_TYPE_SOFTWARE_LOOPBACK"
	case IF_TYPE_EON:
		return "IF_TYPE_EON"
	case IF_TYPE_ETHERNET_3MBIT:
		return "IF_TYPE_ETHERNET_3MBIT"
	case IF_TYPE_NSIP:
		return "IF_TYPE_NSIP"
	case IF_TYPE_SLIP:
		return "IF_TYPE_SLIP"
	case IF_TYPE_ULTRA:
		return "IF_TYPE_ULTRA"
	case IF_TYPE_DS3:
		return "IF_TYPE_DS3"
	case IF_TYPE_SIP:
		return "IF_TYPE_SIP"
	case IF_TYPE_FRAMERELAY:
		return "IF_TYPE_FRAMERELAY"
	case IF_TYPE_RS232:
		return "IF_TYPE_RS232"
	case IF_TYPE_PARA:
		return "IF_TYPE_PARA"
	case IF_TYPE_ARCNET:
		return "IF_TYPE_ARCNET"
	case IF_TYPE_ARCNET_PLUS:
		return "IF_TYPE_ARCNET_PLUS"
	case IF_TYPE_ATM:
		return "IF_TYPE_ATM"
	case IF_TYPE_MIO_X25:
		return "IF_TYPE_MIO_X25"
	case IF_TYPE_SONET:
		return "IF_TYPE_SONET"
	case IF_TYPE_X25_PLE:
		return "IF_TYPE_X25_PLE"
	case IF_TYPE_ISO88022_LLC:
		return "IF_TYPE_ISO88022_LLC"
	case IF_TYPE_LOCALTALK:
		return "IF_TYPE_LOCALTALK"
	case IF_TYPE_SMDS_DXI:
		return "IF_TYPE_SMDS_DXI"
	case IF_TYPE_FRAMERELAY_SERVICE:
		return "IF_TYPE_FRAMERELAY_SERVICE"
	case IF_TYPE_V35:
		return "IF_TYPE_V35"
	case IF_TYPE_HSSI:
		return "IF_TYPE_HSSI"
	case IF_TYPE_HIPPI:
		return "IF_TYPE_HIPPI"
	case IF_TYPE_MODEM:
		return "IF_TYPE_MODEM"
	case IF_TYPE_AAL5:
		return "IF_TYPE_AAL5"
	case IF_TYPE_SONET_PATH:
		return "IF_TYPE_SONET_PATH"
	case IF_TYPE_SONET_VT:
		return "IF_TYPE_SONET_VT"
	case IF_TYPE_SMDS_ICIP:
		return "IF_TYPE_SMDS_ICIP"
	case IF_TYPE_PROP_VIRTUAL:
		return "IF_TYPE_PROP_VIRTUAL"
	case IF_TYPE_PROP_MULTIPLEXOR:
		return "IF_TYPE_PROP_MULTIPLEXOR"
	case IF_TYPE_IEEE80212:
		return "IF_TYPE_IEEE80212"
	case IF_TYPE_FIBRECHANNEL:
		return "IF_TYPE_FIBRECHANNEL"
	case IF_TYPE_HIPPIINTERFACE:
		return "IF_TYPE_HIPPIINTERFACE"
	case IF_TYPE_FRAMERELAY_INTERCONNECT:
		return "IF_TYPE_FRAMERELAY_INTERCONNECT"
	case IF_TYPE_AFLANE_8023:
		return "IF_TYPE_AFLANE_8023"
	case IF_TYPE_AFLANE_8025:
		return "IF_TYPE_AFLANE_8025"
	case IF_TYPE_CCTEMUL:
		return "IF_TYPE_CCTEMUL"
	case IF_TYPE_FASTETHER:
		return "IF_TYPE_FASTETHER"
	case IF_TYPE_ISDN:
		return "IF_TYPE_ISDN"
	case IF_TYPE_V11:
		return "IF_TYPE_V11"
	case IF_TYPE_V36:
		return "IF_TYPE_V36"
	case IF_TYPE_G703_64K:
		return "IF_TYPE_G703_64K"
	case IF_TYPE_G703_2MB:
		return "IF_TYPE_G703_2MB"
	case IF_TYPE_QLLC:
		return "IF_TYPE_QLLC"
	case IF_TYPE_FASTETHER_FX:
		return "IF_TYPE_FASTETHER_FX"
	case IF_TYPE_CHANNEL:
		return "IF_TYPE_CHANNEL"
	case IF_TYPE_IEEE80211:
		return "IF_TYPE_IEEE80211"
	case IF_TYPE_IBM370PARCHAN:
		return "IF_TYPE_IBM370PARCHAN"
	case IF_TYPE_ESCON:
		return "IF_TYPE_ESCON"
	case IF_TYPE_DLSW:
		return "IF_TYPE_DLSW"
	case IF_TYPE_ISDN_S:
		return "IF_TYPE_ISDN_S"
	case IF_TYPE_ISDN_U:
		return "IF_TYPE_ISDN_U"
	case IF_TYPE_LAP_D:
		return "IF_TYPE_LAP_D"
	case IF_TYPE_IPSWITCH:
		return "IF_TYPE_IPSWITCH"
	case IF_TYPE_RSRB:
		return "IF_TYPE_RSRB"
	case IF_TYPE_ATM_LOGICAL:
		return "IF_TYPE_ATM_LOGICAL"
	case IF_TYPE_DS0:
		return "IF_TYPE_DS0"
	case IF_TYPE_DS0_BUNDLE:
		return "IF_TYPE_DS0_BUNDLE"
	case IF_TYPE_BSC:
		return "IF_TYPE_BSC"
	case IF_TYPE_ASYNC:
		return "IF_TYPE_ASYNC"
	case IF_TYPE_CNR:
		return "IF_TYPE_CNR"
	case IF_TYPE_ISO88025R_DTR:
		return "IF_TYPE_ISO88025R_DTR"
	case IF_TYPE_EPLRS:
		return "IF_TYPE_EPLRS"
	case IF_TYPE_ARAP:
		return "IF_TYPE_ARAP"
	case IF_TYPE_PROP_CNLS:
		return "IF_TYPE_PROP_CNLS"
	case IF_TYPE_HOSTPAD:
		return "IF_TYPE_HOSTPAD"
	case IF_TYPE_TERMPAD:
		return "IF_TYPE_TERMPAD"
	case IF_TYPE_FRAMERELAY_MPI:
		return "IF_TYPE_FRAMERELAY_MPI"
	case IF_TYPE_X213:
		return "IF_TYPE_X213"
	case IF_TYPE_ADSL:
		return "IF_TYPE_ADSL"
	case IF_TYPE_RADSL:
		return "IF_TYPE_RADSL"
	case IF_TYPE_SDSL:
		return "IF_TYPE_SDSL"
	case IF_TYPE_VDSL:
		return "IF_TYPE_VDSL"
	case IF_TYPE_ISO88025_CRFPRINT:
		return "IF_TYPE_ISO88025_CRFPRINT"
	case IF_TYPE_MYRINET:
		return "IF_TYPE_MYRINET"
	case IF_TYPE_VOICE_EM:
		return "IF_TYPE_VOICE_EM"
	case IF_TYPE_VOICE_FXO:
		return "IF_TYPE_VOICE_FXO"
	case IF_TYPE_VOICE_FXS:
		return "IF_TYPE_VOICE_FXS"
	case IF_TYPE_VOICE_ENCAP:
		return "IF_TYPE_VOICE_ENCAP"
	case IF_TYPE_VOICE_OVERIP:
		return "IF_TYPE_VOICE_OVERIP"
	case IF_TYPE_ATM_DXI:
		return "IF_TYPE_ATM_DXI"
	case IF_TYPE_ATM_FUNI:
		return "IF_TYPE_ATM_FUNI"
	case IF_TYPE_ATM_IMA:
		return "IF_TYPE_ATM_IMA"
	case IF_TYPE_PPPMULTILINKBUNDLE:
		return "IF_TYPE_PPPMULTILINKBUNDLE"
	case IF_TYPE_IPOVER_CDLC:
		return "IF_TYPE_IPOVER_CDLC"
	case IF_TYPE_IPOVER_CLAW:
		return "IF_TYPE_IPOVER_CLAW"
	case IF_TYPE_STACKTOSTACK:
		return "IF_TYPE_STACKTOSTACK"
	case IF_TYPE_VIRTUALIPADDRESS:
		return "IF_TYPE_VIRTUALIPADDRESS"
	case IF_TYPE_MPC:
		return "IF_TYPE_MPC"
	case IF_TYPE_IPOVER_ATM:
		return "IF_TYPE_IPOVER_ATM"
	case IF_TYPE_ISO88025_FIBER:
		return "IF_TYPE_ISO88025_FIBER"
	case IF_TYPE_TDLC:
		return "IF_TYPE_TDLC"
	case IF_TYPE_GIGABITETHERNET:
		return "IF_TYPE_GIGABITETHERNET"
	case IF_TYPE_HDLC:
		return "IF_TYPE_HDLC"
	case IF_TYPE_LAP_F:
		return "IF_TYPE_LAP_F"
	case IF_TYPE_V37:
		return "IF_TYPE_V37"
	case IF_TYPE_X25_MLP:
		return "IF_TYPE_X25_MLP"
	case IF_TYPE_X25_HUNTGROUP:
		return "IF_TYPE_X25_HUNTGROUP"
	case IF_TYPE_TRANSPHDLC:
		return "IF_TYPE_TRANSPHDLC"
	case IF_TYPE_INTERLEAVE:
		return "IF_TYPE_INTERLEAVE"
	case IF_TYPE_FAST:
		return "IF_TYPE_FAST"
	case IF_TYPE_IP:
		return "IF_TYPE_IP"
	case IF_TYPE_DOCSCABLE_MACLAYER:
		return "IF_TYPE_DOCSCABLE_MACLAYER"
	case IF_TYPE_DOCSCABLE_DOWNSTREAM:
		return "IF_TYPE_DOCSCABLE_DOWNSTREAM"
	case IF_TYPE_DOCSCABLE_UPSTREAM:
		return "IF_TYPE_DOCSCABLE_UPSTREAM"
	case IF_TYPE_A12MPPSWITCH:
		return "IF_TYPE_A12MPPSWITCH"
	case IF_TYPE_TUNNEL:
		return "IF_TYPE_TUNNEL"
	case IF_TYPE_COFFEE:
		return "IF_TYPE_COFFEE"
	case IF_TYPE_CES:
		return "IF_TYPE_CES"
	case IF_TYPE_ATM_SUBINTERFACE:
		return "IF_TYPE_ATM_SUBINTERFACE"
	case IF_TYPE_L2_VLAN:
		return "IF_TYPE_L2_VLAN"
	case IF_TYPE_L3_IPVLAN:
		return "IF_TYPE_L3_IPVLAN"
	case IF_TYPE_L3_IPXVLAN:
		return "IF_TYPE_L3_IPXVLAN"
	case IF_TYPE_DIGITALPOWERLINE:
		return "IF_TYPE_DIGITALPOWERLINE"
	case IF_TYPE_MEDIAMAILOVERIP:
		return "IF_TYPE_MEDIAMAILOVERIP"
	case IF_TYPE_DTM:
		return "IF_TYPE_DTM"
	case IF_TYPE_DCN:
		return "IF_TYPE_DCN"
	case IF_TYPE_IPFORWARD:
		return "IF_TYPE_IPFORWARD"
	case IF_TYPE_MSDSL:
		return "IF_TYPE_MSDSL"
	case IF_TYPE_IEEE1394:
		return "IF_TYPE_IEEE1394"
	case IF_TYPE_IF_GSN:
		return "IF_TYPE_IF_GSN"
	case IF_TYPE_DVBRCC_MACLAYER:
		return "IF_TYPE_DVBRCC_MACLAYER"
	case IF_TYPE_DVBRCC_DOWNSTREAM:
		return "IF_TYPE_DVBRCC_DOWNSTREAM"
	case IF_TYPE_DVBRCC_UPSTREAM:
		return "IF_TYPE_DVBRCC_UPSTREAM"
	case IF_TYPE_ATM_VIRTUAL:
		return "IF_TYPE_ATM_VIRTUAL"
	case IF_TYPE_MPLS_TUNNEL:
		return "IF_TYPE_MPLS_TUNNEL"
	case IF_TYPE_SRP:
		return "IF_TYPE_SRP"
	case IF_TYPE_VOICEOVERATM:
		return "IF_TYPE_VOICEOVERATM"
	case IF_TYPE_VOICEOVERFRAMERELAY:
		return "IF_TYPE_VOICEOVERFRAMERELAY"
	case IF_TYPE_IDSL:
		return "IF_TYPE_IDSL"
	case IF_TYPE_COMPOSITELINK:
		return "IF_TYPE_COMPOSITELINK"
	case IF_TYPE_SS7_SIGLINK:
		return "IF_TYPE_SS7_SIGLINK"
	case IF_TYPE_PROP_WIRELESS_P2P:
		return "IF_TYPE_PROP_WIRELESS_P2P"
	case IF_TYPE_FR_FORWARD:
		return "IF_TYPE_FR_FORWARD"
	case IF_TYPE_RFC1483:
		return "IF_TYPE_RFC1483"
	case IF_TYPE_USB:
		return "IF_TYPE_USB"
	case IF_TYPE_IEEE8023AD_LAG:
		return "IF_TYPE_IEEE8023AD_LAG"
	case IF_TYPE_BGP_POLICY_ACCOUNTING:
		return "IF_TYPE_BGP_POLICY_ACCOUNTING"
	case IF_TYPE_FRF16_MFR_BUNDLE:
		return "IF_TYPE_FRF16_MFR_BUNDLE"
	case IF_TYPE_H323_GATEKEEPER:
		return "IF_TYPE_H323_GATEKEEPER"
	case IF_TYPE_H323_PROXY:
		return "IF_TYPE_H323_PROXY"
	case IF_TYPE_MPLS:
		return "IF_TYPE_MPLS"
	case IF_TYPE_MF_SIGLINK:
		return "IF_TYPE_MF_SIGLINK"
	case IF_TYPE_HDSL2:
		return "IF_TYPE_HDSL2"
	case IF_TYPE_SHDSL:
		return "IF_TYPE_SHDSL"
	case IF_TYPE_DS1_FDL:
		return "IF_TYPE_DS1_FDL"
	case IF_TYPE_POS:
		return "IF_TYPE_POS"
	case IF_TYPE_DVB_ASI_IN:
		return "IF_TYPE_DVB_ASI_IN"
	case IF_TYPE_DVB_ASI_OUT:
		return "IF_TYPE_DVB_ASI_OUT"
	case IF_TYPE_PLC:
		return "IF_TYPE_PLC"
	case IF_TYPE_NFAS:
		return "IF_TYPE_NFAS"
	case IF_TYPE_TR008:
		return "IF_TYPE_TR008"
	case IF_TYPE_GR303_RDT:
		return "IF_TYPE_GR303_RDT"
	case IF_TYPE_GR303_IDT:
		return "IF_TYPE_GR303_IDT"
	case IF_TYPE_ISUP:
		return "IF_TYPE_ISUP"
	case IF_TYPE_PROP_DOCS_WIRELESS_MACLAYER:
		return "IF_TYPE_PROP_DOCS_WIRELESS_MACLAYER"
	case IF_TYPE_PROP_DOCS_WIRELESS_DOWNSTREAM:
		return "IF_TYPE_PROP_DOCS_WIRELESS_DOWNSTREAM"
	case IF_TYPE_PROP_DOCS_WIRELESS_UPSTREAM:
		return "IF_TYPE_PROP_DOCS_WIRELESS_UPSTREAM"
	case IF_TYPE_HIPERLAN2:
		return "IF_TYPE_HIPERLAN2"
	case IF_TYPE_PROP_BWA_P2MP:
		return "IF_TYPE_PROP_BWA_P2MP"
	case IF_TYPE_SONET_OVERHEAD_CHANNEL:
		return "IF_TYPE_SONET_OVERHEAD_CHANNEL"
	case IF_TYPE_DIGITAL_WRAPPER_OVERHEAD_CHANNEL:
		return "IF_TYPE_DIGITAL_WRAPPER_OVERHEAD_CHANNEL"
	case IF_TYPE_AAL2:
		return "IF_TYPE_AAL2"
	case IF_TYPE_RADIO_MAC:
		return "IF_TYPE_RADIO_MAC"
	case IF_TYPE_ATM_RADIO:
		return "IF_TYPE_ATM_RADIO"
	case IF_TYPE_IMT:
		return "IF_TYPE_IMT"
	case IF_TYPE_MVL:
		return "IF_TYPE_MVL"
	case IF_TYPE_REACH_DSL:
		return "IF_TYPE_REACH_DSL"
	case IF_TYPE_FR_DLCI_ENDPT:
		return "IF_TYPE_FR_DLCI_ENDPT"
	case IF_TYPE_ATM_VCI_ENDPT:
		return "IF_TYPE_ATM_VCI_ENDPT"
	case IF_TYPE_OPTICAL_CHANNEL:
		return "IF_TYPE_OPTICAL_CHANNEL"
	case IF_TYPE_OPTICAL_TRANSPORT:
		return "IF_TYPE_OPTICAL_TRANSPORT"
	case IF_TYPE_IEEE80216_WMAN:
		return "IF_TYPE_IEEE80216_WMAN"
	case IF_TYPE_WWANPP:
		return "IF_TYPE_WWANPP"
	case IF_TYPE_WWANPP2:
		return "IF_TYPE_WWANPP2"
	case IF_TYPE_IEEE802154:
		return "IF_TYPE_IEEE802154"
	case IF_TYPE_XBOX_WIRELESS:
		return "IF_TYPE_XBOX_WIRELESS"
	default:
		return fmt.Sprintf("IFTYPE_UNKNOWN(%d)", t)
	}
}
