package dec13

import (
	"fmt"
	"my_utils"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func unique(intSlice [][2]int) [][2]int {
	keys := make(map[[2]int]bool)
	list := [][2]int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func handleCommand(data [][2]int, cmd Command) {
	for idx, _ := range data {
		if data[idx][cmd.direction_] > cmd.line_ {
			data[idx][cmd.direction_] -= 2 * abs(data[idx][cmd.direction_]-cmd.line_)
		}
	}
}

func Solve1() int {
	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/13/data")

	in, cmds := ParseInput(data)

	handleCommand(in, cmds[0])

	result := unique(in)

	return len(result)
}

func findMax(in [][2]int) [2]int {
	max := 0
	may := 0

	for _, elem := range in {
		if elem[0] > max {
			max = elem[0]
		}
		if elem[1] > may {
			may = elem[1]
		}
	}

	return [2]int{max, may}
}

func Solve2() int {

	data := my_utils.ReadFile("C:/Dev/Go/AdventOfCode-2021/13/data")

	in, cmds := ParseInput(data)

	for _, cmd := range cmds {
		handleCommand(in, cmd)
	}

	result := unique(in)

	max := findMax(result)

	fmt.Println(max)

	out := [][]rune{}

	for row := 0; row <= max[1]; row++ {
		line := []rune{}
		for col := 0; col <= max[0]; col++ {
			line = append(line, '.')
		}
		out = append(out, line)
	}

	for _, elem := range result {
		fmt.Println(elem)
		out[elem[1]][elem[0]] = '#'
	}

	for row := 0; row <= max[1]; row++ {
		for col := 0; col <= max[0]; col++ {
			fmt.Printf("%c", out[row][col])
		}
		fmt.Println()
	}

	return 1
}
