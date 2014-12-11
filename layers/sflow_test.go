package layers

import (
	"code.google.com/p/gopacket"
	_ "fmt"
	"reflect"
	"testing"
)

var SFlowTestPacket = []byte{
	0x84, 0x2b, 0x2b, 0x16, 0x8b, 0x62, 0xf0, 0x50, 0x56, 0x85, 0x3a, 0xfd, 0x08, 0x00, 0x45, 0x00,
	0x05, 0xbc, 0x9c, 0x04, 0x40, 0x00, 0xff, 0x11, 0xc7, 0x00, 0x0a, 0x01, 0xff, 0x0e, 0x0a, 0x01,
	0x00, 0x1b, 0xc7, 0x57, 0x18, 0xc7, 0x05, 0xa8, 0x22, 0x3b, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x01, 0x0a, 0x01, 0xf8, 0x16, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x9d, 0xfb, 0x40, 0x49,
	0xc6, 0xcd, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x00, 0x26,
	0x27, 0xe8, 0x00, 0x00, 0x02, 0x13, 0x00, 0x00, 0x3e, 0x80, 0x50, 0xbd, 0xe5, 0x80, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xd2, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6d,
	0xf0, 0x40, 0x08, 0x00, 0x45, 0x00, 0x05, 0xc0, 0x6b, 0xaa, 0x40, 0x00, 0x40, 0x06, 0x8f, 0x41,
	0x0a, 0x01, 0x0e, 0x16, 0x36, 0xf0, 0xeb, 0x45, 0x76, 0xfd, 0x00, 0x50, 0xca, 0x77, 0xef, 0x96,
	0xfc, 0x28, 0x63, 0x40, 0x50, 0x10, 0x00, 0x3c, 0x64, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
	0xf4, 0x00, 0x00, 0x02, 0x77, 0x00, 0x00, 0x00, 0xfd, 0x3b, 0x8c, 0xe7, 0x04, 0x4a, 0x2d, 0xb2,
	0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x01, 0x48, 0xcc, 0x11, 0x0d, 0xe3, 0x00,
	0x26, 0x85, 0x30, 0x00, 0x00, 0x07, 0x66, 0x00, 0x02, 0xd0, 0x8a, 0x00, 0x02, 0xce, 0xf0, 0x00,
	0x29, 0x7e, 0x80, 0x00, 0x02, 0xd0, 0x98, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x26, 0x85, 0x30, 0x00,
	0x00, 0x00, 0xf4, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5e, 0x5c, 0x1e, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xa2, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc,
	0x9f, 0xfd, 0x40, 0x00, 0x40, 0x06, 0x6a, 0xfa, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x08, 0x13,
	0x23, 0x84, 0xb7, 0x22, 0x8a, 0xc9, 0x50, 0xb5, 0x4e, 0x10, 0x2a, 0x87, 0x80, 0x10, 0x06, 0x01,
	0x10, 0xa6, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x07, 0x99, 0x3a, 0xd8, 0x5b,
	0x01, 0x46, 0x09, 0x00, 0x0c, 0x00, 0x0c, 0x3c, 0xac, 0x4a, 0x1b, 0x06, 0x04, 0x78, 0x78, 0x4e,
	0xc2, 0x05, 0x46, 0x43, 0x06, 0x04, 0x78, 0x78, 0xee, 0x9c, 0x00, 0x41, 0xef, 0x05, 0x81, 0x32,
	0x1b, 0x06, 0x04, 0x78, 0x78, 0x56, 0x72, 0x05, 0x4e, 0x92, 0x00, 0x96, 0x39, 0x00, 0xea, 0x3f,
	0x01, 0x15, 0xa3, 0x08, 0x04, 0x42, 0x6a, 0x82, 0x87, 0x08, 0x05, 0xcc, 0x00, 0x04, 0x00, 0x00,
	0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5a,
	0xcd, 0xd0, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x07, 0xd0, 0x95, 0x67, 0xe1, 0x30, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x46, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f,
	0x11, 0x28, 0x08, 0x00, 0x45, 0x00, 0x04, 0x34, 0xdb, 0x36, 0x40, 0x00, 0x40, 0x06, 0x38, 0xac,
	0x0a, 0x01, 0x0e, 0x11, 0x0a, 0x01, 0x00, 0xcf, 0x23, 0x84, 0xa0, 0x3f, 0x3c, 0xce, 0xd5, 0x4a,
	0x72, 0x0b, 0x5d, 0x1a, 0x80, 0x10, 0x06, 0x01, 0x8a, 0x50, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a,
	0xef, 0x1f, 0xa2, 0xba, 0xe6, 0xfa, 0xae, 0xb3, 0xfe, 0xcf, 0x00, 0x19, 0xcf, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x01, 0xb9, 0x79, 0xdd, 0x42, 0x00, 0x00, 0x02, 0x84, 0x9b, 0xa9, 0x02, 0xe2, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x06, 0x32, 0x39, 0x35, 0x34, 0x33, 0x36, 0x00, 0x00, 0x02, 0x70, 0xcd,
	0x16, 0x40, 0xa6, 0x98, 0x88, 0x24, 0x06, 0x50, 0xb0, 0xf4, 0xee, 0x03, 0xa6, 0xfa, 0x87, 0xaf,
	0xc1, 0x99, 0x52, 0x0d, 0x07, 0xa8, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0xa8, 0x00, 0x00, 0x20, 0xf2, 0x00, 0x00, 0x02, 0x0a, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x58, 0x00, 0x00, 0x02, 0x0a, 0x00, 0x00,
	0x00, 0x06, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x01, 0x29, 0x82, 0x6d, 0xb0, 0x6c, 0x0b, 0xcb, 0x0d, 0xdd, 0x96, 0x00, 0x06,
	0xa8, 0xc6, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x34, 0x02, 0x35, 0x58, 0x7c, 0x9e, 0x56, 0x64, 0x25, 0x71, 0x00, 0x70,
	0x5a, 0xc4, 0x00, 0x09, 0x08, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5e, 0x5c, 0x1f, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xaa, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc,
	0x0f, 0xba, 0x40, 0x00, 0x40, 0x06, 0xf4, 0x3f, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x0f, 0x11,
	0x23, 0x84, 0xcd, 0xc0, 0xf4, 0x0e, 0x90, 0x23, 0xd7, 0x32, 0x8b, 0x31, 0x80, 0x10, 0x00, 0x1d,
	0x6b, 0x12, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x28, 0xef, 0x1f, 0xec, 0x76,
	0xaa, 0x25, 0x01, 0x04, 0xc0, 0xac, 0xfe, 0x25, 0x01, 0x8e, 0x25, 0x01, 0x16, 0xc7, 0x28, 0xfe,
	0x7e, 0x70, 0xfe, 0x7e, 0x70, 0x52, 0x7e, 0x70, 0x15, 0x9b, 0xfe, 0x35, 0x01, 0xfe, 0x35, 0x01,
	0x42, 0x35, 0x01, 0xfe, 0x95, 0x77, 0xfe, 0x95, 0x77, 0xfe, 0x95, 0x77, 0x52, 0x95, 0x77, 0x00,
	0xd2, 0xfe, 0x70, 0x02, 0x92, 0x70, 0x02, 0x16, 0x60, 0x22, 0x00, 0x7e, 0xb2, 0x15, 0x00, 0x00,
	0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5a,
	0xcd, 0xd1, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x07, 0xd0, 0x95, 0x67, 0xe9, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0xb8, 0xca, 0x3a, 0x6f,
	0x11, 0x28, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc, 0xfe, 0x05, 0x40, 0x00, 0x40, 0x06, 0x06, 0xf4,
	0x0a, 0x01, 0x0e, 0x11, 0x0a, 0x01, 0x0e, 0x10, 0x23, 0x84, 0xfa, 0x29, 0xae, 0xd4, 0x95, 0x03,
	0x99, 0xb8, 0x77, 0xd0, 0x80, 0x10, 0x00, 0x1d, 0x6f, 0x4f, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a,
	0xef, 0x1f, 0xa2, 0xcc, 0xef, 0x1f, 0xf4, 0x2c, 0xfe, 0xdb, 0x05, 0xa1, 0xdb, 0x04, 0x9e, 0xc0,
	0xfe, 0x30, 0x08, 0xb2, 0x30, 0x08, 0xda, 0x2b, 0xbd, 0xfe, 0x2a, 0x01, 0xfe, 0x2a, 0x01, 0x21,
	0x2a, 0x00, 0xb2, 0xfe, 0x57, 0xb0, 0xb6, 0x57, 0xb0, 0x14, 0x74, 0xf4, 0xf0, 0x4c, 0x05, 0x68,
	0xfe, 0x54, 0x02, 0xfe, 0x54, 0x02, 0xd2, 0x54, 0x02, 0x00, 0xbe, 0xfe, 0x32, 0x0f, 0xb6, 0x32,
	0x0f, 0x14, 0x2e, 0x16, 0xaf, 0x47, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x94, 0x01, 0x5e, 0x5c, 0x20, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xb2, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x54, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x42, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x00, 0x34,
	0xa8, 0x23, 0x40, 0x00, 0x40, 0x06, 0x61, 0x7f, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x0f, 0x10,
	0x97, 0x91, 0x23, 0x84, 0x24, 0xfa, 0x91, 0xf7, 0xb4, 0xe8, 0xf3, 0x2d, 0x80, 0x10, 0x00, 0xab,
	0x7b, 0x7d, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x36, 0xef, 0x1f, 0xdc, 0xde,
	0x00, 0x00, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var rawPacket1 = []byte{
	0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6d, 0xf0, 0x40, 0x8, 0x0, 0x45, 0x0, 0x5,
	0xc0, 0x6b, 0xaa, 0x40, 0x0, 0x40, 0x6, 0x8f, 0x41, 0xa, 0x1, 0xe, 0x16, 0x36, 0xf0, 0xeb, 0x45,
	0x76, 0xfd, 0x0, 0x50, 0xca, 0x77, 0xef, 0x96, 0xfc, 0x28, 0x63, 0x40, 0x50, 0x10, 0x0, 0x3c, 0x64,
	0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x0, 0xf4, 0x0, 0x0, 0x2, 0x77, 0x0, 0x0, 0x0, 0xfd, 0x3b, 0x8c, 0xe7,
	0x4, 0x4a, 0x2d, 0xb2, 0xc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x1, 0x48, 0xcc, 0x11, 0xd, 0xe3,
	0x0, 0x26, 0x85, 0x30, 0x0, 0x0, 0x7, 0x66, 0x0, 0x2, 0xd0, 0x8a, 0x0, 0x2, 0xce, 0xf0, 0x0, 0x29,
	0x7e, 0x80, 0x0, 0x2, 0xd0, 0x98, 0x0, 0x0, 0x0, 0xc, 0x0, 0x26, 0x85, 0x30, 0x0, 0x0, 0x0, 0xf4,
	0x0, 0x0, 0x2,
}

var rawPacket2 = []byte{
	0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x8, 0x0, 0x45, 0x0, 0x5, 0xdc,
	0x9f, 0xfd, 0x40, 0x0, 0x40, 0x6, 0x6a, 0xfa, 0xa, 0x1, 0xe, 0x10, 0xa, 0x1, 0x8, 0x13, 0x23, 0x84, 0xb7,
	0x22, 0x8a, 0xc9, 0x50, 0xb5, 0x4e, 0x10, 0x2a, 0x87, 0x80, 0x10, 0x6, 0x1, 0x10, 0xa6, 0x0, 0x0, 0x1,
	0x1, 0x8, 0xa, 0xef, 0x1f, 0xf4, 0x7, 0x99, 0x3a, 0xd8, 0x5b, 0x1, 0x46, 0x9, 0x0, 0xc, 0x0, 0xc, 0x3c,
	0xac, 0x4a, 0x1b, 0x6, 0x4, 0x78, 0x78, 0x4e, 0xc2, 0x5, 0x46, 0x43, 0x6, 0x4, 0x78, 0x78, 0xee, 0x9c,
	0x0, 0x41, 0xef, 0x5, 0x81, 0x32, 0x1b, 0x6, 0x4, 0x78, 0x78, 0x56, 0x72, 0x5, 0x4e, 0x92, 0x0, 0x96,
	0x39, 0x0, 0xea, 0x3f, 0x1, 0x15, 0xa3, 0x8, 0x4, 0x42, 0x6a, 0x82, 0x87, 0x8, 0x5, 0xcc, 0x0, 0x4,
}

var rawPacket3 = []byte{
	0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0x11, 0x28, 0x8, 0x0, 0x45, 0x0, 0x4, 0x34, 0xdb,
	0x36, 0x40, 0x0, 0x40, 0x6, 0x38, 0xac, 0xa, 0x1, 0xe, 0x11, 0xa, 0x1, 0x0, 0xcf, 0x23, 0x84, 0xa0, 0x3f, 0x3c,
	0xce, 0xd5, 0x4a, 0x72, 0xb, 0x5d, 0x1a, 0x80, 0x10, 0x6, 0x1, 0x8a, 0x50, 0x0, 0x0, 0x1, 0x1, 0x8, 0xa, 0xef,
	0x1f, 0xa2, 0xba, 0xe6, 0xfa, 0xae, 0xb3, 0xfe, 0xcf, 0x0, 0x19, 0xcf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xb9, 0x79,
	0xdd, 0x42, 0x0, 0x0, 0x2, 0x84, 0x9b, 0xa9, 0x2, 0xe2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x32, 0x39, 0x35, 0x34,
	0x33, 0x36, 0x0, 0x0, 0x2, 0x70, 0xcd, 0x16, 0x40, 0xa6, 0x98, 0x88, 0x24, 0x6, 0x50, 0xb0, 0xf4, 0xee, 0x3,
	0xa6, 0xfa, 0x87, 0xaf, 0xc1, 0x99, 0x52, 0xd, 0x7, 0xa8,
}

var rawPacket4 = []byte{
	0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x8, 0x0, 0x45, 0x0, 0x5, 0xdc, 0xf, 0xba,
	0x40, 0x0, 0x40, 0x6, 0xf4, 0x3f, 0xa, 0x1, 0xe, 0x10, 0xa, 0x1, 0xf, 0x11, 0x23, 0x84, 0xcd, 0xc0, 0xf4, 0xe, 0x90,
	0x23, 0xd7, 0x32, 0x8b, 0x31, 0x80, 0x10, 0x0, 0x1d, 0x6b, 0x12, 0x0, 0x0, 0x1, 0x1, 0x8, 0xa, 0xef, 0x1f, 0xf4, 0x28,
	0xef, 0x1f, 0xec, 0x76, 0xaa, 0x25, 0x1, 0x4, 0xc0, 0xac, 0xfe, 0x25, 0x1, 0x8e, 0x25, 0x1, 0x16, 0xc7, 0x28, 0xfe,
	0x7e, 0x70, 0xfe, 0x7e, 0x70, 0x52, 0x7e, 0x70, 0x15, 0x9b, 0xfe, 0x35, 0x1, 0xfe, 0x35, 0x1, 0x42, 0x35, 0x1, 0xfe,
	0x95, 0x77, 0xfe, 0x95, 0x77, 0xfe, 0x95, 0x77, 0x52, 0x95, 0x77, 0x0, 0xd2, 0xfe, 0x70, 0x2, 0x92, 0x70, 0x2, 0x16,
	0x60, 0x22, 0x0, 0x7e, 0xb2, 0x15,
}

var rawPacket5 = []byte{
	0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0xb8, 0xca, 0x3a, 0x6f, 0x11, 0x28, 0x8, 0x0, 0x45, 0x0, 0x5, 0xdc, 0xfe, 0x5, 0x40,
	0x0, 0x40, 0x6, 0x6, 0xf4, 0xa, 0x1, 0xe, 0x11, 0xa, 0x1, 0xe, 0x10, 0x23, 0x84, 0xfa, 0x29, 0xae, 0xd4, 0x95, 0x3, 0x99,
	0xb8, 0x77, 0xd0, 0x80, 0x10, 0x0, 0x1d, 0x6f, 0x4f, 0x0, 0x0, 0x1, 0x1, 0x8, 0xa, 0xef, 0x1f, 0xa2, 0xcc, 0xef, 0x1f,
	0xf4, 0x2c, 0xfe, 0xdb, 0x5, 0xa1, 0xdb, 0x4, 0x9e, 0xc0, 0xfe, 0x30, 0x8, 0xb2, 0x30, 0x8, 0xda, 0x2b, 0xbd, 0xfe,
	0x2a, 0x1, 0xfe, 0x2a, 0x1, 0x21, 0x2a, 0x0, 0xb2, 0xfe, 0x57, 0xb0, 0xb6, 0x57, 0xb0, 0x14, 0x74, 0xf4, 0xf0, 0x4c,
	0x5, 0x68, 0xfe, 0x54, 0x2, 0xfe, 0x54, 0x2, 0xd2, 0x54, 0x2, 0x0, 0xbe, 0xfe, 0x32, 0xf, 0xb6, 0x32, 0xf, 0x14, 0x2e,
	0x16, 0xaf, 0x47,
}

var rawPacket6 = []byte{
	0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x8, 0x0, 0x45, 0x0, 0x0, 0x34, 0xa8, 0x23, 0x40,
	0x0, 0x40, 0x6, 0x61, 0x7f, 0xa, 0x1, 0xe, 0x10, 0xa, 0x1, 0xf, 0x10, 0x97, 0x91, 0x23, 0x84, 0x24, 0xfa, 0x91, 0xf7, 0xb4,
	0xe8, 0xf3, 0x2d, 0x80, 0x10, 0x0, 0xab, 0x7b, 0x7d, 0x0, 0x0, 0x1, 0x1, 0x8, 0xa, 0xef, 0x1f, 0xf4, 0x36, 0xef, 0x1f, 0xdc,
	0xde, 0x0, 0x0,
}

func BenchmarkDecodeSFlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sflow SFlowDatagram
		parser := gopacket.NewDecodingLayerParser(LayerTypeSFlow, &sflow)
		decoded := []gopacket.LayerType{}
		parser.DecodeLayers(SFlowTestPacket, &decoded)
	}
}

