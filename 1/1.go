package dec1

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"os"
	"strconv"
)

func Solve1() int {
	data, err := ReadFile("C:/Dev/Go/AdventOfCodode-2021/1/data")

	if err != nil {
		log.Fatal("Unable to read file")
	}

	count := 0

	for idx, num := range data {
		if idx == 0 {
			continue
		}
		if num > data[idx-1] {
			count += 1
		}
	}

	return count
}

func Solve2() int {
	data, err := ReadFile("C:/Dev/Go/AdventOfCodode-2021/1/data")

	if err != nil {
		log.Fatal("Unable to read file")
	}

	count := 0

	sum := 0

	for idx, num := range data {
		if idx < 3 {
			sum += num
			continue
		}
		tmp := sum + num - data[idx-3]

		if tmp > sum {
			count += 1
		}
		sum = tmp
	}

	return count
}

func ReadFile(filename string) ([]int, error) {
	if filename == "" {
		return nil, errors.New("Invalid string")
	}

	fptr := flag.String("fpath", filename, "File path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	var data []int

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		data = append(data, num)
	}

	return data, nil
}
