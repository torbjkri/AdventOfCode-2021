package dec18

import "testing"

func TestNodeExplode(t *testing.T) {
	input := "[[6,[5,[4,[3,2]]]],1]"

	n, _ := ParseLine(input)

	exploded, _ := n.explode(0)

	result := n.GetString()

	want := "[[6,[5,[7,0]]],3]"

	if !exploded {
		t.Errorf("Incorrect Explosion! Got %v, expected %v", exploded, true)
	}

	if result != want {
		t.Errorf("Incorrect explosion result! Input %v Got %v, expected %v", input, result, want)
	}

	//// More examples
	input = "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"

	n, _ = ParseLine(input)

	exploded, _ = n.explode(0)

	result = n.GetString()

	want = "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"

	if !exploded {
		t.Errorf("Incorrect Explosion! Got %v, expected %v", exploded, true)
	}

	if result != want {
		t.Errorf("Incorrect explosion result! Input %v Got %v, expected %v", input, result, want)
	}

}
