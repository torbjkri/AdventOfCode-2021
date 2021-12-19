package dec16

import "testing"
import "my_utils"

func Test16Data1(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata1")

	_, p, _ := parse_packet(Parse(data[0]))

	if !(p.getVersionSum() == 16) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", p.getVersionSum(), 16)
	}
}

func Test16Data2(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata1")

	_, p, _ := parse_packet(Parse(data[1]))

	if !(p.getVersionSum() == 12) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", p.getVersionSum(), 12)
	}
}

func Test16Data3(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata1")

	_, p, _ := parse_packet(Parse(data[2]))

	if !(p.getVersionSum() == 23) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", p.getVersionSum(), 23)
	}
}

func Test16Data4(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata1")

	_, p, _ := parse_packet(Parse(data[3]))

	if !(p.getVersionSum() == 31) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", p.getVersionSum(), 31)
	}
}

func Test16Data2_1(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[0]))

	val := p.getValue()

	if !(val == 3) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, 3)
	}
}

func Test16Data2_2(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[1]))

	val := p.getValue()

	want := 54

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_3(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[2]))

	val := p.getValue()

	want := 7

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_4(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[3]))

	val := p.getValue()

	want := 9

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_5(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[4]))

	val := p.getValue()

	want := 1

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_6(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[5]))

	val := p.getValue()

	want := 0

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_7(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[6]))

	val := p.getValue()

	want := 0

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}

func Test16Data2_8(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/testdata2")

	_, p, _ := parse_packet(Parse(data[7]))

	val := p.getValue()

	want := 1

	if !(val == want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", val, want)
	}
}
