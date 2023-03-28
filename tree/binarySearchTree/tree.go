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

// Find
// 尋找節點
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

// DeleteNode
// 刪除節點
// 需要去判斷要刪除的節點是在根節點或是左/右節點上
func (n *Node) DeleteNode(i int) *Node {

	if n == nil {
		return nil
	}

	p := n
	// 刪除的節點為根節點
	if p.value == i {
		// 只有右節點
		if p.left == nil {
			p = p.right
			return p
		}

		// 只有左節點
		if p.right == nil {
			p = p.left
			return p
		}

		// 同時有左右節點
		// 找到根節點的右子樹中最左邊的樹葉節點
		c := p.right
		for c.left != nil {
			c = c.left
		}

		// 把根節點的左子樹接到左子樹中最左邊
		c.left = p.left

		// 根節點替代為最初根節點的右側節點
		p = p.right
		return p
	}

	// 節點在右節點
	if p.value < i {
		p.right = p.right.DeleteNode(i)
	}

	// 節點在左節點
	if p.value > i {
		p.left = p.left.DeleteNode(i)
	}

	return p
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
