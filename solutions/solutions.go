package main

import (
	"dec1"
	"dec2"
	"dec3"
	"dec4"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func getDec1() [2]func() int {
	return [2]func() int{dec1.Solve1, dec1.Solve2}
}

func getDec2() [2]func() int {
	return [2]func() int{dec2.Solve1, dec2.Solve2}
}

func getDec3() [2]func() int {
	return [2]func() int{dec3.Solve1, dec3.Solve2}
}

func getDec4() [2]func() int {
	return [2]func() int{dec4.Solve1, dec4.Solve2}
}

func executeDay(solution [2]func() int, day int) {
	prefix := fmt.Sprintf("[AoC] December %d: ", day)
	l := log.New(os.Stdout, prefix, 0)
	start := time.Now()
	l.Printf("Solution  A: %d", solution[0]())
	timeA := time.Since(start).Microseconds()
	l.Println("Part A took: ", timeA, "us")

	start = time.Now()
	l.Printf("Solution B: %d", solution[1]())
	timeB := time.Since(start).Microseconds()
	l.Printf("Part B took: %d us", timeB)
	l.Printf("Took in total : %d us", timeA+timeB)
}

func main() {

	var input int
	var err error
	err = nil

	if input, err = strconv.Atoi(os.Args[1]); err != nil {
		panic("Invalid Input")
	}

	var days [][2]func() int

	switch input {
	case 0:
		days = append(days, getDec1())
		days = append(days, getDec2())
		days = append(days, getDec3())
		days = append(days, getDec4())
	case 1:
		days = append(days, getDec1())
	case 2:
		days = append(days, getDec1())
	case 3:
		days = append(days, getDec1())
	case 4:
		days = append(days, getDec1())

	}

	if input > 0 {
		executeDay(days[0], input)
	} else {
		for idx, _ := range days {
			executeDay(days[idx], idx+1)
			fmt.Println()
		}

	}

}
