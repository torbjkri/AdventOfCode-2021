package dec16

type Packet interface {
	getVersionSum() int
	getTypeVersion() int
	getTypeID() int
	getValue() int
}

type LiteralPacket struct {
	type_id_ int
	version_ int
	value_   int
}

func (p LiteralPacket) getVersionSum() int {
	return p.version_
}

func (p LiteralPacket) getTypeVersion() int {
	return p.version_
}

func (p LiteralPacket) getTypeID() int {
	return p.type_id_
}

func (p LiteralPacket) getValue() int {
	return p.value_
}

type OperatorPacket struct {
	packets_ []Packet
	type_id_ int
	version_ int
}

func (p OperatorPacket) getVersionSum() int {
	sum := p.getTypeVersion()
	for _, elem := range p.packets_ {
		sum += elem.getVersionSum()
	}

	return sum
}

func (p OperatorPacket) getTypeVersion() int {
	return p.version_
}

func (p OperatorPacket) getTypeID() int {
	return p.version_
}

func (p OperatorPacket) getValue() int {
	values := []int{}

	for _, elem := range p.packets_ {
		values = append(values, elem.getValue())
	}

	switch p.type_id_ {
	case 0:
		sum := 0
		for _, val := range values {
			sum += val
		}
		return sum
	case 1:
		sum := 0
		for i, val := range values {
			if i == 0 {
				sum = val
			} else {
				sum *= val
			}
		}
		return sum
	case 2:
		min := 0
		for i, val := range values {
			if i == 0 {
				min = val
			} else {
				if val < min {
					min = val
				}
			}
		}
		return min
	case 3:
		max := 0
		for i, val := range values {
			if i == 0 {
				max = val
			} else {
				if val > max {
					max = val
				}
			}
		}
		return max
	case 5:
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	default:
		return -100

	}
}

// return rest data, packet created, bits used
func parse_packet(data []int) ([]int, Packet, int) {
	version := GetPacketVersion(data)
	typeID := GetPacketTypeID(data)

	bits_used := 6

	if typeID == 4 {
		p := LiteralPacket{}
		p.version_ = version
		p.type_id_ = typeID

		num := make([]int, 0)
		for {
			last_num := data[bits_used]
			num = append(num, data[bits_used+1:bits_used+5]...)
			bits_used += 5
			if last_num == 0 {
				break
			}
		}
		p.value_ = BitSliceToInt(num)
		return data[bits_used:], p, bits_used

	} else {
		p := OperatorPacket{}

		p.packets_ = []Packet{}

		p.version_ = version
		p.type_id_ = typeID

		operator_type := data[bits_used]
		bits_used += 1

		if operator_type == 0 {
			total_bits_in_packets := BitSliceToInt(data[bits_used : bits_used+15])
			packets_use := 0
			bits_used += 15

			for packets_use < total_bits_in_packets {
				_, np, this_packet_use := parse_packet(data[bits_used:])
				bits_used += this_packet_use
				packets_use += this_packet_use
				p.packets_ = append(p.packets_, np)
			}
			return data[bits_used:], p, bits_used

		} else {
			total_packets := BitSliceToInt(data[bits_used : bits_used+11])
			bits_used += 11

			for i := 0; i < total_packets; i++ {
				_, np, this_packet_use := parse_packet(data[bits_used:])
				bits_used += this_packet_use
				p.packets_ = append(p.packets_, np)
			}
			return data[bits_used:], p, bits_used
		}

		return []int{}, p, 0
	}

}

func GetPacketVersion(d []int) int {
	return BitSliceToInt(d[:3])
}

func GetPacketTypeID(d []int) int {
	return BitSliceToInt(d[3:6])
}
