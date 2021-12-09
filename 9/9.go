package dec9

import (
	"sort"
)

func Solve1() int {

	level := getData("C:/Dev/Go/AdventOfCode-2021/9/data")

	sum, _ := level.GetLowPoints()

	return sum
}

func Solve2() int {

	level := getData("C:/Dev/Go/AdventOfCode-2021/9/data")

	_, low := level.GetLowPoints()

	sizes := []int{}

	for _, p := range low {
		sizes = append(sizes, level.ExploreBasin(p, 0))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}
