package dec10

type IoTypes struct {
	open_brackets_          []rune
	close_brackets_         []rune
	close_to_open_          map[rune]rune
	close_bracket_to_score_ map[rune]int
	open_bracket_to_score_  map[rune]int
}

func runeSliceConains(s []rune, e rune) bool {
	for _, elem := range s {
		if elem == e {
			return true
		}
	}
	return false
}

func CreateIoTypes() *IoTypes {
	io := new(IoTypes)

	io.open_brackets_ = []rune{'(', '{', '<', '['}
	io.close_brackets_ = []rune{')', '}', '>', ']'}
	io.close_to_open_ = map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
		'>': '<',
	}

	io.close_bracket_to_score_ = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	io.open_bracket_to_score_ = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	return io
}

func CheckCorruptedLineScore(data string) int {

	io := CreateIoTypes()

	lifo := NewLIFOQ()

	for _, elem := range data {
		if runeSliceConains(io.open_brackets_, elem) {
			lifo.Push(elem)
		} else if runeSliceConains(io.close_brackets_, elem) {
			last_pushed := lifo.Pop()
			open_elem := io.close_to_open_[elem]

			if last_pushed != open_elem {
				return io.close_bracket_to_score_[elem]
			}
		}
	}

	return 0
}

func CheckIncompleteLineScore(data string) int {

	io := CreateIoTypes()

	lifo := NewLIFOQ()

	for _, elem := range data {
		if runeSliceConains(io.open_brackets_, elem) {
			lifo.Push(elem)
		} else {
			last_pushed := lifo.Pop()
			open_elem := io.close_to_open_[elem]

			if last_pushed != open_elem {
				return 0
			}
		}
	}
	sum := 0
	for !lifo.IsEmpty() {
		sum *= 5
		sum += io.open_bracket_to_score_[lifo.Pop()]
	}

	return sum
}
