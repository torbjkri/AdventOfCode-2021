package dec1

import (
	"my_utils"
	"strconv"
)

func getData(filename string) []int {
	data := my_utils.ReadFile(filename)

	var result []int

	for _, ln := range data {
		num, _ := strconv.Atoi(ln)
		result = append(result, num)
	}

	return result
}

func Solve1() int {
	data := getData("C:/Dev/Go/AdventOfCode-2021/1/data")

	count := 0

	for idx, num := range data {
		if idx == 0 {
			continue
		}
		if num > data[idx-1] {
			count += 1
		}
	}

	return count
}

func Solve2() int {
	data := getData("C:/Dev/Go/AdventOfCode-2021/1/data")

	count := 0

	sum := 0

	for idx, num := range data {
		if idx < 3 {
			sum += num
			continue
		}
		tmp := sum + num - data[idx-3]

		if tmp > sum {
			count += 1
		}
		sum = tmp
	}

	return count
}
