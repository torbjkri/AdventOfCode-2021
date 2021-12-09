package dec9

import "testing"

func TestSolution1(t *testing.T) {
	solution := Solve1()

	if solution != 15 {
		t.Fatalf("WRONG! %d", solution)
	}
}

func TestSolution2(t *testing.T) {
	solution := Solve2()

	if solution != 1 {
		t.Fatalf("WRONG! %d", solution)
	}
}
