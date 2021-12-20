package dec18

import "testing"

func TestParse(t *testing.T) {
	input := "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"

	n, _ := ParseLine(input)

	result := n.GetString()
	if input != result {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", result, input)
	}
}
