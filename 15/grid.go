package dec15

import (
	"fmt"
	"strconv"
	"strings"
)

type Grid struct {
	finished_  NodeSlice
	visited_   NodeSlice
	unvisited_ NodeSlice

	start_ Position
	end_   Position
}

func (g *Grid) FindShortestPath() int {

	currentNode := g.visited_.data_[0]

	for !g.end_.Equal(&currentNode.pos_) {

		for i, _ := range currentNode.children_ {
			var child *Node = nil
			childpos := currentNode.children_[i]
			if child = g.finished_.Find(childpos); child != nil {
				fmt.Println("Children not properly reset")
				return -1
			} else if child = g.visited_.Find(childpos); child != nil {
				oldVal := child.total_risk_
				child.Visit(currentNode)
				if child.total_risk_ < oldVal {
					g.visited_.Update(child)
				}

			} else if child = g.unvisited_.Find(childpos); child != nil {
				child.Visit(currentNode)
				g.visited_.AddSorted(child)

				g.unvisited_.RemoveNode(child)

			} else {
				fmt.Println("Unable to locate node!!")
				return -1
			}
		}
		g.finished_.Push(currentNode)

		currentNode = g.visited_.Pop()

	}

	//fmt.Println("Backtrace")
	//for !g.start_.Equal(&currentNode.pos_) {
	//	fmt.Printf("Pos : %v, Risk: %v\n", currentNode.pos_, currentNode.risk_)
	//	currentNode = currentNode.shorted_parent_
	//}

	return currentNode.naked_risk_
}

func NewGrid(data []string) (g *Grid) {
	g = new(Grid)

	g.finished_ = NodeSlice{[]*Node{}}
	g.unvisited_ = NodeSlice{[]*Node{}}
	g.visited_ = NodeSlice{[]*Node{}}

	g.start_ = Pos(0, 0)

	max := len(data)
	may := len(data[0])

	g.end_ = Pos(max-1, may-1)

	for j, line := range data {
		line_vec := strings.Split(line, "")
		for i, elem := range line_vec {
			num, _ := strconv.Atoi(elem)
			if i == 0 && j == 0 {
				n := NewNode(Pos(0, 0), Pos(max, may), num)
				n.total_risk_ = 0
				n.naked_risk_ = 0
				g.visited_.Push(n)
			} else {
				g.unvisited_.Push(NewNode(Pos(i, j), Pos(max, may), num))
			}
		}
	}

	return g
}

func NewGrid5x(data []string) (g *Grid) {
	g = new(Grid)

	g.finished_ = NodeSlice{[]*Node{}}
	g.unvisited_ = NodeSlice{[]*Node{}}
	g.visited_ = NodeSlice{[]*Node{}}

	g.start_ = Pos(0, 0)

	max := len(data) * 5
	may := len(data[0]) * 5

	g.end_ = Pos(max-1, may-1)

	for extray := 0; extray < 5; extray++ {
		for j, line := range data {
			line_vec := strings.Split(line, "")

			for extrax := 0; extrax < 5; extrax++ {
				for i, elem := range line_vec {
					num, _ := strconv.Atoi(elem)
					num = num + extrax + extray
					if num > 9 {
						num = num - 9
					}
					if i == 0 && j == 0 && extrax == 0 && extray == 0 {
						n := NewNode(Pos(0, 0), Pos(max, may), num)
						n.total_risk_ = 0
						n.naked_risk_ = 0
						g.visited_.Push(n)
					} else {
						g.unvisited_.Push(NewNode(Pos(i+extrax*len(data), j+extray*len(data[0])), Pos(max, may), num))
					}
				}
			}
		}
	}

	return g
}
