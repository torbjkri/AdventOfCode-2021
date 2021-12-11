package dec10

type LIFOQ struct {
	data_ []rune
}

func NewLIFOQ() *LIFOQ {
	return new(LIFOQ)
}

func (q *LIFOQ) Push(elem rune) {
	q.data_ = append(q.data_, elem)
}

func (q *LIFOQ) Pop() rune {
	res := q.data_[len(q.data_)-1]
	q.data_ = q.data_[:len(q.data_)-1]
	return res
}

func (q *LIFOQ) IsEmpty() bool {
	return len(q.data_) == 0
}
