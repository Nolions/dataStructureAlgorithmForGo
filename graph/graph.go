package graph

import "fmt"

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
		str += node.String() + " -> "
		nNodes := g.Edges[*node]
		for _, nNode := range nNodes {
			str += nNode.String() + " "
		}
		str += "\n"
	}

	return str
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Val)
}
