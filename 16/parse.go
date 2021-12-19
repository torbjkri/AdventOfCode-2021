package dec16

import (
	"encoding/hex"
)

func ByteToBitSlice(bs []byte) []int {
	result := make([]int, len(bs)*8)

	for i, b := range bs {
		for j := 0; j < 8; j++ {
			result[i*8+j] = int(b >> uint(7-j) & 0x01)
		}
	}

	return result
}

func BitSliceToBytes(ints []int) []byte {
	//var numbyte int = len(ints)/8 + 1

	result := []byte{}

	var temp byte = 0

	for i := 0; i < len(ints); i++ {
		temp = (temp << 1) | byte(ints[i])

		if i%8 == 7 {
			result = append(result, temp)
			temp = 0
		}
	}

	if len(ints)%8 != 0 {
		result = append(result, temp)
	}

	return result

}

func BitSliceToInt(ints []int) int {
	sum := 0
	multiplier := 1

	for i := len(ints) - 1; i >= 0; i-- {
		sum += multiplier * ints[i]
		multiplier *= 2
	}

	return sum
}

func Parse(data string) []int {
	if rr, err := hex.DecodeString(data); err != nil {
		panic(err)
	} else {
		result := ByteToBitSlice(rr)
		return result
	}
	return []int{}
}
