// Copyright 2012 Google, Inc. All rights reserved.
// Copyright 2009-2012 Andreas Krennmair. All rights reserved.

package gopacket

import (
	"fmt"
	"github.com/gconnell/gopacket"
)

// EthernetType is an enumeration of ethernet type values, and acts as a decoder
// for any type it supports.
type EthernetType uint16

const (
	// EthernetTypeLLC is not an actual ethernet type.  It is instead a
	// placeholder we use in Ethernet frames that use the 802.3 standard of
	// srcmac|dstmac|length|LLC instead of srcmac|dstmac|ethertype.
	EthernetTypeLLC            EthernetType = 0
	EthernetTypeCDP            EthernetType = 0x2000
	EthernetTypeIPv4           EthernetType = 0x0800
	EthernetTypeARP            EthernetType = 0x0806
	EthernetTypeIPv6           EthernetType = 0x86DD
	EthernetTypeDot1Q          EthernetType = 0x8100
	EthernetTypePPPoEDiscovery EthernetType = 0x8863
	EthernetTypePPPoESession   EthernetType = 0x8864
	EthernetTypeCTP            EthernetType = 0x9000
)

func (e EthernetType) Decode(data []byte) (out gopacket.DecodeResult, err error) {
	switch e {
	case EthernetTypeLLC:
		return LayerTypeLLC.Decode(data)
	case EthernetTypeIPv4:
		return LayerTypeIPv4.Decode(data)
	case EthernetTypeIPv6:
		return LayerTypeIPv6.Decode(data)
	case EthernetTypeARP:
		return LayerTypeARP.Decode(data)
	case EthernetTypeDot1Q:
		return LayerTypeDot1Q.Decode(data)
	case EthernetTypePPPoEDiscovery, EthernetTypePPPoESession:
		return LayerTypePPPoE.Decode(data)
	case EthernetTypeCTP:
		return LayerTypeCTP.Decode(data)
	case EthernetTypeCDP:
		return LayerTypeCDP.Decode(data)
	}
	err = fmt.Errorf("Unsupported ethernet type %v", e)
	return
}

// IPProtocol is an enumeration of IP protocol values, and acts as a decoder
// for any type it supports.
type IPProtocol uint8

const (
	IPProtocolICMP    IPProtocol = 1
	IPProtocolTCP     IPProtocol = 6
	IPProtocolUDP     IPProtocol = 17
	IPProtocolSCTP    IPProtocol = 132
	IPProtocolIPv6    IPProtocol = 41
	IPProtocolIPIP    IPProtocol = 94
	IPProtocolEtherIP IPProtocol = 97
	IPProtocolRUDP    IPProtocol = 27
	IPProtocolGRE     IPProtocol = 47
)

func (ip IPProtocol) Decode(data []byte) (out gopacket.DecodeResult, err error) {
	switch ip {
	case IPProtocolTCP:
		return LayerTypeTCP.Decode(data)
	case IPProtocolUDP:
		return LayerTypeUDP.Decode(data)
	case IPProtocolICMP:
		return LayerTypeICMP.Decode(data)
	case IPProtocolSCTP:
		return LayerTypeSCTP.Decode(data)
	case IPProtocolIPv6:
		return LayerTypeIPv6.Decode(data)
	case IPProtocolIPIP:
		return LayerTypeIPv4.Decode(data)
	case IPProtocolEtherIP:
		return LayerTypeEtherIP.Decode(data)
	case IPProtocolRUDP:
		return LayerTypeRUDP.Decode(data)
	case IPProtocolGRE:
		return LayerTypeGRE.Decode(data)
	}
	err = fmt.Errorf("Unsupported IP protocol %v", ip)
	return
}

// LinkType is an enumeration of link types, and acts as a decoder for any
// link type it supports.
type LinkType int

const (
	// According to pcap-linktype(7).
	LinkTypeNull           LinkType = 0
	LinkTypeEthernet       LinkType = 1
	LinkTypeTokenRing      LinkType = 6
	LinkTypeArcNet         LinkType = 7
	LinkTypeSLIP           LinkType = 8
	LinkTypePPP            LinkType = 9
	LinkTypeFDDI           LinkType = 10
	LinkTypeATM_RFC1483    LinkType = 100
	LinkTypeRaw            LinkType = 101
	LinkTypePPP_HDLC       LinkType = 50
	LinkTypePPPEthernet    LinkType = 51
	LinkTypeC_HDLC         LinkType = 104
	LinkTypeIEEE802_11     LinkType = 105
	LinkTypeFRelay         LinkType = 107
	LinkTypeLoop           LinkType = 108
	LinkTypeLinuxSLL       LinkType = 113
	LinkTypeLTalk          LinkType = 104
	LinkTypePFLog          LinkType = 117
	LinkTypePrismHeader    LinkType = 119
	LinkTypeIPOverFC       LinkType = 122
	LinkTypeSunATM         LinkType = 123
	LinkTypeIEEE80211Radio LinkType = 127
	LinkTypeARCNetLinux    LinkType = 129
	LinkTypeLinuxIRDA      LinkType = 144
	LinkTypeLinuxLAPD      LinkType = 177
)

