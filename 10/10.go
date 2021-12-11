package dec10

import (
	"my_utils"
	"sort"
)

func Solve1() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/10/data")

	sum := 0

	for _, line := range data {
		sum += CheckCorruptedLineScore(line)
	}

	return sum
}

func Solve2() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/10/data")

	scores := []int{}

	for _, line := range data {
		score := CheckIncompleteLineScore(line)

		if score != 0 {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	middle := scores[len(scores)/2]

	return middle
}
