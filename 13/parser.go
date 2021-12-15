package dec13

import (
	"strconv"
	"strings"
)

type Command struct {
	direction_ int
	line_      int
}

func ParseInput(data []string) (input [][2]int, commands []Command) {

	for idx, elem := range data {
		if elem == "" {
			data = data[idx+1:]
			break
		}

		str_idx := strings.Split(elem, ",")

		x, _ := strconv.Atoi(str_idx[0])
		y, _ := strconv.Atoi(str_idx[1])

		input = append(input, [2]int{x, y})
	}

	for _, elem := range data {
		str := strings.Split(elem, " ")
		interesting := strings.Split(str[len(str)-1], "=")

		c := Command{}
		if interesting[0] == "x" {
			c.direction_ = 0
		} else {
			c.direction_ = 1
		}
		c.line_, _ = strconv.Atoi(interesting[1])

		commands = append(commands, c)
	}

	return input, commands
}
