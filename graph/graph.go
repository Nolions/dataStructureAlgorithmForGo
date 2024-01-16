package graph

import "fmt"

type Node struct {
	Val string
}

type Graph struct {
	Vertex []*Node
	Edges  map[Node][]*Node
}

// AddVertex 增加節點
func (g *Graph) AddVertex(n *Node) {
	g.Vertex = append(g.Vertex, n)
}

// AddEdge 增加線段
func (g *Graph) AddEdge(h, t *Node) {
	// 第一次新增圖形的邊時需要初始化Edges map
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 互相把對方放到自己map中以確立兩個節點雙向關係，以代表兩點所形成的線
	g.Edges[*h] = append(g.Edges[*h], t)
	g.Edges[*t] = append(g.Edges[*t], h)
}

// BFS 透過BFS遍歷
// 透過佇列方式實現
func (g *Graph) BFS() []string {
	var data []string
	q := NewQueue()
	head := g.Vertex[0]
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
			if !visited[next] {
				q.Enqueue(*next)
				visited[next] = true
			}
		}

		data = append(data, node.Val)
	}

	return data
}

// DFS 透過DFS遍歷
// 透過堆疊方式實現
func (g *Graph) DFS() []string {
	var data []string
	q := NewStack()
	head := g.Vertex[0]
	q.Push(*head)

	visited := make(map[*Node]bool)
	visited[head] = true

	for {
		if q.IsEmpty() {
			break
		}
		node := q.Pop()
		visited[node] = true
		nNodes := g.Edges[*node]
		for _, next := range nNodes {
			if !visited[next] {
				q.Push(*next)
				visited[next] = true
			}
		}

		data = append(data, node.Val)
	}

	return data
}

// String Graph format to string
func (g *Graph) String() string {
	str := ""
	for _, node := range g.Vertex {
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
