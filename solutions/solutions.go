package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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
	days = append(days, getDec10())
	days = append(days, getDec11())
	days = append(days, getDec12())
	days = append(days, getDec13())
	days = append(days, getDec14())
	days = append(days, getDec15())
	days = append(days, getDec16())
	days = append(days, getDec17())

	if day > 0 {
		executeDay(days[day-1], day, times)
	} else {
		for idx, _ := range days {
			executeDay(days[idx], idx+1, times)
			fmt.Println()
		}
	}

}
