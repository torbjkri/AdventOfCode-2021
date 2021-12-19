package dec16

import "testing"

func TestParseLiteral(t *testing.T) {
	data := []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0}
	new_data, p, used := parse_packet(data)

	literal, ok := p.(LiteralPacket)

	if !ok {
		t.Errorf("Packet illegal type!")
	}

	rest_value := BitSliceToInt(new_data)

	if !(rest_value == 0) {
		t.Errorf("Incorrect parsing rest! Got %v, expected %v", rest_value, 0)
	}

	if !(literal.value_ == 2021) {
		t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 2021)
	}

	if !(literal.version_ == 6) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", literal.version_, 6)
	}

	if !(literal.type_id_ == 4) {
		t.Errorf("Incorrect parsing type ID! Got %v, expected %v", literal.type_id_, 6)
	}

	if !(used == len(data)-len(new_data)) {
		t.Errorf("Incorrect parsing used bits! Got %v, expected %v", used, len(data)-len(new_data))
	}

}

func TestParseNumBitsOperator(t *testing.T) {
	data := []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	new_data, p, used := parse_packet(data)

	operator, ok := p.(OperatorPacket)

	if !ok {
		t.Errorf("Packet illegal type!")
	}

	rest_value := BitSliceToInt(new_data)

	if !(rest_value == 0) {
		t.Errorf("Incorrect parsing rest! Got %v, expected %v", rest_value, 0)
	}

	if !(operator.version_ == 1) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", operator.version_, 1)
	}

	if !(operator.type_id_ == 6) {
		t.Errorf("Incorrect parsing type ID! Got %v, expected %v", operator.type_id_, 6)
	}

	if !(used == len(data)-len(new_data)) {
		t.Errorf("Incorrect parsing used bits! Got %v, expected %v", used, len(data)-len(new_data))
	}

	if !(len(operator.packets_) == 2) {
		t.Errorf("Incorrect parsing wrapped packet! Got %v, expected %v", len(operator.packets_), 2)
	}

	for i, pack := range operator.packets_ {
		literal, ok := pack.(LiteralPacket)

		if !ok {
			t.Errorf("Sub Packet illegal type! %v", i)
		}
		if i == 0 {
			if !(literal.value_ == 10) {
				t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 10)
			}
		}
		if i == 1 {
			if !(literal.value_ == 20) {
				t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 20)
			}
		}
	}

}

func TestParseNumPacketsOperator(t *testing.T) {
	data := []int{1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}
	new_data, p, used := parse_packet(data)

	operator, ok := p.(OperatorPacket)

	if !ok {
		t.Errorf("Packet illegal type!")
	}

	rest_value := BitSliceToInt(new_data)

	if !(rest_value == 0) {
		t.Errorf("Incorrect parsing rest! Got %v, expected %v", rest_value, 0)
	}

	if !(operator.version_ == 7) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", operator.version_, 7)
	}

	if !(operator.type_id_ == 3) {
		t.Errorf("Incorrect parsing type ID! Got %v, expected %v", operator.type_id_, 3)
	}

	if !(used == len(data)-len(new_data)) {
		t.Errorf("Incorrect parsing used bits! Got %v, expected %v", used, len(data)-len(new_data))
	}

	if !(len(operator.packets_) == 3) {
		t.Errorf("Incorrect parsing wrapped packet! Got %v, expected %v", len(operator.packets_), 2)
	}

	for i, pack := range operator.packets_ {
		literal, ok := pack.(LiteralPacket)

		if !ok {
			t.Errorf("Sub Packet illegal type! %v", i)
		}
		if i == 0 {
			if !(literal.value_ == 1) {
				t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 1)
			}
		}
		if i == 1 {
			if !(literal.value_ == 2) {
				t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 2)
			}
		}
		if i == 3 {
			if !(literal.value_ == 3) {
				t.Errorf("Incorrect parsing value! Got %v, expected %v", literal.value_, 3)
			}
		}
	}

}
