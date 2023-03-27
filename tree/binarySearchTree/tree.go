package binarySearchTree

type Node struct {
	value int
	left  *Node
	right *Node
}

func node(value int) *Node {
	var n Node
	n.value = value
	return &n
}

// Insert
// 新增樹節點
func (n *Node) Insert(v int) {
	if v < n.value {
		if n.left == nil {
			n.left = node(v)
		} else {
			n.left.Insert(v)
		}
	} else {
		if n.right == nil {
			n.right = node(v)
		} else {
			n.right.Insert(v)
		}
	}
}

func (n *Node) Find(i int) *Node {
	if n == nil {
		return nil
	}

	p := n
	if p.value == i {
		return p
	} else if p.value < i {
		return p.right.Find(i)
	} else {
		return p.left.Find(i)
	}
}

var data []int

// Postorder
// 後序遍歷
func (n *Node) Postorder() {
	if n != nil {
		// 往左走
		n.left.Postorder()
		// 往右走
		n.right.Postorder()
		// 印出

		data = append(data, n.value)
	}
}

// Inorder
// 中序遍歷
func (n *Node) Inorder() {
	if n != nil {
		// 往左走
		n.left.Inorder()
		// 印出
		data = append(data, n.value)
		// 往右走
		n.right.Inorder()
	}
}

// Preorder
// 前序遍歷
func (n *Node) Preorder() {
	if n != nil {
		// 印出
		data = append(data, n.value)
		// 往左走
		n.left.Preorder()
		// 往右走
		n.right.Preorder()
	}
}
