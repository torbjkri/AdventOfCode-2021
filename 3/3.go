package dec3

import (
	"fmt"
	"my_utils"
	"strconv"
)

func getData(filename string) ([]uint64, int) {
	data := my_utils.ReadFile(filename)

	numbits := len([]rune(data[0]))

	var result []uint64
	for _, strnum := range data {
		if num, err := strconv.ParseUint(strnum, 2, numbits); err == nil {
			result = append(result, num)
		} else {
			fmt.Println("Error converting binary data")
			fmt.Println(err)
			return nil, 0

		}
	}

	return result, numbits
}

func setBit(n uint64, pos int) uint64 {
	n |= (1 << pos)
	return n
}

func hasBit(val uint64, bit int) bool {
	res := val & (1 << bit)
	return res > 0
}

func update(val uint64, bit int) int {
	if hasBit(val, bit) {
		return 1
	} else {
		return -1
	}
}

func bitSum(values []uint64, bit int) int {
	sum := 0
	for _, val := range values {
		sum += update(val, bit)
	}
	return sum
}

func Solve1() int {
	data, numbits := getData("C:/Dev/Go/AdventOfCode-2021/3/data")

	var gamma uint64 = 0
	var beta uint64 = 0

	total := make([]int, numbits)

	for i := 0; i < numbits; i++ {
		total[i] = bitSum(data, i)
	}

	for i := 0; i < numbits; i++ {
		if total[i] >= 0 {
			gamma = setBit(gamma, i)
		} else {
			beta = setBit(beta, i)
		}
	}

	return int(gamma * beta)
}

func scrubData(data []uint64, i int, keep_value bool) []uint64 {
	result := []uint64{}

	for _, value := range data {
		has_value := hasBit(value, i)

		if (has_value && keep_value) || (!has_value && !keep_value) {
			result = append(result, value)
		}
	}

	return result
}

func Solve2() int {
	data, numbits := getData("C:/Dev/Go/AdventOfCode-2021/3/data")

	o2_data := data

	var o2_rating uint64 = 0

	// Check first

	for i := numbits - 1; i >= 0; i-- {
		total := bitSum(o2_data, i)
		o2_data = scrubData(o2_data, i, total >= 0)
		if len(o2_data) == 1 {
			break
		}
	}
	o2_rating = o2_data[0]

	co2_data := data

	var co2_rating uint64 = 0

	for i := numbits - 1; i >= 0; i-- {
		total := bitSum(co2_data, i)
		co2_data = scrubData(co2_data, i, total < 0)
		if len(co2_data) == 1 {
			break
		}
	}
	co2_rating = co2_data[0]

	return int(o2_rating * co2_rating)
}
