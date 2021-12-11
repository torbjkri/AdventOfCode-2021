package dec10

import "testing"

func TestLIFOQPush(t *testing.T) {
	q := NewLIFOQ()

	q.Push('A')

	if q.data_[0] != 'A' {
		t.Fatalf("WRONG! %d", 'A')
	}
}

func TestLIFOQPop(t *testing.T) {
	q := NewLIFOQ()

	q.Push('A')
	q.Pop()

	if !q.IsEmpty() {
		t.Fatalf("WRONG! %d", len(q.data_))
	}
}

func TestLIFOQPushPushPop(t *testing.T) {
	q := NewLIFOQ()

	q.Push('A')
	q.Push('B')
	q.Pop()

	if len(q.data_) != 1 {
		t.Fatalf("WRONG! %d", len(q.data_))
	}
	if q.data_[0] != 'A' {
		t.Fatalf("WRONG! %d", q.data_[0])
	}
}
