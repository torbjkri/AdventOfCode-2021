package my_utils

import (
	"bufio"
	"log"
	"os"
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
