package my_utils

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	log.SetPrefix("[Utils] - ReadFile: ")
	log.SetFlags(0)

	if filename == "" {
		log.Fatal("Empty string")
		return nil
	}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	var data []string

	for s.Scan() {
		data = append(data, s.Text())
	}

	return data
}

func StringToNumList(list []string, separator string) []int {
	var result []int

	for _, line := range list {
		str_line := strings.Split(line, separator)
		for _, str_num := range str_line {
			num, _ := strconv.Atoi(str_num)
			result = append(result, num)
		}
	}

	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
