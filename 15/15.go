package dec15

import "my_utils"

func Solve1() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/15/data")

	g := NewGrid(data)

	return g.FindShortestPath()
}

func Solve2() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/15/data")

	g := NewGrid5x(data)

	return g.FindShortestPath()
}
