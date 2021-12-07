package dec7

import "testing"

func TestSolution1(t *testing.T) {
	solution := Solve1()

	if solution != 37 {
		t.Fatalf("WRONG! %d", solution)
	}
}

func TestSolution2(t *testing.T) {
	solution := Solve2()

	if solution != 168 {
		t.Fatalf("WRONG! %d", solution)
	}
}
