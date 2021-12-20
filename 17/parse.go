package dec17

import (
	"strconv"
	"strings"
)

func Parse(data []string) []int {

	input := strings.Replace(data[0], "target area: ", "", 1)
	input = strings.Replace(input, "x=", "", 1)
	input = strings.Replace(input, " y=", "", 1)
	input = strings.ReplaceAll(input, "..", ",")

	in_vec := strings.Split(input, ",")
	res := []int{}

	for _, elem := range in_vec {
		num, _ := strconv.Atoi(elem)
		res = append(res, num)
	}

	return res
}
