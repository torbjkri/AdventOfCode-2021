package dec16

//
//func (p OperatorPacket) detangleOperatorPacket() []Packet {
//	result := []Packet{}
//
//	if p.data_[6] == 0 {
//		// 15 bith saying total size of rest of packets
//	} else {
//		// 11 bits given number of sub packets
//		num_packets := BitSliceToInt(p.data_[7 : 7+11])
//
//		for i := 0; i < num_packets; i++ {
//			np := Packet{}
//			np.data_ = p.data_[18+i*11 : 18+(i+1)*11]
//			np.
//		}
//
//	}
//}
//
//func extractLiteralData(d []int) ([]int, Packet) {
//	out_data = []int{}
//	p := LiteralPacket{}
//
//	length := 0
//
//	last_element := false
//
//	for !last_element {
//		if d[6+length*5] == 0 {
//			last_element = true
//		}
//		length += 1
//	}
//	p.data_ = d[:6+length*5]
//
//	out_data = d[6+length*5:]
//
//	return out_data, p
//}

//
//func DetangleOperatorPacket(bits []int) []Packet {
//
//	result := [][]int{}
//	// Ta inn full pakke
//	if bits[6] == 0 {
//		// 15 bit size saying total size of rest of packet
//		packetSizes := BitSliceToInt(bits[7 : 7+15])
//		sub_packet_bits := bits[22 : 22+packetSizes]
//
//		lastPacket := false
//		for !lastPacket {
//			version := GetPacketTypeID(sub_packet_bits)
//			if
//		}
//		// Må detanle disse pakkene også
//
//	} else {
//		// 11 bits saying the number of sub packets
//		// each being 11 bits?\
//		num_packets := BitSliceToInt(bits[7 : 7+11])
//
//		for i := 0; i < num_packets; i++ {
//			result = append(result, bits[18+i*11:18+(i+1)*11])
//		}
//	}
//}
//
