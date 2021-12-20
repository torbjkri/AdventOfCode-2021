package dec18

import "strconv"

func ParseLine(data string) (*Node, string) {
	n := new(Node)
	n.is_num_ = false
	n.value_ = 0
	n.left_ = nil
	n.right_ = nil
	n.parent_ = nil

	if data[0] == '[' {
		n.left_, data = ParseLine(data[1:]) // removing [
		n.left_.parent_ = n
	}

	if data[0] == ',' {
		n.right_, data = ParseLine(data[1:]) // removing ,
		n.right_.parent_ = n
	}

	if data[0] == ']' {
		return n, data[1:] // removing ]
	}

	num, _ := strconv.Atoi(string(data[0]))

	n.is_num_ = true
	n.value_ = num

	return n, data[1:] // removing the number

}
