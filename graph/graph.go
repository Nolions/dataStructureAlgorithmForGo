package graph

import (
	"fmt"
)

type Node struct {
	Val string
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

// AddNode 增加節點
func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(h, t *Node) {
	// 第一次新增圖形的邊時需要初始化Edges map
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 互相把對方放到自己map中以確立兩個節點雙向關係，以代表兩點所形成的線
	g.Edges[*h] = append(g.Edges[*h], t)
	g.Edges[*t] = append(g.Edges[*t], h)
}

func (g *Graph) String() string {
	str := ""
	for _, node := range g.Nodes {
		str += node.string() + " -> "
		nNodes := g.Edges[*node]
		for _, nNode := range nNodes {
			str += nNode.string() + " "
		}
		str += "\n"
	}

	return str
}

func (n *Node) string() string {
	return fmt.Sprintf("%v", n.Val)
}

func (g *Graph) Traverse() []string {
	var data []string
	q := NewQueue()
	head := g.Nodes[0]
	q.Enqueue(*head)

	visited := make(map[*Node]bool)
	visited[head] = true

	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		nNodes := g.Edges[*node]
		for _, next := range nNodes {
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}

		data = append(data, node.Val)
	}

	return data
}
