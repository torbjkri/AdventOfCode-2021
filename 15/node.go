package dec15

import (
	"math"
	"sort"
)

type Position struct {
	x_ int
	y_ int
}

func (p *Position) Equal(v *Position) bool {
	return p.x_ == v.x_ && p.y_ == v.y_
}

func RemovePosition(ps []Position, p *Position) []Position {
	for i, _ := range ps {
		if ps[i].Equal(p) {
			ps[i] = ps[len(ps)-1]
			return ps[:len(ps)-1]
		}
	}

	return ps
}

func Pos(x, y int) Position {
	return Position{x, y}
}

type Node struct {
	pos_ Position

	risk_       int
	total_risk_ int

	children_ []Position

	shorted_parent_ *Node
}

func NewNode(p Position, maxP Position, r int) (n *Node) {
	n = new(Node)
	n.pos_ = p
	n.risk_ = r
	n.total_risk_ = math.MaxInt

	if p.x_ > 0 {
		n.children_ = append(n.children_, Pos(p.x_-1, p.y_))
	}
	if p.x_ < maxP.x_-1 {
		n.children_ = append(n.children_, Pos(p.x_+1, p.y_))
	}
	if p.y_ > 0 {
		n.children_ = append(n.children_, Pos(p.x_, p.y_-1))
	}
	if p.y_ < maxP.y_-1 {
		n.children_ = append(n.children_, Pos(p.x_, p.y_+1))
	}

	n.shorted_parent_ = nil

	return n
}

func (n *Node) Visit(visitor *Node) {
	tot := visitor.total_risk_ + n.risk_

	if tot < n.total_risk_ {
		n.total_risk_ = tot
		n.shorted_parent_ = visitor
	}

	n.children_ = RemovePosition(n.children_, &visitor.pos_)
}

func Sort(ns []*Node) []*Node {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].total_risk_ > ns[j].total_risk_
	})

	return ns
}

func Pop(ns []*Node) ([]*Node, *Node) {
	last := ns[len(ns)-1]
	return ns[:len(ns)-1], last
}

func Push(ns []*Node, n *Node) []*Node {
	return append(ns, n)
}

func Find(ns []*Node, p Position) (o *Node) {
	for i, _ := range ns {
		if ns[i].pos_.Equal(&p) {
			return ns[i]
		}
	}
	return nil
}

func RemoveNode(ns []*Node, n *Node) []*Node {
	for i, _ := range ns {
		if ns[i].pos_.Equal(&n.pos_) {
			ns[i] = ns[len(ns)-1]
			return ns[:len(ns)-1]
		}
	}
	return ns
}
