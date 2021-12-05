package main

import (
	"dec1"
	"dec2"
	"dec3"
	"dec4"
	"fmt"
	"log"
	"os"
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

	day := getDec1()
	executeDay(day, 1)

}
