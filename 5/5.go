package dec5

import (
	"my_utils"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

type Vents struct {
	clouds_ map[Point]int
}

func (v Vents) addCloud(position Point) {
	if _, ok := v.clouds_[position]; ok {
		v.clouds_[position] += 1
	} else {
		v.clouds_[position] = 1
	}
}

func (l Line) getLinePoints() []Point {
	var points []Point

	max := func(x1 int, x2 int) int {
		if x1 > x2 {
			return x1
		}
		return x2
	}
	min := func(x1 int, x2 int) int {
		if x1 > x2 {
			return x2
		}
		return x1
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if l.p1.x == l.p2.x {
		for i := min(l.p1.y, l.p2.y); i <= max(l.p1.y, l.p2.y); i++ {
			points = append(points, Point{l.p1.x, i})
		}
	} else if l.p1.y == l.p2.y {
		for i := min(l.p1.x, l.p2.x); i <= max(l.p1.x, l.p2.x); i++ {
			points = append(points, Point{i, l.p1.y})
		}
	} else {

		dx := (l.p2.x - l.p1.x) / abs(l.p2.x-l.p1.x)
		dy := (l.p2.y - l.p1.y) / abs(l.p2.y-l.p1.y)

		x := l.p1.x
		y := l.p1.y
		points = append(points, Point{x, y})
		for x != l.p2.x && y != l.p2.y {
			x += dx
			y += dy
			points = append(points, Point{x, y})

		}
	}

	return points
}

func getInput(filename string) []Line {
	data := my_utils.ReadFile(filename)

	var lines []Line

	for _, line := range data {
		str_input := strings.Split(line, " -> ")
		p1_str := strings.Split(str_input[0], ",")
		p1_x, _ := strconv.Atoi(p1_str[0])
		p1_y, _ := strconv.Atoi(p1_str[1])

		p1 := Point{p1_x, p1_y}

		p2_str := strings.Split(str_input[1], ",")
		p2_x, _ := strconv.Atoi(p2_str[0])
		p2_y, _ := strconv.Atoi(p2_str[1])

		p2 := Point{p2_x, p2_y}

		lines = append(lines, Line{p1, p2})
	}

	return lines
}

func generateTask1Points(lines []Line) []Point {
	var points []Point
	for _, line := range lines {
		if line.p1.x == line.p2.x || line.p1.y == line.p2.y {
			points = append(points, line.getLinePoints()...)
		}
	}
	return points
}

func generateTask2Points(lines []Line) []Point {
	var points []Point
	for _, line := range lines {
		points = append(points, line.getLinePoints()...)
	}
	return points
}

func Solve1() int {
	var t Vents
	t.clouds_ = map[Point]int{}

	input := getInput("C:/Dev/Go/AdventOfCode-2021/5/data")

	input_points := generateTask1Points(input)

	for _, point := range input_points {
		t.addCloud(point)
	}

	sum := 0

	for _, elem := range t.clouds_ {
		if elem > 1 {
			sum += 1
		}
	}

	return sum
}

func Solve2() int {
	var t Vents
	t.clouds_ = map[Point]int{}

	input := getInput("C:/Dev/Go/AdventOfCode-2021/5/data")

	input_points := generateTask2Points(input)

	for _, point := range input_points {
		t.addCloud(point)
	}

	sum := 0

	for _, elem := range t.clouds_ {
		if elem > 1 {
			sum += 1
		}
	}

	return sum
}
