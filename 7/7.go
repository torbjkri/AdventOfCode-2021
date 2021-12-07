package dec7

import (
	"math"
	"my_utils"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetInitialData(filename string) []int {
	data := my_utils.ReadFile(filename)

	return my_utils.StringToNumList(data, ",")
}

func totalDistance(list []int, target int) int {
	sum := 0

	for _, elem := range list {
		sum += abs(elem - target)
	}

	return sum
}

func getAverage(list []int) int {
	sum := 0
	for _, elem := range list {
		sum += elem
	}

	return int(math.Round(float64(sum) / float64(len(list))))
}

func getWeightedDistance(list []int, target int) int {
	sum_calc := func(a, b int) int {
		return (abs(a-b) + 1) * (a + b) / 2
	}

	sum := 0
	for _, elem := range list {
		sum += sum_calc(abs(target-elem), 0)
	}

	return sum
}

func Solve1() int {

	crabbies := GetInitialData("C:/Dev/Go/AdventOfCode-2021/7/testdata")

	sort.Slice(crabbies, func(i, j int) bool { return crabbies[i] < crabbies[j] })

	var target_idx int

	target_idx = len(crabbies) / 2

	target := crabbies[target_idx]

	return totalDistance(crabbies, target)
}

func Solve2() int {

	crabbies := GetInitialData("C:/Dev/Go/AdventOfCode-2021/7/testdata")
	return getWeightedDistance(crabbies, getAverage(crabbies)+1)
}