func TestDecodeUDPSFlow(t *testing.T) {
	p := gopacket.NewPacket(SFlowTestPacket, LayerTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv4, LayerTypeUDP, LayerTypeSFlow}, t)
	if got, ok := p.TransportLayer().(*UDP); ok {
		want := &UDP{
			BaseLayer: BaseLayer{SFlowTestPacket[34:42], SFlowTestPacket[42:]},
			sPort:     []byte{199, 87},
			dPort:     []byte{24, 199},
			SrcPort:   51031,
			DstPort:   6343,
			Checksum:  8763,
			Length:    1448,
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("UDP layer mismatch, \nwant  %#v\ngot %#v\n", want, got)
		}
	} else {
		t.Error("Transport layer packet not UDP")
	}

}

func TestDecodeSFlowDatagram(t *testing.T) {
	p := gopacket.NewPacket(SFlowTestPacket, LayerTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv4, LayerTypeUDP, LayerTypeSFlow}, t)
	if got, ok := p.ApplicationLayer().(*SFlowDatagram); ok {
		want := &SFlowDatagram{
			DatagramVersion: uint32(5),
			AgentAddress:    []byte{0xa, 0x1, 0xf8, 0x16},
			SubAgentID:      uint32(17),
			SequenceNumber:  uint32(40443),
			AgentUptime:     uint32(1078576845),
			SampleCount:     uint32(7),
			FlowSamples: []SFlowFlowSample{
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x2627e8,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x213,
					SamplingRate:    0x3e80,
					SamplePool:      0x50bde580,
					Dropped:         0x0,
					InputInterface:  0x213,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5d2,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(rawPacket1, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15e5c1e,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12fa290,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(rawPacket2, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15acdd0,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x255,
					SamplingRate:    0x7d0,
					SamplePool:      0x9567e130,
					Dropped:         0x0,
					InputInterface:  0x255,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x446,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(rawPacket3, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15e5c1f,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12faa60,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(rawPacket4, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15acdd1,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x255,
					SamplingRate:    0x7d0,
					SamplePool:      0x9567e900,
					Dropped:         0x0,
					InputInterface:  0x255,
					OutputInterface: 0x257,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(rawPacket5, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x202,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0x94,
					SequenceNumber:  0x15e5c20,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12fb230,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x54,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x46,
							PayloadRemoved: 0x4,
							HeaderLength:   0x42,
							Header:         gopacket.NewPacket(rawPacket6, LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
			},
			CounterSamples: []SFlowCounterSample{
				SFlowCounterSample{
					Format:         0x2,
					SampleLength:   0xa8,
					SequenceNumber: 0x20f2,
					SourceIDClass:  0x0,
					SourceIDIndex:  0x20a,
					RecordCount:    0x2,
					Records: []SFlowRecord{
						SFlowGenericInterfaceCounters{
							SFlowBaseCounterRecord: SFlowBaseCounterRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x58,
							},
							IfIndex:            0x20a,
							IfType:             0x6,
							IfSpeed:            0x2540be400,
							IfDirection:        0x1,
							IfStatus:           0x3,
							IfInOctets:         0x129826db06c0b,
							IfInUcastPkts:      0xcb0ddd96,
							IfInMulticastPkts:  0x6a8c6,
							IfInBroadcastPkts:  0x7b,
							IfInDiscards:       0x0,
							IfInErrors:         0x0,
							IfInUnknownProtos:  0x0,
							IfOutOctets:        0x340235587c9e,
							IfOutUcastPkts:     0x56642571,
							IfOutMulticastPkts: 0x705ac4,
							IfOutBroadcastPkts: 0x908f1,
							IfOutDiscards:      0x0,
							IfOutErrors:        0x0,
							IfPromiscuousMode:  0x0,
						},
						SFlowEthernetCounters{
							SFlowBaseCounterRecord: SFlowBaseCounterRecord{
								EnterpriseID:   0x0,
								Format:         0x2,
								FlowDataLength: 0x34,
							},
							AlignmentErrors:           0x0,
							FCSErrors:                 0x0,
							SingleCollisionFrames:     0x0,
							MultipleCollisionFrames:   0x0,
							SQETestErrors:             0x0,
							DeferredTransmissions:     0x0,
							LateCollisions:            0x0,
							ExcessiveCollisions:       0x0,
							InternalMacTransmitErrors: 0x0,
							CarrierSenseErrors:        0x0,
							FrameTooLongs:             0x0,
							InternalMacReceiveErrors:  0x0,
							SymbolErrors:              0x0,
						},
					},
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("SFlow layer mismatch, \nwant:\n\n%#v\ngot:\n\n\n%#v\n\n", want, got)
		}
	} else {
		t.Error("Application layer packet not UDP")
	}
}
