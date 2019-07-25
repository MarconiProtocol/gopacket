// Copyright 2014, Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"net"
	"reflect"
	"testing"

	"github.com/MarconiProtocol/gopacket"
)

// Generator: python layers/test_creator.py --layerType=LayerTypePrismHeader --linkType=LinkTypePrismHeader --name=Prism%s ~/tmp/dump.pcap
// http://wiki.wireshark.org/SampleCaptures#Sample_Captures

// testPacketPrism is the packet:
//   21:32:37.616872 BSSID:Broadcast DA:Broadcast SA:cc:fa:00:ad:79:e8 (oui Unknown) Probe Request () [1.0 2.0 5.5 11.0 Mbit]
//        0x0000:  4400 0000 9000 0000 7261 3000 0000 0000
//        0x0010:  0000 0000 0000 0000 4400 0100 0000 0400
//        0x0020:  f9c1 2900 4400 0200 0000 0000 0000 0000
//        0x0030:  4400 0300 0000 0400 0a00 0000 4400 0400
//        0x0040:  0000 0400 e1ff ffff 0000 0000 0000 0000
//        0x0050:  0000 0000 4400 0600 0000 0400 0000 0000
//        0x0060:  4400 0700 0000 0400 0000 0000 4400 0800
//        0x0070:  0000 0400 0200 0000 4400 0900 0000 0000
//        0x0080:  0000 0000 4400 0a00 0000 0400 7e00 0000
//        0x0090:  4000 0000 ffff ffff ffff ccfa 00ad 79e8
//        0x00a0:  ffff ffff ffff a041 0000 0104 0204 0b16
//        0x00b0:  3208 0c12 1824 3048 606c 0301 012d 1a2d
//        0x00c0:  1117 ff00 0000 0000 0000 0000 0000 0000
//        0x00d0:  0000 0000 0000 0000 007f 0800 0000 0000
//        0x00e0:  0000 40dd 0900 1018 0200 0010 0000 dd1e
//        0x00f0:  0090 4c33 2d11 17ff 0000 0000 0000 0000
//        0x0100:  0000 0000 0000 0000 0000 0000 0000

