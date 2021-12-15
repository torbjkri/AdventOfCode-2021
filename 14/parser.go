package dec14

import "strings"

func ParseInput(data []string) (input string, rules map[[2]rune]rune) {

	input = data[0]

	for _, line := range data[2:] {
		elems := strings.Split(line, " ")
		rules[[2]rune{rune(elems[0][0]), rune(elems[0][1])}] = rune(elems[2][0])
	}

	return input, rules
}
