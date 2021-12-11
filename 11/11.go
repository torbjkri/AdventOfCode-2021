package dec11

import (
	"my_utils"
)

func Solve1() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/11/testdata")

	g := CreateGrid(data)

	for i := 0; i < 100; i++ {
		g.Update()
	}

	return g.flashes_
}

func Solve2() int {
	return 1
}
