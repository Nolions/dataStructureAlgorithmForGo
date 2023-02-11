package linkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func initNode() {
	Head = NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
}

func TestSize(t *testing.T) {
	initNode()
	assert.Equal(t, 4, Size(Head))
}

func TestAppend(t *testing.T) {
	initNode()

	Append(2)

	assert.Equal(t, 5, Size(Head))
	assert.Equal(t, []int{30, 10, 40, 40, 2}, toSlice(Head))
}

func TestPrepend(t *testing.T) {
	initNode()

	Prepend(0)

	assert.Equal(t, 5, Size(Head))
	assert.Equal(t, []int{0, 30, 10, 40, 40}, toSlice(Head))
}

func TestReverse(t *testing.T) {
	initNode()
	reverseNode := Reverse()

	assert.Equal(t, []int{40, 40, 10, 30}, toSlice(reverseNode))
}

func TestRemoveVal(t *testing.T) {
	initNode()

	RemoveVal(30)

	assert.Equal(t, 3, Size(Head))
	assert.Equal(t, []int{10, 40, 40}, toSlice(Head))
}

func TestRemoveValCase2(t *testing.T) {
	initNode()

	RemoveVal(10)

	assert.Equal(t, 3, Size(Head))
	assert.Equal(t, []int{30, 40, 40}, toSlice(Head))
}

func TestRemoveValNoFound(t *testing.T) {
	initNode()

	RemoveVal(0)

	assert.Equal(t, 4, Size(Head))
	assert.Equal(t, []int{30, 10, 40, 40}, toSlice(Head))
}

func TestMerge(t *testing.T) {
	initNode()

	Merge(NewNode(2, NewNode(10, nil)))

	assert.Equal(t, 6, Size(Head))
	assert.Equal(t, []int{30, 10, 40, 40, 2, 10}, toSlice(Head))
}

func toSlice(node *Node) []int {
	var arr []int
	temp := node
	for temp != nil {
		arr = append(arr, temp.value)
		temp = temp.next
	}

	return arr
}
