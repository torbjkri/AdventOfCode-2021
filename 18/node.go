package dec18

import "strconv"

var nex_node_id int = 0

type Node struct {
	is_num_ bool
	value_  int

	left_   *Node
	right_  *Node
	parent_ *Node
}

func (n *Node) explode(depth int) (bool, [2]int) {

	if n.is_num_ {
		return false, [2]int{0, 0}
	}

	ok1 := n.left_.is_num_
	ok2 := n.right_.is_num_

	if ok1 && ok2 {
		if depth > 3 {

			res := [2]int{n.left_.value_, n.right_.value_}
			n.left_ = nil
			n.right_ = nil
			n.is_num_ = true
			n.value_ = 0
			return true, res
		}
	}

	if !ok1 {
		exp, res := n.left_.explode(depth + 1)
		if exp {

			if res[1] != 0 {
				if !n.right_.is_num_ {
					n.right_.add_leftmost(res[1])
				} else {
					n.right_.value_ += res[1]
				}
			}
			if res[0] != 0 {
				return true, [2]int{res[0], 0}
			}
			return true, [2]int{0, 0}
		}
	}

	if !ok2 {
		exp, res := n.right_.explode(depth + 1)
		if exp {

			if res[0] != 0 {
				if !n.left_.is_num_ {
					n.left_.add_rightmost(res[0])
				} else {
					n.left_.value_ += res[0]
				}
			}
			if res[1] != 0 {
				return true, [2]int{0, res[1]}
			}
			return true, [2]int{0, 0}

		}
	}

	return false, [2]int{0, 0}
}

func (n *Node) add_leftmost(x int) {
	if n.left_.is_num_ {
		n.left_.value_ += x
	} else {
		n.left_.add_leftmost(x)
	}
}

func (n *Node) add_rightmost(x int) {
	if n.right_.is_num_ {
		n.right_.value_ += x
	} else {
		n.right_.add_rightmost(x)
	}
}

func (n *Node) split() bool {
	if !n.is_num_ {
		sp := n.left_.split()
		if sp {
			return true
		}
		sp = n.right_.split()
		if sp {
			return true
		}

		return false
	}

	if n.value_ >= 10 {
		left := &Node{}
		left.left_ = nil
		left.right_ = nil
		left.parent_ = n
		left.is_num_ = true
		left.value_ = int(n.value_ / 2)

		right := &Node{}
		right.left_ = nil
		right.right_ = nil
		right.parent_ = n
		right.is_num_ = true
		right.value_ = int(n.value_/2) + (n.value_ % 2)

		n.is_num_ = false
		n.value_ = 0
		n.left_ = left
		n.right_ = right
		return true
	}
	return false
}

func (n *Node) GetString() string {
	if n.is_num_ {
		return strconv.Itoa(n.value_)
	} else {
		return "[" + n.left_.GetString() + "," + n.right_.GetString() + "]"
	}
}
