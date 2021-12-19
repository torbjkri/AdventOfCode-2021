package dec16

import "my_utils"

func Solve1() int {

	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/data")

	_, p, _ := parse_packet(Parse(data[0]))

	return p.getVersionSum()
}

func Solve2() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/16/data")

	_, p, _ := parse_packet(Parse(data[0]))

	return p.getValue()
}
