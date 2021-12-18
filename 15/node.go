package dec15

import (
	"fmt"
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

	weight_     int
	naked_risk_ int
}

func NewNode(p Position, maxP Position, r int) (n *Node) {
	n = new(Node)
	n.pos_ = p
	n.risk_ = r
	n.total_risk_ = math.MaxInt
	n.naked_risk_ = math.MaxInt

	n.weight_ = maxP.x_ + maxP.y_ - p.x_ - p.y_

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
	tot := visitor.naked_risk_ + n.risk_ + n.weight_

	if tot < n.total_risk_ {
		n.total_risk_ = tot
		n.naked_risk_ = visitor.naked_risk_ + n.risk_
		n.shorted_parent_ = visitor
	}

	n.children_ = RemovePosition(n.children_, &visitor.pos_)
}

type NodeSlice struct {
	data_ []*Node
}

func (ns *NodeSlice) RemoveNode(n *Node) {
	for i, _ := range ns.data_ {
		if ns.data_[i].pos_.Equal(&n.pos_) {
			ns.data_[i] = ns.data_[len(ns.data_)-1]

			ns.data_ = ns.data_[:len(ns.data_)-1]
			return
		}
	}
	return
}

func (ns *NodeSlice) Find(p Position) (o *Node) {
	for i, _ := range ns.data_ {
		if ns.data_[i].pos_.Equal(&p) {
			return ns.data_[i]
		}
	}
	return nil
}

func (ns *NodeSlice) Push(n *Node) {
	ns.data_ = append(ns.data_, n)
}

func (ns *NodeSlice) Pop() *Node {
	last := ns.data_[len(ns.data_)-1]
	ns.data_ = ns.data_[:len(ns.data_)-1]
	return last
}

func (ns *NodeSlice) Sort() {
	sort.Slice(ns.data_, func(i, j int) bool {
		return ns.data_[i].total_risk_ > ns.data_[j].total_risk_
	})
}

func (ns *NodeSlice) AddSorted(n *Node) {
	if len(ns.data_) == 0 {
		ns.data_ = append(ns.data_, n)
		return
	}
	for i := len(ns.data_) - 1; i >= 0; i-- {
		if ns.data_[i].total_risk_ > n.total_risk_ {
			ns.data_ = append(ns.data_[:i+1], ns.data_[i:]...)
			ns.data_[i] = n
			return
		}
	}
	ns.data_ = append([]*Node{n}, ns.data_...)
}

func (ns *NodeSlice) Update(n *Node) {

	newPos := -1
	for i := len(ns.data_) - 1; i >= 0; i-- {
		if newPos == -1 {
			if ns.data_[i].total_risk_ > n.total_risk_ {
				newPos = i
			}
		}
		if ns.data_[i].pos_.Equal(&n.pos_) {
			if newPos != i {
				fmt.Printf("%v, %v, %v\n", i, newPos, len(ns.data_))
				temp := append(ns.data_[i+1:newPos:newPos-i], ns.data_[newPos-1:]...)
				ns.data_ = append(ns.data_[:i], temp...)
				ns.data_[newPos] = n
				return
			}
		}
	}
}
