package dec9

import (
	"my_utils"
	"strconv"
	"strings"
)

type HeightMap struct {
	data_     [][]int
	explored_ [][]bool
}

type Point struct {
	row_ int
	col_ int
}

func (h *HeightMap) GetLowPoints() (sum int, result []Point) {

	max_rows := len(h.data_)
	max_cols := len(h.data_[1])

	for i, row := range h.data_ {
		for j, val := range row {
			if val == 9 {
				continue
			}

			// check above
			if i > 0 {
				if h.data_[i-1][j] <= val {
					continue
				}

			}
			if i < max_rows-1 {
				if h.data_[i+1][j] <= val {
					continue
				}
			}
			if j > 0 {
				if h.data_[i][j-1] <= val {
					continue
				}

			}
			if j < max_cols-1 {
				if h.data_[i][j+1] <= val {
					continue
				}

			}

			result = append(result, Point{i, j})
			sum += val + 1

		}
	}

	return sum, result
}

func (h *HeightMap) ExploreBasin(p Point, sum int) int {
	// check above
	i := p.row_
	j := p.col_
	val := h.data_[i][j]

	max_rows := len(h.data_)
	max_cols := len(h.data_[0])

	if i > 0 {
		v := h.data_[i-1][j]
		e := h.explored_[i-1][j]
		if v > val && v != 9 && !e {
			sum = h.ExploreBasin(Point{row_: i - 1, col_: j}, sum)
		}

	}
	if i < max_rows-1 {
		v := h.data_[i+1][j]
		e := h.explored_[i+1][j]
		if v > val && v != 9 && !e {
			sum = h.ExploreBasin(Point{row_: i + 1, col_: j}, sum)
		}
	}
	if j > 0 {
		v := h.data_[i][j-1]
		e := h.explored_[i][j-1]
		if v > val && v != 9 && !e {
			sum = h.ExploreBasin(Point{row_: i, col_: j - 1}, sum)
		}

	}
	if j < max_cols-1 {
		v := h.data_[i][j+1]
		e := h.explored_[i][j+1]
		if v > val && v != 9 && !e {
			sum = h.ExploreBasin(Point{row_: i, col_: j + 1}, sum)
		}
	}

	h.explored_[i][j] = true

	return sum + 1
}

func CreateHeightMap(data []string) *HeightMap {
	HeightMap := new(HeightMap)

	for _, stringline := range data {
		var HeightMapline []int
		var boolLine []bool
		line := strings.Split(stringline, "")

		for _, strnum := range line {
			if strnum == "" {
				continue
			}
			num, _ := strconv.Atoi(strnum)
			HeightMapline = append(HeightMapline, num)
			boolLine = append(boolLine, false)
		}

		HeightMap.data_ = append(HeightMap.data_, HeightMapline)
		HeightMap.explored_ = append(HeightMap.explored_, boolLine)
	}
	return HeightMap
}

func getData(filename string) *HeightMap {
	data := my_utils.ReadFile(filename)

	level_map := CreateHeightMap(data)

	return level_map
}
