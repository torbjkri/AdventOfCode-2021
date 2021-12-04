package main

import (
	"dec1"
	"dec2"
	"dec3"
	"dec4"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("December 1 - A: ", dec1.Solve1())
	timeA := time.Since(start).Microseconds()
	fmt.Println("Dec 1 part A took: ", timeA, "us")

	start = time.Now()
	fmt.Println("December 1 - B: ", dec1.Solve2())
	timeB := time.Since(start).Microseconds()
	fmt.Println("December 1 part B took: ", timeB, "us")
	fmt.Println("December 1 took in total : ", timeA+timeB, "us")
	fmt.Println()
	fmt.Println()

	//// December 2
	start = time.Now()
	fmt.Println("December 2 - A: ", dec2.Solve1())
	timeA = time.Since(start).Microseconds()
	fmt.Println("Dec 2 part A took: ", timeA, "us")

	start = time.Now()
	fmt.Println("December 2 - B: ", dec2.Solve2())
	timeB = time.Since(start).Microseconds()
	fmt.Println("December 2 part B took: ", timeB, "us")
	fmt.Println("December 2 took in total : ", timeA+timeB, "us")
	fmt.Println()
	fmt.Println()
	// December 3
	start = time.Now()
	fmt.Println("December 3 - A: ", dec3.Solve1())
	timeA = time.Since(start).Microseconds()
	fmt.Println("Dec 23 part A took: ", timeA, "us")

	start = time.Now()
	fmt.Println("December 3 - B: ", dec3.Solve2())
	timeB = time.Since(start).Microseconds()
	fmt.Println("December 3 part B took: ", timeB, "us")
	fmt.Println("December 3 took in total : ", timeA+timeB, "us")
	fmt.Println()
	fmt.Println()
	// December 4
	start = time.Now()
	fmt.Println("December 4 - A: ", dec4.Solve1())
	timeA = time.Since(start).Microseconds()
	fmt.Println("Dec 4 part A took: ", timeA, "us")

	start = time.Now()
	fmt.Println("December 4 - B: ", dec4.Solve2())
	timeB = time.Since(start).Microseconds()
	fmt.Println("December 4 part B took: ", timeB, "us")
	fmt.Println("December 4 took in total : ", timeA+timeB, "us")
	fmt.Println()
	fmt.Println()
}
