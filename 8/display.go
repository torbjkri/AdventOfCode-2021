package dec8

import (
	"sort"
)

var digits map[string]int = map[string]int{
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"bdefg":   6,
	"acf":     7,
	"abcdefg": 8,
	"abcdf":   9,
}

type Display struct {
	signal_pattern_ []string
	output_         []string

	segment_occurence_ map[rune]int
	maps_to_           map[string]int
}

func CreateDisplay(signal []string, output []string) *Display {
	result := new(Display)

	result.signal_pattern_ = signal
	result.output_ = output

	result.segment_occurence_ = map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
	}

	for _, signal := range result.signal_pattern_ {
		for _, char := range signal {
			result.segment_occurence_[char] += 1
		}
	}

	result.maps_to_ = map[string]int{}

	// Map all
	i := 0
	for i < len(signal) {
		l := len(signal[i])
		if l == 2 {
			result.maps_to_[signal[i]] = 1
			signal = remove(signal, i)
		} else if l == 3 {
			result.maps_to_[signal[i]] = 7
			signal = remove(signal, i)
		} else if l == 4 {
			result.maps_to_[signal[i]] = 4
			signal = remove(signal, i)
		} else if l == 7 {
			result.maps_to_[signal[i]] = 8
			signal = remove(signal, i)
		} else {

			input := map_to_occurences(signal[i], result.segment_occurence_)
			if equals(input, []int{4, 6, 7, 8, 8, 9}) {
				result.maps_to_[signal[i]] = 0
				signal = remove(signal, i)
			} else if equals(input, []int{4, 7, 7, 8, 8}) {
				result.maps_to_[signal[i]] = 2
				signal = remove(signal, i)
			} else if equals(input, []int{7, 7, 8, 8, 9}) {
				result.maps_to_[signal[i]] = 3
				signal = remove(signal, i)
			} else if equals(input, []int{6, 7, 7, 8, 9}) {
				result.maps_to_[signal[i]] = 5
				signal = remove(signal, i)
			} else if equals(input, []int{4, 6, 7, 7, 8, 9}) {
				result.maps_to_[signal[i]] = 6
				signal = remove(signal, i)
			} else if equals(input, []int{6, 7, 7, 8, 8, 9}) {
				result.maps_to_[signal[i]] = 9
				signal = remove(signal, i)
			}
		}
	}

	return result
}

func (d Display) Decode() []int {
	result := []int{}

	decode := func(in string, maps map[string]int) int {
		for k, elem := range maps {
			if k == in {
				return elem
			}
		}
		return -1
	}

	for _, elem := range d.output_ {
		num := decode(elem, d.maps_to_)

		result = append(result, num)
	}

	return result
}

func (d Display) CountSimpleNumbers() int {
	result := 0

	for _, elem := range d.output_ {
		l := len(elem)
		if l == 2 || l == 3 || l == 4 || l == 7 {
			result += 1
		}
	}

	return result
}

func map_to_occurences(input string, maps map[rune]int) []int {
	result := []int{}

	for _, char := range input {
		result = append(result, maps[char])
	}

	sort.Ints(result)
	return result
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, _ := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func remove(s []string, idx int) []string {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}