func (l LinkType) Decode(data []byte) (out gopacket.DecodeResult, err error) {
	switch l {
	case LinkTypeEthernet:
		return LayerTypeEthernet.Decode(data)
	case LinkTypePPP:
		return LayerTypePPP.Decode(data)
	}
	err = fmt.Errorf("Unsupported link-layer type %v", l)
	return
}

// PPPoECode is the PPPoE code enum, taken from http://tools.ietf.org/html/rfc2516
type PPPoECode int

const (
	PPPoECodePADI    PPPoECode = 0x09
	PPPoECodePADO    PPPoECode = 0x07
	PPPoECodePADR    PPPoECode = 0x19
	PPPoECodePADS    PPPoECode = 0x65
	PPPoECodePADT    PPPoECode = 0xA7
	PPPoECodeSession PPPoECode = 0x00
)

// Decode decodes a PPPoE payload, based on the PPPoECode.
func (p PPPoECode) Decode(data []byte) (_ gopacket.DecodeResult, err error) {
	switch p {
	case PPPoECodeSession:
		return LayerTypePPP.Decode(data)
	}
	err = fmt.Errorf("Cannot currently handle PPPoE error code %v", p)
	return
}

// PPPType is an enumeration of PPP type values, and acts as a decoder for any
// type it supports.
type PPPType uint16

const (
	PPPTypeIPv4 PPPType = 0x0021
	PPPTypeIPv6 PPPType = 0x0057
)

func (p PPPType) Decode(data []byte) (out gopacket.DecodeResult, err error) {
	switch p {
	case PPPTypeIPv4:
		return LayerTypeIPv4.Decode(data)
	case PPPTypeIPv6:
		return LayerTypeIPv6.Decode(data)
	}
	err = fmt.Errorf("Unsupported PPP type %v", p)
	return
}

// SCTPChunkType is an enumeration of chunk types inside SCTP packets.
type SCTPChunkType uint8

const (
	SCTPChunkTypeData             SCTPChunkType = 0
	SCTPChunkTypeInit             SCTPChunkType = 1
	SCTPChunkTypeInitAck          SCTPChunkType = 2
	SCTPChunkTypeSack             SCTPChunkType = 3
	SCTPChunkTypeHeartbeat        SCTPChunkType = 4
	SCTPChunkTypeHeartbeatAck     SCTPChunkType = 5
	SCTPChunkTypeAbort            SCTPChunkType = 6
	SCTPChunkTypeShutdown         SCTPChunkType = 7
	SCTPChunkTypeShutdownAck      SCTPChunkType = 8
	SCTPChunkTypeError            SCTPChunkType = 9
	SCTPChunkTypeCookieEcho       SCTPChunkType = 10
	SCTPChunkTypeCookieAck        SCTPChunkType = 11
	SCTPChunkTypeShutdownComplete SCTPChunkType = 14
)

func (s SCTPChunkType) Decode(data []byte) (_ gopacket.DecodeResult, err error) {
	switch s {
	case SCTPChunkTypeData:
		return LayerTypeSCTPData.Decode(data)
	case SCTPChunkTypeInit, SCTPChunkTypeInitAck:
		return LayerTypeSCTPInit.Decode(data)
	case SCTPChunkTypeSack:
		return LayerTypeSCTPSack.Decode(data)
	case SCTPChunkTypeHeartbeat, SCTPChunkTypeHeartbeatAck:
		return LayerTypeSCTPHeartbeat.Decode(data)
	case SCTPChunkTypeAbort, SCTPChunkTypeError:
		return LayerTypeSCTPError.Decode(data)
	case SCTPChunkTypeShutdown:
		return LayerTypeSCTPShutdown.Decode(data)
	case SCTPChunkTypeShutdownAck:
		return LayerTypeSCTPShutdownAck.Decode(data)
	case SCTPChunkTypeCookieEcho:
		return LayerTypeSCTPCookieEcho.Decode(data)
	case SCTPChunkTypeCookieAck, SCTPChunkTypeShutdownComplete:
		return LayerTypeSCTPEmptyLayer.Decode(data)
	}
	err = fmt.Errorf("Unable to decode SCTP chunk type %v", s)
	return
}
