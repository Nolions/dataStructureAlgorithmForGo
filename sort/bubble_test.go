package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{8, 4, 3, 0, 5, 7, 6}

	actual := bubbleSort(data)
	expected := []int{0, 3, 4, 5, 6, 7, 8}

	assert.Equal(t, expected, actual)
}
