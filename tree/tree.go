package tree

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
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
	data = []int{}
	if n != nil {
		// 往左走
		n.left.Postorder()

		// 印出
		data = append(data, n.value)

		// 往右走
		n.right.Postorder()
	}
}

// Preorder
// 前序遍歷
func (n *Node) Preorder() {
	data = []int{}
	if n != nil {
		// 印出
		data = append(data, n.value)

		// 往左走
		n.left.Postorder()

		// 往右走
		n.right.Postorder()
	}
}
