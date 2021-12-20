package dec18

import "testing"
import "my_utils"

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

func TestNodeSplit(t *testing.T) {
	input := "[[[[0,7],4],[15,[0,13]]],[1,1]]"

	n, _ := ParseLine(input)

	splitted := n.split()

	result := n.GetString()

	want := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"

	if !splitted {
		t.Errorf("Incorrect Splitting! Got %v, expected %v", splitted, true)
	}

	if result != want {
		t.Errorf("Incorrect split result! Input %v Got %v, expected %v", input, result, want)
	}
}

func TestNodeAdd(t *testing.T) {
	input1 := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	input2 := "[1,1]"

	n, _ := ParseLine(input1)
	n2, _ := ParseLine(input2)

	n = AddNodes(n, n2)

	result := n.GetString()

	want := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"

	if result != want {
		t.Errorf("Incorrect split result! Input %v, %v Got %v, expected %v", input1, input2, result, want)
	}
}

func TestHomeworkATestdata(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/18/testdata")

	n, _ := ParseLine(data[0])

	for _, line := range data[1:] {
		n2, _ := ParseLine(line)

		n = AddNodes(n, n2)

		for {
			if exploded, _ := n.explode(0); exploded == true {
				continue
			}
			if splitted := n.split(); splitted == true {
				continue
			}
			break
		}
	}
	result := n.GetString()
	want := "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"

	if result != want {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", result, want)
	}

	magn := n.getMagnitude()
	want_magn := 4140
	if magn != want_magn {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", magn, want_magn)
	}
}

func TestHomeworkA(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/18/data")

	n, _ := ParseLine(data[0])

	for _, line := range data[1:] {
		n2, _ := ParseLine(line)

		n = AddNodes(n, n2)

		for {
			if exploded, _ := n.explode(0); exploded == true {
				continue
			}
			if splitted := n.split(); splitted == true {
				continue
			}
			break
		}
	}
	result := n.GetString()
	want := "[[[[6,7],[7,7]],[[7,7],[7,7]]],[[[0,7],[8,8]],[[8,7],[6,1]]]]"

	if result != want {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", result, want)
	}

	magn := n.getMagnitude()
	want_magn := 3892
	if magn != want_magn {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", magn, want_magn)
	}
}

func TestHomeworkBTest(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/18/testdata")

	largest_mag := 0
	var n_best *Node = nil

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			n1, _ := ParseLine(data[i])
			n2, _ := ParseLine(data[j])

			n3 := AddNodes(n1, n2)

			for {
				if exploded, _ := n3.explode(0); exploded == true {
					continue
				}
				if splitted := n3.split(); splitted == true {
					continue
				}
				break
			}
			mag := n3.getMagnitude()
			if mag > largest_mag {
				largest_mag = mag
				n_best = n3
			}

			n4, _ := ParseLine(data[i])
			n5, _ := ParseLine(data[j])
			n6 := AddNodes(n5, n4)

			for {
				if exploded, _ := n6.explode(0); exploded == true {
					continue
				}
				if splitted := n6.split(); splitted == true {
					continue
				}
				break
			}

			mag = n6.getMagnitude()
			if mag > largest_mag {
				largest_mag = mag
				n_best = n6
			}

		}
	}

	result := n_best.GetString()
	want := "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]"

	if result != want {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", result, want)
	}

	magn := n_best.getMagnitude()
	want_magn := 3993
	if magn != want_magn {
		t.Errorf("Incorrect split result! Got %v, expected %v", magn, want_magn)
	}
}

func TestHomeworkB(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/18/data")

	largest_mag := 0
	var n_best *Node = nil

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			n1, _ := ParseLine(data[i])
			n2, _ := ParseLine(data[j])

			n3 := AddNodes(n1, n2)

			for {
				if exploded, _ := n3.explode(0); exploded == true {
					continue
				}
				if splitted := n3.split(); splitted == true {
					continue
				}
				break
			}
			mag := n3.getMagnitude()
			if mag > largest_mag {
				largest_mag = mag
				n_best = n3
			}

			n4, _ := ParseLine(data[i])
			n5, _ := ParseLine(data[j])
			n6 := AddNodes(n5, n4)

			for {
				if exploded, _ := n6.explode(0); exploded == true {
					continue
				}
				if splitted := n6.split(); splitted == true {
					continue
				}
				break
			}

			mag = n6.getMagnitude()
			if mag > largest_mag {
				largest_mag = mag
				n_best = n6
			}

		}
	}

	result := n_best.GetString()
	want := "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]"

	if result != want {
		t.Errorf("Incorrect split result! Input Got %v, expected %v", result, want)
	}

	magn := n_best.getMagnitude()
	want_magn := 3993
	if magn != want_magn {
		t.Errorf("Incorrect split result! Got %v, expected %v", magn, want_magn)
	}
}
