package dec16

import "testing"

func equalSliceInt(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func equalSliceByte(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestParse(t *testing.T) {
	result := Parse("0F")

	want := []int{0, 0, 0, 0, 1, 1, 1, 1}

	if !equalSliceInt(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}

	result = Parse("F1")

	want = []int{1, 1, 1, 1, 0, 0, 0, 1}

	if !equalSliceInt(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}

	want = []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	result = Parse("38006F45291200")

	if !equalSliceInt(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}
}

func TestBitSliceToByte(t *testing.T) {
	result := BitSliceToBytes([]int{1, 1, 1, 1})

	var want []byte = []byte{15}

	if !equalSliceByte(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}

	result = BitSliceToBytes([]int{0, 0, 0, 0, 1, 1, 1, 1})

	want = []byte{15}

	if !equalSliceByte(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}

	result = BitSliceToBytes([]int{0, 0, 0, 0, 1, 1, 1, 1, 1})

	want = []byte{15, 1}

	if !equalSliceByte(result, want) {
		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
	}
}
