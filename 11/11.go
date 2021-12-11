package dec11

import (
	"my_utils"
)

func Solve1() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/11/data")

	g := CreateGrid(data)

	for i := 0; i < 100; i++ {
		g.Update()
	}

	return g.flashes_
}

func Solve2() int {
	//data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/11/data")

	g := CreateRandomGrid(20, 100)

	for i := 1; ; i++ {
		g.Update()
		if g.AllFlashed() {
			return i
		}
		g.UpdateImage()
		g.Visualize()
	}

	return 0
}
