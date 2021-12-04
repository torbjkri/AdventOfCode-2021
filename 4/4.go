package dec4

import (
	"fmt"
	"my_utils"
	"strconv"
	"strings"
)

type Field struct {
	num_    int
	ticked_ bool
}

type Board struct {
	data_     [][]Field
	data_sum_ int
}

func (b *Board) CheckInput(input int) (bool, [2]int) {
	for row_idx, _ := range b.data_ {
		for col_idx, _ := range b.data_[row_idx] {
			if b.data_[row_idx][col_idx].num_ == input {
				b.data_[row_idx][col_idx].ticked_ = true
				b.data_sum_ = b.data_sum_ - input
				return true, [2]int{row_idx, col_idx}
			}
		}
	}
	return false, [2]int{-1, -1}
}

func (b *Board) checkBingoRow(idx int) bool {
	for i := 0; i < 5; i++ {
		if b.data_[idx][i].ticked_ == false {
			return false
		}
	}

	return true
}

func (b *Board) checkBingoCol(idx int) bool {
	for i := 0; i < 5; i++ {
		if b.data_[i][idx].ticked_ == false {
			return false
		}
	}

	return true
}

func (b *Board) CheckBingo(indexes [2]int) bool {
	if bingo := b.checkBingoRow(indexes[0]); bingo == true {
		return bingo
	} else if bingo := b.checkBingoCol(indexes[1]); bingo == true {
		return bingo
	}
	return false
}

func CreateBoard(data []string) Board {
	var board Board
	sum := 0
	for _, stringline := range data {
		line := strings.Split(stringline, " ")
		var boardline []Field

		for _, strnum := range line {
			if strnum == "" {
				continue
			}
			num, _ := strconv.Atoi(strnum)
			sum += num
			field := Field{num_: num, ticked_: false}
			boardline = append(boardline, field)
		}

		board.data_ = append(board.data_, boardline)
	}
	board.data_sum_ = sum
	return board
}

func createInput(data string) []int {
	strinput := strings.Split(data, ",")
	var input []int
	for _, strnum := range strinput {
		num, _ := strconv.Atoi(strnum)
		input = append(input, num)
	}
	return input
}

func createBoards(data []string) []*Board {
	var boards []*Board
	var boardData []string
	for _, strline := range data {
		if strline == "" {
			b := CreateBoard(boardData)
			boards = append(boards, &b)
			boardData = nil
		} else {
			boardData = append(boardData, strline)
		}
	}
	b := CreateBoard(boardData)
	boards = append(boards, &b)
	return boards
}

func getData(filename string) ([]int, []*Board) {
	data := my_utils.ReadFile(filename)

	input := createInput(data[0])

	boards := createBoards(data[2:len(data)])

	return input, boards
}

func Solve1() int {
	input, boards := getData("C:/Dev/Go/AdventOfCode-2021/4/data")

	for _, in := range input {
		for board_idx, _ := range boards {
			if hit, idx := boards[board_idx].CheckInput(in); hit == true {
				if bingo := boards[board_idx].CheckBingo(idx); bingo == true {
					return boards[board_idx].data_sum_ * in
				}
			}
		}
	}

	return -1
}

func remove(b []*Board, i int) []*Board {
	b[i] = b[len(b)-1]
	return b[:len(b)-1]
}

func print(b *Board) {
	for _, line := range b.data_ {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Println()
}

func Solve2() int {
	input, boards := getData("C:/Dev/Go/AdventOfCode-2021/4/data")

	for _, in := range input {
		var bingo_boards []int
		for board_idx, _ := range boards {
			if hit, idx := boards[board_idx].CheckInput(in); hit == true {
				if bingo := boards[board_idx].CheckBingo(idx); bingo == true {
					bingo_boards = append(bingo_boards, board_idx)
					if len(boards) == 1 {
						return boards[0].data_sum_ * in
					}
				}
			}
		}
		for k, _ := range bingo_boards {
			idx := len(bingo_boards) - 1 - k
			boards = remove(boards, bingo_boards[idx])
		}
	}

	return -1
}
