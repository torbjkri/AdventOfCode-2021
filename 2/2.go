package dec2

import (
	"my_utils"
	"strconv"
	"strings"
)

func getData() ([]string, []int) {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/2/data")

	var result_value []int
	var result_direction []string

	for _, ln := range data {
		val := strings.Split(ln, " ")
		num, _ := strconv.Atoi(val[1])
		result_value = append(result_value, num)
		result_direction = append(result_direction, val[0])
	}

	return result_direction, result_value
}

func Solve1() int {
	direction, value := getData()

	hor := 0
	depth := 0

	for idx, dir := range direction {
		switch dir {
		case "forward":
			hor += value[idx]
		case "down":
			depth += value[idx]
		case "up":
			depth -= value[idx]
		}
	}

	return hor * depth
}

func Solve2() int {
	direction, value := getData()

	hor := 0
	depth := 0
	aim := 0

	for idx, dir := range direction {
		switch dir {
		case "forward":
			hor += value[idx]
			depth += +value[idx] * aim
		case "down":
			aim += value[idx]
		case "up":
			aim -= value[idx]
		}
	}

	return hor * depth
}