var testPacketPrism = []byte{
	0x44, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x72, 0x61, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00,
	0xf9, 0xc1, 0x29, 0x00, 0x44, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x44, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x44, 0x00, 0x04, 0x00,
	0x00, 0x00, 0x04, 0x00, 0xe1, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x06, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x44, 0x00, 0x07, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x08, 0x00,
	0x00, 0x00, 0x04, 0x00, 0x02, 0x00, 0x00, 0x00, 0x44, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x04, 0x00, 0x7e, 0x00, 0x00, 0x00,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xcc, 0xfa, 0x00, 0xad, 0x79, 0xe8,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa0, 0x41, 0x00, 0x00, 0x01, 0x04, 0x02, 0x04, 0x0b, 0x16,
	0x32, 0x08, 0x0c, 0x12, 0x18, 0x24, 0x30, 0x48, 0x60, 0x6c, 0x03, 0x01, 0x01, 0x2d, 0x1a, 0x2d,
	0x11, 0x17, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7f, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x40, 0xdd, 0x09, 0x00, 0x10, 0x18, 0x02, 0x00, 0x00, 0x10, 0x00, 0x00, 0xdd, 0x1e,
	0x00, 0x90, 0x4c, 0x33, 0x2d, 0x11, 0x17, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

func TestPacketPrism(t *testing.T) {
	p := gopacket.NewPacket(testPacketPrism, LinkTypePrismHeader, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypePrismHeader, LayerTypeDot11, LayerTypeDot11MgmtProbeReq}, t)

	if got, ok := p.Layer(LayerTypePrismHeader).(*PrismHeader); ok {
		want := &PrismHeader{
			BaseLayer: BaseLayer{
				Contents: []uint8{0x44, 0x0, 0x0, 0x0, 0x90, 0x0, 0x0, 0x0, 0x72, 0x61, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x1, 0x0, 0x0, 0x0, 0x4, 0x0, 0xf9, 0xc1, 0x29, 0x0, 0x44, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x3, 0x0, 0x0, 0x0, 0x4, 0x0, 0xa, 0x0, 0x0, 0x0, 0x44, 0x0, 0x4, 0x0, 0x0, 0x0, 0x4, 0x0, 0xe1, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x6, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x7, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x8, 0x0, 0x0, 0x0, 0x4, 0x0, 0x2, 0x0, 0x0, 0x0, 0x44, 0x0, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0xa, 0x0, 0x0, 0x0, 0x4, 0x0, 0x7e, 0x0, 0x0, 0x0},
				Payload:  []uint8{0x40, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xcc, 0xfa, 0x0, 0xad, 0x79, 0xe8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa0, 0x41, 0x0, 0x0, 0x1, 0x4, 0x2, 0x4, 0xb, 0x16, 0x32, 0x8, 0xc, 0x12, 0x18, 0x24, 0x30, 0x48, 0x60, 0x6c, 0x3, 0x1, 0x1, 0x2d, 0x1a, 0x2d, 0x11, 0x17, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0xdd, 0x9, 0x0, 0x10, 0x18, 0x2, 0x0, 0x0, 0x10, 0x0, 0x0, 0xdd, 0x1e, 0x0, 0x90, 0x4c, 0x33, 0x2d, 0x11, 0x17, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, Code: 0x44, Length: 0x90, DeviceName: "ra0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00",
			Values: []PrismValue{
				PrismValue{DID: PrismDIDType1HostTime, Status: 0x0, Length: 0x4, Data: []uint8{0xf9, 0xc1, 0x29, 0x0}},
				PrismValue{DID: PrismDIDType1MACTime, Status: 0x0, Length: 0x0, Data: []uint8{}},
				PrismValue{DID: PrismDIDType1Channel, Status: 0x0, Length: 0x4, Data: []uint8{0xa, 0x0, 0x0, 0x0}},
				PrismValue{DID: PrismDIDType1RSSI, Status: 0x0, Length: 0x4, Data: []uint8{0xe1, 0xff, 0xff, 0xff}},
				PrismValue{DID: 0x0, Status: 0x0, Length: 0x0, Data: []uint8{}},
				PrismValue{DID: PrismDIDType1Signal, Status: 0x0, Length: 0x4, Data: []uint8{0x0, 0x0, 0x0, 0x0}},
				PrismValue{DID: PrismDIDType1Noise, Status: 0x0, Length: 0x4, Data: []uint8{0x0, 0x0, 0x0, 0x0}},
				PrismValue{DID: PrismDIDType1Rate, Status: 0x0, Length: 0x4, Data: []uint8{0x2, 0x0, 0x0, 0x0}},
				PrismValue{DID: PrismDIDType1TransmittedFrameIndicator, Status: 0x0, Length: 0x0, Data: []uint8{}},
				PrismValue{DID: PrismDIDType1FrameLength, Status: 0x0, Length: 0x4, Data: []uint8{0x7e, 0x0, 0x0, 0x0}},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RadioTap packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	}

	if got, ok := p.Layer(LayerTypeDot11).(*Dot11); ok {
		want := &Dot11{
			BaseLayer: BaseLayer{
				Contents: []uint8{0x40, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xcc, 0xfa, 0x0, 0xad, 0x79, 0xe8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa0, 0x41},
				Payload:  []uint8{0x0, 0x0, 0x1, 0x4, 0x2, 0x4, 0xb, 0x16, 0x32, 0x8, 0xc, 0x12, 0x18, 0x24, 0x30, 0x48, 0x60, 0x6c, 0x3, 0x1, 0x1, 0x2d, 0x1a, 0x2d, 0x11, 0x17, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0xdd, 0x9, 0x0, 0x10, 0x18, 0x2, 0x0, 0x0, 0x10, 0x0, 0x0, 0xdd, 0x1e, 0x0, 0x90, 0x4c, 0x33, 0x2d, 0x11, 0x17, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			Type:           0x10,
			Proto:          0x0,
			Flags:          0x0,
			DurationID:     0x0,
			Address1:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Address2:       net.HardwareAddr{0xcc, 0xfa, 0x0, 0xad, 0x79, 0xe8},
			Address3:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Address4:       net.HardwareAddr(nil),
			SequenceNumber: 0x041a,
			FragmentNumber: 0x0,
			Checksum:       0x0,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Dot11 packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	}
}

func BenchmarkDecodePacketPrism(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketPrism, LinkTypePrismHeader, gopacket.NoCopy)
	}
}
