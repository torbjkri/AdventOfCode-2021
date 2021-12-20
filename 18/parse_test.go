package dec18

import "testing"

func TestParse(t *testing.T) {
	input := "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"

	n, _ := ParseLine(input)

	result := n.GetString()
	if input != result {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", result, input)
	}

	input = "[[[[0,7],4],[15,[0,13]]],[1,1]]"

	n, _ = ParseLine(input)

	result = n.GetString()
	if input != result {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", result, input)
	}
}
