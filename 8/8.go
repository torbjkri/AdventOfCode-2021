package dec8

func Solve1() int {
	input := getInput("C:/Dev/Go/AdventOfCode-2021/8/data")

	result := 0

	for _, elem := range input {
		result += elem.CountSimpleNumbers()
	}

	return result
}

func Solve2() int {

	input := getInput("C:/Dev/Go/AdventOfCode-2021/8/data")

	result := 0

	for _, elem := range input {
		decoded := elem.Decode()
		result += decoded[3] + (decoded[2]+(decoded[1]+decoded[0]*10)*10)*10
	}

	return result
}
