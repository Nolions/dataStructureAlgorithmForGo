package binaryTree

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

func (n *Node) AppendLeftNode(node *Node) {
	n.left = node
}

func (n *Node) AppendRightNode(node *Node) {
	n.right = node
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
