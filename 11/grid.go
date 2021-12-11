package dec11

import (
	"strconv"
	"strings"
)

type Grid struct {
	octopi_  [][]*Octopus
	width_   int
	height_  int
	flashes_ int
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
			if g.octopi_[i][j].CheckFlash() {
				g.flashes_ += 1
				g.SpreadFlash(i, j)
			}
		}
	}

}

func (g *Grid) PumpOctopus(row, col int) {
	if row > -1 && row < g.height_ && col > -1 && col < g.width_ {
		g.octopi_[row][col].Pump()
		if g.octopi_[row][col].CheckFlash() {
			g.flashes_ += 1
			g.SpreadFlash(row, col)
		}
	}
}

func (g *Grid) SpreadFlash(row, col int) {
	for i := row - 1; i < row+2; i++ {
		for j := col - 1; j < col+2; j++ {
			g.PumpOctopus(i, j)
		}
	}
}
