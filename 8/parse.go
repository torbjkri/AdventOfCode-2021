package dec8

import (
	"my_utils"
	"strings"
)

func parseSignal(line string) []string {
	line_vec := strings.Split(line, " ")

	for i, _ := range line_vec {
		line_vec[i] = my_utils.SortString(line_vec[i])
	}

	return line_vec
}

func parseOutput(line string) []string {
	line_vec := strings.Split(line, " ")

	for i, _ := range line_vec {
		line_vec[i] = my_utils.SortString(line_vec[i])
	}

	return line_vec
}

func getInput(filename string) []*Display {
	data := my_utils.ReadFile(filename)

	var result []*Display

	for _, line := range data {
		splitted := strings.Split(line, " | ")
		signal := parseSignal(splitted[0])
		output := parseOutput(splitted[1])

		temp := CreateDisplay(signal, output)
		result = append(result, temp)

	}

	return result
}
