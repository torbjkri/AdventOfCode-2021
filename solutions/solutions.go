package main

import (
	"dec1"
	"dec2"
	"dec3"
	"dec4"
	"dec5"
	"dec6"
	"dec7"
	"dec8"
	"dec9"
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

func getDec5() [2]func() int {
	return [2]func() int{dec5.Solve1, dec5.Solve2}
}

func getDec6() [2]func() int {
	return [2]func() int{dec6.Solve1, dec6.Solve2}
}

func getDec7() [2]func() int {
	return [2]func() int{dec7.Solve1, dec7.Solve2}
}

func getDec8() [2]func() int {
	return [2]func() int{dec8.Solve1, dec8.Solve2}
}

func getDec9() [2]func() int {
	return [2]func() int{dec9.Solve1, dec9.Solve2}
}

func executeDay(solution [2]func() int, day int, times int) {
	prefix := fmt.Sprintf("[AoC] December %d: ", day)
	l := log.New(os.Stdout, prefix, 0)
	l.Printf("Solution A: %d", solution[0]())
	start := time.Now()
	//l.Printf("Solution  A: %d", solution[0]())
	for i := 0; i < times; i++ {
		solution[0]()
	}
	timeA := time.Since(start).Microseconds()
	l.Println("Part A took: ", timeA/int64(times), "us")

	l.Printf("Solution B: %d", solution[1]())
	start = time.Now()
	for i := 0; i < times; i++ {
		solution[1]()
	}
	timeB := time.Since(start).Microseconds()
	l.Printf("Part B took: %d us", timeB/int64(times))
	l.Printf("Took in total : %d us", (timeA+timeB)/int64(times))
}

func main() {

	var day int
	var times int
	var err error
	err = nil

	if day, err = strconv.Atoi(os.Args[1]); err != nil {
		panic("Invalid Input")
	}

	if times, err = strconv.Atoi(os.Args[2]); err != nil {
		panic("Invalid Input")
	}

	var days [][2]func() int

	days = append(days, getDec1())
	days = append(days, getDec2())
	days = append(days, getDec3())
	days = append(days, getDec4())
	days = append(days, getDec5())
	days = append(days, getDec6())
	days = append(days, getDec7())
	days = append(days, getDec8())
	days = append(days, getDec9())

	if day > 0 {
		executeDay(days[day-1], day, times)
	} else {
		for idx, _ := range days {
			executeDay(days[idx], idx+1, times)
			fmt.Println()
		}
	}

}
