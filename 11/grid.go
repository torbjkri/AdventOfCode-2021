package dec11

import (
	"fmt"
	"image"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Grid struct {
	octopi_  [][]*Octopus
	width_   int
	height_  int
	flashes_ int

	img_ *image.Gray
}

func CreateGrid(data []string) *Grid {

	g := new(Grid)
	g.flashes_ = 0
	g.octopi_ = [][]*Octopus{}

	for _, line := range data {
		line_vec := strings.Split(line, "")

		octline := []*Octopus{}

		for _, elem := range line_vec {
			num, _ := strconv.Atoi(elem)
			octline = append(octline, CreateOctopus(num))
		}

		g.octopi_ = append(g.octopi_, octline)
	}

	g.height_ = len(g.octopi_)
	g.width_ = len(g.octopi_[0])

	g.img_ = image.NewGray(image.Rect(0, 0, g.width_, g.height_))

	return g
}

func CreateRandomGrid(sizey int, sizex int) *Grid {

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	g := new(Grid)
	g.flashes_ = 0
	g.octopi_ = [][]*Octopus{}

	for i := 0; i < sizey; i++ {

		octline := []*Octopus{}

		for j := 0; j < sizex; j++ {
			octline = append(octline, CreateOctopus(r2.Intn(9)))
		}

		g.octopi_ = append(g.octopi_, octline)
	}

	g.height_ = len(g.octopi_)
	g.width_ = len(g.octopi_[0])

	g.img_ = image.NewGray(image.Rect(0, 0, g.width_, g.height_))

	return g
}

func (g *Grid) Update() {
	// update state
	for i := 0; i < g.height_; i++ {
		for j := 0; j < g.width_; j++ {
			g.octopi_[i][j].Update()
		}
	}

	// Check for flashes and do the dance
	for i := 0; i < g.height_; i++ {
		for j := 0; j < g.width_; j++ {
			g.CheckFlash(i, j)
		}
	}

}

func (g *Grid) PumpOctopus(row, col int) {
	if row > -1 && row < g.height_ && col > -1 && col < g.width_ {
		g.octopi_[row][col].Pump()
		g.CheckFlash(row, col)
	}
}

func (g *Grid) SpreadFlash(row, col int) {
	for i := row - 1; i < row+2; i++ {
		for j := col - 1; j < col+2; j++ {
			g.PumpOctopus(i, j)
		}
	}
}

func (g *Grid) CheckFlash(row, col int) {
	if g.octopi_[row][col].CheckFlash() {
		g.flashes_ += 1
		g.SpreadFlash(row, col)
	}
}

func (g *Grid) AllFlashed() bool {
	for i := 0; i < g.height_; i++ {
		for j := 0; j < g.width_; j++ {
			if !g.octopi_[i][j].HasFlashed() {
				return false
			}
		}
	}
	return true
}

func (g *Grid) UpdateImage() {
	for i := 0; i < g.height_; i++ {
		for j := 0; j < g.width_; j++ {
			g.img_.SetGray(i, j, g.octopi_[i][j].GetEnergyColor())
		}
	}
}

func (g *Grid) Visualize() {
	for i := 0; i < g.height_; i++ {
		for j := 0; j < g.width_; j++ {
			e := g.octopi_[i][j].energy_
			if e == 0 {
				fmt.Print("0")
			} else if e > 7 {
				fmt.Print("-")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(100 * time.Millisecond)
}
