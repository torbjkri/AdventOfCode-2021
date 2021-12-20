package dec17

import "testing"
import "my_utils"
import "fmt"

func TestParser(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/17/data")

	res := Parse(data)
	want := []int{211, 232, -124, -69}

	if !equalSliceInt(res, want) {
		t.Errorf("Incorrect parsing version! Got %v, expected %v", res, want)
	}
}

func TestProbe(t *testing.T) {

	probe := Probe{x_: 0, y_: 0, vel_x_: 0, vel_y_: 2}

	probe.Update()

	if probe.x_ != 0 {
		t.Errorf("Incorrect x  value! Got %v, expected %v", probe.x_, 0)
	}

	if probe.y_ != 2 {
		t.Errorf("Incorrect y  value! Got %v, expected %v", probe.y_, 2)
	}

	if probe.vel_y_ != 1 {
		t.Errorf("Incorrect y  velocity! Got %v, expected %v", probe.vel_y_, 1)
	}

	if probe.vel_x_ != 0 {
		t.Errorf("Incorrect x velocity! Got %v, expected %v", probe.vel_x_, 0)
	}

	probe = Probe{x_: 0, y_: 0, vel_x_: 2, vel_y_: 0}

	probe.Update()

	if probe.x_ != 2 {
		t.Errorf("Incorrect x  value! Got %v, expected %v", probe.x_, 2)
	}

	if probe.y_ != 0 {
		t.Errorf("Incorrect y  value! Got %v, expected %v", probe.y_, 0)
	}

	if probe.vel_y_ != -1 {
		t.Errorf("Incorrect y  velocity! Got %v, expected %v", probe.vel_y_, -1)
	}

	if probe.vel_x_ != 1 {
		t.Errorf("Incorrect  x velocity! Got %v, expected %v", probe.vel_x_, 1)
	}
}

func TestData(t *testing.T) {
	data := []string{"target area: x=20..30, y=-10..-5"}

	input := Parse(data)

	fmt.Println(input)

	target := Target{x_min_: input[0], x_max_: input[1], y_min_: input[2], y_max_: input[3]}

	fmt.Println(target)

	total_max_y := 0

	for i := 0; i < 15; i++ {
		for j := -10; j < 11; j++ {

			max_y := 0
			p := Probe{x_: 0, y_: 0, vel_x_: i, vel_y_: j}

			for !p.ShouldStop(target) && !target.IsInside(p.x_, p.y_) {
				p.Update()
				if p.y_ > max_y {
					max_y = p.y_
				}
			}

			if target.IsInside(p.x_, p.y_) {
				if max_y > total_max_y {
					total_max_y = max_y
				}
			}

		}
	}

	if total_max_y != 45 {
		t.Errorf("Incorrect Max height! Got %v, expected %v", total_max_y, 45)
	}
}

func Test17A(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/17/data")

	input := Parse(data)

	target := Target{x_min_: input[0], x_max_: input[1], y_min_: input[2], y_max_: input[3]}

	fmt.Println(target)

	total_max_y := 0

	for i := 0; i < 150; i++ {
		for j := 0; j < 150; j++ {

			max_y := 0
			p := Probe{x_: 0, y_: 0, vel_x_: i, vel_y_: j}

			for !p.ShouldStop(target) && !target.IsInside(p.x_, p.y_) {
				p.Update()
				if p.y_ > max_y {
					max_y = p.y_
				}
			}

			if target.IsInside(p.x_, p.y_) {
				if max_y > total_max_y {
					total_max_y = max_y
				}
			}

		}
	}

	if total_max_y != 7626 {
		t.Errorf("Incorrect Max height! Got %v, expected %v", total_max_y, 7626)
	}
}

func Test17B(t *testing.T) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/17/data")

	input := Parse(data)

	target := Target{x_min_: input[0], x_max_: input[1], y_min_: input[2], y_max_: input[3]}

	fmt.Println(target)

	num_hits := 0

	for i := 0; i < 250; i++ {
		for j := -200; j < 250; j++ {
			p := Probe{x_: 0, y_: 0, vel_x_: i, vel_y_: j}

			for !p.ShouldStop(target) && !target.IsInside(p.x_, p.y_) {
				p.Update()
			}

			if target.IsInside(p.x_, p.y_) {
				num_hits += 1
			}

		}
	}

	if num_hits != 45 {
		t.Errorf("Incorrect Max height! Got %v, expected %v", num_hits, 45)
	}
}

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
