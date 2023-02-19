package array

import (
	"github.com/Nolions/dataStructureAlgorithmForGo/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s *Stack
var l []int

//var a

func init() {
	s = New(8)
	l = []int{1, 2, 3, 4, 5, 6}
}

func TestPush(t *testing.T) {
	for _, v := range l {
		s.Push(v)
	}

	assert.Equal(t, []stack.Data{1, 2, 3, 4, 5, 6}, s.list)
}

func TestPop(t *testing.T) {
	for _, v := range l {
		s.Push(v)
	}

	for i := len(s.list); i > 0; i-- {
		assert.Equal(t, l[i-1], s.Pop())
	}
}
