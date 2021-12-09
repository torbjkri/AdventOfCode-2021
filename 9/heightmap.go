package dec9

import (
	"my_utils"
	"strconv"
	"strings"
)

type HeightMap struct {
	data_ [][]int
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

funct (h *HeightMap) ExploreBasin(p Point) {
	
}

func CreateHeightMap(data []string) *HeightMap {
	HeightMap := new(HeightMap)

	for _, stringline := range data {
		var HeightMapline []int
		line := strings.Split(stringline, "")

		for _, strnum := range line {
			if strnum == "" {
				continue
			}
			num, _ := strconv.Atoi(strnum)
			HeightMapline = append(HeightMapline, num)
		}

		HeightMap.data_ = append(HeightMap.data_, HeightMapline)
	}
	return HeightMap
}

func getData(filename string) *HeightMap {
	data := my_utils.ReadFile(filename)

	level_map := CreateHeightMap(data)

	return level_map
}
