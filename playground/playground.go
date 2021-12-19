package main

import "fmt"

type A interface {
	getName() string
}

type B struct {
}

func (b B) getName() string {
	return "B"
}

func (b B) getExtra() string {
	return "Extra B"
}

type C struct {
}

func (c C) getName() string {
	return "C"
}

func main() {
	d := []A{}

	b := B{}
	c := C{}

	d = append(d, b)
	d = append(d, c)

	for _, elem := range d {
		fmt.Println(elem.getName())

	}
}
